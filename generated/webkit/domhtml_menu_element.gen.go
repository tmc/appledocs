// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLMenuElement */


/* debug [class_header]: Header for DOMHTMLMenuElement */
// The class instance for the [DOMHTMLMenuElement] class.
var (
	DOMHTMLMenuElementClass     _DOMHTMLMenuElementClass
	DOMHTMLMenuElementClassOnce sync.Once
)

func getDOMHTMLMenuElementClass() _DOMHTMLMenuElementClass {
	DOMHTMLMenuElementClassOnce.Do(func() {
		DOMHTMLMenuElementClass = _DOMHTMLMenuElementClass{objc.GetClass("DOMHTMLMenuElement")}
	})
	return DOMHTMLMenuElementClass
}

type _DOMHTMLMenuElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLMenuElement */
// An interface definition for the [DOMHTMLMenuElement] class.
type IDOMHTMLMenuElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLMenuElement */
	// properties:
	Compact() bool
	SetCompact(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLMenuElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLMenuElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLMenuElementClass) Alloc() DOMHTMLMenuElement {
	rv := objc.Send[DOMHTMLMenuElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLMenuElementClass) New() DOMHTMLMenuElement {
	rv := objc.Send[DOMHTMLMenuElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLMenuElement) Init() DOMHTMLMenuElement {
	rv := objc.Send[DOMHTMLMenuElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLMenuElement) Autorelease() DOMHTMLMenuElement {
	rv := objc.Send[DOMHTMLMenuElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLMenuElement creates a new DOMHTMLMenuElement instance.
func NewDOMHTMLMenuElement() DOMHTMLMenuElement {
	return getDOMHTMLMenuElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLMenuElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLMenuElement
type DOMHTMLMenuElement struct {
	DOMHTMLElement
}

// DOMHTMLMenuElementFrom constructs a [DOMHTMLMenuElement] from an unsafe.Pointer.
func DOMHTMLMenuElementFrom(ptr unsafe.Pointer) DOMHTMLMenuElement {
	return DOMHTMLMenuElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLMenuElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLMenuElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLMenuElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLMenuElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLMenuElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLMenuElement/compact
func (d_ DOMHTMLMenuElement) Compact() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("compact"))
	return rv
}/* debug [instance_properties/getter]: compact */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLMenuElement/compact
func (d_ DOMHTMLMenuElement) SetCompact(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCompact:"), value)
}/* debug [instance_properties/setter]: compact */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLMenuElement */



