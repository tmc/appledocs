// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ColorSpace] class.
var (
	colorSpaceClass     _ColorSpaceClass
	colorSpaceClassOnce sync.Once
)

func getColorSpaceClass() _ColorSpaceClass {
	colorSpaceClassOnce.Do(func() {
		colorSpaceClass = _ColorSpaceClass{objc.GetClass("NSColorSpace")}
	})
	return colorSpaceClass
}

type _ColorSpaceClass struct {
	class objc.Class
}

// An interface definition for the [ColorSpace] class.
type IColorSpace interface {
	objectivec.IObject
}

// An object that represents a custom color space.
//
// You can make custom color spaces from ColorSync profiles or from ICC profiles. also has factory methods that return objects representing the system color spaces. You can use the method of the class to create color objects using custom objects. You can also send the message to an object to convert it between two color spaces, either of which may be a custom color space.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getColorSpaceClass().New()
}




