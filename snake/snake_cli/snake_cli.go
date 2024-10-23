package snake_cli

import (
	"fmt"
	"time"

	"os"
	"os/exec"

	"golang.design/x/hotkey/mainthread"
	"snakegame.com/constants"
	keylistener "snakegame.com/keylistener"
	snakeModule "snakegame.com/snake"
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

type Point struct {
	X int32
	Y int32
}
type Board struct {
	TopLeft     Point
	BottomRight Point
}
type SnakeCli struct {
	snake     *snakeModule.Snake
	table     *[][]uint8
	board     *Board
	lastInput uint8

	doneChannel          chan bool
	keyboardInputChannel chan uint8
	ticker               *time.Ticker
	movementTicket       *time.Ticker
}

func (board *Board) Print(snakes []*snakeModule.Snake) {
	// Add more variations for points beyond 0, 0
	for y := board.TopLeft.Y; y >= board.BottomRight.Y; y-- {
		for x := board.TopLeft.X; x <= board.BottomRight.X; x++ {
			for _, snake := range snakes {
				for index, snakePoint := range *snake.Body {
					if x == snakePoint.X && y == snakePoint.Y {
						if index == 0 {
							fmt.Print("⏿")
						} else {
							fmt.Print("x")
						}
					}
				}
			}

			fmt.Print("⠀")
		}

		fmt.Println()
	}
}

func (cli *SnakeCli) InitializeSnake() {
	var startPointX int32 = 3
	var startPointY int32 = 5
	snakeLength := 3

	snejk := snakeModule.MakeSnake(startPointX, startPointY, int32(snakeLength))

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
	topLeft := Point{
		X: 0,
		Y: 5,
	}

	bottomRight := Point{
		X: 20,
		Y: 0,
	}
	board := Board{
		TopLeft:     topLeft,
		BottomRight: bottomRight,
	}
	cli.board = &board
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func renderGame(snake *snakeModule.Snake, board *Board) {
	clearScreen()
	array := make([]*snakeModule.Snake, 1)
	array[0] = snake
	// append(array, snake)

	board.Print(array)
	// snake.Print()
	// snakeModule.PrintTable(table, snake)
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
			renderGame(cli.snake, cli.board)
		}
	}()

	go func() {
		for {
			exitWhenDone(isDone)
			<-cli.movementTicket.C
			cli.snake.Move(cli.lastInput)
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
