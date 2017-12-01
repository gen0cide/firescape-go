// math.go

package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import (
	"github.com/davecgh/go-spew/spew"
	uuid "github.com/satori/go.uuid"
)

//export Puts
func Puts(buf *C.char) *C.char {
	spew.Dump(C.GoString(buf))
	ret := C.CString(uuid.NewV4().String())
	return ret
}

// main function is required, don't know why!
func main() {} // a dummy function
