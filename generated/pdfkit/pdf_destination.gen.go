// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
}

// A object describes a point on a PDF page.
//
// In typical usage, you do not initialize objects but rather get them as either attributes of or objects, or in response to the method .
//
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


// Returns the page that the destination refers to.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDestination/page
func (p_ PDFDestination) Page() PDFPage {
	rv := objc.Send[PDFPage](p_.ID, objc.Sel("page"))
	return rv
}

// An object that represents an action for a PDF element, such as a link annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/action
func (p_ PDFDestination) Action() PDFAction {
	rv := objc.Send[PDFAction](p_.ID, objc.Sel("action"))
	return rv
}


// SetAction sets the value of the action property.
// An object that represents an action for a PDF element, such as a link annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/action
func (p_ PDFDestination) SetAction(value IPDFAction) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAction:"), value)
}

// Returns the modification date of the annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/modificationdate
func (p_ PDFDestination) ModificationDate() foundation.Date {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("modificationDate"))
	return rv
}


// SetModificationDate sets the value of the modificationDate property.
// Returns the modification date of the annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/modificationdate
func (p_ PDFDestination) SetModificationDate(value foundation.IDate) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setModificationDate:"), value)
}

// Returns the type of the annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/type
func (p_ PDFDestination) Type() appkit.string {
	rv := objc.Send[appkit.string](p_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
// Returns the type of the annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/type
func (p_ PDFDestination) SetType(value appkit.string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setType:"), value)
}

// Returns the name of the user who created the annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/username
func (p_ PDFDestination) UserName() appkit.string {
	rv := objc.Send[appkit.string](p_.ID, objc.Sel("userName"))
	return rv
}


// SetUserName sets the value of the userName property.
// Returns the name of the user who created the annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/username
func (p_ PDFDestination) SetUserName(value appkit.string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUserName:"), value)
}

// Returns the point, in page space, that the destination refers to.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfdestination/point
func (p_ PDFDestination) Point() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](p_.ID, objc.Sel("point"))
	return rv
}


// SetPoint sets the value of the point property.
// Returns the point, in page space, that the destination refers to.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfdestination/point
func (p_ PDFDestination) SetPoint(value coregraphics.CGPoint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPoint:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfdestination/zoom
func (p_ PDFDestination) Zoom() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("zoom"))
	return rv
}


// SetZoom sets the value of the zoom property.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfdestination/zoom
func (p_ PDFDestination) SetZoom(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setZoom:"), value)
}

// Returns a
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfview/currentdestination
func (p_ PDFDestination) CurrentDestination() PDFDestination {
	rv := objc.Send[PDFDestination](p_.ID, objc.Sel("currentDestination"))
	return rv
}


// SetCurrentDestination sets the value of the currentDestination property.
// Returns a

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfview/currentdestination
func (p_ PDFDestination) SetCurrentDestination(value IPDFDestination) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentDestination:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/kpdfdestinationunspecifiedvalue
func (p_ PDFDestination) KPDFDestinationUnspecifiedValue() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("kPDFDestinationUnspecifiedValue"))
	return rv
}



