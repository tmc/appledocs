// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMHTMLObjectElement */


/* debug [class_header]: Header for DOMHTMLObjectElement */
// The class instance for the [DOMHTMLObjectElement] class.
var (
	DOMHTMLObjectElementClass     _DOMHTMLObjectElementClass
	DOMHTMLObjectElementClassOnce sync.Once
)

func getDOMHTMLObjectElementClass() _DOMHTMLObjectElementClass {
	DOMHTMLObjectElementClassOnce.Do(func() {
		DOMHTMLObjectElementClass = _DOMHTMLObjectElementClass{objc.GetClass("DOMHTMLObjectElement")}
	})
	return DOMHTMLObjectElementClass
}

type _DOMHTMLObjectElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLObjectElement */
// An interface definition for the [DOMHTMLObjectElement] class.
type IDOMHTMLObjectElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLObjectElement */
	// properties:
	AbsoluteImageURL() objc.IObject /* cross-framework: NSURL */
	Align() objc.IObject /* cross-framework: NSString */
	SetAlign(value objc.IObject /* cross-framework: NSString */)
	Archive() objc.IObject /* cross-framework: NSString */
	SetArchive(value objc.IObject /* cross-framework: NSString */)
	Border() objc.IObject /* cross-framework: NSString */
	SetBorder(value objc.IObject /* cross-framework: NSString */)
	Code() objc.IObject /* cross-framework: NSString */
	SetCode(value objc.IObject /* cross-framework: NSString */)
	CodeBase() objc.IObject /* cross-framework: NSString */
	SetCodeBase(value objc.IObject /* cross-framework: NSString */)
	CodeType() objc.IObject /* cross-framework: NSString */
	SetCodeType(value objc.IObject /* cross-framework: NSString */)
	ContentDocument() IDOMDocument
	ContentFrame() IWebFrame
	Data() objc.IObject /* cross-framework: NSString */
	SetData(value objc.IObject /* cross-framework: NSString */)
	Declare() bool
	SetDeclare(value bool)
	Form() IDOMHTMLFormElement
	Height() objc.IObject /* cross-framework: NSString */
	SetHeight(value objc.IObject /* cross-framework: NSString */)
	Hspace() int
	SetHspace(value int)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Standby() objc.IObject /* cross-framework: NSString */
	SetStandby(value objc.IObject /* cross-framework: NSString */)
	Type() objc.IObject /* cross-framework: NSString */
	SetType(value objc.IObject /* cross-framework: NSString */)
	UseMap() objc.IObject /* cross-framework: NSString */
	SetUseMap(value objc.IObject /* cross-framework: NSString */)
	Vspace() int
	SetVspace(value int)
	Width() objc.IObject /* cross-framework: NSString */
	SetWidth(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLObjectElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLObjectElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLObjectElementClass) Alloc() DOMHTMLObjectElement {
	rv := objc.Send[DOMHTMLObjectElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLObjectElementClass) New() DOMHTMLObjectElement {
	rv := objc.Send[DOMHTMLObjectElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLObjectElement) Init() DOMHTMLObjectElement {
	rv := objc.Send[DOMHTMLObjectElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLObjectElement) Autorelease() DOMHTMLObjectElement {
	rv := objc.Send[DOMHTMLObjectElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLObjectElement creates a new DOMHTMLObjectElement instance.
func NewDOMHTMLObjectElement() DOMHTMLObjectElement {
	return getDOMHTMLObjectElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLObjectElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement
type DOMHTMLObjectElement struct {
	DOMHTMLElement
}

// DOMHTMLObjectElementFrom constructs a [DOMHTMLObjectElement] from an unsafe.Pointer.
func DOMHTMLObjectElementFrom(ptr unsafe.Pointer) DOMHTMLObjectElement {
	return DOMHTMLObjectElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLObjectElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLObjectElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLObjectElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLObjectElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLObjectElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/absoluteImageURL
func (d_ DOMHTMLObjectElement) AbsoluteImageURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](d_.ID, objc.Sel("absoluteImageURL"))
	return rv
}/* debug [instance_properties/getter]: absoluteImageURL */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/align
func (d_ DOMHTMLObjectElement) Align() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("align"))
	return rv
}/* debug [instance_properties/getter]: align */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/align
func (d_ DOMHTMLObjectElement) SetAlign(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAlign:"), value)
}/* debug [instance_properties/setter]: align */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/archive
func (d_ DOMHTMLObjectElement) Archive() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("archive"))
	return rv
}/* debug [instance_properties/getter]: archive */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/archive
func (d_ DOMHTMLObjectElement) SetArchive(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setArchive:"), value)
}/* debug [instance_properties/setter]: archive */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/border
func (d_ DOMHTMLObjectElement) Border() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("border"))
	return rv
}/* debug [instance_properties/getter]: border */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/border
func (d_ DOMHTMLObjectElement) SetBorder(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBorder:"), value)
}/* debug [instance_properties/setter]: border */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/code
func (d_ DOMHTMLObjectElement) Code() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("code"))
	return rv
}/* debug [instance_properties/getter]: code */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/code
func (d_ DOMHTMLObjectElement) SetCode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCode:"), value)
}/* debug [instance_properties/setter]: code */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/codeBase
func (d_ DOMHTMLObjectElement) CodeBase() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("codeBase"))
	return rv
}/* debug [instance_properties/getter]: codeBase */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/codeBase
func (d_ DOMHTMLObjectElement) SetCodeBase(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCodeBase:"), value)
}/* debug [instance_properties/setter]: codeBase */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/codeType
func (d_ DOMHTMLObjectElement) CodeType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("codeType"))
	return rv
}/* debug [instance_properties/getter]: codeType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/codeType
func (d_ DOMHTMLObjectElement) SetCodeType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCodeType:"), value)
}/* debug [instance_properties/setter]: codeType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/contentDocument
func (d_ DOMHTMLObjectElement) ContentDocument() IDOMDocument {
	rv := objc.Send[DOMDocument](d_.ID, objc.Sel("contentDocument"))
	return rv
}/* debug [instance_properties/getter]: contentDocument */


// The content frame of the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/contentFrame
func (d_ DOMHTMLObjectElement) ContentFrame() IWebFrame {
	rv := objc.Send[WebFrame](d_.ID, objc.Sel("contentFrame"))
	return rv
}/* debug [instance_properties/getter]: contentFrame */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/data
func (d_ DOMHTMLObjectElement) Data() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/data
func (d_ DOMHTMLObjectElement) SetData(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setData:"), value)
}/* debug [instance_properties/setter]: data */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/declare
func (d_ DOMHTMLObjectElement) Declare() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("declare"))
	return rv
}/* debug [instance_properties/getter]: declare */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/declare
func (d_ DOMHTMLObjectElement) SetDeclare(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDeclare:"), value)
}/* debug [instance_properties/setter]: declare */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/form
func (d_ DOMHTMLObjectElement) Form() IDOMHTMLFormElement {
	rv := objc.Send[DOMHTMLFormElement](d_.ID, objc.Sel("form"))
	return rv
}/* debug [instance_properties/getter]: form */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/height
func (d_ DOMHTMLObjectElement) Height() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("height"))
	return rv
}/* debug [instance_properties/getter]: height */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/height
func (d_ DOMHTMLObjectElement) SetHeight(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHeight:"), value)
}/* debug [instance_properties/setter]: height */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/hspace
func (d_ DOMHTMLObjectElement) Hspace() int {
	rv := objc.Send[int](d_.ID, objc.Sel("hspace"))
	return rv
}/* debug [instance_properties/getter]: hspace */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/hspace
func (d_ DOMHTMLObjectElement) SetHspace(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHspace:"), value)
}/* debug [instance_properties/setter]: hspace */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/name
func (d_ DOMHTMLObjectElement) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/name
func (d_ DOMHTMLObjectElement) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/standby
func (d_ DOMHTMLObjectElement) Standby() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("standby"))
	return rv
}/* debug [instance_properties/getter]: standby */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/standby
func (d_ DOMHTMLObjectElement) SetStandby(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setStandby:"), value)
}/* debug [instance_properties/setter]: standby */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/type
func (d_ DOMHTMLObjectElement) Type() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/type
func (d_ DOMHTMLObjectElement) SetType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setType:"), value)
}/* debug [instance_properties/setter]: type */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/useMap
func (d_ DOMHTMLObjectElement) UseMap() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("useMap"))
	return rv
}/* debug [instance_properties/getter]: useMap */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/useMap
func (d_ DOMHTMLObjectElement) SetUseMap(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setUseMap:"), value)
}/* debug [instance_properties/setter]: useMap */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/vspace
func (d_ DOMHTMLObjectElement) Vspace() int {
	rv := objc.Send[int](d_.ID, objc.Sel("vspace"))
	return rv
}/* debug [instance_properties/getter]: vspace */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/vspace
func (d_ DOMHTMLObjectElement) SetVspace(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setVspace:"), value)
}/* debug [instance_properties/setter]: vspace */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/width
func (d_ DOMHTMLObjectElement) Width() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("width"))
	return rv
}/* debug [instance_properties/getter]: width */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLObjectElement/width
func (d_ DOMHTMLObjectElement) SetWidth(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWidth:"), value)
}/* debug [instance_properties/setter]: width */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLObjectElement */



