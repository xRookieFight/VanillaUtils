package gamemode

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/xrookiefight/vanillautils/commands/op"
	"github.com/xrookiefight/vanillautils/lang"
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
	name := lang.GameModeName(mode)

	if targets, _ := t.Target.Load(); len(targets) > 0 {
		if pt, ok := targets[0].(*player.Player); ok {
			pt.SetGameMode(mode)
			output.Printt(lang.MessageGameModeOther, pt.Name(), name)
			pt.Messaget(lang.MessageGameModeChanged, name)
		} else {
			output.Errort(lang.MessagePlayerNotFound)
		}

		return
	}

	if p, ok := source.(*player.Player); ok {
		p.SetGameMode(mode)
		output.Printt(lang.MessageGameModeSelf, name)
	} else {
		output.Errort(cmd.MessageUsage, "/gamemode <GameMode: mode> <Target: target>")
	}
}
