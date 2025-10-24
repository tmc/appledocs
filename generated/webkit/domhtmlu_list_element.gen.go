// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLUListElement */

/* debug [class_header]: Header for DOMHTMLUListElement */
// The class instance for the [DOMHTMLUListElement] class.
var (
	DOMHTMLUListElementClass     _DOMHTMLUListElementClass
	DOMHTMLUListElementClassOnce sync.Once
)

func getDOMHTMLUListElementClass() _DOMHTMLUListElementClass {
	DOMHTMLUListElementClassOnce.Do(func() {
		DOMHTMLUListElementClass = _DOMHTMLUListElementClass{objc.GetClass("DOMHTMLUListElement")}
	})
	return DOMHTMLUListElementClass
}

type _DOMHTMLUListElementClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMHTMLUListElement */
// An interface definition for the [DOMHTMLUListElement] class.
type IDOMHTMLUListElement interface {
	IDOMHTMLElement

	/* debug [class_interface_properties]: Properties for DOMHTMLUListElement */
	// properties:
	Compact() bool
	SetCompact(value bool)
	Type() objc.IObject /* cross-framework: NSString */
	SetType(value objc.IObject /* cross-framework: NSString */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMHTMLUListElement */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMHTMLUListElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLUListElementClass) Alloc() DOMHTMLUListElement {
	rv := objc.Send[DOMHTMLUListElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLUListElementClass) New() DOMHTMLUListElement {
	rv := objc.Send[DOMHTMLUListElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLUListElement) Init() DOMHTMLUListElement {
	rv := objc.Send[DOMHTMLUListElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLUListElement) Autorelease() DOMHTMLUListElement {
	rv := objc.Send[DOMHTMLUListElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLUListElement creates a new DOMHTMLUListElement instance.
func NewDOMHTMLUListElement() DOMHTMLUListElement {
	return getDOMHTMLUListElementClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMHTMLUListElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLUListElement
type DOMHTMLUListElement struct {
	DOMHTMLElement
}

// DOMHTMLUListElementFrom constructs a [DOMHTMLUListElement] from an unsafe.Pointer.
func DOMHTMLUListElementFrom(ptr unsafe.Pointer) DOMHTMLUListElement {
	return DOMHTMLUListElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMHTMLUListElement */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMHTMLUListElement */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMHTMLUListElement */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMHTMLUListElement */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMHTMLUListElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLUListElement/compact
func (d_ DOMHTMLUListElement) Compact() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("compact"))
	return rv
} /* debug [instance_properties/getter]: compact */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLUListElement/compact
func (d_ DOMHTMLUListElement) SetCompact(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCompact:"), value)
} /* debug [instance_properties/setter]: compact */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLUListElement/type
func (d_ DOMHTMLUListElement) Type() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("type"))
	return rv
} /* debug [instance_properties/getter]: type */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLUListElement/type
func (d_ DOMHTMLUListElement) SetType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setType:"), value)
} /* debug [instance_properties/setter]: type */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMHTMLUListElement */
