
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [views] class.
var viewsClass _viewsClass

func init() {
	viewsClass = _viewsClass{objc.GetClass("views")}
}

type _viewsClass struct {
	objc.Class
}

// An interface definition for the [views] class.
type Iviews interface {
	ID() objc.ID
}

type views struct {
	id objc.ID
}

func viewsFrom(ptr unsafe.Pointer) views {
	return views{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ views) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _viewsClass) Alloc() views {
	rv := objc.Send[views](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _viewsClass) New() views {
	rv := objc.Send[views](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newviews creates and returns a new initialized instance.
func Newviews() views {
	return viewsClass.New()
}

// Init initializes the instance.
func (v_ views) Init() views {
	rv := objc.Send[views](v_.ID(), selInit)
	return rv
}
