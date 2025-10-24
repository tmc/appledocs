// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLMapElement */

/* debug [class_header]: Header for DOMHTMLMapElement */
// The class instance for the [DOMHTMLMapElement] class.
var (
	DOMHTMLMapElementClass     _DOMHTMLMapElementClass
	DOMHTMLMapElementClassOnce sync.Once
)

func getDOMHTMLMapElementClass() _DOMHTMLMapElementClass {
	DOMHTMLMapElementClassOnce.Do(func() {
		DOMHTMLMapElementClass = _DOMHTMLMapElementClass{objc.GetClass("DOMHTMLMapElement")}
	})
	return DOMHTMLMapElementClass
}

type _DOMHTMLMapElementClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMHTMLMapElement */
// An interface definition for the [DOMHTMLMapElement] class.
type IDOMHTMLMapElement interface {
	IDOMHTMLElement

	/* debug [class_interface_properties]: Properties for DOMHTMLMapElement */
	// properties:
	Areas() IDOMHTMLCollection
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMHTMLMapElement */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMHTMLMapElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLMapElementClass) Alloc() DOMHTMLMapElement {
	rv := objc.Send[DOMHTMLMapElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLMapElementClass) New() DOMHTMLMapElement {
	rv := objc.Send[DOMHTMLMapElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLMapElement) Init() DOMHTMLMapElement {
	rv := objc.Send[DOMHTMLMapElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLMapElement) Autorelease() DOMHTMLMapElement {
	rv := objc.Send[DOMHTMLMapElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLMapElement creates a new DOMHTMLMapElement instance.
func NewDOMHTMLMapElement() DOMHTMLMapElement {
	return getDOMHTMLMapElementClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMHTMLMapElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLMapElement
type DOMHTMLMapElement struct {
	DOMHTMLElement
}

// DOMHTMLMapElementFrom constructs a [DOMHTMLMapElement] from an unsafe.Pointer.
func DOMHTMLMapElementFrom(ptr unsafe.Pointer) DOMHTMLMapElement {
	return DOMHTMLMapElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMHTMLMapElement */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMHTMLMapElement */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMHTMLMapElement */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMHTMLMapElement */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMHTMLMapElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLMapElement/areas
func (d_ DOMHTMLMapElement) Areas() IDOMHTMLCollection {
	rv := objc.Send[DOMHTMLCollection](d_.ID, objc.Sel("areas"))
	return rv
} /* debug [instance_properties/getter]: areas */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLMapElement/name
func (d_ DOMHTMLMapElement) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("name"))
	return rv
} /* debug [instance_properties/getter]: name */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLMapElement/name
func (d_ DOMHTMLMapElement) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setName:"), value)
} /* debug [instance_properties/setter]: name */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMHTMLMapElement */
