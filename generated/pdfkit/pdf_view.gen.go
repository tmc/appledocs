// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PDFView] class.
var (
	PDFViewClass     _PDFViewClass
	PDFViewClassOnce sync.Once
)

func getPDFViewClass() _PDFViewClass {
	PDFViewClassOnce.Do(func() {
		PDFViewClass = _PDFViewClass{objc.GetClass("PDFView")}
	})
	return PDFViewClass
}

type _PDFViewClass struct {
	class objc.Class
}

// An interface definition for the [PDFView] class.
type IPDFView interface {
	appkit.IView
	// properties:
	CurrentDestination() IPDFDestination
	CurrentPage() IPDFPage
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	Document() IPDFDocument
	SetDocument(value IPDFDocument)
	InMarkupMode() bool
	SetInMarkupMode(value bool)
	PageOverlayViewProvider() objc.ID
	SetPageOverlayViewProvider(value objc.ID)
	PageShadowsEnabled() bool
	SetPageShadowsEnabled(value bool)
	VisiblePages() []IPDFPage
	IsFindInteractionEnabled() bool
	SetIsFindInteractionEnabled(value bool)
	IsInMarkupMode() bool
	SetIsInMarkupMode(value bool)
	// methods:
}

// An object that encapsulates the functionality of PDF Kit into a single widget that you can add to your application using Interface Builder.
//
// may be the only class you need to deal with for adding PDF functionality to your application. It lets you display PDF data and allows users to select content, navigate through a document, set zoom level, and copy textual content to the Pasteboard. also keeps track of page history. You can subclass to create a custom PDF viewer. You can also create a custom PDF viewer by using the PDF Kit utility classes directly and not using at all.


// An object that encapsulates the functionality of PDF Kit into a single widget that you can add to your application using Interface Builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView
type PDFView struct {
	appkit.View
}

// PDFViewFrom constructs a [PDFView] from an unsafe.Pointer.
//
// An object that encapsulates the functionality of PDF Kit into a single widget that you can add to your application using Interface Builder.
func PDFViewFrom(ptr unsafe.Pointer) PDFView {
	return PDFView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFViewClass) Alloc() PDFView {
	rv := objc.Send[PDFView](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFViewClass) New() PDFView {
	rv := objc.Send[PDFView](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFView) Init() PDFView {
	rv := objc.Send[PDFView](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFView) Autorelease() PDFView {
	rv := objc.Send[PDFView](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFView creates a new PDFView instance.
func NewPDFView() PDFView {
	return getPDFViewClass().New()
}



// Returns a object representing the current page and the current point in the view specified in page space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/currentDestination
func (p_ PDFView) CurrentDestination() IPDFDestination {
	rv := objc.Send[PDFDestination](p_.ID, objc.Sel("currentDestination"))
	return rv
}


// Returns the current page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/currentPage
func (p_ PDFView) CurrentPage() IPDFPage {
	rv := objc.Send[PDFPage](p_.ID, objc.Sel("currentPage"))
	return rv
}


// Returns the view’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/delegate
func (p_ PDFView) Delegate() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("delegate"))
	return rv
}


// Returns the view’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/delegate
func (p_ PDFView) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}


// Returns the document associated with a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/document
func (p_ PDFView) Document() IPDFDocument {
	rv := objc.Send[PDFDocument](p_.ID, objc.Sel("document"))
	return rv
}


// Returns the document associated with a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/document
func (p_ PDFView) SetDocument(value IPDFDocument) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDocument:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/isInMarkupMode
func (p_ PDFView) InMarkupMode() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("inMarkupMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/isInMarkupMode
func (p_ PDFView) SetInMarkupMode(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInMarkupMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/pageOverlayViewProvider
func (p_ PDFView) PageOverlayViewProvider() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("pageOverlayViewProvider"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/pageOverlayViewProvider
func (p_ PDFView) SetPageOverlayViewProvider(value objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPageOverlayViewProvider:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/pageShadowsEnabled
func (p_ PDFView) PageShadowsEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("pageShadowsEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/pageShadowsEnabled
func (p_ PDFView) SetPageShadowsEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPageShadowsEnabled:"), value)
}


// Returns an array of objects that represent the currently visible pages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/visiblePages
func (p_ PDFView) VisiblePages() []IPDFPage {
	rv := objc.Send[[]PDFPage](p_.ID, objc.Sel("visiblePages"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfview/isfindinteractionenabled
func (p_ PDFView) IsFindInteractionEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isFindInteractionEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfview/isfindinteractionenabled
func (p_ PDFView) SetIsFindInteractionEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsFindInteractionEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfview/isinmarkupmode
func (p_ PDFView) IsInMarkupMode() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isInMarkupMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfview/isinmarkupmode
func (p_ PDFView) SetIsInMarkupMode(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsInMarkupMode:"), value)
}


