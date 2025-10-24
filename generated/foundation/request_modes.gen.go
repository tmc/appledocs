// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class requestModes */


/* debug [class_header]: Header for requestModes */
// The class instance for the [requestModes] class.
var (
	RequestModesClass     _requestModesClass
	RequestModesClassOnce sync.Once
)

func getrequestModesClass() _requestModesClass {
	RequestModesClassOnce.Do(func() {
		RequestModesClass = _requestModesClass{objc.GetClass("requestModes")}
	})
	return RequestModesClass
}

type _requestModesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for requestModes */
// An interface definition for the [requestModes] class.
type IrequestModes interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for requestModes */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for requestModes */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for requestModes */
// Alloc allocates a new instance without initialization.
func (rc _requestModesClass) Alloc() requestModes {
	rv := objc.Send[requestModes](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _requestModesClass) New() requestModes {
	rv := objc.Send[requestModes](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ requestModes) Init() requestModes {
	rv := objc.Send[requestModes](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ requestModes) Autorelease() requestModes {
	rv := objc.Send[requestModes](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewrequestModes creates a new requestModes instance.
func NewrequestModes() requestModes {
	return getrequestModesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for requestModes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/requestModes-c.ivar
type requestModes struct {
	objectivec.Object
}

// requestModesFrom constructs a [requestModes] from an unsafe.Pointer.
func requestModesFrom(ptr unsafe.Pointer) requestModes {
	return requestModes{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for requestModes *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for requestModes */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for requestModes */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for requestModes */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for requestModes */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class requestModes */



