package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"example.com/go/model"
	"example.com/go/monitor/cpu"
	"example.com/go/monitor/disk"
	"example.com/go/monitor/mem"
	"example.com/go/monitor/net"
	"example.com/go/processer"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	timeInterval := 1000 * time.Millisecond
	monitorList := []model.Monitor{
		&cpu.CPUMonitor{Interval: timeInterval},
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
				fmt.Printf("[%s]: %v\n", stat.Name, stat.Value)
			}
			fmt.Println("-----")
			model.StartMutex.Unlock()
		}
	}()

	wg.Wait()

	cancel()

	close(statCh)
}
