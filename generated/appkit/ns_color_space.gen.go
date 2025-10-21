// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ColorSpace] class.
var (
	ColorSpaceClass     _ColorSpaceClass
	ColorSpaceClassOnce sync.Once
)

func getColorSpaceClass() _ColorSpaceClass {
	ColorSpaceClassOnce.Do(func() {
		ColorSpaceClass = _ColorSpaceClass{objc.GetClass("NSColorSpace")}
	})
	return ColorSpaceClass
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

// Initializes and returns a color space object initialized from a Core Graphics color-space object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpace/init(cgColorSpace:)
func NewColorSpaceWithCGColorSpace(cgColorSpace coregraphics.CGColorSpaceRef) ColorSpace {
	instance := getColorSpaceClass().Alloc()
	rv := objc.Send[ColorSpace](instance.ID, objc.Sel("initWithCGColorSpace:"), cgColorSpace)
	rv.Autorelease()
	return rv
}

// Initializes and returns a color space object from the specified ColorSync profile.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpace/init(colorSyncProfile:)
func NewColorSpaceWithColorSyncProfile(prof unsafe.Pointer) ColorSpace {
	instance := getColorSpaceClass().Alloc()
	rv := objc.Send[ColorSpace](instance.ID, objc.Sel("initWithColorSyncProfile:"), prof)
	rv.Autorelease()
	return rv
}

// Initializes and returns a color space object from the specified ICC profile.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpace/init(iccProfileData:)
func NewColorSpaceWithICCProfileData(iccData unsafe.Pointer) ColorSpace {
	instance := getColorSpaceClass().Alloc()
	rv := objc.Send[ColorSpace](instance.ID, objc.Sel("initWithICCProfileData:"), iccData)
	rv.Autorelease()
	return rv
}

// Returns the list of color spaces available on the system that are displayed in the color panel, in the order they are displayed in the color panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpace/availableColorSpaces(with:)
func (cc _ColorSpaceClass) AvailableColorSpacesWithModel(model unsafe.Pointer) []ColorSpace {
	rv := objc.Send[[]ColorSpace](objc.ID(cc.class), objc.Sel("availableColorSpacesWithModel:"), model)
	return rv
}

// A color space object that represents an Adobe RGB (1998) color space.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpace/adobeRGB1998
func (cc _ColorSpaceClass) AdobeRGB1998ColorSpace() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("adobeRGB1998ColorSpace"))
	return rv
}

// A color space object that represents a calibrated or device-dependent gray color space.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpace/deviceGray
func (cc _ColorSpaceClass) DeviceGrayColorSpace() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("deviceGrayColorSpace"))
	return rv
}

// A color space object that represents an extended gray color space with a gamma value of 2.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpace/extendedGenericGamma22Gray
func (cc _ColorSpaceClass) ExtendedGenericGamma22GrayColorSpace() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("extendedGenericGamma22GrayColorSpace"))
	return rv
}

// A color space object that represents an extended sRGB color space.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpace/extendedSRGB
func (cc _ColorSpaceClass) ExtendedSRGBColorSpace() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("extendedSRGBColorSpace"))
	return rv
}

// A color space object that represents an Adobe RGB (1998) color space.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpace/adobeRGB1998
func (c_ ColorSpace) AdobeRGB1998ColorSpace() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("adobeRGB1998ColorSpace"))
	return rv
}

// A color space object that represents a calibrated or device-dependent gray color space.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpace/deviceGray
func (c_ ColorSpace) DeviceGrayColorSpace() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("deviceGrayColorSpace"))
	return rv
}

// A color space object that represents an extended gray color space with a gamma value of 2.2.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpace/extendedGenericGamma22Gray
func (c_ ColorSpace) ExtendedGenericGamma22GrayColorSpace() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("extendedGenericGamma22GrayColorSpace"))
	return rv
}

// A color space object that represents an extended sRGB color space.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpace/extendedSRGB
func (c_ ColorSpace) ExtendedSRGBColorSpace() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("extendedSRGBColorSpace"))
	return rv
}
