// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewNotification

// ExampleNewNotificationWithCoder demonstrates how to create a Notification instance using NewNotificationWithCoder.
// Initializes a notification with the data from an unarchiver.
func ExampleNewNotificationWithCoder() {
	_ = foundation.NewNotificationWithCoder(
		foundation.NSCoder{}, // coder NSCoder
	)
	// Output:
}
