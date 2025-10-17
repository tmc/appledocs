// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Morphology] class.
var morphologyClass = _MorphologyClass{objc.GetClass("NSMorphology")}

type _MorphologyClass struct {
	class objc.Class
}

// An interface definition for the [Morphology] class.
type IMorphology interface {
	objectivec.IObject
}

// A description of the grammatical properties of a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphology

type Morphology struct {
	objectivec.Object
}

// MorphologyFrom constructs a [Morphology] from an unsafe.Pointer.
//
// A description of the grammatical properties of a string.
func MorphologyFrom(ptr unsafe.Pointer) Morphology {
	return Morphology{objectivec.Object{objc.ID(ptr)}}
}



