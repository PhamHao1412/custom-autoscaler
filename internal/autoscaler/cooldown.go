package autoscaler

import (
	"sync"
	"time"
)

type CooldownManager struct {
	lastActionTime time.Time
	mu             sync.Mutex
	duration       time.Duration
}

func NewCooldownManager(seconds int) *CooldownManager {
	return &CooldownManager{
		lastActionTime: time.Now().Add(-time.Hour),
		duration:       time.Duration(seconds) * time.Second,
	}
}

func (c *CooldownManager) CanScale() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return time.Since(c.lastActionTime) >= c.duration
}

func (c *CooldownManager) RecordAction() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.lastActionTime = time.Now()
}
