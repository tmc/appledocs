// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [InputManager] class.
var (
	InputManagerClass     _InputManagerClass
	InputManagerClassOnce sync.Once
)

func getInputManagerClass() _InputManagerClass {
	InputManagerClassOnce.Do(func() {
		InputManagerClass = _InputManagerClass{objc.GetClass("NSInputManager")}
	})
	return InputManagerClass
}

type _InputManagerClass struct {
	class objc.Class
}

// An interface definition for the [InputManager] class.
type IInputManager interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSInputManager
type InputManager struct {
	objectivec.Object
}

// InputManagerFrom constructs a [InputManager] from an unsafe.Pointer.
func InputManagerFrom(ptr unsafe.Pointer) InputManager {
	return InputManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _InputManagerClass) Alloc() InputManager {
	rv := objc.Send[InputManager](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _InputManagerClass) New() InputManager {
	rv := objc.Send[InputManager](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InputManager) Init() InputManager {
	rv := objc.Send[InputManager](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InputManager) Autorelease() InputManager {
	rv := objc.Send[InputManager](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInputManager creates a new InputManager instance.
func NewInputManager() InputManager {
	return getInputManagerClass().New()
}




