// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"unsafe"
)

// PFileProviderEnumerator is the NSFileProviderEnumerator protocol interface.
//
// A protocol for enumerating items and changes.
//
// Availability:
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 11.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.fileprovider/documentation/FileProvider/NSFileProviderEnumerator
type PFileProviderEnumerator interface {
	// Required methods
	EnumerateItemsForObserverStartingAtPage(observer unsafe.Pointer, page FileProviderPage /* typedef */)/* debug [protocol_interface/required_method]: EnumerateItemsForObserverStartingAtPage */
	Invalidate()/* debug [protocol_interface/required_method]: Invalidate */
	// Optional methods
	CurrentSyncAnchorWithCompletionHandler(completionHandler unsafe.Pointer)
	HasCurrentSyncAnchorWithCompletionHandler() bool
	EnumerateChangesForObserverFromSyncAnchor(observer unsafe.Pointer, syncAnchor FileProviderSyncAnchor /* typedef */)
	HasEnumerateChangesForObserverFromSyncAnchor() bool
}
