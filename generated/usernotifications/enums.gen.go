// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

/* debug [enums.gen.go]: Generating 10 enums for UserNotifications */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum UNAlertStyle (3 cases) */
// UNAlertStyle - Constants indicating the presentation styles for alerts.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAlertStyle
type UNAlertStyle uint

const (
	// UNAlertStyleAlert - Modal alerts.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAlertStyle/alert
	UNAlertStyleAlert UNAlertStyle = 0
	// UNAlertStyleBanner - Banner alerts.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAlertStyle/banner
	UNAlertStyleBanner UNAlertStyle = 0
	// UNAlertStyleNone - No alert.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAlertStyle/none
	UNAlertStyleNone UNAlertStyle = 0
)

/* debug [enums.gen.go]: Processing enum UNAuthorizationOptions (9 cases) */
// UNAuthorizationOptions - Options that determine the authorized features of local and remote notifications.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAuthorizationOptions
type UNAuthorizationOptions uint

const (
	// UNAuthorizationOptionAlert - The ability to display alerts.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAuthorizationOptions/alert
	UNAuthorizationOptionAlert UNAuthorizationOptions = 0
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
	// UNAuthorizationOptionProvidesAppNotificationSettings - An option indicating the system should display a button for in-app notification settings.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAuthorizationOptions/providesAppNotificationSettings
	UNAuthorizationOptionProvidesAppNotificationSettings UNAuthorizationOptions = 0
	// UNAuthorizationOptionProvisional - The ability to post noninterrupting notifications provisionally to the Notification Center.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAuthorizationOptions/provisional
	UNAuthorizationOptionProvisional UNAuthorizationOptions = 0
	// UNAuthorizationOptionSound - The ability to play sounds.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAuthorizationOptions/sound
	UNAuthorizationOptionSound UNAuthorizationOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAuthorizationOptions/timeSensitive
	UNAuthorizationOptionTimeSensitive UNAuthorizationOptions = 0
)

/* debug [enums.gen.go]: Processing enum UNAuthorizationStatus (5 cases) */
// UNAuthorizationStatus - Constants indicating whether the app is allowed to schedule notifications.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAuthorizationStatus
type UNAuthorizationStatus uint

const (
	// UNAuthorizationStatusAuthorized - The app is authorized to schedule or receive notifications.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAuthorizationStatus/authorized
	UNAuthorizationStatusAuthorized UNAuthorizationStatus = 0
	// UNAuthorizationStatusDenied - The app isn’t authorized to schedule or receive notifications.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAuthorizationStatus/denied
	UNAuthorizationStatusDenied UNAuthorizationStatus = 0
	// UNAuthorizationStatusEphemeral - The app is authorized to schedule or receive notifications for a limited amount of time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAuthorizationStatus/ephemeral
	UNAuthorizationStatusEphemeral UNAuthorizationStatus = 0
	// UNAuthorizationStatusNotDetermined - The user hasn’t yet made a choice about whether the app is allowed to schedule notifications.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAuthorizationStatus/notDetermined
	UNAuthorizationStatusNotDetermined UNAuthorizationStatus = 0
	// UNAuthorizationStatusProvisional - The application is provisionally authorized to post noninterruptive user notifications.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNAuthorizationStatus/provisional
	UNAuthorizationStatusProvisional UNAuthorizationStatus = 0
)

/* debug [enums.gen.go]: Processing enum UNErrorCode (12 cases) */
// UNErrorCode - Constants that identify notification errors.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNError/Code
type UNErrorCode uint

const (
	// UNErrorCodeAttachmentCorrupt - The file for an attachment is corrupt.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNError/Code/attachmentCorrupt
	UNErrorCodeAttachmentCorrupt UNErrorCode = 0
	// UNErrorCodeAttachmentInvalidFileSize - An attachment is too large.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNError/Code/attachmentInvalidFileSize
	UNErrorCodeAttachmentInvalidFileSize UNErrorCode = 0
	// UNErrorCodeAttachmentInvalidURL - The URL for an attachment was invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNError/Code/attachmentInvalidURL
	UNErrorCodeAttachmentInvalidURL UNErrorCode = 0
	// UNErrorCodeAttachmentMoveIntoDataStoreFailed - An error occurred when trying to move an attachment to the system data store.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNError/Code/attachmentMoveIntoDataStoreFailed
	UNErrorCodeAttachmentMoveIntoDataStoreFailed UNErrorCode = 0
	// UNErrorCodeAttachmentNotInDataStore - The specified attachment isn’t in the system data store.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNError/Code/attachmentNotInDataStore
	UNErrorCodeAttachmentNotInDataStore UNErrorCode = 0
	// UNErrorCodeAttachmentUnrecognizedType - The file type of an attachment isn’t supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNError/Code/attachmentUnrecognizedType
	UNErrorCodeAttachmentUnrecognizedType UNErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNError/Code/badgeInputInvalid
	UNErrorCodeBadgeInputInvalid UNErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNError/Code/contentProvidingInvalid
	UNErrorCodeContentProvidingInvalid UNErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNError/Code/contentProvidingObjectNotAllowed
	UNErrorCodeContentProvidingObjectNotAllowed UNErrorCode = 0
	// UNErrorCodeNotificationInvalidNoContent - The notification has no user-facing content, but should.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNError/Code/notificationInvalidNoContent
	UNErrorCodeNotificationInvalidNoContent UNErrorCode = 0
	// UNErrorCodeNotificationInvalidNoDate - The notification doesn’t have an associated date, but should.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNError/Code/notificationInvalidNoDate
	UNErrorCodeNotificationInvalidNoDate UNErrorCode = 0
	// UNErrorCodeNotificationsNotAllowed - Notifications aren’t allowed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNError/Code/notificationsNotAllowed
	UNErrorCodeNotificationsNotAllowed UNErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum UNNotificationActionOptions (3 cases) */
// UNNotificationActionOptions - The behaviors you can apply to an action.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationActionOptions
type UNNotificationActionOptions uint

const (
	// UNNotificationActionOptionAuthenticationRequired - The action can be performed only on an unlocked device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationActionOptions/authenticationRequired
	UNNotificationActionOptionAuthenticationRequired UNNotificationActionOptions = 0
	// UNNotificationActionOptionDestructive - The action performs a destructive task.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationActionOptions/destructive
	UNNotificationActionOptionDestructive UNNotificationActionOptions = 0
	// UNNotificationActionOptionForeground - The action causes the app to launch in the foreground.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationActionOptions/foreground
	UNNotificationActionOptionForeground UNNotificationActionOptions = 0
)

/* debug [enums.gen.go]: Processing enum UNNotificationCategoryOptions (5 cases) */
// UNNotificationCategoryOptions - Constants indicating how to handle notifications associated with this category.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategoryOptions
type UNNotificationCategoryOptions uint

const (
	// UNNotificationCategoryOptionAllowAnnouncement - An option that grants Siri permission to read incoming messages out loud when the user has a compatible audio output device connected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategoryOptions/allowAnnouncement
	UNNotificationCategoryOptionAllowAnnouncement UNNotificationCategoryOptions = 0
	// UNNotificationCategoryOptionAllowInCarPlay - Allow CarPlay to display notifications of this type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategoryOptions/allowInCarPlay
	UNNotificationCategoryOptionAllowInCarPlay UNNotificationCategoryOptions = 0
	// UNNotificationCategoryOptionCustomDismissAction - Send dismiss actions to the   object’s delegate for handling.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategoryOptions/customDismissAction
	UNNotificationCategoryOptionCustomDismissAction UNNotificationCategoryOptions = 0
	// UNNotificationCategoryOptionHiddenPreviewsShowSubtitle - Show the notification’s subtitle, even if the user has disabled notification previews for the app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategoryOptions/hiddenPreviewsShowSubtitle
	UNNotificationCategoryOptionHiddenPreviewsShowSubtitle UNNotificationCategoryOptions = 0
	// UNNotificationCategoryOptionHiddenPreviewsShowTitle - Show the notification’s title, even if the user has disabled notification previews for the app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategoryOptions/hiddenPreviewsShowTitle
	UNNotificationCategoryOptionHiddenPreviewsShowTitle UNNotificationCategoryOptions = 0
)

/* debug [enums.gen.go]: Processing enum UNNotificationInterruptionLevel (4 cases) */
// UNNotificationInterruptionLevel - Constants that indicate the importance and delivery timing of a notification.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationInterruptionLevel
type UNNotificationInterruptionLevel uint

const (
	// UNNotificationInterruptionLevelActive - The system presents the notification immediately, lights up the screen, and can play a sound.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationInterruptionLevel/active
	UNNotificationInterruptionLevelActive UNNotificationInterruptionLevel = 0
	// UNNotificationInterruptionLevelCritical - The system presents the notification immediately, lights up the screen, and bypasses the mute switch to play a sound.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationInterruptionLevel/critical
	UNNotificationInterruptionLevelCritical UNNotificationInterruptionLevel = 0
	// UNNotificationInterruptionLevelPassive - The system adds the notification to the notification list without lighting up the screen or playing a sound.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationInterruptionLevel/passive
	UNNotificationInterruptionLevelPassive UNNotificationInterruptionLevel = 0
	// UNNotificationInterruptionLevelTimeSensitive - The system presents the notification immediately, lights up the screen, can play a sound, and breaks through system notification controls.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationInterruptionLevel/timeSensitive
	UNNotificationInterruptionLevelTimeSensitive UNNotificationInterruptionLevel = 0
)

/* debug [enums.gen.go]: Processing enum UNNotificationPresentationOptions (5 cases) */
// UNNotificationPresentationOptions - Constants indicating how to present a notification in a foreground app.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationPresentationOptions
type UNNotificationPresentationOptions uint

const (
	// UNNotificationPresentationOptionAlert - Display the alert using the content provided by the notification.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationPresentationOptions/alert
	UNNotificationPresentationOptionAlert UNNotificationPresentationOptions = 0
	// UNNotificationPresentationOptionBadge - Apply the notification’s badge value to the app’s icon.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationPresentationOptions/badge
	UNNotificationPresentationOptionBadge UNNotificationPresentationOptions = 0
	// UNNotificationPresentationOptionBanner - Present the notification as a banner.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationPresentationOptions/banner
	UNNotificationPresentationOptionBanner UNNotificationPresentationOptions = 0
	// UNNotificationPresentationOptionList - Show the notification in Notification Center.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationPresentationOptions/list
	UNNotificationPresentationOptionList UNNotificationPresentationOptions = 0
	// UNNotificationPresentationOptionSound - Play the sound associated with the notification.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationPresentationOptions/sound
	UNNotificationPresentationOptionSound UNNotificationPresentationOptions = 0
)

/* debug [enums.gen.go]: Processing enum UNNotificationSetting (3 cases) */
// UNNotificationSetting - Constants that indicate the current status of a notification setting.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSetting
type UNNotificationSetting uint

const (
	// UNNotificationSettingDisabled - The setting is disabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSetting/disabled
	UNNotificationSettingDisabled UNNotificationSetting = 0
	// UNNotificationSettingEnabled - The setting is enabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSetting/enabled
	UNNotificationSettingEnabled UNNotificationSetting = 0
	// UNNotificationSettingNotSupported - The setting is not available to your app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSetting/notSupported
	UNNotificationSettingNotSupported UNNotificationSetting = 0
)

/* debug [enums.gen.go]: Processing enum UNShowPreviewsSetting (3 cases) */
// UNShowPreviewsSetting - Constants indicating the style previewing a notification’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNShowPreviewsSetting
type UNShowPreviewsSetting uint

const (
	// UNShowPreviewsSettingAlways - The notification’s content is always shown, even when the device is locked.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNShowPreviewsSetting/always
	UNShowPreviewsSettingAlways UNShowPreviewsSetting = 0
	// UNShowPreviewsSettingNever - The notification’s content is never shown, even when the device is unlocked
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNShowPreviewsSetting/never
	UNShowPreviewsSettingNever UNShowPreviewsSetting = 0
	// UNShowPreviewsSettingWhenAuthenticated - The notification’s content is shown only when the device is unlocked.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNShowPreviewsSetting/whenAuthenticated
	UNShowPreviewsSettingWhenAuthenticated UNShowPreviewsSetting = 0
)


