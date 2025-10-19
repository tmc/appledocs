// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [maskPattern] class.
var maskPatternClass = _maskPatternClass{objc.GetClass("maskPattern")}

type _maskPatternClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/maskPattern-c.ivar

type maskPattern struct {
	objectivec.Object
}

// maskPatternFrom constructs a [maskPattern] from an unsafe.Pointer.
func maskPatternFrom(ptr unsafe.Pointer) maskPattern {
	return maskPattern{objectivec.Object{objc.ID(ptr)}}
}



