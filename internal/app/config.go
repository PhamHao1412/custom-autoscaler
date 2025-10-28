package app

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Autoscaler struct {
		IntervalSeconds   int     `yaml:"interval_seconds"`
		CooldownSeconds   int     `yaml:"cooldown_seconds"`
		ScaleUpCPU        float64 `yaml:"scale_up_cpu"`
		ScaleDownCPU      float64 `yaml:"scale_down_cpu"`
		ScaleUpMemory     float64 `yaml:"scale_up_mem"`
		ScaleDownMemory   float64 `yaml:"scale_down_mem""`
		ScaleUpRespTime   float64 `yaml:"scale_up_resp_time""`
		ScaleDownRespTime float64 `yaml:"scale_down_resp_time""`
		MinNodes          int     `yaml:"min_nodes"`
		MaxNodes          int     `yaml:"max_nodes"`
		Provider          string  `yaml:"provider"`
		PrometheusPort    int     `yaml:"prometheus_port"`
	} `yaml:"autoscaler"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read app file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse app file: %w", err)
	}
	return &cfg, nil
}
