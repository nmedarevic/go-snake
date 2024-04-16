package main

import (
	"fmt"
	"os"
	"os/exec"

	keylistener "sneakgame.com/keylistener"
	snakeModule "sneakgame.com/snake"

	"golang.design/x/hotkey/mainthread"
)

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func main() {
	doneChannel := make(chan bool, 1)
	keyboardInputChannel := make(chan uint8, 10)

	snejk := snakeModule.MakeSnake(3, 5, 3)
	table := make([][]uint8, 10)
	for i := range table {
		table[i] = make([]uint8, 10)
	}

	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	snakeModule.PrintTable(&table, snejk)

	go func() {
		for {
			<-doneChannel
			fmt.Println("Done!")
		}
	}()

	go func() {
		for {
			input := <-keyboardInputChannel

			clearScreen()
			snakeModule.MoveSnake(snejk, &table, input)
			snakeModule.PrintTable(&table, snejk)
			fmt.Println("Key pressed:", input)
		}
	}()

	mainthread.Init(keylistener.GetKeyListener(doneChannel, keyboardInputChannel))
}
