// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [topRight] class.
var topRightClass = _topRightClass{objc.GetClass("topRight")}

type _topRightClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature/topRight-c.ivar

type topRight struct {
	objectivec.Object
}

// topRightFrom constructs a [topRight] from an unsafe.Pointer.
func topRightFrom(ptr unsafe.Pointer) topRight {
	return topRight{objectivec.Object{objc.ID(ptr)}}
}



