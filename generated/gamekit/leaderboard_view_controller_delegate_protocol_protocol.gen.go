// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PLeaderboardViewControllerDelegate is the GKLeaderboardViewControllerDelegate protocol interface.
//
// The   protocol is implemented by delegates of the   class. The delegate is called when the player dismisses the leaderboard.
//
// Availability:
//   - macOS 10.8+ (Deprecated in 10.10)
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// See: doc://com.apple.gamekit/documentation/GameKit/GKLeaderboardViewControllerDelegate
type PLeaderboardViewControllerDelegate interface {
	// Required methods
	LeaderboardViewControllerDidFinish(viewController IGKLeaderboardViewController)/* debug [protocol_interface/required_method]: LeaderboardViewControllerDidFinish */
}

// LeaderboardViewControllerDelegate is a delegate implementation builder for the PLeaderboardViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type LeaderboardViewControllerDelegate struct {
	_LeaderboardViewControllerDidFinish func(viewController IGKLeaderboardViewController)
}

// SetLeaderboardViewControllerDidFinish sets the handler for the LeaderboardViewControllerDidFinish delegate method.
//
// Called when the leaderboard view is dismissed.
func (d *LeaderboardViewControllerDelegate) SetLeaderboardViewControllerDidFinish(f func(viewController IGKLeaderboardViewController)) {
	d._LeaderboardViewControllerDidFinish = f
}

// LeaderboardViewControllerDidFinish implements the PLeaderboardViewControllerDelegate interface.
func (d *LeaderboardViewControllerDelegate) LeaderboardViewControllerDidFinish(viewController IGKLeaderboardViewController) {
	if d._LeaderboardViewControllerDidFinish != nil {
		d._LeaderboardViewControllerDidFinish(viewController)
	}
}

// HasLeaderboardViewControllerDidFinish returns true if a handler for LeaderboardViewControllerDidFinish has been set.
func (d *LeaderboardViewControllerDelegate) HasLeaderboardViewControllerDidFinish() bool {
	return d._LeaderboardViewControllerDidFinish != nil
}
