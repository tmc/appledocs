// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CKSyncEngineWillFetchChangesEvent */


/* debug [class_header]: Header for CKSyncEngineWillFetchChangesEvent */
// The class instance for the [CKSyncEngineWillFetchChangesEvent] class.
var (
	CKSyncEngineWillFetchChangesEventClass     _CKSyncEngineWillFetchChangesEventClass
	CKSyncEngineWillFetchChangesEventClassOnce sync.Once
)

func getCKSyncEngineWillFetchChangesEventClass() _CKSyncEngineWillFetchChangesEventClass {
	CKSyncEngineWillFetchChangesEventClassOnce.Do(func() {
		CKSyncEngineWillFetchChangesEventClass = _CKSyncEngineWillFetchChangesEventClass{objc.GetClass("CKSyncEngineWillFetchChangesEvent")}
	})
	return CKSyncEngineWillFetchChangesEventClass
}

type _CKSyncEngineWillFetchChangesEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKSyncEngineWillFetchChangesEvent */
// An interface definition for the [CKSyncEngineWillFetchChangesEvent] class.
type ICKSyncEngineWillFetchChangesEvent interface {
	ICKSyncEngineEvent
	
/* debug [class_interface_properties]: Properties for CKSyncEngineWillFetchChangesEvent */
	// properties:
	Context() ICKSyncEngineFetchChangesContext
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKSyncEngineWillFetchChangesEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKSyncEngineWillFetchChangesEvent */
// Alloc allocates a new instance without initialization.
func (cc _CKSyncEngineWillFetchChangesEventClass) Alloc() CKSyncEngineWillFetchChangesEvent {
	rv := objc.Send[CKSyncEngineWillFetchChangesEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKSyncEngineWillFetchChangesEventClass) New() CKSyncEngineWillFetchChangesEvent {
	rv := objc.Send[CKSyncEngineWillFetchChangesEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKSyncEngineWillFetchChangesEvent) Init() CKSyncEngineWillFetchChangesEvent {
	rv := objc.Send[CKSyncEngineWillFetchChangesEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKSyncEngineWillFetchChangesEvent) Autorelease() CKSyncEngineWillFetchChangesEvent {
	rv := objc.Send[CKSyncEngineWillFetchChangesEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineWillFetchChangesEvent creates a new CKSyncEngineWillFetchChangesEvent instance.
func NewCKSyncEngineWillFetchChangesEvent() CKSyncEngineWillFetchChangesEvent {
	return getCKSyncEngineWillFetchChangesEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKSyncEngineWillFetchChangesEvent */
// An object that represents an imminent database fetch.


// An object that represents an imminent database fetch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineWillFetchChangesEvent
type CKSyncEngineWillFetchChangesEvent struct {
	CKSyncEngineEvent
}

// CKSyncEngineWillFetchChangesEventFrom constructs a [CKSyncEngineWillFetchChangesEvent] from an unsafe.Pointer.
//
// An object that represents an imminent database fetch.
func CKSyncEngineWillFetchChangesEventFrom(ptr unsafe.Pointer) CKSyncEngineWillFetchChangesEvent {
	return CKSyncEngineWillFetchChangesEvent{
		CKSyncEngineEvent: CKSyncEngineEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKSyncEngineWillFetchChangesEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKSyncEngineWillFetchChangesEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKSyncEngineWillFetchChangesEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKSyncEngineWillFetchChangesEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKSyncEngineWillFetchChangesEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKSyncEngineWillFetchChangesEvent/context
func (c_ CKSyncEngineWillFetchChangesEvent) Context() ICKSyncEngineFetchChangesContext {
	rv := objc.Send[CKSyncEngineFetchChangesContext](c_.ID, objc.Sel("context"))
	return rv
}/* debug [instance_properties/getter]: context */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKSyncEngineWillFetchChangesEvent */



