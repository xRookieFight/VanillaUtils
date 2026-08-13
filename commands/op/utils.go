package op

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/xrookiefight/vanillautils/commands/utils"
	"github.com/xrookiefight/vanillautils/console"
	"os"
	"strings"
)

var ops = make([]string, 0)

func GetOps() []string {
	return ops
}

func AddOp(p string) {
	ops = append(ops, strings.ToLower(p))
	SaveOps()
}

func DelOp(p string) {
	var newOps []string
	for _, p_ := range ops {
		if p_ != strings.ToLower(p) {
			newOps = append(newOps, p_)
		}
	}

	ops = newOps
	SaveOps()
}

func IsOp(s cmd.Source) bool {
	if _, ok := s.(*console.Console); ok {
		return true
	}
	if p, ok := s.(*player.Player); ok {
		return IsOpName(p.Name())
	}

	return false
}

func IsOpName(name string) bool {
	for _, p_ := range ops {
		if p_ == strings.ToLower(name) {
			return true
		}
	}

	return false
}

func LoadOps() {
	if _, err := os.Stat("ops.txt"); os.IsNotExist(err) {
		if _, err = os.Create("ops.txt"); err != nil {
			fmt.Println("Error on creating ops.txt: " + err.Error())
			return
		}
	}
	data, err := os.ReadFile("ops.txt")
	if err == nil {
		ops = utils.SliceUnique(strings.Split(string(data), "\n"))
	}
}

func SaveOps() {
	_ = os.WriteFile("ops.txt", []byte(strings.Join(utils.SliceUnique(ops), "\n")), 0655)
}
