// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PUserNotificationCenterDelegate is the NSUserNotificationCenterDelegate protocol interface.
//
// An interface that enables customizing the behavior of the default notification center.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/NSUserNotificationCenterDelegate
type PUserNotificationCenterDelegate interface {
	// Optional methods
	UserNotificationCenterShouldPresentNotification(center IUserNotificationCenter, notification IUserNotification) bool
	HasUserNotificationCenterShouldPresentNotification() bool
}

// UserNotificationCenterDelegate is a delegate implementation builder for the PUserNotificationCenterDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type UserNotificationCenterDelegate struct {
	_UserNotificationCenterShouldPresentNotification func(center IUserNotificationCenter, notification IUserNotification) bool
}

// SetUserNotificationCenterShouldPresentNotification sets the handler for the UserNotificationCenterShouldPresentNotification delegate method.
//
// Sent to the delegate when the user notification center has decided not to present your notification.
func (d *UserNotificationCenterDelegate) SetUserNotificationCenterShouldPresentNotification(f func(center IUserNotificationCenter, notification IUserNotification) bool) {
	d._UserNotificationCenterShouldPresentNotification = f
}

// UserNotificationCenterShouldPresentNotification implements the PUserNotificationCenterDelegate interface.
func (d *UserNotificationCenterDelegate) UserNotificationCenterShouldPresentNotification(center IUserNotificationCenter, notification IUserNotification) bool {
	if d._UserNotificationCenterShouldPresentNotification != nil {
		return d._UserNotificationCenterShouldPresentNotification(center, notification)
	}
	var zero bool
	return zero
}

// HasUserNotificationCenterShouldPresentNotification returns true if a handler for UserNotificationCenterShouldPresentNotification has been set.
func (d *UserNotificationCenterDelegate) HasUserNotificationCenterShouldPresentNotification() bool {
	return d._UserNotificationCenterShouldPresentNotification != nil
}
