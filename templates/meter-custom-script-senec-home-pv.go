package templates 

import (
	"github.com/evcc-io/config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "custom",
		Name:   "SENEC.Home (PV Meter)",
		Sample: `power:
  source: script
  cmd: >
    /bin/bash -c "set +H; curl --data '{\"ENERGY\":{\"GUI_INVERTER_POWER\":\"\"}}' --header \"Content-Type: application/json\" --request POST http://192.0.2.2/lala.cgi | jq .ENERGY.GUI_INVERTER_POWER | python3 -c 'import struct;print(struct.unpack(\"!f\",bytes.fromhex(input()[4:12]))[0])'"
  timeout: 5s`,
	}

	registry.Add(template)
}
