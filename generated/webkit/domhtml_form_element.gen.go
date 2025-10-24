// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMHTMLFormElement */


/* debug [class_header]: Header for DOMHTMLFormElement */
// The class instance for the [DOMHTMLFormElement] class.
var (
	DOMHTMLFormElementClass     _DOMHTMLFormElementClass
	DOMHTMLFormElementClassOnce sync.Once
)

func getDOMHTMLFormElementClass() _DOMHTMLFormElementClass {
	DOMHTMLFormElementClassOnce.Do(func() {
		DOMHTMLFormElementClass = _DOMHTMLFormElementClass{objc.GetClass("DOMHTMLFormElement")}
	})
	return DOMHTMLFormElementClass
}

type _DOMHTMLFormElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLFormElement */
// An interface definition for the [DOMHTMLFormElement] class.
type IDOMHTMLFormElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLFormElement */
	// properties:
	AcceptCharset() objc.IObject /* cross-framework: NSString */
	SetAcceptCharset(value objc.IObject /* cross-framework: NSString */)
	Action() objc.IObject /* cross-framework: NSString */
	SetAction(value objc.IObject /* cross-framework: NSString */)
	Elements() IDOMHTMLCollection
	Encoding() objc.IObject /* cross-framework: NSString */
	SetEncoding(value objc.IObject /* cross-framework: NSString */)
	Enctype() objc.IObject /* cross-framework: NSString */
	SetEnctype(value objc.IObject /* cross-framework: NSString */)
	Length() int
	Method() objc.IObject /* cross-framework: NSString */
	SetMethod(value objc.IObject /* cross-framework: NSString */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Target() objc.IObject /* cross-framework: NSString */
	SetTarget(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLFormElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLFormElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLFormElementClass) Alloc() DOMHTMLFormElement {
	rv := objc.Send[DOMHTMLFormElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLFormElementClass) New() DOMHTMLFormElement {
	rv := objc.Send[DOMHTMLFormElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLFormElement) Init() DOMHTMLFormElement {
	rv := objc.Send[DOMHTMLFormElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLFormElement) Autorelease() DOMHTMLFormElement {
	rv := objc.Send[DOMHTMLFormElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLFormElement creates a new DOMHTMLFormElement instance.
func NewDOMHTMLFormElement() DOMHTMLFormElement {
	return getDOMHTMLFormElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLFormElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFormElement
type DOMHTMLFormElement struct {
	DOMHTMLElement
}

// DOMHTMLFormElementFrom constructs a [DOMHTMLFormElement] from an unsafe.Pointer.
func DOMHTMLFormElementFrom(ptr unsafe.Pointer) DOMHTMLFormElement {
	return DOMHTMLFormElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLFormElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLFormElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLFormElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLFormElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLFormElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFormElement/acceptCharset
func (d_ DOMHTMLFormElement) AcceptCharset() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("acceptCharset"))
	return rv
}/* debug [instance_properties/getter]: acceptCharset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFormElement/acceptCharset
func (d_ DOMHTMLFormElement) SetAcceptCharset(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAcceptCharset:"), value)
}/* debug [instance_properties/setter]: acceptCharset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFormElement/action
func (d_ DOMHTMLFormElement) Action() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("action"))
	return rv
}/* debug [instance_properties/getter]: action */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFormElement/action
func (d_ DOMHTMLFormElement) SetAction(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAction:"), value)
}/* debug [instance_properties/setter]: action */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFormElement/elements
func (d_ DOMHTMLFormElement) Elements() IDOMHTMLCollection {
	rv := objc.Send[DOMHTMLCollection](d_.ID, objc.Sel("elements"))
	return rv
}/* debug [instance_properties/getter]: elements */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFormElement/encoding
func (d_ DOMHTMLFormElement) Encoding() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("encoding"))
	return rv
}/* debug [instance_properties/getter]: encoding */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFormElement/encoding
func (d_ DOMHTMLFormElement) SetEncoding(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setEncoding:"), value)
}/* debug [instance_properties/setter]: encoding */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFormElement/enctype
func (d_ DOMHTMLFormElement) Enctype() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("enctype"))
	return rv
}/* debug [instance_properties/getter]: enctype */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFormElement/enctype
func (d_ DOMHTMLFormElement) SetEnctype(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setEnctype:"), value)
}/* debug [instance_properties/setter]: enctype */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFormElement/length
func (d_ DOMHTMLFormElement) Length() int {
	rv := objc.Send[int](d_.ID, objc.Sel("length"))
	return rv
}/* debug [instance_properties/getter]: length */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFormElement/method
func (d_ DOMHTMLFormElement) Method() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("method"))
	return rv
}/* debug [instance_properties/getter]: method */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFormElement/method
func (d_ DOMHTMLFormElement) SetMethod(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMethod:"), value)
}/* debug [instance_properties/setter]: method */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFormElement/name
func (d_ DOMHTMLFormElement) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFormElement/name
func (d_ DOMHTMLFormElement) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFormElement/target
func (d_ DOMHTMLFormElement) Target() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("target"))
	return rv
}/* debug [instance_properties/getter]: target */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFormElement/target
func (d_ DOMHTMLFormElement) SetTarget(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTarget:"), value)
}/* debug [instance_properties/setter]: target */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLFormElement */



