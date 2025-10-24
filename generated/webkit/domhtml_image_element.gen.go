// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMHTMLImageElement */


/* debug [class_header]: Header for DOMHTMLImageElement */
// The class instance for the [DOMHTMLImageElement] class.
var (
	DOMHTMLImageElementClass     _DOMHTMLImageElementClass
	DOMHTMLImageElementClassOnce sync.Once
)

func getDOMHTMLImageElementClass() _DOMHTMLImageElementClass {
	DOMHTMLImageElementClassOnce.Do(func() {
		DOMHTMLImageElementClass = _DOMHTMLImageElementClass{objc.GetClass("DOMHTMLImageElement")}
	})
	return DOMHTMLImageElementClass
}

type _DOMHTMLImageElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLImageElement */
// An interface definition for the [DOMHTMLImageElement] class.
type IDOMHTMLImageElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLImageElement */
	// properties:
	AbsoluteImageURL() objc.IObject /* cross-framework: NSURL */
	Align() objc.IObject /* cross-framework: NSString */
	SetAlign(value objc.IObject /* cross-framework: NSString */)
	Alt() objc.IObject /* cross-framework: NSString */
	SetAlt(value objc.IObject /* cross-framework: NSString */)
	AltDisplayString() objc.IObject /* cross-framework: NSString */
	Border() objc.IObject /* cross-framework: NSString */
	SetBorder(value objc.IObject /* cross-framework: NSString */)
	Complete() bool
	Height() int
	SetHeight(value int)
	Hspace() int
	SetHspace(value int)
	IsMap() bool
	SetIsMap(value bool)
	LongDesc() objc.IObject /* cross-framework: NSString */
	SetLongDesc(value objc.IObject /* cross-framework: NSString */)
	Lowsrc() objc.IObject /* cross-framework: NSString */
	SetLowsrc(value objc.IObject /* cross-framework: NSString */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	NaturalHeight() int
	NaturalWidth() int
	Src() objc.IObject /* cross-framework: NSString */
	SetSrc(value objc.IObject /* cross-framework: NSString */)
	UseMap() objc.IObject /* cross-framework: NSString */
	SetUseMap(value objc.IObject /* cross-framework: NSString */)
	Vspace() int
	SetVspace(value int)
	Width() int
	SetWidth(value int)
	X() int
	Y() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLImageElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLImageElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLImageElementClass) Alloc() DOMHTMLImageElement {
	rv := objc.Send[DOMHTMLImageElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLImageElementClass) New() DOMHTMLImageElement {
	rv := objc.Send[DOMHTMLImageElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLImageElement) Init() DOMHTMLImageElement {
	rv := objc.Send[DOMHTMLImageElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLImageElement) Autorelease() DOMHTMLImageElement {
	rv := objc.Send[DOMHTMLImageElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLImageElement creates a new DOMHTMLImageElement instance.
func NewDOMHTMLImageElement() DOMHTMLImageElement {
	return getDOMHTMLImageElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLImageElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement
type DOMHTMLImageElement struct {
	DOMHTMLElement
}

// DOMHTMLImageElementFrom constructs a [DOMHTMLImageElement] from an unsafe.Pointer.
func DOMHTMLImageElementFrom(ptr unsafe.Pointer) DOMHTMLImageElement {
	return DOMHTMLImageElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLImageElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLImageElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLImageElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLImageElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLImageElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/absoluteImageURL
func (d_ DOMHTMLImageElement) AbsoluteImageURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](d_.ID, objc.Sel("absoluteImageURL"))
	return rv
}/* debug [instance_properties/getter]: absoluteImageURL */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/align
func (d_ DOMHTMLImageElement) Align() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("align"))
	return rv
}/* debug [instance_properties/getter]: align */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/align
func (d_ DOMHTMLImageElement) SetAlign(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAlign:"), value)
}/* debug [instance_properties/setter]: align */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/alt
func (d_ DOMHTMLImageElement) Alt() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("alt"))
	return rv
}/* debug [instance_properties/getter]: alt */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/alt
func (d_ DOMHTMLImageElement) SetAlt(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAlt:"), value)
}/* debug [instance_properties/setter]: alt */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/altDisplayString
func (d_ DOMHTMLImageElement) AltDisplayString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("altDisplayString"))
	return rv
}/* debug [instance_properties/getter]: altDisplayString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/border
func (d_ DOMHTMLImageElement) Border() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("border"))
	return rv
}/* debug [instance_properties/getter]: border */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/border
func (d_ DOMHTMLImageElement) SetBorder(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBorder:"), value)
}/* debug [instance_properties/setter]: border */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/complete
func (d_ DOMHTMLImageElement) Complete() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("complete"))
	return rv
}/* debug [instance_properties/getter]: complete */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/height
func (d_ DOMHTMLImageElement) Height() int {
	rv := objc.Send[int](d_.ID, objc.Sel("height"))
	return rv
}/* debug [instance_properties/getter]: height */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/height
func (d_ DOMHTMLImageElement) SetHeight(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHeight:"), value)
}/* debug [instance_properties/setter]: height */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/hspace
func (d_ DOMHTMLImageElement) Hspace() int {
	rv := objc.Send[int](d_.ID, objc.Sel("hspace"))
	return rv
}/* debug [instance_properties/getter]: hspace */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/hspace
func (d_ DOMHTMLImageElement) SetHspace(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHspace:"), value)
}/* debug [instance_properties/setter]: hspace */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/isMap
func (d_ DOMHTMLImageElement) IsMap() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isMap"))
	return rv
}/* debug [instance_properties/getter]: isMap */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/isMap
func (d_ DOMHTMLImageElement) SetIsMap(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsMap:"), value)
}/* debug [instance_properties/setter]: isMap */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/longDesc
func (d_ DOMHTMLImageElement) LongDesc() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("longDesc"))
	return rv
}/* debug [instance_properties/getter]: longDesc */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/longDesc
func (d_ DOMHTMLImageElement) SetLongDesc(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLongDesc:"), value)
}/* debug [instance_properties/setter]: longDesc */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/lowsrc
func (d_ DOMHTMLImageElement) Lowsrc() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("lowsrc"))
	return rv
}/* debug [instance_properties/getter]: lowsrc */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/lowsrc
func (d_ DOMHTMLImageElement) SetLowsrc(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLowsrc:"), value)
}/* debug [instance_properties/setter]: lowsrc */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/name
func (d_ DOMHTMLImageElement) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/name
func (d_ DOMHTMLImageElement) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/naturalHeight
func (d_ DOMHTMLImageElement) NaturalHeight() int {
	rv := objc.Send[int](d_.ID, objc.Sel("naturalHeight"))
	return rv
}/* debug [instance_properties/getter]: naturalHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/naturalWidth
func (d_ DOMHTMLImageElement) NaturalWidth() int {
	rv := objc.Send[int](d_.ID, objc.Sel("naturalWidth"))
	return rv
}/* debug [instance_properties/getter]: naturalWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/src
func (d_ DOMHTMLImageElement) Src() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("src"))
	return rv
}/* debug [instance_properties/getter]: src */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/src
func (d_ DOMHTMLImageElement) SetSrc(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSrc:"), value)
}/* debug [instance_properties/setter]: src */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/useMap
func (d_ DOMHTMLImageElement) UseMap() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("useMap"))
	return rv
}/* debug [instance_properties/getter]: useMap */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/useMap
func (d_ DOMHTMLImageElement) SetUseMap(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setUseMap:"), value)
}/* debug [instance_properties/setter]: useMap */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/vspace
func (d_ DOMHTMLImageElement) Vspace() int {
	rv := objc.Send[int](d_.ID, objc.Sel("vspace"))
	return rv
}/* debug [instance_properties/getter]: vspace */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/vspace
func (d_ DOMHTMLImageElement) SetVspace(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setVspace:"), value)
}/* debug [instance_properties/setter]: vspace */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/width
func (d_ DOMHTMLImageElement) Width() int {
	rv := objc.Send[int](d_.ID, objc.Sel("width"))
	return rv
}/* debug [instance_properties/getter]: width */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/width
func (d_ DOMHTMLImageElement) SetWidth(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWidth:"), value)
}/* debug [instance_properties/setter]: width */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/x
func (d_ DOMHTMLImageElement) X() int {
	rv := objc.Send[int](d_.ID, objc.Sel("x"))
	return rv
}/* debug [instance_properties/getter]: x */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLImageElement/y
func (d_ DOMHTMLImageElement) Y() int {
	rv := objc.Send[int](d_.ID, objc.Sel("y"))
	return rv
}/* debug [instance_properties/getter]: y */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLImageElement */



