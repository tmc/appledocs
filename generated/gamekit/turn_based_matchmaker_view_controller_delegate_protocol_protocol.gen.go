// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PTurnBasedMatchmakerViewControllerDelegate is the GKTurnBasedMatchmakerViewControllerDelegate protocol interface.
//
// A protocol that handles when the status of turn-based matchmaking changes.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.gamekit/documentation/GameKit/GKTurnBasedMatchmakerViewControllerDelegate
type PTurnBasedMatchmakerViewControllerDelegate interface {
	// Required methods
	TurnBasedMatchmakerViewControllerDidFailWithError(viewController IGKTurnBasedMatchmakerViewController, error_ objc.IObject /* cross-framework: Error */)/* debug [protocol_interface/required_method]: TurnBasedMatchmakerViewControllerDidFailWithError */
	TurnBasedMatchmakerViewControllerWasCancelled(viewController IGKTurnBasedMatchmakerViewController)/* debug [protocol_interface/required_method]: TurnBasedMatchmakerViewControllerWasCancelled */
	// Optional methods
	TurnBasedMatchmakerViewControllerDidFindMatch(viewController IGKTurnBasedMatchmakerViewController, match IGKTurnBasedMatch)
	HasTurnBasedMatchmakerViewControllerDidFindMatch() bool
	TurnBasedMatchmakerViewControllerPlayerQuitForMatch(viewController IGKTurnBasedMatchmakerViewController, match IGKTurnBasedMatch)
	HasTurnBasedMatchmakerViewControllerPlayerQuitForMatch() bool
}

// TurnBasedMatchmakerViewControllerDelegate is a delegate implementation builder for the PTurnBasedMatchmakerViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TurnBasedMatchmakerViewControllerDelegate struct {
	_TurnBasedMatchmakerViewControllerDidFindMatch func(viewController IGKTurnBasedMatchmakerViewController, match IGKTurnBasedMatch)
	_TurnBasedMatchmakerViewControllerPlayerQuitForMatch func(viewController IGKTurnBasedMatchmakerViewController, match IGKTurnBasedMatch)
	_TurnBasedMatchmakerViewControllerDidFailWithError func(viewController IGKTurnBasedMatchmakerViewController, error_ objc.IObject /* cross-framework: Error */)
	_TurnBasedMatchmakerViewControllerWasCancelled func(viewController IGKTurnBasedMatchmakerViewController)
}

// SetTurnBasedMatchmakerViewControllerDidFindMatch sets the handler for the TurnBasedMatchmakerViewControllerDidFindMatch delegate method.
//
// Handles when the view controller finds participants for a turn-based match.
func (d *TurnBasedMatchmakerViewControllerDelegate) SetTurnBasedMatchmakerViewControllerDidFindMatch(f func(viewController IGKTurnBasedMatchmakerViewController, match IGKTurnBasedMatch)) {
	d._TurnBasedMatchmakerViewControllerDidFindMatch = f
}

// SetTurnBasedMatchmakerViewControllerPlayerQuitForMatch sets the handler for the TurnBasedMatchmakerViewControllerPlayerQuitForMatch delegate method.
//
// Handles when a player quits the match.
func (d *TurnBasedMatchmakerViewControllerDelegate) SetTurnBasedMatchmakerViewControllerPlayerQuitForMatch(f func(viewController IGKTurnBasedMatchmakerViewController, match IGKTurnBasedMatch)) {
	d._TurnBasedMatchmakerViewControllerPlayerQuitForMatch = f
}

// SetTurnBasedMatchmakerViewControllerDidFailWithError sets the handler for the TurnBasedMatchmakerViewControllerDidFailWithError delegate method.
//
// Handles when an error occurs while the local player invites other players.
func (d *TurnBasedMatchmakerViewControllerDelegate) SetTurnBasedMatchmakerViewControllerDidFailWithError(f func(viewController IGKTurnBasedMatchmakerViewController, error_ objc.IObject /* cross-framework: Error */)) {
	d._TurnBasedMatchmakerViewControllerDidFailWithError = f
}

// SetTurnBasedMatchmakerViewControllerWasCancelled sets the handler for the TurnBasedMatchmakerViewControllerWasCancelled delegate method.
//
// Handles when the player dismisses the view controller without inviting players.
func (d *TurnBasedMatchmakerViewControllerDelegate) SetTurnBasedMatchmakerViewControllerWasCancelled(f func(viewController IGKTurnBasedMatchmakerViewController)) {
	d._TurnBasedMatchmakerViewControllerWasCancelled = f
}

// TurnBasedMatchmakerViewControllerDidFindMatch implements the PTurnBasedMatchmakerViewControllerDelegate interface.
func (d *TurnBasedMatchmakerViewControllerDelegate) TurnBasedMatchmakerViewControllerDidFindMatch(viewController IGKTurnBasedMatchmakerViewController, match IGKTurnBasedMatch) {
	if d._TurnBasedMatchmakerViewControllerDidFindMatch != nil {
		d._TurnBasedMatchmakerViewControllerDidFindMatch(viewController, match)
	}
}

// HasTurnBasedMatchmakerViewControllerDidFindMatch returns true if a handler for TurnBasedMatchmakerViewControllerDidFindMatch has been set.
func (d *TurnBasedMatchmakerViewControllerDelegate) HasTurnBasedMatchmakerViewControllerDidFindMatch() bool {
	return d._TurnBasedMatchmakerViewControllerDidFindMatch != nil
}

// TurnBasedMatchmakerViewControllerPlayerQuitForMatch implements the PTurnBasedMatchmakerViewControllerDelegate interface.
func (d *TurnBasedMatchmakerViewControllerDelegate) TurnBasedMatchmakerViewControllerPlayerQuitForMatch(viewController IGKTurnBasedMatchmakerViewController, match IGKTurnBasedMatch) {
	if d._TurnBasedMatchmakerViewControllerPlayerQuitForMatch != nil {
		d._TurnBasedMatchmakerViewControllerPlayerQuitForMatch(viewController, match)
	}
}

// HasTurnBasedMatchmakerViewControllerPlayerQuitForMatch returns true if a handler for TurnBasedMatchmakerViewControllerPlayerQuitForMatch has been set.
func (d *TurnBasedMatchmakerViewControllerDelegate) HasTurnBasedMatchmakerViewControllerPlayerQuitForMatch() bool {
	return d._TurnBasedMatchmakerViewControllerPlayerQuitForMatch != nil
}

// TurnBasedMatchmakerViewControllerDidFailWithError implements the PTurnBasedMatchmakerViewControllerDelegate interface.
func (d *TurnBasedMatchmakerViewControllerDelegate) TurnBasedMatchmakerViewControllerDidFailWithError(viewController IGKTurnBasedMatchmakerViewController, error_ objc.IObject /* cross-framework: Error */) {
	if d._TurnBasedMatchmakerViewControllerDidFailWithError != nil {
		d._TurnBasedMatchmakerViewControllerDidFailWithError(viewController, error_)
	}
}

// HasTurnBasedMatchmakerViewControllerDidFailWithError returns true if a handler for TurnBasedMatchmakerViewControllerDidFailWithError has been set.
func (d *TurnBasedMatchmakerViewControllerDelegate) HasTurnBasedMatchmakerViewControllerDidFailWithError() bool {
	return d._TurnBasedMatchmakerViewControllerDidFailWithError != nil
}

// TurnBasedMatchmakerViewControllerWasCancelled implements the PTurnBasedMatchmakerViewControllerDelegate interface.
func (d *TurnBasedMatchmakerViewControllerDelegate) TurnBasedMatchmakerViewControllerWasCancelled(viewController IGKTurnBasedMatchmakerViewController) {
	if d._TurnBasedMatchmakerViewControllerWasCancelled != nil {
		d._TurnBasedMatchmakerViewControllerWasCancelled(viewController)
	}
}

// HasTurnBasedMatchmakerViewControllerWasCancelled returns true if a handler for TurnBasedMatchmakerViewControllerWasCancelled has been set.
func (d *TurnBasedMatchmakerViewControllerDelegate) HasTurnBasedMatchmakerViewControllerWasCancelled() bool {
	return d._TurnBasedMatchmakerViewControllerWasCancelled != nil
}
