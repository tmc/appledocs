
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ImageView] class.
var ImageViewClass _ImageViewClass

func init() {
	ImageViewClass = _ImageViewClass{objc.GetClass("NSImageView")}
}

type _ImageViewClass struct {
	objc.Class
}

// An interface definition for the [ImageView] class.
type IImageView interface {
	ID() objc.ID
}

type ImageView struct {
	id objc.ID
}

func ImageViewFrom(ptr unsafe.Pointer) ImageView {
	return ImageView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ ImageView) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _ImageViewClass) Alloc() ImageView {
	rv := objc.Send[ImageView](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _ImageViewClass) New() ImageView {
	rv := objc.Send[ImageView](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewImageView creates and returns a new initialized instance.
func NewImageView() ImageView {
	return ImageViewClass.New()
}

// Init initializes the instance.
func (i_ ImageView) Init() ImageView {
	rv := objc.Send[ImageView](i_.ID(), selInit)
	return rv
}
