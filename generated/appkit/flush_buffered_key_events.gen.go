
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [flushBufferedKeyEvents] class.
var flushBufferedKeyEventsClass _flushBufferedKeyEventsClass

func init() {
	flushBufferedKeyEventsClass = _flushBufferedKeyEventsClass{objc.GetClass("flushBufferedKeyEvents")}
}

type _flushBufferedKeyEventsClass struct {
	objc.Class
}

// An interface definition for the [flushBufferedKeyEvents] class.
type IflushBufferedKeyEvents interface {
	ID() objc.ID
}

type flushBufferedKeyEvents struct {
	id objc.ID
}

func flushBufferedKeyEventsFrom(ptr unsafe.Pointer) flushBufferedKeyEvents {
	return flushBufferedKeyEvents{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ flushBufferedKeyEvents) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _flushBufferedKeyEventsClass) Alloc() flushBufferedKeyEvents {
	rv := objc.Send[flushBufferedKeyEvents](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _flushBufferedKeyEventsClass) New() flushBufferedKeyEvents {
	rv := objc.Send[flushBufferedKeyEvents](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewflushBufferedKeyEvents creates and returns a new initialized instance.
func NewflushBufferedKeyEvents() flushBufferedKeyEvents {
	return flushBufferedKeyEventsClass.New()
}

// Init initializes the instance.
func (f_ flushBufferedKeyEvents) Init() flushBufferedKeyEvents {
	rv := objc.Send[flushBufferedKeyEvents](f_.ID(), selInit)
	return rv
}
