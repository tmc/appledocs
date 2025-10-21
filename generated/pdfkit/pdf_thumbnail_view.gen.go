// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/appkit"
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
}

// An object that contains a set of thumbnails, each of which represents a page in a PDF document.
//
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
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/allowsDragging
func (p_ PDFThumbnailView) AllowsDragging() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsDragging"))
	return rv
}


// SetAllowsDragging sets the value of the allowsDragging property.
// Returns a Boolean value indicating whether users can drag thumbnails (that is, re-order pages in the document) within the thumbnail view.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/allowsDragging
func (p_ PDFThumbnailView) SetAllowsDragging(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowsDragging:"), value)
}
// Returns a Boolean value indicating whether users can select multiple thumbnails in the thumbnail view at one time.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/allowsMultipleSelection
func (p_ PDFThumbnailView) AllowsMultipleSelection() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsMultipleSelection"))
	return rv
}


// SetAllowsMultipleSelection sets the value of the allowsMultipleSelection property.
// Returns a Boolean value indicating whether users can select multiple thumbnails in the thumbnail view at one time.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/allowsMultipleSelection
func (p_ PDFThumbnailView) SetAllowsMultipleSelection(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowsMultipleSelection:"), value)
}
// Returns the color used in the background of the thumbnail view.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/backgroundColor
func (p_ PDFThumbnailView) BackgroundColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("backgroundColor"))
	return rv
}


// SetBackgroundColor sets the value of the backgroundColor property.
// Returns the color used in the background of the thumbnail view.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/backgroundColor
func (p_ PDFThumbnailView) SetBackgroundColor(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBackgroundColor:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/contentInset
func (p_ PDFThumbnailView) ContentInset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("contentInset"))
	return rv
}


// SetContentInset sets the value of the contentInset property.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/contentInset
func (p_ PDFThumbnailView) SetContentInset(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentInset:"), value)
}
// Returns the font used to label the thumbnails.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/labelFont
func (p_ PDFThumbnailView) LabelFont() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("labelFont"))
	return rv
}


// SetLabelFont sets the value of the labelFont property.
// Returns the font used to label the thumbnails.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/labelFont
func (p_ PDFThumbnailView) SetLabelFont(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLabelFont:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/layoutMode
func (p_ PDFThumbnailView) LayoutMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("layoutMode"))
	return rv
}


// SetLayoutMode sets the value of the layoutMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/layoutMode
func (p_ PDFThumbnailView) SetLayoutMode(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLayoutMode:"), value)
}
// Returns the maximum number of columns of thumbnails the thumbnail view can display.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/maximumNumberOfColumns
func (p_ PDFThumbnailView) MaximumNumberOfColumns() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("maximumNumberOfColumns"))
	return rv
}


// SetMaximumNumberOfColumns sets the value of the maximumNumberOfColumns property.
// Returns the maximum number of columns of thumbnails the thumbnail view can display.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/maximumNumberOfColumns
func (p_ PDFThumbnailView) SetMaximumNumberOfColumns(value uint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMaximumNumberOfColumns:"), value)
}
// Returns the object associated with the thumbnail view.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/pdfView
func (p_ PDFThumbnailView) PDFView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("PDFView"))
	return rv
}


// SetPDFView sets the value of the PDFView property.
// Returns the object associated with the thumbnail view.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/pdfView
func (p_ PDFThumbnailView) SetPDFView(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPDFView:"), value)
}
// Returns an array of PDF pages that correspond to the selected thumbnails in the thumbnail view.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/selectedPages
func (p_ PDFThumbnailView) SelectedPages() []PDFPage {
	rv := objc.Send[[]PDFPage](p_.ID, objc.Sel("selectedPages"))
	return rv
}

// Returns the maximum width and height of the thumbnails in the thumbnail view.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/thumbnailSize
func (p_ PDFThumbnailView) ThumbnailSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](p_.ID, objc.Sel("thumbnailSize"))
	return rv
}


// SetThumbnailSize sets the value of the thumbnailSize property.
// Returns the maximum width and height of the thumbnails in the thumbnail view.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFThumbnailView/thumbnailSize
func (p_ PDFThumbnailView) SetThumbnailSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setThumbnailSize:"), value)
}


