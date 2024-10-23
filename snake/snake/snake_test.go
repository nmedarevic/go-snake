package snake

import (
	"testing"

	"snakegame.com/constants"
)

func Test_SnakeGoingLeft(t *testing.T) {
	t.Run("Should move the head down by pressing Down", func(t *testing.T) {
		snakeBody := make([]SnakePoint, 0)
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(0)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(1), Y: int32(0)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(2), Y: int32(0)})

		snake := Snake{
			Body: &snakeBody,
			Id:   RandSeq(10),
		}

		// snake.Print()
		snake.Move(constants.Down)
		// snake.Print()

		if (*snake.Body)[0].X != 0 || (*snake.Body)[0].Y != -1 ||
			(*snake.Body)[1].X != 0 || (*snake.Body)[1].Y != 0 ||
			(*snake.Body)[2].X != 1 || (*snake.Body)[2].Y != 0 {
			t.Errorf("Head is not in correct place")
		}
	})

	t.Run("Should move the head up by pressing Up", func(t *testing.T) {
		snakeBody := make([]SnakePoint, 0)
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(0)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(1), Y: int32(0)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(2), Y: int32(0)})

		snake := Snake{
			Body: &snakeBody,
			Id:   RandSeq(10),
		}

		// snake.Print()
		snake.Move(constants.Up)
		// snake.Print()

		if (*snake.Body)[0].X != 0 || (*snake.Body)[0].Y != 1 ||
			(*snake.Body)[1].X != 0 || (*snake.Body)[1].Y != 0 ||
			(*snake.Body)[2].X != 1 || (*snake.Body)[2].Y != 0 {
			t.Errorf("Head is not in correct place")
		}
	})

	t.Run("Should not move the head back by pressing Left", func(t *testing.T) {
		snakeBody := make([]SnakePoint, 0)
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(0)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(1), Y: int32(0)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(2), Y: int32(0)})

		snake := Snake{
			Body: &snakeBody,
			Id:   RandSeq(10),
		}

		// snake.Print()
		snake.Move(constants.Left)
		// snake.Print()

		if (*snake.Body)[0].X != -1 || (*snake.Body)[0].Y != 0 ||
			(*snake.Body)[1].X != 0 || (*snake.Body)[1].Y != 0 ||
			(*snake.Body)[2].X != 1 || (*snake.Body)[2].Y != 0 {
			t.Errorf("Head is not in correct place")
		}
	})

	t.Run("Should not move the head forward by pressing Right", func(t *testing.T) {
		snakeBody := make([]SnakePoint, 0)
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(0)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(1), Y: int32(0)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(2), Y: int32(0)})

		snake := Snake{
			Body: &snakeBody,
			Id:   RandSeq(10),
		}

		// snake.Print()
		snake.Move(constants.Right)
		// snake.Print()

		if (*snake.Body)[0].X != 0 || (*snake.Body)[0].Y != 0 ||
			(*snake.Body)[1].X != 1 || (*snake.Body)[1].Y != 0 ||
			(*snake.Body)[2].X != 2 || (*snake.Body)[2].Y != 0 {
			t.Errorf("Head is not in correct place")
		}
	})
}

func Test_SnakeGoingRight(t *testing.T) {
	t.Run("Should move the head down by pressing Down", func(t *testing.T) {
		snakeBody := make([]SnakePoint, 0)
		snakeBody = append(snakeBody, SnakePoint{X: int32(2), Y: int32(0)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(1), Y: int32(0)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(0)})

		snake := Snake{
			Body: &snakeBody,
			Id:   RandSeq(10),
		}

		// snake.Print()
		snake.Move(constants.Down)
		// snake.Print()

		if (*snake.Body)[0].X != 2 || (*snake.Body)[0].Y != -1 ||
			(*snake.Body)[1].X != 2 || (*snake.Body)[1].Y != 0 ||
			(*snake.Body)[2].X != 1 || (*snake.Body)[2].Y != 0 {
			t.Errorf("Head is not in correct place")
		}
	})

	t.Run("Should move the head up by pressing Up", func(t *testing.T) {
		snakeBody := make([]SnakePoint, 0)
		snakeBody = append(snakeBody, SnakePoint{X: int32(2), Y: int32(0)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(1), Y: int32(0)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(0)})

		snake := Snake{
			Body: &snakeBody,
			Id:   RandSeq(10),
		}

		// snake.Print()
		snake.Move(constants.Up)
		// snake.Print()

		if (*snake.Body)[0].X != 2 || (*snake.Body)[0].Y != 1 ||
			(*snake.Body)[1].X != 2 || (*snake.Body)[1].Y != 0 ||
			(*snake.Body)[2].X != 1 || (*snake.Body)[2].Y != 0 {
			t.Errorf("Head is not in correct place")
		}
	})

	t.Run("Should not move the head back by pressing Left", func(t *testing.T) {
		snakeBody := make([]SnakePoint, 0)
		snakeBody = append(snakeBody, SnakePoint{X: int32(2), Y: int32(0)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(1), Y: int32(0)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(0)})

		snake := Snake{
			Body: &snakeBody,
			Id:   RandSeq(10),
		}

		// snake.Print()
		snake.Move(constants.Left)
		// snake.Print()

		if (*snake.Body)[0].X != 2 || (*snake.Body)[0].Y != 0 ||
			(*snake.Body)[1].X != 1 || (*snake.Body)[1].Y != 0 ||
			(*snake.Body)[2].X != 0 || (*snake.Body)[2].Y != 0 {
			t.Errorf("Head is not in correct place")
		}
	})

	t.Run("Should not move the head forward by pressing Right", func(t *testing.T) {
		snakeBody := make([]SnakePoint, 0)
		snakeBody = append(snakeBody, SnakePoint{X: int32(2), Y: int32(0)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(1), Y: int32(0)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(0)})

		snake := Snake{
			Body: &snakeBody,
			Id:   RandSeq(10),
		}

		// snake.Print()
		snake.Move(constants.Right)
		// snake.Print()

		if (*snake.Body)[0].X != 3 || (*snake.Body)[0].Y != 0 ||
			(*snake.Body)[1].X != 2 || (*snake.Body)[1].Y != 0 ||
			(*snake.Body)[2].X != 1 || (*snake.Body)[2].Y != 0 {
			t.Errorf("Head is not in correct place")
		}
	})
}

func Test_SnakeGoingUp(t *testing.T) {
	t.Run("Should move the head down by pressing Down", func(t *testing.T) {
		snakeBody := make([]SnakePoint, 0)
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(2)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(1)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(0)})

		snake := Snake{
			Body: &snakeBody,
			Id:   RandSeq(10),
		}

		// snake.Print()
		snake.Move(constants.Down)
		// snake.Print()

		if (*snake.Body)[0].X != 0 || (*snake.Body)[0].Y != 2 ||
			(*snake.Body)[1].X != 0 || (*snake.Body)[1].Y != 1 ||
			(*snake.Body)[2].X != 0 || (*snake.Body)[2].Y != 0 {
			t.Errorf("Head is not in correct place")
		}
	})

	t.Run("Should move the head up by pressing Up", func(t *testing.T) {
		snakeBody := make([]SnakePoint, 0)
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(2)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(1)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(0)})

		snake := Snake{
			Body: &snakeBody,
			Id:   RandSeq(10),
		}

		// snake.Print()
		snake.Move(constants.Up)
		// snake.Print()

		if (*snake.Body)[0].X != 0 || (*snake.Body)[0].Y != 3 ||
			(*snake.Body)[1].X != 0 || (*snake.Body)[1].Y != 2 ||
			(*snake.Body)[2].X != 0 || (*snake.Body)[2].Y != 1 {
			t.Errorf("Head is not in correct place")
		}
	})

	t.Run("Should not move the head back by pressing Left", func(t *testing.T) {
		snakeBody := make([]SnakePoint, 0)
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(2)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(1)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(0)})

		snake := Snake{
			Body: &snakeBody,
			Id:   RandSeq(10),
		}

		// snake.Print()
		snake.Move(constants.Left)
		// snake.Print()

		if (*snake.Body)[0].X != -1 || (*snake.Body)[0].Y != 2 ||
			(*snake.Body)[1].X != 0 || (*snake.Body)[1].Y != 2 ||
			(*snake.Body)[2].X != 0 || (*snake.Body)[2].Y != 1 {
			t.Errorf("Head is not in correct place")
		}
	})

	t.Run("Should not move the head back by pressing Right", func(t *testing.T) {
		snakeBody := make([]SnakePoint, 0)
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(2)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(1)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(0)})

		snake := Snake{
			Body: &snakeBody,
			Id:   RandSeq(10),
		}

		// snake.Print()
		snake.Move(constants.Right)
		// snake.Print()

		if (*snake.Body)[0].X != 1 || (*snake.Body)[0].Y != 2 ||
			(*snake.Body)[1].X != 0 || (*snake.Body)[1].Y != 2 ||
			(*snake.Body)[2].X != 0 || (*snake.Body)[2].Y != 1 {
			t.Errorf("Head is not in correct place")
		}
	})
}

func Test_SnakeGoingDown(t *testing.T) {
	t.Run("Should move the head down by pressing Down", func(t *testing.T) {
		snakeBody := make([]SnakePoint, 0)
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(0)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(1)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(2)})

		snake := Snake{
			Body: &snakeBody,
			Id:   RandSeq(10),
		}

		// snake.Print()
		snake.Move(constants.Down)
		// snake.Print()

		if (*snake.Body)[0].X != 0 || (*snake.Body)[0].Y != -1 ||
			(*snake.Body)[1].X != 0 || (*snake.Body)[1].Y != 0 ||
			(*snake.Body)[2].X != 0 || (*snake.Body)[2].Y != 1 {
			t.Errorf("Head is not in correct place")
		}
	})

	t.Run("Should move the head up by pressing Up", func(t *testing.T) {
		snakeBody := make([]SnakePoint, 0)
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(0)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(1)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(2)})

		snake := Snake{
			Body: &snakeBody,
			Id:   RandSeq(10),
		}

		// snake.Print()
		snake.Move(constants.Up)
		// snake.Print()

		if (*snake.Body)[0].X != 0 || (*snake.Body)[0].Y != 0 ||
			(*snake.Body)[1].X != 0 || (*snake.Body)[1].Y != 1 ||
			(*snake.Body)[2].X != 0 || (*snake.Body)[2].Y != 2 {
			t.Errorf("Head is not in correct place")
		}
	})

	t.Run("Should not move the head back by pressing Left", func(t *testing.T) {
		snakeBody := make([]SnakePoint, 0)
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(0)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(1)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(2)})

		snake := Snake{
			Body: &snakeBody,
			Id:   RandSeq(10),
		}

		// snake.Print()
		snake.Move(constants.Left)
		// snake.Print()

		if (*snake.Body)[0].X != -1 || (*snake.Body)[0].Y != 0 ||
			(*snake.Body)[1].X != 0 || (*snake.Body)[1].Y != 0 ||
			(*snake.Body)[2].X != 0 || (*snake.Body)[2].Y != 1 {
			t.Errorf("Head is not in correct place")
		}
	})

	t.Run("Should not move the head back by pressing Right", func(t *testing.T) {
		snakeBody := make([]SnakePoint, 0)
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(0)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(1)})
		snakeBody = append(snakeBody, SnakePoint{X: int32(0), Y: int32(2)})

		snake := Snake{
			Body: &snakeBody,
			Id:   RandSeq(10),
		}

		// snake.Print()
		snake.Move(constants.Right)
		// snake.Print()

		if (*snake.Body)[0].X != 1 || (*snake.Body)[0].Y != 0 ||
			(*snake.Body)[1].X != 0 || (*snake.Body)[1].Y != 0 ||
			(*snake.Body)[2].X != 0 || (*snake.Body)[2].Y != 1 {
			t.Errorf("Head is not in correct place")
		}
	})
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

	if IsHeadTouchingOtherSnake(snake, snake2) {
		HandleSnakeContact(snake, snake2)
	}

	var headDidNotMoveUp = (*(snake).Body)[0].Y != int32(startingY-2)

	if headDidNotMoveUp {
		t.Errorf("Head is not in correct place")
	}

}
