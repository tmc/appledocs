
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [discardCursorRects] class.
var discardCursorRectsClass _discardCursorRectsClass

func init() {
	discardCursorRectsClass = _discardCursorRectsClass{objc.GetClass("discardCursorRects")}
}

type _discardCursorRectsClass struct {
	objc.Class
}

// An interface definition for the [discardCursorRects] class.
type IdiscardCursorRects interface {
	ID() objc.ID
}

type discardCursorRects struct {
	id objc.ID
}

func discardCursorRectsFrom(ptr unsafe.Pointer) discardCursorRects {
	return discardCursorRects{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ discardCursorRects) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _discardCursorRectsClass) Alloc() discardCursorRects {
	rv := objc.Send[discardCursorRects](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _discardCursorRectsClass) New() discardCursorRects {
	rv := objc.Send[discardCursorRects](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdiscardCursorRects creates and returns a new initialized instance.
func NewdiscardCursorRects() discardCursorRects {
	return discardCursorRectsClass.New()
}

// Init initializes the instance.
func (d_ discardCursorRects) Init() discardCursorRects {
	rv := objc.Send[discardCursorRects](d_.ID(), selInit)
	return rv
}
