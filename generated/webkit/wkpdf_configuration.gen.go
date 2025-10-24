// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKPDFConfiguration */


/* debug [class_header]: Header for WKPDFConfiguration */
// The class instance for the [PDFConfiguration] class.
var (
	PDFConfigurationClass     _PDFConfigurationClass
	PDFConfigurationClassOnce sync.Once
)

func getPDFConfigurationClass() _PDFConfigurationClass {
	PDFConfigurationClassOnce.Do(func() {
		PDFConfigurationClass = _PDFConfigurationClass{objc.GetClass("WKPDFConfiguration")}
	})
	return PDFConfigurationClass
}

type _PDFConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFConfiguration */
// An interface definition for the [PDFConfiguration] class.
type IPDFConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PDFConfiguration */
	// properties:
	AllowTransparentBackground() bool
	SetAllowTransparentBackground(value bool)
	Rect() corefoundation.CGRect
	SetRect(value corefoundation.CGRect)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFConfiguration */
// Alloc allocates a new instance without initialization.
func (pc _PDFConfigurationClass) Alloc() PDFConfiguration {
	rv := objc.Send[PDFConfiguration](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PDFConfigurationClass) New() PDFConfiguration {
	rv := objc.Send[PDFConfiguration](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFConfiguration) Init() PDFConfiguration {
	rv := objc.Send[PDFConfiguration](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFConfiguration) Autorelease() PDFConfiguration {
	rv := objc.Send[PDFConfiguration](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFConfiguration creates a new PDFConfiguration instance.
func NewPDFConfiguration() PDFConfiguration {
	return getPDFConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFConfiguration */
// The configuration data to use when generating a PDF representation of a web view’s contents.
//
// Create a object when you want to generate a PDF version of your web view’s content. Use this object to specify the portion of the web view to capture. To generate the PDF content, pass the configuration object to the method of , which returns the PDF data for you to use.


// The configuration data to use when generating a PDF representation of a web view’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPDFConfiguration
type PDFConfiguration struct {
	objectivec.Object
}

// PDFConfigurationFrom constructs a [PDFConfiguration] from an unsafe.Pointer.
//
// The configuration data to use when generating a PDF representation of a web view’s contents.
func PDFConfigurationFrom(ptr unsafe.Pointer) PDFConfiguration {
	return PDFConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFConfiguration */

// A Boolean value that indicates whether the PDF may have a transparent background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPDFConfiguration/allowTransparentBackground
func (p_ PDFConfiguration) AllowTransparentBackground() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowTransparentBackground"))
	return rv
}/* debug [instance_properties/getter]: allowTransparentBackground */


// A Boolean value that indicates whether the PDF may have a transparent background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPDFConfiguration/allowTransparentBackground
func (p_ PDFConfiguration) SetAllowTransparentBackground(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowTransparentBackground:"), value)
}/* debug [instance_properties/setter]: allowTransparentBackground */


// The portion of your web view to capture, specified as a rectangle in the view’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPDFConfiguration/rect-3xww9
func (p_ PDFConfiguration) Rect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p_.ID, objc.Sel("rect"))
	return rv
}/* debug [instance_properties/getter]: rect */


// The portion of your web view to capture, specified as a rectangle in the view’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPDFConfiguration/rect-3xww9
func (p_ PDFConfiguration) SetRect(value corefoundation.CGRect) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRect:"), value)
}/* debug [instance_properties/setter]: rect */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WKPDFConfiguration */



