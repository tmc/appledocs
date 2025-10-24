// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMHTMLFrameSetElement */


/* debug [class_header]: Header for DOMHTMLFrameSetElement */
// The class instance for the [DOMHTMLFrameSetElement] class.
var (
	DOMHTMLFrameSetElementClass     _DOMHTMLFrameSetElementClass
	DOMHTMLFrameSetElementClassOnce sync.Once
)

func getDOMHTMLFrameSetElementClass() _DOMHTMLFrameSetElementClass {
	DOMHTMLFrameSetElementClassOnce.Do(func() {
		DOMHTMLFrameSetElementClass = _DOMHTMLFrameSetElementClass{objc.GetClass("DOMHTMLFrameSetElement")}
	})
	return DOMHTMLFrameSetElementClass
}

type _DOMHTMLFrameSetElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLFrameSetElement */
// An interface definition for the [DOMHTMLFrameSetElement] class.
type IDOMHTMLFrameSetElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLFrameSetElement */
	// properties:
	Cols() objc.IObject /* cross-framework: NSString */
	SetCols(value objc.IObject /* cross-framework: NSString */)
	Rows() objc.IObject /* cross-framework: NSString */
	SetRows(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLFrameSetElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLFrameSetElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLFrameSetElementClass) Alloc() DOMHTMLFrameSetElement {
	rv := objc.Send[DOMHTMLFrameSetElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLFrameSetElementClass) New() DOMHTMLFrameSetElement {
	rv := objc.Send[DOMHTMLFrameSetElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLFrameSetElement) Init() DOMHTMLFrameSetElement {
	rv := objc.Send[DOMHTMLFrameSetElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLFrameSetElement) Autorelease() DOMHTMLFrameSetElement {
	rv := objc.Send[DOMHTMLFrameSetElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLFrameSetElement creates a new DOMHTMLFrameSetElement instance.
func NewDOMHTMLFrameSetElement() DOMHTMLFrameSetElement {
	return getDOMHTMLFrameSetElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLFrameSetElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameSetElement
type DOMHTMLFrameSetElement struct {
	DOMHTMLElement
}

// DOMHTMLFrameSetElementFrom constructs a [DOMHTMLFrameSetElement] from an unsafe.Pointer.
func DOMHTMLFrameSetElementFrom(ptr unsafe.Pointer) DOMHTMLFrameSetElement {
	return DOMHTMLFrameSetElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLFrameSetElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLFrameSetElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLFrameSetElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLFrameSetElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLFrameSetElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameSetElement/cols
func (d_ DOMHTMLFrameSetElement) Cols() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("cols"))
	return rv
}/* debug [instance_properties/getter]: cols */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameSetElement/cols
func (d_ DOMHTMLFrameSetElement) SetCols(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCols:"), value)
}/* debug [instance_properties/setter]: cols */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameSetElement/rows
func (d_ DOMHTMLFrameSetElement) Rows() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("rows"))
	return rv
}/* debug [instance_properties/getter]: rows */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLFrameSetElement/rows
func (d_ DOMHTMLFrameSetElement) SetRows(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setRows:"), value)
}/* debug [instance_properties/setter]: rows */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLFrameSetElement */



