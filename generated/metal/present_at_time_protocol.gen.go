// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// presentAtTimeProtocol is the presentAtTime: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.11+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to presentAtTime:.
var presentAtTimeProtocol *objc.Protocol

func init() {
	presentAtTimeProtocol = objc.GetProtocol("presentAtTime:")
}
