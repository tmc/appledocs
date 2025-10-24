// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PFileProviderChangeObserver is the NSFileProviderChangeObserver protocol interface.
//
// An observer that receives changes and deletions during enumeration.
//
// Availability:
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 11.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.fileprovider/documentation/FileProvider/NSFileProviderChangeObserver
type PFileProviderChangeObserver interface {
	// Required methods
	DidDeleteItemsWithIdentifiers(deletedItemIdentifiers []string)/* debug [protocol_interface/required_method]: DidDeleteItemsWithIdentifiers */
	DidUpdateItems(updatedItems []objc.ID)/* debug [protocol_interface/required_method]: DidUpdateItems */
	FinishEnumeratingChangesUpToSyncAnchorMoreComing(anchor FileProviderSyncAnchor /* typedef */, moreComing bool)/* debug [protocol_interface/required_method]: FinishEnumeratingChangesUpToSyncAnchorMoreComing */
	FinishEnumeratingWithError(error_ objc.IObject /* cross-framework: Error */)/* debug [protocol_interface/required_method]: FinishEnumeratingWithError */
}
