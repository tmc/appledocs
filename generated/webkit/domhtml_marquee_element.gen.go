// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLMarqueeElement */


/* debug [class_header]: Header for DOMHTMLMarqueeElement */
// The class instance for the [DOMHTMLMarqueeElement] class.
var (
	DOMHTMLMarqueeElementClass     _DOMHTMLMarqueeElementClass
	DOMHTMLMarqueeElementClassOnce sync.Once
)

func getDOMHTMLMarqueeElementClass() _DOMHTMLMarqueeElementClass {
	DOMHTMLMarqueeElementClassOnce.Do(func() {
		DOMHTMLMarqueeElementClass = _DOMHTMLMarqueeElementClass{objc.GetClass("DOMHTMLMarqueeElement")}
	})
	return DOMHTMLMarqueeElementClass
}

type _DOMHTMLMarqueeElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLMarqueeElement */
// An interface definition for the [DOMHTMLMarqueeElement] class.
type IDOMHTMLMarqueeElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLMarqueeElement */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLMarqueeElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLMarqueeElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLMarqueeElementClass) Alloc() DOMHTMLMarqueeElement {
	rv := objc.Send[DOMHTMLMarqueeElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLMarqueeElementClass) New() DOMHTMLMarqueeElement {
	rv := objc.Send[DOMHTMLMarqueeElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLMarqueeElement) Init() DOMHTMLMarqueeElement {
	rv := objc.Send[DOMHTMLMarqueeElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLMarqueeElement) Autorelease() DOMHTMLMarqueeElement {
	rv := objc.Send[DOMHTMLMarqueeElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLMarqueeElement creates a new DOMHTMLMarqueeElement instance.
func NewDOMHTMLMarqueeElement() DOMHTMLMarqueeElement {
	return getDOMHTMLMarqueeElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLMarqueeElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLMarqueeElement
type DOMHTMLMarqueeElement struct {
	DOMHTMLElement
}

// DOMHTMLMarqueeElementFrom constructs a [DOMHTMLMarqueeElement] from an unsafe.Pointer.
func DOMHTMLMarqueeElementFrom(ptr unsafe.Pointer) DOMHTMLMarqueeElement {
	return DOMHTMLMarqueeElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLMarqueeElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLMarqueeElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLMarqueeElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLMarqueeElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLMarqueeElement */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLMarqueeElement */



