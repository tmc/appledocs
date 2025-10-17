// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SpringAnimation] class.
var springAnimationClass = _SpringAnimationClass{objc.GetClass("CASpringAnimation")}

type _SpringAnimationClass struct {
	class objc.Class
}

// An animation that applies a spring-like force to a layer’s properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation

type SpringAnimation struct {
	BasicAnimation
}

// SpringAnimationFrom constructs a [SpringAnimation] from an unsafe.Pointer.
//
// An animation that applies a spring-like force to a layer’s properties.
func SpringAnimationFrom(ptr unsafe.Pointer) SpringAnimation {
	return SpringAnimation{
		BasicAnimation: BasicAnimationFrom(ptr),
	}
}



