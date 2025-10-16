
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [viewsNeedDisplay] class.
var viewsNeedDisplayClass _viewsNeedDisplayClass

func init() {
	viewsNeedDisplayClass = _viewsNeedDisplayClass{objc.GetClass("viewsNeedDisplay")}
}

type _viewsNeedDisplayClass struct {
	objc.Class
}

// An interface definition for the [viewsNeedDisplay] class.
type IviewsNeedDisplay interface {
	ID() objc.ID
}

type viewsNeedDisplay struct {
	id objc.ID
}

func viewsNeedDisplayFrom(ptr unsafe.Pointer) viewsNeedDisplay {
	return viewsNeedDisplay{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ viewsNeedDisplay) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _viewsNeedDisplayClass) Alloc() viewsNeedDisplay {
	rv := objc.Send[viewsNeedDisplay](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _viewsNeedDisplayClass) New() viewsNeedDisplay {
	rv := objc.Send[viewsNeedDisplay](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewviewsNeedDisplay creates and returns a new initialized instance.
func NewviewsNeedDisplay() viewsNeedDisplay {
	return viewsNeedDisplayClass.New()
}

// Init initializes the instance.
func (v_ viewsNeedDisplay) Init() viewsNeedDisplay {
	rv := objc.Send[viewsNeedDisplay](v_.ID(), selInit)
	return rv
}
