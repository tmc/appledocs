// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CTCallCenter */


/* debug [class_header]: Header for CTCallCenter */
// The class instance for the [CallCenter] class.
var (
	CallCenterClass     _CallCenterClass
	CallCenterClassOnce sync.Once
)

func getCallCenterClass() _CallCenterClass {
	CallCenterClassOnce.Do(func() {
		CallCenterClass = _CallCenterClass{objc.GetClass("CTCallCenter")}
	})
	return CallCenterClass
}

type _CallCenterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CallCenter */
// An interface definition for the [CallCenter] class.
type ICallCenter interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CallCenter */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CallCenter */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CallCenter */
// Alloc allocates a new instance without initialization.
func (cc _CallCenterClass) Alloc() CallCenter {
	rv := objc.Send[CallCenter](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CallCenterClass) New() CallCenter {
	rv := objc.Send[CallCenter](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CallCenter) Init() CallCenter {
	rv := objc.Send[CallCenter](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CallCenter) Autorelease() CallCenter {
	rv := objc.Send[CallCenter](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCallCenter creates a new CallCenter instance.
func NewCallCenter() CallCenter {
	return getCallCenterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CallCenter */
// An object that provides a list of current cellular calls, and provides the ability to respond to state changes for calls.


// An object that provides a list of current cellular calls, and provides the ability to respond to state changes for calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCallCenter
type CallCenter struct {
	objectivec.Object
}

// CallCenterFrom constructs a [CallCenter] from an unsafe.Pointer.
//
// An object that provides a list of current cellular calls, and provides the ability to respond to state changes for calls.
func CallCenterFrom(ptr unsafe.Pointer) CallCenter {
	return CallCenter{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CallCenter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CallCenter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CallCenter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CallCenter */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CallCenter */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CTCallCenter */


