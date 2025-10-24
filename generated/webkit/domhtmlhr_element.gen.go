// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMHTMLHRElement */


/* debug [class_header]: Header for DOMHTMLHRElement */
// The class instance for the [DOMHTMLHRElement] class.
var (
	DOMHTMLHRElementClass     _DOMHTMLHRElementClass
	DOMHTMLHRElementClassOnce sync.Once
)

func getDOMHTMLHRElementClass() _DOMHTMLHRElementClass {
	DOMHTMLHRElementClassOnce.Do(func() {
		DOMHTMLHRElementClass = _DOMHTMLHRElementClass{objc.GetClass("DOMHTMLHRElement")}
	})
	return DOMHTMLHRElementClass
}

type _DOMHTMLHRElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLHRElement */
// An interface definition for the [DOMHTMLHRElement] class.
type IDOMHTMLHRElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLHRElement */
	// properties:
	Align() objc.IObject /* cross-framework: NSString */
	SetAlign(value objc.IObject /* cross-framework: NSString */)
	NoShade() bool
	SetNoShade(value bool)
	Size() objc.IObject /* cross-framework: NSString */
	SetSize(value objc.IObject /* cross-framework: NSString */)
	Width() objc.IObject /* cross-framework: NSString */
	SetWidth(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLHRElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLHRElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLHRElementClass) Alloc() DOMHTMLHRElement {
	rv := objc.Send[DOMHTMLHRElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLHRElementClass) New() DOMHTMLHRElement {
	rv := objc.Send[DOMHTMLHRElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLHRElement) Init() DOMHTMLHRElement {
	rv := objc.Send[DOMHTMLHRElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLHRElement) Autorelease() DOMHTMLHRElement {
	rv := objc.Send[DOMHTMLHRElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLHRElement creates a new DOMHTMLHRElement instance.
func NewDOMHTMLHRElement() DOMHTMLHRElement {
	return getDOMHTMLHRElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLHRElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLHRElement
type DOMHTMLHRElement struct {
	DOMHTMLElement
}

// DOMHTMLHRElementFrom constructs a [DOMHTMLHRElement] from an unsafe.Pointer.
func DOMHTMLHRElementFrom(ptr unsafe.Pointer) DOMHTMLHRElement {
	return DOMHTMLHRElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLHRElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLHRElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLHRElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLHRElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLHRElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLHRElement/align
func (d_ DOMHTMLHRElement) Align() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("align"))
	return rv
}/* debug [instance_properties/getter]: align */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLHRElement/align
func (d_ DOMHTMLHRElement) SetAlign(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAlign:"), value)
}/* debug [instance_properties/setter]: align */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLHRElement/noShade
func (d_ DOMHTMLHRElement) NoShade() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("noShade"))
	return rv
}/* debug [instance_properties/getter]: noShade */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLHRElement/noShade
func (d_ DOMHTMLHRElement) SetNoShade(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setNoShade:"), value)
}/* debug [instance_properties/setter]: noShade */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLHRElement/size
func (d_ DOMHTMLHRElement) Size() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("size"))
	return rv
}/* debug [instance_properties/getter]: size */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLHRElement/size
func (d_ DOMHTMLHRElement) SetSize(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSize:"), value)
}/* debug [instance_properties/setter]: size */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLHRElement/width
func (d_ DOMHTMLHRElement) Width() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("width"))
	return rv
}/* debug [instance_properties/getter]: width */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLHRElement/width
func (d_ DOMHTMLHRElement) SetWidth(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWidth:"), value)
}/* debug [instance_properties/setter]: width */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLHRElement */



