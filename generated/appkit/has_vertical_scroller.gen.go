
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [hasVerticalScroller] class.
var hasVerticalScrollerClass _hasVerticalScrollerClass

func init() {
	hasVerticalScrollerClass = _hasVerticalScrollerClass{objc.GetClass("hasVerticalScroller")}
}

type _hasVerticalScrollerClass struct {
	objc.Class
}

// An interface definition for the [hasVerticalScroller] class.
type IhasVerticalScroller interface {
	ID() objc.ID
}

type hasVerticalScroller struct {
	id objc.ID
}

func hasVerticalScrollerFrom(ptr unsafe.Pointer) hasVerticalScroller {
	return hasVerticalScroller{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (h_ hasVerticalScroller) ID() objc.ID {
	return h_.id
}

// Alloc allocates a new instance without initialization.
func (hc _hasVerticalScrollerClass) Alloc() hasVerticalScroller {
	rv := objc.Send[hasVerticalScroller](objc.ID(hc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (hc _hasVerticalScrollerClass) New() hasVerticalScroller {
	rv := objc.Send[hasVerticalScroller](objc.ID(hc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewhasVerticalScroller creates and returns a new initialized instance.
func NewhasVerticalScroller() hasVerticalScroller {
	return hasVerticalScrollerClass.New()
}

// Init initializes the instance.
func (h_ hasVerticalScroller) Init() hasVerticalScroller {
	rv := objc.Send[hasVerticalScroller](h_.ID(), selInit)
	return rv
}
