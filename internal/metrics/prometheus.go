package metrics

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	CPUUsageGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "autoscaler_cpu_usage_percent",
		Help: "Simulated CPU usage percentage",
	})
	NodeCountGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "autoscaler_node_count",
		Help: "Current number of nodes in mock cloud",
	})
	LastActionGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "autoscaler_last_action",
		Help: "Last scaling action: 1=scale-up, -1=scale-down, 0=no-op",
	}, []string{"action"})
	MemoryUsageGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "autoscaler_memory_usage_percent",
		Help: "Simulated Memory usage percentage",
	})
	ResponseTimeGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "autoscaler_request_response_time_ms",
		Help: "Simulated average response time of system requests (milliseconds)",
	})
)

func InitPrometheus(port int) {
	prometheus.MustRegister(CPUUsageGauge, NodeCountGauge, LastActionGauge, MemoryUsageGauge, ResponseTimeGauge)

	http.Handle("/metrics", promhttp.Handler())

	go func() {
		addr := fmt.Sprintf(":%d", port)
		fmt.Printf("📊 Prometheus metrics server running at http://localhost%s/metrics\n", addr)
		if err := http.ListenAndServe(addr, nil); err != nil {
			fmt.Printf("❌ metrics server error: %v\n", err)
		}
	}()
}
