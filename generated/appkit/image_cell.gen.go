
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ImageCell] class.
var ImageCellClass _ImageCellClass

func init() {
	ImageCellClass = _ImageCellClass{objc.GetClass("NSImageCell")}
}

type _ImageCellClass struct {
	objc.Class
}

// An interface definition for the [ImageCell] class.
type IImageCell interface {
	ID() objc.ID
}

type ImageCell struct {
	id objc.ID
}

func ImageCellFrom(ptr unsafe.Pointer) ImageCell {
	return ImageCell{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ ImageCell) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _ImageCellClass) Alloc() ImageCell {
	rv := objc.Send[ImageCell](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _ImageCellClass) New() ImageCell {
	rv := objc.Send[ImageCell](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewImageCell creates and returns a new initialized instance.
func NewImageCell() ImageCell {
	return ImageCellClass.New()
}

// Init initializes the instance.
func (i_ ImageCell) Init() ImageCell {
	rv := objc.Send[ImageCell](i_.ID(), selInit)
	return rv
}
