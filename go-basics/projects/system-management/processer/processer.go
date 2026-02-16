package processer

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"example.com/go/model"
	"example.com/go/monitor/cpu"
	"example.com/go/monitor/mem"
	"github.com/shirou/gopsutil/process"
)

type Processer interface {
	cpu.CPUMonitor | mem.MemMonitor
}

func RunMonitoring(ctx context.Context, wg *sync.WaitGroup, statCh chan<- model.SystemStat, monitor model.Monitor) {

	defer wg.Done()
	timeInterval := 1000 * time.Millisecond
	ticker := time.NewTicker(timeInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Monitoring stopped.")
			return
		case <-ticker.C:
			usage, alert := monitor.CheckUsage(ctx)

			stat := model.SystemStat{Name: monitor.GetName(), Value: usage}

			if alert {
				LogAlerts(stat)
			}
			statCh <- stat
		}
	}
}

func GetTopProcesses(ctx context.Context, cpuMaxPercent float64, memMaxPercent float32) string {
	processes, err := process.ProcessesWithContext(ctx)
	if err != nil {
		return fmt.Sprintf("[Get Top Processes] Could not retrieve process list: %v \n", err)
	}

	var wg sync.WaitGroup
	var cpuList, memList []model.ProcessStat

	procChan := make(chan model.ProcessStat, len(processes))

	const processPerGoroutine = 10
	for i := 0; i < len(processes); i += processPerGoroutine {
		end := i + processPerGoroutine
		if end > len(processes) {
			end = len(processes)
		}
		batch := processes[i:end]
		wg.Add(1)
		go func(batch []*process.Process) {
			defer wg.Done()
			for _, proc := range batch {
				select {
				case <-ctx.Done():
					return
				default:
				}

				name, err := proc.NameWithContext(ctx)
				if err != nil {
					continue
				}
				cpuPercent, err := proc.CPUPercentWithContext(ctx)
				if err != nil {
					continue
				}
				memPercent, err := proc.MemoryPercentWithContext(ctx)
				if err != nil {
					continue
				}

				createTime, err := proc.CreateTimeWithContext(ctx)
				if err != nil {
					continue
				}

				if cpuPercent < cpuMaxPercent && memPercent < memMaxPercent {
					continue
				}

				// Time Unix trả về thời gian tính bằng milliseconds kể từ 1/1/1970, nên chia cho 1000 để chuyển sang seconds
				runningTime := time.Since(time.Unix(createTime/1000, 0))

				procChan <- model.ProcessStat{
					ID:          proc.Pid,
					Name:        name,
					CPUPercent:  cpuPercent,
					MemPercent:  memPercent,
					RunningTime: float32(runningTime.Seconds()),
				}
			}
		}(batch)
	}

	go func() {
		wg.Wait()
		close(procChan)
	}()

	for procStat := range procChan {
		if procStat.CPUPercent > cpuMaxPercent {
			cpuList = append(cpuList, procStat)
		}

		if procStat.MemPercent > memMaxPercent {
			memList = append(memList, procStat)
		}
	}

	sort.Slice(cpuList, func(i, j int) bool {
		return cpuList[i].CPUPercent > cpuList[j].CPUPercent
	})

	sort.Slice(memList, func(i, j int) bool {
		return memList[i].MemPercent > memList[j].MemPercent
	})

	var output strings.Builder
	output.WriteString("Top 5 CPU-consuming processes:\n")
	for i, proc := range cpuList {
		if i >= 5 {
			break
		}
		fmt.Fprintf(&output, "%d [PID: %d], Name: %s, CPU: %.2f%%, Mem: %.2f%%, Running Time: %.2fs\n",
			i+1, proc.ID, proc.Name, proc.CPUPercent, proc.MemPercent, proc.RunningTime)
	}

	output.WriteString("\nTop 5 Memory-consuming processes:\n")
	for i, proc := range memList {
		if i >= 5 {
			break
		}
		fmt.Fprintf(&output, "%d [PID: %d], Name: %s, CPU: %.2f%%, Mem: %.2f%%, Running Time: %.2fs\n",
			i+1, proc.ID, proc.Name, proc.CPUPercent, proc.MemPercent, proc.RunningTime)
	}

	ExportTopProcessesToCSV(cpuList, memList)

	return output.String()
}

func ExportTopProcessesToCSV(cpuList, memList []model.ProcessStat) string {
	// O_CREATE: Tạo file mới nếu chưa tồn tại, hoặc xóa nội dung cũ nếu đã tồn tại
	// O_WRONLY: Mở file ở chế độ chỉ ghi
	// O_APPEND: Ghi thêm vào cuối file thay vì ghi đè
	// 0644: Quyền truy cập file (owner có quyền đọc và ghi, group và others chỉ có quyền đọc)
	file, err := os.OpenFile("top_process_stats.csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Sprintf("[Export to CSV] Could not open file: %v \n", err)
	}

	defer file.Close()

	if stat, err := file.Stat(); err == nil && stat.Size() == 0 {
		// Ghi header nếu file mới tạo hoặc đã bị xóa nội dung
		file.WriteString("Timestamp,PID,Name,CPU (%),Memory (%),Running Time\n")
	}

	timestamp := time.Now().Format(time.RFC3339) // Định dạng thời gian theo chuẩn ISO 8601

	for _, proc := range cpuList {
		line := fmt.Sprintf("%s,%d,%s,%.2f,%.2f,%.2f\n",
			timestamp, proc.ID, proc.Name, proc.CPUPercent, proc.MemPercent, proc.RunningTime)
		file.WriteString(line)
	}

	for _, proc := range memList {
		line := fmt.Sprintf("%s,%d,%s,%.2f,%.2f,%.2f\n",
			timestamp, proc.ID, proc.Name, proc.CPUPercent, proc.MemPercent, proc.RunningTime)
		file.WriteString(line)
	}

	return "Exported top processes to top_process_stats.csv"
}

func LogAlerts(stat model.SystemStat) {
	file, err := os.OpenFile("alerts.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("[Log Alerts] Could not open log file: %v \n", err)
		return
	}
	defer file.Close()

	logEntry := fmt.Sprintf("%s - ALERT: %s usage is high: %s\n", time.Now().Format(time.RFC3339), stat.Name, stat.Value)
	if _, err := file.WriteString(logEntry); err != nil {
		fmt.Printf("[Log Alerts] Could not write to log file: %v \n", err)
	}
}
