// Code generated from Apple documentation for UserNotificationsUI. DO NOT EDIT.

package usernotificationsui

/* debug [enums.gen.go]: Generating 2 enums for UserNotificationsUI */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum UNNotificationContentExtensionMediaPlayPauseButtonType (3 cases) */
// UNNotificationContentExtensionMediaPlayPauseButtonType - Constants indicating the type of media button to display.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotificationsUI/UNNotificationContentExtensionMediaPlayPauseButtonType
type UNNotificationContentExtensionMediaPlayPauseButtonType uint

const (
	// UNNotificationContentExtensionMediaPlayPauseButtonTypeDefault - A standard play/pause button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotificationsUI/UNNotificationContentExtensionMediaPlayPauseButtonType/default
	UNNotificationContentExtensionMediaPlayPauseButtonTypeDefault UNNotificationContentExtensionMediaPlayPauseButtonType = 0
	// UNNotificationContentExtensionMediaPlayPauseButtonTypeNone - No media button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotificationsUI/UNNotificationContentExtensionMediaPlayPauseButtonType/none
	UNNotificationContentExtensionMediaPlayPauseButtonTypeNone UNNotificationContentExtensionMediaPlayPauseButtonType = 0
	// UNNotificationContentExtensionMediaPlayPauseButtonTypeOverlay - A partially transparent play/pause button that is layered on top of your   media content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotificationsUI/UNNotificationContentExtensionMediaPlayPauseButtonType/overlay
	UNNotificationContentExtensionMediaPlayPauseButtonTypeOverlay UNNotificationContentExtensionMediaPlayPauseButtonType = 0
)

/* debug [enums.gen.go]: Processing enum UNNotificationContentExtensionResponseOption (3 cases) */
// UNNotificationContentExtensionResponseOption - Constants indicating the preferred response to a notification.
//
// [Full Topic]: https://developer.apple.com/documentation/UserNotificationsUI/UNNotificationContentExtensionResponseOption
type UNNotificationContentExtensionResponseOption uint

const (
	// UNNotificationContentExtensionResponseOptionDismiss - Dismiss the notification interface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotificationsUI/UNNotificationContentExtensionResponseOption/dismiss
	UNNotificationContentExtensionResponseOptionDismiss UNNotificationContentExtensionResponseOption = 0
	// UNNotificationContentExtensionResponseOptionDismissAndForwardAction - Dismiss the notification interface and forward the notification to the app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotificationsUI/UNNotificationContentExtensionResponseOption/dismissAndForwardAction
	UNNotificationContentExtensionResponseOptionDismissAndForwardAction UNNotificationContentExtensionResponseOption = 0
	// UNNotificationContentExtensionResponseOptionDoNotDismiss - Don’t dismiss the notification interface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/UserNotificationsUI/UNNotificationContentExtensionResponseOption/doNotDismiss
	UNNotificationContentExtensionResponseOptionDoNotDismiss UNNotificationContentExtensionResponseOption = 0
)
