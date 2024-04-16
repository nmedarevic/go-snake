module sneakgame.com

go 1.21.4

replace sneakgame.com/constants v0.0.0 => ./constants

replace sneakgame.com/snake v0.0.0 => ./snake

replace sneakgame.com/keylistener v0.0.0 => ./keylistener

require (
	golang.design/x/hotkey v0.4.1
	sneakgame.com/keylistener v0.0.0
)

require (
	golang.design/x/mainthread v0.3.0 // indirect
	sneakgame.com/constants v0.0.0 // indirect
	sneakgame.com/snake v0.0.0
)
