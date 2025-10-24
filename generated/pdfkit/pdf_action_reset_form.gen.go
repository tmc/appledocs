// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PDFActionResetForm */


/* debug [class_header]: Header for PDFActionResetForm */
// The class instance for the [PDFActionResetForm] class.
var (
	PDFActionResetFormClass     _PDFActionResetFormClass
	PDFActionResetFormClassOnce sync.Once
)

func getPDFActionResetFormClass() _PDFActionResetFormClass {
	PDFActionResetFormClassOnce.Do(func() {
		PDFActionResetFormClass = _PDFActionResetFormClass{objc.GetClass("PDFActionResetForm")}
	})
	return PDFActionResetFormClass
}

type _PDFActionResetFormClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFActionResetForm */
// An interface definition for the [PDFActionResetForm] class.
type IPDFActionResetForm interface {
	IPDFAction
	
/* debug [class_interface_properties]: Properties for PDFActionResetForm */
	// properties:
	Fields() []string
	SetFields(value []string)
	FieldsIncludedAreCleared() bool
	SetFieldsIncludedAreCleared(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFActionResetForm */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFActionResetForm */
// Alloc allocates a new instance without initialization.
func (pc _PDFActionResetFormClass) Alloc() PDFActionResetForm {
	rv := objc.Send[PDFActionResetForm](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PDFActionResetFormClass) New() PDFActionResetForm {
	rv := objc.Send[PDFActionResetForm](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFActionResetForm) Init() PDFActionResetForm {
	rv := objc.Send[PDFActionResetForm](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFActionResetForm) Autorelease() PDFActionResetForm {
	rv := objc.Send[PDFActionResetForm](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFActionResetForm creates a new PDFActionResetForm instance.
func NewPDFActionResetForm() PDFActionResetForm {
	return getPDFActionResetFormClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFActionResetForm */
// , a subclass of , defines methods for getting and clearing fields in a PDF form.
//
// A object represents an action associated with a PDF form.


// , a subclass of , defines methods for getting and clearing fields in a PDF form.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionResetForm
type PDFActionResetForm struct {
	PDFAction
}

// PDFActionResetFormFrom constructs a [PDFActionResetForm] from an unsafe.Pointer.
//
// , a subclass of , defines methods for getting and clearing fields in a PDF form.
func PDFActionResetFormFrom(ptr unsafe.Pointer) PDFActionResetForm {
	return PDFActionResetForm{
		PDFAction: PDFActionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFActionResetForm */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFActionResetForm */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFActionResetForm */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFActionResetForm */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFActionResetForm */

// Returns an array of fields associated with the reset action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionResetForm/fields
func (p_ PDFActionResetForm) Fields() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("fields"))
	return rv
}/* debug [instance_properties/getter]: fields */


// Returns an array of fields associated with the reset action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionResetForm/fields
func (p_ PDFActionResetForm) SetFields(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setFields:"), nsArray)
}/* debug [instance_properties/setter]: fields */


// Sets whether the fields associated with the reset action are cleared when the action is performed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionResetForm/fieldsIncludedAreCleared
func (p_ PDFActionResetForm) FieldsIncludedAreCleared() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("fieldsIncludedAreCleared"))
	return rv
}/* debug [instance_properties/getter]: fieldsIncludedAreCleared */


// Sets whether the fields associated with the reset action are cleared when the action is performed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionResetForm/fieldsIncludedAreCleared
func (p_ PDFActionResetForm) SetFieldsIncludedAreCleared(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFieldsIncludedAreCleared:"), value)
}/* debug [instance_properties/setter]: fieldsIncludedAreCleared */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFActionResetForm */


