
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Font] class.
var FontClass _FontClass

func init() {
	FontClass = _FontClass{objc.GetClass("NSFont")}
}

type _FontClass struct {
	objc.Class
}

// An interface definition for the [Font] class.
type IFont interface {
	ID() objc.ID
}

type Font struct {
	id objc.ID
}

func FontFrom(ptr unsafe.Pointer) Font {
	return Font{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ Font) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _FontClass) Alloc() Font {
	rv := objc.Send[Font](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _FontClass) New() Font {
	rv := objc.Send[Font](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewFont creates and returns a new initialized instance.
func NewFont() Font {
	return FontClass.New()
}

// Init initializes the instance.
func (f_ Font) Init() Font {
	rv := objc.Send[Font](f_.ID(), selInit)
	return rv
}
