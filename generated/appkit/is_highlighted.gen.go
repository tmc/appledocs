
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isHighlighted] class.
var isHighlightedClass _isHighlightedClass

func init() {
	isHighlightedClass = _isHighlightedClass{objc.GetClass("isHighlighted")}
}

type _isHighlightedClass struct {
	objc.Class
}

// An interface definition for the [isHighlighted] class.
type IisHighlighted interface {
	ID() objc.ID
}

type isHighlighted struct {
	id objc.ID
}

func isHighlightedFrom(ptr unsafe.Pointer) isHighlighted {
	return isHighlighted{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isHighlighted) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isHighlightedClass) Alloc() isHighlighted {
	rv := objc.Send[isHighlighted](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isHighlightedClass) New() isHighlighted {
	rv := objc.Send[isHighlighted](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisHighlighted creates and returns a new initialized instance.
func NewisHighlighted() isHighlighted {
	return isHighlightedClass.New()
}

// Init initializes the instance.
func (i_ isHighlighted) Init() isHighlighted {
	rv := objc.Send[isHighlighted](i_.ID(), selInit)
	return rv
}
