// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// supportsFunctionPointersProtocol is the supportsFunctionPointers protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to supportsFunctionPointers.
var supportsFunctionPointersProtocol *objc.Protocol

func init() {
	supportsFunctionPointersProtocol = objc.GetProtocol("supportsFunctionPointers")
}
