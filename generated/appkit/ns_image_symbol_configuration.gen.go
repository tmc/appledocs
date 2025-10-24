// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [ImageSymbolConfiguration] class.
type IImageSymbolConfiguration interface {
	objectivec.IObject
	// properties:
	SymbolConfiguration() IImageSymbolConfiguration
	SetSymbolConfiguration(value IImageSymbolConfiguration)
	// methods:
	ConfigurationByApplyingConfiguration(configuration IImageSymbolConfiguration) unsafe.Pointer
}

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

// Alloc allocates a new instance without initialization.
func (ic _ImageSymbolConfigurationClass) Alloc() ImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates a symbol configuration with the specified point size and font weight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(pointSize:weight:)
func NewImageSymbolConfigurationWithPointSizeWeight(pointSize float64, weight FontWeight /* not a class type */) ImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](objc.ID(getImageSymbolConfigurationClass().class), objc.Sel("configurationWithPointSize:weight:"), pointSize, weight)
	return rv
}



// Creates a symbol configuration with the specified point size and font weight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(pointSize:weight:)
func (ic _ImageSymbolConfigurationClass) ConfigurationWithPointSizeWeight(pointSize float64, weight FontWeight /* not a class type */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("configurationWithPointSize:weight:"), pointSize, weight)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/preferringHierarchical()
func (ic _ImageSymbolConfigurationClass) ConfigurationPreferringHierarchical() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("configurationPreferringHierarchical"))
	return rv
}


// Creates a configuration that specifies that the symbol should prefer its multicolor variant if one exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/preferringMulticolor()
func (ic _ImageSymbolConfigurationClass) ConfigurationPreferringMulticolor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("configurationPreferringMulticolor"))
	return rv
}


// Creates a configuration object by applying the values from the configuration you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/applying(_:)
func (i_ ImageSymbolConfiguration) ConfigurationByApplyingConfiguration(configuration IImageSymbolConfiguration) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("configurationByApplyingConfiguration:"), configuration)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/symbolconfiguration
func (i_ ImageSymbolConfiguration) SymbolConfiguration() IImageSymbolConfiguration {
	rv := objc.Send[ImageSymbolConfiguration](i_.ID, objc.Sel("symbolConfiguration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimageview/symbolconfiguration
func (i_ ImageSymbolConfiguration) SetSymbolConfiguration(value IImageSymbolConfiguration) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSymbolConfiguration:"), value)
}


