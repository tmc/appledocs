// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import "github.com/ebitengine/purego/objc"

// isLossForPlayerProtocol is the isLossForPlayer: protocol.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to isLossForPlayer:.
var isLossForPlayerProtocol *objc.Protocol

func init() {
	isLossForPlayerProtocol = objc.GetProtocol("isLossForPlayer:")
}
