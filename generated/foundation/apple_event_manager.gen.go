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

// An interface definition for the [AppleEventManager] class.
type IAppleEventManager interface {
	objectivec.IObject
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
// Alloc allocates a new instance without initialization.
func (ac _AppleEventManagerClass) Alloc() AppleEventManager {
	rv := objc.Send[AppleEventManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AppleEventManagerClass) New() AppleEventManager {
	rv := objc.Send[AppleEventManager](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AppleEventManager) Init() AppleEventManager {
	rv := objc.Send[AppleEventManager](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AppleEventManager) Autorelease() AppleEventManager {
	rv := objc.Send[AppleEventManager](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAppleEventManager creates a new AppleEventManager instance.
func NewAppleEventManager() AppleEventManager {
	return appleEventManagerClass.New()
}




