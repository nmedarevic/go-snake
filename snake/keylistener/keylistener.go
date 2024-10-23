package keylistener

import (
	"log"
	"strconv"

	"golang.design/x/hotkey"
	constants "snakegame.com/constants"
)

func GetKeyListener(doneChannel chan bool, keyboardInputChannel chan uint8) func() {
	usedKeys := []hotkey.Key{
		hotkey.KeyUp,
		hotkey.KeyDown,
		hotkey.KeyLeft,
		hotkey.KeyRight,
		hotkey.KeyX,
		hotkey.KeySpace,
		hotkey.KeyEscape,
		hotkey.KeyQ,
	}

	return func() {
		keyRegistrations := make([]*hotkey.Hotkey, 0)

		for _, value := range usedKeys {
			newHotkeyPointer := hotkey.New([]hotkey.Modifier{}, value)

			keyRegistrations = append(keyRegistrations, newHotkeyPointer)
		}

		for _, hotkeyPointer := range keyRegistrations {
			error := hotkeyPointer.Register()

			if error != nil {
				log.Fatalf("hotkey: failed to register hotkey")
			}

			go listenToKey(hotkeyPointer, keyboardInputChannel, doneChannel)
		}

		for {
			<-doneChannel
			log.Print("Unregistering all listeners")

			for _, hotkeyPointer := range keyRegistrations {
				hotkeyPointer.Unregister()
			}
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
		if hexKeyVal == int(hotkey.KeySpace) {
			input <- constants.Space
		}

		// fmt.Println("Key pressed:", hexKeyVal, key.String(), int(hotkey.KeyQ), int(hotkey.KeyEscape))

		if key.String() == "X" || hexKeyVal == int(hotkey.KeyQ) || hexKeyVal == int(hotkey.KeyEscape) {
			doneChannel <- true
		}
	}
}
