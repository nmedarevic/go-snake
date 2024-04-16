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

func MoveSnake(snake *Snake, field *[][]uint8, direction uint8) {
	var previousPoint *SnakePoint

	for i := 0; i < len(*snake.Body); i++ {
		if previousPoint != nil {
			var temporaryPoint = SnakePoint{X: (*snake.Body)[i].X, Y: (*snake.Body)[i].Y}

			(*snake.Body)[i].X = previousPoint.X
			(*snake.Body)[i].Y = previousPoint.Y

			previousPoint = &temporaryPoint

			continue
		}

		if previousPoint == nil {
			previousPoint = &SnakePoint{X: (*snake.Body)[i].X, Y: (*snake.Body)[i].Y}
		}

		if direction == Down {
			if i+1 < len(*snake.Body) && (*snake.Body)[i].Y+1 == (*snake.Body)[i+1].Y {
				break
			}

			(*snake.Body)[i].Y = (*snake.Body)[i].Y + 1
		}

		if direction == Up {
			if i+1 < len(*snake.Body) && (*snake.Body)[i].Y-1 == (*snake.Body)[i+1].Y {
				break
			}

			(*snake.Body)[i].Y = (*snake.Body)[i].Y - 1
		}

		if direction == Left {
			if i+1 < len(*snake.Body) && (*snake.Body)[i].X-1 == (*snake.Body)[i+1].X {
				break
			}

			(*snake.Body)[i].X = (*snake.Body)[i].X - 1
		}

		if direction == Right {
			if i+1 < len(*snake.Body) && (*snake.Body)[i].X+1 == (*snake.Body)[i+1].X {
				break
			}

			(*snake.Body)[i].X = (*snake.Body)[i].X + 1
		}

		if (*snake.Body)[i].Y >= uint8(len(*field)) {
			(*snake.Body)[i].Y = 0
		}

		if (*snake.Body)[i].X >= uint8(len((*field)[0])) {
			(*snake.Body)[i].X = 0
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
					if index == 0 {
						(*table)[y][x] = uint8(1)
						continue
					}
					(*table)[y][x] = uint8(2)
					continue
				}
			}
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
