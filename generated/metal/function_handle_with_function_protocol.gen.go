// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// functionHandleWithFunctionProtocol is the functionHandleWithFunction: protocol.
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to functionHandleWithFunction:.
var functionHandleWithFunctionProtocol *objc.Protocol

func init() {
	functionHandleWithFunctionProtocol = objc.GetProtocol("functionHandleWithFunction:")
}
