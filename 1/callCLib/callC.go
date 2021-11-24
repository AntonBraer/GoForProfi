package main

// #cgo CFLAGS: -I${SRCDIR}/callClib
// #cgo LDFLAGS: ${SRCDIR}/callC.a
// #include <stdlib.h>
// #include <callC.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Going to call a C Function")
	C.cHello()
	fmt.Println("Another C function")
	message := C.CString("This is Anton!")
	defer C.free(unsafe.Pointer(message))
	C.printMessage(message)
	fmt.Println("Done!")
}
