package main

import (
	"fmt"
	"log"
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

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Done!")
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()
	// go f("goroutine")
	mainthread.Init(getFn(done))
} // Not necessary when use in Fyne, Ebiten or Gio.

func getFn(done chan bool) func() {
	return func() {
		hk := hotkey.New([]hotkey.Modifier{}, hotkey.KeyS)
		err := hk.Register()

		if err != nil {
			log.Fatalf("hotkey: failed to register hotkey: %v", err)
		}

		for {
			<-hk.Keydown()
			log.Printf("hotkey: %v is registered\n", hk)
			done <- true
			break
		}
		// log.Printf("hotkey: %v is down\n", hk)
		// <-hk.Keyup()
		// log.Printf("hotkey: %v is up\n", hk)
		hk.Unregister()
		// log.Printf("hotkey: %v is unregistered\n", hk)
	}
}
