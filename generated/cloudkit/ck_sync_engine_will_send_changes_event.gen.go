// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CKSyncEngineWillSendChangesEvent */


/* debug [class_header]: Header for CKSyncEngineWillSendChangesEvent */
// The class instance for the [CKSyncEngineWillSendChangesEvent] class.
var (
	CKSyncEngineWillSendChangesEventClass     _CKSyncEngineWillSendChangesEventClass
	CKSyncEngineWillSendChangesEventClassOnce sync.Once
)

func getCKSyncEngineWillSendChangesEventClass() _CKSyncEngineWillSendChangesEventClass {
	CKSyncEngineWillSendChangesEventClassOnce.Do(func() {
		CKSyncEngineWillSendChangesEventClass = _CKSyncEngineWillSendChangesEventClass{objc.GetClass("CKSyncEngineWillSendChangesEvent")}
	})
	return CKSyncEngineWillSendChangesEventClass
}

type _CKSyncEngineWillSendChangesEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineWillSendChangesEvent */
// An interface definition for the [CKSyncEngineWillSendChangesEvent] class.
type ICKSyncEngineWillSendChangesEvent interface {
	ICKSyncEngineEvent
	
/* debug [class_interface_properties]: Properties for CKSyncEngineWillSendChangesEvent */
	// properties:
	Context() ICKSyncEngineSendChangesContext
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineWillSendChangesEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineWillSendChangesEvent */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineWillSendChangesEventClass) Alloc() CKSyncEngineWillSendChangesEvent {
	rv := objc.Send[CKSyncEngineWillSendChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEngineWillSendChangesEventClass) New() CKSyncEngineWillSendChangesEvent {
	rv := objc.Send[CKSyncEngineWillSendChangesEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineWillSendChangesEvent) Init() CKSyncEngineWillSendChangesEvent {
	rv := objc.Send[CKSyncEngineWillSendChangesEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineWillSendChangesEvent) Autorelease() CKSyncEngineWillSendChangesEvent {
	rv := objc.Send[CKSyncEngineWillSendChangesEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineWillSendChangesEvent creates a new CKSyncEngineWillSendChangesEvent instance.
func NewCKSyncEngineWillSendChangesEvent() CKSyncEngineWillSendChangesEvent {
	return getCKSyncEngineWillSendChangesEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineWillSendChangesEvent */
// An object that provides information about an imminent send of local changes.


// An object that provides information about an imminent send of local changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineWillSendChangesEvent
type CKSyncEngineWillSendChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineWillSendChangesEventFrom constructs a [CKSyncEngineWillSendChangesEvent] from an unsafe.Pointer.
//
// An object that provides information about an imminent send of local changes.
func CKSyncEngineWillSendChangesEventFrom(ptr unsafe.Pointer) CKSyncEngineWillSendChangesEvent {
	return CKSyncEngineWillSendChangesEvent{
		CKSyncEngineEvent: CKSyncEngineEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineWillSendChangesEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineWillSendChangesEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineWillSendChangesEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineWillSendChangesEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineWillSendChangesEvent */

// The context of the imminent send request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineWillSendChangesEvent/context
func (c_ CKSyncEngineWillSendChangesEvent) Context() ICKSyncEngineSendChangesContext {
	rv := objc.Send[CKSyncEngineSendChangesContext](c_.ID, objc.Sel("context"))
	return rv
}/* debug [instance_properties/getter]: context */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineWillSendChangesEvent */



