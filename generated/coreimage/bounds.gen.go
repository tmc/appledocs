// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [bounds] class.
var boundsClass = _boundsClass{objc.GetClass("bounds")}

type _boundsClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature/bounds-c.ivar

type bounds struct {
	objectivec.Object
}

// boundsFrom constructs a [bounds] from an unsafe.Pointer.
func boundsFrom(ptr unsafe.Pointer) bounds {
	return bounds{objectivec.Object{objc.ID(ptr)}}
}



