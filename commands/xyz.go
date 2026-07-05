package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/xrookiefight/vanillautils/console"
)

type XYZ struct{}

var opened = map[string]bool{}

func (t XYZ) Run(source cmd.Source, output *cmd.Output, _ *world.Tx) {
	if _, ok := source.(*console.Console); ok {
		output.Error("Use in game!")
		return
	}
	var msg string
	p, _ := source.(*player.Player)
	id := p.Name()

	opened[id] = !opened[id]

	if opened[id] {
		p.ShowCoordinates()
		msg = "shown"
	} else {
		p.HideCoordinates()
		msg = "hidden"
	}

	output.Print("Coordinates " + msg + ".")
}
