// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLTableElement */

/* debug [class_header]: Header for DOMHTMLTableElement */
// The class instance for the [DOMHTMLTableElement] class.
var (
	DOMHTMLTableElementClass     _DOMHTMLTableElementClass
	DOMHTMLTableElementClassOnce sync.Once
)

func getDOMHTMLTableElementClass() _DOMHTMLTableElementClass {
	DOMHTMLTableElementClassOnce.Do(func() {
		DOMHTMLTableElementClass = _DOMHTMLTableElementClass{objc.GetClass("DOMHTMLTableElement")}
	})
	return DOMHTMLTableElementClass
}

type _DOMHTMLTableElementClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMHTMLTableElement */
// An interface definition for the [DOMHTMLTableElement] class.
type IDOMHTMLTableElement interface {
	IDOMHTMLElement

	/* debug [class_interface_properties]: Properties for DOMHTMLTableElement */
	// properties:
	Align() objc.IObject /* cross-framework: NSString */
	SetAlign(value objc.IObject /* cross-framework: NSString */)
	BgColor() objc.IObject /* cross-framework: NSString */
	SetBgColor(value objc.IObject /* cross-framework: NSString */)
	Border() objc.IObject /* cross-framework: NSString */
	SetBorder(value objc.IObject /* cross-framework: NSString */)
	Caption() IDOMHTMLTableCaptionElement
	SetCaption(value IDOMHTMLTableCaptionElement)
	CellPadding() objc.IObject /* cross-framework: NSString */
	SetCellPadding(value objc.IObject /* cross-framework: NSString */)
	CellSpacing() objc.IObject /* cross-framework: NSString */
	SetCellSpacing(value objc.IObject /* cross-framework: NSString */)
	FrameBorders() objc.IObject /* cross-framework: NSString */
	SetFrameBorders(value objc.IObject /* cross-framework: NSString */)
	Rows() IDOMHTMLCollection
	Rules() objc.IObject /* cross-framework: NSString */
	SetRules(value objc.IObject /* cross-framework: NSString */)
	Summary() objc.IObject /* cross-framework: NSString */
	SetSummary(value objc.IObject /* cross-framework: NSString */)
	TBodies() IDOMHTMLCollection
	TFoot() IDOMHTMLTableSectionElement
	SetTFoot(value IDOMHTMLTableSectionElement)
	THead() IDOMHTMLTableSectionElement
	SetTHead(value IDOMHTMLTableSectionElement)
	Width() objc.IObject /* cross-framework: NSString */
	SetWidth(value objc.IObject /* cross-framework: NSString */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMHTMLTableElement */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMHTMLTableElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLTableElementClass) Alloc() DOMHTMLTableElement {
	rv := objc.Send[DOMHTMLTableElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLTableElementClass) New() DOMHTMLTableElement {
	rv := objc.Send[DOMHTMLTableElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLTableElement) Init() DOMHTMLTableElement {
	rv := objc.Send[DOMHTMLTableElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLTableElement) Autorelease() DOMHTMLTableElement {
	rv := objc.Send[DOMHTMLTableElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLTableElement creates a new DOMHTMLTableElement instance.
func NewDOMHTMLTableElement() DOMHTMLTableElement {
	return getDOMHTMLTableElementClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMHTMLTableElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement
type DOMHTMLTableElement struct {
	DOMHTMLElement
}

// DOMHTMLTableElementFrom constructs a [DOMHTMLTableElement] from an unsafe.Pointer.
func DOMHTMLTableElementFrom(ptr unsafe.Pointer) DOMHTMLTableElement {
	return DOMHTMLTableElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMHTMLTableElement */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMHTMLTableElement */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMHTMLTableElement */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMHTMLTableElement */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMHTMLTableElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/align
func (d_ DOMHTMLTableElement) Align() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("align"))
	return rv
} /* debug [instance_properties/getter]: align */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/align
func (d_ DOMHTMLTableElement) SetAlign(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAlign:"), value)
} /* debug [instance_properties/setter]: align */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/bgColor
func (d_ DOMHTMLTableElement) BgColor() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("bgColor"))
	return rv
} /* debug [instance_properties/getter]: bgColor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/bgColor
func (d_ DOMHTMLTableElement) SetBgColor(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBgColor:"), value)
} /* debug [instance_properties/setter]: bgColor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/border
func (d_ DOMHTMLTableElement) Border() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("border"))
	return rv
} /* debug [instance_properties/getter]: border */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/border
func (d_ DOMHTMLTableElement) SetBorder(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBorder:"), value)
} /* debug [instance_properties/setter]: border */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/caption
func (d_ DOMHTMLTableElement) Caption() IDOMHTMLTableCaptionElement {
	rv := objc.Send[DOMHTMLTableCaptionElement](d_.ID, objc.Sel("caption"))
	return rv
} /* debug [instance_properties/getter]: caption */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/caption
func (d_ DOMHTMLTableElement) SetCaption(value IDOMHTMLTableCaptionElement) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCaption:"), value)
} /* debug [instance_properties/setter]: caption */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/cellPadding
func (d_ DOMHTMLTableElement) CellPadding() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("cellPadding"))
	return rv
} /* debug [instance_properties/getter]: cellPadding */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/cellPadding
func (d_ DOMHTMLTableElement) SetCellPadding(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCellPadding:"), value)
} /* debug [instance_properties/setter]: cellPadding */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/cellSpacing
func (d_ DOMHTMLTableElement) CellSpacing() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("cellSpacing"))
	return rv
} /* debug [instance_properties/getter]: cellSpacing */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/cellSpacing
func (d_ DOMHTMLTableElement) SetCellSpacing(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCellSpacing:"), value)
} /* debug [instance_properties/setter]: cellSpacing */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/frameBorders
func (d_ DOMHTMLTableElement) FrameBorders() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("frameBorders"))
	return rv
} /* debug [instance_properties/getter]: frameBorders */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/frameBorders
func (d_ DOMHTMLTableElement) SetFrameBorders(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFrameBorders:"), value)
} /* debug [instance_properties/setter]: frameBorders */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/rows
func (d_ DOMHTMLTableElement) Rows() IDOMHTMLCollection {
	rv := objc.Send[DOMHTMLCollection](d_.ID, objc.Sel("rows"))
	return rv
} /* debug [instance_properties/getter]: rows */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/rules
func (d_ DOMHTMLTableElement) Rules() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("rules"))
	return rv
} /* debug [instance_properties/getter]: rules */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/rules
func (d_ DOMHTMLTableElement) SetRules(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setRules:"), value)
} /* debug [instance_properties/setter]: rules */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/summary
func (d_ DOMHTMLTableElement) Summary() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("summary"))
	return rv
} /* debug [instance_properties/getter]: summary */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/summary
func (d_ DOMHTMLTableElement) SetSummary(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSummary:"), value)
} /* debug [instance_properties/setter]: summary */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/tBodies
func (d_ DOMHTMLTableElement) TBodies() IDOMHTMLCollection {
	rv := objc.Send[DOMHTMLCollection](d_.ID, objc.Sel("tBodies"))
	return rv
} /* debug [instance_properties/getter]: tBodies */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/tFoot
func (d_ DOMHTMLTableElement) TFoot() IDOMHTMLTableSectionElement {
	rv := objc.Send[DOMHTMLTableSectionElement](d_.ID, objc.Sel("tFoot"))
	return rv
} /* debug [instance_properties/getter]: tFoot */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/tFoot
func (d_ DOMHTMLTableElement) SetTFoot(value IDOMHTMLTableSectionElement) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTFoot:"), value)
} /* debug [instance_properties/setter]: tFoot */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/tHead
func (d_ DOMHTMLTableElement) THead() IDOMHTMLTableSectionElement {
	rv := objc.Send[DOMHTMLTableSectionElement](d_.ID, objc.Sel("tHead"))
	return rv
} /* debug [instance_properties/getter]: tHead */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/tHead
func (d_ DOMHTMLTableElement) SetTHead(value IDOMHTMLTableSectionElement) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTHead:"), value)
} /* debug [instance_properties/setter]: tHead */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/width
func (d_ DOMHTMLTableElement) Width() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("width"))
	return rv
} /* debug [instance_properties/getter]: width */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableElement/width
func (d_ DOMHTMLTableElement) SetWidth(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWidth:"), value)
} /* debug [instance_properties/setter]: width */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMHTMLTableElement */
