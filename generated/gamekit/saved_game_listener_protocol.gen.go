// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import "github.com/ebitengine/purego/objc"

// SavedGameListenerProtocol is the GKSavedGameListener protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to GKSavedGameListener.
var SavedGameListenerProtocol *objc.Protocol

func init() {
	SavedGameListenerProtocol = objc.GetProtocol("GKSavedGameListener")
}
