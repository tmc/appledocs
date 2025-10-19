// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVPortraitEffectsMatte] class.
var (
	aVPortraitEffectsMatteClass     _AVPortraitEffectsMatteClass
	aVPortraitEffectsMatteClassOnce sync.Once
)

func getAVPortraitEffectsMatteClass() _AVPortraitEffectsMatteClass {
	aVPortraitEffectsMatteClassOnce.Do(func() {
		aVPortraitEffectsMatteClass = _AVPortraitEffectsMatteClass{objc.GetClass("AVPortraitEffectsMatte")}
	})
	return aVPortraitEffectsMatteClass
}

type _AVPortraitEffectsMatteClass struct {
	class objc.Class
}

// An interface definition for the [AVPortraitEffectsMatte] class.
type IAVPortraitEffectsMatte interface {
	objectivec.IObject
}

// An auxiliary image used to separate foreground from background with high resolution.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPortraitEffectsMatte
type AVPortraitEffectsMatte struct {
	objectivec.Object
}

// AVPortraitEffectsMatteFrom constructs a [AVPortraitEffectsMatte] from an unsafe.Pointer.
//
// An auxiliary image used to separate foreground from background with high resolution.
func AVPortraitEffectsMatteFrom(ptr unsafe.Pointer) AVPortraitEffectsMatte {
	return AVPortraitEffectsMatte{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AVPortraitEffectsMatteClass) Alloc() AVPortraitEffectsMatte {
	rv := objc.Send[AVPortraitEffectsMatte](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVPortraitEffectsMatteClass) New() AVPortraitEffectsMatte {
	rv := objc.Send[AVPortraitEffectsMatte](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVPortraitEffectsMatte) Init() AVPortraitEffectsMatte {
	rv := objc.Send[AVPortraitEffectsMatte](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVPortraitEffectsMatte) Autorelease() AVPortraitEffectsMatte {
	rv := objc.Send[AVPortraitEffectsMatte](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPortraitEffectsMatte creates a new AVPortraitEffectsMatte instance.
func NewAVPortraitEffectsMatte() AVPortraitEffectsMatte {
	return getAVPortraitEffectsMatteClass().New()
}




