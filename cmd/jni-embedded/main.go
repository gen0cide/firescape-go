// main.go

package main

// #cgo LDFLAGS: -ljvm -L/Library/Java/JavaVirtualMachines/jdk1.8.0_151.jdk/Contents/Home/jre/lib/ -L/Library/Java/JavaVirtualMachines/jdk1.8.0_151.jdk/Contents/Home/jre/lib/server/ -L/Library/Java/JavaVirtualMachines/jdk1.8.0_151.jdk/Contents/Home/bundle/Libraries
// #cgo CFLAGS: -I. -I /Library/Java/JavaVirtualMachines/jdk1.8.0_151.jdk/Contents/Home/jre/lib/server/ -I /Library/Java/JavaVirtualMachines/jdk1.8.0_151.jdk/Contents/Home/include/ -I /Library/Java/JavaVirtualMachines/jdk1.8.0_151.jdk/Contents/Home/include/darwin/
// #include "lib.h"
import "C"
import (
	"fmt"
	"os"
)

func main() {
	javaOpts := fmt.Sprintf("-Djava.class.path=%s", os.Args[1])
	fmt.Println(C.GoString(C.FirescapeInit(C.CString(javaOpts))))
}
