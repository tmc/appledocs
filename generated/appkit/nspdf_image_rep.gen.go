// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSPDFImageRep */


/* debug [class_header]: Header for NSPDFImageRep */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFImageRep */
// An interface definition for the [PDFImageRep] class.
type IPDFImageRep interface {
	IImageRep
	
/* debug [class_interface_properties]: Properties for PDFImageRep */
	// properties:
	Bounds() Rect /* not a class type */
	CurrentPage() int
	SetCurrentPage(value int)
	PageCount() int
	PDFRepresentation() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFImageRep */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFImageRep */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFImageRep */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFImageRep */

// Returns a representation of an image initialized with the specified PDF data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFImageRep/init(data:)
func NewPDFImageRepWithData(pdfData objc.IObject /* cross-framework: NSData */) PDFImageRep {
	instance := getPDFImageRepClass().Alloc()
	rv := objc.Send[PDFImageRep](instance.ID, objc.Sel("initWithData:"), pdfData)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPDFImageRepWithData */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFImageRep */

// Creates and returns a representation of an image initialized with the specified PDF data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFImageRep/imageRepWithData:
func (pc _PDFImageRepClass) ImageRepWithData(pdfData objc.IObject /* cross-framework: NSData */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("imageRepWithData:"), pdfData)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ImageRepWithData) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFImageRep */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFImageRep */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFImageRep */

// The image representation’s bounding rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFImageRep/bounds
func (p_ PDFImageRep) Bounds() Rect /* not a class type */ {
	rv := objc.Send[Rect](p_.ID, objc.Sel("bounds"))
	return rv
}/* debug [instance_properties/getter]: bounds */


// The page currently displayed by the image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFImageRep/currentPage
func (p_ PDFImageRep) CurrentPage() int {
	rv := objc.Send[int](p_.ID, objc.Sel("currentPage"))
	return rv
}/* debug [instance_properties/getter]: currentPage */


// The page currently displayed by the image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFImageRep/currentPage
func (p_ PDFImageRep) SetCurrentPage(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentPage:"), value)
}/* debug [instance_properties/setter]: currentPage */


// The number of pages in the image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFImageRep/pageCount
func (p_ PDFImageRep) PageCount() int {
	rv := objc.Send[int](p_.ID, objc.Sel("pageCount"))
	return rv
}/* debug [instance_properties/getter]: pageCount */


// The PDF representation of the representation’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFImageRep/pdfRepresentation
func (p_ PDFImageRep) PDFRepresentation() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("PDFRepresentation"))
	return rv
}/* debug [instance_properties/getter]: PDFRepresentation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPDFImageRep */


