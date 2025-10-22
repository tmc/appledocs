// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications_test

import (
	"github.com/tmc/appledocs/generated/usernotifications"
)

// Suppress unused import errors
var _ = usernotifications.NewUNNotificationAction

// ExampleNewUNNotificationActionWithIdentifierTitleOptions demonstrates how to create a UNNotificationAction instance using NewUNNotificationActionWithIdentifierTitleOptions.
// Creates an action object by using the specified title and options.
func ExampleNewUNNotificationActionWithIdentifierTitleOptions() {
	_ = usernotifications.NewUNNotificationActionWithIdentifierTitleOptions(
		"identifier", // identifier string
		"title", // title string
		usernotifications.UNNotificationActionOptions{}, // options UNNotificationActionOptions
	)
	// Output:
}
// ExampleNewUNNotificationActionWithIdentifierTitleOptionsIcon demonstrates how to create a UNNotificationAction instance using NewUNNotificationActionWithIdentifierTitleOptionsIcon.
// Creates an action object by using the specified title, options, and icon.
func ExampleNewUNNotificationActionWithIdentifierTitleOptionsIcon() {
	_ = usernotifications.NewUNNotificationActionWithIdentifierTitleOptionsIcon(
		"identifier", // identifier string
		"title", // title string
		usernotifications.UNNotificationActionOptions{}, // options UNNotificationActionOptions
		usernotifications.UNNotificationActionIcon{}, // icon UNNotificationActionIcon
	)
	// Output:
}
