// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [InputManager] class.
var InputManagerClass objc.Class

func init() {
	InputManagerClass = objc.GetClass("NSInputManager")
}

type InputManager struct {
	objc.ID
}

func InputManagerFrom(ptr unsafe.Pointer) InputManager {
	return InputManager{
		ID: objc.ID(ptr),
	}
}




