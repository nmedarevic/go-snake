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

func MakeSnake() *Snake {
	snakeBody := make([]SnakePoint, 3)
	snakeBody[0] = SnakePoint{X: uint8(3), Y: uint8(5)}
	snakeBody[1] = SnakePoint{X: uint8(3), Y: uint8(4)}
	snakeBody[2] = SnakePoint{X: uint8(3), Y: uint8(3)}

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
		// (*snake.Body)[0].Y = (*snake.Body)[0].Y + 1
	}

	if direction == Up {
		nextPositionY = (*snake.Body)[0].Y - 1
		changedY = true
		// (*snake.Body)[0].Y = (*snake.Body)[0].Y - 1
	}

	if direction == Left {
		nextPositionX = (*snake.Body)[0].X + 1
		changedX = true
		// (*snake.Body)[0].X = (*snake.Body)[0].X + 1
	}

	if direction == Right {
		nextPositionX = (*snake.Body)[0].X - 1
		changedX = true
		// (*snake.Body)[0].X = (*snake.Body)[0].X - 1
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

func Move(snake *Snake, field *[][]uint8, direction uint8) {
	MoveHead(snake, field, direction)

	// Moves the body
	for i := 1; i < len(*snake.Body); i++ {
		fmt.Println(i, "Curr", (*snake.Body)[i])
		fmt.Println(i-1, "Prev", (*snake.Body)[i-1])

		// When the head is going up from heading horizontally
		if (*snake.Body)[i].X+1 == (*snake.Body)[i-1].X && (*snake.Body)[i].Y != (*snake.Body)[i-1].Y {
			(*snake.Body)[i].X = (*snake.Body)[i].X + 1

			continue
		}

		// Moving down
		if (*snake.Body)[i].X == (*snake.Body)[i-1].X && (*snake.Body)[i].Y+2 == (*snake.Body)[i-1].Y {
			(*snake.Body)[i].Y = (*snake.Body)[i].Y + 1

			continue
		}
	}
}

func PrintTable(table *[][]uint8, snake *Snake) {
	fmt.Println(snake.Body)
	for y := range *table {
		for x := range (*table)[y] {
			(*table)[y][x] = uint8(0)
			for index, s := range *snake.Body {
				if uint8(x) == s.X && uint8(y) == s.Y {
					// (*table)[y][x] = uint8(1)

					if index == 0 {
						// (*table)[y][x] = uint8(2)

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