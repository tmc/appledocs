// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [bottomLeft] class.
var bottomLeftClass = _bottomLeftClass{objc.GetClass("bottomLeft")}

type _bottomLeftClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature/bottomLeft-c.ivar

type bottomLeft struct {
	objectivec.Object
}

// bottomLeftFrom constructs a [bottomLeft] from an unsafe.Pointer.
func bottomLeftFrom(ptr unsafe.Pointer) bottomLeft {
	return bottomLeft{objectivec.Object{objc.ID(ptr)}}
}



