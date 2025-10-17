// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Animation] class.
var animationClass = _AnimationClass{objc.GetClass("CAAnimation")}

type _AnimationClass struct {
	class objc.Class
}

// The abstract superclass for animations in Core Animation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation

type Animation struct {
	objectivec.Object
}

// AnimationFrom constructs a [Animation] from an unsafe.Pointer.
//
// The abstract superclass for animations in Core Animation.
func AnimationFrom(ptr unsafe.Pointer) Animation {
	return Animation{objectivec.Object{objc.ID(ptr)}}
}



