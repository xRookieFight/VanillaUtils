package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/xrookiefight/vanillautils/global"
	"math"
)

type Help struct {
	Page cmd.Optional[int]         `cmd:"page"`
	Args cmd.Optional[cmd.Varargs] `cmd:"args"`
}

func (t Help) Run(source cmd.Source, output *cmd.Output, _ *world.Tx) {
	p, _ := t.Page.Load()
	page := int(math.Max(float64(p), 1))
	output.Printf("--- Help Page %d ---", page)
	max := page * 5

	// Only list commands the source is actually allowed to run.
	var visible []cmd.Command
	for _, c := range global.Commands {
		if len(c.Runnables(source)) > 0 {
			visible = append(visible, c)
		}
	}

	for i, c := range visible {
		if i < max-5 {
			continue
		} else if i >= max {
			break
		}
		output.Printf("%s: %s", c.Name(), c.Description())
	}

	if output.MessageCount() == 1 {
		output.Errorf("There are only %d pages.", int(math.Round(float64(len(visible))/5)))
	}
}
