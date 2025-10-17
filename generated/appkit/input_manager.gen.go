// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [InputManager] class.
var inputManagerClass = _InputManagerClass{objc.GetClass("NSInputManager")}

type _InputManagerClass struct {
	class objc.Class
}

// An interface definition for the [InputManager] class.
type IInputManager interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSInputManager

type InputManager struct {
	objectivec.Object
}

// InputManagerFrom constructs a [InputManager] from an unsafe.Pointer.
func InputManagerFrom(ptr unsafe.Pointer) InputManager {
	return InputManager{objectivec.Object{objc.ID(ptr)}}
}



