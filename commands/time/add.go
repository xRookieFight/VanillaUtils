package time

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/xrookiefight/vanillautils/commands/op"
)

type Add struct {
	Sub    cmd.SubCommand `cmd:"add"`
	Amount int            `name:"amount"`
}

func (t Add) Run(source cmd.Source, output *cmd.Output, tx *world.Tx) {
	if !op.IsOp(source) {
		output.Error("You don't have permission to run this command.")
		return
	}

	w := tx.World()
	w.SetTime(w.Time() + t.Amount)
	output.Printf("Added %d to the time", t.Amount)
}
