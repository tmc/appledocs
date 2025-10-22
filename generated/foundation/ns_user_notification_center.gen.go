// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UserNotificationCenter] class.
var (
	UserNotificationCenterClass     _UserNotificationCenterClass
	UserNotificationCenterClassOnce sync.Once
)

func getUserNotificationCenterClass() _UserNotificationCenterClass {
	UserNotificationCenterClassOnce.Do(func() {
		UserNotificationCenterClass = _UserNotificationCenterClass{objc.GetClass("NSUserNotificationCenter")}
	})
	return UserNotificationCenterClass
}

type _UserNotificationCenterClass struct {
	class objc.Class
}

// An interface definition for the [UserNotificationCenter] class.
type IUserNotificationCenter interface {
	objectivec.IObject
	RemoveDeliveredNotification(notification IUserNotification)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	ActualDeliveryDate() Date
	SetActualDeliveryDate(value IDate)
	DeliveryDate() Date
	SetDeliveryDate(value IDate)
	IsPresented() bool
	SetIsPresented(value bool)
	DeliveredNotifications() NSUserNotification
	SetDeliveredNotifications(value IUserNotification)
	ScheduledNotifications() NSUserNotification
	SetScheduledNotifications(value IUserNotification)
}

// An object that delivers notifications from apps to the user.
//
// When a user notification’s delivery date has been reached, or it’s manually delivered, the notification center may display the notification to the user. The user notification center reserves the right to decide if a delivered user notification is presented to the user. For example, it may suppress the notification if the application is already frontmost (the delegate can override this action). The application can check the result of this decision by examining the property of a delivered user notification. instances the are tracking will be in one of two states: scheduled or delivered. A scheduled user notification has a . On that delivery date, the notification will move from being scheduled to being delivered. Note that the user notification may be displayed later than the delivery date depending on many factors. A delivered user notification has an . That’s the date when it moved from being scheduled to delivered, or when it was manually delivered using the method. The application and the user notification center are both ultimately subject to the user’s preferences. If the user decides to hide all alerts from your application, the property will still behave as above, but the user won’t see any animation or hear any sound. The provides more information about the delivered user notification and allows forcing the display of a user notification even if the application is frontmost.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotificationCenter
type UserNotificationCenter struct {
	objectivec.Object
}

// UserNotificationCenterFrom constructs a [UserNotificationCenter] from an unsafe.Pointer.
//
// An object that delivers notifications from apps to the user.
func UserNotificationCenterFrom(ptr unsafe.Pointer) UserNotificationCenter {
	return UserNotificationCenter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _UserNotificationCenterClass) Alloc() UserNotificationCenter {
	rv := objc.Send[UserNotificationCenter](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UserNotificationCenterClass) New() UserNotificationCenter {
	rv := objc.Send[UserNotificationCenter](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UserNotificationCenter) Init() UserNotificationCenter {
	rv := objc.Send[UserNotificationCenter](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UserNotificationCenter) Autorelease() UserNotificationCenter {
	rv := objc.Send[UserNotificationCenter](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUserNotificationCenter creates a new UserNotificationCenter instance.
func NewUserNotificationCenter() UserNotificationCenter {
	return getUserNotificationCenterClass().New()
}


// Remove a delivered user notification from the user notification center.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotificationCenter/removeDeliveredNotification(_:)
func (u_ UserNotificationCenter) RemoveDeliveredNotification(notification IUserNotification) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeDeliveredNotification:"), notification)
}

// Specifies the notification center delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotificationCenter/delegate
func (u_ UserNotificationCenter) Delegate() objc.ID {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// Specifies the notification center delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotificationCenter/delegate
func (u_ UserNotificationCenter) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDelegate:"), value)
}

// The date this notification was actually delivered.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/actualdeliverydate
func (u_ UserNotificationCenter) ActualDeliveryDate() Date {
	rv := objc.Send[Date](u_.ID, objc.Sel("actualDeliveryDate"))
	return rv
}


// SetActualDeliveryDate sets the value of the actualDeliveryDate property.
// The date this notification was actually delivered.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/actualdeliverydate
func (u_ UserNotificationCenter) SetActualDeliveryDate(value IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setActualDeliveryDate:"), value)
}

// Specifies when the notification should be delivered.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/deliverydate
func (u_ UserNotificationCenter) DeliveryDate() Date {
	rv := objc.Send[Date](u_.ID, objc.Sel("deliveryDate"))
	return rv
}


// SetDeliveryDate sets the value of the deliveryDate property.
// Specifies when the notification should be delivered.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/deliverydate
func (u_ UserNotificationCenter) SetDeliveryDate(value IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDeliveryDate:"), value)
}

// Specifies whether the user notification has been presented.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/ispresented
func (u_ UserNotificationCenter) IsPresented() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isPresented"))
	return rv
}


// SetIsPresented sets the value of the isPresented property.
// Specifies whether the user notification has been presented.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/ispresented
func (u_ UserNotificationCenter) SetIsPresented(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsPresented:"), value)
}

// An array of all user notifications delivered to the notification center.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotificationcenter/deliverednotifications
func (u_ UserNotificationCenter) DeliveredNotifications() NSUserNotification {
	rv := objc.Send[NSUserNotification](u_.ID, objc.Sel("deliveredNotifications"))
	return rv
}


// SetDeliveredNotifications sets the value of the deliveredNotifications property.
// An array of all user notifications delivered to the notification center.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotificationcenter/deliverednotifications
func (u_ UserNotificationCenter) SetDeliveredNotifications(value IUserNotification) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDeliveredNotifications:"), value)
}

// Specifies an array of scheduled user notifications that have not yet been delivered.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotificationcenter/schedulednotifications
func (u_ UserNotificationCenter) ScheduledNotifications() NSUserNotification {
	rv := objc.Send[NSUserNotification](u_.ID, objc.Sel("scheduledNotifications"))
	return rv
}


// SetScheduledNotifications sets the value of the scheduledNotifications property.
// Specifies an array of scheduled user notifications that have not yet been delivered.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotificationcenter/schedulednotifications
func (u_ UserNotificationCenter) SetScheduledNotifications(value IUserNotification) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setScheduledNotifications:"), value)
}



