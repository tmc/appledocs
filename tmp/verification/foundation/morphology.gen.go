// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var morphologyClass _MorphologyClass

func init() {
	morphologyClass = _MorphologyClass{objc.GetClass("NSMorphology")}
}

type _MorphologyClass struct {
	class objc.Class
}

type Morphology struct {
	objc.ID
}

func MorphologyFrom(ptr unsafe.Pointer) Morphology {
	return Morphology{
		ID: objc.ID(ptr),
	}
}




