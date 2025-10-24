//go:build darwin && ios

// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
)

// iOS-only methods for UNLocationNotificationTrigger


// iOS-only properties

// The region used to determine when the system sends the notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNLocationNotificationTrigger/region
func (u_ UNLocationNotificationTrigger) Region() objc.IObject /* cross-framework: Region */ {
	rv := objc.Send[corelocation.Region](u_.ID, objc.Sel("region"))
	return rv
}




