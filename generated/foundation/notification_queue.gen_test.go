// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewNotificationQueue

// ExampleNewNotificationQueueWithNotificationCenter demonstrates how to create a NotificationQueue instance using NewNotificationQueueWithNotificationCenter.
// Initializes and returns a notification queue for the specified notification center.
func ExampleNewNotificationQueueWithNotificationCenter() {
	_ = foundation.NewNotificationQueueWithNotificationCenter(
		foundation.NSNotificationCenter{}, // notificationCenter NSNotificationCenter
	)
	// Output:
}
