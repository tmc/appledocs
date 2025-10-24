// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

// Enum types and constants
// UNAlertStyle - Constants indicating the presentation styles for alerts.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAlertStyle
type UNAlertStyle uint

// UNAuthorizationOptions - Options that determine the authorized features of local and remote notifications.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAuthorizationOptions
type UNAuthorizationOptions uint

const (
	// UNAuthorizationOptionAnnouncement - The ability for Siri to automatically read out messages over AirPods.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAuthorizationOptions/announcement
	UNAuthorizationOptionAnnouncement UNAuthorizationOptions = 0
	// UNAuthorizationOptionBadge - The ability to update the app’s badge.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAuthorizationOptions/badge
	UNAuthorizationOptionBadge UNAuthorizationOptions = 0
	// UNAuthorizationOptionCarPlay - The ability to display notifications in a CarPlay environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAuthorizationOptions/carPlay
	UNAuthorizationOptionCarPlay UNAuthorizationOptions = 0
	// UNAuthorizationOptionCriticalAlert - The ability to play sounds for critical alerts.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAuthorizationOptions/criticalAlert
	UNAuthorizationOptionCriticalAlert UNAuthorizationOptions = 0
	// UNAuthorizationOptionProvisional - The ability to post noninterrupting notifications provisionally to the Notification Center.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAuthorizationOptions/provisional
	UNAuthorizationOptionProvisional UNAuthorizationOptions = 0
	// UNAuthorizationOptionSound - The ability to play sounds.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAuthorizationOptions/sound
	UNAuthorizationOptionSound UNAuthorizationOptions = 0
)

// UNAuthorizationStatus - Constants indicating whether the app is allowed to schedule notifications.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAuthorizationStatus
type UNAuthorizationStatus uint

const (
	// UNAuthorizationStatusProvisional - The application is provisionally authorized to post noninterruptive user notifications.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAuthorizationStatus/provisional
	UNAuthorizationStatusProvisional UNAuthorizationStatus = 0
)

// UNErrorCode - Constants that identify notification errors.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNError/Code
type UNErrorCode uint

// UNNotificationActionOptions - The behaviors you can apply to an action.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationActionOptions
type UNNotificationActionOptions uint

const (
	// UNNotificationActionOptionDestructive - The action performs a destructive task.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationActionOptions/destructive
	UNNotificationActionOptionDestructive UNNotificationActionOptions = 0
)

// UNNotificationCategoryOptions - Constants indicating how to handle notifications associated with this category.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategoryOptions
type UNNotificationCategoryOptions uint

// UNNotificationInterruptionLevel - Constants that indicate the importance and delivery timing of a notification.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationInterruptionLevel
type UNNotificationInterruptionLevel uint

// UNNotificationPresentationOptions - Constants indicating how to present a notification in a foreground app.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationPresentationOptions
type UNNotificationPresentationOptions uint

const (
	// UNNotificationPresentationOptionSound - Play the sound associated with the notification.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationPresentationOptions/sound
	UNNotificationPresentationOptionSound UNNotificationPresentationOptions = 0
)

// UNNotificationSetting - Constants that indicate the current status of a notification setting.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSetting
type UNNotificationSetting uint

const (
	// UNNotificationSettingEnabled - The setting is enabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSetting/enabled
	UNNotificationSettingEnabled UNNotificationSetting = 0
)

// UNShowPreviewsSetting - Constants indicating the style previewing a notification’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNShowPreviewsSetting
type UNShowPreviewsSetting uint


