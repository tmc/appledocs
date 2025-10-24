// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import "github.com/ebitengine/purego/objc"

// LocalPlayerListenerProtocol is the GKLocalPlayerListener protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 3.0+
//
// Use this protocol when registering custom classes that conform to GKLocalPlayerListener.
var LocalPlayerListenerProtocol *objc.Protocol

func init() {
	LocalPlayerListenerProtocol = objc.GetProtocol("GKLocalPlayerListener")
}

