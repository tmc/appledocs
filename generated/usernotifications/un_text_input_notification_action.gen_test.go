// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications_test

import (
	"github.com/tmc/appledocs/generated/usernotifications"
)

// Suppress unused import errors
var _ = usernotifications.NewUNTextInputNotificationAction

// ExampleNewUNTextInputNotificationActionWithIdentifierTitleOptionsIconTextInputButtonTitleTextInputPlaceholder demonstrates how to create a UNTextInputNotificationAction instance using NewUNTextInputNotificationActionWithIdentifierTitleOptionsIconTextInputButtonTitleTextInputPlaceholder.
// Creates an action object with an icon that accepts text input from the user.
func ExampleNewUNTextInputNotificationActionWithIdentifierTitleOptionsIconTextInputButtonTitleTextInputPlaceholder() {
	_ = usernotifications.NewUNTextInputNotificationActionWithIdentifierTitleOptionsIconTextInputButtonTitleTextInputPlaceholder(
		"identifier", // identifier string
		"title", // title string
		usernotifications.UNNotificationActionOptions{}, // options UNNotificationActionOptions
		usernotifications.UNNotificationActionIcon{}, // icon UNNotificationActionIcon
		"textInputButtonTitle", // textInputButtonTitle string
		"textInputPlaceholder", // textInputPlaceholder string
	)
	// Output:
}
// ExampleNewUNTextInputNotificationActionWithIdentifierTitleOptionsTextInputButtonTitleTextInputPlaceholder demonstrates how to create a UNTextInputNotificationAction instance using NewUNTextInputNotificationActionWithIdentifierTitleOptionsTextInputButtonTitleTextInputPlaceholder.
// Creates an action object that accepts text input from the user.
func ExampleNewUNTextInputNotificationActionWithIdentifierTitleOptionsTextInputButtonTitleTextInputPlaceholder() {
	_ = usernotifications.NewUNTextInputNotificationActionWithIdentifierTitleOptionsTextInputButtonTitleTextInputPlaceholder(
		"identifier", // identifier string
		"title", // title string
		usernotifications.UNNotificationActionOptions{}, // options UNNotificationActionOptions
		"textInputButtonTitle", // textInputButtonTitle string
		"textInputPlaceholder", // textInputPlaceholder string
	)
	// Output:
}
