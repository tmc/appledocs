// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [bottomRight] class.
var bottomRightClass = _bottomRightClass{objc.GetClass("bottomRight")}

type _bottomRightClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature/bottomRight-c.ivar

type bottomRight struct {
	objectivec.Object
}

// bottomRightFrom constructs a [bottomRight] from an unsafe.Pointer.
func bottomRightFrom(ptr unsafe.Pointer) bottomRight {
	return bottomRight{objectivec.Object{objc.ID(ptr)}}
}



