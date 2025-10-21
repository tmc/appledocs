// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import "github.com/ebitengine/purego/objc"

// BatchProviderProtocol is the MLBatchProvider protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+
//
// Use this protocol when registering custom classes that conform to MLBatchProvider.
var BatchProviderProtocol *objc.Protocol

func init() {
	BatchProviderProtocol = objc.GetProtocol("MLBatchProvider")
}
