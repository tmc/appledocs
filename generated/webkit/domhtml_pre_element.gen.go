// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLPreElement */


/* debug [class_header]: Header for DOMHTMLPreElement */
// The class instance for the [DOMHTMLPreElement] class.
var (
	DOMHTMLPreElementClass     _DOMHTMLPreElementClass
	DOMHTMLPreElementClassOnce sync.Once
)

func getDOMHTMLPreElementClass() _DOMHTMLPreElementClass {
	DOMHTMLPreElementClassOnce.Do(func() {
		DOMHTMLPreElementClass = _DOMHTMLPreElementClass{objc.GetClass("DOMHTMLPreElement")}
	})
	return DOMHTMLPreElementClass
}

type _DOMHTMLPreElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLPreElement */
// An interface definition for the [DOMHTMLPreElement] class.
type IDOMHTMLPreElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLPreElement */
	// properties:
	Width() int
	SetWidth(value int)
	Wrap() bool
	SetWrap(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLPreElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLPreElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLPreElementClass) Alloc() DOMHTMLPreElement {
	rv := objc.Send[DOMHTMLPreElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLPreElementClass) New() DOMHTMLPreElement {
	rv := objc.Send[DOMHTMLPreElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLPreElement) Init() DOMHTMLPreElement {
	rv := objc.Send[DOMHTMLPreElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLPreElement) Autorelease() DOMHTMLPreElement {
	rv := objc.Send[DOMHTMLPreElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLPreElement creates a new DOMHTMLPreElement instance.
func NewDOMHTMLPreElement() DOMHTMLPreElement {
	return getDOMHTMLPreElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLPreElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLPreElement
type DOMHTMLPreElement struct {
	DOMHTMLElement
}

// DOMHTMLPreElementFrom constructs a [DOMHTMLPreElement] from an unsafe.Pointer.
func DOMHTMLPreElementFrom(ptr unsafe.Pointer) DOMHTMLPreElement {
	return DOMHTMLPreElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLPreElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLPreElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLPreElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLPreElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLPreElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLPreElement/width
func (d_ DOMHTMLPreElement) Width() int {
	rv := objc.Send[int](d_.ID, objc.Sel("width"))
	return rv
}/* debug [instance_properties/getter]: width */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLPreElement/width
func (d_ DOMHTMLPreElement) SetWidth(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWidth:"), value)
}/* debug [instance_properties/setter]: width */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLPreElement/wrap
func (d_ DOMHTMLPreElement) Wrap() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("wrap"))
	return rv
}/* debug [instance_properties/getter]: wrap */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLPreElement/wrap
func (d_ DOMHTMLPreElement) SetWrap(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWrap:"), value)
}/* debug [instance_properties/setter]: wrap */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLPreElement */



