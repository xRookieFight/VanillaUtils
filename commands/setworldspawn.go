package commands

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/xrookiefight/vanillautils/commands/op"
)

type SetWorldSpawnXYZ struct {
	X, Y, Z float64
	Args    cmd.Optional[cmd.Varargs] `cmd:"args"`
}

func (SetWorldSpawnXYZ) Allow(src cmd.Source) bool {
	return op.IsOp(src)
}

func (t SetWorldSpawnXYZ) Run(source cmd.Source, output *cmd.Output, tx *world.Tx) {
	bp := cube.Pos{int(t.X), int(t.Y), int(t.Z)}
	tx.World().SetSpawn(bp)
	output.Printf("Set the default world spawn point to (%d, %d, %d)", bp.X(), bp.Y(), bp.Z())
}

type SetWorldSpawn struct {
	Args cmd.Optional[cmd.Varargs] `cmd:"args"`
}

func (SetWorldSpawn) Allow(src cmd.Source) bool {
	return op.IsOp(src)
}

func (t SetWorldSpawn) Run(source cmd.Source, output *cmd.Output, tx *world.Tx) {
	if args, ok := t.Args.Load(); ok && args != "" {
		// Arguments were passed but did not match the XYZ overload.
		output.Error("Usage: /setworldspawn <X: float> <Y: float> <Z: float>")
		return
	}
	if p, ok := source.(*player.Player); ok {
		pos := p.Position()
		bp := cube.Pos{int(pos.X()), int(pos.Y()), int(pos.Z())}
		tx.World().SetSpawn(bp)
		output.Printf("Set the world spawn point to (%d, %d, %d)", bp.X(), bp.Y(), bp.Z())
	} else {
		output.Error("Usage: /setworldspawn <X: float> <Y: float> <Z: float>")
	}
}
