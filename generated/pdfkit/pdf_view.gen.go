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
	TakePasswordFrom(sender objectivec.IObject)
	CurrentDestination() PDFDestination
	CurrentPage() PDFPage
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	Document() PDFDocument
	SetDocument(value IPDFDocument)
	FindInteraction() unsafe.Pointer
	FindInteractionEnabled() bool
	SetFindInteractionEnabled(value bool)
	InMarkupMode() bool
	SetInMarkupMode(value bool)
	PageOverlayViewProvider() objc.ID
	SetPageOverlayViewProvider(value objc.ID)
	PageShadowsEnabled() bool
	SetPageShadowsEnabled(value bool)
	VisiblePages() []PDFPage
	IsFindInteractionEnabled() bool
	SetIsFindInteractionEnabled(value bool)
	IsInMarkupMode() bool
	SetIsInMarkupMode(value bool)
}

// An object that encapsulates the functionality of PDF Kit into a single widget that you can add to your application using Interface Builder.
//
// may be the only class you need to deal with for adding PDF functionality to your application. It lets you display PDF data and allows users to select content, navigate through a document, set zoom level, and copy textual content to the Pasteboard. also keeps track of page history. You can subclass to create a custom PDF viewer. You can also create a custom PDF viewer by using the PDF Kit utility classes directly and not using at all.
//
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


// Unlocks with the password from the specified sender.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/takePasswordFrom(_:)
func (p_ PDFView) TakePasswordFrom(sender objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("takePasswordFrom:"), sender)
}

// Returns a object representing the current page and the current point in the view specified in page space.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/currentDestination
func (p_ PDFView) CurrentDestination() PDFDestination {
	rv := objc.Send[PDFDestination](p_.ID, objc.Sel("currentDestination"))
	return rv
}

// Returns the current page.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/currentPage
func (p_ PDFView) CurrentPage() PDFPage {
	rv := objc.Send[PDFPage](p_.ID, objc.Sel("currentPage"))
	return rv
}

// Returns the view’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/delegate
func (p_ PDFView) Delegate() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// Returns the view’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/delegate
func (p_ PDFView) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}

// Returns the document associated with a object.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/document
func (p_ PDFView) Document() PDFDocument {
	rv := objc.Send[PDFDocument](p_.ID, objc.Sel("document"))
	return rv
}


// SetDocument sets the value of the document property.
// Returns the document associated with a object.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/document
func (p_ PDFView) SetDocument(value IPDFDocument) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDocument:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/findInteraction
func (p_ PDFView) FindInteraction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("findInteraction"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/isFindInteractionEnabled
func (p_ PDFView) FindInteractionEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("findInteractionEnabled"))
	return rv
}


// SetFindInteractionEnabled sets the value of the findInteractionEnabled property.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/isFindInteractionEnabled
func (p_ PDFView) SetFindInteractionEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFindInteractionEnabled:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/isInMarkupMode
func (p_ PDFView) InMarkupMode() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("inMarkupMode"))
	return rv
}


// SetInMarkupMode sets the value of the inMarkupMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/isInMarkupMode
func (p_ PDFView) SetInMarkupMode(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInMarkupMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/pageOverlayViewProvider
func (p_ PDFView) PageOverlayViewProvider() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("pageOverlayViewProvider"))
	return rv
}


// SetPageOverlayViewProvider sets the value of the pageOverlayViewProvider property.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/pageOverlayViewProvider
func (p_ PDFView) SetPageOverlayViewProvider(value objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPageOverlayViewProvider:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/pageShadowsEnabled
func (p_ PDFView) PageShadowsEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("pageShadowsEnabled"))
	return rv
}


// SetPageShadowsEnabled sets the value of the pageShadowsEnabled property.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/pageShadowsEnabled
func (p_ PDFView) SetPageShadowsEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPageShadowsEnabled:"), value)
}

// Returns an array of objects that represent the currently visible pages.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/visiblePages
func (p_ PDFView) VisiblePages() []PDFPage {
	rv := objc.Send[[]PDFPage](p_.ID, objc.Sel("visiblePages"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfview/isfindinteractionenabled
func (p_ PDFView) IsFindInteractionEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isFindInteractionEnabled"))
	return rv
}


// SetIsFindInteractionEnabled sets the value of the isFindInteractionEnabled property.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfview/isfindinteractionenabled
func (p_ PDFView) SetIsFindInteractionEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsFindInteractionEnabled:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfview/isinmarkupmode
func (p_ PDFView) IsInMarkupMode() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isInMarkupMode"))
	return rv
}


// SetIsInMarkupMode sets the value of the isInMarkupMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfview/isinmarkupmode
func (p_ PDFView) SetIsInMarkupMode(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsInMarkupMode:"), value)
}




