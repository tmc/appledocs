// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AffineTransform] class.
var affineTransformClass = _AffineTransformClass{objc.GetClass("NSAffineTransform")}

type _AffineTransformClass struct {
	class objc.Class
}

// A graphics coordinate transformation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAffineTransform

type AffineTransform struct {
	objectivec.Object
}

// AffineTransformFrom constructs a [AffineTransform] from an unsafe.Pointer.
//
// A graphics coordinate transformation.
func AffineTransformFrom(ptr unsafe.Pointer) AffineTransform {
	return AffineTransform{objectivec.Object{objc.ID(ptr)}}
}



