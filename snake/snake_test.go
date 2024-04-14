package snake

import (
	"testing"
)

func TestMove(t *testing.T) {
	snake := MakeSnake()

	table := make([][]uint8, 10)
	for i := range table {
		table[i] = make([]uint8, 10)
	}

	Move(snake, &table, Down)
	// Move(snake, &table, Down)
	// Move(snake, &table, Down)
	// Move(snake, &table, Down)
	Move(snake, &table, Up)
	PrintTable(&table, snake)
}
