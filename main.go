package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"sneakgame.com/constants"
	keylistener "sneakgame.com/keylistener"
	snakeModule "sneakgame.com/snake"

	"golang.design/x/hotkey/mainthread"
)

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

func main() {
	var lastInput = constants.Down
	doneChannel := make(chan bool, 1)
	keyboardInputChannel := make(chan uint8, 10)

	snejk := snakeModule.MakeSnake(3, 5, 3)
	table := make([][]uint8, 10)
	for i := range table {
		table[i] = make([]uint8, 10)
	}

	clearScreen()

	snakeModule.PrintTable(&table, snejk)

	ticker := time.NewTicker(1000 * time.Millisecond)

	go func() {
		for {
			<-ticker.C
			renderGame(snejk, &table, lastInput)
		}
	}()

	go func() {
		for {
			<-doneChannel
			fmt.Println("Done!")
		}
	}()

	go func() {
		for {
			lastInput = <-keyboardInputChannel
			fmt.Println("Key pressed:", lastInput)
			renderGame(snejk, &table, lastInput)
		}
	}()

	mainthread.Init(keylistener.GetKeyListener(doneChannel, keyboardInputChannel))
}
