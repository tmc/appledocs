// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications_test

import (
	"github.com/tmc/appledocs/generated/usernotifications"
)

// Suppress unused import errors
var _ = usernotifications.NewUNNotificationActionIcon


// ExampleNewUNNotificationActionIconWithSystemImageName demonstrates how to create a UNNotificationActionIcon instance using NewUNNotificationActionIconWithSystemImageName.
// Creates an action icon by using a system symbol image.
func ExampleNewUNNotificationActionIconWithSystemImageName() {
	_ = usernotifications.NewUNNotificationActionIconWithSystemImageName(
		"systemImageName", // systemImageName string
	)
	// Output:
}

// ExampleNewUNNotificationActionIconWithTemplateImageName demonstrates how to create a UNNotificationActionIcon instance using NewUNNotificationActionIconWithTemplateImageName.
// Creates an action icon based on an image in your app’s bundle, preferably in an asset catalog.
func ExampleNewUNNotificationActionIconWithTemplateImageName() {
	_ = usernotifications.NewUNNotificationActionIconWithTemplateImageName(
		"templateImageName", // templateImageName string
	)
	// Output:
}


