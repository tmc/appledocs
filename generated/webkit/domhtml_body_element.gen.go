// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLBodyElement */

/* debug [class_header]: Header for DOMHTMLBodyElement */
// The class instance for the [DOMHTMLBodyElement] class.
var (
	DOMHTMLBodyElementClass     _DOMHTMLBodyElementClass
	DOMHTMLBodyElementClassOnce sync.Once
)

func getDOMHTMLBodyElementClass() _DOMHTMLBodyElementClass {
	DOMHTMLBodyElementClassOnce.Do(func() {
		DOMHTMLBodyElementClass = _DOMHTMLBodyElementClass{objc.GetClass("DOMHTMLBodyElement")}
	})
	return DOMHTMLBodyElementClass
}

type _DOMHTMLBodyElementClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMHTMLBodyElement */
// An interface definition for the [DOMHTMLBodyElement] class.
type IDOMHTMLBodyElement interface {
	IDOMHTMLElement

	/* debug [class_interface_properties]: Properties for DOMHTMLBodyElement */
	// properties:
	ALink() objc.IObject /* cross-framework: NSString */
	SetALink(value objc.IObject /* cross-framework: NSString */)
	Background() objc.IObject /* cross-framework: NSString */
	SetBackground(value objc.IObject /* cross-framework: NSString */)
	BgColor() objc.IObject /* cross-framework: NSString */
	SetBgColor(value objc.IObject /* cross-framework: NSString */)
	Link() objc.IObject /* cross-framework: NSString */
	SetLink(value objc.IObject /* cross-framework: NSString */)
	Text() objc.IObject /* cross-framework: NSString */
	SetText(value objc.IObject /* cross-framework: NSString */)
	VLink() objc.IObject /* cross-framework: NSString */
	SetVLink(value objc.IObject /* cross-framework: NSString */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMHTMLBodyElement */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMHTMLBodyElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLBodyElementClass) Alloc() DOMHTMLBodyElement {
	rv := objc.Send[DOMHTMLBodyElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLBodyElementClass) New() DOMHTMLBodyElement {
	rv := objc.Send[DOMHTMLBodyElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLBodyElement) Init() DOMHTMLBodyElement {
	rv := objc.Send[DOMHTMLBodyElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLBodyElement) Autorelease() DOMHTMLBodyElement {
	rv := objc.Send[DOMHTMLBodyElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLBodyElement creates a new DOMHTMLBodyElement instance.
func NewDOMHTMLBodyElement() DOMHTMLBodyElement {
	return getDOMHTMLBodyElementClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMHTMLBodyElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBodyElement
type DOMHTMLBodyElement struct {
	DOMHTMLElement
}

// DOMHTMLBodyElementFrom constructs a [DOMHTMLBodyElement] from an unsafe.Pointer.
func DOMHTMLBodyElementFrom(ptr unsafe.Pointer) DOMHTMLBodyElement {
	return DOMHTMLBodyElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMHTMLBodyElement */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMHTMLBodyElement */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMHTMLBodyElement */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMHTMLBodyElement */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMHTMLBodyElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBodyElement/aLink
func (d_ DOMHTMLBodyElement) ALink() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("aLink"))
	return rv
} /* debug [instance_properties/getter]: aLink */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBodyElement/aLink
func (d_ DOMHTMLBodyElement) SetALink(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setALink:"), value)
} /* debug [instance_properties/setter]: aLink */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBodyElement/background
func (d_ DOMHTMLBodyElement) Background() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("background"))
	return rv
} /* debug [instance_properties/getter]: background */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBodyElement/background
func (d_ DOMHTMLBodyElement) SetBackground(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBackground:"), value)
} /* debug [instance_properties/setter]: background */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBodyElement/bgColor
func (d_ DOMHTMLBodyElement) BgColor() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("bgColor"))
	return rv
} /* debug [instance_properties/getter]: bgColor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBodyElement/bgColor
func (d_ DOMHTMLBodyElement) SetBgColor(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBgColor:"), value)
} /* debug [instance_properties/setter]: bgColor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBodyElement/link
func (d_ DOMHTMLBodyElement) Link() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("link"))
	return rv
} /* debug [instance_properties/getter]: link */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBodyElement/link
func (d_ DOMHTMLBodyElement) SetLink(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLink:"), value)
} /* debug [instance_properties/setter]: link */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBodyElement/text
func (d_ DOMHTMLBodyElement) Text() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("text"))
	return rv
} /* debug [instance_properties/getter]: text */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBodyElement/text
func (d_ DOMHTMLBodyElement) SetText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setText:"), value)
} /* debug [instance_properties/setter]: text */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBodyElement/vLink
func (d_ DOMHTMLBodyElement) VLink() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("vLink"))
	return rv
} /* debug [instance_properties/getter]: vLink */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBodyElement/vLink
func (d_ DOMHTMLBodyElement) SetVLink(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setVLink:"), value)
} /* debug [instance_properties/setter]: vLink */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMHTMLBodyElement */
