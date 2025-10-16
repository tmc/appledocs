
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [enableCursorRects] class.
var enableCursorRectsClass _enableCursorRectsClass

func init() {
	enableCursorRectsClass = _enableCursorRectsClass{objc.GetClass("enableCursorRects")}
}

type _enableCursorRectsClass struct {
	objc.Class
}

// An interface definition for the [enableCursorRects] class.
type IenableCursorRects interface {
	ID() objc.ID
}

type enableCursorRects struct {
	id objc.ID
}

func enableCursorRectsFrom(ptr unsafe.Pointer) enableCursorRects {
	return enableCursorRects{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (e_ enableCursorRects) ID() objc.ID {
	return e_.id
}

// Alloc allocates a new instance without initialization.
func (ec _enableCursorRectsClass) Alloc() enableCursorRects {
	rv := objc.Send[enableCursorRects](objc.ID(ec.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ec _enableCursorRectsClass) New() enableCursorRects {
	rv := objc.Send[enableCursorRects](objc.ID(ec.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewenableCursorRects creates and returns a new initialized instance.
func NewenableCursorRects() enableCursorRects {
	return enableCursorRectsClass.New()
}

// Init initializes the instance.
func (e_ enableCursorRects) Init() enableCursorRects {
	rv := objc.Send[enableCursorRects](e_.ID(), selInit)
	return rv
}
