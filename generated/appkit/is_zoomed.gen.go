
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isZoomed] class.
var isZoomedClass _isZoomedClass

func init() {
	isZoomedClass = _isZoomedClass{objc.GetClass("isZoomed")}
}

type _isZoomedClass struct {
	objc.Class
}

// An interface definition for the [isZoomed] class.
type IisZoomed interface {
	ID() objc.ID
}

type isZoomed struct {
	id objc.ID
}

func isZoomedFrom(ptr unsafe.Pointer) isZoomed {
	return isZoomed{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isZoomed) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isZoomedClass) Alloc() isZoomed {
	rv := objc.Send[isZoomed](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isZoomedClass) New() isZoomed {
	rv := objc.Send[isZoomed](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisZoomed creates and returns a new initialized instance.
func NewisZoomed() isZoomed {
	return isZoomedClass.New()
}

// Init initializes the instance.
func (i_ isZoomed) Init() isZoomed {
	rv := objc.Send[isZoomed](i_.ID(), selInit)
	return rv
}
