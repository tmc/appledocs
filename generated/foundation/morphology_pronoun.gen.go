// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MorphologyPronoun] class.
var (
	morphologyPronounClass     _MorphologyPronounClass
	morphologyPronounClassOnce sync.Once
)

func getMorphologyPronounClass() _MorphologyPronounClass {
	morphologyPronounClassOnce.Do(func() {
		morphologyPronounClass = _MorphologyPronounClass{objc.GetClass("NSMorphologyPronoun")}
	})
	return morphologyPronounClass
}

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

// Alloc allocates a new instance without initialization.
func (mc _MorphologyPronounClass) Alloc() MorphologyPronoun {
	rv := objc.Send[MorphologyPronoun](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MorphologyPronounClass) New() MorphologyPronoun {
	rv := objc.Send[MorphologyPronoun](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MorphologyPronoun) Init() MorphologyPronoun {
	rv := objc.Send[MorphologyPronoun](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MorphologyPronoun) Autorelease() MorphologyPronoun {
	rv := objc.Send[MorphologyPronoun](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMorphologyPronoun creates a new MorphologyPronoun instance.
func NewMorphologyPronoun() MorphologyPronoun {
	return getMorphologyPronounClass().New()
}




