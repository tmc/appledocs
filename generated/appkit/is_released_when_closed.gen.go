
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isReleasedWhenClosed] class.
var isReleasedWhenClosedClass _isReleasedWhenClosedClass

func init() {
	isReleasedWhenClosedClass = _isReleasedWhenClosedClass{objc.GetClass("isReleasedWhenClosed")}
}

type _isReleasedWhenClosedClass struct {
	objc.Class
}

// An interface definition for the [isReleasedWhenClosed] class.
type IisReleasedWhenClosed interface {
	ID() objc.ID
}

type isReleasedWhenClosed struct {
	id objc.ID
}

func isReleasedWhenClosedFrom(ptr unsafe.Pointer) isReleasedWhenClosed {
	return isReleasedWhenClosed{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isReleasedWhenClosed) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isReleasedWhenClosedClass) Alloc() isReleasedWhenClosed {
	rv := objc.Send[isReleasedWhenClosed](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isReleasedWhenClosedClass) New() isReleasedWhenClosed {
	rv := objc.Send[isReleasedWhenClosed](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisReleasedWhenClosed creates and returns a new initialized instance.
func NewisReleasedWhenClosed() isReleasedWhenClosed {
	return isReleasedWhenClosedClass.New()
}

// Init initializes the instance.
func (i_ isReleasedWhenClosed) Init() isReleasedWhenClosed {
	rv := objc.Send[isReleasedWhenClosed](i_.ID(), selInit)
	return rv
}
