package snake

import "math/rand"

func MakeSnake(x int32, y int32, length int32) *Snake {
	snakeBody := make([]SnakePoint, length)
	snakeBody[0] = SnakePoint{X: x, Y: y}
	snakeBody[1] = SnakePoint{X: x, Y: int32(y - 1)}
	snakeBody[2] = SnakePoint{X: x, Y: int32(y - 2)}

	snake := Snake{
		Body: &snakeBody,
	}

	return &snake
}

func MakeSnakeHeadPointUp(x int32, y int32, length int32) *Snake {
	snakeBody := make([]SnakePoint, length)
	for i := 0; i < int(length); i++ {
		snakeBody[i] = SnakePoint{X: x, Y: y + int32(i)}
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

func MakeSnakePointingUp(x int32, y int32, length int32) *Snake {
	snakeBody := make([]SnakePoint, length)

	for i := 0; i < int(length); i++ {
		snakeBody[i] = SnakePoint{X: int32(x), Y: int32(y + int32(i))}
	}

	snake := Snake{
		Body: &snakeBody,
		Id:   randSeq(10),
	}

	return &snake
}

func MakeSnakePointingRight(x int32, y int32, length int32) *Snake {
	snakeBody := make([]SnakePoint, length)

	for i := 0; i < int(length); i++ {
		snakeBody[i] = SnakePoint{X: int32(x) + int32(i), Y: int32(y)}
	}

	snake := Snake{
		Body: &snakeBody,
		Id:   randSeq(10),
	}

	return &snake
}
