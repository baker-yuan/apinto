package limiting_strategy

import (
	"fmt"
	"github.com/eolinker/apinto/drivers"
	"github.com/eolinker/apinto/strategy"
	"github.com/eolinker/eosc"
)

func Check(cfg *Config, workers map[eosc.RequireId]eosc.IWorker) error {
 
	return checkConfig(cfg)
}

func checkConfig(conf *Config) error {
	if conf.Priority > 999 || conf.Priority < 1 {
		return fmt.Errorf("priority value %d not allow ", conf.Priority)
	}
	_, err := strategy.ParseFilter(conf.Filters)
	if err != nil {
		return err
	}

	return nil
}

func Create(id, name string, v *Config, workers map[eosc.RequireId]eosc.IWorker) (eosc.IWorker, error) {
	if err := Check(v, workers); err != nil {
		return nil, err
	}

	lg := &Limiting{
		WorkerBase: drivers.Worker(id, name),
	}

	err := lg.Reset(v, workers)
	if err != nil {
		return nil, err
	}

	controller.Store(id)
	return lg, nil
}
