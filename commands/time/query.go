package time

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/xrookiefight/vanillautils/lang"
)

// dayLength is the amount of ticks a full Minecraft day takes.
const dayLength = 24000

type Query struct {
	Sub  cmd.SubCommand            `cmd:"query"`
	Time timeQuery                 `cmd:"time"`
	Args cmd.Optional[cmd.Varargs] `cmd:"args"`
}

func (t Query) Run(source cmd.Source, output *cmd.Output, tx *world.Tx) {
	time := tx.World().Time()

	switch t.Time {
	case "day":
		output.Printt(lang.MessageTimeQuery, time/dayLength)
	case "daytime":
		output.Printt(lang.MessageTimeQuery, time%dayLength)
	default:
		output.Printt(lang.MessageTimeQuery, time)
	}
}

type timeQuery string

func (timeQuery) Type() string {
	return "TimeQuery"
}

func (timeQuery) Options(source cmd.Source) []string {
	return []string{"day", "daytime", "gametime"}
}
