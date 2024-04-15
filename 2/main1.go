package snake

// import (
// 	"fmt"
// 	"time"
// )

// func printTable(table *[][]uint8, snake *Snake) {
// 	fmt.Println(snake.Body)
// 	for y := range *table {
// 		for x := range (*table)[y] {
// 			(*table)[y][x] = uint8(0)
// 			for index, s := range *snake.Body {
// 				if uint8(x) == s.X && uint8(y) == s.Y {
// 					// (*table)[y][x] = uint8(1)

// 					if index == 0 {
// 						// (*table)[y][x] = uint8(2)

// 						fmt.Print("S")
// 						continue
// 					}
// 					fmt.Print("N")
// 				}
// 			}
// 			fmt.Print(" ")
// 		}
// 		fmt.Println()
// 	}
// }

// type MoveCommand struct {
// 	Direction uint8
// }

// // type Snake struct {
// // 	Body *[]SneakPoint
// // }

// // const Up = 1
// // const Down = 2
// // const Left = 3
// // const Right = 4

// // func Move(snake *Snake, field *[][]uint8, direction uint8) {
// // 	if direction == Down {
// // 		(*snake.Body)[0].Y = (*snake.Body)[0].Y + 1
// // 	}

// // 	if direction == Left {
// // 		(*snake.Body)[0].X = (*snake.Body)[0].X + 1
// // 	}

// // 	if direction == Up {
// // 		(*snake.Body)[0].Y = (*snake.Body)[0].Y - 1
// // 	}

// // 	if direction == Right {
// // 		(*snake.Body)[0].X = (*snake.Body)[0].X - 1
// // 	}
// // 	if (*snake.Body)[0].Y >= uint8(len(*field)) {
// // 		(*snake.Body)[0].Y = 0
// // 	}
// // 	if (*snake.Body)[0].X >= uint8(len((*field)[0])) {
// // 		(*snake.Body)[0].X = 0
// // 	}
// // 	fmt.Print("XXXXX")
// // 	// Moves the body
// // 	for i := 1; i < len(*snake.Body); i++ {
// // 		fmt.Println("Curr", (*snake.Body)[i])
// // 		fmt.Println("Prev", (*snake.Body)[i-1])
// // 		// When the head is going up from heading horizontally
// // 		if (*snake.Body)[i].X+1 == (*snake.Body)[i-1].X && (*snake.Body)[i].Y != (*snake.Body)[i-1].Y {
// // 			(*snake.Body)[i].X = (*snake.Body)[i].X + 1

// // 			continue
// // 		}

// // 		// Moving down
// // 		if (*snake.Body)[i].X == (*snake.Body)[i-1].X && (*snake.Body)[i].Y+2 == (*snake.Body)[i-1].Y {
// // 			(*snake.Body)[i].Y = (*snake.Body)[i].Y + 1

// // 			continue
// // 		}
// // 	}
// // }

// func moveSnake(table *[][]uint8, s *Snake) {
// 	for i := 0; i < len(*s.Body); i++ {
// 		(*s.Body)[i].Y = (*s.Body)[i].Y + 1
// 		if (*s.Body)[i].Y >= uint8(len(*table)) {
// 			(*s.Body)[i].Y = 0
// 		}

// 		(*s.Body)[i].X = (*s.Body)[i].X + 1
// 		if (*s.Body)[i].X >= uint8(len((*table)[0])) {
// 			(*s.Body)[i].X = 0
// 		}
// 	}
// }

// func sneakGame() {
// 	snakeBody := make([]SnakePoint, 3)
// 	snakeBody[0] = SnakePoint{X: uint8(3), Y: uint8(5)}
// 	snakeBody[1] = SnakePoint{X: uint8(3), Y: uint8(3)}
// 	snakeBody[2] = SnakePoint{X: uint8(3), Y: uint8(4)}

// 	snake := Snake{
// 		Body: &snakeBody,
// 	}

// 	table := make([][]uint8, 10)
// 	for i := range table {
// 		table[i] = make([]uint8, 10)
// 	}

// 	ticker := time.NewTicker(500 * time.Millisecond)
// 	tickerSnakeChange := time.NewTicker(500 * time.Millisecond)
// 	done := make(chan bool)

// 	go func() {
// 		for {
// 			select {
// 			case <-done:
// 				return
// 			case <-ticker.C:
// 				// c := exec.Command("clear")
// 				// c.Stdout = os.Stdout
// 				// c.Run()
// 				printTable(&table, &snake)
// 				dt := time.Now()
// 				fmt.Println("AAAA", dt.UnixMilli())
// 			}
// 		}
// 	}()

// 	go func() {
// 		for {
// 			<-tickerSnakeChange.C
// 			fmt.Println("MOVES")
// 			fmt.Println("MOVES")
// 			fmt.Println("MOVES")
// 			// moveSnake(&table, &snake)

// 			Move(&snake, &table, Down)
// 		}
// 	}()

// 	for {
// 		<-done
// 	}
// }

// func main() {
// 	sneakGame()
// }
