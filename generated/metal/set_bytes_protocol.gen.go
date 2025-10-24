// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// setBytesProtocol is the setBytes: protocol.
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to setBytes:.
var setBytesProtocol *objc.Protocol

func init() {
	setBytesProtocol = objc.GetProtocol("setBytes:")
}

