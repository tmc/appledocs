// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import "github.com/ebitengine/purego/objc"

// timebaseProtocol is the timebase protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.13+
//   - tvOS 11.0+
//   - visionOS 1.0+
//   - watchOS 4.0+
//
// Use this protocol when registering custom classes that conform to timebase.
var timebaseProtocol *objc.Protocol

func init() {
	timebaseProtocol = objc.GetProtocol("timebase")
}

