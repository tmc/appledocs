// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMHTMLQuoteElement */


/* debug [class_header]: Header for DOMHTMLQuoteElement */
// The class instance for the [DOMHTMLQuoteElement] class.
var (
	DOMHTMLQuoteElementClass     _DOMHTMLQuoteElementClass
	DOMHTMLQuoteElementClassOnce sync.Once
)

func getDOMHTMLQuoteElementClass() _DOMHTMLQuoteElementClass {
	DOMHTMLQuoteElementClassOnce.Do(func() {
		DOMHTMLQuoteElementClass = _DOMHTMLQuoteElementClass{objc.GetClass("DOMHTMLQuoteElement")}
	})
	return DOMHTMLQuoteElementClass
}

type _DOMHTMLQuoteElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLQuoteElement */
// An interface definition for the [DOMHTMLQuoteElement] class.
type IDOMHTMLQuoteElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLQuoteElement */
	// properties:
	Cite() objc.IObject /* cross-framework: NSString */
	SetCite(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLQuoteElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLQuoteElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLQuoteElementClass) Alloc() DOMHTMLQuoteElement {
	rv := objc.Send[DOMHTMLQuoteElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLQuoteElementClass) New() DOMHTMLQuoteElement {
	rv := objc.Send[DOMHTMLQuoteElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLQuoteElement) Init() DOMHTMLQuoteElement {
	rv := objc.Send[DOMHTMLQuoteElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLQuoteElement) Autorelease() DOMHTMLQuoteElement {
	rv := objc.Send[DOMHTMLQuoteElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLQuoteElement creates a new DOMHTMLQuoteElement instance.
func NewDOMHTMLQuoteElement() DOMHTMLQuoteElement {
	return getDOMHTMLQuoteElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLQuoteElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLQuoteElement
type DOMHTMLQuoteElement struct {
	DOMHTMLElement
}

// DOMHTMLQuoteElementFrom constructs a [DOMHTMLQuoteElement] from an unsafe.Pointer.
func DOMHTMLQuoteElementFrom(ptr unsafe.Pointer) DOMHTMLQuoteElement {
	return DOMHTMLQuoteElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLQuoteElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLQuoteElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLQuoteElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLQuoteElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLQuoteElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLQuoteElement/cite
func (d_ DOMHTMLQuoteElement) Cite() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("cite"))
	return rv
}/* debug [instance_properties/getter]: cite */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLQuoteElement/cite
func (d_ DOMHTMLQuoteElement) SetCite(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCite:"), value)
}/* debug [instance_properties/setter]: cite */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLQuoteElement */



