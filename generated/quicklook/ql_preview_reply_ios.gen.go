//go:build darwin && ios

// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
)

// iOS-only methods for PreviewReply


// iOS-only properties

// Attachments for HTML data previews. The keys of the dictionary are the attachment identifiers (eg foo) that can be referenced with the cid:id URL (eg cid:foo).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewReply/attachments
func (p_ PreviewReply) Attachments() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("attachments"))
	return rv
}
func (p_ PreviewReply) SetAttachments(value foundation.IDictionary) {
	p_.ID.Send(objc.RegisterName("setAttachments:"), value)
}

// String encoding for text or html based previews. Defaults to NSUTF8StringEncoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewReply/stringEncoding-1k9kb
func (p_ PreviewReply) StringEncoding() StringEncoding /* not a class type */ {
	rv := objc.Send[StringEncoding](p_.ID, objc.Sel("stringEncoding"))
	return rv
}
func (p_ PreviewReply) SetStringEncoding(value StringEncoding /* not a class type */) {
	p_.ID.Send(objc.RegisterName("setStringEncoding:"), value)
}

// Custom display title for the preview. If left as the empty string, QuickLook will use the file name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewReply/title
func (p_ PreviewReply) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("title"))
	return rv
}
func (p_ PreviewReply) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	p_.ID.Send(objc.RegisterName("setTitle:"), value)
}




