// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [InputServer] class.
var inputServerClass = _InputServerClass{objc.GetClass("NSInputServer")}

type _InputServerClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSInputServer

type InputServer struct {
	objectivec.Object
}

// InputServerFrom constructs a [InputServer] from an unsafe.Pointer.
func InputServerFrom(ptr unsafe.Pointer) InputServer {
	return InputServer{objectivec.Object{objc.ID(ptr)}}
}



