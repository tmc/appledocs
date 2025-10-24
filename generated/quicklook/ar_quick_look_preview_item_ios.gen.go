//go:build darwin && ios

// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for QuickLookPreviewItem


// iOS-only properties

// Whether or not AR Quick Look allows content scaling in AR mode. Defaults to which allows scaling content in AR mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/ARQuickLookPreviewItem/allowsContentScaling
func (q_ QuickLookPreviewItem) AllowsContentScaling() bool {
	rv := objc.Send[bool](q_.ID, objc.Sel("allowsContentScaling"))
	return rv
}
func (q_ QuickLookPreviewItem) SetAllowsContentScaling(value bool) {
	q_.ID.Send(objc.RegisterName("setAllowsContentScaling:"), value)
}

// An optional canonical web page URL for the 3D content that will be shared.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/ARQuickLookPreviewItem/canonicalWebPageURL
func (q_ QuickLookPreviewItem) CanonicalWebPageURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](q_.ID, objc.Sel("canonicalWebPageURL"))
	return rv
}
func (q_ QuickLookPreviewItem) SetCanonicalWebPageURL(value objc.IObject /* cross-framework: NSURL */) {
	q_.ID.Send(objc.RegisterName("setCanonicalWebPageURL:"), value)
}




