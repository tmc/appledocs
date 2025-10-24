// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import "github.com/ebitengine/purego/objc"

// cacheProtocol is the cache: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// Use this protocol when registering custom classes that conform to cache:.
var cacheProtocol *objc.Protocol

func init() {
	cacheProtocol = objc.GetProtocol("cache:")
}
