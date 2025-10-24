//go:build darwin && ios

// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for FilePreviewRequest


// iOS-only properties

// The url of the file for which a preview is being requested.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLFilePreviewRequest/fileURL
func (f_ FilePreviewRequest) FileURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](f_.ID, objc.Sel("fileURL"))
	return rv
}





