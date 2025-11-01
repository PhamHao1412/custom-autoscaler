package autoscaler

import (
	"custom-autoscaler/internal/cloud"
	"custom-autoscaler/internal/metrics"
	"fmt"
	"log"
	"time"
)

func StartScheduler(provider cloud.Provider, engine DecisionEngine, interval time.Duration, cooldown *CooldownManager) {
	fmt.Println("🚀 Autoscaler started...")

	for {
		m := metrics.GetCurrentMacMetrics()
		nodes, _ := provider.ListNodes()

		metrics.CPUUsageGauge.Set(m.CPUUsage)
		metrics.NodeCountGauge.Set(float64(len(nodes)))
		metrics.MemoryUsageGauge.Set(m.MemoryUsage)
		metrics.ResponseTimeGauge.Set(m.ResponseTimeMS)
		action, reason := engine.Evaluate(m, len(nodes))
		log.Printf("[AUTOSCALER] CPU=%.2f%% | Nodes=%d | Action=%s | Reason=%s\n",
			m.CPUUsage, len(nodes), action, reason)

		if !cooldown.CanScale() && action != "no-op" {
			log.Println("⏳ In cooldown period, skipping scale action...")
		} else {
			switch action {
			case "scale-up":
				name := fmt.Sprintf("node-%d", len(nodes)+1)
				_ = provider.AddNode(name)
				metrics.LastActionGauge.WithLabelValues("scale-up").Set(1)
				cooldown.RecordAction()
			case "scale-down":
				if len(nodes) > 0 {
					name := nodes[len(nodes)-1]
					_ = provider.RemoveNode(name)
					metrics.LastActionGauge.WithLabelValues("scale-down").Set(-1)
					cooldown.RecordAction()
				}
			default:
				metrics.LastActionGauge.WithLabelValues("no-op").Set(0)
			}
		}

		time.Sleep(interval)
	}
}
