// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [topLeft] class.
var topLeftClass = _topLeftClass{objc.GetClass("topLeft")}

type _topLeftClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature/topLeft-c.ivar

type topLeft struct {
	objectivec.Object
}

// topLeftFrom constructs a [topLeft] from an unsafe.Pointer.
func topLeftFrom(ptr unsafe.Pointer) topLeft {
	return topLeft{objectivec.Object{objc.ID(ptr)}}
}



