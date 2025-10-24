// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PFileProviderUserInteractionSuppressing is the NSFileProviderUserInteractionSuppressing protocol interface.
//
// Support for suppressing user-interaction alerts.
//
// Availability:
//   - macOS 12.0+
//
// See: doc://com.apple.fileprovider/documentation/FileProvider/NSFileProviderUserInteractionSuppressing
type PFileProviderUserInteractionSuppressing interface {
	// Required methods
	IsInteractionSuppressedForIdentifier(suppressionIdentifier objc.IObject /* cross-framework: NSString */) bool/* debug [protocol_interface/required_method]: IsInteractionSuppressedForIdentifier */
	SetInteractionSuppressedForIdentifier(suppression bool, suppressionIdentifier objc.IObject /* cross-framework: NSString */)/* debug [protocol_interface/required_method]: SetInteractionSuppressedForIdentifier */
}
