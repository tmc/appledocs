// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import "github.com/ebitengine/purego/objc"

// serviceNameProtocol is the serviceName protocol.
//
// Availability:
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 11.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to serviceName.
var serviceNameProtocol *objc.Protocol

func init() {
	serviceNameProtocol = objc.GetProtocol("serviceName")
}

