// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [UNTimeIntervalNotificationTrigger] class.
type IUNTimeIntervalNotificationTrigger interface {
	IUNNotificationTrigger
	NextTriggerDate() foundation.Date
	TimeInterval() foundation.TimeInterval
}

// A trigger condition that causes the system to deliver a notification after the amount of time you specify elapses.
//
// Create a object when you want to schedule the delivery of a local notification after the number of seconds you specify elapses. You use this type of trigger to implement timers. Listing 1 creates a trigger that delivers its notification one time after 30 minutes have elapsed. Listing 1. Creating a trigger that fires in 30 minutes
//
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

// Alloc allocates a new instance without initialization.
func (uc _UNTimeIntervalNotificationTriggerClass) Alloc() UNTimeIntervalNotificationTrigger {
	rv := objc.Send[UNTimeIntervalNotificationTrigger](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a time interval trigger using the time value parameter.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNTimeIntervalNotificationTrigger/init(timeInterval:repeats:)
func NewUNTimeIntervalNotificationTriggerWithTimeIntervalRepeats(timeInterval foundation.ITimeInterval, repeats bool) UNTimeIntervalNotificationTrigger {
	rv := objc.Send[UNTimeIntervalNotificationTrigger](objc.ID(getUNTimeIntervalNotificationTriggerClass().class), objc.Sel("triggerWithTimeInterval:repeats:"), timeInterval, repeats)
	return rv
}


// Creates a time interval trigger using the time value parameter.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNTimeIntervalNotificationTrigger/init(timeInterval:repeats:)
func (uc _UNTimeIntervalNotificationTriggerClass) TriggerWithTimeIntervalRepeats(timeInterval foundation.ITimeInterval, repeats bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("triggerWithTimeInterval:repeats:"), timeInterval, repeats)
	return rv
}

// The next date at which the trigger conditions are met.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNTimeIntervalNotificationTrigger/nextTriggerDate()
func (u_ UNTimeIntervalNotificationTrigger) NextTriggerDate() foundation.Date {
	rv := objc.Send[foundation.Date](u_.ID, objc.Sel("nextTriggerDate"))
	return rv
}

// The time interval to create the trigger.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNTimeIntervalNotificationTrigger/timeInterval
func (u_ UNTimeIntervalNotificationTrigger) TimeInterval() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](u_.ID, objc.Sel("timeInterval"))
	return rv
}



