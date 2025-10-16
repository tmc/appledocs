
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [titleRect] class.
var titleRectClass _titleRectClass

func init() {
	titleRectClass = _titleRectClass{objc.GetClass("titleRect")}
}

type _titleRectClass struct {
	objc.Class
}

// An interface definition for the [titleRect] class.
type ItitleRect interface {
	ID() objc.ID
}

type titleRect struct {
	id objc.ID
}

func titleRectFrom(ptr unsafe.Pointer) titleRect {
	return titleRect{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ titleRect) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _titleRectClass) Alloc() titleRect {
	rv := objc.Send[titleRect](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _titleRectClass) New() titleRect {
	rv := objc.Send[titleRect](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtitleRect creates and returns a new initialized instance.
func NewtitleRect() titleRect {
	return titleRectClass.New()
}

// Init initializes the instance.
func (t_ titleRect) Init() titleRect {
	rv := objc.Send[titleRect](t_.ID(), selInit)
	return rv
}
