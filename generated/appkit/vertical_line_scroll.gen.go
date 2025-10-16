
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [verticalLineScroll] class.
var verticalLineScrollClass _verticalLineScrollClass

func init() {
	verticalLineScrollClass = _verticalLineScrollClass{objc.GetClass("verticalLineScroll")}
}

type _verticalLineScrollClass struct {
	objc.Class
}

// An interface definition for the [verticalLineScroll] class.
type IverticalLineScroll interface {
	ID() objc.ID
}

type verticalLineScroll struct {
	id objc.ID
}

func verticalLineScrollFrom(ptr unsafe.Pointer) verticalLineScroll {
	return verticalLineScroll{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ verticalLineScroll) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _verticalLineScrollClass) Alloc() verticalLineScroll {
	rv := objc.Send[verticalLineScroll](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _verticalLineScrollClass) New() verticalLineScroll {
	rv := objc.Send[verticalLineScroll](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewverticalLineScroll creates and returns a new initialized instance.
func NewverticalLineScroll() verticalLineScroll {
	return verticalLineScrollClass.New()
}

// Init initializes the instance.
func (v_ verticalLineScroll) Init() verticalLineScroll {
	rv := objc.Send[verticalLineScroll](v_.ID(), selInit)
	return rv
}
