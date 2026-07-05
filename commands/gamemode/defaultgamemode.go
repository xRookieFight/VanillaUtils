package gamemode

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/xrookiefight/vanillautils/commands/op"
)

type DefaultGameMode struct {
	GameMode mode // look at gamemode.go for "mode"
}

func (t DefaultGameMode) Run(source cmd.Source, output *cmd.Output, tx *world.Tx) {
	if !op.IsOp(source) {
		output.Error("You don't have permission to run this command.")
		return
	}
	mode := StringToGameMode(string(t.GameMode))
	tx.World().SetDefaultGameMode(mode)
	output.Printf("Default world game mode is %s now.", GameModeToName(mode))
}
