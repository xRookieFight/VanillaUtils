package op

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/xrookiefight/vanillautils/commands/utils"
	"github.com/xrookiefight/vanillautils/lang"
)

type Op struct {
	Player string
	Args   cmd.Optional[cmd.Varargs] `cmd:"args"`
}

func (Op) Allow(src cmd.Source) bool {
	return IsOp(src)
}

func (t Op) Run(_ cmd.Source, output *cmd.Output, tx *world.Tx) {
	if t.Player == "" {
		output.Errort(cmd.MessageUsage, "/op <Player: string>")
		return
	}
	if IsOpName(t.Player) {
		output.Errort(lang.MessageOpFailed, t.Player)
		return
	}

	AddOp(t.Player)
	output.Printt(lang.MessageOpSuccess, t.Player)

	if p, ok := utils.PlayerByName(t.Player, tx); ok {
		p.Messaget(lang.MessageOpSuccess, p.Name())
	}
}

type Deop struct {
	Player string
	Args   cmd.Optional[cmd.Varargs] `cmd:"args"`
}

func (Deop) Allow(src cmd.Source) bool {
	return IsOp(src)
}

func (t Deop) Run(_ cmd.Source, output *cmd.Output, tx *world.Tx) {
	if t.Player == "" {
		output.Errort(cmd.MessageUsage, "/deop <Player: string>")
		return
	}
	if !IsOpName(t.Player) {
		output.Errort(lang.MessageDeopFailed, t.Player)
		return
	}

	DelOp(t.Player)
	output.Printt(lang.MessageDeopSuccess, t.Player)

	if p, ok := utils.PlayerByName(t.Player, tx); ok {
		p.Messaget(lang.MessageDeopSuccess, p.Name())
	}
}
