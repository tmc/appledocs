
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [horizontalPageScroll] class.
var horizontalPageScrollClass _horizontalPageScrollClass

func init() {
	horizontalPageScrollClass = _horizontalPageScrollClass{objc.GetClass("horizontalPageScroll")}
}

type _horizontalPageScrollClass struct {
	objc.Class
}

// An interface definition for the [horizontalPageScroll] class.
type IhorizontalPageScroll interface {
	ID() objc.ID
}

type horizontalPageScroll struct {
	id objc.ID
}

func horizontalPageScrollFrom(ptr unsafe.Pointer) horizontalPageScroll {
	return horizontalPageScroll{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (h_ horizontalPageScroll) ID() objc.ID {
	return h_.id
}

// Alloc allocates a new instance without initialization.
func (hc _horizontalPageScrollClass) Alloc() horizontalPageScroll {
	rv := objc.Send[horizontalPageScroll](objc.ID(hc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (hc _horizontalPageScrollClass) New() horizontalPageScroll {
	rv := objc.Send[horizontalPageScroll](objc.ID(hc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewhorizontalPageScroll creates and returns a new initialized instance.
func NewhorizontalPageScroll() horizontalPageScroll {
	return horizontalPageScrollClass.New()
}

// Init initializes the instance.
func (h_ horizontalPageScroll) Init() horizontalPageScroll {
	rv := objc.Send[horizontalPageScroll](h_.ID(), selInit)
	return rv
}
