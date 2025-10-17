// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PathControl] class.
var pathControlClass = _PathControlClass{objc.GetClass("NSPathControl")}

type _PathControlClass struct {
	class objc.Class
}

// A display of a file system path or virtual path information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl

type PathControl struct {
	Control
}

// PathControlFrom constructs a [PathControl] from an unsafe.Pointer.
//
// A display of a file system path or virtual path information.
func PathControlFrom(ptr unsafe.Pointer) PathControl {
	return PathControl{
		Control: ControlFrom(ptr),
	}
}



