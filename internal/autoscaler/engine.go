package autoscaler

import (
	"custom-autoscaler/internal/metrics"
	"fmt"
)

type DecisionEngine struct {
	ScaleUpCPU        float64
	ScaleDownCPU      float64
	ScaleUpMem        float64
	ScaleDownMem      float64
	ScaleUpRespTime   float64
	ScaleDownRespTime float64
	MinNodes          int
	MaxNodes          int
}

func (d *DecisionEngine) Evaluate(m metrics.Metrics, currentNodes int) (action string, reason string) {
	cpu := m.CPUUsage
	mem := m.MemoryUsage
	rt := m.ResponseTimeMS

	switch {
	case (cpu > d.ScaleUpCPU || mem > d.ScaleUpMem || rt > d.ScaleUpRespTime) && currentNodes < d.MaxNodes:
		return "scale-up", fmt.Sprintf(
			"CPU %.2f%% / Mem %.2f%% / RT %.2fms > (%.2f%% | %.2f%% | %.2fms)",
			cpu, mem, rt, d.ScaleUpCPU, d.ScaleUpMem, d.ScaleUpRespTime,
		)

	case (cpu < d.ScaleDownCPU && mem < d.ScaleDownMem && rt < d.ScaleDownRespTime) && currentNodes > d.MinNodes:
		return "scale-down", fmt.Sprintf(
			"CPU %.2f%% / Mem %.2f%% / RT %.2fms < (%.2f%% | %.2f%% | %.2fms)",
			cpu, mem, rt, d.ScaleDownCPU, d.ScaleDownMem, d.ScaleDownRespTime,
		)

	default:
		return "no-op", fmt.Sprintf(
			"CPU %.2f%% / Mem %.2f%% / RT %.2fms within normal range",
			cpu, mem, rt,
		)
	}
}
