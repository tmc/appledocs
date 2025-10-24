// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import "github.com/ebitengine/purego/objc"

// GameSessionEventListenerProtocol is the GKGameSessionEventListener protocol.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 10.0+ (Deprecated in 12.0)
//   - iPadOS 10.0+ (Deprecated in 12.0)
//   - macOS 10.12+ (Deprecated in 10.14)
//   - tvOS 10.0+ (Deprecated in 12.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// Use this protocol when registering custom classes that conform to GKGameSessionEventListener.
var GameSessionEventListenerProtocol *objc.Protocol

func init() {
	GameSessionEventListenerProtocol = objc.GetProtocol("GKGameSessionEventListener")
}

