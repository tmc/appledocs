// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMHTMLHtmlElement */


/* debug [class_header]: Header for DOMHTMLHtmlElement */
// The class instance for the [DOMHTMLHtmlElement] class.
var (
	DOMHTMLHtmlElementClass     _DOMHTMLHtmlElementClass
	DOMHTMLHtmlElementClassOnce sync.Once
)

func getDOMHTMLHtmlElementClass() _DOMHTMLHtmlElementClass {
	DOMHTMLHtmlElementClassOnce.Do(func() {
		DOMHTMLHtmlElementClass = _DOMHTMLHtmlElementClass{objc.GetClass("DOMHTMLHtmlElement")}
	})
	return DOMHTMLHtmlElementClass
}

type _DOMHTMLHtmlElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLHtmlElement */
// An interface definition for the [DOMHTMLHtmlElement] class.
type IDOMHTMLHtmlElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLHtmlElement */
	// properties:
	Version() objc.IObject /* cross-framework: NSString */
	SetVersion(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLHtmlElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLHtmlElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLHtmlElementClass) Alloc() DOMHTMLHtmlElement {
	rv := objc.Send[DOMHTMLHtmlElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLHtmlElementClass) New() DOMHTMLHtmlElement {
	rv := objc.Send[DOMHTMLHtmlElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLHtmlElement) Init() DOMHTMLHtmlElement {
	rv := objc.Send[DOMHTMLHtmlElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLHtmlElement) Autorelease() DOMHTMLHtmlElement {
	rv := objc.Send[DOMHTMLHtmlElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLHtmlElement creates a new DOMHTMLHtmlElement instance.
func NewDOMHTMLHtmlElement() DOMHTMLHtmlElement {
	return getDOMHTMLHtmlElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLHtmlElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLHtmlElement
type DOMHTMLHtmlElement struct {
	DOMHTMLElement
}

// DOMHTMLHtmlElementFrom constructs a [DOMHTMLHtmlElement] from an unsafe.Pointer.
func DOMHTMLHtmlElementFrom(ptr unsafe.Pointer) DOMHTMLHtmlElement {
	return DOMHTMLHtmlElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLHtmlElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLHtmlElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLHtmlElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLHtmlElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLHtmlElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLHtmlElement/version
func (d_ DOMHTMLHtmlElement) Version() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("version"))
	return rv
}/* debug [instance_properties/getter]: version */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLHtmlElement/version
func (d_ DOMHTMLHtmlElement) SetVersion(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setVersion:"), value)
}/* debug [instance_properties/setter]: version */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLHtmlElement */



