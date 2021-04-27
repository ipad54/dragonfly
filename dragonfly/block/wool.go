package block

import (
	"github.com/df-mc/dragonfly/dragonfly/block/colour"
	"github.com/df-mc/dragonfly/dragonfly/block/instrument"
	"github.com/df-mc/dragonfly/dragonfly/item"
	"github.com/df-mc/dragonfly/dragonfly/world"
)

// Wool is a colourful block that can be obtained by killing/shearing sheep, or crafted using four string.
type Wool struct {
	noNBT
	solid

	// Colour is the colour of the wool.
	Colour colour.Colour
}

// Instrument ...
func (w Wool) Instrument() instrument.Instrument {
	return instrument.Guitar()
}

// FlammabilityInfo ...
func (w Wool) FlammabilityInfo() FlammabilityInfo {
	return FlammabilityInfo{
		Encouragement: 30,
		Flammability:  60,
		LavaFlammable: true,
	}
}

// BreakInfo ...
func (w Wool) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    0.8,
		Harvestable: alwaysHarvestable,
		Effective:   shearsEffective,
		Drops:       simpleDrops(item.NewStack(w, 1)),
	}
}

// EncodeItem ...
func (w Wool) EncodeItem() (id int32, meta int16) {
	return 35, int16(w.Colour.Uint8())
}

// EncodeBlock ...
func (w Wool) EncodeBlock() (name string, properties map[string]interface{}) {
	return "minecraft:wool", map[string]interface{}{"color": w.Colour.String()}
}

// allWool returns wool blocks with all possible colours.
func allWool() []world.Block {
	b := make([]world.Block, 0, 16)
	for _, c := range colour.All() {
		b = append(b, Wool{Colour: c})
	}
	return b
}
