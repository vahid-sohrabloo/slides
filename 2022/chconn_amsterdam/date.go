package main

import "time"

// DATE START OMIT
type DateType[T any] interface {
	comparable
	FromTime(val time.Time, precision int) T
	ToTime(val *time.Location, precision int) time.Time
}

type Date[T DateType[T]] struct {
	Base[T]
	loc       *time.Location
	precision int
}

// DATE END OMIT
