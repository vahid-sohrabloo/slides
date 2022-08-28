package main

import (
	"fmt"
	"os"
	"reflect"
	"syscall"
	"unsafe"
)

// UNSAFE1 START OMIT
func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

func Float64frombits(b uint64) float64 {
	return *(*float64)(unsafe.Pointer(&b))
}

// UNSAFE1 END OMIT

// UNSAFE12 START OMIT

func ByteSlice2String(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs)) // invalid result
}

// UNSAFE12 END OMIT

func Unsafe2() {
	// UNSAFE2 START OMIT

	type T struct{ a int }
	var t T
	fmt.Printf("%p\n", &t)                          // 0xc6233120a8
	println(&t)                                     // 0xc6233120a8
	fmt.Printf("%x\n", uintptr(unsafe.Pointer(&t))) // c6233120a8
	// UNSAFE2 END OMIT
}

func Unsafe3() {
	// UNSAFE3 START OMIT

	type T struct {
		x bool
		y [3]int16
	}

	const N = unsafe.Offsetof(T{}.y)
	const M = unsafe.Sizeof(T{}.y[0])
	t := T{y: [3]int16{123, 456, 789}}
	p := unsafe.Pointer(&t)
	// "uintptr(p)+N+M+M" is the address of t.y[2].
	ty2 := (*int16)(unsafe.Pointer(uintptr(p) + N + M + M))
	fmt.Println(*ty2) // 789
	// UNSAFE3 END OMIT
	{
		// UNSAFE31 START OMIT
		// THIS IS DANGEROUS
		t := T{y: [3]int16{123, 456, 789}}
		p := unsafe.Pointer(&t)
		// ty2 := (*int16)(unsafe.Pointer(uintptr(p)+N+M+M))
		addr := uintptr(p) + N + M + M
		ty2 := (*int16)(unsafe.Pointer(addr))
		fmt.Println(*ty2)
		// UNSAFE31 END OMIT
	}
}

// UNSAFE4 START OMIT
// Assume this function will not inlined.
func DoSomething(addr uintptr) {
	// read or write values at the passed address ...
}

func Syscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err error)

// UNSAFE4 END OMIT
func CallUnsafe(path string) {
	// UNSAFE41 START OMIT

	const FS_IOC_GETFLAGS = 0x80086601
	file, _ := os.Open(path)
	var attr int
	syscall.Syscall(syscall.SYS_IOCTL, file.Fd(), FS_IOC_GETFLAGS, uintptr(unsafe.Pointer(&attr)))
	// UNSAFE41 END OMIT

	{
		// UNSAFE42 START OMIT

		u := file.Fd()
		syscall.Syscall(syscall.SYS_IOCTL, u, FS_IOC_GETFLAGS, uintptr(unsafe.Pointer(&attr)))

		// UNSAFE42 END OMIT
	}
}

// Assume this function will not inlined.
func Unsafe5(addr uintptr) {
	// UNSAFE5 START OMIT
	p := (*int)(unsafe.Pointer(reflect.ValueOf(new(int)).Pointer()))
	// UNSAFE5 END OMIT
	fmt.Println(p)
}

func Unsafe6() {
	// UNSAFE6 START OMIT
	a := [...]byte{'G', 'o', 'l', 'a', 'n', 'g'}
	s := "Java"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	hdr.Data = uintptr(unsafe.Pointer(&a))
	hdr.Len = len(a)
	fmt.Println(s) // Golang
	// Now s and a share the same byte sequence, which
	// makes the bytes in the string s become mutable.
	a[2], a[3], a[4], a[5] = 'o', 'g', 'l', 'e'
	fmt.Println(s) // Google
	// UNSAFE6 END OMIT
}

// UNSAFE61 START OMIT
// This code is from easyjson
func BytesToStr(data []byte) string {
	h := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	shdr := reflect.StringHeader{Data: h.Data, Len: h.Len}
	return *(*string)(unsafe.Pointer(&shdr))
}

func BytesToStr2(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

// UNSAFE61 END OMIT

func Unsafe6Bug() {
	// UNSAFE62 START OMIT
	var d = []byte{'G', 'o', 'l', 'a', 'n', 'g'}
	s := BytesToStr(d)
	fmt.Println(s)
	// UNSAFE62 END OMIT
	{
		// UNSAFE63 START OMIT
		var d = []byte{'G', 'o', 'l', 'a', 'n', 'g'}
		h := (*reflect.SliceHeader)(unsafe.Pointer(&d))
		shdr := reflect.StringHeader{Data: h.Data, Len: h.Len}
		s := *(*string)(unsafe.Pointer(&shdr))
		fmt.Println(s)
		// UNSAFE63 END OMIT

	}
}
