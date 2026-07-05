# VanillaUtils
Essential commands for [Dragonfly](https://github.com/df-mc/dragonfly), updated for the modern Dragonfly API (v0.10+).

Based on [EssentialsGO](https://github.com/xerenahmed/EssentialsGO) by [xerenahmed](https://github.com/xerenahmed) - all credit for the original project goes to him.

## Ready for:
- [x] /help - Show server commands and descriptions.
- [x] /gamemode - Changes the player to a specific game mode.
- [x] /teleport - Teleport everywhere.
- [x] /defaultgamemode - Set the default gamemode.
- [x] /setworldspawn - Sets a worlds' spawn point.
- [x] /xyz - Show/hide coordinates.
- [x] /op - Give op permissions to a player.
- [x] /deop - Take op permissions from a player.
- [x] /stop - Stop the server from in-game.
- [x] /time - Changes or queries the worlds game time.

## Usage
### Get the package
`go get -u github.com/xrookiefight/vanillautils`
### Import package
```go
import "github.com/xrookiefight/vanillautils"
```
### Register Commands
```go
vanillautils.RegisterCommands(server) // the server is *server.Server
```
### Exclude Commands
```go
vanillautils.RegisterCommandsWithout(server, []string{"stop", "defaultgamemode"}) // All commands will be loaded, except "stop" and "defaultgamemode"
```

### Simple Console Command Sender
```go
vanillautils.LoadConsole()
```

## Credits
- [xerenahmed](https://github.com/xerenahmed) - author of the original [EssentialsGO](https://github.com/xerenahmed/EssentialsGO) this project is based on.
