// Code generated from Apple documentation for CoreText. DO NOT EDIT.

package coretext

import (

	"github.com/tmc/appledocs/generated/corefoundation"
)

// PAdaptiveImageProviding is the CTAdaptiveImageProviding protocol interface.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//   - watchOS +
//
// See: doc://com.apple.coretext/documentation/CoreText/CTAdaptiveImageProviding
type PAdaptiveImageProviding interface {
	// Required methods
	ImageForProposedSizeScaleFactorImageOffsetImageSize(proposedSize corefoundation.CGSize, scaleFactor float64, outImageOffset corefoundation.CGPoint, outImageSize corefoundation.CGSize) ImageRef/* debug [protocol_interface/required_method]: ImageForProposedSizeScaleFactorImageOffsetImageSize */
}
