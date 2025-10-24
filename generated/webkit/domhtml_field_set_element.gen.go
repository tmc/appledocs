// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLFieldSetElement */


/* debug [class_header]: Header for DOMHTMLFieldSetElement */
// The class instance for the [DOMHTMLFieldSetElement] class.
var (
	DOMHTMLFieldSetElementClass     _DOMHTMLFieldSetElementClass
	DOMHTMLFieldSetElementClassOnce sync.Once
)

func getDOMHTMLFieldSetElementClass() _DOMHTMLFieldSetElementClass {
	DOMHTMLFieldSetElementClassOnce.Do(func() {
		DOMHTMLFieldSetElementClass = _DOMHTMLFieldSetElementClass{objc.GetClass("DOMHTMLFieldSetElement")}
	})
	return DOMHTMLFieldSetElementClass
}

type _DOMHTMLFieldSetElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLFieldSetElement */
// An interface definition for the [DOMHTMLFieldSetElement] class.
type IDOMHTMLFieldSetElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLFieldSetElement */
	// properties:
	Form() IDOMHTMLFormElement
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLFieldSetElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLFieldSetElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLFieldSetElementClass) Alloc() DOMHTMLFieldSetElement {
	rv := objc.Send[DOMHTMLFieldSetElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLFieldSetElementClass) New() DOMHTMLFieldSetElement {
	rv := objc.Send[DOMHTMLFieldSetElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLFieldSetElement) Init() DOMHTMLFieldSetElement {
	rv := objc.Send[DOMHTMLFieldSetElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLFieldSetElement) Autorelease() DOMHTMLFieldSetElement {
	rv := objc.Send[DOMHTMLFieldSetElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLFieldSetElement creates a new DOMHTMLFieldSetElement instance.
func NewDOMHTMLFieldSetElement() DOMHTMLFieldSetElement {
	return getDOMHTMLFieldSetElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLFieldSetElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFieldSetElement
type DOMHTMLFieldSetElement struct {
	DOMHTMLElement
}

// DOMHTMLFieldSetElementFrom constructs a [DOMHTMLFieldSetElement] from an unsafe.Pointer.
func DOMHTMLFieldSetElementFrom(ptr unsafe.Pointer) DOMHTMLFieldSetElement {
	return DOMHTMLFieldSetElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLFieldSetElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLFieldSetElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLFieldSetElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLFieldSetElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLFieldSetElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFieldSetElement/form
func (d_ DOMHTMLFieldSetElement) Form() IDOMHTMLFormElement {
	rv := objc.Send[DOMHTMLFormElement](d_.ID, objc.Sel("form"))
	return rv
}/* debug [instance_properties/getter]: form */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLFieldSetElement */



