package snake

import (
	"fmt"

	constants "snakegame.com/constants"
)

type SnakePoint struct {
	X int32
	Y int32
}

type Snake struct {
	Body *[]SnakePoint
	Id   string
}

type EnclosingRectangle struct {
	Xmin int32
	Xmax int32
	Ymin int32
	Ymax int32
}

func (snake *Snake) Move(direction uint8) {
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
			if i+1 < len(*snake.Body) && (*snake.Body)[i].Y-1 == (*snake.Body)[i+1].Y {
				break
			}

			(*snake.Body)[i].Y = (*snake.Body)[i].Y - 1
		}

		if direction == constants.Up {
			if i+1 < len(*snake.Body) && (*snake.Body)[i].Y+1 == (*snake.Body)[i+1].Y {
				break
			}

			(*snake.Body)[i].Y = (*snake.Body)[i].Y + 1
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
	}
}

func findEnclosingRectangle(snake *Snake) *EnclosingRectangle {
	var rectangleXMax int32 = 0
	var rectangleXMin int32 = 2147483647
	var rectangleYMax int32 = 0
	var rectangleYMin int32 = 2147483647

	for _, point := range *snake.Body {
		if point.X > rectangleXMax {
			rectangleXMax = point.X
		}

		if point.X < rectangleXMin {
			rectangleXMin = point.X
		}

		if point.Y > rectangleYMax {
			rectangleYMax = point.Y
		}

		if point.Y < rectangleYMin {
			rectangleYMin = point.Y
		}
	}

	var rectangle = EnclosingRectangle{
		Xmin: rectangleXMin,
		Xmax: rectangleXMax,
		Ymin: rectangleYMin,
		Ymax: rectangleYMax,
	}

	return &rectangle
}

func (snake *Snake) Print() {
	var enclosingRectangle = findEnclosingRectangle(snake)
	fmt.Println(enclosingRectangle)
	fmt.Println(snake.Body)
	for y := int(enclosingRectangle.Ymax) - int(enclosingRectangle.Ymin); y >= 0; y-- {
	loop:
		for x := 0; x <= int(enclosingRectangle.Xmax)-int(enclosingRectangle.Xmin); x++ {
			for index, s := range *snake.Body {
				if int32(x) == s.X-enclosingRectangle.Xmin && int32(y) == s.Y-enclosingRectangle.Ymin {
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

// func RenderGameToTerminal(table *[][]uint8, snakes [](*Snake)) {
// 	for y := range *table {
// 	loopSnakes:
// 		for x := range (*table)[y] {
// 			for _, snake := range snakes {

// 				for index, s := range *snake.Body {
// 					if uint8(x) == s.X && uint8(y) == s.Y {

// 						if index == 0 {
// 							fmt.Print("⏿")
// 							continue loopSnakes
// 						}

// 						fmt.Print("x")

// 						continue loopSnakes
// 					}
// 				}
// 			}
// 			fmt.Print("⠀")
// 		}
// 		fmt.Println()
// 	}
// }

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
