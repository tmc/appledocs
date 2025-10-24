// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMHTMLOptionElement */


/* debug [class_header]: Header for DOMHTMLOptionElement */
// The class instance for the [DOMHTMLOptionElement] class.
var (
	DOMHTMLOptionElementClass     _DOMHTMLOptionElementClass
	DOMHTMLOptionElementClassOnce sync.Once
)

func getDOMHTMLOptionElementClass() _DOMHTMLOptionElementClass {
	DOMHTMLOptionElementClassOnce.Do(func() {
		DOMHTMLOptionElementClass = _DOMHTMLOptionElementClass{objc.GetClass("DOMHTMLOptionElement")}
	})
	return DOMHTMLOptionElementClass
}

type _DOMHTMLOptionElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLOptionElement */
// An interface definition for the [DOMHTMLOptionElement] class.
type IDOMHTMLOptionElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLOptionElement */
	// properties:
	DefaultSelected() bool
	SetDefaultSelected(value bool)
	Disabled() bool
	SetDisabled(value bool)
	Form() IDOMHTMLFormElement
	Index() int
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	Selected() bool
	SetSelected(value bool)
	Text() objc.IObject /* cross-framework: NSString */
	Value() objc.IObject /* cross-framework: NSString */
	SetValue(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLOptionElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLOptionElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLOptionElementClass) Alloc() DOMHTMLOptionElement {
	rv := objc.Send[DOMHTMLOptionElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLOptionElementClass) New() DOMHTMLOptionElement {
	rv := objc.Send[DOMHTMLOptionElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLOptionElement) Init() DOMHTMLOptionElement {
	rv := objc.Send[DOMHTMLOptionElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLOptionElement) Autorelease() DOMHTMLOptionElement {
	rv := objc.Send[DOMHTMLOptionElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLOptionElement creates a new DOMHTMLOptionElement instance.
func NewDOMHTMLOptionElement() DOMHTMLOptionElement {
	return getDOMHTMLOptionElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLOptionElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOptionElement
type DOMHTMLOptionElement struct {
	DOMHTMLElement
}

// DOMHTMLOptionElementFrom constructs a [DOMHTMLOptionElement] from an unsafe.Pointer.
func DOMHTMLOptionElementFrom(ptr unsafe.Pointer) DOMHTMLOptionElement {
	return DOMHTMLOptionElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLOptionElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLOptionElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLOptionElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLOptionElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLOptionElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOptionElement/defaultSelected
func (d_ DOMHTMLOptionElement) DefaultSelected() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("defaultSelected"))
	return rv
}/* debug [instance_properties/getter]: defaultSelected */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOptionElement/defaultSelected
func (d_ DOMHTMLOptionElement) SetDefaultSelected(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDefaultSelected:"), value)
}/* debug [instance_properties/setter]: defaultSelected */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOptionElement/disabled
func (d_ DOMHTMLOptionElement) Disabled() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("disabled"))
	return rv
}/* debug [instance_properties/getter]: disabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOptionElement/disabled
func (d_ DOMHTMLOptionElement) SetDisabled(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDisabled:"), value)
}/* debug [instance_properties/setter]: disabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOptionElement/form
func (d_ DOMHTMLOptionElement) Form() IDOMHTMLFormElement {
	rv := objc.Send[DOMHTMLFormElement](d_.ID, objc.Sel("form"))
	return rv
}/* debug [instance_properties/getter]: form */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOptionElement/index
func (d_ DOMHTMLOptionElement) Index() int {
	rv := objc.Send[int](d_.ID, objc.Sel("index"))
	return rv
}/* debug [instance_properties/getter]: index */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOptionElement/label
func (d_ DOMHTMLOptionElement) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOptionElement/label
func (d_ DOMHTMLOptionElement) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOptionElement/selected
func (d_ DOMHTMLOptionElement) Selected() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("selected"))
	return rv
}/* debug [instance_properties/getter]: selected */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOptionElement/selected
func (d_ DOMHTMLOptionElement) SetSelected(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSelected:"), value)
}/* debug [instance_properties/setter]: selected */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOptionElement/text
func (d_ DOMHTMLOptionElement) Text() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("text"))
	return rv
}/* debug [instance_properties/getter]: text */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOptionElement/value
func (d_ DOMHTMLOptionElement) Value() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOptionElement/value
func (d_ DOMHTMLOptionElement) SetValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLOptionElement */



