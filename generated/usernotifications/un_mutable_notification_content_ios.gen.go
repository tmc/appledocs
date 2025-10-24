//go:build darwin && ios

// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for UNMutableNotificationContent

// iOS-only properties

// The name of the image or storyboard to use when your app launches because of the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent/launchImageName
func (u_ UNMutableNotificationContent) LaunchImageName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](u_.ID, objc.Sel("launchImageName"))
	return rv
}
func (u_ UNMutableNotificationContent) SetLaunchImageName(value objc.IObject /* cross-framework: NSString */) {
	u_.ID.Send(objc.RegisterName("setLaunchImageName:"), value)
}
