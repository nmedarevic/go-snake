package constants

const Up uint8 = 1
const Down uint8 = 2
const Left uint8 = 3
const Right uint8 = 4
const Space uint8 = 5
const Escape uint8 = 27
const Q uint8 = 113

func IsIllegalMovementKey(key uint8, previousKey uint8) bool {
	if key == Up && previousKey == Down {
		return true
	}

	if key == Down && previousKey == Up {
		return true
	}

	if key == Left && previousKey == Right {
		return true
	}

	if key == Right && previousKey == Left {
		return true
	}

	return false
}
