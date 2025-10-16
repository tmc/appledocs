
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [horizontalScroller] class.
var horizontalScrollerClass _horizontalScrollerClass

func init() {
	horizontalScrollerClass = _horizontalScrollerClass{objc.GetClass("horizontalScroller")}
}

type _horizontalScrollerClass struct {
	objc.Class
}

// An interface definition for the [horizontalScroller] class.
type IhorizontalScroller interface {
	ID() objc.ID
}

type horizontalScroller struct {
	id objc.ID
}

func horizontalScrollerFrom(ptr unsafe.Pointer) horizontalScroller {
	return horizontalScroller{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (h_ horizontalScroller) ID() objc.ID {
	return h_.id
}

// Alloc allocates a new instance without initialization.
func (hc _horizontalScrollerClass) Alloc() horizontalScroller {
	rv := objc.Send[horizontalScroller](objc.ID(hc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (hc _horizontalScrollerClass) New() horizontalScroller {
	rv := objc.Send[horizontalScroller](objc.ID(hc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewhorizontalScroller creates and returns a new initialized instance.
func NewhorizontalScroller() horizontalScroller {
	return horizontalScrollerClass.New()
}

// Init initializes the instance.
func (h_ horizontalScroller) Init() horizontalScroller {
	rv := objc.Send[horizontalScroller](h_.ID(), selInit)
	return rv
}
