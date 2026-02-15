package model

import (
	"context"
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

var (
	StartMutex sync.Mutex
	StatMap    = make(map[string]SystemStat)
)
