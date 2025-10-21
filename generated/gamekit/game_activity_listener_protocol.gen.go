// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import "github.com/ebitengine/purego/objc"

// GameActivityListenerProtocol is the GKGameActivityListener protocol.
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//
// Use this protocol when registering custom classes that conform to GKGameActivityListener.
var GameActivityListenerProtocol *objc.Protocol

func init() {
	GameActivityListenerProtocol = objc.GetProtocol("GKGameActivityListener")
}
