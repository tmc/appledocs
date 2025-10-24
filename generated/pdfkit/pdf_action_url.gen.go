// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PDFActionURL */


/* debug [class_header]: Header for PDFActionURL */
// The class instance for the [PDFActionURL] class.
var (
	PDFActionURLClass     _PDFActionURLClass
	PDFActionURLClassOnce sync.Once
)

func getPDFActionURLClass() _PDFActionURLClass {
	PDFActionURLClassOnce.Do(func() {
		PDFActionURLClass = _PDFActionURLClass{objc.GetClass("PDFActionURL")}
	})
	return PDFActionURLClass
}

type _PDFActionURLClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFActionURL */
// An interface definition for the [PDFActionURL] class.
type IPDFActionURL interface {
	IPDFAction
	
/* debug [class_interface_properties]: Properties for PDFActionURL */
	// properties:
	URL() objc.IObject /* cross-framework: NSURL */
	SetURL(value objc.IObject /* cross-framework: NSURL */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFActionURL */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFActionURL */
// Alloc allocates a new instance without initialization.
func (pc _PDFActionURLClass) Alloc() PDFActionURL {
	rv := objc.Send[PDFActionURL](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PDFActionURLClass) New() PDFActionURL {
	rv := objc.Send[PDFActionURL](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFActionURL) Init() PDFActionURL {
	rv := objc.Send[PDFActionURL](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFActionURL) Autorelease() PDFActionURL {
	rv := objc.Send[PDFActionURL](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFActionURL creates a new PDFActionURL instance.
func NewPDFActionURL() PDFActionURL {
	return getPDFActionURLClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFActionURL */
// , a subclass of , defines methods for getting and setting the URL associated with a URL action.


// , a subclass of , defines methods for getting and setting the URL associated with a URL action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionURL
type PDFActionURL struct {
	PDFAction
}

// PDFActionURLFrom constructs a [PDFActionURL] from an unsafe.Pointer.
//
// , a subclass of , defines methods for getting and setting the URL associated with a URL action.
func PDFActionURLFrom(ptr unsafe.Pointer) PDFActionURL {
	return PDFActionURL{
		PDFAction: PDFActionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFActionURL */

// Initializes a URL action with the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionURL/init(url:)
func NewPDFActionURLWithURL(url objc.IObject /* cross-framework: NSURL */) PDFActionURL {
	instance := getPDFActionURLClass().Alloc()
	rv := objc.Send[PDFActionURL](instance.ID, objc.Sel("initWithURL:"), url)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPDFActionURLWithURL */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFActionURL */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFActionURL */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFActionURL */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFActionURL */

// Returns the URL associated with the URL action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionURL/url
func (p_ PDFActionURL) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](p_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_properties/getter]: URL */


// Returns the URL associated with the URL action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionURL/url
func (p_ PDFActionURL) SetURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setURL:"), value)
}/* debug [instance_properties/setter]: URL */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFActionURL */


