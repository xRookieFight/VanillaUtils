# VanillaUtils
Essential commands for [Dragonfly](https://github.com/df-mc/dragonfly), updated for the modern Dragonfly API (v0.10+).

Based on [EssentialsGO](https://github.com/xerenahmed/EssentialsGO) by [xerenahmed](https://github.com/xerenahmed) - all credit for the original project goes to him.

## Available commands:
- /help - Show server commands and descriptions.
- /gamemode - Changes the player to a specific game mode.
- /teleport - Teleport everywhere.
- /defaultgamemode - Set the default gamemode.
- /setworldspawn - Sets a worlds' spawn point.
- /xyz - Show/hide coordinates.
- /op - Give op permissions to a player.
- /deop - Take op permissions from a player.
- /stop - Stop the server from in-game.
- /time - Changes or queries the worlds game time.

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
