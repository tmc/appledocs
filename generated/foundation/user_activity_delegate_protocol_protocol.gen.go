// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PUserActivityDelegate is the NSUserActivityDelegate protocol interface.
//
// The interface through which a user activity instance notifies its delegate of updates.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 3.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/NSUserActivityDelegate
type PUserActivityDelegate interface {
	// Optional methods
	UserActivityWasContinued(userActivity IUserActivity)
	HasUserActivityWasContinued() bool
}

// UserActivityDelegate is a delegate implementation builder for the PUserActivityDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type UserActivityDelegate struct {
	_UserActivityWasContinued func(userActivity IUserActivity)
}

// SetUserActivityWasContinued sets the handler for the UserActivityWasContinued delegate method.
//
// Notifies the delegate that the user activity was continued on another device.
func (d *UserActivityDelegate) SetUserActivityWasContinued(f func(userActivity IUserActivity)) {
	d._UserActivityWasContinued = f
}

// UserActivityWasContinued implements the PUserActivityDelegate interface.
func (d *UserActivityDelegate) UserActivityWasContinued(userActivity IUserActivity) {
	if d._UserActivityWasContinued != nil {
		d._UserActivityWasContinued(userActivity)
	}
}

// HasUserActivityWasContinued returns true if a handler for UserActivityWasContinued has been set.
func (d *UserActivityDelegate) HasUserActivityWasContinued() bool {
	return d._UserActivityWasContinued != nil
}
