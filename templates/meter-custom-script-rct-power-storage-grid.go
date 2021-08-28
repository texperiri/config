package templates 

import (
	"github.com/evcc-io/config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "custom",
		Name:   "RCT Power Storage (Grid Meter)",
		Sample: `power:
  source: script
  cmd: /bin/bash -c "rctclient read-value --host 192.0.2.2 --name g_sync.p_ac_grid_sum_lp"
  timeout: 5s`,
	}

	registry.Add(template)
}
