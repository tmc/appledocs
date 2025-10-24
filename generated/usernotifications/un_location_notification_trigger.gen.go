// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class UNLocationNotificationTrigger */


/* debug [class_header]: Header for UNLocationNotificationTrigger */
// The class instance for the [UNLocationNotificationTrigger] class.
var (
	UNLocationNotificationTriggerClass     _UNLocationNotificationTriggerClass
	UNLocationNotificationTriggerClassOnce sync.Once
)

func getUNLocationNotificationTriggerClass() _UNLocationNotificationTriggerClass {
	UNLocationNotificationTriggerClassOnce.Do(func() {
		UNLocationNotificationTriggerClass = _UNLocationNotificationTriggerClass{objc.GetClass("UNLocationNotificationTrigger")}
	})
	return UNLocationNotificationTriggerClass
}

type _UNLocationNotificationTriggerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UNLocationNotificationTrigger */
// An interface definition for the [UNLocationNotificationTrigger] class.
type IUNLocationNotificationTrigger interface {
	IUNNotificationTrigger
	
/* debug [class_interface_properties]: Properties for UNLocationNotificationTrigger */
	// properties:
	NotifyOnEntry() bool
	SetNotifyOnEntry(value bool)
	NotifyOnExit() bool
	SetNotifyOnExit(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UNLocationNotificationTrigger */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UNLocationNotificationTrigger */
// Alloc allocates a new instance without initialization.
func (uc _UNLocationNotificationTriggerClass) Alloc() UNLocationNotificationTrigger {
	rv := objc.Send[UNLocationNotificationTrigger](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UNLocationNotificationTriggerClass) New() UNLocationNotificationTrigger {
	rv := objc.Send[UNLocationNotificationTrigger](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UNLocationNotificationTrigger) Init() UNLocationNotificationTrigger {
	rv := objc.Send[UNLocationNotificationTrigger](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UNLocationNotificationTrigger) Autorelease() UNLocationNotificationTrigger {
	rv := objc.Send[UNLocationNotificationTrigger](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNLocationNotificationTrigger creates a new UNLocationNotificationTrigger instance.
func NewUNLocationNotificationTrigger() UNLocationNotificationTrigger {
	return getUNLocationNotificationTriggerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UNLocationNotificationTrigger */
// A trigger condition that causes the system to deliver a notification when the user’s device enters or exits a geographic region you specify.
//
// Create a object when you want to schedule the delivery of a local notification when the device enters or leaves a specific geographic region. The system limits the number of location-based triggers that it schedules at the same time. When configuring the region, use the and properties to specify whether you want the system to deliver notifications on entry, on exit, or both. Listing 1 shows the creation of a trigger that fires only once when the user’s device enters a circular region with a 2-kilometer radius. Listing 1. Creating a location-based trigger The system doesn’t immediately trigger region-based notifications when the edge of the boundary is crossed. The system applies heuristics to ensure that the boundary crossing represents a deliberate event and isn’t the result of spurious location data. For more information about the heuristics, see .


// A trigger condition that causes the system to deliver a notification when the user’s device enters or exits a geographic region you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNLocationNotificationTrigger
type UNLocationNotificationTrigger struct {
	UNNotificationTrigger
}

// UNLocationNotificationTriggerFrom constructs a [UNLocationNotificationTrigger] from an unsafe.Pointer.
//
// A trigger condition that causes the system to deliver a notification when the user’s device enters or exits a geographic region you specify.
func UNLocationNotificationTriggerFrom(ptr unsafe.Pointer) UNLocationNotificationTrigger {
	return UNLocationNotificationTrigger{
		UNNotificationTrigger: UNNotificationTriggerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UNLocationNotificationTrigger */

// Creates a location trigger using the region parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNLocationNotificationTrigger/init(region:repeats:)
func NewUNLocationNotificationTriggerWithRegionRepeats(region corelocation.Region, repeats bool) UNLocationNotificationTrigger {
	rv := objc.Send[UNLocationNotificationTrigger](objc.ID(getUNLocationNotificationTriggerClass().class), objc.Sel("triggerWithRegion:repeats:"), region, repeats)
	return rv
}/* debug [class_init_methods/constructor]: NewUNLocationNotificationTriggerWithRegionRepeats */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UNLocationNotificationTrigger */

// Creates a location trigger using the region parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNLocationNotificationTrigger/init(region:repeats:)
func (uc _UNLocationNotificationTriggerClass) TriggerWithRegionRepeats(region corelocation.Region, repeats bool) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(uc.class), objc.Sel("triggerWithRegion:repeats:"), region, repeats)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TriggerWithRegionRepeats) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UNLocationNotificationTrigger */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UNLocationNotificationTrigger */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UNLocationNotificationTrigger */

// A Boolean indicating that notifications are generated upon entry into the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLRegion/notifyOnEntry
func (u_ UNLocationNotificationTrigger) NotifyOnEntry() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("notifyOnEntry"))
	return rv
}/* debug [instance_properties/getter]: notifyOnEntry */


// A Boolean indicating that notifications are generated upon entry into the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLRegion/notifyOnEntry
func (u_ UNLocationNotificationTrigger) SetNotifyOnEntry(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNotifyOnEntry:"), value)
}/* debug [instance_properties/setter]: notifyOnEntry */


// A Boolean indicating that notifications are generated upon exit from the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLRegion/notifyOnExit
func (u_ UNLocationNotificationTrigger) NotifyOnExit() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("notifyOnExit"))
	return rv
}/* debug [instance_properties/getter]: notifyOnExit */


// A Boolean indicating that notifications are generated upon exit from the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLRegion/notifyOnExit
func (u_ UNLocationNotificationTrigger) SetNotifyOnExit(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setNotifyOnExit:"), value)
}/* debug [instance_properties/setter]: notifyOnExit */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class UNLocationNotificationTrigger */


