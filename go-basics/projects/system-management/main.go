package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/danielleit241/model"
	"github.com/danielleit241/monitor/cpu"
	"github.com/danielleit241/monitor/disk"
	"github.com/danielleit241/monitor/mem"
	"github.com/danielleit241/monitor/net"
	"github.com/danielleit241/processer"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	timeInterval := 1000 * time.Millisecond
	cpuMaxPercent := 5.0
	memMaxPercent := float32(1.0)

	monitorList := []model.Monitor{
		&cpu.CPUMonitor{Interval: timeInterval}, // new CPUMonitor()
		&mem.MemMonitor{Interval: timeInterval},
		&net.NetMonitor{Interval: timeInterval},
		&disk.DiskMonitor{Interval: timeInterval},
	}

	statCh := make(chan model.SystemStat)

	var wg sync.WaitGroup

	for _, monitor := range monitorList {
		wg.Add(1)
		go processer.RunMonitoring(ctx, &wg, statCh, monitor)
	}

	go func() {
		for stat := range statCh {
			model.StartMutex.Lock()
			model.StatMap[stat.Name] = stat
			model.StartMutex.Unlock()
		}
	}()

	go func() {
		for {
			time.Sleep(5 * time.Second)
			model.StartMutex.Lock()
			for _, stat := range model.StatMap {
				fmt.Print(stat)
			}
			fmt.Println("-----")
			model.StartMutex.Unlock()

			topProcesses := processer.GetTopProcesses(ctx, cpuMaxPercent, memMaxPercent)
			fmt.Println(topProcesses)
		}
	}()

	wg.Wait()

	cancel()

	close(statCh)
}
