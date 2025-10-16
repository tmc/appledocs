
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [disableCursorRects] class.
var disableCursorRectsClass _disableCursorRectsClass

func init() {
	disableCursorRectsClass = _disableCursorRectsClass{objc.GetClass("disableCursorRects")}
}

type _disableCursorRectsClass struct {
	objc.Class
}

// An interface definition for the [disableCursorRects] class.
type IdisableCursorRects interface {
	ID() objc.ID
}

type disableCursorRects struct {
	id objc.ID
}

func disableCursorRectsFrom(ptr unsafe.Pointer) disableCursorRects {
	return disableCursorRects{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ disableCursorRects) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _disableCursorRectsClass) Alloc() disableCursorRects {
	rv := objc.Send[disableCursorRects](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _disableCursorRectsClass) New() disableCursorRects {
	rv := objc.Send[disableCursorRects](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdisableCursorRects creates and returns a new initialized instance.
func NewdisableCursorRects() disableCursorRects {
	return disableCursorRectsClass.New()
}

// Init initializes the instance.
func (d_ disableCursorRects) Init() disableCursorRects {
	rv := objc.Send[disableCursorRects](d_.ID(), selInit)
	return rv
}
