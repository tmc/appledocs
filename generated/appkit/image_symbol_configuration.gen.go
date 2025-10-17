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



