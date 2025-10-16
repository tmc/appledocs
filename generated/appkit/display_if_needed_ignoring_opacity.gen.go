
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [displayIfNeededIgnoringOpacity] class.
var displayIfNeededIgnoringOpacityClass _displayIfNeededIgnoringOpacityClass

func init() {
	displayIfNeededIgnoringOpacityClass = _displayIfNeededIgnoringOpacityClass{objc.GetClass("displayIfNeededIgnoringOpacity")}
}

type _displayIfNeededIgnoringOpacityClass struct {
	objc.Class
}

// An interface definition for the [displayIfNeededIgnoringOpacity] class.
type IdisplayIfNeededIgnoringOpacity interface {
	ID() objc.ID
}

type displayIfNeededIgnoringOpacity struct {
	id objc.ID
}

func displayIfNeededIgnoringOpacityFrom(ptr unsafe.Pointer) displayIfNeededIgnoringOpacity {
	return displayIfNeededIgnoringOpacity{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ displayIfNeededIgnoringOpacity) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _displayIfNeededIgnoringOpacityClass) Alloc() displayIfNeededIgnoringOpacity {
	rv := objc.Send[displayIfNeededIgnoringOpacity](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _displayIfNeededIgnoringOpacityClass) New() displayIfNeededIgnoringOpacity {
	rv := objc.Send[displayIfNeededIgnoringOpacity](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdisplayIfNeededIgnoringOpacity creates and returns a new initialized instance.
func NewdisplayIfNeededIgnoringOpacity() displayIfNeededIgnoringOpacity {
	return displayIfNeededIgnoringOpacityClass.New()
}

// Init initializes the instance.
func (d_ displayIfNeededIgnoringOpacity) Init() displayIfNeededIgnoringOpacity {
	rv := objc.Send[displayIfNeededIgnoringOpacity](d_.ID(), selInit)
	return rv
}
