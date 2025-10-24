// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLParagraphElement */

/* debug [class_header]: Header for DOMHTMLParagraphElement */
// The class instance for the [DOMHTMLParagraphElement] class.
var (
	DOMHTMLParagraphElementClass     _DOMHTMLParagraphElementClass
	DOMHTMLParagraphElementClassOnce sync.Once
)

func getDOMHTMLParagraphElementClass() _DOMHTMLParagraphElementClass {
	DOMHTMLParagraphElementClassOnce.Do(func() {
		DOMHTMLParagraphElementClass = _DOMHTMLParagraphElementClass{objc.GetClass("DOMHTMLParagraphElement")}
	})
	return DOMHTMLParagraphElementClass
}

type _DOMHTMLParagraphElementClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMHTMLParagraphElement */
// An interface definition for the [DOMHTMLParagraphElement] class.
type IDOMHTMLParagraphElement interface {
	IDOMHTMLElement

	/* debug [class_interface_properties]: Properties for DOMHTMLParagraphElement */
	// properties:
	Align() objc.IObject /* cross-framework: NSString */
	SetAlign(value objc.IObject /* cross-framework: NSString */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMHTMLParagraphElement */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMHTMLParagraphElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLParagraphElementClass) Alloc() DOMHTMLParagraphElement {
	rv := objc.Send[DOMHTMLParagraphElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLParagraphElementClass) New() DOMHTMLParagraphElement {
	rv := objc.Send[DOMHTMLParagraphElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLParagraphElement) Init() DOMHTMLParagraphElement {
	rv := objc.Send[DOMHTMLParagraphElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLParagraphElement) Autorelease() DOMHTMLParagraphElement {
	rv := objc.Send[DOMHTMLParagraphElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLParagraphElement creates a new DOMHTMLParagraphElement instance.
func NewDOMHTMLParagraphElement() DOMHTMLParagraphElement {
	return getDOMHTMLParagraphElementClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMHTMLParagraphElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLParagraphElement
type DOMHTMLParagraphElement struct {
	DOMHTMLElement
}

// DOMHTMLParagraphElementFrom constructs a [DOMHTMLParagraphElement] from an unsafe.Pointer.
func DOMHTMLParagraphElementFrom(ptr unsafe.Pointer) DOMHTMLParagraphElement {
	return DOMHTMLParagraphElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMHTMLParagraphElement */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMHTMLParagraphElement */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMHTMLParagraphElement */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMHTMLParagraphElement */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMHTMLParagraphElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLParagraphElement/align
func (d_ DOMHTMLParagraphElement) Align() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("align"))
	return rv
} /* debug [instance_properties/getter]: align */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLParagraphElement/align
func (d_ DOMHTMLParagraphElement) SetAlign(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAlign:"), value)
} /* debug [instance_properties/setter]: align */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMHTMLParagraphElement */
