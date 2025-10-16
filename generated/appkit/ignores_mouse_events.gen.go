
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ignoresMouseEvents] class.
var ignoresMouseEventsClass _ignoresMouseEventsClass

func init() {
	ignoresMouseEventsClass = _ignoresMouseEventsClass{objc.GetClass("ignoresMouseEvents")}
}

type _ignoresMouseEventsClass struct {
	objc.Class
}

// An interface definition for the [ignoresMouseEvents] class.
type IignoresMouseEvents interface {
	ID() objc.ID
}

type ignoresMouseEvents struct {
	id objc.ID
}

func ignoresMouseEventsFrom(ptr unsafe.Pointer) ignoresMouseEvents {
	return ignoresMouseEvents{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ ignoresMouseEvents) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _ignoresMouseEventsClass) Alloc() ignoresMouseEvents {
	rv := objc.Send[ignoresMouseEvents](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _ignoresMouseEventsClass) New() ignoresMouseEvents {
	rv := objc.Send[ignoresMouseEvents](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewignoresMouseEvents creates and returns a new initialized instance.
func NewignoresMouseEvents() ignoresMouseEvents {
	return ignoresMouseEventsClass.New()
}

// Init initializes the instance.
func (i_ ignoresMouseEvents) Init() ignoresMouseEvents {
	rv := objc.Send[ignoresMouseEvents](i_.ID(), selInit)
	return rv
}
