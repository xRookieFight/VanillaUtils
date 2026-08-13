package lang

import (
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-mc/dragonfly/server/world"
	"golang.org/x/text/language"
)

// str is a plain translation identifier that resolves to itself, mirroring the
// unexported type dragonfly uses for its own translations.
type str string

func (s str) Resolve(language.Tag) string { return string(s) }

// Translation strings below use the vanilla Bedrock keys, so clients render
// them in their own language. The fallbacks match the English values and are
// used for sources without a language, such as the console.
var (
	MessageGameModeSelf      = chat.Translate(str("%commands.gamemode.success.self"), 1, `Set own game mode to %v`)
	MessageGameModeOther     = chat.Translate(str("%commands.gamemode.success.other"), 2, `Set %v's game mode to %v`)
	MessageGameModeChanged   = chat.Translate(str("%gameMode.changed"), 1, `Your game mode has been updated to %v`)
	MessageDefaultGameMode   = chat.Translate(str("%commands.defaultgamemode.success"), 1, `The world's default game mode is now %v`)
	MessageTimeSet           = chat.Translate(str("%commands.time.set"), 1, `Set the time to %v`)
	MessageTimeAdded         = chat.Translate(str("%commands.time.added"), 1, `Added %v to the time`)
	MessageTimeQuery         = chat.Translate(str("%commands.time.query"), 1, `Time is %v`)
	MessageOpSuccess         = chat.Translate(str("%commands.op.success"), 1, `Opped: %v`)
	MessageOpFailed          = chat.Translate(str("%commands.op.failed"), 1, `Could not op %v`)
	MessageDeopSuccess       = chat.Translate(str("%commands.deop.success"), 1, `De-opped: %v`)
	MessageDeopFailed        = chat.Translate(str("%commands.deop.failed"), 1, `Could not de-op %v`)
	MessageSetWorldSpawn     = chat.Translate(str("%commands.setworldspawn.success"), 3, `Set the world spawn point to (%v, %v, %v)`)
	MessageStop              = chat.Translate(str("%commands.stop.start"), 0, `Stopping the server`)
	MessagePlayerNotFound    = chat.Translate(str("%commands.generic.player.notFound"), 0, `That player cannot be found`)
	MessageGameModeSurvival  = chat.Translate(str("%gameMode.survival"), 0, `Survival Mode`)
	MessageGameModeCreative  = chat.Translate(str("%gameMode.creative"), 0, `Creative Mode`)
	MessageGameModeAdventure = chat.Translate(str("%gameMode.adventure"), 0, `Adventure Mode`)
	MessageGameModeSpectator = chat.Translate(str("%gameMode.spectator"), 0, `Spectator Mode`)
)

// GameModeName returns the translation holding the vanilla display name of the
// game mode passed. It may be used as a parameter of another translation.
func GameModeName(gm world.GameMode) chat.Translation {
	return map[world.GameMode]chat.Translation{
		world.GameModeSurvival:  MessageGameModeSurvival,
		world.GameModeCreative:  MessageGameModeCreative,
		world.GameModeAdventure: MessageGameModeAdventure,
		world.GameModeSpectator: MessageGameModeSpectator,
	}[gm]
}
