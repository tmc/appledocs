// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [AffineTransform] class.
var AffineTransformClass objc.Class

func init() {
	AffineTransformClass = objc.GetClass("NSAffineTransform")
}

type AffineTransform struct {
	objc.ID
}

func AffineTransformFrom(ptr unsafe.Pointer) AffineTransform {
	return AffineTransform{
		ID: objc.ID(ptr),
	}
}



