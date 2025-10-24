// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CKSyncEngineStateUpdateEvent */


/* debug [class_header]: Header for CKSyncEngineStateUpdateEvent */
// The class instance for the [CKSyncEngineStateUpdateEvent] class.
var (
	CKSyncEngineStateUpdateEventClass     _CKSyncEngineStateUpdateEventClass
	CKSyncEngineStateUpdateEventClassOnce sync.Once
)

func getCKSyncEngineStateUpdateEventClass() _CKSyncEngineStateUpdateEventClass {
	CKSyncEngineStateUpdateEventClassOnce.Do(func() {
		CKSyncEngineStateUpdateEventClass = _CKSyncEngineStateUpdateEventClass{objc.GetClass("CKSyncEngineStateUpdateEvent")}
	})
	return CKSyncEngineStateUpdateEventClass
}

type _CKSyncEngineStateUpdateEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineStateUpdateEvent */
// An interface definition for the [CKSyncEngineStateUpdateEvent] class.
type ICKSyncEngineStateUpdateEvent interface {
	ICKSyncEngineEvent
	
/* debug [class_interface_properties]: Properties for CKSyncEngineStateUpdateEvent */
	// properties:
	StateSerialization() ICKSyncEngineStateSerialization
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineStateUpdateEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineStateUpdateEvent */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineStateUpdateEventClass) Alloc() CKSyncEngineStateUpdateEvent {
	rv := objc.Send[CKSyncEngineStateUpdateEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEngineStateUpdateEventClass) New() CKSyncEngineStateUpdateEvent {
	rv := objc.Send[CKSyncEngineStateUpdateEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineStateUpdateEvent) Init() CKSyncEngineStateUpdateEvent {
	rv := objc.Send[CKSyncEngineStateUpdateEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineStateUpdateEvent) Autorelease() CKSyncEngineStateUpdateEvent {
	rv := objc.Send[CKSyncEngineStateUpdateEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineStateUpdateEvent creates a new CKSyncEngineStateUpdateEvent instance.
func NewCKSyncEngineStateUpdateEvent() CKSyncEngineStateUpdateEvent {
	return getCKSyncEngineStateUpdateEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineStateUpdateEvent */
// An object that provides information about an update to the sync engine’s state.


// An object that provides information about an update to the sync engine’s state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineStateUpdateEvent
type CKSyncEngineStateUpdateEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineStateUpdateEventFrom constructs a [CKSyncEngineStateUpdateEvent] from an unsafe.Pointer.
//
// An object that provides information about an update to the sync engine’s state.
func CKSyncEngineStateUpdateEventFrom(ptr unsafe.Pointer) CKSyncEngineStateUpdateEvent {
	return CKSyncEngineStateUpdateEvent{
		CKSyncEngineEvent: CKSyncEngineEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineStateUpdateEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineStateUpdateEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineStateUpdateEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineStateUpdateEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineStateUpdateEvent */

// The current state of the sync engine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineStateUpdateEvent/stateSerialization
func (c_ CKSyncEngineStateUpdateEvent) StateSerialization() ICKSyncEngineStateSerialization {
	rv := objc.Send[CKSyncEngineStateSerialization](c_.ID, objc.Sel("stateSerialization"))
	return rv
}/* debug [instance_properties/getter]: stateSerialization */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineStateUpdateEvent */



