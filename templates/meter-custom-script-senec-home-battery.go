package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "custom",
		Name:   "SENEC.Home (Battery)",
		Sample: `power:
  source: script
  cmd: /bin/bash -c "curl -d '{\"ENERGY\":{\"GUI_BAT_DATA_POWER\":\"\"}}' -H \"Content-Type: application/json\" -X POST http://192.0.2.2/lala.cgi | jq .ENERGY.GUI_BAT_DATA_POWER | python3 -c 'import struct;print(struct.unpack(\"!f\",bytes.fromhex(input()[3:]))[0])'"
  timeout: 5s
  scale: -1
soc:
  source: script
  cmd: /bin/bash -c "curl -d '{\"ENERGY\":{\"GUI_BAT_DATA_FUEL_CHARGE\":\"\"}}' -H \"Content-Type: application/json\" -X POST http://192.0.2.2/lala.cgi | jq .ENERGY.GUI_BAT_DATA_FUEL_CHARGE | python3 -c 'import struct;print(struct.unpack(\"!f\",bytes.fromhex(input()[3:]))[0])'"
  timeout: 5s`,
	}

	registry.Add(template)
}