// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import "github.com/ebitengine/purego/objc"

// OS_nw_interfaceProtocol is the OS_nw_interface protocol.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 6.0+
//
// Use this protocol when registering custom classes that conform to OS_nw_interface.
var OS_nw_interfaceProtocol *objc.Protocol

func init() {
	OS_nw_interfaceProtocol = objc.GetProtocol("OS_nw_interface")
}
