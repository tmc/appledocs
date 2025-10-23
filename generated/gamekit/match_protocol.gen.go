// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import "github.com/ebitengine/purego/objc"

// matchProtocol is the match: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to match:.
var matchProtocol *objc.Protocol

func init() {
	matchProtocol = objc.GetProtocol("match:")
}
