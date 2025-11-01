package main

import (
	"custom-autoscaler/internal/app"
	"custom-autoscaler/internal/autoscaler"
	"custom-autoscaler/internal/cloud"
	"custom-autoscaler/internal/logging"
	"custom-autoscaler/internal/metrics"
	"log"
	"time"

	"github.com/viebiz/lit/env"
)

func main() {
	logging.InitLogger("logs/autoscaler.log")
	cfg, err := env.ReadAppConfig[app.Config]()
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	log.Println("✅ Config loaded:")
	log.Printf("  Provider: %s\n", cfg.Autoscaler.Provider)
	log.Printf("  Interval: %ds\n", cfg.Autoscaler.IntervalSeconds)
	log.Printf("  Cooldown: %ds\n", cfg.Autoscaler.CooldownSeconds)
	log.Printf("  Prometheus port: %d\n", cfg.Autoscaler.PrometheusPort)
	log.Println("🚀 Autoscaler initialized and ready.")
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
