// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import "github.com/ebitengine/purego/objc"

// activePlayerProtocol is the activePlayer protocol.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to activePlayer.
var activePlayerProtocol *objc.Protocol

func init() {
	activePlayerProtocol = objc.GetProtocol("activePlayer")
}
