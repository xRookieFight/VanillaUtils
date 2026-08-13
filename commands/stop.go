package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/xrookiefight/vanillautils/commands/op"
	"github.com/xrookiefight/vanillautils/global"
	"github.com/xrookiefight/vanillautils/lang"
)

type Stop struct {
	Args cmd.Optional[cmd.Varargs] `cmd:"args"`
}

func (Stop) Allow(src cmd.Source) bool {
	return op.IsOp(src)
}

func (t Stop) Run(_ cmd.Source, output *cmd.Output, _ *world.Tx) {
	output.Printt(lang.MessageStop)
	// Close must not run inside the world transaction this command executes in,
	// otherwise shutting the world down would deadlock.
	go global.Server.Close()
}
