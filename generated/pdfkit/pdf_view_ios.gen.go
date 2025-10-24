//go:build darwin && ios

// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for PDFView


// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/findInteraction
func (p_ PDFView) FindInteraction() FindInteraction /* not a class type */ {
	rv := objc.Send[FindInteraction](p_.ID, objc.Sel("findInteraction"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/isFindInteractionEnabled
func (p_ PDFView) FindInteractionEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("findInteractionEnabled"))
	return rv
}
func (p_ PDFView) SetFindInteractionEnabled(value bool) {
	p_.ID.Send(objc.RegisterName("setFindInteractionEnabled:"), value)
}






