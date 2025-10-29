package main

import (
	"custom-autoscaler/internal/app"
	"custom-autoscaler/internal/autoscaler"
	"custom-autoscaler/internal/cloud"
	"custom-autoscaler/internal/metrics"
	"fmt"
	"log"
	"time"

	"github.com/viebiz/lit/env"
)

func main() {
	cfg, err := env.ReadAppConfig[app.Config]()
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	fmt.Println("✅ Config loaded:")
	fmt.Printf("  Provider: %s\n", cfg.Autoscaler.Provider)
	fmt.Printf("  Interval: %ds\n", cfg.Autoscaler.IntervalSeconds)
	fmt.Printf("  Cooldown: %ds\n", cfg.Autoscaler.CooldownSeconds)
	fmt.Printf("  Prometheus port: %d\n", cfg.Autoscaler.PrometheusPort)

	var provider cloud.Provider
	switch cfg.Autoscaler.Provider {
	case "mock":
		provider = cloud.NewMockCloudProvider()
	default:
		log.Fatalf("Unknown provider: %s", cfg.Autoscaler.Provider)
	}

	engine := autoscaler.DecisionEngine{
		ScaleUpCPU:        cfg.Autoscaler.ScaleUpCPU,
		ScaleDownCPU:      cfg.Autoscaler.ScaleDownCPU,
		ScaleUpMem:        cfg.Autoscaler.ScaleUpMemory,
		ScaleDownMem:      cfg.Autoscaler.ScaleDownMemory,
		ScaleUpRespTime:   cfg.Autoscaler.ScaleUpRespTime,
		ScaleDownRespTime: cfg.Autoscaler.ScaleDownRespTime,
		MinNodes:          cfg.Autoscaler.MinNodes,
		MaxNodes:          cfg.Autoscaler.MaxNodes,
	}

	cooldown := autoscaler.NewCooldownManager(cfg.Autoscaler.CooldownSeconds)
	metrics.InitPrometheus(cfg.Autoscaler.PrometheusPort)

	interval := time.Duration(cfg.Autoscaler.IntervalSeconds) * time.Second
	autoscaler.StartScheduler(provider, engine, interval, cooldown)
}
