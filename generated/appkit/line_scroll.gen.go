
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [lineScroll] class.
var lineScrollClass _lineScrollClass

func init() {
	lineScrollClass = _lineScrollClass{objc.GetClass("lineScroll")}
}

type _lineScrollClass struct {
	objc.Class
}

// An interface definition for the [lineScroll] class.
type IlineScroll interface {
	ID() objc.ID
}

type lineScroll struct {
	id objc.ID
}

func lineScrollFrom(ptr unsafe.Pointer) lineScroll {
	return lineScroll{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ lineScroll) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _lineScrollClass) Alloc() lineScroll {
	rv := objc.Send[lineScroll](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _lineScrollClass) New() lineScroll {
	rv := objc.Send[lineScroll](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewlineScroll creates and returns a new initialized instance.
func NewlineScroll() lineScroll {
	return lineScrollClass.New()
}

// Init initializes the instance.
func (l_ lineScroll) Init() lineScroll {
	rv := objc.Send[lineScroll](l_.ID(), selInit)
	return rv
}
