// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Controller] class.
var controllerClass = _ControllerClass{objc.GetClass("NSController")}

type _ControllerClass struct {
	class objc.Class
}

// An abstract class that implements the and informal protocols required for controller classes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSController

type Controller struct {
	objectivec.Object
}

// ControllerFrom constructs a [Controller] from an unsafe.Pointer.
//
// An abstract class that implements the and informal protocols required for controller classes.
func ControllerFrom(ptr unsafe.Pointer) Controller {
	return Controller{objectivec.Object{objc.ID(ptr)}}
}



