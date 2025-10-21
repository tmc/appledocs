// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// setBytesProtocol is the setBytes: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.3+
//   - iPadOS 8.3+
//   - macOS 10.11+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to setBytes:.
var setBytesProtocol *objc.Protocol

func init() {
	setBytesProtocol = objc.GetProtocol("setBytes:")
}
