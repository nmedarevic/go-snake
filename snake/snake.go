package snake

import (
	"fmt"

	constants "sneakgame.com/constants"
)

type SnakePoint struct {
	X uint8
	Y uint8
}

type Snake struct {
	Body *[]SnakePoint
	Id   string
}

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

		if direction == constants.Down {
			if i+1 < len(*snake.Body) && (*snake.Body)[i].Y+1 == (*snake.Body)[i+1].Y {
				break
			}

			(*snake.Body)[i].Y = (*snake.Body)[i].Y + 1
		}

		if direction == constants.Up {
			if i+1 < len(*snake.Body) && (*snake.Body)[i].Y-1 == (*snake.Body)[i+1].Y {
				break
			}

			(*snake.Body)[i].Y = (*snake.Body)[i].Y - 1
		}

		if direction == constants.Left {
			if i+1 < len(*snake.Body) && (*snake.Body)[i].X-1 == (*snake.Body)[i+1].X {
				break
			}

			(*snake.Body)[i].X = (*snake.Body)[i].X - 1
		}

		if direction == constants.Right {
			if i+1 < len(*snake.Body) && (*snake.Body)[i].X+1 == (*snake.Body)[i+1].X {
				break
			}

			(*snake.Body)[i].X = (*snake.Body)[i].X + 1
		}

		if (*snake.Body)[i].Y >= uint8(len(*field)) {
			(*snake.Body)[i].Y = 0
		}

		// When snake touches the left wall
		if (*snake.Body)[i].Y == uint8(0) && previousPoint.Y == uint8(0) {
			(*snake.Body)[i].Y = uint8(len(*field) - 1)
		}

		if (*snake.Body)[i].X >= uint8(len((*field)[0])) {
			(*snake.Body)[i].X = 0
		}

		// When snake touches the right wall
		if (*snake.Body)[i].X == uint8(0) && previousPoint.X == uint8(0) {
			(*snake.Body)[i].X = uint8(len((*field)[0]) - 1)
		}
	}
}

func PrintTable(table *[][]uint8, snake *Snake) {
	fmt.Println(snake.Body)

	for y := range *table {
	loop:
		for x := range (*table)[y] {
			for index, s := range *snake.Body {
				if uint8(x) == s.X && uint8(y) == s.Y {

					if index == 0 {
						fmt.Print("⏿")
						continue loop
					}

					fmt.Print("x")

					continue loop
				}
			}
			fmt.Print("⠀")
		}
		fmt.Println()
	}
}

func RenderGameToTerminal(table *[][]uint8, snakes [](*Snake)) {
	for y := range *table {
	loopSnakes:
		for x := range (*table)[y] {
			for _, snake := range snakes {

				for index, s := range *snake.Body {
					if uint8(x) == s.X && uint8(y) == s.Y {

						if index == 0 {
							fmt.Print("⏿")
							continue loopSnakes
						}

						fmt.Print("x")

						continue loopSnakes
					}
				}
			}
			fmt.Print("⠀")
		}
		fmt.Println()
	}
}

func IsHeadTouchingOtherSnake(snake1 *Snake, snake2 *Snake) bool {
	var head = (*(*snake1).Body)[0]

	for i := 0; i < len((*(*snake2).Body)); i++ {
		snakePoint := (*(*snake2).Body)[i]

		if head.X == snakePoint.X && head.Y == snakePoint.Y {
			return true
		}
	}

	return false
}

func HandleSnakeContact(snake1 *Snake, snake2 *Snake) {
	var head = (*(*snake1).Body)[0]

	for i := 0; i < len((*(*snake2).Body)); i++ {
		snakePoint := (*(*snake2).Body)[i]

		if head.X == snakePoint.X && head.Y == snakePoint.Y {
			*((*snake2).Body) = (*(*snake2).Body)[0:i]
			return
		}
	}
}

func AreSnakesTouching(snakes [](*Snake)) bool {
	for _, snake := range snakes {
		var head = (*(*snake).Body)[0]
		var id = (*snake).Id

		for _, snake2 := range snakes {
			if (*snake2).Id == id {
				continue
			}

			// var otherHead = (*(*snake).Body)[0]

			// if heads touch, longer snake winds

			// if current snake is in other snake's body
			// make other snake smaller
			// var shouldDelete = false
			for i := 0; i < len((*(*snake).Body)); i++ {
				snakePoint := (*(*snake).Body)[i]

				if head.X == snakePoint.X && head.Y == snakePoint.Y {
					return true
				}
			}
		}
	}

	return false
}
