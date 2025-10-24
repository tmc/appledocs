// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMHTMLEmbedElement */


/* debug [class_header]: Header for DOMHTMLEmbedElement */
// The class instance for the [DOMHTMLEmbedElement] class.
var (
	DOMHTMLEmbedElementClass     _DOMHTMLEmbedElementClass
	DOMHTMLEmbedElementClassOnce sync.Once
)

func getDOMHTMLEmbedElementClass() _DOMHTMLEmbedElementClass {
	DOMHTMLEmbedElementClassOnce.Do(func() {
		DOMHTMLEmbedElementClass = _DOMHTMLEmbedElementClass{objc.GetClass("DOMHTMLEmbedElement")}
	})
	return DOMHTMLEmbedElementClass
}

type _DOMHTMLEmbedElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLEmbedElement */
// An interface definition for the [DOMHTMLEmbedElement] class.
type IDOMHTMLEmbedElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLEmbedElement */
	// properties:
	Align() objc.IObject /* cross-framework: NSString */
	SetAlign(value objc.IObject /* cross-framework: NSString */)
	Height() int
	SetHeight(value int)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Src() objc.IObject /* cross-framework: NSString */
	SetSrc(value objc.IObject /* cross-framework: NSString */)
	Type() objc.IObject /* cross-framework: NSString */
	SetType(value objc.IObject /* cross-framework: NSString */)
	Width() int
	SetWidth(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLEmbedElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLEmbedElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLEmbedElementClass) Alloc() DOMHTMLEmbedElement {
	rv := objc.Send[DOMHTMLEmbedElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLEmbedElementClass) New() DOMHTMLEmbedElement {
	rv := objc.Send[DOMHTMLEmbedElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLEmbedElement) Init() DOMHTMLEmbedElement {
	rv := objc.Send[DOMHTMLEmbedElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLEmbedElement) Autorelease() DOMHTMLEmbedElement {
	rv := objc.Send[DOMHTMLEmbedElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLEmbedElement creates a new DOMHTMLEmbedElement instance.
func NewDOMHTMLEmbedElement() DOMHTMLEmbedElement {
	return getDOMHTMLEmbedElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLEmbedElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLEmbedElement
type DOMHTMLEmbedElement struct {
	DOMHTMLElement
}

// DOMHTMLEmbedElementFrom constructs a [DOMHTMLEmbedElement] from an unsafe.Pointer.
func DOMHTMLEmbedElementFrom(ptr unsafe.Pointer) DOMHTMLEmbedElement {
	return DOMHTMLEmbedElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLEmbedElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLEmbedElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLEmbedElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLEmbedElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLEmbedElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLEmbedElement/align
func (d_ DOMHTMLEmbedElement) Align() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("align"))
	return rv
}/* debug [instance_properties/getter]: align */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLEmbedElement/align
func (d_ DOMHTMLEmbedElement) SetAlign(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAlign:"), value)
}/* debug [instance_properties/setter]: align */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLEmbedElement/height
func (d_ DOMHTMLEmbedElement) Height() int {
	rv := objc.Send[int](d_.ID, objc.Sel("height"))
	return rv
}/* debug [instance_properties/getter]: height */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLEmbedElement/height
func (d_ DOMHTMLEmbedElement) SetHeight(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHeight:"), value)
}/* debug [instance_properties/setter]: height */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLEmbedElement/name
func (d_ DOMHTMLEmbedElement) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLEmbedElement/name
func (d_ DOMHTMLEmbedElement) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLEmbedElement/src
func (d_ DOMHTMLEmbedElement) Src() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("src"))
	return rv
}/* debug [instance_properties/getter]: src */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLEmbedElement/src
func (d_ DOMHTMLEmbedElement) SetSrc(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSrc:"), value)
}/* debug [instance_properties/setter]: src */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLEmbedElement/type
func (d_ DOMHTMLEmbedElement) Type() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLEmbedElement/type
func (d_ DOMHTMLEmbedElement) SetType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setType:"), value)
}/* debug [instance_properties/setter]: type */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLEmbedElement/width
func (d_ DOMHTMLEmbedElement) Width() int {
	rv := objc.Send[int](d_.ID, objc.Sel("width"))
	return rv
}/* debug [instance_properties/getter]: width */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLEmbedElement/width
func (d_ DOMHTMLEmbedElement) SetWidth(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWidth:"), value)
}/* debug [instance_properties/setter]: width */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLEmbedElement */



