package templates 

import (
	"github.com/evcc-io/config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "custom",
		Name:   "RCT Power Storage (Battery Meter)",
		Sample: `power:
  source: script
  cmd: /bin/bash -c "rctclient read-value --host 192.0.2.2 --name g_sync.p_acc_lp"
  timeout: 5s
soc:
  source: script
  cmd: /bin/bash -c "echo $(rctclient read-value --host 192.0.2.2 --name battery.soc) \* 100. | bc -l"
  timeout: 5s`,
	}

	registry.Add(template)
}
