// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import "github.com/ebitengine/purego/objc"

// OS_dispatch_queue_serialProtocol is the OS_dispatch_queue_serial protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+
//
// Use this protocol when registering custom classes that conform to OS_dispatch_queue_serial.
var OS_dispatch_queue_serialProtocol *objc.Protocol

func init() {
	OS_dispatch_queue_serialProtocol = objc.GetProtocol("OS_dispatch_queue_serial")
}

