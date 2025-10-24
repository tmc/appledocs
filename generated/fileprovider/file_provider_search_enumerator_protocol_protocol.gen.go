// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"unsafe"
)

// PFileProviderSearchEnumerator is the NSFileProviderSearchEnumerator protocol interface.
//
// A protocol that defines methods for providing search results and canceling searches.
//
// Availability:
//   - macOS 26.0+
//
// See: doc://com.apple.fileprovider/documentation/FileProvider/NSFileProviderSearchEnumerator
type PFileProviderSearchEnumerator interface {
	// Required methods
	EnumerateSearchResultsForObserverStartingAtPage(observer unsafe.Pointer, page FileProviderPage /* typedef */)/* debug [protocol_interface/required_method]: EnumerateSearchResultsForObserverStartingAtPage */
	Invalidate()/* debug [protocol_interface/required_method]: Invalidate */
}
