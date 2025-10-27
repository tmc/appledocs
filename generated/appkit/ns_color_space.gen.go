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
	CgColorSpace() IColorSpace
	SetCgColorSpace(value IColorSpace)
	ColorSpaceModel() coreml.Model
	SetColorSpaceModel(value coreml.Model)
	ColorSyncProfile() objectivec.IObject
	SetColorSyncProfile(value objectivec.IObject)
	IccProfileData() foundation.Data
	SetIccProfileData(value foundation.Data)
	LocalizedName() foundation.foundation.INSString
	SetLocalizedName(value foundation.foundation.INSString)
	NumberOfColorComponents() int
	SetNumberOfColorComponents(value int)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _ColorSpaceClass) Alloc() ColorSpace {
	rv := objc.Send[ColorSpace](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

























// The Core Graphics color-space object that represents a color space equivalent to the color space’s.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/cgcolorspace
func (c_ ColorSpace) CgColorSpace() IColorSpace {
	rv := objc.Send[ColorSpace](c_.ID, objc.Sel("cgColorSpace"))
	return rv
}


// The Core Graphics color-space object that represents a color space equivalent to the color space’s.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/cgcolorspace
func (c_ ColorSpace) SetCgColorSpace(value IColorSpace) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCgColorSpace:"), value)
}


// The model on which the color space is based.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/colorspacemodel
func (c_ ColorSpace) ColorSpaceModel() coreml.Model {
	rv := objc.Send[coreml.Model](c_.ID, objc.Sel("colorSpaceModel"))
	return rv
}


// The model on which the color space is based.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/colorspacemodel
func (c_ ColorSpace) SetColorSpaceModel(value coreml.Model) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setColorSpaceModel:"), value)
}


// The ColorSync profile from which the color space was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/colorsyncprofile
func (c_ ColorSpace) ColorSyncProfile() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("colorSyncProfile"))
	return rv
}


// The ColorSync profile from which the color space was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/colorsyncprofile
func (c_ ColorSpace) SetColorSyncProfile(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setColorSyncProfile:"), value)
}


// The ICC profile data from which the color space was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/iccprofiledata
func (c_ ColorSpace) IccProfileData() foundation.Data {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("iccProfileData"))
	return rv
}


// The ICC profile data from which the color space was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/iccprofiledata
func (c_ ColorSpace) SetIccProfileData(value foundation.Data) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIccProfileData:"), value)
}


// The localized name of the color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/localizedname
func (c_ ColorSpace) LocalizedName() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("localizedName"))
	return rv
}


// The localized name of the color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/localizedname
func (c_ ColorSpace) SetLocalizedName(value foundation.foundation.INSString) {
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








