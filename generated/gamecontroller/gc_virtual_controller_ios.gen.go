//go:build darwin && ios

// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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

// iOS-only properties




