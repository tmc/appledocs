
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isCompatibleWithResponsiveScrolling] class.
var isCompatibleWithResponsiveScrollingClass _isCompatibleWithResponsiveScrollingClass

func init() {
	isCompatibleWithResponsiveScrollingClass = _isCompatibleWithResponsiveScrollingClass{objc.GetClass("isCompatibleWithResponsiveScrolling")}
}

type _isCompatibleWithResponsiveScrollingClass struct {
	objc.Class
}

// An interface definition for the [isCompatibleWithResponsiveScrolling] class.
type IisCompatibleWithResponsiveScrolling interface {
	ID() objc.ID
}

type isCompatibleWithResponsiveScrolling struct {
	id objc.ID
}

func isCompatibleWithResponsiveScrollingFrom(ptr unsafe.Pointer) isCompatibleWithResponsiveScrolling {
	return isCompatibleWithResponsiveScrolling{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isCompatibleWithResponsiveScrolling) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isCompatibleWithResponsiveScrollingClass) Alloc() isCompatibleWithResponsiveScrolling {
	rv := objc.Send[isCompatibleWithResponsiveScrolling](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isCompatibleWithResponsiveScrollingClass) New() isCompatibleWithResponsiveScrolling {
	rv := objc.Send[isCompatibleWithResponsiveScrolling](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisCompatibleWithResponsiveScrolling creates and returns a new initialized instance.
func NewisCompatibleWithResponsiveScrolling() isCompatibleWithResponsiveScrolling {
	return isCompatibleWithResponsiveScrollingClass.New()
}

// Init initializes the instance.
func (i_ isCompatibleWithResponsiveScrolling) Init() isCompatibleWithResponsiveScrolling {
	rv := objc.Send[isCompatibleWithResponsiveScrolling](i_.ID(), selInit)
	return rv
}
