// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CKSyncEngineDidFetchChangesEvent */


/* debug [class_header]: Header for CKSyncEngineDidFetchChangesEvent */
// The class instance for the [CKSyncEngineDidFetchChangesEvent] class.
var (
	CKSyncEngineDidFetchChangesEventClass     _CKSyncEngineDidFetchChangesEventClass
	CKSyncEngineDidFetchChangesEventClassOnce sync.Once
)

func getCKSyncEngineDidFetchChangesEventClass() _CKSyncEngineDidFetchChangesEventClass {
	CKSyncEngineDidFetchChangesEventClassOnce.Do(func() {
		CKSyncEngineDidFetchChangesEventClass = _CKSyncEngineDidFetchChangesEventClass{objc.GetClass("CKSyncEngineDidFetchChangesEvent")}
	})
	return CKSyncEngineDidFetchChangesEventClass
}

type _CKSyncEngineDidFetchChangesEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineDidFetchChangesEvent */
// An interface definition for the [CKSyncEngineDidFetchChangesEvent] class.
type ICKSyncEngineDidFetchChangesEvent interface {
	ICKSyncEngineEvent
	
/* debug [class_interface_properties]: Properties for CKSyncEngineDidFetchChangesEvent */
	// properties:
	Context() ICKSyncEngineFetchChangesContext
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineDidFetchChangesEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineDidFetchChangesEvent */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineDidFetchChangesEventClass) Alloc() CKSyncEngineDidFetchChangesEvent {
	rv := objc.Send[CKSyncEngineDidFetchChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEngineDidFetchChangesEventClass) New() CKSyncEngineDidFetchChangesEvent {
	rv := objc.Send[CKSyncEngineDidFetchChangesEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineDidFetchChangesEvent) Init() CKSyncEngineDidFetchChangesEvent {
	rv := objc.Send[CKSyncEngineDidFetchChangesEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineDidFetchChangesEvent) Autorelease() CKSyncEngineDidFetchChangesEvent {
	rv := objc.Send[CKSyncEngineDidFetchChangesEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineDidFetchChangesEvent creates a new CKSyncEngineDidFetchChangesEvent instance.
func NewCKSyncEngineDidFetchChangesEvent() CKSyncEngineDidFetchChangesEvent {
	return getCKSyncEngineDidFetchChangesEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineDidFetchChangesEvent */
// An object that represents a completed database fetch.


// An object that represents a completed database fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDidFetchChangesEvent
type CKSyncEngineDidFetchChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineDidFetchChangesEventFrom constructs a [CKSyncEngineDidFetchChangesEvent] from an unsafe.Pointer.
//
// An object that represents a completed database fetch.
func CKSyncEngineDidFetchChangesEventFrom(ptr unsafe.Pointer) CKSyncEngineDidFetchChangesEvent {
	return CKSyncEngineDidFetchChangesEvent{
		CKSyncEngineEvent: CKSyncEngineEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineDidFetchChangesEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineDidFetchChangesEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineDidFetchChangesEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineDidFetchChangesEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineDidFetchChangesEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineDidFetchChangesEvent/context
func (c_ CKSyncEngineDidFetchChangesEvent) Context() ICKSyncEngineFetchChangesContext {
	rv := objc.Send[CKSyncEngineFetchChangesContext](c_.ID, objc.Sel("context"))
	return rv
}/* debug [instance_properties/getter]: context */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineDidFetchChangesEvent */



