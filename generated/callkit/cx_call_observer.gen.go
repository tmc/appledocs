// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CXCallObserver */


/* debug [class_header]: Header for CXCallObserver */
// The class instance for the [CXCallObserver] class.
var (
	CXCallObserverClass     _CXCallObserverClass
	CXCallObserverClassOnce sync.Once
)

func getCXCallObserverClass() _CXCallObserverClass {
	CXCallObserverClassOnce.Do(func() {
		CXCallObserverClass = _CXCallObserverClass{objc.GetClass("CXCallObserver")}
	})
	return CXCallObserverClass
}

type _CXCallObserverClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CXCallObserver */
// An interface definition for the [CXCallObserver] class.
type ICXCallObserver interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CXCallObserver */
	// properties:
	CallObserver() ICXCallObserver
	SetCallObserver(value ICXCallObserver)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CXCallObserver */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CXCallObserver */
// Alloc allocates a new instance without initialization.
func (cc _CXCallObserverClass) Alloc() CXCallObserver {
	rv := objc.Send[CXCallObserver](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CXCallObserverClass) New() CXCallObserver {
	rv := objc.Send[CXCallObserver](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXCallObserver) Init() CXCallObserver {
	rv := objc.Send[CXCallObserver](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXCallObserver) Autorelease() CXCallObserver {
	rv := objc.Send[CXCallObserver](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXCallObserver creates a new CXCallObserver instance.
func NewCXCallObserver() CXCallObserver {
	return getCXCallObserverClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CXCallObserver */
// A programmatic interface for an object that manages a list of active calls and observes call changes.
//
// You can retrieve a list of active calls on an object using the property. You can also provide an object conforming to the protocol as the call observer delegate using the method to respond to any active call changes. VoIP apps typically interact with the object returned by the property of a instance. However, any app can create a new object to be notified of any calls activity on the system.


// A programmatic interface for an object that manages a list of active calls and observes call changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallObserver
type CXCallObserver struct {
	objectivec.Object
}

// CXCallObserverFrom constructs a [CXCallObserver] from an unsafe.Pointer.
//
// A programmatic interface for an object that manages a list of active calls and observes call changes.
func CXCallObserverFrom(ptr unsafe.Pointer) CXCallObserver {
	return CXCallObserver{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CXCallObserver *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CXCallObserver */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CXCallObserver */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CXCallObserver */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CXCallObserver */

// Returns an observer for active calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxcallcontroller/callobserver
func (c_ CXCallObserver) CallObserver() ICXCallObserver {
	rv := objc.Send[CXCallObserver](c_.ID, objc.Sel("callObserver"))
	return rv
}/* debug [instance_properties/getter]: callObserver */


// Returns an observer for active calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxcallcontroller/callobserver
func (c_ CXCallObserver) SetCallObserver(value ICXCallObserver) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCallObserver:"), value)
}/* debug [instance_properties/setter]: callObserver */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CXCallObserver */


