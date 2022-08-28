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

// Data get all the data in current block as a slice.
//
// NOTE: the return slice only valid in current block, if you want to use it after, you should copy it. or use Read
// READ START OMIT
func (c *Base[T]) Data() []T {
	value := *(*[]T)(unsafe.Pointer(&c.b))
	return value[:c.numRow]
}

// Read reads all the data in current block and append to the input.
func (c *Base[T]) Read(value []T) []T {
	v := *(*[]T)(unsafe.Pointer(&c.b))
	return append(value, v[:c.numRow]...)
}

// Row return the value of given row.
// NOTE: Row number start from zero
func (c *Base[T]) Row(row int) T {
	i := row * c.size
	return *(*T)(unsafe.Pointer(&c.b[i]))
}

// READ END OMIT

// WRITE START OMIT

// Append value for insert
func (c *Base[T]) Append(v T) {
	c.numRow++
	c.values = append(c.values, v)
}

// AppendSlice append slice of value for insert
func (c *Base[T]) AppendSlice(v []T) {
	if len(v) == 0 {
		return
	}
	c.values = append(c.values, v...)
	c.numRow += len(v)
}

// WRITE END OMIT
