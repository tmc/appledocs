//go:build darwin && ios

// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for GCVirtualControllerConfiguration


// iOS-only properties

// The input elements of a virtual controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCVirtualController/Configuration/elements
func (g_ GCVirtualControllerConfiguration) Elements() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("elements"))
	return rv
}
func (g_ GCVirtualControllerConfiguration) SetElements(value unsafe.Pointer) {
	g_.ID.Send(objc.RegisterName("setElements:"), value)
}

// A Boolean value that indicates whether the system or the app presents the virtual interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCVirtualController/Configuration/isHidden
func (g_ GCVirtualControllerConfiguration) Hidden() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("hidden"))
	return rv
}
func (g_ GCVirtualControllerConfiguration) SetHidden(value bool) {
	g_.ID.Send(objc.RegisterName("setHidden:"), value)
}





