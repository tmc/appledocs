// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import "github.com/ebitengine/purego/objc"

// URLSessionTaskDelegateProtocol is the NSURLSessionTaskDelegate protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// Use this protocol when registering custom classes that conform to NSURLSessionTaskDelegate.
var URLSessionTaskDelegateProtocol *objc.Protocol

func init() {
	URLSessionTaskDelegateProtocol = objc.GetProtocol("NSURLSessionTaskDelegate")
}
