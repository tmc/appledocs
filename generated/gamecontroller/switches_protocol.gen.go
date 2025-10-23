// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import "github.com/ebitengine/purego/objc"

// switchesProtocol is the switches protocol.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to switches.
var switchesProtocol *objc.Protocol

func init() {
	switchesProtocol = objc.GetProtocol("switches")
}
