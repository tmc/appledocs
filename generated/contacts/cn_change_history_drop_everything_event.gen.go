// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CNChangeHistoryDropEverythingEvent */


/* debug [class_header]: Header for CNChangeHistoryDropEverythingEvent */
// The class instance for the [CNChangeHistoryDropEverythingEvent] class.
var (
	CNChangeHistoryDropEverythingEventClass     _CNChangeHistoryDropEverythingEventClass
	CNChangeHistoryDropEverythingEventClassOnce sync.Once
)

func getCNChangeHistoryDropEverythingEventClass() _CNChangeHistoryDropEverythingEventClass {
	CNChangeHistoryDropEverythingEventClassOnce.Do(func() {
		CNChangeHistoryDropEverythingEventClass = _CNChangeHistoryDropEverythingEventClass{objc.GetClass("CNChangeHistoryDropEverythingEvent")}
	})
	return CNChangeHistoryDropEverythingEventClass
}

type _CNChangeHistoryDropEverythingEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNChangeHistoryDropEverythingEvent */
// An interface definition for the [CNChangeHistoryDropEverythingEvent] class.
type ICNChangeHistoryDropEverythingEvent interface {
	ICNChangeHistoryEvent
	
/* debug [class_interface_properties]: Properties for CNChangeHistoryDropEverythingEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNChangeHistoryDropEverythingEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNChangeHistoryDropEverythingEvent */
// Alloc allocates a new instance without initialization.
func (cc _CNChangeHistoryDropEverythingEventClass) Alloc() CNChangeHistoryDropEverythingEvent {
	rv := objc.Send[CNChangeHistoryDropEverythingEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNChangeHistoryDropEverythingEventClass) New() CNChangeHistoryDropEverythingEvent {
	rv := objc.Send[CNChangeHistoryDropEverythingEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNChangeHistoryDropEverythingEvent) Init() CNChangeHistoryDropEverythingEvent {
	rv := objc.Send[CNChangeHistoryDropEverythingEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNChangeHistoryDropEverythingEvent) Autorelease() CNChangeHistoryDropEverythingEvent {
	rv := objc.Send[CNChangeHistoryDropEverythingEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNChangeHistoryDropEverythingEvent creates a new CNChangeHistoryDropEverythingEvent instance.
func NewCNChangeHistoryDropEverythingEvent() CNChangeHistoryDropEverythingEvent {
	return getCNChangeHistoryDropEverythingEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNChangeHistoryDropEverythingEvent */
// An object that indicates the delegate should drop all contacts and groups before handling change events.
//
// The system sends this event to your delegate when the system determines that enough has changed since the last time your app fetched the history changes that an incremental update is no longer possible. Following the drop-everything event, your app receives an add event for each contact and group currently in the database. This allows you to implement full syncs and incremental syncs using the same code.


// An object that indicates the delegate should drop all contacts and groups before handling change events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNChangeHistoryDropEverythingEvent
type CNChangeHistoryDropEverythingEvent struct {
	CNChangeHistoryEvent
}

// CNChangeHistoryDropEverythingEventFrom constructs a [CNChangeHistoryDropEverythingEvent] from an unsafe.Pointer.
//
// An object that indicates the delegate should drop all contacts and groups before handling change events.
func CNChangeHistoryDropEverythingEventFrom(ptr unsafe.Pointer) CNChangeHistoryDropEverythingEvent {
	return CNChangeHistoryDropEverythingEvent{
		CNChangeHistoryEvent: CNChangeHistoryEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNChangeHistoryDropEverythingEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNChangeHistoryDropEverythingEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNChangeHistoryDropEverythingEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNChangeHistoryDropEverythingEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNChangeHistoryDropEverythingEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNChangeHistoryDropEverythingEvent */



