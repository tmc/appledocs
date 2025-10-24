// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMHTMLAppletElement */


/* debug [class_header]: Header for DOMHTMLAppletElement */
// The class instance for the [DOMHTMLAppletElement] class.
var (
	DOMHTMLAppletElementClass     _DOMHTMLAppletElementClass
	DOMHTMLAppletElementClassOnce sync.Once
)

func getDOMHTMLAppletElementClass() _DOMHTMLAppletElementClass {
	DOMHTMLAppletElementClassOnce.Do(func() {
		DOMHTMLAppletElementClass = _DOMHTMLAppletElementClass{objc.GetClass("DOMHTMLAppletElement")}
	})
	return DOMHTMLAppletElementClass
}

type _DOMHTMLAppletElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLAppletElement */
// An interface definition for the [DOMHTMLAppletElement] class.
type IDOMHTMLAppletElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLAppletElement */
	// properties:
	Align() objc.IObject /* cross-framework: NSString */
	SetAlign(value objc.IObject /* cross-framework: NSString */)
	Alt() objc.IObject /* cross-framework: NSString */
	SetAlt(value objc.IObject /* cross-framework: NSString */)
	Archive() objc.IObject /* cross-framework: NSString */
	SetArchive(value objc.IObject /* cross-framework: NSString */)
	Code() objc.IObject /* cross-framework: NSString */
	SetCode(value objc.IObject /* cross-framework: NSString */)
	CodeBase() objc.IObject /* cross-framework: NSString */
	SetCodeBase(value objc.IObject /* cross-framework: NSString */)
	Height() objc.IObject /* cross-framework: NSString */
	SetHeight(value objc.IObject /* cross-framework: NSString */)
	Hspace() int
	SetHspace(value int)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	GetObject() objc.IObject /* cross-framework: NSString */
	SetGetObject(value objc.IObject /* cross-framework: NSString */)
	Vspace() int
	SetVspace(value int)
	Width() objc.IObject /* cross-framework: NSString */
	SetWidth(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLAppletElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLAppletElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLAppletElementClass) Alloc() DOMHTMLAppletElement {
	rv := objc.Send[DOMHTMLAppletElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLAppletElementClass) New() DOMHTMLAppletElement {
	rv := objc.Send[DOMHTMLAppletElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLAppletElement) Init() DOMHTMLAppletElement {
	rv := objc.Send[DOMHTMLAppletElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLAppletElement) Autorelease() DOMHTMLAppletElement {
	rv := objc.Send[DOMHTMLAppletElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLAppletElement creates a new DOMHTMLAppletElement instance.
func NewDOMHTMLAppletElement() DOMHTMLAppletElement {
	return getDOMHTMLAppletElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLAppletElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAppletElement
type DOMHTMLAppletElement struct {
	DOMHTMLElement
}

// DOMHTMLAppletElementFrom constructs a [DOMHTMLAppletElement] from an unsafe.Pointer.
func DOMHTMLAppletElementFrom(ptr unsafe.Pointer) DOMHTMLAppletElement {
	return DOMHTMLAppletElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLAppletElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLAppletElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLAppletElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLAppletElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLAppletElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAppletElement/align
func (d_ DOMHTMLAppletElement) Align() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("align"))
	return rv
}/* debug [instance_properties/getter]: align */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAppletElement/align
func (d_ DOMHTMLAppletElement) SetAlign(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAlign:"), value)
}/* debug [instance_properties/setter]: align */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAppletElement/alt
func (d_ DOMHTMLAppletElement) Alt() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("alt"))
	return rv
}/* debug [instance_properties/getter]: alt */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAppletElement/alt
func (d_ DOMHTMLAppletElement) SetAlt(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAlt:"), value)
}/* debug [instance_properties/setter]: alt */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAppletElement/archive
func (d_ DOMHTMLAppletElement) Archive() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("archive"))
	return rv
}/* debug [instance_properties/getter]: archive */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAppletElement/archive
func (d_ DOMHTMLAppletElement) SetArchive(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setArchive:"), value)
}/* debug [instance_properties/setter]: archive */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAppletElement/code
func (d_ DOMHTMLAppletElement) Code() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("code"))
	return rv
}/* debug [instance_properties/getter]: code */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAppletElement/code
func (d_ DOMHTMLAppletElement) SetCode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCode:"), value)
}/* debug [instance_properties/setter]: code */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAppletElement/codeBase
func (d_ DOMHTMLAppletElement) CodeBase() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("codeBase"))
	return rv
}/* debug [instance_properties/getter]: codeBase */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAppletElement/codeBase
func (d_ DOMHTMLAppletElement) SetCodeBase(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCodeBase:"), value)
}/* debug [instance_properties/setter]: codeBase */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAppletElement/height
func (d_ DOMHTMLAppletElement) Height() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("height"))
	return rv
}/* debug [instance_properties/getter]: height */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAppletElement/height
func (d_ DOMHTMLAppletElement) SetHeight(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHeight:"), value)
}/* debug [instance_properties/setter]: height */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAppletElement/hspace
func (d_ DOMHTMLAppletElement) Hspace() int {
	rv := objc.Send[int](d_.ID, objc.Sel("hspace"))
	return rv
}/* debug [instance_properties/getter]: hspace */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAppletElement/hspace
func (d_ DOMHTMLAppletElement) SetHspace(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHspace:"), value)
}/* debug [instance_properties/setter]: hspace */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAppletElement/name
func (d_ DOMHTMLAppletElement) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAppletElement/name
func (d_ DOMHTMLAppletElement) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAppletElement/object
func (d_ DOMHTMLAppletElement) GetObject() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("object"))
	return rv
}/* debug [instance_properties/getter]: object */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAppletElement/object
func (d_ DOMHTMLAppletElement) SetGetObject(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setGetObject:"), value)
}/* debug [instance_properties/setter]: object */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAppletElement/vspace
func (d_ DOMHTMLAppletElement) Vspace() int {
	rv := objc.Send[int](d_.ID, objc.Sel("vspace"))
	return rv
}/* debug [instance_properties/getter]: vspace */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAppletElement/vspace
func (d_ DOMHTMLAppletElement) SetVspace(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setVspace:"), value)
}/* debug [instance_properties/setter]: vspace */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAppletElement/width
func (d_ DOMHTMLAppletElement) Width() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("width"))
	return rv
}/* debug [instance_properties/getter]: width */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLAppletElement/width
func (d_ DOMHTMLAppletElement) SetWidth(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWidth:"), value)
}/* debug [instance_properties/setter]: width */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLAppletElement */



