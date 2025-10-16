
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [resetCursorRects] class.
var resetCursorRectsClass _resetCursorRectsClass

func init() {
	resetCursorRectsClass = _resetCursorRectsClass{objc.GetClass("resetCursorRects")}
}

type _resetCursorRectsClass struct {
	objc.Class
}

// An interface definition for the [resetCursorRects] class.
type IresetCursorRects interface {
	ID() objc.ID
}

type resetCursorRects struct {
	id objc.ID
}

func resetCursorRectsFrom(ptr unsafe.Pointer) resetCursorRects {
	return resetCursorRects{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ resetCursorRects) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _resetCursorRectsClass) Alloc() resetCursorRects {
	rv := objc.Send[resetCursorRects](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _resetCursorRectsClass) New() resetCursorRects {
	rv := objc.Send[resetCursorRects](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewresetCursorRects creates and returns a new initialized instance.
func NewresetCursorRects() resetCursorRects {
	return resetCursorRectsClass.New()
}

// Init initializes the instance.
func (r_ resetCursorRects) Init() resetCursorRects {
	rv := objc.Send[resetCursorRects](r_.ID(), selInit)
	return rv
}
