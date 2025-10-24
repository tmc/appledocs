//go:build darwin && ios

// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for PDFThumbnailView


// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/contentInset
func (p_ PDFThumbnailView) ContentInset() objc.IObject /* cross-framework: EdgeInsets */ {
	rv := objc.Send[foundation.EdgeInsets](p_.ID, objc.Sel("contentInset"))
	return rv
}
func (p_ PDFThumbnailView) SetContentInset(value objc.IObject /* cross-framework: EdgeInsets */) {
	p_.ID.Send(objc.RegisterName("setContentInset:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/layoutMode
func (p_ PDFThumbnailView) LayoutMode() PDFThumbnailLayoutMode {
	rv := objc.Send[PDFThumbnailLayoutMode](p_.ID, objc.Sel("layoutMode"))
	return rv
}
func (p_ PDFThumbnailView) SetLayoutMode(value PDFThumbnailLayoutMode) {
	p_.ID.Send(objc.RegisterName("setLayoutMode:"), value)
}





