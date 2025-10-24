// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLAreaElement */

/* debug [class_header]: Header for DOMHTMLAreaElement */
// The class instance for the [DOMHTMLAreaElement] class.
var (
	DOMHTMLAreaElementClass     _DOMHTMLAreaElementClass
	DOMHTMLAreaElementClassOnce sync.Once
)

func getDOMHTMLAreaElementClass() _DOMHTMLAreaElementClass {
	DOMHTMLAreaElementClassOnce.Do(func() {
		DOMHTMLAreaElementClass = _DOMHTMLAreaElementClass{objc.GetClass("DOMHTMLAreaElement")}
	})
	return DOMHTMLAreaElementClass
}

type _DOMHTMLAreaElementClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMHTMLAreaElement */
// An interface definition for the [DOMHTMLAreaElement] class.
type IDOMHTMLAreaElement interface {
	IDOMHTMLElement

	/* debug [class_interface_properties]: Properties for DOMHTMLAreaElement */
	// properties:
	AbsoluteLinkURL() objc.IObject /* cross-framework: NSURL */
	Alt() objc.IObject             /* cross-framework: NSString */
	SetAlt(value objc.IObject /* cross-framework: NSString */)
	Coords() objc.IObject /* cross-framework: NSString */
	SetCoords(value objc.IObject /* cross-framework: NSString */)
	HashName() objc.IObject /* cross-framework: NSString */
	Host() objc.IObject     /* cross-framework: NSString */
	Hostname() objc.IObject /* cross-framework: NSString */
	Href() objc.IObject     /* cross-framework: NSString */
	SetHref(value objc.IObject /* cross-framework: NSString */)
	NoHref() bool
	SetNoHref(value bool)
	Pathname() objc.IObject /* cross-framework: NSString */
	Port() objc.IObject     /* cross-framework: NSString */
	Protocol() objc.IObject /* cross-framework: NSString */
	Search() objc.IObject   /* cross-framework: NSString */
	Shape() objc.IObject    /* cross-framework: NSString */
	SetShape(value objc.IObject /* cross-framework: NSString */)
	Target() objc.IObject /* cross-framework: NSString */
	SetTarget(value objc.IObject /* cross-framework: NSString */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMHTMLAreaElement */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMHTMLAreaElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLAreaElementClass) Alloc() DOMHTMLAreaElement {
	rv := objc.Send[DOMHTMLAreaElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLAreaElementClass) New() DOMHTMLAreaElement {
	rv := objc.Send[DOMHTMLAreaElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLAreaElement) Init() DOMHTMLAreaElement {
	rv := objc.Send[DOMHTMLAreaElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLAreaElement) Autorelease() DOMHTMLAreaElement {
	rv := objc.Send[DOMHTMLAreaElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLAreaElement creates a new DOMHTMLAreaElement instance.
func NewDOMHTMLAreaElement() DOMHTMLAreaElement {
	return getDOMHTMLAreaElementClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMHTMLAreaElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAreaElement
type DOMHTMLAreaElement struct {
	DOMHTMLElement
}

// DOMHTMLAreaElementFrom constructs a [DOMHTMLAreaElement] from an unsafe.Pointer.
func DOMHTMLAreaElementFrom(ptr unsafe.Pointer) DOMHTMLAreaElement {
	return DOMHTMLAreaElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMHTMLAreaElement */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMHTMLAreaElement */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMHTMLAreaElement */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMHTMLAreaElement */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMHTMLAreaElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAreaElement/absoluteLinkURL
func (d_ DOMHTMLAreaElement) AbsoluteLinkURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](d_.ID, objc.Sel("absoluteLinkURL"))
	return rv
} /* debug [instance_properties/getter]: absoluteLinkURL */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAreaElement/alt
func (d_ DOMHTMLAreaElement) Alt() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("alt"))
	return rv
} /* debug [instance_properties/getter]: alt */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAreaElement/alt
func (d_ DOMHTMLAreaElement) SetAlt(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAlt:"), value)
} /* debug [instance_properties/setter]: alt */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAreaElement/coords
func (d_ DOMHTMLAreaElement) Coords() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("coords"))
	return rv
} /* debug [instance_properties/getter]: coords */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAreaElement/coords
func (d_ DOMHTMLAreaElement) SetCoords(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCoords:"), value)
} /* debug [instance_properties/setter]: coords */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAreaElement/hashName
func (d_ DOMHTMLAreaElement) HashName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("hashName"))
	return rv
} /* debug [instance_properties/getter]: hashName */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAreaElement/host
func (d_ DOMHTMLAreaElement) Host() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("host"))
	return rv
} /* debug [instance_properties/getter]: host */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAreaElement/hostname
func (d_ DOMHTMLAreaElement) Hostname() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("hostname"))
	return rv
} /* debug [instance_properties/getter]: hostname */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAreaElement/href
func (d_ DOMHTMLAreaElement) Href() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("href"))
	return rv
} /* debug [instance_properties/getter]: href */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAreaElement/href
func (d_ DOMHTMLAreaElement) SetHref(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHref:"), value)
} /* debug [instance_properties/setter]: href */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAreaElement/noHref
func (d_ DOMHTMLAreaElement) NoHref() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("noHref"))
	return rv
} /* debug [instance_properties/getter]: noHref */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAreaElement/noHref
func (d_ DOMHTMLAreaElement) SetNoHref(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setNoHref:"), value)
} /* debug [instance_properties/setter]: noHref */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAreaElement/pathname
func (d_ DOMHTMLAreaElement) Pathname() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("pathname"))
	return rv
} /* debug [instance_properties/getter]: pathname */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAreaElement/port
func (d_ DOMHTMLAreaElement) Port() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("port"))
	return rv
} /* debug [instance_properties/getter]: port */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAreaElement/protocol
func (d_ DOMHTMLAreaElement) Protocol() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("protocol"))
	return rv
} /* debug [instance_properties/getter]: protocol */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAreaElement/search
func (d_ DOMHTMLAreaElement) Search() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("search"))
	return rv
} /* debug [instance_properties/getter]: search */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAreaElement/shape
func (d_ DOMHTMLAreaElement) Shape() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("shape"))
	return rv
} /* debug [instance_properties/getter]: shape */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAreaElement/shape
func (d_ DOMHTMLAreaElement) SetShape(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setShape:"), value)
} /* debug [instance_properties/setter]: shape */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAreaElement/target
func (d_ DOMHTMLAreaElement) Target() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("target"))
	return rv
} /* debug [instance_properties/getter]: target */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAreaElement/target
func (d_ DOMHTMLAreaElement) SetTarget(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTarget:"), value)
} /* debug [instance_properties/setter]: target */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMHTMLAreaElement */
