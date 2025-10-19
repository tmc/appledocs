// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [hasSmile] class.
var hasSmileClass = _hasSmileClass{objc.GetClass("hasSmile")}

type _hasSmileClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasSmile-c.ivar

type hasSmile struct {
	objectivec.Object
}

// hasSmileFrom constructs a [hasSmile] from an unsafe.Pointer.
func hasSmileFrom(ptr unsafe.Pointer) hasSmile {
	return hasSmile{objectivec.Object{objc.ID(ptr)}}
}



