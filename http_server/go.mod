module main

go 1.22.2

replace snakegame.com/snake v0.0.0 => ../snake/snake
replace snakegame.com/constants v0.0.0 => ../snake/constants
replace snakegame.com/http/start => ./start

require (
	github.com/a-h/templ v0.2.778 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	snakegame.com/http/start v0.0.0 // indirect
	snakegame.com/snake v0.0.0
	snakegame.com/constants v0.0.0
)
