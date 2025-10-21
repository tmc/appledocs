// Code generated from Apple documentation for PassKit. DO NOT EDIT.

package passkit

import "github.com/ebitengine/purego/objc"

// completionHandlerProtocol is the completionHandler protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to completionHandler.
var completionHandlerProtocol *objc.Protocol

func init() {
	completionHandlerProtocol = objc.GetProtocol("completionHandler")
}
