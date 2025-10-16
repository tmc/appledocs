
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [acceptsMouseMovedEvents] class.
var acceptsMouseMovedEventsClass _acceptsMouseMovedEventsClass

func init() {
	acceptsMouseMovedEventsClass = _acceptsMouseMovedEventsClass{objc.GetClass("acceptsMouseMovedEvents")}
}

type _acceptsMouseMovedEventsClass struct {
	objc.Class
}

// An interface definition for the [acceptsMouseMovedEvents] class.
type IacceptsMouseMovedEvents interface {
	ID() objc.ID
}

type acceptsMouseMovedEvents struct {
	id objc.ID
}

func acceptsMouseMovedEventsFrom(ptr unsafe.Pointer) acceptsMouseMovedEvents {
	return acceptsMouseMovedEvents{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ acceptsMouseMovedEvents) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _acceptsMouseMovedEventsClass) Alloc() acceptsMouseMovedEvents {
	rv := objc.Send[acceptsMouseMovedEvents](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _acceptsMouseMovedEventsClass) New() acceptsMouseMovedEvents {
	rv := objc.Send[acceptsMouseMovedEvents](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewacceptsMouseMovedEvents creates and returns a new initialized instance.
func NewacceptsMouseMovedEvents() acceptsMouseMovedEvents {
	return acceptsMouseMovedEventsClass.New()
}

// Init initializes the instance.
func (a_ acceptsMouseMovedEvents) Init() acceptsMouseMovedEvents {
	rv := objc.Send[acceptsMouseMovedEvents](a_.ID(), selInit)
	return rv
}
