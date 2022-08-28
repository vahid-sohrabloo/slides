package main

import "io"

type ColumnBasic interface {
	ReadRaw(num int, r io.Reader) error
	HeaderReader(io.Reader, bool) error
	HeaderWriter(io.Writer)
	WriteTo(io.Writer) (int64, error)
	NumRow() int
	Reset()
	SetType(v []byte)
	Type() []byte
	SetName(v []byte)
	Name() []byte
	Validate() error
	columnType() string
	SetWriteBufferSize(int)
}

// COLUMN START OMIT
type Column[T any] interface {
	ColumnBasic
	Data() []T
	Read([]T) []T
	Row(int) T
	Append(...T)
}

type NullableColumn[T any] interface {
	Column[T]
	DataP() []*T
	ReadP([]*T) []*T
	RowP(int) *T
	AppendP(...*T)
}

//COLUMN END OMIT
