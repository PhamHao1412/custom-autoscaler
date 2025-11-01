package metrics

import (
	"log"
	"math/rand"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

type Metrics struct {
	CPUUsage       float64
	MemoryUsage    float64
	ResponseTimeMS float64
}

func GetCurrentMetrics() Metrics {
	rand.Seed(time.Now().UnixNano())

	return Metrics{
		CPUUsage:       20 + rand.Float64()*80,  // 20–100%
		MemoryUsage:    30 + rand.Float64()*60,  // 30–90%
		ResponseTimeMS: 50 + rand.Float64()*450, // 50–500 ms
	}
}

func GetCurrentMacMetrics() Metrics {
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Printf("⚠️ Failed to read CPU usage: %v", err)
		return Metrics{}
	}

	memStat, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("⚠️ Failed to read memory usage: %v", err)
		return Metrics{}
	}

	// Simulate request latency (ms)
	responseTime := 50 + rand.Float64()*850

	cpuUsage := 0.0
	if len(cpuPercent) > 0 {
		cpuUsage = cpuPercent[0]
	}

	log.Printf("[METRIC] CPU=%.2f%% | Mem=%.2f%% | RespTime=%.2fms\n",
		cpuUsage, memStat.UsedPercent, responseTime)

	return Metrics{
		CPUUsage:       cpuUsage,
		MemoryUsage:    memStat.UsedPercent,
		ResponseTimeMS: responseTime,
	}
}
