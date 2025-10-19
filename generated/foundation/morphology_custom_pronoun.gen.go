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

// An interface definition for the [MorphologyCustomPronoun] class.
type IMorphologyCustomPronoun interface {
	objectivec.IObject
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
// Alloc allocates a new instance without initialization.
func (mc _MorphologyCustomPronounClass) Alloc() MorphologyCustomPronoun {
	rv := objc.Send[MorphologyCustomPronoun](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (mc _MorphologyCustomPronounClass) New() MorphologyCustomPronoun {
	rv := objc.Send[MorphologyCustomPronoun](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MorphologyCustomPronoun) Init() MorphologyCustomPronoun {
	rv := objc.Send[MorphologyCustomPronoun](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MorphologyCustomPronoun) Autorelease() MorphologyCustomPronoun {
	rv := objc.Send[MorphologyCustomPronoun](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMorphologyCustomPronoun creates a new MorphologyCustomPronoun instance.
func NewMorphologyCustomPronoun() MorphologyCustomPronoun {
	return morphologyCustomPronounClass.New()
}




