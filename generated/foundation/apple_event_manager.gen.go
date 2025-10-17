// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AppleEventManager] class.
var appleEventManagerClass = _AppleEventManagerClass{objc.GetClass("NSAppleEventManager")}

type _AppleEventManagerClass struct {
	class objc.Class
}

// A mechanism for registering handler routines for specific types of Apple events and dispatching events to those handlers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventManager

type AppleEventManager struct {
	objectivec.Object
}

// AppleEventManagerFrom constructs a [AppleEventManager] from an unsafe.Pointer.
//
// A mechanism for registering handler routines for specific types of Apple events and dispatching events to those handlers.
func AppleEventManagerFrom(ptr unsafe.Pointer) AppleEventManager {
	return AppleEventManager{objectivec.Object{objc.ID(ptr)}}
}



