
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [maskImage] class.
var maskImageClass _maskImageClass

func init() {
	maskImageClass = _maskImageClass{objc.GetClass("maskImage")}
}

type _maskImageClass struct {
	objc.Class
}

// An interface definition for the [maskImage] class.
type ImaskImage interface {
	ID() objc.ID
}

type maskImage struct {
	id objc.ID
}

func maskImageFrom(ptr unsafe.Pointer) maskImage {
	return maskImage{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ maskImage) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _maskImageClass) Alloc() maskImage {
	rv := objc.Send[maskImage](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _maskImageClass) New() maskImage {
	rv := objc.Send[maskImage](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewmaskImage creates and returns a new initialized instance.
func NewmaskImage() maskImage {
	return maskImageClass.New()
}

// Init initializes the instance.
func (m_ maskImage) Init() maskImage {
	rv := objc.Send[maskImage](m_.ID(), selInit)
	return rv
}
