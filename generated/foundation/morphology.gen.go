// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Morphology] class.
var MorphologyClass objc.Class

func init() {
	MorphologyClass = objc.GetClass("NSMorphology")
}

type Morphology struct {
	objc.ID
}

func MorphologyFrom(ptr unsafe.Pointer) Morphology {
	return Morphology{
		ID: objc.ID(ptr),
	}
}



