package snake

import (
	"testing"

	"sneakgame.com/constants"
)

func TestShouldNotMoveHeadBackwards(t *testing.T) {
	originalSnake := MakeSnake(3, 5, 3)
	snake := MakeSnake(3, 5, 3)

	table := make([][]uint8, 10)
	for i := range table {
		table[i] = make([]uint8, 10)
	}

	snake.Move((constants.Up))
	snake.Print()

	for i := 0; i < len(*(originalSnake.Body)); i++ {
		if (*(originalSnake.Body))[i].X != (*(snake.Body))[i].X || (*(originalSnake.Body))[i].Y != (*snake.Body)[i].Y {
			t.Errorf("Snake head moved backwards")
		}
	}
}

func TestShouldMoveHeadBottomRight(t *testing.T) {
	snake := MakeSnake(3, 5, 3)

	table := make([][]uint8, 10)
	for i := range table {
		table[i] = make([]uint8, 10)
	}

	snake.Print()
	snake.Move(constants.Right)
	snake.Print()
}

func TestShouldMoveHeadBottomLeft(t *testing.T) {
	snake := MakeSnake(3, 5, 3)

	table := make([][]uint8, 10)
	for i := range table {
		table[i] = make([]uint8, 10)
	}

	snake.Print()
	snake.Move(constants.Left)
	snake.Print()

}

func TestShouldMoveHeadTopLeft(t *testing.T) {
	var snakeLength int32 = 5
	snakeBody := make([]SnakePoint, snakeLength)
	snakeBody[0] = SnakePoint{X: int32(0), Y: int32(4)}
	snakeBody[1] = SnakePoint{X: int32(0), Y: int32(3)}
	snakeBody[2] = SnakePoint{X: int32(0), Y: int32(2)}
	snakeBody[3] = SnakePoint{X: int32(0), Y: int32(1)}
	snakeBody[4] = SnakePoint{X: int32(0), Y: int32(0)}

	snake := Snake{
		Body: &snakeBody,
		Id:   randSeq(10),
	}

	snake.Move(constants.Left)
	snake.Print()

	snake.Move(constants.Up)
	snake.Print()

	if (*snake.Body)[0].X != -1 && (*snake.Body)[0].Y != 5 ||
		(*snake.Body)[1].X != -1 && (*snake.Body)[1].Y != 4 ||
		(*snake.Body)[2].X != 0 && (*snake.Body)[2].Y == 4 {
		t.Errorf("Head is not in correct place")
	}
}

func TestShouldMoveToRight(t *testing.T) {
	var startingX int32 = 3
	var startingY int32 = 5
	var snakeLength int32 = 5

	snakeBody := make([]SnakePoint, snakeLength)
	snakeBody[0] = SnakePoint{X: int32(0), Y: int32(0)}
	snakeBody[1] = SnakePoint{X: int32(0), Y: int32(1)}
	snakeBody[2] = SnakePoint{X: int32(0), Y: int32(2)}
	snakeBody[3] = SnakePoint{X: int32(0), Y: int32(3)}
	snakeBody[4] = SnakePoint{X: int32(0), Y: int32(4)}

	snake := Snake{
		Body: &snakeBody,
		Id:   randSeq(10),
	}

	table := make([][]uint8, 10)
	for i := range table {
		table[i] = make([]uint8, 10)
	}

	snake.Print()
	snake.Move(constants.Right)
	snake.Move(constants.Right)

	var headDidNotMoveRight = (*(snake).Body)[0].X != int32(startingX+2)
	var headDidNotMoveUp = (*(snake).Body)[0].Y != int32(startingY)

	if headDidNotMoveRight || headDidNotMoveUp {
		t.Errorf("Head is not in correct place")
	}
}

func TestShouldMoveToRightIndependentOfTable(t *testing.T) {
	var startingX int32 = 3
	var startingY int32 = 5
	var snakeLength int32 = 5

	snake := MakeSnakePointingUp(startingX, startingY, snakeLength)

	table := make([][]uint8, 10)
	for i := range table {
		table[i] = make([]uint8, 10)
	}

	snake.Move(constants.Right)
	snake.Move(constants.Right)

	var headDidNotMoveRight = (*(snake).Body)[0].X != int32(startingX+2)
	var headDidNotMoveUp = (*(snake).Body)[0].Y != int32(startingY)

	if headDidNotMoveRight || headDidNotMoveUp {
		t.Errorf("Head is not in correct place")
	}
}

func TestRenderingTwoSnakes(t *testing.T) {
	// Snake 1
	var startingX int32 = 3
	var startingY int32 = 5
	var snakeLength int32 = 5

	// Snake 2
	var startingX2 int32 = 1
	var startingY2 int32 = 3
	var snakeLength2 int32 = 5

	snake := MakeSnakePointingUp(startingX, startingY, snakeLength)
	snake2 := MakeSnakePointingRight(startingX2, startingY2, snakeLength2)

	table := make([][]uint8, 10)
	for i := range table {
		table[i] = make([]uint8, 10)
	}

	// PrintTable(&table, snake)
	// PrintTable(&table, snake2)
	// RenderGameToTerminal(&table, []*Snake{snake, snake2})
	// Move snake 1 up
	// MoveSnake(snake, &table, constants.Up)
	// MoveSnake(snake, &table, constants.Up)

	if IsHeadTouchingOtherSnake(snake, snake2) {
		HandleSnakeContact(snake, snake2)
	}
	// MoveSnake(snake, &table, constants.Down)
	// MoveSnake(snake, &table, constants.Right)
	// RenderGameToTerminal(&table, []*Snake{snake, snake2})

	// var headDidNotMoveRight = (*(snake).Body)[0].X != startingX+2
	var headDidNotMoveUp = (*(snake).Body)[0].Y != int32(startingY-2)

	if headDidNotMoveUp {
		t.Errorf("Head is not in correct place")
	}

}
