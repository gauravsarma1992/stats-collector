package statscollector

import "log"

var (
	IncrActionT    = ActionT(1)
	DecrActionT    = ActionT(2)
	AddActionT     = ActionT(3)
	ReplaceActionT = ActionT(4)
)

type (
	LabelT  map[string]interface{}
	ActionT uint8

	MetricsStore struct {
		Metrics map[string]*Metric `json:"metrics"`
	}

	Metric struct {
		Name   string  `json:"name"`
		Value  float64 `json:"value"`
		Action ActionT `json:"action"`

		Labels map[string]interface{} `json:"labels"`
	}
)

func NewMetricStore() (ms *MetricsStore, err error) {
	log.Println("Initializing Metric Store")
	ms = &MetricsStore{
		Metrics: make(map[string]*Metric),
	}
	return
}

func (ms *MetricsStore) Add(metric *Metric) (err error) {
	var (
		prevMetric *Metric
		isPresent  bool
		currValue  float64
	)
	if prevMetric, isPresent = ms.Metrics[metric.Name]; !isPresent {
		prevMetric = metric
		prevMetric.Value = 0
		ms.Metrics[metric.Name] = prevMetric
	}

	currValue = prevMetric.Value

	switch metric.Action {
	case IncrActionT:
		currValue += 1
	case DecrActionT:
		currValue -= 1
	case AddActionT:
		currValue += metric.Value
	case ReplaceActionT:
		currValue = metric.Value
	default:
		currValue += metric.Value
	}
	prevMetric.Value = currValue
	return
}

func (ms *MetricsStore) List() (metrics []*Metric, err error) {
	for _, metric := range ms.Metrics {
		metrics = append(metrics, metric)
	}
	return
}
