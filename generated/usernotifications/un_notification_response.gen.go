// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/gameplaykit"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class UNNotificationResponse */


/* debug [class_header]: Header for UNNotificationResponse */
// The class instance for the [UNNotificationResponse] class.
var (
	UNNotificationResponseClass     _UNNotificationResponseClass
	UNNotificationResponseClassOnce sync.Once
)

func getUNNotificationResponseClass() _UNNotificationResponseClass {
	UNNotificationResponseClassOnce.Do(func() {
		UNNotificationResponseClass = _UNNotificationResponseClass{objc.GetClass("UNNotificationResponse")}
	})
	return UNNotificationResponseClass
}

type _UNNotificationResponseClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UNNotificationResponse */
// An interface definition for the [UNNotificationResponse] class.
type IUNNotificationResponse interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for UNNotificationResponse */
	// properties:
	ActionIdentifier() objc.IObject /* cross-framework: NSString */
	Notification() IUNNotification
	UNNotificationDefaultActionIdentifier() objc.IObject /* cross-framework: NSString */
	UNNotificationDismissActionIdentifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UNNotificationResponse */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UNNotificationResponse */
// Alloc allocates a new instance without initialization.
func (uc _UNNotificationResponseClass) Alloc() UNNotificationResponse {
	rv := objc.Send[UNNotificationResponse](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UNNotificationResponseClass) New() UNNotificationResponse {
	rv := objc.Send[UNNotificationResponse](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UNNotificationResponse) Init() UNNotificationResponse {
	rv := objc.Send[UNNotificationResponse](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UNNotificationResponse) Autorelease() UNNotificationResponse {
	rv := objc.Send[UNNotificationResponse](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotificationResponse creates a new UNNotificationResponse instance.
func NewUNNotificationResponse() UNNotificationResponse {
	return getUNNotificationResponseClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UNNotificationResponse */
// The user’s response to an actionable notification.
//
// When the user interacts with a delivered notification, the system delivers a object to your app so that you can process the response. Users can interact with delivered notifications in many ways. If the notification’s category had associated action buttons, they can select one of those buttons. Users can also dismiss the notification without selecting one of your actions and they can open your app. A response object tells you which option the user selected. You don’t create objects yourself. Instead, the shared user notification center object creates them and delivers them to the method of its delegate object. Use that method to extract any needed information from the response object and take appropriate action. For more information about responding to actions, see .


// The user’s response to an actionable notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationResponse
type UNNotificationResponse struct {
	objectivec.Object
}

// UNNotificationResponseFrom constructs a [UNNotificationResponse] from an unsafe.Pointer.
//
// The user’s response to an actionable notification.
func UNNotificationResponseFrom(ptr unsafe.Pointer) UNNotificationResponse {
	return UNNotificationResponse{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UNNotificationResponse *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UNNotificationResponse */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UNNotificationResponse */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UNNotificationResponse */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UNNotificationResponse */

// The identifier string of the action that the user selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationResponse/actionIdentifier
func (u_ UNNotificationResponse) ActionIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("actionIdentifier"))
	return rv
}/* debug [instance_properties/getter]: actionIdentifier */


// The notification to which the user responded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationResponse/notification
func (u_ UNNotificationResponse) Notification() IUNNotification {
	rv := objc.Send[UNNotification](u_.ID, objc.Sel("notification"))
	return rv
}/* debug [instance_properties/getter]: notification */


// An action that indicates the user opened the app from the notification interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/usernotifications/unnotificationdefaultactionidentifier
func (u_ UNNotificationResponse) UNNotificationDefaultActionIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("UNNotificationDefaultActionIdentifier"))
	return rv
}/* debug [instance_properties/getter]: UNNotificationDefaultActionIdentifier */


// The action that indicates the user explicitly dismissed the notification interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/usernotifications/unnotificationdismissactionidentifier
func (u_ UNNotificationResponse) UNNotificationDismissActionIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("UNNotificationDismissActionIdentifier"))
	return rv
}/* debug [instance_properties/getter]: UNNotificationDismissActionIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class UNNotificationResponse */


