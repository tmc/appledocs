// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PMDLAssetResolver is the MDLAssetResolver protocol interface.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.13+
//   - tvOS 11.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.modelio/documentation/ModelIO/MDLAssetResolver
type PMDLAssetResolver interface {
	// Required methods
	CanResolveAssetNamed(name objc.IObject /* cross-framework: NSString */) bool/* debug [protocol_interface/required_method]: CanResolveAssetNamed */
	ResolveAssetNamed(name objc.IObject /* cross-framework: NSString */) foundation.URL/* debug [protocol_interface/required_method]: ResolveAssetNamed */
}
