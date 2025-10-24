// Code generated from Apple documentation for UserNotificationsUI. DO NOT EDIT.

package usernotificationsui

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/usernotifications"
)

// PUNNotificationContentExtension is the UNNotificationContentExtension protocol interface.
//
// An object that presents a custom interface for a delivered local or remote   notification.
//
// Availability:
//   - Mac Catalyst 10.0+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 11.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.User-Notifications-UI/documentation/UserNotificationsUI/UNNotificationContentExtension
type PUNNotificationContentExtension interface {
	// Required methods
	DidReceiveNotification(notification usernotifications.UNNotification)
	// Optional methods
	DidReceiveNotificationResponseCompletionHandler(response usernotifications.UNNotificationResponse, completion unsafe.Pointer)
	HasDidReceiveNotificationResponseCompletionHandler() bool
	MediaPause()
	HasMediaPause() bool
	MediaPlay()
	HasMediaPlay() bool
}
