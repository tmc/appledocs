// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLScriptElement */

/* debug [class_header]: Header for DOMHTMLScriptElement */
// The class instance for the [DOMHTMLScriptElement] class.
var (
	DOMHTMLScriptElementClass     _DOMHTMLScriptElementClass
	DOMHTMLScriptElementClassOnce sync.Once
)

func getDOMHTMLScriptElementClass() _DOMHTMLScriptElementClass {
	DOMHTMLScriptElementClassOnce.Do(func() {
		DOMHTMLScriptElementClass = _DOMHTMLScriptElementClass{objc.GetClass("DOMHTMLScriptElement")}
	})
	return DOMHTMLScriptElementClass
}

type _DOMHTMLScriptElementClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMHTMLScriptElement */
// An interface definition for the [DOMHTMLScriptElement] class.
type IDOMHTMLScriptElement interface {
	IDOMHTMLElement

	/* debug [class_interface_properties]: Properties for DOMHTMLScriptElement */
	// properties:
	Charset() objc.IObject /* cross-framework: NSString */
	SetCharset(value objc.IObject /* cross-framework: NSString */)
	Defer() bool
	SetDefer(value bool)
	Event() objc.IObject /* cross-framework: NSString */
	SetEvent(value objc.IObject /* cross-framework: NSString */)
	HtmlFor() objc.IObject /* cross-framework: NSString */
	SetHtmlFor(value objc.IObject /* cross-framework: NSString */)
	Src() objc.IObject /* cross-framework: NSString */
	SetSrc(value objc.IObject /* cross-framework: NSString */)
	Text() objc.IObject /* cross-framework: NSString */
	SetText(value objc.IObject /* cross-framework: NSString */)
	Type() objc.IObject /* cross-framework: NSString */
	SetType(value objc.IObject /* cross-framework: NSString */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMHTMLScriptElement */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMHTMLScriptElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLScriptElementClass) Alloc() DOMHTMLScriptElement {
	rv := objc.Send[DOMHTMLScriptElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLScriptElementClass) New() DOMHTMLScriptElement {
	rv := objc.Send[DOMHTMLScriptElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLScriptElement) Init() DOMHTMLScriptElement {
	rv := objc.Send[DOMHTMLScriptElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLScriptElement) Autorelease() DOMHTMLScriptElement {
	rv := objc.Send[DOMHTMLScriptElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLScriptElement creates a new DOMHTMLScriptElement instance.
func NewDOMHTMLScriptElement() DOMHTMLScriptElement {
	return getDOMHTMLScriptElementClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMHTMLScriptElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLScriptElement
type DOMHTMLScriptElement struct {
	DOMHTMLElement
}

// DOMHTMLScriptElementFrom constructs a [DOMHTMLScriptElement] from an unsafe.Pointer.
func DOMHTMLScriptElementFrom(ptr unsafe.Pointer) DOMHTMLScriptElement {
	return DOMHTMLScriptElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMHTMLScriptElement */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMHTMLScriptElement */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMHTMLScriptElement */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMHTMLScriptElement */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMHTMLScriptElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLScriptElement/charset
func (d_ DOMHTMLScriptElement) Charset() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("charset"))
	return rv
} /* debug [instance_properties/getter]: charset */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLScriptElement/charset
func (d_ DOMHTMLScriptElement) SetCharset(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCharset:"), value)
} /* debug [instance_properties/setter]: charset */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLScriptElement/defer
func (d_ DOMHTMLScriptElement) Defer() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("defer"))
	return rv
} /* debug [instance_properties/getter]: defer */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLScriptElement/defer
func (d_ DOMHTMLScriptElement) SetDefer(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDefer:"), value)
} /* debug [instance_properties/setter]: defer */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLScriptElement/event
func (d_ DOMHTMLScriptElement) Event() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("event"))
	return rv
} /* debug [instance_properties/getter]: event */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLScriptElement/event
func (d_ DOMHTMLScriptElement) SetEvent(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setEvent:"), value)
} /* debug [instance_properties/setter]: event */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLScriptElement/htmlFor
func (d_ DOMHTMLScriptElement) HtmlFor() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("htmlFor"))
	return rv
} /* debug [instance_properties/getter]: htmlFor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLScriptElement/htmlFor
func (d_ DOMHTMLScriptElement) SetHtmlFor(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHtmlFor:"), value)
} /* debug [instance_properties/setter]: htmlFor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLScriptElement/src
func (d_ DOMHTMLScriptElement) Src() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("src"))
	return rv
} /* debug [instance_properties/getter]: src */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLScriptElement/src
func (d_ DOMHTMLScriptElement) SetSrc(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSrc:"), value)
} /* debug [instance_properties/setter]: src */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLScriptElement/text
func (d_ DOMHTMLScriptElement) Text() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("text"))
	return rv
} /* debug [instance_properties/getter]: text */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLScriptElement/text
func (d_ DOMHTMLScriptElement) SetText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setText:"), value)
} /* debug [instance_properties/setter]: text */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLScriptElement/type
func (d_ DOMHTMLScriptElement) Type() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("type"))
	return rv
} /* debug [instance_properties/getter]: type */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLScriptElement/type
func (d_ DOMHTMLScriptElement) SetType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setType:"), value)
} /* debug [instance_properties/setter]: type */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMHTMLScriptElement */
