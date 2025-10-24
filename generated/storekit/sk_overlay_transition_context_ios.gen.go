//go:build darwin && ios

// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for OverlayTransitionContext


// Adds a closure you can use to animate view properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/TransitionContext/addAnimation(_:)
func (o_ OverlayTransitionContext) AddAnimationBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("addAnimationBlock:"), block)
}

// iOS-only properties

// The size and location of the overlay at the end of the transition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/TransitionContext/endFrame
func (o_ OverlayTransitionContext) EndFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o_.ID, objc.Sel("endFrame"))
	return rv
}

// The size and location of the overlay before the transition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/TransitionContext/startFrame
func (o_ OverlayTransitionContext) StartFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o_.ID, objc.Sel("startFrame"))
	return rv
}





