package snake

import "math/rand"

func MakeSnake(x int32, y int32, length int32) *Snake {
	snakeBody := make([]Point, length)
	snakeBody[0] = Point{X: x, Y: y}
	snakeBody[1] = Point{X: x, Y: int32(y - 1)}
	snakeBody[2] = Point{X: x, Y: int32(y - 2)}

	snake := Snake{
		Body: &snakeBody,
	}

	return &snake
}

func MakeSnakeHeadPointUp(x int32, y int32, length int32) *Snake {
	snakeBody := make([]Point, length)
	for i := 0; i < int(length); i++ {
		snakeBody[i] = Point{X: x, Y: y + int32(i)}
	}

	snake := Snake{
		Body: &snakeBody,
		Id:   RandSeq(10),
	}

	return &snake
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func MakePointingUp(x int32, y int32, length int32) *Snake {
	snakeBody := make([]Point, length)

	for i := 0; i < int(length); i++ {
		snakeBody[i] = Point{X: int32(x), Y: int32(y + int32(i))}
	}

	snake := Snake{
		Body: &snakeBody,
		Id:   RandSeq(10),
	}

	return &snake
}

func MakePointingRight(x int32, y int32, length int32) *Snake {
	snakeBody := make([]Point, length)

	for i := 0; i < int(length); i++ {
		snakeBody[i] = Point{X: int32(x) + int32(i), Y: int32(y)}
	}

	snake := Snake{
		Body: &snakeBody,
		Id:   RandSeq(10),
	}

	return &snake
}
