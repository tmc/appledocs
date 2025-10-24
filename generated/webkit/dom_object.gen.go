// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMObject */


/* debug [class_header]: Header for DOMObject */
// The class instance for the [DOMObject] class.
var (
	DOMObjectClass     _DOMObjectClass
	DOMObjectClassOnce sync.Once
)

func getDOMObjectClass() _DOMObjectClass {
	DOMObjectClassOnce.Do(func() {
		DOMObjectClass = _DOMObjectClass{objc.GetClass("DOMObject")}
	})
	return DOMObjectClass
}

type _DOMObjectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMObject */
// An interface definition for the [DOMObject] class.
type IDOMObject interface {
	IWebScriptObject
	
/* debug [class_interface_properties]: Properties for DOMObject */
	// properties:
	Sheet() IDOMStyleSheet
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMObject */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMObject */
// Alloc allocates a new instance without initialization.
func (dc _DOMObjectClass) Alloc() DOMObject {
	rv := objc.Send[DOMObject](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMObjectClass) New() DOMObject {
	rv := objc.Send[DOMObject](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMObject) Init() DOMObject {
	rv := objc.Send[DOMObject](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMObject) Autorelease() DOMObject {
	rv := objc.Send[DOMObject](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMObject creates a new DOMObject instance.
func NewDOMObject() DOMObject {
	return getDOMObjectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMObject */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMObject
type DOMObject struct {
	WebScriptObject
}

// DOMObjectFrom constructs a [DOMObject] from an unsafe.Pointer.
func DOMObjectFrom(ptr unsafe.Pointer) DOMObject {
	return DOMObject{
		WebScriptObject: WebScriptObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMObject *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMObject */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMObject */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMObject */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMObject */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMObject/sheet
func (d_ DOMObject) Sheet() IDOMStyleSheet {
	rv := objc.Send[DOMStyleSheet](d_.ID, objc.Sel("sheet"))
	return rv
}/* debug [instance_properties/getter]: sheet */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMObject */



