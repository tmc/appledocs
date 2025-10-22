// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [UNCalendarNotificationTrigger] class.
type IUNCalendarNotificationTrigger interface {
	IUNNotificationTrigger
	NextTriggerDate() foundation.Date
	DateComponents() foundation.DateComponents
}

// A trigger condition that causes a notification the system delivers at a specific date and time.
//
// Create a object when you want to schedule the delivery of a local notification at the date and time you specify. You use an object to specify only the time values that you want the system to use to determine the matching date and time. Listing 1 creates a trigger that delivers its notification every morning at 8:30. The repeating behavior is achieved by specifying for the parameter when creating the trigger. Listing 1. Creating a trigger that repeats at a specific time
//
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

// Alloc allocates a new instance without initialization.
func (uc _UNCalendarNotificationTriggerClass) Alloc() UNCalendarNotificationTrigger {
	rv := objc.Send[UNCalendarNotificationTrigger](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a calendar trigger using the date components parameter.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNCalendarNotificationTrigger/init(dateMatching:repeats:)
func NewUNCalendarNotificationTriggerWithDateMatchingComponentsRepeats(dateComponents foundation.IDateComponents, repeats bool) UNCalendarNotificationTrigger {
	rv := objc.Send[UNCalendarNotificationTrigger](objc.ID(getUNCalendarNotificationTriggerClass().class), objc.Sel("triggerWithDateMatchingComponents:repeats:"), dateComponents, repeats)
	return rv
}


// Creates a calendar trigger using the date components parameter.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNCalendarNotificationTrigger/init(dateMatching:repeats:)
func (uc _UNCalendarNotificationTriggerClass) TriggerWithDateMatchingComponentsRepeats(dateComponents foundation.IDateComponents, repeats bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("triggerWithDateMatchingComponents:repeats:"), dateComponents, repeats)
	return rv
}

// The next date at which the trigger conditions are met.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNCalendarNotificationTrigger/nextTriggerDate()
func (u_ UNCalendarNotificationTrigger) NextTriggerDate() foundation.Date {
	rv := objc.Send[foundation.Date](u_.ID, objc.Sel("nextTriggerDate"))
	return rv
}

// The date components to construct this object.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNCalendarNotificationTrigger/dateComponents
func (u_ UNCalendarNotificationTrigger) DateComponents() foundation.DateComponents {
	rv := objc.Send[foundation.DateComponents](u_.ID, objc.Sel("dateComponents"))
	return rv
}


