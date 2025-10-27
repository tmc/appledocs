// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PFeatureProvider is the MLFeatureProvider protocol interface.
//
// An interface that represents a collection of values for either a model’s input or its output.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.13+
//   - tvOS 11.0+
//   - visionOS 1.0+
//   - watchOS 4.0+
//
// See: doc://com.apple.coreml/documentation/CoreML/MLFeatureProvider
type PFeatureProvider interface {
	// Required methods
	FeatureValueForName(featureName foundation.foundation.INSString) IFeatureValue
}
