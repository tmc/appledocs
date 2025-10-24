// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PGameCenterControllerDelegate is the GKGameCenterControllerDelegate protocol interface.
//
// The delegate that GameKit calls when the player dismisses the dashboard.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 26.0)
//   - iOS 6.0+ (Deprecated in 26.0)
//   - iPadOS 6.0+ (Deprecated in 26.0)
//   - macOS 10.8+ (Deprecated in 26.0)
//   - tvOS 9.0+ (Deprecated in 26.0)
//   - visionOS 1.0+ (Deprecated in 26.0)
//
// See: doc://com.apple.gamekit/documentation/GameKit/GKGameCenterControllerDelegate
type PGameCenterControllerDelegate interface {
	// Required methods
	GameCenterViewControllerDidFinish(gameCenterViewController IGKGameCenterViewController)/* debug [protocol_interface/required_method]: GameCenterViewControllerDidFinish */
}

// GameCenterControllerDelegate is a delegate implementation builder for the PGameCenterControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type GameCenterControllerDelegate struct {
	_GameCenterViewControllerDidFinish func(gameCenterViewController IGKGameCenterViewController)
}

// SetGameCenterViewControllerDidFinish sets the handler for the GameCenterViewControllerDidFinish delegate method.
//
// Handles when the player dismisses the dashboard.
func (d *GameCenterControllerDelegate) SetGameCenterViewControllerDidFinish(f func(gameCenterViewController IGKGameCenterViewController)) {
	d._GameCenterViewControllerDidFinish = f
}

// GameCenterViewControllerDidFinish implements the PGameCenterControllerDelegate interface.
func (d *GameCenterControllerDelegate) GameCenterViewControllerDidFinish(gameCenterViewController IGKGameCenterViewController) {
	if d._GameCenterViewControllerDidFinish != nil {
		d._GameCenterViewControllerDidFinish(gameCenterViewController)
	}
}

// HasGameCenterViewControllerDidFinish returns true if a handler for GameCenterViewControllerDidFinish has been set.
func (d *GameCenterControllerDelegate) HasGameCenterViewControllerDidFinish() bool {
	return d._GameCenterViewControllerDidFinish != nil
}
