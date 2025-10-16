
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ImageRep] class.
var ImageRepClass _ImageRepClass

func init() {
	ImageRepClass = _ImageRepClass{objc.GetClass("NSImageRep")}
}

type _ImageRepClass struct {
	objc.Class
}

// An interface definition for the [ImageRep] class.
type IImageRep interface {
	ID() objc.ID
}

type ImageRep struct {
	id objc.ID
}

func ImageRepFrom(ptr unsafe.Pointer) ImageRep {
	return ImageRep{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ ImageRep) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _ImageRepClass) Alloc() ImageRep {
	rv := objc.Send[ImageRep](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _ImageRepClass) New() ImageRep {
	rv := objc.Send[ImageRep](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewImageRep creates and returns a new initialized instance.
func NewImageRep() ImageRep {
	return ImageRepClass.New()
}

// Init initializes the instance.
func (i_ ImageRep) Init() ImageRep {
	rv := objc.Send[ImageRep](i_.ID(), selInit)
	return rv
}
