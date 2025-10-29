package app

type Config struct {
	Autoscaler struct {
		IntervalSeconds   int     `mapstructure:"INTERVAL_SECONDS"`
		CooldownSeconds   int     `mapstructure:"COOLDOWN_SECONDS"`
		ScaleUpCPU        float64 `mapstructure:"SCALE_UP_CPU"`
		ScaleDownCPU      float64 `mapstructure:"SCALE_DOWN_CPU"`
		ScaleUpMemory     float64 `mapstructure:"SCALE_UP_MEM"`
		ScaleDownMemory   float64 `mapstructure:"SCALE_DOWN_MEM"`
		ScaleUpRespTime   float64 `mapstructure:"SCALE_UP_RESP_TIME"`
		ScaleDownRespTime float64 `mapstructure:"SCALE_DOWN_RESP_TIME"`
		MinNodes          int     `mapstructure:"MIN_NODES"`
		MaxNodes          int     `mapstructure:"MAX_NODES"`
		Provider          string  `mapstructure:"PROVIDER"`
		PrometheusPort    int     `mapstructure:"PROMETHEUS_PORT"`
	} `mapstructure:"AUTOSCALER"`
}
