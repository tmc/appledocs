// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Responder] class.
var (
	ResponderClass     _ResponderClass
	ResponderClassOnce sync.Once
)

func getResponderClass() _ResponderClass {
	ResponderClassOnce.Do(func() {
		ResponderClass = _ResponderClass{objc.GetClass("NSResponder")}
	})
	return ResponderClass
}

type _ResponderClass struct {
	class objc.Class
}

// An interface definition for the [Responder] class.
type IResponder interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other AppKit classes.


// A parent class referenced by other AppKit classes. [Full Topic]
type Responder struct {
	objectivec.Object
}

// ResponderFrom constructs a [Responder] from an unsafe.Pointer.
//
// A parent class referenced by other AppKit classes.
func ResponderFrom(ptr unsafe.Pointer) Responder {
	return Responder{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _ResponderClass) Alloc() Responder {
	rv := objc.Send[Responder](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _ResponderClass) New() Responder {
	rv := objc.Send[Responder](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ Responder) Init() Responder {
	rv := objc.Send[Responder](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ Responder) Autorelease() Responder {
	rv := objc.Send[Responder](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewResponder creates a new Responder instance.
func NewResponder() Responder {
	return getResponderClass().New()
}




