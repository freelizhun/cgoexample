#include <dlfcn.h>
  
typedef int (*someFunc) (const char *s,const char *b);
 
int bridge_someFunc(someFunc f, const char *s, const char *b) {
    return f(s,b);
}
*/
import "C"
import (
    "fmt"
    "os"
    "unsafe"
)
 
var handle = -1
 
func myOpenImage(s string, b string) int {
    fmt.Fprintf(os.Stderr, "tagged image %s:%s error...\n", s,b)
    handle++
    return handle
}
 
func main() {
    libpath := C.CString("/usr/lib/httc_register_container.so")
    defer C.free(unsafe.Pointer(libpath))
    imglib := C.dlopen(libpath, C.RTLD_LAZY)
    var imghandle int
    if imglib != nil {
        httc_register_container := C.CString("httc_register_container")
        defer C.free(unsafe.Pointer(httc_register_container))
        fp := C.dlsym(imglib, httc_register_container)
        if fp != nil {
            imageName := C.CString("busybox")
            defer C.free(unsafe.Pointer(imageName))
            imageVersion := C.CString("1.28.4")
            defer C.free(unsafe.Pointer(imageVersion))
            imghandle = int(C.bridge_someFunc(C.someFunc(fp), imageName, imageVersion))
 
        } else {
            imghandle = myOpenImage("busybox","1.28.4")
        }
        C.dlclose(imglib)
    } else {
        imghandle = myOpenImage("busybox","1.28.4")
    }
    fmt.Printf("opened with handle %d\n", imghandle)
}
