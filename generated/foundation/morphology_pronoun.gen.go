// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MorphologyPronoun] class.
var morphologyPronounClass = _MorphologyPronounClass{objc.GetClass("NSMorphologyPronoun")}

type _MorphologyPronounClass struct {
	class objc.Class
}

// An interface definition for the [MorphologyPronoun] class.
type IMorphologyPronoun interface {
	objectivec.IObject
}

// A custom pronoun for referring to a third person. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphologyPronoun

type MorphologyPronoun struct {
	objectivec.Object
}

// MorphologyPronounFrom constructs a [MorphologyPronoun] from an unsafe.Pointer.
//
// A custom pronoun for referring to a third person.
func MorphologyPronounFrom(ptr unsafe.Pointer) MorphologyPronoun {
	return MorphologyPronoun{objectivec.Object{objc.ID(ptr)}}
}



