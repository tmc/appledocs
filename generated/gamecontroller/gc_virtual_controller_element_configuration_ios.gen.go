//go:build darwin && ios

// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for GCVirtualControllerElementConfiguration


// iOS-only properties

// A Boolean value that determines whether the thumbstick element behaves as a touchpad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCVirtualController/ElementConfiguration/actsAsTouchpad
func (g_ GCVirtualControllerElementConfiguration) ActsAsTouchpad() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("actsAsTouchpad"))
	return rv
}
func (g_ GCVirtualControllerElementConfiguration) SetActsAsTouchpad(value bool) {
	g_.ID.Send(objc.RegisterName("setActsAsTouchpad:"), value)
}

// A Boolean value that determines whether the virtual controller hides the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCVirtualController/ElementConfiguration/isHidden
func (g_ GCVirtualControllerElementConfiguration) Hidden() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("hidden"))
	return rv
}
func (g_ GCVirtualControllerElementConfiguration) SetHidden(value bool) {
	g_.ID.Send(objc.RegisterName("setHidden:"), value)
}

// The Bezier path for the shape of an element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCVirtualController/ElementConfiguration/path
func (g_ GCVirtualControllerElementConfiguration) Path() appkit.BezierPath {
	rv := objc.Send[appkit.BezierPath](g_.ID, objc.Sel("path"))
	return rv
}
func (g_ GCVirtualControllerElementConfiguration) SetPath(value appkit.BezierPath) {
	g_.ID.Send(objc.RegisterName("setPath:"), value)
}





