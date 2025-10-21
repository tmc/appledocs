// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

import "github.com/ebitengine/purego/objc"

// removeObjectProtocol is the removeObject: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to removeObject:.
var removeObjectProtocol *objc.Protocol

func init() {
	removeObjectProtocol = objc.GetProtocol("removeObject:")
}
