package processer

import (
	"context"
	"fmt"
	"sync"
	"time"

	"example.com/go/model"
	"example.com/go/monitor/cpu"
	"example.com/go/monitor/mem"
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
