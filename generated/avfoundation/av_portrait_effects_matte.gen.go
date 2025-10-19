// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVPortraitEffectsMatte] class.
var aVPortraitEffectsMatteClass = _AVPortraitEffectsMatteClass{objc.GetClass("AVPortraitEffectsMatte")}

type _AVPortraitEffectsMatteClass struct {
	class objc.Class
}

// An auxiliary image used to separate foreground from background with high resolution. [Full Topic]
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



