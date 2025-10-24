//go:build darwin && ios

// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
)

// iOS-only methods for PreviewReplyAttachment


// iOS-only properties

// The content type of the attachment for an html preview
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewReplyAttachment/contentType
func (p_ PreviewReplyAttachment) ContentType() uniformtypeidentifiers.UTType {
	rv := objc.Send[uniformtypeidentifiers.UTType](p_.ID, objc.Sel("contentType"))
	return rv
}

// The data content of an html preview
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewReplyAttachment/data
func (p_ PreviewReplyAttachment) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("data"))
	return rv
}




