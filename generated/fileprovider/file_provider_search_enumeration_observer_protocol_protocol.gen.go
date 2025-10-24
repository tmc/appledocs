// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PFileProviderSearchEnumerationObserver is the NSFileProviderSearchEnumerationObserver protocol interface.
//
// A protocol that defines a type that receives enumerations of search results from your extension.
//
// Availability:
//   - macOS 26.0+
//
// See: doc://com.apple.fileprovider/documentation/FileProvider/NSFileProviderSearchEnumerationObserver
type PFileProviderSearchEnumerationObserver interface {
	// Required methods
	DidEnumerateSearchResults(searchResults []objc.ID)/* debug [protocol_interface/required_method]: DidEnumerateSearchResults */
	FinishEnumeratingUpToPage(nextPage FileProviderPage /* typedef */)/* debug [protocol_interface/required_method]: FinishEnumeratingUpToPage */
	FinishEnumeratingWithError(error_ objc.IObject /* cross-framework: Error */)/* debug [protocol_interface/required_method]: FinishEnumeratingWithError */
}
