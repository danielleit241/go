package cpu

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

type CPUMonitor struct {
	Interval time.Duration
}

func NewCPUMonitor(interval time.Duration) *CPUMonitor {
	return &CPUMonitor{Interval: interval}
}

func (m *CPUMonitor) GetName() string {
	return "CPU"
}

func (m *CPUMonitor) CheckUsage(ctx context.Context) (string, bool) {
	percent, err := cpu.PercentWithContext(ctx, m.Interval, false)
	if err != nil {
		value := fmt.Sprintf("[CPU Monitor] Could not retrieve CPU usage: %v \n", err)
		return value, false
	}

	if len(percent) > 0 {
		value := fmt.Sprintf("%.2f%%", percent[0])
		return value, percent[0] > 60
	}
	return "0.00", false
}
