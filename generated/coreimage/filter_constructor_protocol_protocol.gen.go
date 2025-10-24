// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PFilterConstructor is the CIFilterConstructor protocol interface.
//
// A general interface for objects that produce filters.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//
// See: doc://com.apple.coreimage/documentation/CoreImage/CIFilterConstructor
type PFilterConstructor interface {
	// Required methods
	FilterWithName(name objc.IObject /* cross-framework: NSString */) Filter/* debug [protocol_interface/required_method]: FilterWithName */
}
