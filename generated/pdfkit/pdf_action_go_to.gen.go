// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PDFActionGoTo */


/* debug [class_header]: Header for PDFActionGoTo */
// The class instance for the [PDFActionGoTo] class.
var (
	PDFActionGoToClass     _PDFActionGoToClass
	PDFActionGoToClassOnce sync.Once
)

func getPDFActionGoToClass() _PDFActionGoToClass {
	PDFActionGoToClassOnce.Do(func() {
		PDFActionGoToClass = _PDFActionGoToClass{objc.GetClass("PDFActionGoTo")}
	})
	return PDFActionGoToClass
}

type _PDFActionGoToClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFActionGoTo */
// An interface definition for the [PDFActionGoTo] class.
type IPDFActionGoTo interface {
	IPDFAction
	
/* debug [class_interface_properties]: Properties for PDFActionGoTo */
	// properties:
	Destination() IPDFDestination
	SetDestination(value IPDFDestination)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFActionGoTo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFActionGoTo */
// Alloc allocates a new instance without initialization.
func (pc _PDFActionGoToClass) Alloc() PDFActionGoTo {
	rv := objc.Send[PDFActionGoTo](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PDFActionGoToClass) New() PDFActionGoTo {
	rv := objc.Send[PDFActionGoTo](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFActionGoTo) Init() PDFActionGoTo {
	rv := objc.Send[PDFActionGoTo](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFActionGoTo) Autorelease() PDFActionGoTo {
	rv := objc.Send[PDFActionGoTo](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFActionGoTo creates a new PDFActionGoTo instance.
func NewPDFActionGoTo() PDFActionGoTo {
	return getPDFActionGoToClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFActionGoTo */
// , a subclass of , defines methods for getting and setting the destination of a go-to action.
//
// A object represents the action of going to a specific location within the PDF document.


// , a subclass of , defines methods for getting and setting the destination of a go-to action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionGoTo
type PDFActionGoTo struct {
	PDFAction
}

// PDFActionGoToFrom constructs a [PDFActionGoTo] from an unsafe.Pointer.
//
// , a subclass of , defines methods for getting and setting the destination of a go-to action.
func PDFActionGoToFrom(ptr unsafe.Pointer) PDFActionGoTo {
	return PDFActionGoTo{
		PDFAction: PDFActionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFActionGoTo */

// Initializes the go-to action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionGoTo/init(destination:)
func NewPDFActionGoToWithDestination(destination IPDFDestination) PDFActionGoTo {
	instance := getPDFActionGoToClass().Alloc()
	rv := objc.Send[PDFActionGoTo](instance.ID, objc.Sel("initWithDestination:"), destination)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPDFActionGoToWithDestination */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFActionGoTo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFActionGoTo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFActionGoTo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFActionGoTo */

// Returns the destination associated with the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionGoTo/destination
func (p_ PDFActionGoTo) Destination() IPDFDestination {
	rv := objc.Send[PDFDestination](p_.ID, objc.Sel("destination"))
	return rv
}/* debug [instance_properties/getter]: destination */


// Returns the destination associated with the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionGoTo/destination
func (p_ PDFActionGoTo) SetDestination(value IPDFDestination) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDestination:"), value)
}/* debug [instance_properties/setter]: destination */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFActionGoTo */


