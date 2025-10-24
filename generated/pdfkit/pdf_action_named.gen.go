// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PDFActionNamed */


/* debug [class_header]: Header for PDFActionNamed */
// The class instance for the [PDFActionNamed] class.
var (
	PDFActionNamedClass     _PDFActionNamedClass
	PDFActionNamedClassOnce sync.Once
)

func getPDFActionNamedClass() _PDFActionNamedClass {
	PDFActionNamedClassOnce.Do(func() {
		PDFActionNamedClass = _PDFActionNamedClass{objc.GetClass("PDFActionNamed")}
	})
	return PDFActionNamedClass
}

type _PDFActionNamedClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFActionNamed */
// An interface definition for the [PDFActionNamed] class.
type IPDFActionNamed interface {
	IPDFAction
	
/* debug [class_interface_properties]: Properties for PDFActionNamed */
	// properties:
	Name() PDFActionNamedName
	SetName(value PDFActionNamedName)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFActionNamed */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFActionNamed */
// Alloc allocates a new instance without initialization.
func (pc _PDFActionNamedClass) Alloc() PDFActionNamed {
	rv := objc.Send[PDFActionNamed](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PDFActionNamedClass) New() PDFActionNamed {
	rv := objc.Send[PDFActionNamed](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFActionNamed) Init() PDFActionNamed {
	rv := objc.Send[PDFActionNamed](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFActionNamed) Autorelease() PDFActionNamed {
	rv := objc.Send[PDFActionNamed](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFActionNamed creates a new PDFActionNamed instance.
func NewPDFActionNamed() PDFActionNamed {
	return getPDFActionNamedClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFActionNamed */
// defines methods used to work with actions in PDF documents, some of which are named in the Adobe PDF Specification.
//
// A object represents an action with a defined name, such as “Go back” or “Zoom in.”


// defines methods used to work with actions in PDF documents, some of which are named in the Adobe PDF Specification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionNamed
type PDFActionNamed struct {
	PDFAction
}

// PDFActionNamedFrom constructs a [PDFActionNamed] from an unsafe.Pointer.
//
// defines methods used to work with actions in PDF documents, some of which are named in the Adobe PDF Specification.
func PDFActionNamedFrom(ptr unsafe.Pointer) PDFActionNamed {
	return PDFActionNamed{
		PDFAction: PDFActionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFActionNamed */

// Initializes the object with the specified named action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionNamed/init(name:)
func NewPDFActionNamedWithName(name PDFActionNamedName) PDFActionNamed {
	instance := getPDFActionNamedClass().Alloc()
	rv := objc.Send[PDFActionNamed](instance.ID, objc.Sel("initWithName:"), name)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPDFActionNamedWithName */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFActionNamed */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFActionNamed */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFActionNamed */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFActionNamed */

// Returns the name of the named action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionNamed/name
func (p_ PDFActionNamed) Name() PDFActionNamedName {
	rv := objc.Send[PDFActionNamedName](p_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// Returns the name of the named action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionNamed/name
func (p_ PDFActionNamed) SetName(value PDFActionNamedName) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFActionNamed */


