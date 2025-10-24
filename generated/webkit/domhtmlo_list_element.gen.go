// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMHTMLOListElement */


/* debug [class_header]: Header for DOMHTMLOListElement */
// The class instance for the [DOMHTMLOListElement] class.
var (
	DOMHTMLOListElementClass     _DOMHTMLOListElementClass
	DOMHTMLOListElementClassOnce sync.Once
)

func getDOMHTMLOListElementClass() _DOMHTMLOListElementClass {
	DOMHTMLOListElementClassOnce.Do(func() {
		DOMHTMLOListElementClass = _DOMHTMLOListElementClass{objc.GetClass("DOMHTMLOListElement")}
	})
	return DOMHTMLOListElementClass
}

type _DOMHTMLOListElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLOListElement */
// An interface definition for the [DOMHTMLOListElement] class.
type IDOMHTMLOListElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLOListElement */
	// properties:
	Compact() bool
	SetCompact(value bool)
	Start() int
	SetStart(value int)
	Type() objc.IObject /* cross-framework: NSString */
	SetType(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLOListElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLOListElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLOListElementClass) Alloc() DOMHTMLOListElement {
	rv := objc.Send[DOMHTMLOListElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLOListElementClass) New() DOMHTMLOListElement {
	rv := objc.Send[DOMHTMLOListElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLOListElement) Init() DOMHTMLOListElement {
	rv := objc.Send[DOMHTMLOListElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLOListElement) Autorelease() DOMHTMLOListElement {
	rv := objc.Send[DOMHTMLOListElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLOListElement creates a new DOMHTMLOListElement instance.
func NewDOMHTMLOListElement() DOMHTMLOListElement {
	return getDOMHTMLOListElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLOListElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOListElement
type DOMHTMLOListElement struct {
	DOMHTMLElement
}

// DOMHTMLOListElementFrom constructs a [DOMHTMLOListElement] from an unsafe.Pointer.
func DOMHTMLOListElementFrom(ptr unsafe.Pointer) DOMHTMLOListElement {
	return DOMHTMLOListElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLOListElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLOListElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLOListElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLOListElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLOListElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOListElement/compact
func (d_ DOMHTMLOListElement) Compact() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("compact"))
	return rv
}/* debug [instance_properties/getter]: compact */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOListElement/compact
func (d_ DOMHTMLOListElement) SetCompact(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCompact:"), value)
}/* debug [instance_properties/setter]: compact */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOListElement/start
func (d_ DOMHTMLOListElement) Start() int {
	rv := objc.Send[int](d_.ID, objc.Sel("start"))
	return rv
}/* debug [instance_properties/getter]: start */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOListElement/start
func (d_ DOMHTMLOListElement) SetStart(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setStart:"), value)
}/* debug [instance_properties/setter]: start */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOListElement/type
func (d_ DOMHTMLOListElement) Type() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOListElement/type
func (d_ DOMHTMLOListElement) SetType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setType:"), value)
}/* debug [instance_properties/setter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLOListElement */



