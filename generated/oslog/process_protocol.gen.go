// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import "github.com/ebitengine/purego/objc"

// processProtocol is the process protocol.
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 10.15+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+
//
// Use this protocol when registering custom classes that conform to process.
var processProtocol *objc.Protocol

func init() {
	processProtocol = objc.GetProtocol("process")
}
