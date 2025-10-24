// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMImplementation */


/* debug [class_header]: Header for DOMImplementation */
// The class instance for the [DOMImplementation] class.
var (
	DOMImplementationClass     _DOMImplementationClass
	DOMImplementationClassOnce sync.Once
)

func getDOMImplementationClass() _DOMImplementationClass {
	DOMImplementationClassOnce.Do(func() {
		DOMImplementationClass = _DOMImplementationClass{objc.GetClass("DOMImplementation")}
	})
	return DOMImplementationClass
}

type _DOMImplementationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMImplementation */
// An interface definition for the [DOMImplementation] class.
type IDOMImplementation interface {
	IDOMObject
	
/* debug [class_interface_properties]: Properties for DOMImplementation */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMImplementation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMImplementation */
// Alloc allocates a new instance without initialization.
func (dc _DOMImplementationClass) Alloc() DOMImplementation {
	rv := objc.Send[DOMImplementation](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMImplementationClass) New() DOMImplementation {
	rv := objc.Send[DOMImplementation](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMImplementation) Init() DOMImplementation {
	rv := objc.Send[DOMImplementation](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMImplementation) Autorelease() DOMImplementation {
	rv := objc.Send[DOMImplementation](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMImplementation creates a new DOMImplementation instance.
func NewDOMImplementation() DOMImplementation {
	return getDOMImplementationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMImplementation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMImplementation
type DOMImplementation struct {
	DOMObject
}

// DOMImplementationFrom constructs a [DOMImplementation] from an unsafe.Pointer.
func DOMImplementationFrom(ptr unsafe.Pointer) DOMImplementation {
	return DOMImplementation{
		DOMObject: DOMObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMImplementation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMImplementation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMImplementation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMImplementation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMImplementation */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMImplementation */



