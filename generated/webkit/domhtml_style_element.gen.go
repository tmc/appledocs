// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMHTMLStyleElement */


/* debug [class_header]: Header for DOMHTMLStyleElement */
// The class instance for the [DOMHTMLStyleElement] class.
var (
	DOMHTMLStyleElementClass     _DOMHTMLStyleElementClass
	DOMHTMLStyleElementClassOnce sync.Once
)

func getDOMHTMLStyleElementClass() _DOMHTMLStyleElementClass {
	DOMHTMLStyleElementClassOnce.Do(func() {
		DOMHTMLStyleElementClass = _DOMHTMLStyleElementClass{objc.GetClass("DOMHTMLStyleElement")}
	})
	return DOMHTMLStyleElementClass
}

type _DOMHTMLStyleElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLStyleElement */
// An interface definition for the [DOMHTMLStyleElement] class.
type IDOMHTMLStyleElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLStyleElement */
	// properties:
	Disabled() bool
	SetDisabled(value bool)
	Media() objc.IObject /* cross-framework: NSString */
	SetMedia(value objc.IObject /* cross-framework: NSString */)
	Sheet() IDOMStyleSheet
	Type() objc.IObject /* cross-framework: NSString */
	SetType(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLStyleElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLStyleElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLStyleElementClass) Alloc() DOMHTMLStyleElement {
	rv := objc.Send[DOMHTMLStyleElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLStyleElementClass) New() DOMHTMLStyleElement {
	rv := objc.Send[DOMHTMLStyleElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLStyleElement) Init() DOMHTMLStyleElement {
	rv := objc.Send[DOMHTMLStyleElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLStyleElement) Autorelease() DOMHTMLStyleElement {
	rv := objc.Send[DOMHTMLStyleElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLStyleElement creates a new DOMHTMLStyleElement instance.
func NewDOMHTMLStyleElement() DOMHTMLStyleElement {
	return getDOMHTMLStyleElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLStyleElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLStyleElement
type DOMHTMLStyleElement struct {
	DOMHTMLElement
}

// DOMHTMLStyleElementFrom constructs a [DOMHTMLStyleElement] from an unsafe.Pointer.
func DOMHTMLStyleElementFrom(ptr unsafe.Pointer) DOMHTMLStyleElement {
	return DOMHTMLStyleElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLStyleElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLStyleElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLStyleElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLStyleElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLStyleElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLStyleElement/disabled
func (d_ DOMHTMLStyleElement) Disabled() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("disabled"))
	return rv
}/* debug [instance_properties/getter]: disabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLStyleElement/disabled
func (d_ DOMHTMLStyleElement) SetDisabled(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDisabled:"), value)
}/* debug [instance_properties/setter]: disabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLStyleElement/media
func (d_ DOMHTMLStyleElement) Media() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("media"))
	return rv
}/* debug [instance_properties/getter]: media */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLStyleElement/media
func (d_ DOMHTMLStyleElement) SetMedia(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMedia:"), value)
}/* debug [instance_properties/setter]: media */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLStyleElement/sheet
func (d_ DOMHTMLStyleElement) Sheet() IDOMStyleSheet {
	rv := objc.Send[DOMStyleSheet](d_.ID, objc.Sel("sheet"))
	return rv
}/* debug [instance_properties/getter]: sheet */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLStyleElement/type
func (d_ DOMHTMLStyleElement) Type() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLStyleElement/type
func (d_ DOMHTMLStyleElement) SetType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setType:"), value)
}/* debug [instance_properties/setter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLStyleElement */



