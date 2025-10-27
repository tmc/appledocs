// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [PDFImageRep] class.
var (
	PDFImageRepClass     _PDFImageRepClass
	PDFImageRepClassOnce sync.Once
)

func getPDFImageRepClass() _PDFImageRepClass {
	PDFImageRepClassOnce.Do(func() {
		PDFImageRepClass = _PDFImageRepClass{objc.GetClass("NSPDFImageRep")}
	})
	return PDFImageRepClass
}

type _PDFImageRepClass struct {
	class objc.Class
}





// An interface definition for the [PDFImageRep] class.
type IPDFImageRep interface {
	IImageRep
	

	// properties:
	Bounds() corefoundation.CGRect
	CurrentPage() int
	SetCurrentPage(value int)
	PageCount() int
	PDFRepresentation() foundation.foundation.INSData


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (pc _PDFImageRepClass) Alloc() PDFImageRep {
	rv := objc.Send[PDFImageRep](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PDFImageRepClass) New() PDFImageRep {
	rv := objc.Send[PDFImageRep](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFImageRep) Init() PDFImageRep {
	rv := objc.Send[PDFImageRep](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFImageRep) Autorelease() PDFImageRep {
	rv := objc.Send[PDFImageRep](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFImageRep creates a new PDFImageRep instance.
func NewPDFImageRep() PDFImageRep {
	return getPDFImageRepClass().New()
}





// An object that can render an image from a PDF format data stream.


// An object that can render an image from a PDF format data stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFImageRep
type PDFImageRep struct {
	ImageRep
}

// PDFImageRepFrom constructs a [PDFImageRep] from an unsafe.Pointer.
//
// An object that can render an image from a PDF format data stream.
func PDFImageRepFrom(ptr unsafe.Pointer) PDFImageRep {
	return PDFImageRep{
		ImageRep: ImageRepFrom(ptr),
	}
}






// Returns a representation of an image initialized with the specified PDF data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFImageRep/init(data:)
func NewPDFImageRepWithData(pdfData foundation.foundation.INSData) PDFImageRep {
	instance := getPDFImageRepClass().Alloc()
	rv := objc.Send[PDFImageRep](instance.ID, objc.Sel("initWithData:"), pdfData)
	rv.Autorelease()
	return rv
}







// Creates and returns a representation of an image initialized with the specified PDF data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFImageRep/imageRepWithData:
func (pc _PDFImageRepClass) ImageRepWithData(pdfData foundation.foundation.INSData) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("imageRepWithData:"), pdfData)
	return rv
}

















// The image representation’s bounding rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFImageRep/bounds
func (p_ PDFImageRep) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p_.ID, objc.Sel("bounds"))
	return rv
}


// The page currently displayed by the image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFImageRep/currentPage
func (p_ PDFImageRep) CurrentPage() int {
	rv := objc.Send[int](p_.ID, objc.Sel("currentPage"))
	return rv
}


// The page currently displayed by the image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFImageRep/currentPage
func (p_ PDFImageRep) SetCurrentPage(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentPage:"), value)
}


// The number of pages in the image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFImageRep/pageCount
func (p_ PDFImageRep) PageCount() int {
	rv := objc.Send[int](p_.ID, objc.Sel("pageCount"))
	return rv
}


// The PDF representation of the representation’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFImageRep/pdfRepresentation
func (p_ PDFImageRep) PDFRepresentation() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("PDFRepresentation"))
	return rv
}







