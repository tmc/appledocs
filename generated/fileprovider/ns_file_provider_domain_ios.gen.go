//go:build darwin && ios

// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for FileProviderDomain


// iOS-only properties

// The path of the domain’s subdirectory relative to the file provider’s shared container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/pathRelativeToDocumentStorage
func (f_ FileProviderDomain) PathRelativeToDocumentStorage() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("pathRelativeToDocumentStorage"))
	return rv
}




