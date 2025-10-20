// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var AffineTransformClass _AffineTransformClass

func init() {
	AffineTransformClass = _AffineTransformClass{objc.GetClass("NSAffineTransform")}
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




