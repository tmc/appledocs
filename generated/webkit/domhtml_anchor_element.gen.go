// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMHTMLAnchorElement */


/* debug [class_header]: Header for DOMHTMLAnchorElement */
// The class instance for the [DOMHTMLAnchorElement] class.
var (
	DOMHTMLAnchorElementClass     _DOMHTMLAnchorElementClass
	DOMHTMLAnchorElementClassOnce sync.Once
)

func getDOMHTMLAnchorElementClass() _DOMHTMLAnchorElementClass {
	DOMHTMLAnchorElementClassOnce.Do(func() {
		DOMHTMLAnchorElementClass = _DOMHTMLAnchorElementClass{objc.GetClass("DOMHTMLAnchorElement")}
	})
	return DOMHTMLAnchorElementClass
}

type _DOMHTMLAnchorElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLAnchorElement */
// An interface definition for the [DOMHTMLAnchorElement] class.
type IDOMHTMLAnchorElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLAnchorElement */
	// properties:
	AbsoluteLinkURL() objc.IObject /* cross-framework: NSURL */
	AccessKey() objc.IObject /* cross-framework: NSString */
	SetAccessKey(value objc.IObject /* cross-framework: NSString */)
	Charset() objc.IObject /* cross-framework: NSString */
	SetCharset(value objc.IObject /* cross-framework: NSString */)
	Coords() objc.IObject /* cross-framework: NSString */
	SetCoords(value objc.IObject /* cross-framework: NSString */)
	HashName() objc.IObject /* cross-framework: NSString */
	Host() objc.IObject /* cross-framework: NSString */
	Hostname() objc.IObject /* cross-framework: NSString */
	Href() objc.IObject /* cross-framework: NSString */
	SetHref(value objc.IObject /* cross-framework: NSString */)
	Hreflang() objc.IObject /* cross-framework: NSString */
	SetHreflang(value objc.IObject /* cross-framework: NSString */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Pathname() objc.IObject /* cross-framework: NSString */
	Port() objc.IObject /* cross-framework: NSString */
	Protocol() objc.IObject /* cross-framework: NSString */
	Rel() objc.IObject /* cross-framework: NSString */
	SetRel(value objc.IObject /* cross-framework: NSString */)
	Rev() objc.IObject /* cross-framework: NSString */
	SetRev(value objc.IObject /* cross-framework: NSString */)
	Search() objc.IObject /* cross-framework: NSString */
	Shape() objc.IObject /* cross-framework: NSString */
	SetShape(value objc.IObject /* cross-framework: NSString */)
	Target() objc.IObject /* cross-framework: NSString */
	SetTarget(value objc.IObject /* cross-framework: NSString */)
	Text() objc.IObject /* cross-framework: NSString */
	Type() objc.IObject /* cross-framework: NSString */
	SetType(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLAnchorElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLAnchorElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLAnchorElementClass) Alloc() DOMHTMLAnchorElement {
	rv := objc.Send[DOMHTMLAnchorElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLAnchorElementClass) New() DOMHTMLAnchorElement {
	rv := objc.Send[DOMHTMLAnchorElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLAnchorElement) Init() DOMHTMLAnchorElement {
	rv := objc.Send[DOMHTMLAnchorElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLAnchorElement) Autorelease() DOMHTMLAnchorElement {
	rv := objc.Send[DOMHTMLAnchorElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLAnchorElement creates a new DOMHTMLAnchorElement instance.
func NewDOMHTMLAnchorElement() DOMHTMLAnchorElement {
	return getDOMHTMLAnchorElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLAnchorElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement
type DOMHTMLAnchorElement struct {
	DOMHTMLElement
}

// DOMHTMLAnchorElementFrom constructs a [DOMHTMLAnchorElement] from an unsafe.Pointer.
func DOMHTMLAnchorElementFrom(ptr unsafe.Pointer) DOMHTMLAnchorElement {
	return DOMHTMLAnchorElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLAnchorElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLAnchorElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLAnchorElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLAnchorElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLAnchorElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/absoluteLinkURL
func (d_ DOMHTMLAnchorElement) AbsoluteLinkURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](d_.ID, objc.Sel("absoluteLinkURL"))
	return rv
}/* debug [instance_properties/getter]: absoluteLinkURL */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/accessKey
func (d_ DOMHTMLAnchorElement) AccessKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("accessKey"))
	return rv
}/* debug [instance_properties/getter]: accessKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/accessKey
func (d_ DOMHTMLAnchorElement) SetAccessKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAccessKey:"), value)
}/* debug [instance_properties/setter]: accessKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/charset
func (d_ DOMHTMLAnchorElement) Charset() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("charset"))
	return rv
}/* debug [instance_properties/getter]: charset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/charset
func (d_ DOMHTMLAnchorElement) SetCharset(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCharset:"), value)
}/* debug [instance_properties/setter]: charset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/coords
func (d_ DOMHTMLAnchorElement) Coords() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("coords"))
	return rv
}/* debug [instance_properties/getter]: coords */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/coords
func (d_ DOMHTMLAnchorElement) SetCoords(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCoords:"), value)
}/* debug [instance_properties/setter]: coords */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/hashName
func (d_ DOMHTMLAnchorElement) HashName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("hashName"))
	return rv
}/* debug [instance_properties/getter]: hashName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/host
func (d_ DOMHTMLAnchorElement) Host() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("host"))
	return rv
}/* debug [instance_properties/getter]: host */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/hostname
func (d_ DOMHTMLAnchorElement) Hostname() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("hostname"))
	return rv
}/* debug [instance_properties/getter]: hostname */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/href
func (d_ DOMHTMLAnchorElement) Href() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("href"))
	return rv
}/* debug [instance_properties/getter]: href */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/href
func (d_ DOMHTMLAnchorElement) SetHref(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHref:"), value)
}/* debug [instance_properties/setter]: href */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/hreflang
func (d_ DOMHTMLAnchorElement) Hreflang() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("hreflang"))
	return rv
}/* debug [instance_properties/getter]: hreflang */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/hreflang
func (d_ DOMHTMLAnchorElement) SetHreflang(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHreflang:"), value)
}/* debug [instance_properties/setter]: hreflang */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/name
func (d_ DOMHTMLAnchorElement) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/name
func (d_ DOMHTMLAnchorElement) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/pathname
func (d_ DOMHTMLAnchorElement) Pathname() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("pathname"))
	return rv
}/* debug [instance_properties/getter]: pathname */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/port
func (d_ DOMHTMLAnchorElement) Port() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("port"))
	return rv
}/* debug [instance_properties/getter]: port */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/protocol
func (d_ DOMHTMLAnchorElement) Protocol() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("protocol"))
	return rv
}/* debug [instance_properties/getter]: protocol */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/rel
func (d_ DOMHTMLAnchorElement) Rel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("rel"))
	return rv
}/* debug [instance_properties/getter]: rel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/rel
func (d_ DOMHTMLAnchorElement) SetRel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setRel:"), value)
}/* debug [instance_properties/setter]: rel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/rev
func (d_ DOMHTMLAnchorElement) Rev() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("rev"))
	return rv
}/* debug [instance_properties/getter]: rev */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/rev
func (d_ DOMHTMLAnchorElement) SetRev(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setRev:"), value)
}/* debug [instance_properties/setter]: rev */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/search
func (d_ DOMHTMLAnchorElement) Search() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("search"))
	return rv
}/* debug [instance_properties/getter]: search */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/shape
func (d_ DOMHTMLAnchorElement) Shape() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("shape"))
	return rv
}/* debug [instance_properties/getter]: shape */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/shape
func (d_ DOMHTMLAnchorElement) SetShape(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setShape:"), value)
}/* debug [instance_properties/setter]: shape */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/target
func (d_ DOMHTMLAnchorElement) Target() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("target"))
	return rv
}/* debug [instance_properties/getter]: target */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/target
func (d_ DOMHTMLAnchorElement) SetTarget(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTarget:"), value)
}/* debug [instance_properties/setter]: target */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/text
func (d_ DOMHTMLAnchorElement) Text() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("text"))
	return rv
}/* debug [instance_properties/getter]: text */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/type
func (d_ DOMHTMLAnchorElement) Type() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAnchorElement/type
func (d_ DOMHTMLAnchorElement) SetType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setType:"), value)
}/* debug [instance_properties/setter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLAnchorElement */



