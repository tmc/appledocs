// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MorphologyCustomPronoun] class.
var morphologyCustomPronounClass = _MorphologyCustomPronounClass{objc.GetClass("NSMorphologyCustomPronoun")}

type _MorphologyCustomPronounClass struct {
	class objc.Class
}

// A custom pronoun behavior for use in a specific langauge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMorphologyCustomPronoun

type MorphologyCustomPronoun struct {
	objectivec.Object
}

// MorphologyCustomPronounFrom constructs a [MorphologyCustomPronoun] from an unsafe.Pointer.
//
// A custom pronoun behavior for use in a specific langauge.
func MorphologyCustomPronounFrom(ptr unsafe.Pointer) MorphologyCustomPronoun {
	return MorphologyCustomPronoun{objectivec.Object{objc.ID(ptr)}}
}



