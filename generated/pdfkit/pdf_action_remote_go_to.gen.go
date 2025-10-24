// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class PDFActionRemoteGoTo */


/* debug [class_header]: Header for PDFActionRemoteGoTo */
// The class instance for the [PDFActionRemoteGoTo] class.
var (
	PDFActionRemoteGoToClass     _PDFActionRemoteGoToClass
	PDFActionRemoteGoToClassOnce sync.Once
)

func getPDFActionRemoteGoToClass() _PDFActionRemoteGoToClass {
	PDFActionRemoteGoToClassOnce.Do(func() {
		PDFActionRemoteGoToClass = _PDFActionRemoteGoToClass{objc.GetClass("PDFActionRemoteGoTo")}
	})
	return PDFActionRemoteGoToClass
}

type _PDFActionRemoteGoToClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFActionRemoteGoTo */
// An interface definition for the [PDFActionRemoteGoTo] class.
type IPDFActionRemoteGoTo interface {
	IPDFAction
	
/* debug [class_interface_properties]: Properties for PDFActionRemoteGoTo */
	// properties:
	PageIndex() uint
	SetPageIndex(value uint)
	Point() vision.Point
	SetPoint(value vision.Point)
	URL() objc.IObject /* cross-framework: NSURL */
	SetURL(value objc.IObject /* cross-framework: NSURL */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFActionRemoteGoTo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFActionRemoteGoTo */
// Alloc allocates a new instance without initialization.
func (pc _PDFActionRemoteGoToClass) Alloc() PDFActionRemoteGoTo {
	rv := objc.Send[PDFActionRemoteGoTo](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PDFActionRemoteGoToClass) New() PDFActionRemoteGoTo {
	rv := objc.Send[PDFActionRemoteGoTo](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFActionRemoteGoTo) Init() PDFActionRemoteGoTo {
	rv := objc.Send[PDFActionRemoteGoTo](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFActionRemoteGoTo) Autorelease() PDFActionRemoteGoTo {
	rv := objc.Send[PDFActionRemoteGoTo](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFActionRemoteGoTo creates a new PDFActionRemoteGoTo instance.
func NewPDFActionRemoteGoTo() PDFActionRemoteGoTo {
	return getPDFActionRemoteGoToClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFActionRemoteGoTo */
// , a subclass of , defines methods for getting and setting the destination of a go-to action that targets another document.


// , a subclass of , defines methods for getting and setting the destination of a go-to action that targets another document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionRemoteGoTo
type PDFActionRemoteGoTo struct {
	PDFAction
}

// PDFActionRemoteGoToFrom constructs a [PDFActionRemoteGoTo] from an unsafe.Pointer.
//
// , a subclass of , defines methods for getting and setting the destination of a go-to action that targets another document.
func PDFActionRemoteGoToFrom(ptr unsafe.Pointer) PDFActionRemoteGoTo {
	return PDFActionRemoteGoTo{
		PDFAction: PDFActionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFActionRemoteGoTo */

// Initializes the remote go-to action with the specified page index, point, and document URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionRemoteGoTo/init(pageIndex:at:fileURL:)
func NewPDFActionRemoteGoToWithPageIndexAtPointFileURL(pageIndex uint, point corefoundation.CGPoint, url objc.IObject /* cross-framework: NSURL */) PDFActionRemoteGoTo {
	instance := getPDFActionRemoteGoToClass().Alloc()
	rv := objc.Send[PDFActionRemoteGoTo](instance.ID, objc.Sel("initWithPageIndex:atPoint:fileURL:"), pageIndex, point, url)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPDFActionRemoteGoToWithPageIndexAtPointFileURL */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFActionRemoteGoTo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFActionRemoteGoTo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFActionRemoteGoTo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFActionRemoteGoTo */

// Returns the zero-based page index referenced by the remote go-to action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionRemoteGoTo/pageIndex
func (p_ PDFActionRemoteGoTo) PageIndex() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("pageIndex"))
	return rv
}/* debug [instance_properties/getter]: pageIndex */


// Returns the zero-based page index referenced by the remote go-to action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionRemoteGoTo/pageIndex
func (p_ PDFActionRemoteGoTo) SetPageIndex(value uint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPageIndex:"), value)
}/* debug [instance_properties/setter]: pageIndex */


// Sets the point, in page space, on the page referenced by the remote go-to action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionRemoteGoTo/point
func (p_ PDFActionRemoteGoTo) Point() vision.Point {
	rv := objc.Send[vision.Point](p_.ID, objc.Sel("point"))
	return rv
}/* debug [instance_properties/getter]: point */


// Sets the point, in page space, on the page referenced by the remote go-to action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionRemoteGoTo/point
func (p_ PDFActionRemoteGoTo) SetPoint(value vision.Point) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPoint:"), value)
}/* debug [instance_properties/setter]: point */


// Returns the URL of the document referenced by the remote go-to action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionRemoteGoTo/url
func (p_ PDFActionRemoteGoTo) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](p_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_properties/getter]: URL */


// Returns the URL of the document referenced by the remote go-to action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionRemoteGoTo/url
func (p_ PDFActionRemoteGoTo) SetURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setURL:"), value)
}/* debug [instance_properties/setter]: URL */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFActionRemoteGoTo */


