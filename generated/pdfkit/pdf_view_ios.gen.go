//go:build darwin && ios

// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

// iOS-only methods for PDFView


// Changes the scroll view to use a to layout and navigate pages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/usePageViewController(_:withViewOptions:)
func (p_ PDFView) UsePageViewControllerWithViewOptions(enable bool, viewOptions objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("usePageViewController:withViewOptions:"), enable, viewOptions)
}

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

// A Boolean value indicating whether the scroll view is using a .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/isUsingPageViewController
func (p_ PDFView) IsUsingPageViewController() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isUsingPageViewController"))
	return rv
}





