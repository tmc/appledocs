
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [layout] class.
var layoutClass _layoutClass

func init() {
	layoutClass = _layoutClass{objc.GetClass("layout")}
}

type _layoutClass struct {
	objc.Class
}

// An interface definition for the [layout] class.
type Ilayout interface {
	ID() objc.ID
}

type layout struct {
	id objc.ID
}

func layoutFrom(ptr unsafe.Pointer) layout {
	return layout{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ layout) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _layoutClass) Alloc() layout {
	rv := objc.Send[layout](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _layoutClass) New() layout {
	rv := objc.Send[layout](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newlayout creates and returns a new initialized instance.
func Newlayout() layout {
	return layoutClass.New()
}

// Init initializes the instance.
func (l_ layout) Init() layout {
	rv := objc.Send[layout](l_.ID(), selInit)
	return rv
}
