module sneakgame.com

go 1.22.2

replace sneakgame.com/constants v0.0.0 => ./constants

replace sneakgame.com/snake v0.0.0 => ./snake

replace sneakgame.com/snake_cli v0.0.0 => ./snake_cli

replace sneakgame.com/keylistener v0.0.0 => ./keylistener

require (
	golang.design/x/hotkey v0.4.1 // indirect
	sneakgame.com/keylistener v0.0.0 // indirect
)

require (
	golang.design/x/mainthread v0.3.0 // indirect
	sneakgame.com/constants v0.0.0 // indirect
	sneakgame.com/snake v0.0.0 // indirect
	sneakgame.com/snake_cli v0.0.0
)
