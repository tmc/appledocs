// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coreml"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	CGColorSpace() ColorSpaceRef /* not a class type */
	ColorSyncProfile() unsafe.Pointer
	ColorSpaceModel() objc.IObject /* cross-framework: Model */
	SetColorSpaceModel(value objc.IObject /* cross-framework: Model */)
	IccProfileData() objc.IObject /* cross-framework: Data */
	SetIccProfileData(value objc.IObject /* cross-framework: Data */)
	LocalizedName() objc.IObject /* cross-framework: NSString */
	SetLocalizedName(value objc.IObject /* cross-framework: NSString */)
	NumberOfColorComponents() int
	SetNumberOfColorComponents(value int)
	// methods:
}

// An object that represents a custom color space.
//
// You can make custom color spaces from ColorSync profiles or from ICC profiles. also has factory methods that return objects representing the system color spaces. You can use the method of the class to create color objects using custom objects. You can also send the message to an object to convert it between two color spaces, either of which may be a custom color space.


// An object that represents a custom color space.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpace/init(cgColorSpace:)
func NewColorSpaceWithCGColorSpace(cgColorSpace ColorSpaceRef /* not a class type */) ColorSpace {
	instance := getColorSpaceClass().Alloc()
	rv := objc.Send[ColorSpace](instance.ID, objc.Sel("initWithCGColorSpace:"), cgColorSpace)
	rv.Autorelease()
	return rv
}



// Returns the list of color spaces available on the system that are displayed in the color panel, in the order they are displayed in the color panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpace/availableColorSpaces(with:)
func (cc _ColorSpaceClass) AvailableColorSpacesWithModel(model ColorSpaceModel /* not a class type */) []IColorSpace {
	rv := objc.Send[[]ColorSpace](objc.ID(cc.class), objc.Sel("availableColorSpacesWithModel:"), model)
	return rv
}


// A color space object that represents an extended gray color space with a gamma value of 2.2.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpace/extendedGenericGamma22Gray
func (cc _ColorSpaceClass) ExtendedGenericGamma22GrayColorSpace() ColorSpace {
	rv := objc.Send[ColorSpace](objc.ID(cc.class), objc.Sel("extendedGenericGamma22GrayColorSpace"))
	return rv
}

// A color space object that represents an extended sRGB color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpace/extendedSRGB
func (cc _ColorSpaceClass) ExtendedSRGBColorSpace() ColorSpace {
	rv := objc.Send[ColorSpace](objc.ID(cc.class), objc.Sel("extendedSRGBColorSpace"))
	return rv
}

// A color space object that represents a device-independent gray color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpace/genericGray
func (cc _ColorSpaceClass) GenericGrayColorSpace() ColorSpace {
	rv := objc.Send[ColorSpace](objc.ID(cc.class), objc.Sel("genericGrayColorSpace"))
	return rv
}

// The Core Graphics color-space object that represents a color space equivalent to the color space’s.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpace/cgColorSpace
func (c_ ColorSpace) CGColorSpace() ColorSpaceRef /* not a class type */ {
	rv := objc.Send[ColorSpaceRef](c_.ID, objc.Sel("CGColorSpace"))
	return rv
}


// The ColorSync profile from which the color space was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpace/colorSyncProfile
func (c_ ColorSpace) ColorSyncProfile() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("colorSyncProfile"))
	return rv
}


// A color space object that represents an extended gray color space with a gamma value of 2.2.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpace/extendedGenericGamma22Gray
func (c_ ColorSpace) ExtendedGenericGamma22GrayColorSpace() IColorSpace {
	rv := objc.Send[ColorSpace](c_.ID, objc.Sel("extendedGenericGamma22GrayColorSpace"))
	return rv
}


// A color space object that represents an extended sRGB color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpace/extendedSRGB
func (c_ ColorSpace) ExtendedSRGBColorSpace() IColorSpace {
	rv := objc.Send[ColorSpace](c_.ID, objc.Sel("extendedSRGBColorSpace"))
	return rv
}


// A color space object that represents a device-independent gray color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpace/genericGray
func (c_ ColorSpace) GenericGrayColorSpace() IColorSpace {
	rv := objc.Send[ColorSpace](c_.ID, objc.Sel("genericGrayColorSpace"))
	return rv
}


// The model on which the color space is based.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/colorspacemodel
func (c_ ColorSpace) ColorSpaceModel() objc.IObject /* cross-framework: Model */ {
	rv := objc.Send[coreml.Model](c_.ID, objc.Sel("colorSpaceModel"))
	return rv
}


// The model on which the color space is based.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/colorspacemodel
func (c_ ColorSpace) SetColorSpaceModel(value objc.IObject /* cross-framework: Model */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setColorSpaceModel:"), value)
}


// The ICC profile data from which the color space was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/iccprofiledata
func (c_ ColorSpace) IccProfileData() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("iccProfileData"))
	return rv
}


// The ICC profile data from which the color space was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/iccprofiledata
func (c_ ColorSpace) SetIccProfileData(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIccProfileData:"), value)
}


// The localized name of the color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/localizedname
func (c_ ColorSpace) LocalizedName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("localizedName"))
	return rv
}


// The localized name of the color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/localizedname
func (c_ ColorSpace) SetLocalizedName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLocalizedName:"), value)
}


// The number of components, excluding alpha, the color space supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/numberofcolorcomponents
func (c_ ColorSpace) NumberOfColorComponents() int {
	rv := objc.Send[int](c_.ID, objc.Sel("numberOfColorComponents"))
	return rv
}


// The number of components, excluding alpha, the color space supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/numberofcolorcomponents
func (c_ ColorSpace) SetNumberOfColorComponents(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNumberOfColorComponents:"), value)
}


