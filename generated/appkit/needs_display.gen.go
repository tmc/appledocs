
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [needsDisplay] class.
var needsDisplayClass _needsDisplayClass

func init() {
	needsDisplayClass = _needsDisplayClass{objc.GetClass("needsDisplay")}
}

type _needsDisplayClass struct {
	objc.Class
}

// An interface definition for the [needsDisplay] class.
type IneedsDisplay interface {
	ID() objc.ID
}

type needsDisplay struct {
	id objc.ID
}

func needsDisplayFrom(ptr unsafe.Pointer) needsDisplay {
	return needsDisplay{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (n_ needsDisplay) ID() objc.ID {
	return n_.id
}

// Alloc allocates a new instance without initialization.
func (nc _needsDisplayClass) Alloc() needsDisplay {
	rv := objc.Send[needsDisplay](objc.ID(nc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (nc _needsDisplayClass) New() needsDisplay {
	rv := objc.Send[needsDisplay](objc.ID(nc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewneedsDisplay creates and returns a new initialized instance.
func NewneedsDisplay() needsDisplay {
	return needsDisplayClass.New()
}

// Init initializes the instance.
func (n_ needsDisplay) Init() needsDisplay {
	rv := objc.Send[needsDisplay](n_.ID(), selInit)
	return rv
}
