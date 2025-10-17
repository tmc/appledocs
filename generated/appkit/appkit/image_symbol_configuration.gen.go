// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ImageSymbolConfiguration] class.
var ImageSymbolConfigurationClass objc.Class

func init() {
	ImageSymbolConfigurationClass = objc.GetClass("NSImageSymbolConfiguration")
}

type ImageSymbolConfiguration struct {
	objc.ID
}

func ImageSymbolConfigurationFrom(ptr unsafe.Pointer) ImageSymbolConfiguration {
	return ImageSymbolConfiguration{
		ID: objc.ID(ptr),
	}
}




