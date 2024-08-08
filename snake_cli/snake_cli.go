package snake_cli

import (
	"time"

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

const TABLE_SIZE = 10
const RENDER_TICKER_REFRESH_MILIS = 40
const MOVE_TICKER_MILIS = 200

type SnakeCli struct {
	snake     *snakeModule.Snake
	table     *[][]uint8
	lastInput uint8

	doneChannel          chan bool
	keyboardInputChannel chan uint8
	ticker               *time.Ticker
	movementTicket       *time.Ticker
}

func (cli *SnakeCli) InitializeSnake() {
	var startPointX uint8 = 3
	var startPointY uint8 = 5
	snakeLength := 3

	snejk := snakeModule.MakeSnake(startPointX, startPointY, uint8(snakeLength))

	cli.snake = snejk
}

func (cli *SnakeCli) InitializeTable() {
	table := make([][]uint8, TABLE_SIZE)

	for i := range table {
		table[i] = make([]uint8, TABLE_SIZE)
	}

	cli.table = &table
}

func (cli *SnakeCli) InitializeGame() {
	doneChannel := make(chan bool, 1)
	keyboardInputChannel := make(chan uint8, 10)
	ticker := time.NewTicker(RENDER_TICKER_REFRESH_MILIS * time.Millisecond)

	cli.lastInput = constants.Down
	cli.doneChannel = doneChannel
	cli.keyboardInputChannel = keyboardInputChannel
	cli.ticker = ticker
	cli.movementTicket = time.NewTicker(MOVE_TICKER_MILIS * time.Millisecond)
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func renderGame(snake *snakeModule.Snake, table *[][]uint8) {
	clearScreen()
	snakeModule.PrintTable(table, snake)
}

func exitWhenDone(isDone bool) {
	if isDone {
		clearScreen()
		os.Exit(0)
		return
	}
}

func (cli *SnakeCli) RunGame() {
	var isDone bool

	go func() {
		for {
			exitWhenDone(isDone)
			<-cli.ticker.C
			renderGame(cli.snake, cli.table)
		}
	}()

	go func() {
		for {
			exitWhenDone(isDone)
			<-cli.movementTicket.C
			snakeModule.MoveSnake(cli.snake, cli.table, cli.lastInput)
		}
	}()

	go func() {
		for {
			isDone = <-cli.doneChannel

			exitWhenDone(isDone)
		}
	}()

	go func() {
		for {
			latestMovementKey := <-cli.keyboardInputChannel

			if constants.IsIllegalMovementKey(latestMovementKey, cli.lastInput) {
				continue
			}

			cli.lastInput = latestMovementKey

			exitWhenDone(isDone)
		}
	}()

	mainthread.Init(keylistener.GetKeyListener(cli.doneChannel, cli.keyboardInputChannel))
}
