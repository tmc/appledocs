// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMHTMLButtonElement */


/* debug [class_header]: Header for DOMHTMLButtonElement */
// The class instance for the [DOMHTMLButtonElement] class.
var (
	DOMHTMLButtonElementClass     _DOMHTMLButtonElementClass
	DOMHTMLButtonElementClassOnce sync.Once
)

func getDOMHTMLButtonElementClass() _DOMHTMLButtonElementClass {
	DOMHTMLButtonElementClassOnce.Do(func() {
		DOMHTMLButtonElementClass = _DOMHTMLButtonElementClass{objc.GetClass("DOMHTMLButtonElement")}
	})
	return DOMHTMLButtonElementClass
}

type _DOMHTMLButtonElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLButtonElement */
// An interface definition for the [DOMHTMLButtonElement] class.
type IDOMHTMLButtonElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLButtonElement */
	// properties:
	AccessKey() objc.IObject /* cross-framework: NSString */
	SetAccessKey(value objc.IObject /* cross-framework: NSString */)
	Autofocus() bool
	SetAutofocus(value bool)
	Disabled() bool
	SetDisabled(value bool)
	Form() IDOMHTMLFormElement
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Type() objc.IObject /* cross-framework: NSString */
	SetType(value objc.IObject /* cross-framework: NSString */)
	Value() objc.IObject /* cross-framework: NSString */
	SetValue(value objc.IObject /* cross-framework: NSString */)
	WillValidate() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLButtonElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLButtonElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLButtonElementClass) Alloc() DOMHTMLButtonElement {
	rv := objc.Send[DOMHTMLButtonElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLButtonElementClass) New() DOMHTMLButtonElement {
	rv := objc.Send[DOMHTMLButtonElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLButtonElement) Init() DOMHTMLButtonElement {
	rv := objc.Send[DOMHTMLButtonElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLButtonElement) Autorelease() DOMHTMLButtonElement {
	rv := objc.Send[DOMHTMLButtonElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLButtonElement creates a new DOMHTMLButtonElement instance.
func NewDOMHTMLButtonElement() DOMHTMLButtonElement {
	return getDOMHTMLButtonElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLButtonElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLButtonElement
type DOMHTMLButtonElement struct {
	DOMHTMLElement
}

// DOMHTMLButtonElementFrom constructs a [DOMHTMLButtonElement] from an unsafe.Pointer.
func DOMHTMLButtonElementFrom(ptr unsafe.Pointer) DOMHTMLButtonElement {
	return DOMHTMLButtonElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLButtonElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLButtonElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLButtonElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLButtonElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLButtonElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLButtonElement/accessKey
func (d_ DOMHTMLButtonElement) AccessKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("accessKey"))
	return rv
}/* debug [instance_properties/getter]: accessKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLButtonElement/accessKey
func (d_ DOMHTMLButtonElement) SetAccessKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAccessKey:"), value)
}/* debug [instance_properties/setter]: accessKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLButtonElement/autofocus
func (d_ DOMHTMLButtonElement) Autofocus() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("autofocus"))
	return rv
}/* debug [instance_properties/getter]: autofocus */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLButtonElement/autofocus
func (d_ DOMHTMLButtonElement) SetAutofocus(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAutofocus:"), value)
}/* debug [instance_properties/setter]: autofocus */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLButtonElement/disabled
func (d_ DOMHTMLButtonElement) Disabled() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("disabled"))
	return rv
}/* debug [instance_properties/getter]: disabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLButtonElement/disabled
func (d_ DOMHTMLButtonElement) SetDisabled(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDisabled:"), value)
}/* debug [instance_properties/setter]: disabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLButtonElement/form
func (d_ DOMHTMLButtonElement) Form() IDOMHTMLFormElement {
	rv := objc.Send[DOMHTMLFormElement](d_.ID, objc.Sel("form"))
	return rv
}/* debug [instance_properties/getter]: form */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLButtonElement/name
func (d_ DOMHTMLButtonElement) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLButtonElement/name
func (d_ DOMHTMLButtonElement) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLButtonElement/type
func (d_ DOMHTMLButtonElement) Type() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLButtonElement/type
func (d_ DOMHTMLButtonElement) SetType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setType:"), value)
}/* debug [instance_properties/setter]: type */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLButtonElement/value
func (d_ DOMHTMLButtonElement) Value() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLButtonElement/value
func (d_ DOMHTMLButtonElement) SetValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLButtonElement/willValidate
func (d_ DOMHTMLButtonElement) WillValidate() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("willValidate"))
	return rv
}/* debug [instance_properties/getter]: willValidate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLButtonElement */



