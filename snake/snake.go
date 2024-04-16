package snake

import "fmt"

type SnakePoint struct {
	X uint8
	Y uint8
}

type Snake struct {
	Body *[]SnakePoint
}

const Up = 1
const Down = 2
const Left = 3
const Right = 4

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
	snakeBody[0] = SnakePoint{X: uint8(x), Y: uint8(y)}
	snakeBody[1] = SnakePoint{X: uint8(x), Y: uint8(y + 1)}
	snakeBody[2] = SnakePoint{X: uint8(x), Y: uint8(y + 2)}

	snake := Snake{
		Body: &snakeBody,
	}

	return &snake
}

func MoveHead(snake *Snake, field *[][]uint8, direction uint8) {
	var changedX bool = false
	var changedY bool = false
	var nextPositionX uint8
	var nextPositionY uint8

	if direction == Down {
		nextPositionY = (*snake.Body)[0].Y + 1
		changedY = true
	}

	if direction == Up {
		nextPositionY = (*snake.Body)[0].Y - 1
		changedY = true
	}

	if direction == Left {
		nextPositionX = (*snake.Body)[0].X - 1
		changedX = true
	}

	if direction == Right {
		nextPositionX = (*snake.Body)[0].X + 1
		changedX = true
	}

	// Move the head only if it does not overlap with the second item
	if changedX && (*snake.Body)[0].X != nextPositionX && nextPositionX != (*snake.Body)[1].X {
		(*snake.Body)[0].X = nextPositionX
	}

	// Move the head only if it does not overlap with the second item
	if changedY && (*snake.Body)[0].Y != nextPositionY && nextPositionY != (*snake.Body)[1].Y {
		(*snake.Body)[0].Y = nextPositionY
	}

	if (*snake.Body)[0].Y >= uint8(len(*field)) {
		(*snake.Body)[0].Y = 0
	}

	if (*snake.Body)[0].X >= uint8(len((*field)[0])) {
		(*snake.Body)[0].X = 0
	}
}

func MoveBody(snake *Snake, field *[][]uint8, direction uint8) {
	for i := 1; i < len(*snake.Body); i++ {
		// fmt.Println(i, "Curr", (*snake.Body)[i])
		// fmt.Println(i-1, "Prev", (*snake.Body)[i-1])

		// Cell in front moved right
		if (*snake.Body)[i].X+2 == (*snake.Body)[i-1].X && (*snake.Body)[i].Y == (*snake.Body)[i-1].Y {
			(*snake.Body)[i].X = (*snake.Body)[i].X + 1
			continue
		}

		// Cell in front moved left
		if (*snake.Body)[i].X-2 == (*snake.Body)[i-1].X && (*snake.Body)[i].Y == (*snake.Body)[i-1].Y {
			(*snake.Body)[i].X = (*snake.Body)[i].X - 1
			continue
		}

		// Cell on the left hand side moved up
		if i+1 == len(*snake.Body)-1 && (*snake.Body)[i].X-1 == (*snake.Body)[i-1].X && (*snake.Body)[i].Y-1 == (*snake.Body)[i-1].Y && (*snake.Body)[i].X+1 == (*snake.Body)[i+1].X {
			(*snake.Body)[i].X = (*snake.Body)[i].X - 1
			continue
		}

		// Cell moved up but next one is horizontal
		if i+1 == len(*snake.Body)-1 && (*snake.Body)[i].X == (*snake.Body)[i-1].X && (*snake.Body)[i].Y-2 == (*snake.Body)[i-1].Y && (*snake.Body)[i].X+1 == (*snake.Body)[i+1].X {
			(*snake.Body)[i].Y = (*snake.Body)[i].Y - 1
			continue
		}

		// Cell in front moved right down
		if (*snake.Body)[i].X+1 == (*snake.Body)[i-1].X && (*snake.Body)[i].Y+1 == (*snake.Body)[i-1].Y {
			(*snake.Body)[i].Y = (*snake.Body)[i].Y + 1
			continue
		}

		// Cell in front moved down left
		if (*snake.Body)[i].X-1 == (*snake.Body)[i-1].X && (*snake.Body)[i].Y+1 == (*snake.Body)[i-1].Y {
			(*snake.Body)[i].Y = (*snake.Body)[i].Y + 1
			continue
		}

		// Cell in front moved down
		if (*snake.Body)[i].X == (*snake.Body)[i-1].X && (*snake.Body)[i].Y+2 == (*snake.Body)[i-1].Y {
			(*snake.Body)[i].Y = (*snake.Body)[i].Y + 1
			continue
		}

		// Cell in front moved top left
		if (*snake.Body)[i].X-1 == (*snake.Body)[i-1].X && (*snake.Body)[i].Y-1 == (*snake.Body)[i-1].Y {
			(*snake.Body)[i].Y = (*snake.Body)[i].Y - 1
			continue
		}

		// Cell in front moved top right
		if (*snake.Body)[i].X+1 == (*snake.Body)[i-1].X && (*snake.Body)[i].Y-1 == (*snake.Body)[i-1].Y {
			(*snake.Body)[i].Y = (*snake.Body)[i].Y - 1
			continue
		}

		// Cell in front moved up
		if (*snake.Body)[i].X == (*snake.Body)[i-1].X && (*snake.Body)[i].Y-2 == (*snake.Body)[i-1].Y {
			(*snake.Body)[i].Y = (*snake.Body)[i].Y - 1
			continue
		}
	}
}

func Move(snake *Snake, field *[][]uint8, direction uint8) {
	MoveHead(snake, field, direction)

	MoveBody(snake, field, direction)
}

func PrintTable(table *[][]uint8, snake *Snake) {
	fmt.Println(snake.Body)
	for y := range *table {
		for x := range (*table)[y] {
			(*table)[y][x] = uint8(0)
			for index, s := range *snake.Body {
				if uint8(x) == s.X && uint8(y) == s.Y {
					if index == 0 {
						(*table)[y][x] = uint8(1)
						continue
					}
					(*table)[y][x] = uint8(2)
					continue
				}
			}
		}
	}

	for y := range *table {
		for x := range (*table)[y] {
			if (*table)[y][x] == 1 {
				fmt.Print("H")
				continue
			}
			if (*table)[y][x] == 2 {
				fmt.Print("B")
				continue
			}
			fmt.Print("_")
		}
		fmt.Println()
	}
}
