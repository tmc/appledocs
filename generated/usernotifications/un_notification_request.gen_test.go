// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications_test

import (
	"github.com/tmc/appledocs/generated/usernotifications"
)

// Suppress unused import errors
var _ = usernotifications.NewUNNotificationRequest

// ExampleNewUNNotificationRequestWithIdentifierContentTrigger demonstrates how to create a UNNotificationRequest instance using NewUNNotificationRequestWithIdentifierContentTrigger.
// Creates a notification request object that you use to schedule a notification.
func ExampleNewUNNotificationRequestWithIdentifierContentTrigger() {
	_ = usernotifications.NewUNNotificationRequestWithIdentifierContentTrigger(
		"identifier", // identifier string
		usernotifications.UNNotificationContent{}, // content UNNotificationContent
		usernotifications.UNNotificationTrigger{}, // trigger UNNotificationTrigger
	)
	// Output:
}
