//go:build darwin && ios

// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for UNNotificationSettings

// iOS-only properties

// The setting that indicates whether Siri can announce your app’s notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/announcementSetting
func (u_ UNNotificationSettings) AnnouncementSetting() UNNotificationSetting {
	rv := objc.Send[UNNotificationSetting](u_.ID, objc.Sel("announcementSetting"))
	return rv
}

// The setting that indicates whether your app’s notifications appear in CarPlay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/carPlaySetting
func (u_ UNNotificationSettings) CarPlaySetting() UNNotificationSetting {
	rv := objc.Send[UNNotificationSetting](u_.ID, objc.Sel("carPlaySetting"))
	return rv
}
