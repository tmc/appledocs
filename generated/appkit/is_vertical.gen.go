
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isVertical] class.
var isVerticalClass _isVerticalClass

func init() {
	isVerticalClass = _isVerticalClass{objc.GetClass("isVertical")}
}

type _isVerticalClass struct {
	objc.Class
}

// An interface definition for the [isVertical] class.
type IisVertical interface {
	ID() objc.ID
}

type isVertical struct {
	id objc.ID
}

func isVerticalFrom(ptr unsafe.Pointer) isVertical {
	return isVertical{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isVertical) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isVerticalClass) Alloc() isVertical {
	rv := objc.Send[isVertical](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isVerticalClass) New() isVertical {
	rv := objc.Send[isVertical](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisVertical creates and returns a new initialized instance.
func NewisVertical() isVertical {
	return isVerticalClass.New()
}

// Init initializes the instance.
func (i_ isVertical) Init() isVertical {
	rv := objc.Send[isVertical](i_.ID(), selInit)
	return rv
}
