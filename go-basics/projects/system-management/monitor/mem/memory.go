package mem

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/mem"
)

type MemMonitor struct {
	Interval time.Duration
}

func NewMemMonitor(interval time.Duration) *MemMonitor {
	return &MemMonitor{Interval: interval}
}

func (m *MemMonitor) GetName() string {
	return "Memory"
}

func (m *MemMonitor) CheckMemUsage(ctx context.Context) (string, bool) {
	vmStat, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		strErr := fmt.Sprintf("[Memory Monitor] Could not retrieve Memory info: %v \n", err)
		return strErr, false
	}
	value := fmt.Sprintf("%.2f%%", vmStat.UsedPercent)
	return value, vmStat.UsedPercent > 60
}
