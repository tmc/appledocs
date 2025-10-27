// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PWritable is the MLWritable protocol interface.
//
// A set of methods that saves a machine learning type to the file system.
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
// See: doc://com.apple.coreml/documentation/CoreML/MLWritable
type PWritable interface {
	// Required methods
	WriteToURLError(url foundation.foundation.INSURL, error_ foundation.foundation.INSError) bool
}
