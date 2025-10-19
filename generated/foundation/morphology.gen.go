// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Morphology] class.
var (
	morphologyClass     _MorphologyClass
	morphologyClassOnce sync.Once
)

func getMorphologyClass() _MorphologyClass {
	morphologyClassOnce.Do(func() {
		morphologyClass = _MorphologyClass{objc.GetClass("NSMorphology")}
	})
	return morphologyClass
}

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

// Alloc allocates a new instance without initialization.
func (mc _MorphologyClass) Alloc() Morphology {
	rv := objc.Send[Morphology](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MorphologyClass) New() Morphology {
	rv := objc.Send[Morphology](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ Morphology) Init() Morphology {
	rv := objc.Send[Morphology](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ Morphology) Autorelease() Morphology {
	rv := objc.Send[Morphology](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMorphology creates a new Morphology instance.
func NewMorphology() Morphology {
	return getMorphologyClass().New()
}




