// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class UNNotificationSettings */


/* debug [class_header]: Header for UNNotificationSettings */
// The class instance for the [UNNotificationSettings] class.
var (
	UNNotificationSettingsClass     _UNNotificationSettingsClass
	UNNotificationSettingsClassOnce sync.Once
)

func getUNNotificationSettingsClass() _UNNotificationSettingsClass {
	UNNotificationSettingsClassOnce.Do(func() {
		UNNotificationSettingsClass = _UNNotificationSettingsClass{objc.GetClass("UNNotificationSettings")}
	})
	return UNNotificationSettingsClass
}

type _UNNotificationSettingsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UNNotificationSettings */
// An interface definition for the [UNNotificationSettings] class.
type IUNNotificationSettings interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for UNNotificationSettings */
	// properties:
	AlertSetting() UNNotificationSetting
	AlertStyle() UNAlertStyle
	AuthorizationStatus() UNAuthorizationStatus
	BadgeSetting() UNNotificationSetting
	CriticalAlertSetting() UNNotificationSetting
	DirectMessagesSetting() UNNotificationSetting
	LockScreenSetting() UNNotificationSetting
	NotificationCenterSetting() UNNotificationSetting
	ProvidesAppNotificationSettings() bool
	ScheduledDeliverySetting() UNNotificationSetting
	ShowPreviewsSetting() UNShowPreviewsSetting
	SoundSetting() UNNotificationSetting
	TimeSensitiveSetting() UNNotificationSetting
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UNNotificationSettings */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UNNotificationSettings */
// Alloc allocates a new instance without initialization.
func (uc _UNNotificationSettingsClass) Alloc() UNNotificationSettings {
	rv := objc.Send[UNNotificationSettings](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UNNotificationSettingsClass) New() UNNotificationSettings {
	rv := objc.Send[UNNotificationSettings](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UNNotificationSettings) Init() UNNotificationSettings {
	rv := objc.Send[UNNotificationSettings](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UNNotificationSettings) Autorelease() UNNotificationSettings {
	rv := objc.Send[UNNotificationSettings](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotificationSettings creates a new UNNotificationSettings instance.
func NewUNNotificationSettings() UNNotificationSettings {
	return getUNNotificationSettingsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UNNotificationSettings */
// The object for managing notification-related settings and the authorization status of your app.
//
// A object contains the current authorization status and notification-related settings for your app. Apps must receive authorization to schedule notifications and to interact with the user. Apps that run in CarPlay must similarly receive authorization to do so. Use this object to determine what notification-related actions your app can perform. You might then use that information to enable, disable, or adjust your app’s notification-related behaviors. Regardless of whether you take action, the system enforces your app’s settings by preventing denied interactions from occurring. You don’t create instances of this class directly. Instead, call the method of your app’s object to get the current settings. For more information about requesting authorization for user interactions, see .


// The object for managing notification-related settings and the authorization status of your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings
type UNNotificationSettings struct {
	objectivec.Object
}

// UNNotificationSettingsFrom constructs a [UNNotificationSettings] from an unsafe.Pointer.
//
// The object for managing notification-related settings and the authorization status of your app.
func UNNotificationSettingsFrom(ptr unsafe.Pointer) UNNotificationSettings {
	return UNNotificationSettings{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UNNotificationSettings *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UNNotificationSettings */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UNNotificationSettings */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UNNotificationSettings */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UNNotificationSettings */

// The authorization status for displaying alerts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/alertSetting
func (u_ UNNotificationSettings) AlertSetting() UNNotificationSetting {
	rv := objc.Send[UNNotificationSetting](u_.ID, objc.Sel("alertSetting"))
	return rv
}/* debug [instance_properties/getter]: alertSetting */


// The type of alert that the app may display when the device is unlocked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/alertStyle
func (u_ UNNotificationSettings) AlertStyle() UNAlertStyle {
	rv := objc.Send[UNAlertStyle](u_.ID, objc.Sel("alertStyle"))
	return rv
}/* debug [instance_properties/getter]: alertStyle */


// The app’s ability to schedule and receive local and remote notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/authorizationStatus
func (u_ UNNotificationSettings) AuthorizationStatus() UNAuthorizationStatus {
	rv := objc.Send[UNAuthorizationStatus](u_.ID, objc.Sel("authorizationStatus"))
	return rv
}/* debug [instance_properties/getter]: authorizationStatus */


// The setting that indicates whether badges appear on your app’s icon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/badgeSetting
func (u_ UNNotificationSettings) BadgeSetting() UNNotificationSetting {
	rv := objc.Send[UNNotificationSetting](u_.ID, objc.Sel("badgeSetting"))
	return rv
}/* debug [instance_properties/getter]: badgeSetting */


// The authorization status for playing sounds for critical alerts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/criticalAlertSetting
func (u_ UNNotificationSettings) CriticalAlertSetting() UNNotificationSetting {
	rv := objc.Send[UNNotificationSetting](u_.ID, objc.Sel("criticalAlertSetting"))
	return rv
}/* debug [instance_properties/getter]: criticalAlertSetting */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/directMessagesSetting
func (u_ UNNotificationSettings) DirectMessagesSetting() UNNotificationSetting {
	rv := objc.Send[UNNotificationSetting](u_.ID, objc.Sel("directMessagesSetting"))
	return rv
}/* debug [instance_properties/getter]: directMessagesSetting */


// The setting that indicates whether your app’s notifications appear on a device’s Lock screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/lockScreenSetting
func (u_ UNNotificationSettings) LockScreenSetting() UNNotificationSetting {
	rv := objc.Send[UNNotificationSetting](u_.ID, objc.Sel("lockScreenSetting"))
	return rv
}/* debug [instance_properties/getter]: lockScreenSetting */


// The setting that indicates whether your app’s notifications appear in Notification Center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/notificationCenterSetting
func (u_ UNNotificationSettings) NotificationCenterSetting() UNNotificationSetting {
	rv := objc.Send[UNNotificationSetting](u_.ID, objc.Sel("notificationCenterSetting"))
	return rv
}/* debug [instance_properties/getter]: notificationCenterSetting */


// A Boolean value indicating the system displays a button for in-app notification settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/providesAppNotificationSettings
func (u_ UNNotificationSettings) ProvidesAppNotificationSettings() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("providesAppNotificationSettings"))
	return rv
}/* debug [instance_properties/getter]: providesAppNotificationSettings */


// The setting that indicates the system schedules the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/scheduledDeliverySetting
func (u_ UNNotificationSettings) ScheduledDeliverySetting() UNNotificationSetting {
	rv := objc.Send[UNNotificationSetting](u_.ID, objc.Sel("scheduledDeliverySetting"))
	return rv
}/* debug [instance_properties/getter]: scheduledDeliverySetting */


// The setting that indicates whether the app shows a preview of the notification’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/showPreviewsSetting
func (u_ UNNotificationSettings) ShowPreviewsSetting() UNShowPreviewsSetting {
	rv := objc.Send[UNShowPreviewsSetting](u_.ID, objc.Sel("showPreviewsSetting"))
	return rv
}/* debug [instance_properties/getter]: showPreviewsSetting */


// The authorization status for playing sounds for incoming notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/soundSetting
func (u_ UNNotificationSettings) SoundSetting() UNNotificationSetting {
	rv := objc.Send[UNNotificationSetting](u_.ID, objc.Sel("soundSetting"))
	return rv
}/* debug [instance_properties/getter]: soundSetting */


// The setting that indicates the system treats the notification as time-sensitive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/timeSensitiveSetting
func (u_ UNNotificationSettings) TimeSensitiveSetting() UNNotificationSetting {
	rv := objc.Send[UNNotificationSetting](u_.ID, objc.Sel("timeSensitiveSetting"))
	return rv
}/* debug [instance_properties/getter]: timeSensitiveSetting */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class UNNotificationSettings */


