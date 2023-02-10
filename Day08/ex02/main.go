package main

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import <Foundation/Foundation.h>
//#include "application.h"
//#include "window.h"
import "C"

func main() {
	C.InitApplication()
	p := C.Window_Create(0, 0, 300, 200, C.CString("School 21"))
	C.Window_MakeKeyAndOrderFront(p)
	C.RunApplication()
}
