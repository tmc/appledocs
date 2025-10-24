// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLTableRowElement */

/* debug [class_header]: Header for DOMHTMLTableRowElement */
// The class instance for the [DOMHTMLTableRowElement] class.
var (
	DOMHTMLTableRowElementClass     _DOMHTMLTableRowElementClass
	DOMHTMLTableRowElementClassOnce sync.Once
)

func getDOMHTMLTableRowElementClass() _DOMHTMLTableRowElementClass {
	DOMHTMLTableRowElementClassOnce.Do(func() {
		DOMHTMLTableRowElementClass = _DOMHTMLTableRowElementClass{objc.GetClass("DOMHTMLTableRowElement")}
	})
	return DOMHTMLTableRowElementClass
}

type _DOMHTMLTableRowElementClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMHTMLTableRowElement */
// An interface definition for the [DOMHTMLTableRowElement] class.
type IDOMHTMLTableRowElement interface {
	IDOMHTMLElement

	/* debug [class_interface_properties]: Properties for DOMHTMLTableRowElement */
	// properties:
	Align() objc.IObject /* cross-framework: NSString */
	SetAlign(value objc.IObject /* cross-framework: NSString */)
	BgColor() objc.IObject /* cross-framework: NSString */
	SetBgColor(value objc.IObject /* cross-framework: NSString */)
	Cells() IDOMHTMLCollection
	Ch() objc.IObject /* cross-framework: NSString */
	SetCh(value objc.IObject /* cross-framework: NSString */)
	ChOff() objc.IObject /* cross-framework: NSString */
	SetChOff(value objc.IObject /* cross-framework: NSString */)
	RowIndex() int
	SectionRowIndex() int
	VAlign() objc.IObject /* cross-framework: NSString */
	SetVAlign(value objc.IObject /* cross-framework: NSString */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMHTMLTableRowElement */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMHTMLTableRowElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLTableRowElementClass) Alloc() DOMHTMLTableRowElement {
	rv := objc.Send[DOMHTMLTableRowElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLTableRowElementClass) New() DOMHTMLTableRowElement {
	rv := objc.Send[DOMHTMLTableRowElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLTableRowElement) Init() DOMHTMLTableRowElement {
	rv := objc.Send[DOMHTMLTableRowElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLTableRowElement) Autorelease() DOMHTMLTableRowElement {
	rv := objc.Send[DOMHTMLTableRowElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLTableRowElement creates a new DOMHTMLTableRowElement instance.
func NewDOMHTMLTableRowElement() DOMHTMLTableRowElement {
	return getDOMHTMLTableRowElementClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMHTMLTableRowElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableRowElement
type DOMHTMLTableRowElement struct {
	DOMHTMLElement
}

// DOMHTMLTableRowElementFrom constructs a [DOMHTMLTableRowElement] from an unsafe.Pointer.
func DOMHTMLTableRowElementFrom(ptr unsafe.Pointer) DOMHTMLTableRowElement {
	return DOMHTMLTableRowElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMHTMLTableRowElement */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMHTMLTableRowElement */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMHTMLTableRowElement */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMHTMLTableRowElement */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMHTMLTableRowElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableRowElement/align
func (d_ DOMHTMLTableRowElement) Align() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("align"))
	return rv
} /* debug [instance_properties/getter]: align */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableRowElement/align
func (d_ DOMHTMLTableRowElement) SetAlign(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAlign:"), value)
} /* debug [instance_properties/setter]: align */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableRowElement/bgColor
func (d_ DOMHTMLTableRowElement) BgColor() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("bgColor"))
	return rv
} /* debug [instance_properties/getter]: bgColor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableRowElement/bgColor
func (d_ DOMHTMLTableRowElement) SetBgColor(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBgColor:"), value)
} /* debug [instance_properties/setter]: bgColor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableRowElement/cells
func (d_ DOMHTMLTableRowElement) Cells() IDOMHTMLCollection {
	rv := objc.Send[DOMHTMLCollection](d_.ID, objc.Sel("cells"))
	return rv
} /* debug [instance_properties/getter]: cells */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableRowElement/ch
func (d_ DOMHTMLTableRowElement) Ch() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("ch"))
	return rv
} /* debug [instance_properties/getter]: ch */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableRowElement/ch
func (d_ DOMHTMLTableRowElement) SetCh(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCh:"), value)
} /* debug [instance_properties/setter]: ch */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableRowElement/chOff
func (d_ DOMHTMLTableRowElement) ChOff() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("chOff"))
	return rv
} /* debug [instance_properties/getter]: chOff */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableRowElement/chOff
func (d_ DOMHTMLTableRowElement) SetChOff(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setChOff:"), value)
} /* debug [instance_properties/setter]: chOff */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableRowElement/rowIndex
func (d_ DOMHTMLTableRowElement) RowIndex() int {
	rv := objc.Send[int](d_.ID, objc.Sel("rowIndex"))
	return rv
} /* debug [instance_properties/getter]: rowIndex */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableRowElement/sectionRowIndex
func (d_ DOMHTMLTableRowElement) SectionRowIndex() int {
	rv := objc.Send[int](d_.ID, objc.Sel("sectionRowIndex"))
	return rv
} /* debug [instance_properties/getter]: sectionRowIndex */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableRowElement/vAlign
func (d_ DOMHTMLTableRowElement) VAlign() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("vAlign"))
	return rv
} /* debug [instance_properties/getter]: vAlign */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableRowElement/vAlign
func (d_ DOMHTMLTableRowElement) SetVAlign(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setVAlign:"), value)
} /* debug [instance_properties/setter]: vAlign */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMHTMLTableRowElement */
