// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
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

// UserNotificationCenterDelegateObject wraps an existing Objective-C object that conforms to the PUserNotificationCenterDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type UserNotificationCenterDelegateObject struct {
	objectivec.Object
}

// NewUserNotificationCenterDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSUserNotificationCenterDelegate protocol.
func NewUserNotificationCenterDelegateObject(obj objectivec.Object) *UserNotificationCenterDelegateObject {
	return &UserNotificationCenterDelegateObject{obj}
}

// Make sure UserNotificationCenterDelegateObject implements PUserNotificationCenterDelegate.
var _ PUserNotificationCenterDelegate = (*UserNotificationCenterDelegateObject)(nil)

// UserNotificationCenterShouldPresentNotification implements the PUserNotificationCenterDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *UserNotificationCenterDelegateObject) UserNotificationCenterShouldPresentNotification(center IUserNotificationCenter, notification IUserNotification) bool {
	return objc.Send[bool](o.ID, objc.Sel("userNotificationCenter:shouldPresentNotification:"), center, notification)
}

// HasUserNotificationCenterShouldPresentNotification returns true; this is a placeholder for optional method checks.
func (o *UserNotificationCenterDelegateObject) HasUserNotificationCenterShouldPresentNotification() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
