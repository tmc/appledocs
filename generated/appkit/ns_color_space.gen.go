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

/* debug [class.gen.go]: Generating class NSColorSpace */


/* debug [class_header]: Header for NSColorSpace */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ColorSpace */
// An interface definition for the [ColorSpace] class.
type IColorSpace interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ColorSpace */
	// properties:
	CgColorSpace() IColorSpace
	SetCgColorSpace(value IColorSpace)
	ColorSpaceModel() coreml.Model
	SetColorSpaceModel(value coreml.Model)
	ColorSyncProfile() objectivec.IObject
	SetColorSyncProfile(value objectivec.IObject)
	IccProfileData() foundation.Data
	SetIccProfileData(value foundation.Data)
	LocalizedName() objc.IObject /* cross-framework: NSString */
	SetLocalizedName(value objc.IObject /* cross-framework: NSString */)
	NumberOfColorComponents() int
	SetNumberOfColorComponents(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ColorSpace */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ColorSpace */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ColorSpace */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ColorSpace *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ColorSpace */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ColorSpace */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ColorSpace */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ColorSpace */

// The Core Graphics color-space object that represents a color space equivalent to the color space’s.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/cgcolorspace
func (c_ ColorSpace) CgColorSpace() IColorSpace {
	rv := objc.Send[ColorSpace](c_.ID, objc.Sel("cgColorSpace"))
	return rv
}/* debug [instance_properties/getter]: cgColorSpace */


// The Core Graphics color-space object that represents a color space equivalent to the color space’s.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/cgcolorspace
func (c_ ColorSpace) SetCgColorSpace(value IColorSpace) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCgColorSpace:"), value)
}/* debug [instance_properties/setter]: cgColorSpace */


// The model on which the color space is based.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/colorspacemodel
func (c_ ColorSpace) ColorSpaceModel() coreml.Model {
	rv := objc.Send[coreml.Model](c_.ID, objc.Sel("colorSpaceModel"))
	return rv
}/* debug [instance_properties/getter]: colorSpaceModel */


// The model on which the color space is based.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/colorspacemodel
func (c_ ColorSpace) SetColorSpaceModel(value coreml.Model) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setColorSpaceModel:"), value)
}/* debug [instance_properties/setter]: colorSpaceModel */


// The ColorSync profile from which the color space was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/colorsyncprofile
func (c_ ColorSpace) ColorSyncProfile() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("colorSyncProfile"))
	return rv
}/* debug [instance_properties/getter]: colorSyncProfile */


// The ColorSync profile from which the color space was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/colorsyncprofile
func (c_ ColorSpace) SetColorSyncProfile(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setColorSyncProfile:"), value)
}/* debug [instance_properties/setter]: colorSyncProfile */


// The ICC profile data from which the color space was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/iccprofiledata
func (c_ ColorSpace) IccProfileData() foundation.Data {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("iccProfileData"))
	return rv
}/* debug [instance_properties/getter]: iccProfileData */


// The ICC profile data from which the color space was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/iccprofiledata
func (c_ ColorSpace) SetIccProfileData(value foundation.Data) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIccProfileData:"), value)
}/* debug [instance_properties/setter]: iccProfileData */


// The localized name of the color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/localizedname
func (c_ ColorSpace) LocalizedName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("localizedName"))
	return rv
}/* debug [instance_properties/getter]: localizedName */


// The localized name of the color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/localizedname
func (c_ ColorSpace) SetLocalizedName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLocalizedName:"), value)
}/* debug [instance_properties/setter]: localizedName */


// The number of components, excluding alpha, the color space supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/numberofcolorcomponents
func (c_ ColorSpace) NumberOfColorComponents() int {
	rv := objc.Send[int](c_.ID, objc.Sel("numberOfColorComponents"))
	return rv
}/* debug [instance_properties/getter]: numberOfColorComponents */


// The number of components, excluding alpha, the color space supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspace/numberofcolorcomponents
func (c_ ColorSpace) SetNumberOfColorComponents(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNumberOfColorComponents:"), value)
}/* debug [instance_properties/setter]: numberOfColorComponents */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSColorSpace */



