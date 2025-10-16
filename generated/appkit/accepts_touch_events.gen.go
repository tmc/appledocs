
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [acceptsTouchEvents] class.
var acceptsTouchEventsClass _acceptsTouchEventsClass

func init() {
	acceptsTouchEventsClass = _acceptsTouchEventsClass{objc.GetClass("acceptsTouchEvents")}
}

type _acceptsTouchEventsClass struct {
	objc.Class
}

// An interface definition for the [acceptsTouchEvents] class.
type IacceptsTouchEvents interface {
	ID() objc.ID
}

type acceptsTouchEvents struct {
	id objc.ID
}

func acceptsTouchEventsFrom(ptr unsafe.Pointer) acceptsTouchEvents {
	return acceptsTouchEvents{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ acceptsTouchEvents) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _acceptsTouchEventsClass) Alloc() acceptsTouchEvents {
	rv := objc.Send[acceptsTouchEvents](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _acceptsTouchEventsClass) New() acceptsTouchEvents {
	rv := objc.Send[acceptsTouchEvents](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewacceptsTouchEvents creates and returns a new initialized instance.
func NewacceptsTouchEvents() acceptsTouchEvents {
	return acceptsTouchEventsClass.New()
}

// Init initializes the instance.
func (a_ acceptsTouchEvents) Init() acceptsTouchEvents {
	rv := objc.Send[acceptsTouchEvents](a_.ID(), selInit)
	return rv
}
