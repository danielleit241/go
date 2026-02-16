package processer

import (
	"context"
	"fmt"
	"sort"
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

func GetTopProcesses(ctx context.Context) string {
	processes, err := process.ProcessesWithContext(ctx)
	if err != nil {
		return fmt.Sprintf("[Get Top Processes] Could not retrieve process list: %v \n", err)
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var cpuList, memList []model.ProcessStat

	procChan := make(chan model.ProcessStat, len(processes))

	for _, proc := range processes {
		wg.Add(1)
		go func(proc *process.Process) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				name, err := proc.NameWithContext(ctx)
				if err != nil {
					return
				}
				cpuPercent, err := proc.CPUPercentWithContext(ctx)
				if err != nil {
					return
				}
				memPercent, err := proc.MemoryPercentWithContext(ctx)
				if err != nil {
					return
				}

				createTime, err := proc.CreateTimeWithContext(ctx)
				if err != nil {
					return
				}

				// Time Unix trả về thời gian tính bằng milliseconds kể từ 1/1/1970, nên chia cho 1000 để chuyển sang seconds
				runningTime := time.Since(time.Unix(createTime/1000, 0))

				if cpuPercent > 5.0 || memPercent > 5.0 {
					mu.Lock()
					defer mu.Unlock()
					var procStat = model.ProcessStat{
						ID:          proc.Pid,
						Name:        name,
						CPUPercent:  cpuPercent,
						MemPercent:  memPercent,
						RunningTime: float32(runningTime.Seconds()),
					}
					procChan <- procStat
				}
			}
		}(proc)
	}

	go func() {
		wg.Wait()
		close(procChan)
	}()

	for procStat := range procChan {
		if procStat.CPUPercent > 5.0 {
			cpuList = append(cpuList, procStat)
		}

		if procStat.MemPercent > 1.0 {
			memList = append(memList, procStat)
		}
	}

	sort.Slice(cpuList, func(i, j int) bool {
		return cpuList[i].CPUPercent > cpuList[j].CPUPercent
	})

	sort.Slice(memList, func(i, j int) bool {
		return memList[i].MemPercent > memList[j].MemPercent
	})

	output := "Top 5 CPU-consuming processes:\n"
	for i, proc := range cpuList {
		if i >= 5 {
			break
		}
		output += fmt.Sprintf("PID: %d, Name: %s, CPU: %.2f%%, Mem: %.2f%%, Running Time: %.2fs\n",
			proc.ID, proc.Name, proc.CPUPercent, proc.MemPercent, proc.RunningTime)
	}

	output += "\nTop 5 Memory-consuming processes:\n"
	for i, proc := range memList {
		if i >= 5 {
			break
		}
		output += fmt.Sprintf("PID: %d, Name: %s, CPU: %.2f%%, Mem: %.2f%%, Running Time: %.2fs\n",
			proc.ID, proc.Name, proc.CPUPercent, proc.MemPercent, proc.RunningTime)
	}

	return output
}
