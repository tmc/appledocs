// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [SFAcousticFeature] class.
var (
	SFAcousticFeatureClass     _SFAcousticFeatureClass
	SFAcousticFeatureClassOnce sync.Once
)

func getSFAcousticFeatureClass() _SFAcousticFeatureClass {
	SFAcousticFeatureClassOnce.Do(func() {
		SFAcousticFeatureClass = _SFAcousticFeatureClass{objc.GetClass("SFAcousticFeature")}
	})
	return SFAcousticFeatureClass
}

type _SFAcousticFeatureClass struct {
	class objc.Class
}

// An interface definition for the [SFAcousticFeature] class.
type ISFAcousticFeature interface {
	objectivec.IObject
}

// The value of a voice analysis metric.
//
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFAcousticFeature
type SFAcousticFeature struct {
	objectivec.Object
}

// SFAcousticFeatureFrom constructs a [SFAcousticFeature] from an unsafe.Pointer.
//
// The value of a voice analysis metric.
func SFAcousticFeatureFrom(ptr unsafe.Pointer) SFAcousticFeature {
	return SFAcousticFeature{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SFAcousticFeatureClass) Alloc() SFAcousticFeature {
	rv := objc.Send[SFAcousticFeature](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFAcousticFeatureClass) New() SFAcousticFeature {
	rv := objc.Send[SFAcousticFeature](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFAcousticFeature) Init() SFAcousticFeature {
	rv := objc.Send[SFAcousticFeature](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFAcousticFeature) Autorelease() SFAcousticFeature {
	rv := objc.Send[SFAcousticFeature](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFAcousticFeature creates a new SFAcousticFeature instance.
func NewSFAcousticFeature() SFAcousticFeature {
	return getSFAcousticFeatureClass().New()
}




