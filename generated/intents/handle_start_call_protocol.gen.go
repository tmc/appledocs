// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import "github.com/ebitengine/purego/objc"

// handleStartCallProtocol is the handleStartCall: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 12.0+
//   - visionOS 1.0+
//   - watchOS 6.0+
//
// Use this protocol when registering custom classes that conform to handleStartCall:.
var handleStartCallProtocol *objc.Protocol

func init() {
	handleStartCallProtocol = objc.GetProtocol("handleStartCall:")
}
