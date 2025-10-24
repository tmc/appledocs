//go:build darwin && ios

// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for PreviewController


// Asks the preview controller to reload its data from its data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewController/reloadData()
func (p_ PreviewController) ReloadData() {
	objc.Send[objc.ID](p_.ID, objc.Sel("reloadData"))
}

// iOS-only properties





