
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [imageAlignment] class.
var imageAlignmentClass _imageAlignmentClass

func init() {
	imageAlignmentClass = _imageAlignmentClass{objc.GetClass("imageAlignment")}
}

type _imageAlignmentClass struct {
	objc.Class
}

// An interface definition for the [imageAlignment] class.
type IimageAlignment interface {
	ID() objc.ID
}

type imageAlignment struct {
	id objc.ID
}

func imageAlignmentFrom(ptr unsafe.Pointer) imageAlignment {
	return imageAlignment{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ imageAlignment) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _imageAlignmentClass) Alloc() imageAlignment {
	rv := objc.Send[imageAlignment](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _imageAlignmentClass) New() imageAlignment {
	rv := objc.Send[imageAlignment](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewimageAlignment creates and returns a new initialized instance.
func NewimageAlignment() imageAlignment {
	return imageAlignmentClass.New()
}

// Init initializes the instance.
func (i_ imageAlignment) Init() imageAlignment {
	rv := objc.Send[imageAlignment](i_.ID(), selInit)
	return rv
}
