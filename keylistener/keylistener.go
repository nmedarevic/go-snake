package keylistener

import (
	"log"
	"strconv"

	"golang.design/x/hotkey"
	constants "sneakgame.com/constants"
)

func GetKeyListener(doneChannel chan bool, keyboardInputChannel chan uint8) func() {
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

		go listenToKey(keyUp, keyboardInputChannel, doneChannel)
		go listenToKey(keyDown, keyboardInputChannel, doneChannel)
		go listenToKey(keyRight, keyboardInputChannel, doneChannel)
		go listenToKey(keyLeft, keyboardInputChannel, doneChannel)
		go listenToKey(keyClose, keyboardInputChannel, doneChannel)

		for {
			<-doneChannel
			log.Print("Unregistering all listeners")
			keyUp.Unregister()
			keyDown.Unregister()
			keyLeft.Unregister()
			keyRight.Unregister()
			keyClose.Unregister()
		}
	}
}

func listenToKey(key *hotkey.Hotkey, input chan uint8, doneChannel chan bool) {
	for {
		<-key.Keydown()
		hexKeyVal, _ := strconv.Atoi(key.String())

		if hexKeyVal == int(hotkey.KeyUp) {
			input <- constants.Up
		}
		if hexKeyVal == int(hotkey.KeyDown) {
			input <- constants.Down
		}
		if hexKeyVal == int(hotkey.KeyLeft) {
			input <- constants.Left
		}
		if hexKeyVal == int(hotkey.KeyRight) {
			input <- constants.Right
		}
		if key.String() == "X" {
			doneChannel <- true
		}
	}
}
