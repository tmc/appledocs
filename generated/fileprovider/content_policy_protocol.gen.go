// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import "github.com/ebitengine/purego/objc"

// contentPolicyProtocol is the contentPolicy protocol.
//
// Availability:
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to contentPolicy.
var contentPolicyProtocol *objc.Protocol

func init() {
	contentPolicyProtocol = objc.GetProtocol("contentPolicy")
}

