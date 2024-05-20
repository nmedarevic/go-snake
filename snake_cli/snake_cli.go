package snake_cli

import (
	"time"

	"fmt"
	"os"
	"os/exec"

	"golang.design/x/hotkey/mainthread"
	"sneakgame.com/constants"
	keylistener "sneakgame.com/keylistener"
	snakeModule "sneakgame.com/snake"
)

type ISnakeCli interface {
	InitializeSnake(uint8, uint8, uint8)
	InitializeTable()
	InitializeGame()
	RunGame()
}

type SnakeCli struct {
	snake     *snakeModule.Snake
	table     *[][]uint8
	lastInput uint8

	doneChannel          chan bool
	keyboardInputChannel chan uint8
	ticker               *time.Ticker
}

func (cli *SnakeCli) InitializeSnake() {
	var startPointX uint8 = 3
	var startPointY uint8 = 5
	snakeLength := 3

	snejk := snakeModule.MakeSnake(startPointX, startPointY, uint8(snakeLength))

	cli.snake = snejk
}

func (cli *SnakeCli) InitializeTable() {
	table := make([][]uint8, 10)
	for i := range table {
		table[i] = make([]uint8, 10)
	}

	cli.table = &table
}

func (cli *SnakeCli) InitializeGame() {
	doneChannel := make(chan bool, 1)
	keyboardInputChannel := make(chan uint8, 10)
	ticker := time.NewTicker(1000 * time.Millisecond)

	cli.lastInput = constants.Down
	cli.doneChannel = doneChannel
	cli.keyboardInputChannel = keyboardInputChannel
	cli.ticker = ticker
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func renderGame(snake *snakeModule.Snake, table *[][]uint8, direction uint8) {
	clearScreen()
	snakeModule.MoveSnake(snake, table, direction)
	snakeModule.PrintTable(table, snake)
}

func (cli *SnakeCli) RunGame() {
	go func() {
		for {
			<-cli.ticker.C
			renderGame(cli.snake, cli.table, cli.lastInput)
		}
	}()

	go func() {
		for {
			<-cli.doneChannel
			fmt.Println("Done!")
		}
	}()

	go func() {
		for {
			cli.lastInput = <-cli.keyboardInputChannel
			fmt.Println("Key pressed:", cli.lastInput)
			renderGame(cli.snake, cli.table, cli.lastInput)
		}
	}()

	mainthread.Init(keylistener.GetKeyListener(cli.doneChannel, cli.keyboardInputChannel))
}
