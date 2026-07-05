package time

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
)

type Query struct {
	Sub  cmd.SubCommand            `cmd:"query"`
	Time timeQuery                 `cmd:"time"`
	Args cmd.Optional[cmd.Varargs] `cmd:"args"`
}

func (t Query) Run(source cmd.Source, output *cmd.Output, _ *world.Tx) {
}

type timeQuery string

func (timeQuery) Type() string {
	return "TimeQuery"
}

func (timeQuery) Options(source cmd.Source) []string {
	return []string{"day", "daytime", "gametime"}
}
