
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [horizontalLineScroll] class.
var horizontalLineScrollClass _horizontalLineScrollClass

func init() {
	horizontalLineScrollClass = _horizontalLineScrollClass{objc.GetClass("horizontalLineScroll")}
}

type _horizontalLineScrollClass struct {
	objc.Class
}

// An interface definition for the [horizontalLineScroll] class.
type IhorizontalLineScroll interface {
	ID() objc.ID
}

type horizontalLineScroll struct {
	id objc.ID
}

func horizontalLineScrollFrom(ptr unsafe.Pointer) horizontalLineScroll {
	return horizontalLineScroll{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (h_ horizontalLineScroll) ID() objc.ID {
	return h_.id
}

// Alloc allocates a new instance without initialization.
func (hc _horizontalLineScrollClass) Alloc() horizontalLineScroll {
	rv := objc.Send[horizontalLineScroll](objc.ID(hc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (hc _horizontalLineScrollClass) New() horizontalLineScroll {
	rv := objc.Send[horizontalLineScroll](objc.ID(hc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewhorizontalLineScroll creates and returns a new initialized instance.
func NewhorizontalLineScroll() horizontalLineScroll {
	return horizontalLineScrollClass.New()
}

// Init initializes the instance.
func (h_ horizontalLineScroll) Init() horizontalLineScroll {
	rv := objc.Send[horizontalLineScroll](h_.ID(), selInit)
	return rv
}
