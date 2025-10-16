
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [hasHorizontalScroller] class.
var hasHorizontalScrollerClass _hasHorizontalScrollerClass

func init() {
	hasHorizontalScrollerClass = _hasHorizontalScrollerClass{objc.GetClass("hasHorizontalScroller")}
}

type _hasHorizontalScrollerClass struct {
	objc.Class
}

// An interface definition for the [hasHorizontalScroller] class.
type IhasHorizontalScroller interface {
	ID() objc.ID
}

type hasHorizontalScroller struct {
	id objc.ID
}

func hasHorizontalScrollerFrom(ptr unsafe.Pointer) hasHorizontalScroller {
	return hasHorizontalScroller{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (h_ hasHorizontalScroller) ID() objc.ID {
	return h_.id
}

// Alloc allocates a new instance without initialization.
func (hc _hasHorizontalScrollerClass) Alloc() hasHorizontalScroller {
	rv := objc.Send[hasHorizontalScroller](objc.ID(hc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (hc _hasHorizontalScrollerClass) New() hasHorizontalScroller {
	rv := objc.Send[hasHorizontalScroller](objc.ID(hc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewhasHorizontalScroller creates and returns a new initialized instance.
func NewhasHorizontalScroller() hasHorizontalScroller {
	return hasHorizontalScrollerClass.New()
}

// Init initializes the instance.
func (h_ hasHorizontalScroller) Init() hasHorizontalScroller {
	rv := objc.Send[hasHorizontalScroller](h_.ID(), selInit)
	return rv
}
