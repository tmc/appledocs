
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isDrawingFindIndicator] class.
var isDrawingFindIndicatorClass _isDrawingFindIndicatorClass

func init() {
	isDrawingFindIndicatorClass = _isDrawingFindIndicatorClass{objc.GetClass("isDrawingFindIndicator")}
}

type _isDrawingFindIndicatorClass struct {
	objc.Class
}

// An interface definition for the [isDrawingFindIndicator] class.
type IisDrawingFindIndicator interface {
	ID() objc.ID
}

type isDrawingFindIndicator struct {
	id objc.ID
}

func isDrawingFindIndicatorFrom(ptr unsafe.Pointer) isDrawingFindIndicator {
	return isDrawingFindIndicator{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isDrawingFindIndicator) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isDrawingFindIndicatorClass) Alloc() isDrawingFindIndicator {
	rv := objc.Send[isDrawingFindIndicator](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isDrawingFindIndicatorClass) New() isDrawingFindIndicator {
	rv := objc.Send[isDrawingFindIndicator](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisDrawingFindIndicator creates and returns a new initialized instance.
func NewisDrawingFindIndicator() isDrawingFindIndicator {
	return isDrawingFindIndicatorClass.New()
}

// Init initializes the instance.
func (i_ isDrawingFindIndicator) Init() isDrawingFindIndicator {
	rv := objc.Send[isDrawingFindIndicator](i_.ID(), selInit)
	return rv
}
