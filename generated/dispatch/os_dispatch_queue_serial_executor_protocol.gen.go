// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import "github.com/ebitengine/purego/objc"

// OS_dispatch_queue_serial_executorProtocol is the OS_dispatch_queue_serial_executor protocol.
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+
//
// Use this protocol when registering custom classes that conform to OS_dispatch_queue_serial_executor.
var OS_dispatch_queue_serial_executorProtocol *objc.Protocol

func init() {
	OS_dispatch_queue_serial_executorProtocol = objc.GetProtocol("OS_dispatch_queue_serial_executor")
}
