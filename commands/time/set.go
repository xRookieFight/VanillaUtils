package time

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/xrookiefight/vanillautils/commands/op"
)

type Set struct {
	Sub    cmd.SubCommand            `cmd:"set"`
	Amount int                       `cmd:"amount"`
	Args   cmd.Optional[cmd.Varargs] `cmd:"args"`
}

func (Set) Allow(src cmd.Source) bool {
	return op.IsOp(src)
}

type SetTimeSpec struct {
	Sub  cmd.SubCommand            `cmd:"set"`
	Time spec                      `cmd:"time"`
	Args cmd.Optional[cmd.Varargs] `cmd:"args"`
}

func (SetTimeSpec) Allow(src cmd.Source) bool {
	return op.IsOp(src)
}

func setTime(output *cmd.Output, tx *world.Tx, t int) {
	tx.World().SetTime(t)
	output.Printf("Set the time to %d", t)
}

func (t Set) Run(_ cmd.Source, output *cmd.Output, tx *world.Tx) {
	setTime(output, tx, t.Amount)
}

func (t SetTimeSpec) Run(_ cmd.Source, output *cmd.Output, tx *world.Tx) {
	tf := map[spec]int64{
		"day": 1000, "night": 13000, "noon": 6000, "midnight": 18000, "sunrise": 23000, "sunset": 12000,
	}[t.Time]
	setTime(output, tx, int(tf))
}

type spec string

func (spec) Type() string {
	return "TimeSpec"
}

func (spec) Options(source cmd.Source) []string {
	return []string{"day", "night", "noon", "midnight", "sunrise", "sunset"}
}
