// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import "github.com/ebitengine/purego/objc"

// headphoneMotionManagerDidDisconnectProtocol is the headphoneMotionManagerDidDisconnect: protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 14.0+
//   - watchOS 7.0+
//
// Use this protocol when registering custom classes that conform to headphoneMotionManagerDidDisconnect:.
var headphoneMotionManagerDidDisconnectProtocol *objc.Protocol

func init() {
	headphoneMotionManagerDidDisconnectProtocol = objc.GetProtocol("headphoneMotionManagerDidDisconnect:")
}
