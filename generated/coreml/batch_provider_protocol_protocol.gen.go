// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
)

// PBatchProvider is the MLBatchProvider protocol interface.
//
// An interface that represents a collection of feature providers.
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
// See: doc://com.apple.coreml/documentation/CoreML/MLBatchProvider
type PBatchProvider interface {
	// Required methods
	FeaturesAtIndex(index int) unsafe.Pointer/* debug [protocol_interface/required_method]: FeaturesAtIndex */
}
