// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMHTMLParamElement */


/* debug [class_header]: Header for DOMHTMLParamElement */
// The class instance for the [DOMHTMLParamElement] class.
var (
	DOMHTMLParamElementClass     _DOMHTMLParamElementClass
	DOMHTMLParamElementClassOnce sync.Once
)

func getDOMHTMLParamElementClass() _DOMHTMLParamElementClass {
	DOMHTMLParamElementClassOnce.Do(func() {
		DOMHTMLParamElementClass = _DOMHTMLParamElementClass{objc.GetClass("DOMHTMLParamElement")}
	})
	return DOMHTMLParamElementClass
}

type _DOMHTMLParamElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLParamElement */
// An interface definition for the [DOMHTMLParamElement] class.
type IDOMHTMLParamElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLParamElement */
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Type() objc.IObject /* cross-framework: NSString */
	SetType(value objc.IObject /* cross-framework: NSString */)
	Value() objc.IObject /* cross-framework: NSString */
	SetValue(value objc.IObject /* cross-framework: NSString */)
	ValueType() objc.IObject /* cross-framework: NSString */
	SetValueType(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLParamElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLParamElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLParamElementClass) Alloc() DOMHTMLParamElement {
	rv := objc.Send[DOMHTMLParamElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLParamElementClass) New() DOMHTMLParamElement {
	rv := objc.Send[DOMHTMLParamElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLParamElement) Init() DOMHTMLParamElement {
	rv := objc.Send[DOMHTMLParamElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLParamElement) Autorelease() DOMHTMLParamElement {
	rv := objc.Send[DOMHTMLParamElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLParamElement creates a new DOMHTMLParamElement instance.
func NewDOMHTMLParamElement() DOMHTMLParamElement {
	return getDOMHTMLParamElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLParamElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLParamElement
type DOMHTMLParamElement struct {
	DOMHTMLElement
}

// DOMHTMLParamElementFrom constructs a [DOMHTMLParamElement] from an unsafe.Pointer.
func DOMHTMLParamElementFrom(ptr unsafe.Pointer) DOMHTMLParamElement {
	return DOMHTMLParamElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLParamElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLParamElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLParamElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLParamElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLParamElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLParamElement/name
func (d_ DOMHTMLParamElement) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLParamElement/name
func (d_ DOMHTMLParamElement) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLParamElement/type
func (d_ DOMHTMLParamElement) Type() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLParamElement/type
func (d_ DOMHTMLParamElement) SetType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setType:"), value)
}/* debug [instance_properties/setter]: type */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLParamElement/value
func (d_ DOMHTMLParamElement) Value() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLParamElement/value
func (d_ DOMHTMLParamElement) SetValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLParamElement/valueType
func (d_ DOMHTMLParamElement) ValueType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("valueType"))
	return rv
}/* debug [instance_properties/getter]: valueType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLParamElement/valueType
func (d_ DOMHTMLParamElement) SetValueType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setValueType:"), value)
}/* debug [instance_properties/setter]: valueType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLParamElement */



