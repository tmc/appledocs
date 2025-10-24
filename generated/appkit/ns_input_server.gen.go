// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [InputServer] class.
var (
	InputServerClass     _InputServerClass
	InputServerClassOnce sync.Once
)

func getInputServerClass() _InputServerClass {
	InputServerClassOnce.Do(func() {
		InputServerClass = _InputServerClass{objc.GetClass("NSInputServer")}
	})
	return InputServerClass
}

type _InputServerClass struct {
	class objc.Class
}

// An interface definition for the [InputServer] class.
type IInputServer interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSInputServer
type InputServer struct {
	objectivec.Object
}

// InputServerFrom constructs a [InputServer] from an unsafe.Pointer.
func InputServerFrom(ptr unsafe.Pointer) InputServer {
	return InputServer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _InputServerClass) Alloc() InputServer {
	rv := objc.Send[InputServer](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _InputServerClass) New() InputServer {
	rv := objc.Send[InputServer](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InputServer) Init() InputServer {
	rv := objc.Send[InputServer](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InputServer) Autorelease() InputServer {
	rv := objc.Send[InputServer](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInputServer creates a new InputServer instance.
func NewInputServer() InputServer {
	return getInputServerClass().New()
}




