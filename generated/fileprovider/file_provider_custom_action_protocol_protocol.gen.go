// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PFileProviderCustomAction is the NSFileProviderCustomAction protocol interface.
//
// Support for custom actions.
//
// Availability:
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 11.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.fileprovider/documentation/FileProvider/NSFileProviderCustomAction
type PFileProviderCustomAction interface {
	// Required methods
	PerformActionWithIdentifierOnItemsWithIdentifiersCompletionHandler(actionIdentifier FileProviderExtensionActionIdentifier /* typedef */, itemIdentifiers []string, completionHandler unsafe.Pointer) foundation.Progress/* debug [protocol_interface/required_method]: PerformActionWithIdentifierOnItemsWithIdentifiersCompletionHandler */
}
