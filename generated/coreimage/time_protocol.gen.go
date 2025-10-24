// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import "github.com/ebitengine/purego/objc"

// timeProtocol is the time protocol.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//
// Use this protocol when registering custom classes that conform to time.
var timeProtocol *objc.Protocol

func init() {
	timeProtocol = objc.GetProtocol("time")
}

