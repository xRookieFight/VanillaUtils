package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/xrookiefight/vanillautils/commands/op"
	"github.com/xrookiefight/vanillautils/global"
)

type Stop struct{}

func (t Stop) Run(source cmd.Source, output *cmd.Output, _ *world.Tx) {
	if !op.IsOp(source) {
		output.Error("You don't have permission to run this command.")
		return
	}
	output.Printf("Stopping server.")
	// Close must not run inside the world transaction this command executes in,
	// otherwise shutting the world down would deadlock.
	go global.Server.Close()
}
