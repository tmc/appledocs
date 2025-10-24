// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMHTMLTableSectionElement */


/* debug [class_header]: Header for DOMHTMLTableSectionElement */
// The class instance for the [DOMHTMLTableSectionElement] class.
var (
	DOMHTMLTableSectionElementClass     _DOMHTMLTableSectionElementClass
	DOMHTMLTableSectionElementClassOnce sync.Once
)

func getDOMHTMLTableSectionElementClass() _DOMHTMLTableSectionElementClass {
	DOMHTMLTableSectionElementClassOnce.Do(func() {
		DOMHTMLTableSectionElementClass = _DOMHTMLTableSectionElementClass{objc.GetClass("DOMHTMLTableSectionElement")}
	})
	return DOMHTMLTableSectionElementClass
}

type _DOMHTMLTableSectionElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLTableSectionElement */
// An interface definition for the [DOMHTMLTableSectionElement] class.
type IDOMHTMLTableSectionElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLTableSectionElement */
	// properties:
	Align() objc.IObject /* cross-framework: NSString */
	SetAlign(value objc.IObject /* cross-framework: NSString */)
	Ch() objc.IObject /* cross-framework: NSString */
	SetCh(value objc.IObject /* cross-framework: NSString */)
	ChOff() objc.IObject /* cross-framework: NSString */
	SetChOff(value objc.IObject /* cross-framework: NSString */)
	Rows() IDOMHTMLCollection
	VAlign() objc.IObject /* cross-framework: NSString */
	SetVAlign(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLTableSectionElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLTableSectionElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLTableSectionElementClass) Alloc() DOMHTMLTableSectionElement {
	rv := objc.Send[DOMHTMLTableSectionElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLTableSectionElementClass) New() DOMHTMLTableSectionElement {
	rv := objc.Send[DOMHTMLTableSectionElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLTableSectionElement) Init() DOMHTMLTableSectionElement {
	rv := objc.Send[DOMHTMLTableSectionElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLTableSectionElement) Autorelease() DOMHTMLTableSectionElement {
	rv := objc.Send[DOMHTMLTableSectionElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLTableSectionElement creates a new DOMHTMLTableSectionElement instance.
func NewDOMHTMLTableSectionElement() DOMHTMLTableSectionElement {
	return getDOMHTMLTableSectionElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLTableSectionElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableSectionElement
type DOMHTMLTableSectionElement struct {
	DOMHTMLElement
}

// DOMHTMLTableSectionElementFrom constructs a [DOMHTMLTableSectionElement] from an unsafe.Pointer.
func DOMHTMLTableSectionElementFrom(ptr unsafe.Pointer) DOMHTMLTableSectionElement {
	return DOMHTMLTableSectionElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLTableSectionElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLTableSectionElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLTableSectionElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLTableSectionElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLTableSectionElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableSectionElement/align
func (d_ DOMHTMLTableSectionElement) Align() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("align"))
	return rv
}/* debug [instance_properties/getter]: align */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableSectionElement/align
func (d_ DOMHTMLTableSectionElement) SetAlign(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAlign:"), value)
}/* debug [instance_properties/setter]: align */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableSectionElement/ch
func (d_ DOMHTMLTableSectionElement) Ch() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("ch"))
	return rv
}/* debug [instance_properties/getter]: ch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableSectionElement/ch
func (d_ DOMHTMLTableSectionElement) SetCh(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCh:"), value)
}/* debug [instance_properties/setter]: ch */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableSectionElement/chOff
func (d_ DOMHTMLTableSectionElement) ChOff() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("chOff"))
	return rv
}/* debug [instance_properties/getter]: chOff */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableSectionElement/chOff
func (d_ DOMHTMLTableSectionElement) SetChOff(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setChOff:"), value)
}/* debug [instance_properties/setter]: chOff */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableSectionElement/rows
func (d_ DOMHTMLTableSectionElement) Rows() IDOMHTMLCollection {
	rv := objc.Send[DOMHTMLCollection](d_.ID, objc.Sel("rows"))
	return rv
}/* debug [instance_properties/getter]: rows */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableSectionElement/vAlign
func (d_ DOMHTMLTableSectionElement) VAlign() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("vAlign"))
	return rv
}/* debug [instance_properties/getter]: vAlign */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTableSectionElement/vAlign
func (d_ DOMHTMLTableSectionElement) SetVAlign(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setVAlign:"), value)
}/* debug [instance_properties/setter]: vAlign */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLTableSectionElement */



