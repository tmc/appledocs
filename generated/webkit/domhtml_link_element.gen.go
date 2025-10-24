// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLLinkElement */

/* debug [class_header]: Header for DOMHTMLLinkElement */
// The class instance for the [DOMHTMLLinkElement] class.
var (
	DOMHTMLLinkElementClass     _DOMHTMLLinkElementClass
	DOMHTMLLinkElementClassOnce sync.Once
)

func getDOMHTMLLinkElementClass() _DOMHTMLLinkElementClass {
	DOMHTMLLinkElementClassOnce.Do(func() {
		DOMHTMLLinkElementClass = _DOMHTMLLinkElementClass{objc.GetClass("DOMHTMLLinkElement")}
	})
	return DOMHTMLLinkElementClass
}

type _DOMHTMLLinkElementClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMHTMLLinkElement */
// An interface definition for the [DOMHTMLLinkElement] class.
type IDOMHTMLLinkElement interface {
	IDOMHTMLElement

	/* debug [class_interface_properties]: Properties for DOMHTMLLinkElement */
	// properties:
	AbsoluteLinkURL() objc.IObject /* cross-framework: NSURL */
	Charset() objc.IObject         /* cross-framework: NSString */
	SetCharset(value objc.IObject /* cross-framework: NSString */)
	Disabled() bool
	SetDisabled(value bool)
	Href() objc.IObject /* cross-framework: NSString */
	SetHref(value objc.IObject /* cross-framework: NSString */)
	Hreflang() objc.IObject /* cross-framework: NSString */
	SetHreflang(value objc.IObject /* cross-framework: NSString */)
	Media() objc.IObject /* cross-framework: NSString */
	SetMedia(value objc.IObject /* cross-framework: NSString */)
	Rel() objc.IObject /* cross-framework: NSString */
	SetRel(value objc.IObject /* cross-framework: NSString */)
	Rev() objc.IObject /* cross-framework: NSString */
	SetRev(value objc.IObject /* cross-framework: NSString */)
	Sheet() IDOMStyleSheet
	Target() objc.IObject /* cross-framework: NSString */
	SetTarget(value objc.IObject /* cross-framework: NSString */)
	Type() objc.IObject /* cross-framework: NSString */
	SetType(value objc.IObject /* cross-framework: NSString */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMHTMLLinkElement */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMHTMLLinkElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLLinkElementClass) Alloc() DOMHTMLLinkElement {
	rv := objc.Send[DOMHTMLLinkElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLLinkElementClass) New() DOMHTMLLinkElement {
	rv := objc.Send[DOMHTMLLinkElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLLinkElement) Init() DOMHTMLLinkElement {
	rv := objc.Send[DOMHTMLLinkElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLLinkElement) Autorelease() DOMHTMLLinkElement {
	rv := objc.Send[DOMHTMLLinkElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLLinkElement creates a new DOMHTMLLinkElement instance.
func NewDOMHTMLLinkElement() DOMHTMLLinkElement {
	return getDOMHTMLLinkElementClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMHTMLLinkElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLinkElement
type DOMHTMLLinkElement struct {
	DOMHTMLElement
}

// DOMHTMLLinkElementFrom constructs a [DOMHTMLLinkElement] from an unsafe.Pointer.
func DOMHTMLLinkElementFrom(ptr unsafe.Pointer) DOMHTMLLinkElement {
	return DOMHTMLLinkElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMHTMLLinkElement */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMHTMLLinkElement */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMHTMLLinkElement */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMHTMLLinkElement */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMHTMLLinkElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLinkElement/absoluteLinkURL
func (d_ DOMHTMLLinkElement) AbsoluteLinkURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](d_.ID, objc.Sel("absoluteLinkURL"))
	return rv
} /* debug [instance_properties/getter]: absoluteLinkURL */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLinkElement/charset
func (d_ DOMHTMLLinkElement) Charset() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("charset"))
	return rv
} /* debug [instance_properties/getter]: charset */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLinkElement/charset
func (d_ DOMHTMLLinkElement) SetCharset(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCharset:"), value)
} /* debug [instance_properties/setter]: charset */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLinkElement/disabled
func (d_ DOMHTMLLinkElement) Disabled() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("disabled"))
	return rv
} /* debug [instance_properties/getter]: disabled */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLinkElement/disabled
func (d_ DOMHTMLLinkElement) SetDisabled(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDisabled:"), value)
} /* debug [instance_properties/setter]: disabled */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLinkElement/href
func (d_ DOMHTMLLinkElement) Href() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("href"))
	return rv
} /* debug [instance_properties/getter]: href */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLinkElement/href
func (d_ DOMHTMLLinkElement) SetHref(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHref:"), value)
} /* debug [instance_properties/setter]: href */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLinkElement/hreflang
func (d_ DOMHTMLLinkElement) Hreflang() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("hreflang"))
	return rv
} /* debug [instance_properties/getter]: hreflang */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLinkElement/hreflang
func (d_ DOMHTMLLinkElement) SetHreflang(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHreflang:"), value)
} /* debug [instance_properties/setter]: hreflang */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLinkElement/media
func (d_ DOMHTMLLinkElement) Media() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("media"))
	return rv
} /* debug [instance_properties/getter]: media */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLinkElement/media
func (d_ DOMHTMLLinkElement) SetMedia(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMedia:"), value)
} /* debug [instance_properties/setter]: media */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLinkElement/rel
func (d_ DOMHTMLLinkElement) Rel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("rel"))
	return rv
} /* debug [instance_properties/getter]: rel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLinkElement/rel
func (d_ DOMHTMLLinkElement) SetRel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setRel:"), value)
} /* debug [instance_properties/setter]: rel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLinkElement/rev
func (d_ DOMHTMLLinkElement) Rev() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("rev"))
	return rv
} /* debug [instance_properties/getter]: rev */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLinkElement/rev
func (d_ DOMHTMLLinkElement) SetRev(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setRev:"), value)
} /* debug [instance_properties/setter]: rev */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLinkElement/sheet
func (d_ DOMHTMLLinkElement) Sheet() IDOMStyleSheet {
	rv := objc.Send[DOMStyleSheet](d_.ID, objc.Sel("sheet"))
	return rv
} /* debug [instance_properties/getter]: sheet */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLinkElement/target
func (d_ DOMHTMLLinkElement) Target() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("target"))
	return rv
} /* debug [instance_properties/getter]: target */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLinkElement/target
func (d_ DOMHTMLLinkElement) SetTarget(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTarget:"), value)
} /* debug [instance_properties/setter]: target */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLinkElement/type
func (d_ DOMHTMLLinkElement) Type() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("type"))
	return rv
} /* debug [instance_properties/getter]: type */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLLinkElement/type
func (d_ DOMHTMLLinkElement) SetType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setType:"), value)
} /* debug [instance_properties/setter]: type */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMHTMLLinkElement */
