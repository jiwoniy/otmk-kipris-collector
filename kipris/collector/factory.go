package collector

import (
	"github.com/jiwoniy/otmk-kipris-collector/kipris/types"
)

func New(cfg types.CollectorConfig) (types.Collector, error) {
	c, err := NewCollector(cfg)
	if err != nil {
		return nil, err
	}
	return c, nil
}
