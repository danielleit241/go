package net

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/net"
)

type NetMonitor struct {
	Interval time.Duration
}

func NewNetMonitor(interval time.Duration) *NetMonitor {
	return &NetMonitor{Interval: interval}
}

func (m *NetMonitor) GetName() string {
	return "Network"
}

func (m *NetMonitor) CheckUsage(ctx context.Context) (string, bool) {
	netIO, err := net.IOCountersWithContext(ctx, false)
	if err != nil {
		strErr := fmt.Sprintf("[Network Monitor] Could not retrieve Network info: %v \n", err)
		return strErr, false
	}
	if len(netIO) > 0 {
		value := fmt.Sprintf("Send: %d KB, Recv: %d KB", netIO[0].BytesSent/1024, netIO[0].BytesRecv/1024)
		return value, false
	}
	return "N/A", false
}
