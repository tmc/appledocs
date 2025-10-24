// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CTCall */


/* debug [class_header]: Header for CTCall */
// The class instance for the [Call] class.
var (
	CallClass     _CallClass
	CallClassOnce sync.Once
)

func getCallClass() _CallClass {
	CallClassOnce.Do(func() {
		CallClass = _CallClass{objc.GetClass("CTCall")}
	})
	return CallClass
}

type _CallClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Call */
// An interface definition for the [Call] class.
type ICall interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Call */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Call */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Call */
// Alloc allocates a new instance without initialization.
func (cc _CallClass) Alloc() Call {
	rv := objc.Send[Call](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CallClass) New() Call {
	rv := objc.Send[Call](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Call) Init() Call {
	rv := objc.Send[Call](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Call) Autorelease() Call {
	rv := objc.Send[Call](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCall creates a new Call instance.
func NewCall() Call {
	return getCallClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Call */
// An object used to identify a cellular call and determine its state.


// An object used to identify a cellular call and determine its state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCall
type Call struct {
	objectivec.Object
}

// CallFrom constructs a [Call] from an unsafe.Pointer.
//
// An object used to identify a cellular call and determine its state.
func CallFrom(ptr unsafe.Pointer) Call {
	return Call{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Call *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Call */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Call */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Call */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Call */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CTCall */


