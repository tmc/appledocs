// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLTextAreaElement */

/* debug [class_header]: Header for DOMHTMLTextAreaElement */
// The class instance for the [DOMHTMLTextAreaElement] class.
var (
	DOMHTMLTextAreaElementClass     _DOMHTMLTextAreaElementClass
	DOMHTMLTextAreaElementClassOnce sync.Once
)

func getDOMHTMLTextAreaElementClass() _DOMHTMLTextAreaElementClass {
	DOMHTMLTextAreaElementClassOnce.Do(func() {
		DOMHTMLTextAreaElementClass = _DOMHTMLTextAreaElementClass{objc.GetClass("DOMHTMLTextAreaElement")}
	})
	return DOMHTMLTextAreaElementClass
}

type _DOMHTMLTextAreaElementClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMHTMLTextAreaElement */
// An interface definition for the [DOMHTMLTextAreaElement] class.
type IDOMHTMLTextAreaElement interface {
	IDOMHTMLElement

	/* debug [class_interface_properties]: Properties for DOMHTMLTextAreaElement */
	// properties:
	Autofocus() bool
	SetAutofocus(value bool)
	Cols() int
	SetCols(value int)
	DefaultValue() objc.IObject /* cross-framework: NSString */
	SetDefaultValue(value objc.IObject /* cross-framework: NSString */)
	Disabled() bool
	SetDisabled(value bool)
	Form() IDOMHTMLFormElement
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	ReadOnly() bool
	SetReadOnly(value bool)
	Rows() int
	SetRows(value int)
	SelectionEnd() int
	SetSelectionEnd(value int)
	SelectionStart() int
	SetSelectionStart(value int)
	Type() objc.IObject  /* cross-framework: NSString */
	Value() objc.IObject /* cross-framework: NSString */
	SetValue(value objc.IObject /* cross-framework: NSString */)
	WillValidate() bool
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMHTMLTextAreaElement */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMHTMLTextAreaElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLTextAreaElementClass) Alloc() DOMHTMLTextAreaElement {
	rv := objc.Send[DOMHTMLTextAreaElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLTextAreaElementClass) New() DOMHTMLTextAreaElement {
	rv := objc.Send[DOMHTMLTextAreaElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLTextAreaElement) Init() DOMHTMLTextAreaElement {
	rv := objc.Send[DOMHTMLTextAreaElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLTextAreaElement) Autorelease() DOMHTMLTextAreaElement {
	rv := objc.Send[DOMHTMLTextAreaElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLTextAreaElement creates a new DOMHTMLTextAreaElement instance.
func NewDOMHTMLTextAreaElement() DOMHTMLTextAreaElement {
	return getDOMHTMLTextAreaElementClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMHTMLTextAreaElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTextAreaElement
type DOMHTMLTextAreaElement struct {
	DOMHTMLElement
}

// DOMHTMLTextAreaElementFrom constructs a [DOMHTMLTextAreaElement] from an unsafe.Pointer.
func DOMHTMLTextAreaElementFrom(ptr unsafe.Pointer) DOMHTMLTextAreaElement {
	return DOMHTMLTextAreaElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMHTMLTextAreaElement */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMHTMLTextAreaElement */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMHTMLTextAreaElement */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMHTMLTextAreaElement */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMHTMLTextAreaElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTextAreaElement/autofocus
func (d_ DOMHTMLTextAreaElement) Autofocus() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("autofocus"))
	return rv
} /* debug [instance_properties/getter]: autofocus */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTextAreaElement/autofocus
func (d_ DOMHTMLTextAreaElement) SetAutofocus(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAutofocus:"), value)
} /* debug [instance_properties/setter]: autofocus */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTextAreaElement/cols
func (d_ DOMHTMLTextAreaElement) Cols() int {
	rv := objc.Send[int](d_.ID, objc.Sel("cols"))
	return rv
} /* debug [instance_properties/getter]: cols */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTextAreaElement/cols
func (d_ DOMHTMLTextAreaElement) SetCols(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCols:"), value)
} /* debug [instance_properties/setter]: cols */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTextAreaElement/defaultValue
func (d_ DOMHTMLTextAreaElement) DefaultValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("defaultValue"))
	return rv
} /* debug [instance_properties/getter]: defaultValue */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTextAreaElement/defaultValue
func (d_ DOMHTMLTextAreaElement) SetDefaultValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDefaultValue:"), value)
} /* debug [instance_properties/setter]: defaultValue */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTextAreaElement/disabled
func (d_ DOMHTMLTextAreaElement) Disabled() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("disabled"))
	return rv
} /* debug [instance_properties/getter]: disabled */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTextAreaElement/disabled
func (d_ DOMHTMLTextAreaElement) SetDisabled(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDisabled:"), value)
} /* debug [instance_properties/setter]: disabled */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTextAreaElement/form
func (d_ DOMHTMLTextAreaElement) Form() IDOMHTMLFormElement {
	rv := objc.Send[DOMHTMLFormElement](d_.ID, objc.Sel("form"))
	return rv
} /* debug [instance_properties/getter]: form */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTextAreaElement/name
func (d_ DOMHTMLTextAreaElement) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("name"))
	return rv
} /* debug [instance_properties/getter]: name */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTextAreaElement/name
func (d_ DOMHTMLTextAreaElement) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setName:"), value)
} /* debug [instance_properties/setter]: name */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTextAreaElement/readOnly
func (d_ DOMHTMLTextAreaElement) ReadOnly() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("readOnly"))
	return rv
} /* debug [instance_properties/getter]: readOnly */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTextAreaElement/readOnly
func (d_ DOMHTMLTextAreaElement) SetReadOnly(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setReadOnly:"), value)
} /* debug [instance_properties/setter]: readOnly */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTextAreaElement/rows
func (d_ DOMHTMLTextAreaElement) Rows() int {
	rv := objc.Send[int](d_.ID, objc.Sel("rows"))
	return rv
} /* debug [instance_properties/getter]: rows */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTextAreaElement/rows
func (d_ DOMHTMLTextAreaElement) SetRows(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setRows:"), value)
} /* debug [instance_properties/setter]: rows */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTextAreaElement/selectionEnd
func (d_ DOMHTMLTextAreaElement) SelectionEnd() int {
	rv := objc.Send[int](d_.ID, objc.Sel("selectionEnd"))
	return rv
} /* debug [instance_properties/getter]: selectionEnd */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTextAreaElement/selectionEnd
func (d_ DOMHTMLTextAreaElement) SetSelectionEnd(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSelectionEnd:"), value)
} /* debug [instance_properties/setter]: selectionEnd */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTextAreaElement/selectionStart
func (d_ DOMHTMLTextAreaElement) SelectionStart() int {
	rv := objc.Send[int](d_.ID, objc.Sel("selectionStart"))
	return rv
} /* debug [instance_properties/getter]: selectionStart */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTextAreaElement/selectionStart
func (d_ DOMHTMLTextAreaElement) SetSelectionStart(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSelectionStart:"), value)
} /* debug [instance_properties/setter]: selectionStart */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTextAreaElement/type
func (d_ DOMHTMLTextAreaElement) Type() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("type"))
	return rv
} /* debug [instance_properties/getter]: type */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTextAreaElement/value
func (d_ DOMHTMLTextAreaElement) Value() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("value"))
	return rv
} /* debug [instance_properties/getter]: value */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTextAreaElement/value
func (d_ DOMHTMLTextAreaElement) SetValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setValue:"), value)
} /* debug [instance_properties/setter]: value */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLTextAreaElement/willValidate
func (d_ DOMHTMLTextAreaElement) WillValidate() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("willValidate"))
	return rv
} /* debug [instance_properties/getter]: willValidate */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMHTMLTextAreaElement */
