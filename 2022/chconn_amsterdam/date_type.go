package main

import "time"

// DATE START OMIT
type DateTime uint32

func TimeToDateTime(t time.Time) DateTime {
	if t.Unix() <= 0 {
		return 0
	}
	return DateTime(t.Unix())
}

func (d DateTime) FromTime(v time.Time, precision int) DateTime {
	return TimeToDateTime(v)
}

func (d DateTime) ToTime(loc *time.Location, precision int) time.Time {
	return time.Unix(int64(d), 0).In(loc)
}

// DATE END OMIT
