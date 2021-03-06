package process

import (
	"github.com/swapbyt3s/zenit/common"
	"github.com/swapbyt3s/zenit/common/log"
	"github.com/swapbyt3s/zenit/config"
	"github.com/swapbyt3s/zenit/plugins/inputs"
	"github.com/swapbyt3s/zenit/plugins/lists/metrics"
)

type InputsPerconaKill struct{}

func (l *InputsPerconaKill) Collect() {
	defer func() {
		if err := recover(); err != nil {
			log.Error("InputsPerconaKill", map[string]interface{}{"error": err})
		}
	}()

	if !config.File.Inputs.Process.PerconaToolKitKill {
		return
	}

	var a = metrics.Load()
	var pid = common.PGrep("pt-kill")
	var value = 0

	if pid > 0 {
		value = 1
	}

	log.Debug("InputsPerconaKill", map[string]interface{}{"pt_kill": value})

	a.Add(metrics.Metric{
		Key: "process_pt_kill",
		Tags: []metrics.Tag{
			{"hostname", config.File.General.Hostname},
		},
		Values: []metrics.Value{
			{"pt_kill", value},
		},
	})
}

func init() {
	inputs.Add("InputsPerconaKill", func() inputs.Input { return &InputsPerconaKill{} })
}
