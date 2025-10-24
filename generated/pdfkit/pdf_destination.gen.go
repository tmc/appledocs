// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class PDFDestination */


/* debug [class_header]: Header for PDFDestination */
// The class instance for the [PDFDestination] class.
var (
	PDFDestinationClass     _PDFDestinationClass
	PDFDestinationClassOnce sync.Once
)

func getPDFDestinationClass() _PDFDestinationClass {
	PDFDestinationClassOnce.Do(func() {
		PDFDestinationClass = _PDFDestinationClass{objc.GetClass("PDFDestination")}
	})
	return PDFDestinationClass
}

type _PDFDestinationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFDestination */
// An interface definition for the [PDFDestination] class.
type IPDFDestination interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PDFDestination */
	// properties:
	Page() IPDFPage
	Point() vision.Point
	Zoom() float64
	SetZoom(value float64)
	Action() IPDFAction
	SetAction(value IPDFAction)
	ModificationDate() foundation.Date
	SetModificationDate(value foundation.Date)
	Type() objc.IObject /* cross-framework: NSString */
	SetType(value objc.IObject /* cross-framework: NSString */)
	UserName() objc.IObject /* cross-framework: NSString */
	SetUserName(value objc.IObject /* cross-framework: NSString */)
	CurrentDestination() IPDFDestination
	SetCurrentDestination(value IPDFDestination)
	KPDFDestinationUnspecifiedValue() float64
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFDestination */
	// methods:
	Compare(destination IPDFDestination) ComparisonResult /* not a class type */
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFDestination */
// Alloc allocates a new instance without initialization.
func (pc _PDFDestinationClass) Alloc() PDFDestination {
	rv := objc.Send[PDFDestination](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PDFDestinationClass) New() PDFDestination {
	rv := objc.Send[PDFDestination](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFDestination) Init() PDFDestination {
	rv := objc.Send[PDFDestination](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFDestination) Autorelease() PDFDestination {
	rv := objc.Send[PDFDestination](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFDestination creates a new PDFDestination instance.
func NewPDFDestination() PDFDestination {
	return getPDFDestinationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFDestination */
// A object describes a point on a PDF page.
//
// In typical usage, you do not initialize objects but rather get them as either attributes of or objects, or in response to the method .


// A object describes a point on a PDF page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDestination
type PDFDestination struct {
	objectivec.Object
}

// PDFDestinationFrom constructs a [PDFDestination] from an unsafe.Pointer.
//
// A object describes a point on a PDF page.
func PDFDestinationFrom(ptr unsafe.Pointer) PDFDestination {
	return PDFDestination{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFDestination */

// Initializes the destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDestination/init(page:at:)
func NewPDFDestinationWithPageAtPoint(page IPDFPage, point vision.Point) PDFDestination {
	instance := getPDFDestinationClass().Alloc()
	rv := objc.Send[PDFDestination](instance.ID, objc.Sel("initWithPage:atPoint:"), page, point)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPDFDestinationWithPageAtPoint */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFDestination */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFDestination */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFDestination */

// Returns a comparison result that indicates the location of the destination in the document, relative to the current position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDestination/compare(_:)
func (p_ PDFDestination) Compare(destination IPDFDestination) ComparisonResult /* not a class type */ {
	rv := objc.Send[ComparisonResult](p_.ID, objc.Sel("compare:"), destination)
	return rv
}/* debug [instance_methods/method]: Compare */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFDestination */

// Returns the page that the destination refers to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDestination/page
func (p_ PDFDestination) Page() IPDFPage {
	rv := objc.Send[PDFPage](p_.ID, objc.Sel("page"))
	return rv
}/* debug [instance_properties/getter]: page */


// Returns the point, in page space, that the destination refers to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDestination/point
func (p_ PDFDestination) Point() vision.Point {
	rv := objc.Send[vision.Point](p_.ID, objc.Sel("point"))
	return rv
}/* debug [instance_properties/getter]: point */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDestination/zoom
func (p_ PDFDestination) Zoom() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("zoom"))
	return rv
}/* debug [instance_properties/getter]: zoom */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDestination/zoom
func (p_ PDFDestination) SetZoom(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setZoom:"), value)
}/* debug [instance_properties/setter]: zoom */


// An object that represents an action for a PDF element, such as a link annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/action
func (p_ PDFDestination) Action() IPDFAction {
	rv := objc.Send[PDFAction](p_.ID, objc.Sel("action"))
	return rv
}/* debug [instance_properties/getter]: action */


// An object that represents an action for a PDF element, such as a link annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/action
func (p_ PDFDestination) SetAction(value IPDFAction) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAction:"), value)
}/* debug [instance_properties/setter]: action */


// Returns the modification date of the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/modificationdate
func (p_ PDFDestination) ModificationDate() foundation.Date {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("modificationDate"))
	return rv
}/* debug [instance_properties/getter]: modificationDate */


// Returns the modification date of the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/modificationdate
func (p_ PDFDestination) SetModificationDate(value foundation.Date) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setModificationDate:"), value)
}/* debug [instance_properties/setter]: modificationDate */


// Returns the type of the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/type
func (p_ PDFDestination) Type() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// Returns the type of the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/type
func (p_ PDFDestination) SetType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setType:"), value)
}/* debug [instance_properties/setter]: type */


// Returns the name of the user who created the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/username
func (p_ PDFDestination) UserName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("userName"))
	return rv
}/* debug [instance_properties/getter]: userName */


// Returns the name of the user who created the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/username
func (p_ PDFDestination) SetUserName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUserName:"), value)
}/* debug [instance_properties/setter]: userName */


// Returns a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfview/currentdestination
func (p_ PDFDestination) CurrentDestination() IPDFDestination {
	rv := objc.Send[PDFDestination](p_.ID, objc.Sel("currentDestination"))
	return rv
}/* debug [instance_properties/getter]: currentDestination */


// Returns a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfview/currentdestination
func (p_ PDFDestination) SetCurrentDestination(value IPDFDestination) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentDestination:"), value)
}/* debug [instance_properties/setter]: currentDestination */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/kpdfdestinationunspecifiedvalue
func (p_ PDFDestination) KPDFDestinationUnspecifiedValue() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("kPDFDestinationUnspecifiedValue"))
	return rv
}/* debug [instance_properties/getter]: kPDFDestinationUnspecifiedValue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFDestination */


