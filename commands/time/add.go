package time

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/xrookiefight/vanillautils/commands/op"
)

type Add struct {
	Sub    cmd.SubCommand            `cmd:"add"`
	Amount int                       `cmd:"amount"`
	Args   cmd.Optional[cmd.Varargs] `cmd:"args"`
}

func (Add) Allow(src cmd.Source) bool {
	return op.IsOp(src)
}

func (t Add) Run(_ cmd.Source, output *cmd.Output, tx *world.Tx) {
	w := tx.World()
	w.SetTime(w.Time() + t.Amount)
	output.Printf("Added %d to the time", t.Amount)
}
