// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSImageSymbolConfiguration */


/* debug [class_header]: Header for NSImageSymbolConfiguration */
// The class instance for the [ImageSymbolConfiguration] class.
var (
	ImageSymbolConfigurationClass     _ImageSymbolConfigurationClass
	ImageSymbolConfigurationClassOnce sync.Once
)

func getImageSymbolConfigurationClass() _ImageSymbolConfigurationClass {
	ImageSymbolConfigurationClassOnce.Do(func() {
		ImageSymbolConfigurationClass = _ImageSymbolConfigurationClass{objc.GetClass("NSImageSymbolConfiguration")}
	})
	return ImageSymbolConfigurationClass
}

type _ImageSymbolConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ImageSymbolConfiguration */
// An interface definition for the [ImageSymbolConfiguration] class.
type IImageSymbolConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ImageSymbolConfiguration */
	// properties:
	SymbolConfiguration() IImageSymbolConfiguration
	SetSymbolConfiguration(value IImageSymbolConfiguration)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ImageSymbolConfiguration */
	// methods:
	ConfigurationByApplyingConfiguration(configuration IImageSymbolConfiguration) objectivec.IObject
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ImageSymbolConfiguration */
// Alloc allocates a new instance without initialization.
func (ic _ImageSymbolConfigurationClass) Alloc() ImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ImageSymbolConfigurationClass) New() ImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageSymbolConfiguration) Init() ImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageSymbolConfiguration) Autorelease() ImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageSymbolConfiguration creates a new ImageSymbolConfiguration instance.
func NewImageSymbolConfiguration() ImageSymbolConfiguration {
	return getImageSymbolConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ImageSymbolConfiguration */
// An object that contains the specific font, style, and weight attributes to apply to a symbol image.
//
// Symbol image configuration objects include details such as the point size, scale, text style, and weight to apply to your symbol image. The system uses these details to determine which variant of the image to use and how to scale or style the image. objects are immutable after you create them. If you use the method on the object, the new image attributes replace any previous attributes you supplied. After creating a symbol configuration object, assign it to the property of the object you use to display the image. If you draw the image directly, use the method to create a new image that contains the new attributes. For design guidance, see .


// An object that contains the specific font, style, and weight attributes to apply to a symbol image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class
type ImageSymbolConfiguration struct {
	objectivec.Object
}

// ImageSymbolConfigurationFrom constructs a [ImageSymbolConfiguration] from an unsafe.Pointer.
//
// An object that contains the specific font, style, and weight attributes to apply to a symbol image.
func ImageSymbolConfigurationFrom(ptr unsafe.Pointer) ImageSymbolConfiguration {
	return ImageSymbolConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ImageSymbolConfiguration */

// Create a configuration with a specific color rendering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(colorRenderingMode:)
func NewImageSymbolConfigurationWithColorRenderingMode(mode ImageSymbolColorRenderingMode) ImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](objc.ID(getImageSymbolConfigurationClass().class), objc.Sel("configurationWithColorRenderingMode:"), mode)
	return rv
}/* debug [class_init_methods/constructor]: NewImageSymbolConfigurationWithColorRenderingMode */


// Creates a hierarchical color configuration using the color you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(hierarchicalColor:)
func NewImageSymbolConfigurationWithHierarchicalColor(hierarchicalColor IColor) ImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](objc.ID(getImageSymbolConfigurationClass().class), objc.Sel("configurationWithHierarchicalColor:"), hierarchicalColor)
	return rv
}/* debug [class_init_methods/constructor]: NewImageSymbolConfigurationWithHierarchicalColor */


// Creates a color configuration by specifying a palette of colors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(paletteColors:)
func NewImageSymbolConfigurationWithPaletteColors(paletteColors []Color) ImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](objc.ID(getImageSymbolConfigurationClass().class), objc.Sel("configurationWithPaletteColors:"), paletteColors)
	return rv
}/* debug [class_init_methods/constructor]: NewImageSymbolConfigurationWithPaletteColors */


// Creates a symbol configuration with the specified point size and font weight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(pointSize:weight:)
func NewImageSymbolConfigurationWithPointSizeWeight(pointSize float64, weight FontWeight /* typedef */) ImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](objc.ID(getImageSymbolConfigurationClass().class), objc.Sel("configurationWithPointSize:weight:"), pointSize, weight)
	return rv
}/* debug [class_init_methods/constructor]: NewImageSymbolConfigurationWithPointSizeWeight */


// Creates a symbol configuration with the specified point size, font weight, and symbol scale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(pointSize:weight:scale:)
func NewImageSymbolConfigurationWithPointSizeWeightScale(pointSize float64, weight FontWeight /* typedef */, scale ImageSymbolScale) ImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](objc.ID(getImageSymbolConfigurationClass().class), objc.Sel("configurationWithPointSize:weight:scale:"), pointSize, weight, scale)
	return rv
}/* debug [class_init_methods/constructor]: NewImageSymbolConfigurationWithPointSizeWeightScale */


// Creates a symbol configuration using the scale you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(scale:)
func NewImageSymbolConfigurationWithScale(scale ImageSymbolScale) ImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](objc.ID(getImageSymbolConfigurationClass().class), objc.Sel("configurationWithScale:"), scale)
	return rv
}/* debug [class_init_methods/constructor]: NewImageSymbolConfigurationWithScale */


// Creates a symbol configuration with the specified text style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(textStyle:)
func NewImageSymbolConfigurationWithTextStyle(style FontTextStyle /* typedef */) ImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](objc.ID(getImageSymbolConfigurationClass().class), objc.Sel("configurationWithTextStyle:"), style)
	return rv
}/* debug [class_init_methods/constructor]: NewImageSymbolConfigurationWithTextStyle */


// Creates a symbol configuration with the specified text style and symbol scale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(textStyle:scale:)
func NewImageSymbolConfigurationWithTextStyleScale(style FontTextStyle /* typedef */, scale ImageSymbolScale) ImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](objc.ID(getImageSymbolConfigurationClass().class), objc.Sel("configurationWithTextStyle:scale:"), style, scale)
	return rv
}/* debug [class_init_methods/constructor]: NewImageSymbolConfigurationWithTextStyleScale */


// Create a configuration with a specified variable value mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(variableValueMode:)
func NewImageSymbolConfigurationWithVariableValueMode(variableValueMode ImageSymbolVariableValueMode) ImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](objc.ID(getImageSymbolConfigurationClass().class), objc.Sel("configurationWithVariableValueMode:"), variableValueMode)
	return rv
}/* debug [class_init_methods/constructor]: NewImageSymbolConfigurationWithVariableValueMode */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ImageSymbolConfiguration */

// Create a configuration with a specific color rendering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(colorRenderingMode:)
func (ic _ImageSymbolConfigurationClass) ConfigurationWithColorRenderingMode(mode ImageSymbolColorRenderingMode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("configurationWithColorRenderingMode:"), mode)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ConfigurationWithColorRenderingMode) */


// Creates a hierarchical color configuration using the color you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(hierarchicalColor:)
func (ic _ImageSymbolConfigurationClass) ConfigurationWithHierarchicalColor(hierarchicalColor IColor) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("configurationWithHierarchicalColor:"), hierarchicalColor)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ConfigurationWithHierarchicalColor) */


// Creates a color configuration by specifying a palette of colors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(paletteColors:)
func (ic _ImageSymbolConfigurationClass) ConfigurationWithPaletteColors(paletteColors []Color) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("configurationWithPaletteColors:"), paletteColors)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ConfigurationWithPaletteColors) */


// Creates a symbol configuration with the specified point size and font weight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(pointSize:weight:)
func (ic _ImageSymbolConfigurationClass) ConfigurationWithPointSizeWeight(pointSize float64, weight FontWeight /* typedef */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("configurationWithPointSize:weight:"), pointSize, weight)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ConfigurationWithPointSizeWeight) */


// Creates a symbol configuration with the specified point size, font weight, and symbol scale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(pointSize:weight:scale:)
func (ic _ImageSymbolConfigurationClass) ConfigurationWithPointSizeWeightScale(pointSize float64, weight FontWeight /* typedef */, scale ImageSymbolScale) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("configurationWithPointSize:weight:scale:"), pointSize, weight, scale)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ConfigurationWithPointSizeWeightScale) */


// Creates a symbol configuration using the scale you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(scale:)
func (ic _ImageSymbolConfigurationClass) ConfigurationWithScale(scale ImageSymbolScale) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("configurationWithScale:"), scale)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ConfigurationWithScale) */


// Creates a symbol configuration with the specified text style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(textStyle:)
func (ic _ImageSymbolConfigurationClass) ConfigurationWithTextStyle(style FontTextStyle /* typedef */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("configurationWithTextStyle:"), style)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ConfigurationWithTextStyle) */


// Creates a symbol configuration with the specified text style and symbol scale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(textStyle:scale:)
func (ic _ImageSymbolConfigurationClass) ConfigurationWithTextStyleScale(style FontTextStyle /* typedef */, scale ImageSymbolScale) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("configurationWithTextStyle:scale:"), style, scale)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ConfigurationWithTextStyleScale) */


// Create a configuration with a specified variable value mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(variableValueMode:)
func (ic _ImageSymbolConfigurationClass) ConfigurationWithVariableValueMode(variableValueMode ImageSymbolVariableValueMode) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("configurationWithVariableValueMode:"), variableValueMode)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ConfigurationWithVariableValueMode) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/preferringHierarchical()
func (ic _ImageSymbolConfigurationClass) ConfigurationPreferringHierarchical() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("configurationPreferringHierarchical"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ConfigurationPreferringHierarchical) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/preferringMonochrome()
func (ic _ImageSymbolConfigurationClass) ConfigurationPreferringMonochrome() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("configurationPreferringMonochrome"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ConfigurationPreferringMonochrome) */


// Creates a configuration that specifies that the symbol should prefer its multicolor variant if one exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/preferringMulticolor()
func (ic _ImageSymbolConfigurationClass) ConfigurationPreferringMulticolor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ic.class), objc.Sel("configurationPreferringMulticolor"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ConfigurationPreferringMulticolor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ImageSymbolConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ImageSymbolConfiguration */

// Creates a configuration object by applying the values from the configuration you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/applying(_:)
func (i_ ImageSymbolConfiguration) ConfigurationByApplyingConfiguration(configuration IImageSymbolConfiguration) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](i_.ID, objc.Sel("configurationByApplyingConfiguration:"), configuration)
	return rv
}/* debug [instance_methods/method]: ConfigurationByApplyingConfiguration */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ImageSymbolConfiguration */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/symbolconfiguration
func (i_ ImageSymbolConfiguration) SymbolConfiguration() IImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](i_.ID, objc.Sel("symbolConfiguration"))
	return rv
}/* debug [instance_properties/getter]: symbolConfiguration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/symbolconfiguration
func (i_ ImageSymbolConfiguration) SetSymbolConfiguration(value IImageSymbolConfiguration) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSymbolConfiguration:"), value)
}/* debug [instance_properties/setter]: symbolConfiguration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSImageSymbolConfiguration */


