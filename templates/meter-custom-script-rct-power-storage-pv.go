package templates 

import (
	"github.com/evcc-io/config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "custom",
		Name:   "RCT Power Storage (PV Meter)",
		Sample: `power:
  source: script
  cmd: /bin/bash -c "echo $(rctclient read-value --host 192.0.2.2 --name g_sync.p_ac_load_sum_lp) \- $(rctclient read-value --host 192.0.2.2 --name g_sync.p_acc_lp) \- $(rctclient read-value --host 192.0.2.2 --name g_sync.p_ac_grid_sum_lp) | bc -l"
  timeout: 5s`,
	}

	registry.Add(template)
}
