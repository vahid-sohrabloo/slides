// BUILD START OMIT
//go:build 386 || amd64 || amd64p32 || arm || arm64 || mipsle || mips64le || mips64p32le || ppc64le || riscv || riscv64
// +build 386 amd64 amd64p32 arm arm64 mipsle mips64le mips64p32le ppc64le riscv riscv64

// BUILD END OMIT

package main

import (
	"io"
	"unsafe"
)

// WRITE START OMIT
type slice struct {
	Data uintptr
	Len  int
	Cap  int
}

func (c *Base[T]) WriteTo(w io.Writer) (int64, error) {
	s := *(*slice)(unsafe.Pointer(&c.values))
	s.Len *= c.size
	s.Cap *= c.size
	var n int64
	src := *(*[]byte)(unsafe.Pointer(&s))
	nw, err := w.Write(src)
	return int64(nw) + n, err
}

// WRITE END OMIT
