// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import "github.com/ebitengine/purego/objc"

// initWithDomainProtocol is the initWithDomain: protocol.
//
// Availability:
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 11.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to initWithDomain:.
var initWithDomainProtocol *objc.Protocol

func init() {
	initWithDomainProtocol = objc.GetProtocol("initWithDomain:")
}
