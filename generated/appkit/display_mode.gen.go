
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [displayMode] class.
var displayModeClass _displayModeClass

func init() {
	displayModeClass = _displayModeClass{objc.GetClass("displayMode")}
}

type _displayModeClass struct {
	objc.Class
}

// An interface definition for the [displayMode] class.
type IdisplayMode interface {
	ID() objc.ID
}

type displayMode struct {
	id objc.ID
}

func displayModeFrom(ptr unsafe.Pointer) displayMode {
	return displayMode{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ displayMode) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _displayModeClass) Alloc() displayMode {
	rv := objc.Send[displayMode](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _displayModeClass) New() displayMode {
	rv := objc.Send[displayMode](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdisplayMode creates and returns a new initialized instance.
func NewdisplayMode() displayMode {
	return displayModeClass.New()
}

// Init initializes the instance.
func (d_ displayMode) Init() displayMode {
	rv := objc.Send[displayMode](d_.ID(), selInit)
	return rv
}
