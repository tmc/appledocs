
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [display] class.
var displayClass _displayClass

func init() {
	displayClass = _displayClass{objc.GetClass("display")}
}

type _displayClass struct {
	objc.Class
}

// An interface definition for the [display] class.
type Idisplay interface {
	ID() objc.ID
}

type display struct {
	id objc.ID
}

func displayFrom(ptr unsafe.Pointer) display {
	return display{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ display) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _displayClass) Alloc() display {
	rv := objc.Send[display](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _displayClass) New() display {
	rv := objc.Send[display](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newdisplay creates and returns a new initialized instance.
func Newdisplay() display {
	return displayClass.New()
}

// Init initializes the instance.
func (d_ display) Init() display {
	rv := objc.Send[display](d_.ID(), selInit)
	return rv
}
