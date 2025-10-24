// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PUNUserNotificationCenterDelegate is the UNUserNotificationCenterDelegate protocol interface.
//
// An interface for processing incoming notifications and responding to notification actions.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.14+
//   - tvOS 10.0+
//   - visionOS 1.0+
//   - watchOS 3.0+
//
// See: doc://com.apple.usernotifications/documentation/UserNotifications/UNUserNotificationCenterDelegate
type PUNUserNotificationCenterDelegate interface {
	// Optional methods
	UserNotificationCenterDidReceiveNotificationResponseWithCompletionHandler(center IUNUserNotificationCenter, response IUNNotificationResponse, completionHandler unsafe.Pointer)
	HasUserNotificationCenterDidReceiveNotificationResponseWithCompletionHandler() bool
	UserNotificationCenterOpenSettingsForNotification(center IUNUserNotificationCenter, notification IUNNotification)
	HasUserNotificationCenterOpenSettingsForNotification() bool
	UserNotificationCenterWillPresentNotificationWithCompletionHandler(center IUNUserNotificationCenter, notification IUNNotification, completionHandler unsafe.Pointer)
	HasUserNotificationCenterWillPresentNotificationWithCompletionHandler() bool
}

// UNUserNotificationCenterDelegate is a delegate implementation builder for the PUNUserNotificationCenterDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type UNUserNotificationCenterDelegate struct {
	_UserNotificationCenterDidReceiveNotificationResponseWithCompletionHandler func(center IUNUserNotificationCenter, response IUNNotificationResponse, completionHandler unsafe.Pointer)
	_UserNotificationCenterOpenSettingsForNotification func(center IUNUserNotificationCenter, notification IUNNotification)
	_UserNotificationCenterWillPresentNotificationWithCompletionHandler func(center IUNUserNotificationCenter, notification IUNNotification, completionHandler unsafe.Pointer)
}

// SetUserNotificationCenterDidReceiveNotificationResponseWithCompletionHandler sets the handler for the UserNotificationCenterDidReceiveNotificationResponseWithCompletionHandler delegate method.
//
// Asks the delegate to process the user’s response to a delivered notification.
func (d *UNUserNotificationCenterDelegate) SetUserNotificationCenterDidReceiveNotificationResponseWithCompletionHandler(f func(center IUNUserNotificationCenter, response IUNNotificationResponse, completionHandler unsafe.Pointer)) {
	d._UserNotificationCenterDidReceiveNotificationResponseWithCompletionHandler = f
}

// SetUserNotificationCenterOpenSettingsForNotification sets the handler for the UserNotificationCenterOpenSettingsForNotification delegate method.
//
// Asks the delegate to display the in-app notification settings.
func (d *UNUserNotificationCenterDelegate) SetUserNotificationCenterOpenSettingsForNotification(f func(center IUNUserNotificationCenter, notification IUNNotification)) {
	d._UserNotificationCenterOpenSettingsForNotification = f
}

// SetUserNotificationCenterWillPresentNotificationWithCompletionHandler sets the handler for the UserNotificationCenterWillPresentNotificationWithCompletionHandler delegate method.
//
// Asks the delegate how to handle a notification that arrived while the app was running in the foreground.
func (d *UNUserNotificationCenterDelegate) SetUserNotificationCenterWillPresentNotificationWithCompletionHandler(f func(center IUNUserNotificationCenter, notification IUNNotification, completionHandler unsafe.Pointer)) {
	d._UserNotificationCenterWillPresentNotificationWithCompletionHandler = f
}

// UserNotificationCenterDidReceiveNotificationResponseWithCompletionHandler implements the PUNUserNotificationCenterDelegate interface.
func (d *UNUserNotificationCenterDelegate) UserNotificationCenterDidReceiveNotificationResponseWithCompletionHandler(center IUNUserNotificationCenter, response IUNNotificationResponse, completionHandler unsafe.Pointer) {
	if d._UserNotificationCenterDidReceiveNotificationResponseWithCompletionHandler != nil {
		d._UserNotificationCenterDidReceiveNotificationResponseWithCompletionHandler(center, response, completionHandler)
	}
}

// HasUserNotificationCenterDidReceiveNotificationResponseWithCompletionHandler returns true if a handler for UserNotificationCenterDidReceiveNotificationResponseWithCompletionHandler has been set.
func (d *UNUserNotificationCenterDelegate) HasUserNotificationCenterDidReceiveNotificationResponseWithCompletionHandler() bool {
	return d._UserNotificationCenterDidReceiveNotificationResponseWithCompletionHandler != nil
}

// UserNotificationCenterOpenSettingsForNotification implements the PUNUserNotificationCenterDelegate interface.
func (d *UNUserNotificationCenterDelegate) UserNotificationCenterOpenSettingsForNotification(center IUNUserNotificationCenter, notification IUNNotification) {
	if d._UserNotificationCenterOpenSettingsForNotification != nil {
		d._UserNotificationCenterOpenSettingsForNotification(center, notification)
	}
}

// HasUserNotificationCenterOpenSettingsForNotification returns true if a handler for UserNotificationCenterOpenSettingsForNotification has been set.
func (d *UNUserNotificationCenterDelegate) HasUserNotificationCenterOpenSettingsForNotification() bool {
	return d._UserNotificationCenterOpenSettingsForNotification != nil
}

// UserNotificationCenterWillPresentNotificationWithCompletionHandler implements the PUNUserNotificationCenterDelegate interface.
func (d *UNUserNotificationCenterDelegate) UserNotificationCenterWillPresentNotificationWithCompletionHandler(center IUNUserNotificationCenter, notification IUNNotification, completionHandler unsafe.Pointer) {
	if d._UserNotificationCenterWillPresentNotificationWithCompletionHandler != nil {
		d._UserNotificationCenterWillPresentNotificationWithCompletionHandler(center, notification, completionHandler)
	}
}

// HasUserNotificationCenterWillPresentNotificationWithCompletionHandler returns true if a handler for UserNotificationCenterWillPresentNotificationWithCompletionHandler has been set.
func (d *UNUserNotificationCenterDelegate) HasUserNotificationCenterWillPresentNotificationWithCompletionHandler() bool {
	return d._UserNotificationCenterWillPresentNotificationWithCompletionHandler != nil
}
