// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/gameplaykit"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [UNNotificationResponse] class.
type IUNNotificationResponse interface {
	objectivec.IObject
	ActionIdentifier() string
	Notification() UNNotification
	TargetScene() gameplaykit.Scene
	UNNotificationDefaultActionIdentifier() string
	UNNotificationDismissActionIdentifier() string
}

// The user’s response to an actionable notification.
//
// When the user interacts with a delivered notification, the system delivers a object to your app so that you can process the response. Users can interact with delivered notifications in many ways. If the notification’s category had associated action buttons, they can select one of those buttons. Users can also dismiss the notification without selecting one of your actions and they can open your app. A response object tells you which option the user selected. You don’t create objects yourself. Instead, the shared user notification center object creates them and delivers them to the method of its delegate object. Use that method to extract any needed information from the response object and take appropriate action. For more information about responding to actions, see .
//
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

// Alloc allocates a new instance without initialization.
func (uc _UNNotificationResponseClass) Alloc() UNNotificationResponse {
	rv := objc.Send[UNNotificationResponse](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The identifier string of the action that the user selected.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationResponse/actionIdentifier
func (u_ UNNotificationResponse) ActionIdentifier() string {
	rv := objc.Send[string](u_.ID, objc.Sel("actionIdentifier"))
	return rv
}

// The notification to which the user responded.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationResponse/notification
func (u_ UNNotificationResponse) Notification() UNNotification {
	rv := objc.Send[UNNotification](u_.ID, objc.Sel("notification"))
	return rv
}

// The scene where the system reflects the user’s response to a notification.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationResponse/targetScene
func (u_ UNNotificationResponse) TargetScene() gameplaykit.Scene {
	rv := objc.Send[gameplaykit.Scene](u_.ID, objc.Sel("targetScene"))
	return rv
}

// An action that indicates the user opened the app from the notification interface.
//
// [Full Topic]: https://developer.apple.com/documentation/usernotifications/unnotificationdefaultactionidentifier
func (u_ UNNotificationResponse) UNNotificationDefaultActionIdentifier() string {
	rv := objc.Send[string](u_.ID, objc.Sel("UNNotificationDefaultActionIdentifier"))
	return rv
}

// The action that indicates the user explicitly dismissed the notification interface.
//
// [Full Topic]: https://developer.apple.com/documentation/usernotifications/unnotificationdismissactionidentifier
func (u_ UNNotificationResponse) UNNotificationDismissActionIdentifier() string {
	rv := objc.Send[string](u_.ID, objc.Sel("UNNotificationDismissActionIdentifier"))
	return rv
}



