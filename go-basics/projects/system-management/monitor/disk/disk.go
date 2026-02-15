package disk

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/disk"
)

type DiskMonitor struct {
	Interval time.Duration
}

func NewDiskMonitor(interval time.Duration) *DiskMonitor {
	return &DiskMonitor{Interval: interval}
}

func (m *DiskMonitor) GetName() string {
	return "Disk"
}

func (m *DiskMonitor) CheckUsage(ctx context.Context) (string, bool) {
	path := "/"
	if runtime.GOOS == "windows" {
		path = "C:"
	}

	diskStat, err := disk.UsageWithContext(ctx, path)
	if err != nil {
		strErr := fmt.Sprintf("[Disk Monitor] Could not retrieve Disk info: %v \n", err)
		return strErr, false
	}

	value := fmt.Sprintf("%.2f%% used", diskStat.UsedPercent)

	return value, diskStat.UsedPercent > 60
}
