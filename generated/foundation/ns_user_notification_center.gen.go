// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSUserNotificationCenter */


/* debug [class_header]: Header for NSUserNotificationCenter */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UserNotificationCenter */
// An interface definition for the [UserNotificationCenter] class.
type IUserNotificationCenter interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for UserNotificationCenter */
	// properties:
	DeliveredNotifications() []UserNotification
	ActualDeliveryDate() IDate
	SetActualDeliveryDate(value IDate)
	DeliveryDate() IDate
	SetDeliveryDate(value IDate)
	IsPresented() bool
	SetIsPresented(value bool)
	ScheduledNotifications() IUserNotification
	SetScheduledNotifications(value IUserNotification)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UserNotificationCenter */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UserNotificationCenter */
// Alloc allocates a new instance without initialization.
func (uc _UserNotificationCenterClass) Alloc() UserNotificationCenter {
	rv := objc.Send[UserNotificationCenter](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UserNotificationCenter */
// An object that delivers notifications from apps to the user.
//
// When a user notification’s delivery date has been reached, or it’s manually delivered, the notification center may display the notification to the user. The user notification center reserves the right to decide if a delivered user notification is presented to the user. For example, it may suppress the notification if the application is already frontmost (the delegate can override this action). The application can check the result of this decision by examining the property of a delivered user notification. instances the are tracking will be in one of two states: scheduled or delivered. A scheduled user notification has a . On that delivery date, the notification will move from being scheduled to being delivered. Note that the user notification may be displayed later than the delivery date depending on many factors. A delivered user notification has an . That’s the date when it moved from being scheduled to delivered, or when it was manually delivered using the method. The application and the user notification center are both ultimately subject to the user’s preferences. If the user decides to hide all alerts from your application, the property will still behave as above, but the user won’t see any animation or hear any sound. The provides more information about the delivered user notification and allows forcing the display of a user notification even if the application is frontmost.


// An object that delivers notifications from apps to the user.
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UserNotificationCenter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UserNotificationCenter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UserNotificationCenter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UserNotificationCenter */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UserNotificationCenter */

// An array of all user notifications delivered to the notification center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserNotificationCenter/deliveredNotifications
func (u_ UserNotificationCenter) DeliveredNotifications() []UserNotification {
	rv := objc.Send[[]UserNotification](u_.ID, objc.Sel("deliveredNotifications"))
	return rv
}/* debug [instance_properties/getter]: deliveredNotifications */


// The date this notification was actually delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/actualdeliverydate
func (u_ UserNotificationCenter) ActualDeliveryDate() IDate {
	rv := objc.Send[Date](u_.ID, objc.Sel("actualDeliveryDate"))
	return rv
}/* debug [instance_properties/getter]: actualDeliveryDate */


// The date this notification was actually delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/actualdeliverydate
func (u_ UserNotificationCenter) SetActualDeliveryDate(value IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setActualDeliveryDate:"), value)
}/* debug [instance_properties/setter]: actualDeliveryDate */


// Specifies when the notification should be delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/deliverydate
func (u_ UserNotificationCenter) DeliveryDate() IDate {
	rv := objc.Send[Date](u_.ID, objc.Sel("deliveryDate"))
	return rv
}/* debug [instance_properties/getter]: deliveryDate */


// Specifies when the notification should be delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/deliverydate
func (u_ UserNotificationCenter) SetDeliveryDate(value IDate) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setDeliveryDate:"), value)
}/* debug [instance_properties/setter]: deliveryDate */


// Specifies whether the user notification has been presented.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/ispresented
func (u_ UserNotificationCenter) IsPresented() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isPresented"))
	return rv
}/* debug [instance_properties/getter]: isPresented */


// Specifies whether the user notification has been presented.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotification/ispresented
func (u_ UserNotificationCenter) SetIsPresented(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsPresented:"), value)
}/* debug [instance_properties/setter]: isPresented */


// Specifies an array of scheduled user notifications that have not yet been delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotificationcenter/schedulednotifications
func (u_ UserNotificationCenter) ScheduledNotifications() IUserNotification {
	rv := objc.Send[UserNotification](u_.ID, objc.Sel("scheduledNotifications"))
	return rv
}/* debug [instance_properties/getter]: scheduledNotifications */


// Specifies an array of scheduled user notifications that have not yet been delivered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsusernotificationcenter/schedulednotifications
func (u_ UserNotificationCenter) SetScheduledNotifications(value IUserNotification) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setScheduledNotifications:"), value)
}/* debug [instance_properties/setter]: scheduledNotifications */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSUserNotificationCenter */



