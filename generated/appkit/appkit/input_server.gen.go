// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [InputServer] class.
var InputServerClass objc.Class

func init() {
	InputServerClass = objc.GetClass("NSInputServer")
}

type InputServer struct {
	objc.ID
}

func InputServerFrom(ptr unsafe.Pointer) InputServer {
	return InputServer{
		ID: objc.ID(ptr),
	}
}



