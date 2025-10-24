// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMHTMLFontElement */


/* debug [class_header]: Header for DOMHTMLFontElement */
// The class instance for the [DOMHTMLFontElement] class.
var (
	DOMHTMLFontElementClass     _DOMHTMLFontElementClass
	DOMHTMLFontElementClassOnce sync.Once
)

func getDOMHTMLFontElementClass() _DOMHTMLFontElementClass {
	DOMHTMLFontElementClassOnce.Do(func() {
		DOMHTMLFontElementClass = _DOMHTMLFontElementClass{objc.GetClass("DOMHTMLFontElement")}
	})
	return DOMHTMLFontElementClass
}

type _DOMHTMLFontElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLFontElement */
// An interface definition for the [DOMHTMLFontElement] class.
type IDOMHTMLFontElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLFontElement */
	// properties:
	Color() objc.IObject /* cross-framework: NSString */
	SetColor(value objc.IObject /* cross-framework: NSString */)
	Face() objc.IObject /* cross-framework: NSString */
	SetFace(value objc.IObject /* cross-framework: NSString */)
	Size() objc.IObject /* cross-framework: NSString */
	SetSize(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLFontElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLFontElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLFontElementClass) Alloc() DOMHTMLFontElement {
	rv := objc.Send[DOMHTMLFontElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLFontElementClass) New() DOMHTMLFontElement {
	rv := objc.Send[DOMHTMLFontElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLFontElement) Init() DOMHTMLFontElement {
	rv := objc.Send[DOMHTMLFontElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLFontElement) Autorelease() DOMHTMLFontElement {
	rv := objc.Send[DOMHTMLFontElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLFontElement creates a new DOMHTMLFontElement instance.
func NewDOMHTMLFontElement() DOMHTMLFontElement {
	return getDOMHTMLFontElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLFontElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFontElement
type DOMHTMLFontElement struct {
	DOMHTMLElement
}

// DOMHTMLFontElementFrom constructs a [DOMHTMLFontElement] from an unsafe.Pointer.
func DOMHTMLFontElementFrom(ptr unsafe.Pointer) DOMHTMLFontElement {
	return DOMHTMLFontElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLFontElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLFontElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLFontElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLFontElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLFontElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFontElement/color
func (d_ DOMHTMLFontElement) Color() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("color"))
	return rv
}/* debug [instance_properties/getter]: color */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFontElement/color
func (d_ DOMHTMLFontElement) SetColor(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setColor:"), value)
}/* debug [instance_properties/setter]: color */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFontElement/face
func (d_ DOMHTMLFontElement) Face() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("face"))
	return rv
}/* debug [instance_properties/getter]: face */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFontElement/face
func (d_ DOMHTMLFontElement) SetFace(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFace:"), value)
}/* debug [instance_properties/setter]: face */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFontElement/size
func (d_ DOMHTMLFontElement) Size() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("size"))
	return rv
}/* debug [instance_properties/getter]: size */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFontElement/size
func (d_ DOMHTMLFontElement) SetSize(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSize:"), value)
}/* debug [instance_properties/setter]: size */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLFontElement */



