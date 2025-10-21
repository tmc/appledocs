// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import "github.com/ebitengine/purego/objc"

// playerIdProtocol is the playerId protocol.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to playerId.
var playerIdProtocol *objc.Protocol

func init() {
	playerIdProtocol = objc.GetProtocol("playerId")
}
