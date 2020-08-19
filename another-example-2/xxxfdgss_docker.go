package image
// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "greeter.h"
import "C"
import (
	"fmt"
	"unsafe"
)
func XgsCallDockerImageTag(imageName string,imageVersion string,imageOrigPath string){
	name := C.CString(imageName)
	defer C.free(unsafe.Pointer(name))
	year := C.int(2018)
	ptr := C.malloc(C.sizeof_char * 1024)
	defer C.free(unsafe.Pointer(ptr))
	size := C.greet(name, year, (*C.char)(ptr))
	b := C.GoBytes(ptr, size)
	fmt.Println(string(b))
	//fmt.Printf("Hello Go world: the imageNmae and tag is %s, %s, %s.\n", f.Name(), f.LongName(),imageOrigPath)
}
