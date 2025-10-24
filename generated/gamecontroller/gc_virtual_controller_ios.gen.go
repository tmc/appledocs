//go:build darwin && ios

// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for GCVirtualController


// Connects the virtual controller to the device and displays it on the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCVirtualController/connect(replyHandler:)
func (g_ GCVirtualController) ConnectWithReplyHandler(reply unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("connectWithReplyHandler:"), reply)
}

// Disconnects the virtual controller from the device and removes it from the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCVirtualController/disconnect()
func (g_ GCVirtualController) Disconnect() {
	objc.Send[objc.ID](g_.ID, objc.Sel("disconnect"))
}

// Changes the value of a directional pad element in the virtual controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCVirtualController/setPosition(_:forDirectionPadElement:)
func (g_ GCVirtualController) SetPositionForDirectionPadElement(position corefoundation.CGPoint, element objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPosition:forDirectionPadElement:"), position, element)
}

// Changes the value of a button element in the virtual controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCVirtualController/setValue(_:forButtonElement:)
func (g_ GCVirtualController) SetValueForButtonElement(value float64, element objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setValue:forButtonElement:"), value, element)
}

// Changes the configuration for one of the virtual controller’s input elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCVirtualController/updateConfiguration(forElement:configuration:)
func (g_ GCVirtualController) UpdateConfigurationForElementConfiguration(element objc.IObject /* cross-framework: NSString */, config unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("updateConfigurationForElement:configuration:"), element, config)
}

// iOS-only properties

// The underlying controller object that you use to access input elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCVirtualController/controller
func (g_ GCVirtualController) Controller() IGCController {
	rv := objc.Send[GCController](g_.ID, objc.Sel("controller"))
	return rv
}




