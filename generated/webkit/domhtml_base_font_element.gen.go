// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMHTMLBaseFontElement */


/* debug [class_header]: Header for DOMHTMLBaseFontElement */
// The class instance for the [DOMHTMLBaseFontElement] class.
var (
	DOMHTMLBaseFontElementClass     _DOMHTMLBaseFontElementClass
	DOMHTMLBaseFontElementClassOnce sync.Once
)

func getDOMHTMLBaseFontElementClass() _DOMHTMLBaseFontElementClass {
	DOMHTMLBaseFontElementClassOnce.Do(func() {
		DOMHTMLBaseFontElementClass = _DOMHTMLBaseFontElementClass{objc.GetClass("DOMHTMLBaseFontElement")}
	})
	return DOMHTMLBaseFontElementClass
}

type _DOMHTMLBaseFontElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLBaseFontElement */
// An interface definition for the [DOMHTMLBaseFontElement] class.
type IDOMHTMLBaseFontElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLBaseFontElement */
	// properties:
	Color() objc.IObject /* cross-framework: NSString */
	SetColor(value objc.IObject /* cross-framework: NSString */)
	Face() objc.IObject /* cross-framework: NSString */
	SetFace(value objc.IObject /* cross-framework: NSString */)
	Size() objc.IObject /* cross-framework: NSString */
	SetSize(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLBaseFontElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLBaseFontElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLBaseFontElementClass) Alloc() DOMHTMLBaseFontElement {
	rv := objc.Send[DOMHTMLBaseFontElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLBaseFontElementClass) New() DOMHTMLBaseFontElement {
	rv := objc.Send[DOMHTMLBaseFontElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLBaseFontElement) Init() DOMHTMLBaseFontElement {
	rv := objc.Send[DOMHTMLBaseFontElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLBaseFontElement) Autorelease() DOMHTMLBaseFontElement {
	rv := objc.Send[DOMHTMLBaseFontElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLBaseFontElement creates a new DOMHTMLBaseFontElement instance.
func NewDOMHTMLBaseFontElement() DOMHTMLBaseFontElement {
	return getDOMHTMLBaseFontElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLBaseFontElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBaseFontElement
type DOMHTMLBaseFontElement struct {
	DOMHTMLElement
}

// DOMHTMLBaseFontElementFrom constructs a [DOMHTMLBaseFontElement] from an unsafe.Pointer.
func DOMHTMLBaseFontElementFrom(ptr unsafe.Pointer) DOMHTMLBaseFontElement {
	return DOMHTMLBaseFontElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLBaseFontElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLBaseFontElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLBaseFontElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLBaseFontElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLBaseFontElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBaseFontElement/color
func (d_ DOMHTMLBaseFontElement) Color() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("color"))
	return rv
}/* debug [instance_properties/getter]: color */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBaseFontElement/color
func (d_ DOMHTMLBaseFontElement) SetColor(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setColor:"), value)
}/* debug [instance_properties/setter]: color */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBaseFontElement/face
func (d_ DOMHTMLBaseFontElement) Face() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("face"))
	return rv
}/* debug [instance_properties/getter]: face */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBaseFontElement/face
func (d_ DOMHTMLBaseFontElement) SetFace(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFace:"), value)
}/* debug [instance_properties/setter]: face */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBaseFontElement/size
func (d_ DOMHTMLBaseFontElement) Size() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("size"))
	return rv
}/* debug [instance_properties/getter]: size */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBaseFontElement/size
func (d_ DOMHTMLBaseFontElement) SetSize(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSize:"), value)
}/* debug [instance_properties/setter]: size */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLBaseFontElement */



