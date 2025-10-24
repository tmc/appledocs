// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PFriendRequestComposeViewControllerDelegate is the GKFriendRequestComposeViewControllerDelegate protocol interface.
//
// The   protocol is implemented by delegates of the   class. The delegate is called when the player dismisses the friend request.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 4.2+ (Deprecated in 10.0)
//   - iPadOS 4.2+ (Deprecated in 10.0)
//   - macOS 10.8+ (Deprecated in 10.12)
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// See: doc://com.apple.gamekit/documentation/GameKit/GKFriendRequestComposeViewControllerDelegate
type PFriendRequestComposeViewControllerDelegate interface {
	// Required methods
	FriendRequestComposeViewControllerDidFinish(viewController IGKFriendRequestComposeViewController)/* debug [protocol_interface/required_method]: FriendRequestComposeViewControllerDidFinish */
}

// FriendRequestComposeViewControllerDelegate is a delegate implementation builder for the PFriendRequestComposeViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type FriendRequestComposeViewControllerDelegate struct {
	_FriendRequestComposeViewControllerDidFinish func(viewController IGKFriendRequestComposeViewController)
}

// SetFriendRequestComposeViewControllerDidFinish sets the handler for the FriendRequestComposeViewControllerDidFinish delegate method.
//
// Called when the player dismisses the request.
func (d *FriendRequestComposeViewControllerDelegate) SetFriendRequestComposeViewControllerDidFinish(f func(viewController IGKFriendRequestComposeViewController)) {
	d._FriendRequestComposeViewControllerDidFinish = f
}

// FriendRequestComposeViewControllerDidFinish implements the PFriendRequestComposeViewControllerDelegate interface.
func (d *FriendRequestComposeViewControllerDelegate) FriendRequestComposeViewControllerDidFinish(viewController IGKFriendRequestComposeViewController) {
	if d._FriendRequestComposeViewControllerDidFinish != nil {
		d._FriendRequestComposeViewControllerDidFinish(viewController)
	}
}

// HasFriendRequestComposeViewControllerDidFinish returns true if a handler for FriendRequestComposeViewControllerDidFinish has been set.
func (d *FriendRequestComposeViewControllerDelegate) HasFriendRequestComposeViewControllerDidFinish() bool {
	return d._FriendRequestComposeViewControllerDidFinish != nil
}
