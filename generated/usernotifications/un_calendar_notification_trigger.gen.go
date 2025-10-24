// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class UNCalendarNotificationTrigger */

/* debug [class_header]: Header for UNCalendarNotificationTrigger */
// The class instance for the [UNCalendarNotificationTrigger] class.
var (
	UNCalendarNotificationTriggerClass     _UNCalendarNotificationTriggerClass
	UNCalendarNotificationTriggerClassOnce sync.Once
)

func getUNCalendarNotificationTriggerClass() _UNCalendarNotificationTriggerClass {
	UNCalendarNotificationTriggerClassOnce.Do(func() {
		UNCalendarNotificationTriggerClass = _UNCalendarNotificationTriggerClass{objc.GetClass("UNCalendarNotificationTrigger")}
	})
	return UNCalendarNotificationTriggerClass
}

type _UNCalendarNotificationTriggerClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for UNCalendarNotificationTrigger */
// An interface definition for the [UNCalendarNotificationTrigger] class.
type IUNCalendarNotificationTrigger interface {
	IUNNotificationTrigger

	/* debug [class_interface_properties]: Properties for UNCalendarNotificationTrigger */
	// properties:
	DateComponents() foundation.DateComponents
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for UNCalendarNotificationTrigger */
	// methods:
	NextTriggerDate() foundation.Date
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for UNCalendarNotificationTrigger */
// Alloc allocates a new instance without initialization.
func (uc _UNCalendarNotificationTriggerClass) Alloc() UNCalendarNotificationTrigger {
	rv := objc.Send[UNCalendarNotificationTrigger](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UNCalendarNotificationTriggerClass) New() UNCalendarNotificationTrigger {
	rv := objc.Send[UNCalendarNotificationTrigger](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UNCalendarNotificationTrigger) Init() UNCalendarNotificationTrigger {
	rv := objc.Send[UNCalendarNotificationTrigger](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UNCalendarNotificationTrigger) Autorelease() UNCalendarNotificationTrigger {
	rv := objc.Send[UNCalendarNotificationTrigger](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNCalendarNotificationTrigger creates a new UNCalendarNotificationTrigger instance.
func NewUNCalendarNotificationTrigger() UNCalendarNotificationTrigger {
	return getUNCalendarNotificationTriggerClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for UNCalendarNotificationTrigger */
// A trigger condition that causes a notification the system delivers at a specific date and time.
//
// Create a object when you want to schedule the delivery of a local notification at the date and time you specify. You use an object to specify only the time values that you want the system to use to determine the matching date and time. Listing 1 creates a trigger that delivers its notification every morning at 8:30. The repeating behavior is achieved by specifying for the parameter when creating the trigger. Listing 1. Creating a trigger that repeats at a specific time

// A trigger condition that causes a notification the system delivers at a specific date and time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNCalendarNotificationTrigger
type UNCalendarNotificationTrigger struct {
	UNNotificationTrigger
}

// UNCalendarNotificationTriggerFrom constructs a [UNCalendarNotificationTrigger] from an unsafe.Pointer.
//
// A trigger condition that causes a notification the system delivers at a specific date and time.
func UNCalendarNotificationTriggerFrom(ptr unsafe.Pointer) UNCalendarNotificationTrigger {
	return UNCalendarNotificationTrigger{
		UNNotificationTrigger: UNNotificationTriggerFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for UNCalendarNotificationTrigger */

// Creates a calendar trigger using the date components parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNCalendarNotificationTrigger/init(dateMatching:repeats:)
func NewUNCalendarNotificationTriggerWithDateMatchingComponentsRepeats(dateComponents foundation.DateComponents, repeats bool) UNCalendarNotificationTrigger {
	rv := objc.Send[UNCalendarNotificationTrigger](objc.ID(getUNCalendarNotificationTriggerClass().class), objc.Sel("triggerWithDateMatchingComponents:repeats:"), dateComponents, repeats)
	return rv
} /* debug [class_init_methods/constructor]: NewUNCalendarNotificationTriggerWithDateMatchingComponentsRepeats */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for UNCalendarNotificationTrigger */

// Creates a calendar trigger using the date components parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNCalendarNotificationTrigger/init(dateMatching:repeats:)
func (uc _UNCalendarNotificationTriggerClass) TriggerWithDateMatchingComponentsRepeats(dateComponents foundation.DateComponents, repeats bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("triggerWithDateMatchingComponents:repeats:"), dateComponents, repeats)
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=TriggerWithDateMatchingComponentsRepeats) */

/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for UNCalendarNotificationTrigger */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for UNCalendarNotificationTrigger */

// The next date at which the trigger conditions are met.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNCalendarNotificationTrigger/nextTriggerDate()
func (u_ UNCalendarNotificationTrigger) NextTriggerDate() foundation.Date {
	rv := objc.Send[foundation.Date](u_.ID, objc.Sel("nextTriggerDate"))
	return rv
} /* debug [instance_methods/method]: NextTriggerDate */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for UNCalendarNotificationTrigger */

// The date components to construct this object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNCalendarNotificationTrigger/dateComponents
func (u_ UNCalendarNotificationTrigger) DateComponents() foundation.DateComponents {
	rv := objc.Send[foundation.DateComponents](u_.ID, objc.Sel("dateComponents"))
	return rv
} /* debug [instance_properties/getter]: dateComponents */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class UNCalendarNotificationTrigger */
