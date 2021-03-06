package main

import (
	"fmt"
	"unsafe"
)

func main() {
	array := [...]int{1, 2, 3, 4, 5}
	pointer := &array[0]
	fmt.Println(*pointer)
	memoryAddress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	for i := 0; i < len(array)-1; i++ {
		pointer = (*int)(unsafe.Pointer(memoryAddress))
		fmt.Println(*pointer)
		memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	}
	pointer = (*int)(unsafe.Pointer(memoryAddress))
	fmt.Println("One more:", *pointer)
}
