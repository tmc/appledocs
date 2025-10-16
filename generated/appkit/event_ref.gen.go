
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [eventRef] class.
var eventRefClass _eventRefClass

func init() {
	eventRefClass = _eventRefClass{objc.GetClass("eventRef")}
}

type _eventRefClass struct {
	objc.Class
}

// An interface definition for the [eventRef] class.
type IeventRef interface {
	ID() objc.ID
}

type eventRef struct {
	id objc.ID
}

func eventRefFrom(ptr unsafe.Pointer) eventRef {
	return eventRef{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (e_ eventRef) ID() objc.ID {
	return e_.id
}

// Alloc allocates a new instance without initialization.
func (ec _eventRefClass) Alloc() eventRef {
	rv := objc.Send[eventRef](objc.ID(ec.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ec _eventRefClass) New() eventRef {
	rv := objc.Send[eventRef](objc.ID(ec.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NeweventRef creates and returns a new initialized instance.
func NeweventRef() eventRef {
	return eventRefClass.New()
}

// Init initializes the instance.
func (e_ eventRef) Init() eventRef {
	rv := objc.Send[eventRef](e_.ID(), selInit)
	return rv
}
