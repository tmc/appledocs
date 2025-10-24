// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CKSyncEngineDidSendChangesEvent */


/* debug [class_header]: Header for CKSyncEngineDidSendChangesEvent */
// The class instance for the [CKSyncEngineDidSendChangesEvent] class.
var (
	CKSyncEngineDidSendChangesEventClass     _CKSyncEngineDidSendChangesEventClass
	CKSyncEngineDidSendChangesEventClassOnce sync.Once
)

func getCKSyncEngineDidSendChangesEventClass() _CKSyncEngineDidSendChangesEventClass {
	CKSyncEngineDidSendChangesEventClassOnce.Do(func() {
		CKSyncEngineDidSendChangesEventClass = _CKSyncEngineDidSendChangesEventClass{objc.GetClass("CKSyncEngineDidSendChangesEvent")}
	})
	return CKSyncEngineDidSendChangesEventClass
}

type _CKSyncEngineDidSendChangesEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineDidSendChangesEvent */
// An interface definition for the [CKSyncEngineDidSendChangesEvent] class.
type ICKSyncEngineDidSendChangesEvent interface {
	ICKSyncEngineEvent
	
/* debug [class_interface_properties]: Properties for CKSyncEngineDidSendChangesEvent */
	// properties:
	Context() ICKSyncEngineSendChangesContext
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineDidSendChangesEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineDidSendChangesEvent */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineDidSendChangesEventClass) Alloc() CKSyncEngineDidSendChangesEvent {
	rv := objc.Send[CKSyncEngineDidSendChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEngineDidSendChangesEventClass) New() CKSyncEngineDidSendChangesEvent {
	rv := objc.Send[CKSyncEngineDidSendChangesEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineDidSendChangesEvent) Init() CKSyncEngineDidSendChangesEvent {
	rv := objc.Send[CKSyncEngineDidSendChangesEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineDidSendChangesEvent) Autorelease() CKSyncEngineDidSendChangesEvent {
	rv := objc.Send[CKSyncEngineDidSendChangesEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineDidSendChangesEvent creates a new CKSyncEngineDidSendChangesEvent instance.
func NewCKSyncEngineDidSendChangesEvent() CKSyncEngineDidSendChangesEvent {
	return getCKSyncEngineDidSendChangesEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineDidSendChangesEvent */
// An object that provides information about a finished send operation.


// An object that provides information about a finished send operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDidSendChangesEvent
type CKSyncEngineDidSendChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineDidSendChangesEventFrom constructs a [CKSyncEngineDidSendChangesEvent] from an unsafe.Pointer.
//
// An object that provides information about a finished send operation.
func CKSyncEngineDidSendChangesEventFrom(ptr unsafe.Pointer) CKSyncEngineDidSendChangesEvent {
	return CKSyncEngineDidSendChangesEvent{
		CKSyncEngineEvent: CKSyncEngineEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineDidSendChangesEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineDidSendChangesEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineDidSendChangesEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineDidSendChangesEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineDidSendChangesEvent */

// The context of the finished send request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDidSendChangesEvent/context
func (c_ CKSyncEngineDidSendChangesEvent) Context() ICKSyncEngineSendChangesContext {
	rv := objc.Send[CKSyncEngineSendChangesContext](c_.ID, objc.Sel("context"))
	return rv
}/* debug [instance_properties/getter]: context */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineDidSendChangesEvent */



