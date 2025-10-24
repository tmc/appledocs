// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PDFAction */


/* debug [class_header]: Header for PDFAction */
// The class instance for the [PDFAction] class.
var (
	PDFActionClass     _PDFActionClass
	PDFActionClassOnce sync.Once
)

func getPDFActionClass() _PDFActionClass {
	PDFActionClassOnce.Do(func() {
		PDFActionClass = _PDFActionClass{objc.GetClass("PDFAction")}
	})
	return PDFActionClass
}

type _PDFActionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFAction */
// An interface definition for the [PDFAction] class.
type IPDFAction interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PDFAction */
	// properties:
	Type() objc.IObject /* cross-framework: NSString */
	Action() IPDFAction
	SetAction(value IPDFAction)
	ModificationDate() foundation.Date
	SetModificationDate(value foundation.Date)
	Page() IPDFPage
	SetPage(value IPDFPage)
	UserName() objc.IObject /* cross-framework: NSString */
	SetUserName(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFAction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFAction */
// Alloc allocates a new instance without initialization.
func (pc _PDFActionClass) Alloc() PDFAction {
	rv := objc.Send[PDFAction](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PDFActionClass) New() PDFAction {
	rv := objc.Send[PDFAction](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFAction) Init() PDFAction {
	rv := objc.Send[PDFAction](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFAction) Autorelease() PDFAction {
	rv := objc.Send[PDFAction](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAction creates a new PDFAction instance.
func NewPDFAction() PDFAction {
	return getPDFActionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFAction */
// An action that is performed when, for example, a PDF annotation is activated or an outline item is clicked.
//
// A object represents an action associated with a PDF element, such as an annotation or a link, that the viewer application can perform. See the Adobe PDF Specification for more about actions and action types. is an abstract superclass of the following concrete classes:


// An action that is performed when, for example, a PDF annotation is activated or an outline item is clicked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAction
type PDFAction struct {
	objectivec.Object
}

// PDFActionFrom constructs a [PDFAction] from an unsafe.Pointer.
//
// An action that is performed when, for example, a PDF annotation is activated or an outline item is clicked.
func PDFActionFrom(ptr unsafe.Pointer) PDFAction {
	return PDFAction{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFAction *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFAction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFAction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFAction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFAction */

// Returns the type of the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAction/type
func (p_ PDFAction) Type() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// An object that represents an action for a PDF element, such as a link annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/action
func (p_ PDFAction) Action() IPDFAction {
	rv := objc.Send[PDFAction](p_.ID, objc.Sel("action"))
	return rv
}/* debug [instance_properties/getter]: action */


// An object that represents an action for a PDF element, such as a link annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/action
func (p_ PDFAction) SetAction(value IPDFAction) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAction:"), value)
}/* debug [instance_properties/setter]: action */


// Returns the modification date of the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/modificationdate
func (p_ PDFAction) ModificationDate() foundation.Date {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("modificationDate"))
	return rv
}/* debug [instance_properties/getter]: modificationDate */


// Returns the modification date of the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/modificationdate
func (p_ PDFAction) SetModificationDate(value foundation.Date) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setModificationDate:"), value)
}/* debug [instance_properties/setter]: modificationDate */


// Returns the page that the annotation is associated with.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/page
func (p_ PDFAction) Page() IPDFPage {
	rv := objc.Send[PDFPage](p_.ID, objc.Sel("page"))
	return rv
}/* debug [instance_properties/getter]: page */


// Returns the page that the annotation is associated with.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/page
func (p_ PDFAction) SetPage(value IPDFPage) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPage:"), value)
}/* debug [instance_properties/setter]: page */


// Returns the name of the user who created the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/username
func (p_ PDFAction) UserName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("userName"))
	return rv
}/* debug [instance_properties/getter]: userName */


// Returns the name of the user who created the annotation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/username
func (p_ PDFAction) SetUserName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUserName:"), value)
}/* debug [instance_properties/setter]: userName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFAction */



