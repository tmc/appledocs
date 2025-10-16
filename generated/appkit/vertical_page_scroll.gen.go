
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [verticalPageScroll] class.
var verticalPageScrollClass _verticalPageScrollClass

func init() {
	verticalPageScrollClass = _verticalPageScrollClass{objc.GetClass("verticalPageScroll")}
}

type _verticalPageScrollClass struct {
	objc.Class
}

// An interface definition for the [verticalPageScroll] class.
type IverticalPageScroll interface {
	ID() objc.ID
}

type verticalPageScroll struct {
	id objc.ID
}

func verticalPageScrollFrom(ptr unsafe.Pointer) verticalPageScroll {
	return verticalPageScroll{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ verticalPageScroll) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _verticalPageScrollClass) Alloc() verticalPageScroll {
	rv := objc.Send[verticalPageScroll](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _verticalPageScrollClass) New() verticalPageScroll {
	rv := objc.Send[verticalPageScroll](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewverticalPageScroll creates and returns a new initialized instance.
func NewverticalPageScroll() verticalPageScroll {
	return verticalPageScrollClass.New()
}

// Init initializes the instance.
func (v_ verticalPageScroll) Init() verticalPageScroll {
	rv := objc.Send[verticalPageScroll](v_.ID(), selInit)
	return rv
}
