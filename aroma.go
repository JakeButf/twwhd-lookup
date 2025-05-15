//aroma.go
//hashmaps of aroma specific addresses
package main
type AromaSource struct{}

type AddressRange struct {
	Start uint32 `json:"start"`
	End uint32 `json:"end`
}

var AddressMap = map[string]uint32{
	"X Position": 0x1096EF48,
	"Y Position": 0x1096EF4C,
	"Z Position": 0x1096EF50,
	"Speed Angle": 0x1096EF0A,
	"Facing Angle": 0x1096EF12,
	"Current Stage": 0x109763E4,
	"Current Spawn": 0x109763ED,
	"Current Layer": 0x109763EF,
	"Next Stage": 0x109763F0,
	"Next Spawn": 0x109763F9,
	"Next Room": 0x109763FA,
	"Next Layer": 0x109763FB,
	"Trigger Loading": 0x109763FC,
	"Event State": 0x10976542,
	"Storage": 0x10976543,
	"Fairy Fountain Get Item Value": 0x10976554,
	"Current Room": 0x10978CF8,
	"Link Pointer": 0x10989C74,
	"Gamepad Input": 0x15073684,
	"Max Hearts": 0x15073681,
	"Current Hearts": 0x15073683,
	"Magic": 0x15073694,
	"Bombs": 0x150736EA,
	"Arrows": 0x150736E9,
	"Time": 0x150736A4,
}

var AddressRangeMap = map[string]AddressRange{
	"Scene Flags": {Start: 0x15073DF8, End: 0x15073E18},
	"Global Flags": {Start: 0x15073CA4, End: 0x15073CE3},
	"Ocean Color": {Start: 0x26917424, End: 0x26917428},
}

func (a AromaSource) GetData() map[string]any {
	return map[string]any{
		"addresses": AddressMap,
		"ranges": AddressRangeMap,
	}
}
