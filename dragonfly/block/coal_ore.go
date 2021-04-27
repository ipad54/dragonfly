package block

import (
	"github.com/df-mc/dragonfly/dragonfly/item"
)

// CoalOre is a common ore.
type CoalOre struct {
	noNBT
	solid
	bassDrum
}

// BreakInfo ...
func (c CoalOre) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    3,
		Harvestable: pickaxeHarvestable,
		Effective:   pickaxeEffective,
		Drops:       simpleDrops(item.NewStack(item.Coal{}, 1)), //TODO: Silk Touch
		XPDrops:     XPDropRange{0, 2},
	}
}

// EncodeItem ...
func (CoalOre) EncodeItem() (id int32, meta int16) {
	return 16, 0
}

// EncodeBlock ...
func (CoalOre) EncodeBlock() (string, map[string]interface{}) {
	return "minecraft:coal_ore", nil
}
