package snake

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
	for i := 0; i < len(snakeBody); i++ {
		snakeBody[i] = SnakePoint{X: uint8(x), Y: uint8(y + uint8(i))}
	}

	snake := Snake{
		Body: &snakeBody,
	}

	return &snake
}
