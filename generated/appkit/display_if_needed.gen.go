
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [displayIfNeeded] class.
var displayIfNeededClass _displayIfNeededClass

func init() {
	displayIfNeededClass = _displayIfNeededClass{objc.GetClass("displayIfNeeded")}
}

type _displayIfNeededClass struct {
	objc.Class
}

// An interface definition for the [displayIfNeeded] class.
type IdisplayIfNeeded interface {
	ID() objc.ID
}

type displayIfNeeded struct {
	id objc.ID
}

func displayIfNeededFrom(ptr unsafe.Pointer) displayIfNeeded {
	return displayIfNeeded{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ displayIfNeeded) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _displayIfNeededClass) Alloc() displayIfNeeded {
	rv := objc.Send[displayIfNeeded](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _displayIfNeededClass) New() displayIfNeeded {
	rv := objc.Send[displayIfNeeded](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdisplayIfNeeded creates and returns a new initialized instance.
func NewdisplayIfNeeded() displayIfNeeded {
	return displayIfNeededClass.New()
}

// Init initializes the instance.
func (d_ displayIfNeeded) Init() displayIfNeeded {
	rv := objc.Send[displayIfNeeded](d_.ID(), selInit)
	return rv
}
