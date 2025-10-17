
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ColorSpace] class.
var ColorSpaceClass _ColorSpaceClass

func init() {
	ColorSpaceClass = _ColorSpaceClass{objc.GetClass("NSColorSpace")}
}

type _ColorSpaceClass struct {
	objc.Class
}

// An interface definition for the [ColorSpace] class.
type IColorSpace interface {
	ID() objc.ID
}

type ColorSpace struct {
	id objc.ID
}

func ColorSpaceFrom(ptr unsafe.Pointer) ColorSpace {
	return ColorSpace{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ ColorSpace) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _ColorSpaceClass) Alloc() ColorSpace {
	rv := objc.Send[ColorSpace](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _ColorSpaceClass) New() ColorSpace {
	rv := objc.Send[ColorSpace](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewColorSpace creates and returns a new initialized instance.
func NewColorSpace() ColorSpace {
	return ColorSpaceClass.New()
}

// Init initializes the instance.
func (c_ ColorSpace) Init() ColorSpace {
	rv := objc.Send[ColorSpace](c_.ID(), selInit)
	return rv
}
