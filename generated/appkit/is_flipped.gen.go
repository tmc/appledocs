
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isFlipped] class.
var isFlippedClass _isFlippedClass

func init() {
	isFlippedClass = _isFlippedClass{objc.GetClass("isFlipped")}
}

type _isFlippedClass struct {
	objc.Class
}

// An interface definition for the [isFlipped] class.
type IisFlipped interface {
	ID() objc.ID
}

type isFlipped struct {
	id objc.ID
}

func isFlippedFrom(ptr unsafe.Pointer) isFlipped {
	return isFlipped{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isFlipped) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isFlippedClass) Alloc() isFlipped {
	rv := objc.Send[isFlipped](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isFlippedClass) New() isFlipped {
	rv := objc.Send[isFlipped](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisFlipped creates and returns a new initialized instance.
func NewisFlipped() isFlipped {
	return isFlippedClass.New()
}

// Init initializes the instance.
func (i_ isFlipped) Init() isFlipped {
	rv := objc.Send[isFlipped](i_.ID(), selInit)
	return rv
}
