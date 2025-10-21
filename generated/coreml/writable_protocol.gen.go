// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import "github.com/ebitengine/purego/objc"

// WritableProtocol is the MLWritable protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 6.0+
//
// Use this protocol when registering custom classes that conform to MLWritable.
var WritableProtocol *objc.Protocol

func init() {
	WritableProtocol = objc.GetProtocol("MLWritable")
}
