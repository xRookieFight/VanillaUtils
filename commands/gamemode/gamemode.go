package gamemode

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/xrookiefight/vanillautils/commands/op"
)

type GameMode struct {
	GameMode mode
	Target   cmd.Optional[[]cmd.Target]
	Args     cmd.Optional[cmd.Varargs] `cmd:"args"`
}

func (GameMode) Allow(src cmd.Source) bool {
	return op.IsOp(src)
}

func (t GameMode) Run(source cmd.Source, output *cmd.Output, _ *world.Tx) {
	mode := StringToGameMode(string(t.GameMode))
	modeString := GameModeToName(mode)

	if targets, _ := t.Target.Load(); len(targets) > 0 {
		if pt, ok := targets[0].(*player.Player); ok {
			pt.SetGameMode(mode)
			output.Printf("Set %s game mode to %s.", pt.Name(), modeString)
			pt.Messagef("Your game mode has been changed to %s.", modeString)
		} else {
			output.Errorf("Target is invalid!")
		}

		return
	}

	if p, ok := source.(*player.Player); ok {
		p.SetGameMode(mode)
		output.Printf("Set own game mode to %s.", modeString)
	} else {
		output.Error("Usage: /gamemode <GameMode: mode> <Target: target>")
	}
}
