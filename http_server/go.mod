module main

go 1.22.2

replace snakegame.com/snake v0.0.0 => ../snake/snake

replace snakegame.com/constants v0.0.0 => ../snake/constants

replace snakegame.com/http/start => ./start

require (
	github.com/a-h/templ v0.2.778
	snakegame.com/http/start v0.0.0
	snakegame.com/snake v0.0.0
)

require snakegame.com/constants v0.0.0 // indirect
