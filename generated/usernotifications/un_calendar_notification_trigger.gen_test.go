// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications_test

import (
	"github.com/tmc/appledocs/generated/usernotifications"
)

// Suppress unused import errors
var _ = usernotifications.NewUNCalendarNotificationTrigger

// ExampleUNCalendarNotificationTrigger_NextTriggerDate demonstrates using NextTriggerDate on a UNCalendarNotificationTrigger instance.
// The next date at which the trigger conditions are met.
func ExampleUNCalendarNotificationTrigger_NextTriggerDate() {
	obj := usernotifications.NewUNCalendarNotificationTrigger()
	_ = obj.NextTriggerDate()
	// Output:
	}

