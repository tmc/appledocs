// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PFileProviderEnumerationObserver is the NSFileProviderEnumerationObserver protocol interface.
//
// An observer that receives batches of items during enumeration.
//
// Availability:
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 11.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.fileprovider/documentation/FileProvider/NSFileProviderEnumerationObserver
type PFileProviderEnumerationObserver interface {
	// Required methods
	DidEnumerateItems(updatedItems []objc.ID)/* debug [protocol_interface/required_method]: DidEnumerateItems */
	FinishEnumeratingUpToPage(nextPage FileProviderPage /* typedef */)/* debug [protocol_interface/required_method]: FinishEnumeratingUpToPage */
	FinishEnumeratingWithError(error_ objc.IObject /* cross-framework: Error */)/* debug [protocol_interface/required_method]: FinishEnumeratingWithError */
}
