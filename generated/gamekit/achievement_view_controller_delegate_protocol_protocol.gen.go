// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PAchievementViewControllerDelegate is the GKAchievementViewControllerDelegate protocol interface.
//
// An object implementing the   protocol is called when the user dismisses the achievements view controller. Typically, this protocol is implemented by the object in your game that originally displayed the achievements user interface.
//
// Availability:
//   - macOS 10.8+ (Deprecated in 10.10)
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// See: doc://com.apple.gamekit/documentation/GameKit/GKAchievementViewControllerDelegate
type PAchievementViewControllerDelegate interface {
	// Required methods
	AchievementViewControllerDidFinish(viewController IGKAchievementViewController)/* debug [protocol_interface/required_method]: AchievementViewControllerDidFinish */
}

// AchievementViewControllerDelegate is a delegate implementation builder for the PAchievementViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type AchievementViewControllerDelegate struct {
	_AchievementViewControllerDidFinish func(viewController IGKAchievementViewController)
}

// SetAchievementViewControllerDidFinish sets the handler for the AchievementViewControllerDidFinish delegate method.
//
// Called when the user dismisses the achievements user interface.
func (d *AchievementViewControllerDelegate) SetAchievementViewControllerDidFinish(f func(viewController IGKAchievementViewController)) {
	d._AchievementViewControllerDidFinish = f
}

// AchievementViewControllerDidFinish implements the PAchievementViewControllerDelegate interface.
func (d *AchievementViewControllerDelegate) AchievementViewControllerDidFinish(viewController IGKAchievementViewController) {
	if d._AchievementViewControllerDidFinish != nil {
		d._AchievementViewControllerDidFinish(viewController)
	}
}

// HasAchievementViewControllerDidFinish returns true if a handler for AchievementViewControllerDidFinish has been set.
func (d *AchievementViewControllerDelegate) HasAchievementViewControllerDidFinish() bool {
	return d._AchievementViewControllerDidFinish != nil
}
