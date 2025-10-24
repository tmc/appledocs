// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMHTMLLabelElement */


/* debug [class_header]: Header for DOMHTMLLabelElement */
// The class instance for the [DOMHTMLLabelElement] class.
var (
	DOMHTMLLabelElementClass     _DOMHTMLLabelElementClass
	DOMHTMLLabelElementClassOnce sync.Once
)

func getDOMHTMLLabelElementClass() _DOMHTMLLabelElementClass {
	DOMHTMLLabelElementClassOnce.Do(func() {
		DOMHTMLLabelElementClass = _DOMHTMLLabelElementClass{objc.GetClass("DOMHTMLLabelElement")}
	})
	return DOMHTMLLabelElementClass
}

type _DOMHTMLLabelElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLLabelElement */
// An interface definition for the [DOMHTMLLabelElement] class.
type IDOMHTMLLabelElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLLabelElement */
	// properties:
	AccessKey() objc.IObject /* cross-framework: NSString */
	SetAccessKey(value objc.IObject /* cross-framework: NSString */)
	Form() IDOMHTMLFormElement
	HtmlFor() objc.IObject /* cross-framework: NSString */
	SetHtmlFor(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLLabelElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLLabelElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLLabelElementClass) Alloc() DOMHTMLLabelElement {
	rv := objc.Send[DOMHTMLLabelElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLLabelElementClass) New() DOMHTMLLabelElement {
	rv := objc.Send[DOMHTMLLabelElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLLabelElement) Init() DOMHTMLLabelElement {
	rv := objc.Send[DOMHTMLLabelElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLLabelElement) Autorelease() DOMHTMLLabelElement {
	rv := objc.Send[DOMHTMLLabelElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLLabelElement creates a new DOMHTMLLabelElement instance.
func NewDOMHTMLLabelElement() DOMHTMLLabelElement {
	return getDOMHTMLLabelElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLLabelElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLabelElement
type DOMHTMLLabelElement struct {
	DOMHTMLElement
}

// DOMHTMLLabelElementFrom constructs a [DOMHTMLLabelElement] from an unsafe.Pointer.
func DOMHTMLLabelElementFrom(ptr unsafe.Pointer) DOMHTMLLabelElement {
	return DOMHTMLLabelElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLLabelElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLLabelElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLLabelElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLLabelElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLLabelElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLabelElement/accessKey
func (d_ DOMHTMLLabelElement) AccessKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("accessKey"))
	return rv
}/* debug [instance_properties/getter]: accessKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLabelElement/accessKey
func (d_ DOMHTMLLabelElement) SetAccessKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAccessKey:"), value)
}/* debug [instance_properties/setter]: accessKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLabelElement/form
func (d_ DOMHTMLLabelElement) Form() IDOMHTMLFormElement {
	rv := objc.Send[DOMHTMLFormElement](d_.ID, objc.Sel("form"))
	return rv
}/* debug [instance_properties/getter]: form */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLabelElement/htmlFor
func (d_ DOMHTMLLabelElement) HtmlFor() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("htmlFor"))
	return rv
}/* debug [instance_properties/getter]: htmlFor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLabelElement/htmlFor
func (d_ DOMHTMLLabelElement) SetHtmlFor(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHtmlFor:"), value)
}/* debug [instance_properties/setter]: htmlFor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLLabelElement */



