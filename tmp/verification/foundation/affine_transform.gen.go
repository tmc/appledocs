// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var affineTransformClass _AffineTransformClass

func init() {
	affineTransformClass = _AffineTransformClass{objc.GetClass("NSAffineTransform")}
}

type _AffineTransformClass struct {
	class objc.Class
}

type AffineTransform struct {
	objc.ID
}

func AffineTransformFrom(ptr unsafe.Pointer) AffineTransform {
	return AffineTransform{
		ID: objc.ID(ptr),
	}
}




