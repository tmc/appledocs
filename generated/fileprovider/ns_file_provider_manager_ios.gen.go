//go:build darwin && ios

// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for FileProviderManager


// iOS-only properties

// The root URL for all shared documents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/documentStorageURL
func (f_ FileProviderManager) DocumentStorageURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](f_.ID, objc.Sel("documentStorageURL"))
	return rv
}

// A purpose identifier for coordinated reads and writes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/providerIdentifier
func (f_ FileProviderManager) ProviderIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("providerIdentifier"))
	return rv
}




