// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import "github.com/ebitengine/purego/objc"

// agentWillUpdateProtocol is the agentWillUpdate: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to agentWillUpdate:.
var agentWillUpdateProtocol *objc.Protocol

func init() {
	agentWillUpdateProtocol = objc.GetProtocol("agentWillUpdate:")
}
