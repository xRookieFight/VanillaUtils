package gamemode

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/xrookiefight/vanillautils/commands/op"
	"github.com/xrookiefight/vanillautils/lang"
)

type DefaultGameMode struct {
	GameMode mode                      // look at gamemode.go for "mode"
	Args     cmd.Optional[cmd.Varargs] `cmd:"args"`
}

func (DefaultGameMode) Allow(src cmd.Source) bool {
	return op.IsOp(src)
}

func (t DefaultGameMode) Run(source cmd.Source, output *cmd.Output, tx *world.Tx) {
	mode := StringToGameMode(string(t.GameMode))
	tx.World().SetDefaultGameMode(mode)
	output.Printt(lang.MessageDefaultGameMode, lang.GameModeName(mode))
}
