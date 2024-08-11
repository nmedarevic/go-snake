package snake

import "math/rand"

func MakeSnake(x uint8, y uint8, length uint8) *Snake {
	snakeBody := make([]SnakePoint, length)
	snakeBody[0] = SnakePoint{X: uint8(x), Y: uint8(y)}
	snakeBody[1] = SnakePoint{X: uint8(x), Y: uint8(y - 1)}
	snakeBody[2] = SnakePoint{X: uint8(x), Y: uint8(y - 2)}

	snake := Snake{
		Body: &snakeBody,
	}

	return &snake
}

func MakeSnakeHeadPointUp(x uint8, y uint8, length uint8) *Snake {
	snakeBody := make([]SnakePoint, length)
	for i := 0; i < int(length); i++ {
		snakeBody[i] = SnakePoint{X: uint8(x), Y: uint8(y + uint8(i))}
	}

	snake := Snake{
		Body: &snakeBody,
		Id:   randSeq(10),
	}

	return &snake
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func MakeSnakePointingUp(x uint8, y uint8, length uint8) *Snake {
	snakeBody := make([]SnakePoint, length)

	for i := 0; i < int(length); i++ {
		snakeBody[i] = SnakePoint{X: uint8(x), Y: uint8(y + uint8(i))}
	}

	snake := Snake{
		Body: &snakeBody,
		Id:   randSeq(10),
	}

	return &snake
}

func MakeSnakePointingRight(x uint8, y uint8, length uint8) *Snake {
	snakeBody := make([]SnakePoint, length)

	for i := 0; i < int(length); i++ {
		snakeBody[i] = SnakePoint{X: uint8(x) + uint8(i), Y: uint8(y)}
	}

	snake := Snake{
		Body: &snakeBody,
		Id:   randSeq(10),
	}

	return &snake
}
