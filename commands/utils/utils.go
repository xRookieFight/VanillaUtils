package utils

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/xrookiefight/vanillautils/global"
	"strings"
)

func SliceUnique(slice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range slice {
		if _, value := keys[entry]; !value && entry != "" {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func PlayerByName(name string, tx *world.Tx) (*player.Player, bool) {
	for p := range global.Server.Players(tx) {
		if strings.EqualFold(p.Name(), name) {
			return p, true
		}
	}
	return nil, false
}
