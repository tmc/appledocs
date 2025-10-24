// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMHTMLBRElement */


/* debug [class_header]: Header for DOMHTMLBRElement */
// The class instance for the [DOMHTMLBRElement] class.
var (
	DOMHTMLBRElementClass     _DOMHTMLBRElementClass
	DOMHTMLBRElementClassOnce sync.Once
)

func getDOMHTMLBRElementClass() _DOMHTMLBRElementClass {
	DOMHTMLBRElementClassOnce.Do(func() {
		DOMHTMLBRElementClass = _DOMHTMLBRElementClass{objc.GetClass("DOMHTMLBRElement")}
	})
	return DOMHTMLBRElementClass
}

type _DOMHTMLBRElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLBRElement */
// An interface definition for the [DOMHTMLBRElement] class.
type IDOMHTMLBRElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLBRElement */
	// properties:
	Clear() objc.IObject /* cross-framework: NSString */
	SetClear(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLBRElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLBRElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLBRElementClass) Alloc() DOMHTMLBRElement {
	rv := objc.Send[DOMHTMLBRElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLBRElementClass) New() DOMHTMLBRElement {
	rv := objc.Send[DOMHTMLBRElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLBRElement) Init() DOMHTMLBRElement {
	rv := objc.Send[DOMHTMLBRElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLBRElement) Autorelease() DOMHTMLBRElement {
	rv := objc.Send[DOMHTMLBRElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLBRElement creates a new DOMHTMLBRElement instance.
func NewDOMHTMLBRElement() DOMHTMLBRElement {
	return getDOMHTMLBRElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLBRElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBRElement
type DOMHTMLBRElement struct {
	DOMHTMLElement
}

// DOMHTMLBRElementFrom constructs a [DOMHTMLBRElement] from an unsafe.Pointer.
func DOMHTMLBRElementFrom(ptr unsafe.Pointer) DOMHTMLBRElement {
	return DOMHTMLBRElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLBRElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLBRElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLBRElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLBRElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLBRElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBRElement/clear
func (d_ DOMHTMLBRElement) Clear() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("clear"))
	return rv
}/* debug [instance_properties/getter]: clear */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLBRElement/clear
func (d_ DOMHTMLBRElement) SetClear(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setClear:"), value)
}/* debug [instance_properties/setter]: clear */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLBRElement */



