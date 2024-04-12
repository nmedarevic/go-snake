package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
)

type SneakPoint struct {
	X uint8
	Y uint8
}

func printTable(table *[][]uint8, sneak *[]SneakPoint) {
	for x := range *table {
		for y := range (*table)[x] {
			for index, s := range *sneak {
				if uint8(x) == s.X && uint8(y) == s.Y {
					(*table)[x][y] = uint8(1)

					if index == 0 {
						(*table)[x][y] = uint8(2)

						fmt.Print("S")
						continue
					}
					fmt.Print("N")
				}
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func sneakGame() {
	sneak := make([]SneakPoint, 2)
	sneak[0] = SneakPoint{X: uint8(5), Y: uint8(5)}
	sneak[1] = SneakPoint{X: uint8(4), Y: uint8(5)}

	table := make([][]uint8, 10)
	for i := range table {
		table[i] = make([]uint8, 10)
	}

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	printTable(&table, &sneak)
}

func f(from string) {
	for {
		fmt.Println(":")
	}
}

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool, 1)
	keyboardInput := make(chan uint8, 10)

	go func() {
		for {
			<-done
			fmt.Println("Done!")
		}
	}()

	go func() {
		for {
			input := <-keyboardInput
			fmt.Println("Key pressed:", input)
		}
	}()

	go func() {
		for {
			t := <-ticker.C
			fmt.Println("Tick at", t)
		}
	}()

	mainthread.Init(getFn(done, keyboardInput))
}

const Up = 1
const Down = 2
const Left = 3
const Right = 4

func listenToKey(key *hotkey.Hotkey, input chan uint8, done chan bool) {
	for {
		<-key.Keydown()
		hexKeyVal, _ := strconv.Atoi(key.String())

		log.Printf("[Listener] Pressed: %x", hexKeyVal)

		if hexKeyVal == int(hotkey.KeyUp) {
			input <- Up
		}
		if hexKeyVal == int(hotkey.KeyDown) {
			input <- Down
		}
		if hexKeyVal == int(hotkey.KeyLeft) {
			input <- Left
		}
		if hexKeyVal == int(hotkey.KeyRight) {
			input <- Right
		}
		if key.String() == "X" {
			done <- true
		}
	}
}

func getFn(done chan bool, keyboardInput chan uint8) func() {
	return func() {
		keyUp := hotkey.New([]hotkey.Modifier{}, hotkey.KeyUp)
		keyDown := hotkey.New([]hotkey.Modifier{}, hotkey.KeyDown)
		keyLeft := hotkey.New([]hotkey.Modifier{}, hotkey.KeyLeft)
		keyRight := hotkey.New([]hotkey.Modifier{}, hotkey.KeyRight)
		keyClose := hotkey.New([]hotkey.Modifier{}, hotkey.KeyX)

		err1 := keyUp.Register()
		err2 := keyDown.Register()
		err3 := keyLeft.Register()
		err4 := keyRight.Register()
		err5 := keyClose.Register()

		if err1 != nil && err2 != nil && err3 != nil && err4 != nil && err5 != nil {
			log.Fatalf("hotkey: failed to register hotkey")
		}

		go listenToKey(keyUp, keyboardInput, done)
		go listenToKey(keyDown, keyboardInput, done)
		go listenToKey(keyRight, keyboardInput, done)
		go listenToKey(keyLeft, keyboardInput, done)
		go listenToKey(keyClose, keyboardInput, done)

		for {
			<-done
			log.Print("Unregistering all listeners")
			keyUp.Unregister()
			keyDown.Unregister()
			keyLeft.Unregister()
			keyRight.Unregister()
			keyClose.Unregister()
		}
	}
}
