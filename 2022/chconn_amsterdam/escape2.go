package main

func Escape() {
	// ESCAPE START OMIT
	var (
		a = 1 // moved to heap: a
		b = false
		c = make(chan struct{})
	)
	go func() {
		if b {
			a++
		}
		close(c)
	}()
	<-c
	println(a, b) // 1 false
	// ESCAPE END OMIT
}
