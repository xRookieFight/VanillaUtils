package op

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
)

type Op struct {
	Player string
	Args   cmd.Optional[cmd.Varargs] `cmd:"args"`
}

func (Op) Allow(src cmd.Source) bool {
	return IsOp(src)
}

func (t Op) Run(source cmd.Source, output *cmd.Output, _ *world.Tx) {
	if t.Player != "" {
		AddOp(t.Player)
		output.Printf("Has been granted op permissions to %s.", t.Player)
	} else {
		output.Error("Usage: /op <Player: string>")
	}
}

type Deop struct {
	Player string
	Args   cmd.Optional[cmd.Varargs] `cmd:"args"`
}

func (Deop) Allow(src cmd.Source) bool {
	return IsOp(src)
}

func (t Deop) Run(source cmd.Source, output *cmd.Output, _ *world.Tx) {
	if t.Player != "" {
		DelOp(t.Player)
		output.Printf("Has been taken op permissions from %s.", t.Player)
	} else {
		output.Error("Usage: /deop <Player: string>")
	}
}
