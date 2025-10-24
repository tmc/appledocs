// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class DOMHTMLSelectElement */


/* debug [class_header]: Header for DOMHTMLSelectElement */
// The class instance for the [DOMHTMLSelectElement] class.
var (
	DOMHTMLSelectElementClass     _DOMHTMLSelectElementClass
	DOMHTMLSelectElementClassOnce sync.Once
)

func getDOMHTMLSelectElementClass() _DOMHTMLSelectElementClass {
	DOMHTMLSelectElementClassOnce.Do(func() {
		DOMHTMLSelectElementClass = _DOMHTMLSelectElementClass{objc.GetClass("DOMHTMLSelectElement")}
	})
	return DOMHTMLSelectElementClass
}

type _DOMHTMLSelectElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLSelectElement */
// An interface definition for the [DOMHTMLSelectElement] class.
type IDOMHTMLSelectElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLSelectElement */
	// properties:
	Autofocus() bool
	SetAutofocus(value bool)
	Disabled() bool
	SetDisabled(value bool)
	Form() IDOMHTMLFormElement
	Length() int
	Multiple() bool
	SetMultiple(value bool)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Options() IDOMHTMLOptionsCollection
	SelectedIndex() int
	SetSelectedIndex(value int)
	Size() int
	SetSize(value int)
	Type() objc.IObject /* cross-framework: NSString */
	Value() objc.IObject /* cross-framework: NSString */
	SetValue(value objc.IObject /* cross-framework: NSString */)
	WillValidate() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLSelectElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLSelectElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLSelectElementClass) Alloc() DOMHTMLSelectElement {
	rv := objc.Send[DOMHTMLSelectElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLSelectElementClass) New() DOMHTMLSelectElement {
	rv := objc.Send[DOMHTMLSelectElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLSelectElement) Init() DOMHTMLSelectElement {
	rv := objc.Send[DOMHTMLSelectElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLSelectElement) Autorelease() DOMHTMLSelectElement {
	rv := objc.Send[DOMHTMLSelectElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLSelectElement creates a new DOMHTMLSelectElement instance.
func NewDOMHTMLSelectElement() DOMHTMLSelectElement {
	return getDOMHTMLSelectElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLSelectElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLSelectElement
type DOMHTMLSelectElement struct {
	DOMHTMLElement
}

// DOMHTMLSelectElementFrom constructs a [DOMHTMLSelectElement] from an unsafe.Pointer.
func DOMHTMLSelectElementFrom(ptr unsafe.Pointer) DOMHTMLSelectElement {
	return DOMHTMLSelectElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLSelectElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLSelectElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLSelectElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLSelectElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLSelectElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLSelectElement/autofocus
func (d_ DOMHTMLSelectElement) Autofocus() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("autofocus"))
	return rv
}/* debug [instance_properties/getter]: autofocus */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLSelectElement/autofocus
func (d_ DOMHTMLSelectElement) SetAutofocus(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAutofocus:"), value)
}/* debug [instance_properties/setter]: autofocus */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLSelectElement/disabled
func (d_ DOMHTMLSelectElement) Disabled() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("disabled"))
	return rv
}/* debug [instance_properties/getter]: disabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLSelectElement/disabled
func (d_ DOMHTMLSelectElement) SetDisabled(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDisabled:"), value)
}/* debug [instance_properties/setter]: disabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLSelectElement/form
func (d_ DOMHTMLSelectElement) Form() IDOMHTMLFormElement {
	rv := objc.Send[DOMHTMLFormElement](d_.ID, objc.Sel("form"))
	return rv
}/* debug [instance_properties/getter]: form */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLSelectElement/length
func (d_ DOMHTMLSelectElement) Length() int {
	rv := objc.Send[int](d_.ID, objc.Sel("length"))
	return rv
}/* debug [instance_properties/getter]: length */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLSelectElement/multiple
func (d_ DOMHTMLSelectElement) Multiple() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("multiple"))
	return rv
}/* debug [instance_properties/getter]: multiple */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLSelectElement/multiple
func (d_ DOMHTMLSelectElement) SetMultiple(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMultiple:"), value)
}/* debug [instance_properties/setter]: multiple */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLSelectElement/name
func (d_ DOMHTMLSelectElement) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLSelectElement/name
func (d_ DOMHTMLSelectElement) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLSelectElement/options
func (d_ DOMHTMLSelectElement) Options() IDOMHTMLOptionsCollection {
	rv := objc.Send[DOMHTMLOptionsCollection](d_.ID, objc.Sel("options"))
	return rv
}/* debug [instance_properties/getter]: options */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLSelectElement/selectedIndex
func (d_ DOMHTMLSelectElement) SelectedIndex() int {
	rv := objc.Send[int](d_.ID, objc.Sel("selectedIndex"))
	return rv
}/* debug [instance_properties/getter]: selectedIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLSelectElement/selectedIndex
func (d_ DOMHTMLSelectElement) SetSelectedIndex(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSelectedIndex:"), value)
}/* debug [instance_properties/setter]: selectedIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLSelectElement/size
func (d_ DOMHTMLSelectElement) Size() int {
	rv := objc.Send[int](d_.ID, objc.Sel("size"))
	return rv
}/* debug [instance_properties/getter]: size */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLSelectElement/size
func (d_ DOMHTMLSelectElement) SetSize(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSize:"), value)
}/* debug [instance_properties/setter]: size */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLSelectElement/type
func (d_ DOMHTMLSelectElement) Type() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLSelectElement/value
func (d_ DOMHTMLSelectElement) Value() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLSelectElement/value
func (d_ DOMHTMLSelectElement) SetValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setValue:"), value)
}/* debug [instance_properties/setter]: value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLSelectElement/willValidate
func (d_ DOMHTMLSelectElement) WillValidate() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("willValidate"))
	return rv
}/* debug [instance_properties/getter]: willValidate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLSelectElement */



