// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ImageSymbolConfiguration] class.
var imageSymbolConfigurationClass = _ImageSymbolConfigurationClass{objc.GetClass("NSImageSymbolConfiguration")}

type _ImageSymbolConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [ImageSymbolConfiguration] class.
type IImageSymbolConfiguration interface {
	objectivec.IObject
}

// An object that contains the specific font, style, and weight attributes to apply to a symbol image. [Full Topic]
//
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

// New creates and returns a new instance with a +1 retain count.
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
	return imageSymbolConfigurationClass.New()
}




