// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMHTMLInputElement */

/* debug [class_header]: Header for DOMHTMLInputElement */
// The class instance for the [DOMHTMLInputElement] class.
var (
	DOMHTMLInputElementClass     _DOMHTMLInputElementClass
	DOMHTMLInputElementClassOnce sync.Once
)

func getDOMHTMLInputElementClass() _DOMHTMLInputElementClass {
	DOMHTMLInputElementClassOnce.Do(func() {
		DOMHTMLInputElementClass = _DOMHTMLInputElementClass{objc.GetClass("DOMHTMLInputElement")}
	})
	return DOMHTMLInputElementClass
}

type _DOMHTMLInputElementClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMHTMLInputElement */
// An interface definition for the [DOMHTMLInputElement] class.
type IDOMHTMLInputElement interface {
	IDOMHTMLElement

	/* debug [class_interface_properties]: Properties for DOMHTMLInputElement */
	// properties:
	AbsoluteImageURL() objc.IObject /* cross-framework: NSURL */
	Accept() objc.IObject           /* cross-framework: NSString */
	SetAccept(value objc.IObject /* cross-framework: NSString */)
	Align() objc.IObject /* cross-framework: NSString */
	SetAlign(value objc.IObject /* cross-framework: NSString */)
	Alt() objc.IObject /* cross-framework: NSString */
	SetAlt(value objc.IObject /* cross-framework: NSString */)
	AltDisplayString() objc.IObject /* cross-framework: NSString */
	Autofocus() bool
	SetAutofocus(value bool)
	Checked() bool
	SetChecked(value bool)
	DefaultChecked() bool
	SetDefaultChecked(value bool)
	DefaultValue() objc.IObject /* cross-framework: NSString */
	SetDefaultValue(value objc.IObject /* cross-framework: NSString */)
	Disabled() bool
	SetDisabled(value bool)
	Files() IDOMFileList
	SetFiles(value IDOMFileList)
	Form() IDOMHTMLFormElement
	Indeterminate() bool
	SetIndeterminate(value bool)
	MaxLength() int
	SetMaxLength(value int)
	Multiple() bool
	SetMultiple(value bool)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	ReadOnly() bool
	SetReadOnly(value bool)
	SelectionEnd() int
	SetSelectionEnd(value int)
	SelectionStart() int
	SetSelectionStart(value int)
	Size() objc.IObject /* cross-framework: NSString */
	SetSize(value objc.IObject /* cross-framework: NSString */)
	Src() objc.IObject /* cross-framework: NSString */
	SetSrc(value objc.IObject /* cross-framework: NSString */)
	Type() objc.IObject /* cross-framework: NSString */
	SetType(value objc.IObject /* cross-framework: NSString */)
	UseMap() objc.IObject /* cross-framework: NSString */
	SetUseMap(value objc.IObject /* cross-framework: NSString */)
	Value() objc.IObject /* cross-framework: NSString */
	SetValue(value objc.IObject /* cross-framework: NSString */)
	WillValidate() bool
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMHTMLInputElement */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMHTMLInputElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLInputElementClass) Alloc() DOMHTMLInputElement {
	rv := objc.Send[DOMHTMLInputElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLInputElementClass) New() DOMHTMLInputElement {
	rv := objc.Send[DOMHTMLInputElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLInputElement) Init() DOMHTMLInputElement {
	rv := objc.Send[DOMHTMLInputElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLInputElement) Autorelease() DOMHTMLInputElement {
	rv := objc.Send[DOMHTMLInputElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLInputElement creates a new DOMHTMLInputElement instance.
func NewDOMHTMLInputElement() DOMHTMLInputElement {
	return getDOMHTMLInputElementClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMHTMLInputElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement
type DOMHTMLInputElement struct {
	DOMHTMLElement
}

// DOMHTMLInputElementFrom constructs a [DOMHTMLInputElement] from an unsafe.Pointer.
func DOMHTMLInputElementFrom(ptr unsafe.Pointer) DOMHTMLInputElement {
	return DOMHTMLInputElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMHTMLInputElement */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMHTMLInputElement */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMHTMLInputElement */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMHTMLInputElement */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMHTMLInputElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/absoluteImageURL
func (d_ DOMHTMLInputElement) AbsoluteImageURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](d_.ID, objc.Sel("absoluteImageURL"))
	return rv
} /* debug [instance_properties/getter]: absoluteImageURL */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/accept
func (d_ DOMHTMLInputElement) Accept() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("accept"))
	return rv
} /* debug [instance_properties/getter]: accept */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/accept
func (d_ DOMHTMLInputElement) SetAccept(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAccept:"), value)
} /* debug [instance_properties/setter]: accept */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/align
func (d_ DOMHTMLInputElement) Align() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("align"))
	return rv
} /* debug [instance_properties/getter]: align */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/align
func (d_ DOMHTMLInputElement) SetAlign(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAlign:"), value)
} /* debug [instance_properties/setter]: align */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/alt
func (d_ DOMHTMLInputElement) Alt() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("alt"))
	return rv
} /* debug [instance_properties/getter]: alt */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/alt
func (d_ DOMHTMLInputElement) SetAlt(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAlt:"), value)
} /* debug [instance_properties/setter]: alt */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/altDisplayString
func (d_ DOMHTMLInputElement) AltDisplayString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("altDisplayString"))
	return rv
} /* debug [instance_properties/getter]: altDisplayString */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/autofocus
func (d_ DOMHTMLInputElement) Autofocus() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("autofocus"))
	return rv
} /* debug [instance_properties/getter]: autofocus */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/autofocus
func (d_ DOMHTMLInputElement) SetAutofocus(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAutofocus:"), value)
} /* debug [instance_properties/setter]: autofocus */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/checked
func (d_ DOMHTMLInputElement) Checked() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("checked"))
	return rv
} /* debug [instance_properties/getter]: checked */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/checked
func (d_ DOMHTMLInputElement) SetChecked(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setChecked:"), value)
} /* debug [instance_properties/setter]: checked */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/defaultChecked
func (d_ DOMHTMLInputElement) DefaultChecked() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("defaultChecked"))
	return rv
} /* debug [instance_properties/getter]: defaultChecked */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/defaultChecked
func (d_ DOMHTMLInputElement) SetDefaultChecked(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDefaultChecked:"), value)
} /* debug [instance_properties/setter]: defaultChecked */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/defaultValue
func (d_ DOMHTMLInputElement) DefaultValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("defaultValue"))
	return rv
} /* debug [instance_properties/getter]: defaultValue */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/defaultValue
func (d_ DOMHTMLInputElement) SetDefaultValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDefaultValue:"), value)
} /* debug [instance_properties/setter]: defaultValue */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/disabled
func (d_ DOMHTMLInputElement) Disabled() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("disabled"))
	return rv
} /* debug [instance_properties/getter]: disabled */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/disabled
func (d_ DOMHTMLInputElement) SetDisabled(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDisabled:"), value)
} /* debug [instance_properties/setter]: disabled */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/files
func (d_ DOMHTMLInputElement) Files() IDOMFileList {
	rv := objc.Send[DOMFileList](d_.ID, objc.Sel("files"))
	return rv
} /* debug [instance_properties/getter]: files */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/files
func (d_ DOMHTMLInputElement) SetFiles(value IDOMFileList) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFiles:"), value)
} /* debug [instance_properties/setter]: files */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/form
func (d_ DOMHTMLInputElement) Form() IDOMHTMLFormElement {
	rv := objc.Send[DOMHTMLFormElement](d_.ID, objc.Sel("form"))
	return rv
} /* debug [instance_properties/getter]: form */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/indeterminate
func (d_ DOMHTMLInputElement) Indeterminate() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("indeterminate"))
	return rv
} /* debug [instance_properties/getter]: indeterminate */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/indeterminate
func (d_ DOMHTMLInputElement) SetIndeterminate(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIndeterminate:"), value)
} /* debug [instance_properties/setter]: indeterminate */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/maxLength
func (d_ DOMHTMLInputElement) MaxLength() int {
	rv := objc.Send[int](d_.ID, objc.Sel("maxLength"))
	return rv
} /* debug [instance_properties/getter]: maxLength */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/maxLength
func (d_ DOMHTMLInputElement) SetMaxLength(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaxLength:"), value)
} /* debug [instance_properties/setter]: maxLength */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/multiple
func (d_ DOMHTMLInputElement) Multiple() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("multiple"))
	return rv
} /* debug [instance_properties/getter]: multiple */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/multiple
func (d_ DOMHTMLInputElement) SetMultiple(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMultiple:"), value)
} /* debug [instance_properties/setter]: multiple */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/name
func (d_ DOMHTMLInputElement) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("name"))
	return rv
} /* debug [instance_properties/getter]: name */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/name
func (d_ DOMHTMLInputElement) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setName:"), value)
} /* debug [instance_properties/setter]: name */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/readOnly
func (d_ DOMHTMLInputElement) ReadOnly() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("readOnly"))
	return rv
} /* debug [instance_properties/getter]: readOnly */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/readOnly
func (d_ DOMHTMLInputElement) SetReadOnly(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setReadOnly:"), value)
} /* debug [instance_properties/setter]: readOnly */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/selectionEnd
func (d_ DOMHTMLInputElement) SelectionEnd() int {
	rv := objc.Send[int](d_.ID, objc.Sel("selectionEnd"))
	return rv
} /* debug [instance_properties/getter]: selectionEnd */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/selectionEnd
func (d_ DOMHTMLInputElement) SetSelectionEnd(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSelectionEnd:"), value)
} /* debug [instance_properties/setter]: selectionEnd */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/selectionStart
func (d_ DOMHTMLInputElement) SelectionStart() int {
	rv := objc.Send[int](d_.ID, objc.Sel("selectionStart"))
	return rv
} /* debug [instance_properties/getter]: selectionStart */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/selectionStart
func (d_ DOMHTMLInputElement) SetSelectionStart(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSelectionStart:"), value)
} /* debug [instance_properties/setter]: selectionStart */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/size
func (d_ DOMHTMLInputElement) Size() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("size"))
	return rv
} /* debug [instance_properties/getter]: size */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/size
func (d_ DOMHTMLInputElement) SetSize(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSize:"), value)
} /* debug [instance_properties/setter]: size */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/src
func (d_ DOMHTMLInputElement) Src() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("src"))
	return rv
} /* debug [instance_properties/getter]: src */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/src
func (d_ DOMHTMLInputElement) SetSrc(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSrc:"), value)
} /* debug [instance_properties/setter]: src */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/type
func (d_ DOMHTMLInputElement) Type() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("type"))
	return rv
} /* debug [instance_properties/getter]: type */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/type
func (d_ DOMHTMLInputElement) SetType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setType:"), value)
} /* debug [instance_properties/setter]: type */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/useMap
func (d_ DOMHTMLInputElement) UseMap() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("useMap"))
	return rv
} /* debug [instance_properties/getter]: useMap */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/useMap
func (d_ DOMHTMLInputElement) SetUseMap(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setUseMap:"), value)
} /* debug [instance_properties/setter]: useMap */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/value
func (d_ DOMHTMLInputElement) Value() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("value"))
	return rv
} /* debug [instance_properties/getter]: value */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/value
func (d_ DOMHTMLInputElement) SetValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setValue:"), value)
} /* debug [instance_properties/setter]: value */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLInputElement/willValidate
func (d_ DOMHTMLInputElement) WillValidate() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("willValidate"))
	return rv
} /* debug [instance_properties/getter]: willValidate */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMHTMLInputElement */
