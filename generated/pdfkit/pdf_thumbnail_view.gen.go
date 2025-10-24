// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PDFThumbnailView] class.
var (
	PDFThumbnailViewClass     _PDFThumbnailViewClass
	PDFThumbnailViewClassOnce sync.Once
)

func getPDFThumbnailViewClass() _PDFThumbnailViewClass {
	PDFThumbnailViewClassOnce.Do(func() {
		PDFThumbnailViewClass = _PDFThumbnailViewClass{objc.GetClass("PDFThumbnailView")}
	})
	return PDFThumbnailViewClass
}

type _PDFThumbnailViewClass struct {
	class objc.Class
}

// An interface definition for the [PDFThumbnailView] class.
type IPDFThumbnailView interface {
	appkit.IView
	// properties:
	AllowsDragging() bool
	SetAllowsDragging(value bool)
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	BackgroundColor() objc.IObject /* cross-framework: Color */
	SetBackgroundColor(value objc.IObject /* cross-framework: Color */)
	LabelFont() objc.IObject /* cross-framework: Font */
	SetLabelFont(value objc.IObject /* cross-framework: Font */)
	MaximumNumberOfColumns() uint
	SetMaximumNumberOfColumns(value uint)
	PDFView() IPDFView
	SetPDFView(value IPDFView)
	SelectedPages() []IPDFPage
	ThumbnailSize() objc.IObject /* cross-framework: Size */
	SetThumbnailSize(value objc.IObject /* cross-framework: Size */)
	// methods:
}

// An object that contains a set of thumbnails, each of which represents a page in a PDF document.


// An object that contains a set of thumbnails, each of which represents a page in a PDF document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView
type PDFThumbnailView struct {
	appkit.View
}

// PDFThumbnailViewFrom constructs a [PDFThumbnailView] from an unsafe.Pointer.
//
// An object that contains a set of thumbnails, each of which represents a page in a PDF document.
func PDFThumbnailViewFrom(ptr unsafe.Pointer) PDFThumbnailView {
	return PDFThumbnailView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFThumbnailViewClass) Alloc() PDFThumbnailView {
	rv := objc.Send[PDFThumbnailView](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFThumbnailViewClass) New() PDFThumbnailView {
	rv := objc.Send[PDFThumbnailView](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFThumbnailView) Init() PDFThumbnailView {
	rv := objc.Send[PDFThumbnailView](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFThumbnailView) Autorelease() PDFThumbnailView {
	rv := objc.Send[PDFThumbnailView](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFThumbnailView creates a new PDFThumbnailView instance.
func NewPDFThumbnailView() PDFThumbnailView {
	return getPDFThumbnailViewClass().New()
}



// Returns a Boolean value indicating whether users can drag thumbnails (that is, re-order pages in the document) within the thumbnail view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/allowsDragging
func (p_ PDFThumbnailView) AllowsDragging() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsDragging"))
	return rv
}


// Returns a Boolean value indicating whether users can drag thumbnails (that is, re-order pages in the document) within the thumbnail view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/allowsDragging
func (p_ PDFThumbnailView) SetAllowsDragging(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowsDragging:"), value)
}


// Returns a Boolean value indicating whether users can select multiple thumbnails in the thumbnail view at one time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/allowsMultipleSelection
func (p_ PDFThumbnailView) AllowsMultipleSelection() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsMultipleSelection"))
	return rv
}


// Returns a Boolean value indicating whether users can select multiple thumbnails in the thumbnail view at one time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/allowsMultipleSelection
func (p_ PDFThumbnailView) SetAllowsMultipleSelection(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowsMultipleSelection:"), value)
}


// Returns the color used in the background of the thumbnail view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/backgroundColor
func (p_ PDFThumbnailView) BackgroundColor() objc.IObject /* cross-framework: Color */ {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("backgroundColor"))
	return rv
}


// Returns the color used in the background of the thumbnail view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/backgroundColor
func (p_ PDFThumbnailView) SetBackgroundColor(value objc.IObject /* cross-framework: Color */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBackgroundColor:"), value)
}


// Returns the font used to label the thumbnails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/labelFont
func (p_ PDFThumbnailView) LabelFont() objc.IObject /* cross-framework: Font */ {
	rv := objc.Send[appkit.Font](p_.ID, objc.Sel("labelFont"))
	return rv
}


// Returns the font used to label the thumbnails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/labelFont
func (p_ PDFThumbnailView) SetLabelFont(value objc.IObject /* cross-framework: Font */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLabelFont:"), value)
}


// Returns the maximum number of columns of thumbnails the thumbnail view can display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/maximumNumberOfColumns
func (p_ PDFThumbnailView) MaximumNumberOfColumns() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("maximumNumberOfColumns"))
	return rv
}


// Returns the maximum number of columns of thumbnails the thumbnail view can display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/maximumNumberOfColumns
func (p_ PDFThumbnailView) SetMaximumNumberOfColumns(value uint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMaximumNumberOfColumns:"), value)
}


// Returns the object associated with the thumbnail view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/pdfView
func (p_ PDFThumbnailView) PDFView() IPDFView {
	rv := objc.Send[PDFView](p_.ID, objc.Sel("PDFView"))
	return rv
}


// Returns the object associated with the thumbnail view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/pdfView
func (p_ PDFThumbnailView) SetPDFView(value IPDFView) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPDFView:"), value)
}


// Returns an array of PDF pages that correspond to the selected thumbnails in the thumbnail view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/selectedPages
func (p_ PDFThumbnailView) SelectedPages() []IPDFPage {
	rv := objc.Send[[]PDFPage](p_.ID, objc.Sel("selectedPages"))
	return rv
}


// Returns the maximum width and height of the thumbnails in the thumbnail view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/thumbnailSize
func (p_ PDFThumbnailView) ThumbnailSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](p_.ID, objc.Sel("thumbnailSize"))
	return rv
}


// Returns the maximum width and height of the thumbnails in the thumbnail view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/thumbnailSize
func (p_ PDFThumbnailView) SetThumbnailSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setThumbnailSize:"), value)
}


