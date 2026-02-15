package processer

import (
	"context"
	"fmt"
	"sync"
	"time"

	"example.com/go/monitor/cpu"
	"example.com/go/monitor/mem"
)

type Processer interface {
	cpu.CPUMonitor | mem.MemMonitor
}

func RunMonitoring(ctx context.Context, wg *sync.WaitGroup) {

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
			cpu := cpu.NewCPUMonitor(timeInterval)
			usage, _ := cpu.CheckCPUUsage(ctx)
			fmt.Printf("%s: %s\n", cpu.GetName(), usage)
			// if isHigh {
			// 	fmt.Println("Warning: CPU usage is high!")
			// }

			mem := mem.NewMemMonitor(timeInterval)
			memUsage, _ := mem.CheckMemUsage(ctx)
			fmt.Printf("%s: %s\n", mem.GetName(), memUsage)
			// if memIsHigh {
			// 	fmt.Println("Warning: Memory usage is high!")
			// }
		}
	}
}
