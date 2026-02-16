package processer

import (
	"context"
	"fmt"
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
			usage, _ := monitor.CheckUsage(ctx)
			statCh <- model.SystemStat{Name: monitor.GetName(), Value: usage}
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

	return output.String()
}
