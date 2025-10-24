// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications_test

import (
	"github.com/tmc/appledocs/generated/usernotifications"
)

// Suppress unused import errors
var _ = usernotifications.NewUNNotificationSound

// ExampleNewUNNotificationSoundNamed demonstrates how to create a UNNotificationSound instance using NewUNNotificationSoundNamed.
// Creates a sound object that represents a custom sound file.
func ExampleNewUNNotificationSoundNamed() {
	_ = usernotifications.NewUNNotificationSoundNamed(
		usernotifications.UNNotificationSoundName /* typedef */ {}, // name UNNotificationSoundName /* typedef */
	)
	// Output:
}
