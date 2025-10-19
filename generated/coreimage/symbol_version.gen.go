// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [symbolVersion] class.
var symbolVersionClass = _symbolVersionClass{objc.GetClass("symbolVersion")}

type _symbolVersionClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/symbolVersion-c.ivar

type symbolVersion struct {
	objectivec.Object
}

// symbolVersionFrom constructs a [symbolVersion] from an unsafe.Pointer.
func symbolVersionFrom(ptr unsafe.Pointer) symbolVersion {
	return symbolVersion{objectivec.Object{objc.ID(ptr)}}
}



