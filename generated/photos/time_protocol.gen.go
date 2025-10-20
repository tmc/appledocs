// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import "github.com/ebitengine/purego/objc"

// timeProtocol is the time protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.12+
//   - tvOS 10.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to time.
var timeProtocol *objc.Protocol

func init() {
	timeProtocol = objc.GetProtocol("time")
}


