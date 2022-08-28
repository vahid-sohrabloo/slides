package main

type ArrayBase struct {
	offsetColumn *Base[uint64]
	dataColumn   ColumnBasic
}

// ARRAY START OMIT

type Array[T any] struct {
	ArrayBase
	columnData []T
}

func NewArray[T any](dataColumn Column[T]) *Array[T] {
	return &Array[T]{
		ArrayBase: ArrayBase{
			dataColumn:   dataColumn,
			offsetColumn: New[uint64](),
		},
	}
}

func (c *Array[T]) Data() [][]T {
	values := make([][]T, c.offsetColumn.numRow)
	// ...
	return values
}

func (c *Array[T]) Read(value [][]T) [][]T {
	// ...
	return value
}

// ARRAY END OMIT
