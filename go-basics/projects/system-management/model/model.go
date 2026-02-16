package model

import (
	"context"
	"fmt"
	"sync"
)

type Monitor interface {
	GetName() string
	CheckUsage(ctx context.Context) (string, bool)
}

type SystemStat struct {
	Name  string
	Value string
}

func (s SystemStat) String() string {
	return fmt.Sprintf("[%s]: %v\n", s.Name, s.Value)
}

type ProcessStat struct {
	ID          int32
	Name        string
	CPUPercent  float64
	MemPercent  float32
	RunningTime float32
}

func (s ProcessStat) String() string {
	return fmt.Sprintf("ID: %d, Name: %s, CPU: %.2f%%, Memory: %.2f%%, Running Time: %.2f seconds",
		s.ID, s.Name, s.CPUPercent, s.MemPercent, s.RunningTime)
}

var (
	StartMutex sync.Mutex
	StatMap    = make(map[string]SystemStat)
)
