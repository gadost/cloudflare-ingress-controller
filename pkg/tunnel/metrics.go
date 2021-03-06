package tunnel

import (
	"time"

	"github.com/cloudflare/cloudflared/origin"
)

// MetricsConfig wraps the argo tunnel metrics in a struct
type MetricsConfig struct {
	Metrics         *origin.TunnelMetrics
	UpdateFrequency time.Duration
}

// NewMetrics created a set of TunnelMetrics,
// allows global prometheus objects, which breaks tests
func NewMetrics() *MetricsConfig {

	return &MetricsConfig{
		Metrics:         origin.NewTunnelMetrics(),
		UpdateFrequency: 5 * time.Second,
	}
}

// NewDummyMetrics creates a sample TunnelMetrics object
// full of default zero values
// does not initializs prometheus and is acceptable for tests
func NewDummyMetrics() *MetricsConfig {

	return &MetricsConfig{
		Metrics:         &origin.TunnelMetrics{},
		UpdateFrequency: 10000 * time.Hour,
	}
}
