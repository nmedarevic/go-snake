module snakegame.com

go 1.22.2

replace snakegame.com/constants v0.0.0 => ./constants

replace snakegame.com/snake v0.0.0 => ./snake

replace snakegame.com/snake_cli v0.0.0 => ./snake_cli

replace snakegame.com/keylistener v0.0.0 => ./keylistener

require (
	golang.design/x/hotkey v0.4.1 // indirect
	snakegame.com/keylistener v0.0.0 // indirect
)

require (
	golang.design/x/mainthread v0.3.0 // indirect
	snakegame.com/constants v0.0.0 // indirect
	snakegame.com/snake v0.0.0 // indirect
	snakegame.com/snake_cli v0.0.0
)
