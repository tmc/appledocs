// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [PDFDestination] class.
type IPDFDestination interface {
	objectivec.IObject
	// properties:
	Action() IPDFAction
	SetAction(value IPDFAction)
	ModificationDate() objc.IObject /* cross-framework: Date */
	SetModificationDate(value objc.IObject /* cross-framework: Date */)
	Page() IPDFPage
	SetPage(value IPDFPage)
	Type() objc.IObject /* cross-framework: NSString */
	SetType(value objc.IObject /* cross-framework: NSString */)
	UserName() objc.IObject /* cross-framework: NSString */
	SetUserName(value objc.IObject /* cross-framework: NSString */)
	Point() objc.IObject /* cross-framework: Point */
	SetPoint(value objc.IObject /* cross-framework: Point */)
	Zoom() float64
	SetZoom(value float64)
	CurrentDestination() IPDFDestination
	SetCurrentDestination(value IPDFDestination)
	KPDFDestinationUnspecifiedValue() float64
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (pc _PDFDestinationClass) Alloc() PDFDestination {
	rv := objc.Send[PDFDestination](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// An object that represents an action for a PDF element, such as a link annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/action
func (p_ PDFDestination) Action() IPDFAction {
	rv := objc.Send[PDFAction](p_.ID, objc.Sel("action"))
	return rv
}


// An object that represents an action for a PDF element, such as a link annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/action
func (p_ PDFDestination) SetAction(value IPDFAction) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAction:"), value)
}


// Returns the modification date of the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/modificationdate
func (p_ PDFDestination) ModificationDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("modificationDate"))
	return rv
}


// Returns the modification date of the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/modificationdate
func (p_ PDFDestination) SetModificationDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setModificationDate:"), value)
}


// Returns the page that the annotation is associated with.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/page
func (p_ PDFDestination) Page() IPDFPage {
	rv := objc.Send[PDFPage](p_.ID, objc.Sel("page"))
	return rv
}


// Returns the page that the annotation is associated with.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/page
func (p_ PDFDestination) SetPage(value IPDFPage) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPage:"), value)
}


// Returns the type of the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/type
func (p_ PDFDestination) Type() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("type"))
	return rv
}


// Returns the type of the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/type
func (p_ PDFDestination) SetType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setType:"), value)
}


// Returns the name of the user who created the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/username
func (p_ PDFDestination) UserName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("userName"))
	return rv
}


// Returns the name of the user who created the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/username
func (p_ PDFDestination) SetUserName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUserName:"), value)
}


// Returns the point, in page space, that the destination refers to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfdestination/point
func (p_ PDFDestination) Point() objc.IObject /* cross-framework: Point */ {
	rv := objc.Send[corefoundation.Point](p_.ID, objc.Sel("point"))
	return rv
}


// Returns the point, in page space, that the destination refers to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfdestination/point
func (p_ PDFDestination) SetPoint(value objc.IObject /* cross-framework: Point */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPoint:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfdestination/zoom
func (p_ PDFDestination) Zoom() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("zoom"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfdestination/zoom
func (p_ PDFDestination) SetZoom(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setZoom:"), value)
}


// Returns a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfview/currentdestination
func (p_ PDFDestination) CurrentDestination() IPDFDestination {
	rv := objc.Send[PDFDestination](p_.ID, objc.Sel("currentDestination"))
	return rv
}


// Returns a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfview/currentdestination
func (p_ PDFDestination) SetCurrentDestination(value IPDFDestination) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentDestination:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/kpdfdestinationunspecifiedvalue
func (p_ PDFDestination) KPDFDestinationUnspecifiedValue() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("kPDFDestinationUnspecifiedValue"))
	return rv
}



