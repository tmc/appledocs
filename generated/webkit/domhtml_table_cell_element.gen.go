// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLTableCellElement */

/* debug [class_header]: Header for DOMHTMLTableCellElement */
// The class instance for the [DOMHTMLTableCellElement] class.
var (
	DOMHTMLTableCellElementClass     _DOMHTMLTableCellElementClass
	DOMHTMLTableCellElementClassOnce sync.Once
)

func getDOMHTMLTableCellElementClass() _DOMHTMLTableCellElementClass {
	DOMHTMLTableCellElementClassOnce.Do(func() {
		DOMHTMLTableCellElementClass = _DOMHTMLTableCellElementClass{objc.GetClass("DOMHTMLTableCellElement")}
	})
	return DOMHTMLTableCellElementClass
}

type _DOMHTMLTableCellElementClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMHTMLTableCellElement */
// An interface definition for the [DOMHTMLTableCellElement] class.
type IDOMHTMLTableCellElement interface {
	IDOMHTMLElement

	/* debug [class_interface_properties]: Properties for DOMHTMLTableCellElement */
	// properties:
	Abbr() objc.IObject /* cross-framework: NSString */
	SetAbbr(value objc.IObject /* cross-framework: NSString */)
	Align() objc.IObject /* cross-framework: NSString */
	SetAlign(value objc.IObject /* cross-framework: NSString */)
	Axis() objc.IObject /* cross-framework: NSString */
	SetAxis(value objc.IObject /* cross-framework: NSString */)
	BgColor() objc.IObject /* cross-framework: NSString */
	SetBgColor(value objc.IObject /* cross-framework: NSString */)
	CellIndex() int
	Ch() objc.IObject /* cross-framework: NSString */
	SetCh(value objc.IObject /* cross-framework: NSString */)
	ChOff() objc.IObject /* cross-framework: NSString */
	SetChOff(value objc.IObject /* cross-framework: NSString */)
	ColSpan() int
	SetColSpan(value int)
	Headers() objc.IObject /* cross-framework: NSString */
	SetHeaders(value objc.IObject /* cross-framework: NSString */)
	Height() objc.IObject /* cross-framework: NSString */
	SetHeight(value objc.IObject /* cross-framework: NSString */)
	NoWrap() bool
	SetNoWrap(value bool)
	RowSpan() int
	SetRowSpan(value int)
	Scope() objc.IObject /* cross-framework: NSString */
	SetScope(value objc.IObject /* cross-framework: NSString */)
	VAlign() objc.IObject /* cross-framework: NSString */
	SetVAlign(value objc.IObject /* cross-framework: NSString */)
	Width() objc.IObject /* cross-framework: NSString */
	SetWidth(value objc.IObject /* cross-framework: NSString */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMHTMLTableCellElement */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMHTMLTableCellElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLTableCellElementClass) Alloc() DOMHTMLTableCellElement {
	rv := objc.Send[DOMHTMLTableCellElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLTableCellElementClass) New() DOMHTMLTableCellElement {
	rv := objc.Send[DOMHTMLTableCellElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLTableCellElement) Init() DOMHTMLTableCellElement {
	rv := objc.Send[DOMHTMLTableCellElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLTableCellElement) Autorelease() DOMHTMLTableCellElement {
	rv := objc.Send[DOMHTMLTableCellElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLTableCellElement creates a new DOMHTMLTableCellElement instance.
func NewDOMHTMLTableCellElement() DOMHTMLTableCellElement {
	return getDOMHTMLTableCellElementClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMHTMLTableCellElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement
type DOMHTMLTableCellElement struct {
	DOMHTMLElement
}

// DOMHTMLTableCellElementFrom constructs a [DOMHTMLTableCellElement] from an unsafe.Pointer.
func DOMHTMLTableCellElementFrom(ptr unsafe.Pointer) DOMHTMLTableCellElement {
	return DOMHTMLTableCellElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMHTMLTableCellElement */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMHTMLTableCellElement */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMHTMLTableCellElement */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMHTMLTableCellElement */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMHTMLTableCellElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/abbr
func (d_ DOMHTMLTableCellElement) Abbr() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("abbr"))
	return rv
} /* debug [instance_properties/getter]: abbr */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/abbr
func (d_ DOMHTMLTableCellElement) SetAbbr(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAbbr:"), value)
} /* debug [instance_properties/setter]: abbr */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/align
func (d_ DOMHTMLTableCellElement) Align() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("align"))
	return rv
} /* debug [instance_properties/getter]: align */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/align
func (d_ DOMHTMLTableCellElement) SetAlign(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAlign:"), value)
} /* debug [instance_properties/setter]: align */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/axis
func (d_ DOMHTMLTableCellElement) Axis() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("axis"))
	return rv
} /* debug [instance_properties/getter]: axis */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/axis
func (d_ DOMHTMLTableCellElement) SetAxis(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAxis:"), value)
} /* debug [instance_properties/setter]: axis */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/bgColor
func (d_ DOMHTMLTableCellElement) BgColor() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("bgColor"))
	return rv
} /* debug [instance_properties/getter]: bgColor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/bgColor
func (d_ DOMHTMLTableCellElement) SetBgColor(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBgColor:"), value)
} /* debug [instance_properties/setter]: bgColor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/cellIndex
func (d_ DOMHTMLTableCellElement) CellIndex() int {
	rv := objc.Send[int](d_.ID, objc.Sel("cellIndex"))
	return rv
} /* debug [instance_properties/getter]: cellIndex */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/ch
func (d_ DOMHTMLTableCellElement) Ch() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("ch"))
	return rv
} /* debug [instance_properties/getter]: ch */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/ch
func (d_ DOMHTMLTableCellElement) SetCh(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCh:"), value)
} /* debug [instance_properties/setter]: ch */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/chOff
func (d_ DOMHTMLTableCellElement) ChOff() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("chOff"))
	return rv
} /* debug [instance_properties/getter]: chOff */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/chOff
func (d_ DOMHTMLTableCellElement) SetChOff(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setChOff:"), value)
} /* debug [instance_properties/setter]: chOff */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/colSpan
func (d_ DOMHTMLTableCellElement) ColSpan() int {
	rv := objc.Send[int](d_.ID, objc.Sel("colSpan"))
	return rv
} /* debug [instance_properties/getter]: colSpan */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/colSpan
func (d_ DOMHTMLTableCellElement) SetColSpan(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setColSpan:"), value)
} /* debug [instance_properties/setter]: colSpan */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/headers
func (d_ DOMHTMLTableCellElement) Headers() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("headers"))
	return rv
} /* debug [instance_properties/getter]: headers */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/headers
func (d_ DOMHTMLTableCellElement) SetHeaders(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHeaders:"), value)
} /* debug [instance_properties/setter]: headers */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/height
func (d_ DOMHTMLTableCellElement) Height() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("height"))
	return rv
} /* debug [instance_properties/getter]: height */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/height
func (d_ DOMHTMLTableCellElement) SetHeight(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHeight:"), value)
} /* debug [instance_properties/setter]: height */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/noWrap
func (d_ DOMHTMLTableCellElement) NoWrap() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("noWrap"))
	return rv
} /* debug [instance_properties/getter]: noWrap */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/noWrap
func (d_ DOMHTMLTableCellElement) SetNoWrap(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setNoWrap:"), value)
} /* debug [instance_properties/setter]: noWrap */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/rowSpan
func (d_ DOMHTMLTableCellElement) RowSpan() int {
	rv := objc.Send[int](d_.ID, objc.Sel("rowSpan"))
	return rv
} /* debug [instance_properties/getter]: rowSpan */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/rowSpan
func (d_ DOMHTMLTableCellElement) SetRowSpan(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setRowSpan:"), value)
} /* debug [instance_properties/setter]: rowSpan */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/scope
func (d_ DOMHTMLTableCellElement) Scope() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("scope"))
	return rv
} /* debug [instance_properties/getter]: scope */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/scope
func (d_ DOMHTMLTableCellElement) SetScope(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setScope:"), value)
} /* debug [instance_properties/setter]: scope */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/vAlign
func (d_ DOMHTMLTableCellElement) VAlign() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("vAlign"))
	return rv
} /* debug [instance_properties/getter]: vAlign */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/vAlign
func (d_ DOMHTMLTableCellElement) SetVAlign(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setVAlign:"), value)
} /* debug [instance_properties/setter]: vAlign */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/width
func (d_ DOMHTMLTableCellElement) Width() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("width"))
	return rv
} /* debug [instance_properties/getter]: width */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableCellElement/width
func (d_ DOMHTMLTableCellElement) SetWidth(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWidth:"), value)
} /* debug [instance_properties/setter]: width */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMHTMLTableCellElement */
