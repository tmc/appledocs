// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CXCallUpdate */


/* debug [class_header]: Header for CXCallUpdate */
// The class instance for the [CXCallUpdate] class.
var (
	CXCallUpdateClass     _CXCallUpdateClass
	CXCallUpdateClassOnce sync.Once
)

func getCXCallUpdateClass() _CXCallUpdateClass {
	CXCallUpdateClassOnce.Do(func() {
		CXCallUpdateClass = _CXCallUpdateClass{objc.GetClass("CXCallUpdate")}
	})
	return CXCallUpdateClass
}

type _CXCallUpdateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CXCallUpdate */
// An interface definition for the [CXCallUpdate] class.
type ICXCallUpdate interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CXCallUpdate */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CXCallUpdate */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CXCallUpdate */
// Alloc allocates a new instance without initialization.
func (cc _CXCallUpdateClass) Alloc() CXCallUpdate {
	rv := objc.Send[CXCallUpdate](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CXCallUpdateClass) New() CXCallUpdate {
	rv := objc.Send[CXCallUpdate](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXCallUpdate) Init() CXCallUpdate {
	rv := objc.Send[CXCallUpdate](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXCallUpdate) Autorelease() CXCallUpdate {
	rv := objc.Send[CXCallUpdate](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXCallUpdate creates a new CXCallUpdate instance.
func NewCXCallUpdate() CXCallUpdate {
	return getCXCallUpdateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CXCallUpdate */
// An encapsulation of new and changed information about a call.
//
// objects are used by the system to communicate changes to calls over time. Not every property on a object must be set each time, as each object includes only new and changed information. For example, when a call is started, only some properties may be known and included in the first object sent to the system, such as . Later in the same call, other properties may change; for example, a call may be upgraded from audio only to audio and video, which would be reflected by a new object with its property set to . When an incoming call is received, you construct a object specifying a and pass that to the method to notify the telephony provider. When an active call is updated, you construct a object specifying any updated information and pass that to the method. For example, if a user changes their contact information during a call, you could notify the telephony provider of this change using a new object with the new value set to its property.


// An encapsulation of new and changed information about a call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallUpdate
type CXCallUpdate struct {
	objectivec.Object
}

// CXCallUpdateFrom constructs a [CXCallUpdate] from an unsafe.Pointer.
//
// An encapsulation of new and changed information about a call.
func CXCallUpdateFrom(ptr unsafe.Pointer) CXCallUpdate {
	return CXCallUpdate{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CXCallUpdate *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CXCallUpdate */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CXCallUpdate */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CXCallUpdate */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CXCallUpdate */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CXCallUpdate */


