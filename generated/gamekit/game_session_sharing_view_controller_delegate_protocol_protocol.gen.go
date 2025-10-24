// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PGameSessionSharingViewControllerDelegate is the GKGameSessionSharingViewControllerDelegate protocol interface.
//
// A protocol you implement to respond to changes to a sharing user interface.
//
// Availability:
//   - tvOS 10.0+ (Deprecated in 12.0)
//
// See: doc://com.apple.gamekit/documentation/GameKit/GKGameSessionSharingViewControllerDelegate
type PGameSessionSharingViewControllerDelegate interface {
	// Required methods
	SharingViewControllerDidFinishWithError(viewController IGKGameSessionSharingViewController, error_ objc.IObject /* cross-framework: Error */)/* debug [protocol_interface/required_method]: SharingViewControllerDidFinishWithError */
}

// GameSessionSharingViewControllerDelegate is a delegate implementation builder for the PGameSessionSharingViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type GameSessionSharingViewControllerDelegate struct {
	_SharingViewControllerDidFinishWithError func(viewController IGKGameSessionSharingViewController, error_ objc.IObject /* cross-framework: Error */)
}

// SetSharingViewControllerDidFinishWithError sets the handler for the SharingViewControllerDidFinishWithError delegate method.
//
// Indicates the sharing view controller is ready to be dismissed.
func (d *GameSessionSharingViewControllerDelegate) SetSharingViewControllerDidFinishWithError(f func(viewController IGKGameSessionSharingViewController, error_ objc.IObject /* cross-framework: Error */)) {
	d._SharingViewControllerDidFinishWithError = f
}

// SharingViewControllerDidFinishWithError implements the PGameSessionSharingViewControllerDelegate interface.
func (d *GameSessionSharingViewControllerDelegate) SharingViewControllerDidFinishWithError(viewController IGKGameSessionSharingViewController, error_ objc.IObject /* cross-framework: Error */) {
	if d._SharingViewControllerDidFinishWithError != nil {
		d._SharingViewControllerDidFinishWithError(viewController, error_)
	}
}

// HasSharingViewControllerDidFinishWithError returns true if a handler for SharingViewControllerDidFinishWithError has been set.
func (d *GameSessionSharingViewControllerDelegate) HasSharingViewControllerDidFinishWithError() bool {
	return d._SharingViewControllerDidFinishWithError != nil
}
