// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLTitleElement */

/* debug [class_header]: Header for DOMHTMLTitleElement */
// The class instance for the [DOMHTMLTitleElement] class.
var (
	DOMHTMLTitleElementClass     _DOMHTMLTitleElementClass
	DOMHTMLTitleElementClassOnce sync.Once
)

func getDOMHTMLTitleElementClass() _DOMHTMLTitleElementClass {
	DOMHTMLTitleElementClassOnce.Do(func() {
		DOMHTMLTitleElementClass = _DOMHTMLTitleElementClass{objc.GetClass("DOMHTMLTitleElement")}
	})
	return DOMHTMLTitleElementClass
}

type _DOMHTMLTitleElementClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMHTMLTitleElement */
// An interface definition for the [DOMHTMLTitleElement] class.
type IDOMHTMLTitleElement interface {
	IDOMHTMLElement

	/* debug [class_interface_properties]: Properties for DOMHTMLTitleElement */
	// properties:
	Text() objc.IObject /* cross-framework: NSString */
	SetText(value objc.IObject /* cross-framework: NSString */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMHTMLTitleElement */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMHTMLTitleElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLTitleElementClass) Alloc() DOMHTMLTitleElement {
	rv := objc.Send[DOMHTMLTitleElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLTitleElementClass) New() DOMHTMLTitleElement {
	rv := objc.Send[DOMHTMLTitleElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLTitleElement) Init() DOMHTMLTitleElement {
	rv := objc.Send[DOMHTMLTitleElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLTitleElement) Autorelease() DOMHTMLTitleElement {
	rv := objc.Send[DOMHTMLTitleElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLTitleElement creates a new DOMHTMLTitleElement instance.
func NewDOMHTMLTitleElement() DOMHTMLTitleElement {
	return getDOMHTMLTitleElementClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMHTMLTitleElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTitleElement
type DOMHTMLTitleElement struct {
	DOMHTMLElement
}

// DOMHTMLTitleElementFrom constructs a [DOMHTMLTitleElement] from an unsafe.Pointer.
func DOMHTMLTitleElementFrom(ptr unsafe.Pointer) DOMHTMLTitleElement {
	return DOMHTMLTitleElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMHTMLTitleElement */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMHTMLTitleElement */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMHTMLTitleElement */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMHTMLTitleElement */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMHTMLTitleElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTitleElement/text
func (d_ DOMHTMLTitleElement) Text() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("text"))
	return rv
} /* debug [instance_properties/getter]: text */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTitleElement/text
func (d_ DOMHTMLTitleElement) SetText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setText:"), value)
} /* debug [instance_properties/setter]: text */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMHTMLTitleElement */
