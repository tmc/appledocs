// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import "github.com/ebitengine/purego/objc"

// connectInitiatorProtocol is the connectInitiator: protocol.
//
// Availability:
//   - Mac Catalyst 14.0+ (Deprecated in 18.0)
//   - iOS 14.0+ (Deprecated in 18.0)
//   - iPadOS 14.0+ (Deprecated in 18.0)
//   - macOS 11.0+ (Deprecated in 15.0)
//   - visionOS 1.0+ (Deprecated in 2.0)
//
// Use this protocol when registering custom classes that conform to connectInitiator:.
var connectInitiatorProtocol *objc.Protocol

func init() {
	connectInitiatorProtocol = objc.GetProtocol("connectInitiator:")
}

