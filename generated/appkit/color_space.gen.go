// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ColorSpace] class.
var colorSpaceClass = _ColorSpaceClass{objc.GetClass("NSColorSpace")}

type _ColorSpaceClass struct {
	class objc.Class
}

// An interface definition for the [ColorSpace] class.
type IColorSpace interface {
	objectivec.IObject
}

// An object that represents a custom color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpace

type ColorSpace struct {
	objectivec.Object
}

// ColorSpaceFrom constructs a [ColorSpace] from an unsafe.Pointer.
//
// An object that represents a custom color space.
func ColorSpaceFrom(ptr unsafe.Pointer) ColorSpace {
	return ColorSpace{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (cc _ColorSpaceClass) Alloc() ColorSpace {
	rv := objc.Send[ColorSpace](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (cc _ColorSpaceClass) New() ColorSpace {
	rv := objc.Send[ColorSpace](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ColorSpace) Init() ColorSpace {
	rv := objc.Send[ColorSpace](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ColorSpace) Autorelease() ColorSpace {
	rv := objc.Send[ColorSpace](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewColorSpace creates a new ColorSpace instance.
func NewColorSpace() ColorSpace {
	return colorSpaceClass.New()
}




