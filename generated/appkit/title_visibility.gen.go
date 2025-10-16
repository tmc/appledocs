
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [titleVisibility] class.
var titleVisibilityClass _titleVisibilityClass

func init() {
	titleVisibilityClass = _titleVisibilityClass{objc.GetClass("titleVisibility")}
}

type _titleVisibilityClass struct {
	objc.Class
}

// An interface definition for the [titleVisibility] class.
type ItitleVisibility interface {
	ID() objc.ID
}

type titleVisibility struct {
	id objc.ID
}

func titleVisibilityFrom(ptr unsafe.Pointer) titleVisibility {
	return titleVisibility{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ titleVisibility) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _titleVisibilityClass) Alloc() titleVisibility {
	rv := objc.Send[titleVisibility](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _titleVisibilityClass) New() titleVisibility {
	rv := objc.Send[titleVisibility](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtitleVisibility creates and returns a new initialized instance.
func NewtitleVisibility() titleVisibility {
	return titleVisibilityClass.New()
}

// Init initializes the instance.
func (t_ titleVisibility) Init() titleVisibility {
	rv := objc.Send[titleVisibility](t_.ID(), selInit)
	return rv
}
