// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class UNTimeIntervalNotificationTrigger */

/* debug [class_header]: Header for UNTimeIntervalNotificationTrigger */
// The class instance for the [UNTimeIntervalNotificationTrigger] class.
var (
	UNTimeIntervalNotificationTriggerClass     _UNTimeIntervalNotificationTriggerClass
	UNTimeIntervalNotificationTriggerClassOnce sync.Once
)

func getUNTimeIntervalNotificationTriggerClass() _UNTimeIntervalNotificationTriggerClass {
	UNTimeIntervalNotificationTriggerClassOnce.Do(func() {
		UNTimeIntervalNotificationTriggerClass = _UNTimeIntervalNotificationTriggerClass{objc.GetClass("UNTimeIntervalNotificationTrigger")}
	})
	return UNTimeIntervalNotificationTriggerClass
}

type _UNTimeIntervalNotificationTriggerClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for UNTimeIntervalNotificationTrigger */
// An interface definition for the [UNTimeIntervalNotificationTrigger] class.
type IUNTimeIntervalNotificationTrigger interface {
	IUNNotificationTrigger

	/* debug [class_interface_properties]: Properties for UNTimeIntervalNotificationTrigger */
	// properties:
	TimeInterval() float64
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for UNTimeIntervalNotificationTrigger */
	// methods:
	NextTriggerDate() foundation.Date
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for UNTimeIntervalNotificationTrigger */
// Alloc allocates a new instance without initialization.
func (uc _UNTimeIntervalNotificationTriggerClass) Alloc() UNTimeIntervalNotificationTrigger {
	rv := objc.Send[UNTimeIntervalNotificationTrigger](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UNTimeIntervalNotificationTriggerClass) New() UNTimeIntervalNotificationTrigger {
	rv := objc.Send[UNTimeIntervalNotificationTrigger](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UNTimeIntervalNotificationTrigger) Init() UNTimeIntervalNotificationTrigger {
	rv := objc.Send[UNTimeIntervalNotificationTrigger](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UNTimeIntervalNotificationTrigger) Autorelease() UNTimeIntervalNotificationTrigger {
	rv := objc.Send[UNTimeIntervalNotificationTrigger](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNTimeIntervalNotificationTrigger creates a new UNTimeIntervalNotificationTrigger instance.
func NewUNTimeIntervalNotificationTrigger() UNTimeIntervalNotificationTrigger {
	return getUNTimeIntervalNotificationTriggerClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for UNTimeIntervalNotificationTrigger */
// A trigger condition that causes the system to deliver a notification after the amount of time you specify elapses.
//
// Create a object when you want to schedule the delivery of a local notification after the number of seconds you specify elapses. You use this type of trigger to implement timers. Listing 1 creates a trigger that delivers its notification one time after 30 minutes have elapsed. Listing 1. Creating a trigger that fires in 30 minutes

// A trigger condition that causes the system to deliver a notification after the amount of time you specify elapses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNTimeIntervalNotificationTrigger
type UNTimeIntervalNotificationTrigger struct {
	UNNotificationTrigger
}

// UNTimeIntervalNotificationTriggerFrom constructs a [UNTimeIntervalNotificationTrigger] from an unsafe.Pointer.
//
// A trigger condition that causes the system to deliver a notification after the amount of time you specify elapses.
func UNTimeIntervalNotificationTriggerFrom(ptr unsafe.Pointer) UNTimeIntervalNotificationTrigger {
	return UNTimeIntervalNotificationTrigger{
		UNNotificationTrigger: UNNotificationTriggerFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for UNTimeIntervalNotificationTrigger */

// Creates a time interval trigger using the time value parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNTimeIntervalNotificationTrigger/init(timeInterval:repeats:)
func NewUNTimeIntervalNotificationTriggerWithTimeIntervalRepeats(timeInterval float64, repeats bool) UNTimeIntervalNotificationTrigger {
	rv := objc.Send[UNTimeIntervalNotificationTrigger](objc.ID(getUNTimeIntervalNotificationTriggerClass().class), objc.Sel("triggerWithTimeInterval:repeats:"), timeInterval, repeats)
	return rv
} /* debug [class_init_methods/constructor]: NewUNTimeIntervalNotificationTriggerWithTimeIntervalRepeats */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for UNTimeIntervalNotificationTrigger */

// Creates a time interval trigger using the time value parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNTimeIntervalNotificationTrigger/init(timeInterval:repeats:)
func (uc _UNTimeIntervalNotificationTriggerClass) TriggerWithTimeIntervalRepeats(timeInterval float64, repeats bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("triggerWithTimeInterval:repeats:"), timeInterval, repeats)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=TriggerWithTimeIntervalRepeats) */

/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for UNTimeIntervalNotificationTrigger */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for UNTimeIntervalNotificationTrigger */

// The next date at which the trigger conditions are met.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNTimeIntervalNotificationTrigger/nextTriggerDate()
func (u_ UNTimeIntervalNotificationTrigger) NextTriggerDate() foundation.Date {
	rv := objc.Send[foundation.Date](u_.ID, objc.Sel("nextTriggerDate"))
	return rv
} /* debug [instance_methods/method]: NextTriggerDate */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for UNTimeIntervalNotificationTrigger */

// The time interval to create the trigger.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNTimeIntervalNotificationTrigger/timeInterval
func (u_ UNTimeIntervalNotificationTrigger) TimeInterval() float64 {
	rv := objc.Send[float64](u_.ID, objc.Sel("timeInterval"))
	return rv
} /* debug [instance_properties/getter]: timeInterval */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class UNTimeIntervalNotificationTrigger */
