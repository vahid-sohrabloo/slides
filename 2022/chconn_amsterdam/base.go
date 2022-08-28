package main

import "unsafe"

// BASE START OMIT
type Base[T comparable] struct {
	size   int
	numRow int
	b      []byte
	values []T
}

// New create a new column
func New[T comparable]() *Base[T] {
	var tmpValue T
	size := int(unsafe.Sizeof(tmpValue))
	return &Base[T]{
		size: size,
	}
}

// BASE END OMIT

// READ START OMIT
func (c *Base[T]) Data() []T {
	value := *(*[]T)(unsafe.Pointer(&c.b))
	return value[:c.numRow]
}

func (c *Base[T]) Read(value []T) []T {
	return append(value, c.Data()...)
}

func (c *Base[T]) Row(row int) T {
	i := row * c.size
	return *(*T)(unsafe.Pointer(&c.b[i]))
}

// READ END OMIT

// WRITE START OMIT

// Append value for insert
func (c *Base[T]) Append(v ...T) {
	c.values = append(c.values, v...)
	c.numRow += len(v)
}

// WRITE END OMIT
