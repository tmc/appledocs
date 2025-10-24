//go:build darwin && ios

// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/gameplaykit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for UNNotificationResponse


// iOS-only properties

// The scene where the system reflects the user’s response to a notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UserNotifications/UNNotificationResponse/targetScene
func (u_ UNNotificationResponse) TargetScene() objc.IObject /* cross-framework: Scene */ {
	rv := objc.Send[gameplaykit.Scene](u_.ID, objc.Sel("targetScene"))
	return rv
}





