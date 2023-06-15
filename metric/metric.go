package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	Root prometheus.Counter
)

func NewCounter(name, help string, labels map[string]string) prometheus.Counter {
	opts := prometheus.CounterOpts{
		Name:        name,
		Help:        help,
		ConstLabels: make(prometheus.Labels),
	}
	for k, v := range labels {
		opts.ConstLabels[k] = v
	}
	return promauto.NewCounter(opts)
}
