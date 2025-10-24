// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLDListElement */


/* debug [class_header]: Header for DOMHTMLDListElement */
// The class instance for the [DOMHTMLDListElement] class.
var (
	DOMHTMLDListElementClass     _DOMHTMLDListElementClass
	DOMHTMLDListElementClassOnce sync.Once
)

func getDOMHTMLDListElementClass() _DOMHTMLDListElementClass {
	DOMHTMLDListElementClassOnce.Do(func() {
		DOMHTMLDListElementClass = _DOMHTMLDListElementClass{objc.GetClass("DOMHTMLDListElement")}
	})
	return DOMHTMLDListElementClass
}

type _DOMHTMLDListElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLDListElement */
// An interface definition for the [DOMHTMLDListElement] class.
type IDOMHTMLDListElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLDListElement */
	// properties:
	Compact() bool
	SetCompact(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLDListElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLDListElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLDListElementClass) Alloc() DOMHTMLDListElement {
	rv := objc.Send[DOMHTMLDListElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLDListElementClass) New() DOMHTMLDListElement {
	rv := objc.Send[DOMHTMLDListElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLDListElement) Init() DOMHTMLDListElement {
	rv := objc.Send[DOMHTMLDListElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLDListElement) Autorelease() DOMHTMLDListElement {
	rv := objc.Send[DOMHTMLDListElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLDListElement creates a new DOMHTMLDListElement instance.
func NewDOMHTMLDListElement() DOMHTMLDListElement {
	return getDOMHTMLDListElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLDListElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDListElement
type DOMHTMLDListElement struct {
	DOMHTMLElement
}

// DOMHTMLDListElementFrom constructs a [DOMHTMLDListElement] from an unsafe.Pointer.
func DOMHTMLDListElementFrom(ptr unsafe.Pointer) DOMHTMLDListElement {
	return DOMHTMLDListElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLDListElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLDListElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLDListElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLDListElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLDListElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDListElement/compact
func (d_ DOMHTMLDListElement) Compact() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("compact"))
	return rv
}/* debug [instance_properties/getter]: compact */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLDListElement/compact
func (d_ DOMHTMLDListElement) SetCompact(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCompact:"), value)
}/* debug [instance_properties/setter]: compact */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLDListElement */



