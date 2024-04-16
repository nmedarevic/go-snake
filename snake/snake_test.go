package snake

import (
	"testing"
)

func TestShouldNotMoveHeadBackwards(t *testing.T) {
	originalSnake := MakeSnake(3, 5, 3)
	snake := MakeSnake(3, 5, 3)

	table := make([][]uint8, 10)
	for i := range table {
		table[i] = make([]uint8, 10)
	}

	MoveSnake(snake, &table, Up)
	PrintTable(&table, snake)

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

	PrintTable(&table, snake)
	MoveSnake(snake, &table, Right)
	PrintTable(&table, snake)
}

func TestShouldMoveHeadBottomLeft(t *testing.T) {
	snake := MakeSnake(3, 5, 3)

	table := make([][]uint8, 10)
	for i := range table {
		table[i] = make([]uint8, 10)
	}

	PrintTable(&table, snake)
	MoveSnake(snake, &table, Left)
	PrintTable(&table, snake)
}

func TestShouldMoveHeadTopLeft(t *testing.T) {
	snake := MakeSnakeHeadPointUp(3, 5, 3)

	table := make([][]uint8, 10)
	for i := range table {
		table[i] = make([]uint8, 10)
	}

	PrintTable(&table, snake)
	MoveSnake(snake, &table, Left)
	PrintTable(&table, snake)

	MoveSnake(snake, &table, Left)
	PrintTable(&table, snake)

	MoveSnake(snake, &table, Up)
	PrintTable(&table, snake)
	MoveSnake(snake, &table, Up)
	PrintTable(&table, snake)
}

func TestShouldMoveHeadTopRight(t *testing.T) {
	snake := MakeSnakeHeadPointUp(3, 5, 5)

	table := make([][]uint8, 10)
	for i := range table {
		table[i] = make([]uint8, 10)
	}

	PrintTable(&table, snake)
	MoveSnake(snake, &table, Right)
	MoveSnake(snake, &table, Right)
	PrintTable(&table, snake)
}
