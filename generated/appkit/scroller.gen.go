
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Scroller] class.
var ScrollerClass _ScrollerClass

func init() {
	ScrollerClass = _ScrollerClass{objc.GetClass("NSScroller")}
}

type _ScrollerClass struct {
	objc.Class
}

// An interface definition for the [Scroller] class.
type IScroller interface {
	ID() objc.ID
}

type Scroller struct {
	id objc.ID
}

func ScrollerFrom(ptr unsafe.Pointer) Scroller {
	return Scroller{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ Scroller) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _ScrollerClass) Alloc() Scroller {
	rv := objc.Send[Scroller](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _ScrollerClass) New() Scroller {
	rv := objc.Send[Scroller](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewScroller creates and returns a new initialized instance.
func NewScroller() Scroller {
	return ScrollerClass.New()
}

// Init initializes the instance.
func (s_ Scroller) Init() Scroller {
	rv := objc.Send[Scroller](s_.ID(), selInit)
	return rv
}
