// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMAttr */

/* debug [class_header]: Header for DOMAttr */
// The class instance for the [DOMAttr] class.
var (
	DOMAttrClass     _DOMAttrClass
	DOMAttrClassOnce sync.Once
)

func getDOMAttrClass() _DOMAttrClass {
	DOMAttrClassOnce.Do(func() {
		DOMAttrClass = _DOMAttrClass{objc.GetClass("DOMAttr")}
	})
	return DOMAttrClass
}

type _DOMAttrClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMAttr */
// An interface definition for the [DOMAttr] class.
type IDOMAttr interface {
	IDOMNode

	/* debug [class_interface_properties]: Properties for DOMAttr */
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
	OwnerElement() IDOMElement
	Specified() bool
	Style() IDOMCSSStyleDeclaration
	Value() objc.IObject /* cross-framework: NSString */
	SetValue(value objc.IObject /* cross-framework: NSString */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMAttr */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMAttr */
// Alloc allocates a new instance without initialization.
func (dc _DOMAttrClass) Alloc() DOMAttr {
	rv := objc.Send[DOMAttr](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMAttrClass) New() DOMAttr {
	rv := objc.Send[DOMAttr](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMAttr) Init() DOMAttr {
	rv := objc.Send[DOMAttr](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMAttr) Autorelease() DOMAttr {
	rv := objc.Send[DOMAttr](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMAttr creates a new DOMAttr instance.
func NewDOMAttr() DOMAttr {
	return getDOMAttrClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMAttr */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMAttr
type DOMAttr struct {
	DOMNode
}

// DOMAttrFrom constructs a [DOMAttr] from an unsafe.Pointer.
func DOMAttrFrom(ptr unsafe.Pointer) DOMAttr {
	return DOMAttr{
		DOMNode: DOMNodeFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMAttr */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMAttr */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMAttr */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMAttr */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMAttr */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMAttr/name
func (d_ DOMAttr) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("name"))
	return rv
} /* debug [instance_properties/getter]: name */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMAttr/ownerElement
func (d_ DOMAttr) OwnerElement() IDOMElement {
	rv := objc.Send[DOMElement](d_.ID, objc.Sel("ownerElement"))
	return rv
} /* debug [instance_properties/getter]: ownerElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMAttr/specified
func (d_ DOMAttr) Specified() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("specified"))
	return rv
} /* debug [instance_properties/getter]: specified */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMAttr/style
func (d_ DOMAttr) Style() IDOMCSSStyleDeclaration {
	rv := objc.Send[DOMCSSStyleDeclaration](d_.ID, objc.Sel("style"))
	return rv
} /* debug [instance_properties/getter]: style */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMAttr/value
func (d_ DOMAttr) Value() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("value"))
	return rv
} /* debug [instance_properties/getter]: value */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMAttr/value
func (d_ DOMAttr) SetValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setValue:"), value)
} /* debug [instance_properties/setter]: value */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMAttr */
