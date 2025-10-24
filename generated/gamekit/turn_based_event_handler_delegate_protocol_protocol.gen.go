// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PTurnBasedEventHandlerDelegate is the GKTurnBasedEventHandlerDelegate protocol interface.
//
// The   protocol is implemented by an object to receive notifications events for turn-based matches. All methods are called on the main thread.
//
// Availability:
//   - macOS 10.8+ (Deprecated in 10.10)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 3.0+ (Deprecated in 3.0)
//
// See: doc://com.apple.gamekit/documentation/GameKit/GKTurnBasedEventHandlerDelegate
type PTurnBasedEventHandlerDelegate interface {
	// Required methods
	HandleInviteFromGameCenter(playersToInvite []string)/* debug [protocol_interface/required_method]: HandleInviteFromGameCenter */
	HandleTurnEventForMatchDidBecomeActive(match IGKTurnBasedMatch, didBecomeActive bool)/* debug [protocol_interface/required_method]: HandleTurnEventForMatchDidBecomeActive */
	// Optional methods
	HandleMatchEnded(match IGKTurnBasedMatch)
	HasHandleMatchEnded() bool
	HandleTurnEventForMatch(match IGKTurnBasedMatch)
	HasHandleTurnEventForMatch() bool
}

// TurnBasedEventHandlerDelegate is a delegate implementation builder for the PTurnBasedEventHandlerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TurnBasedEventHandlerDelegate struct {
	_HandleMatchEnded func(match IGKTurnBasedMatch)
	_HandleTurnEventForMatch func(match IGKTurnBasedMatch)
	_HandleInviteFromGameCenter func(playersToInvite []string)
	_HandleTurnEventForMatchDidBecomeActive func(match IGKTurnBasedMatch, didBecomeActive bool)
}

// SetHandleMatchEnded sets the handler for the HandleMatchEnded delegate method.
//
// Sent to the delegate when a match the local player is participating in has ended.
func (d *TurnBasedEventHandlerDelegate) SetHandleMatchEnded(f func(match IGKTurnBasedMatch)) {
	d._HandleMatchEnded = f
}

// SetHandleTurnEventForMatch sets the handler for the HandleTurnEventForMatch delegate method.
//
// Sent to the delegate when it is the local player’s turn to act in a turn-based match.
func (d *TurnBasedEventHandlerDelegate) SetHandleTurnEventForMatch(f func(match IGKTurnBasedMatch)) {
	d._HandleTurnEventForMatch = f
}

// SetHandleInviteFromGameCenter sets the handler for the HandleInviteFromGameCenter delegate method.
//
// Sent to the delegate when the local player receives an invitation to join a new turn-based match.
func (d *TurnBasedEventHandlerDelegate) SetHandleInviteFromGameCenter(f func(playersToInvite []string)) {
	d._HandleInviteFromGameCenter = f
}

// SetHandleTurnEventForMatchDidBecomeActive sets the handler for the HandleTurnEventForMatchDidBecomeActive delegate method.
//
// Sent to the delegate when it is the local player’s turn to act in a turn-based match.
func (d *TurnBasedEventHandlerDelegate) SetHandleTurnEventForMatchDidBecomeActive(f func(match IGKTurnBasedMatch, didBecomeActive bool)) {
	d._HandleTurnEventForMatchDidBecomeActive = f
}

// HandleMatchEnded implements the PTurnBasedEventHandlerDelegate interface.
func (d *TurnBasedEventHandlerDelegate) HandleMatchEnded(match IGKTurnBasedMatch) {
	if d._HandleMatchEnded != nil {
		d._HandleMatchEnded(match)
	}
}

// HasHandleMatchEnded returns true if a handler for HandleMatchEnded has been set.
func (d *TurnBasedEventHandlerDelegate) HasHandleMatchEnded() bool {
	return d._HandleMatchEnded != nil
}

// HandleTurnEventForMatch implements the PTurnBasedEventHandlerDelegate interface.
func (d *TurnBasedEventHandlerDelegate) HandleTurnEventForMatch(match IGKTurnBasedMatch) {
	if d._HandleTurnEventForMatch != nil {
		d._HandleTurnEventForMatch(match)
	}
}

// HasHandleTurnEventForMatch returns true if a handler for HandleTurnEventForMatch has been set.
func (d *TurnBasedEventHandlerDelegate) HasHandleTurnEventForMatch() bool {
	return d._HandleTurnEventForMatch != nil
}

// HandleInviteFromGameCenter implements the PTurnBasedEventHandlerDelegate interface.
func (d *TurnBasedEventHandlerDelegate) HandleInviteFromGameCenter(playersToInvite []string) {
	if d._HandleInviteFromGameCenter != nil {
		d._HandleInviteFromGameCenter(playersToInvite)
	}
}

// HasHandleInviteFromGameCenter returns true if a handler for HandleInviteFromGameCenter has been set.
func (d *TurnBasedEventHandlerDelegate) HasHandleInviteFromGameCenter() bool {
	return d._HandleInviteFromGameCenter != nil
}

// HandleTurnEventForMatchDidBecomeActive implements the PTurnBasedEventHandlerDelegate interface.
func (d *TurnBasedEventHandlerDelegate) HandleTurnEventForMatchDidBecomeActive(match IGKTurnBasedMatch, didBecomeActive bool) {
	if d._HandleTurnEventForMatchDidBecomeActive != nil {
		d._HandleTurnEventForMatchDidBecomeActive(match, didBecomeActive)
	}
}

// HasHandleTurnEventForMatchDidBecomeActive returns true if a handler for HandleTurnEventForMatchDidBecomeActive has been set.
func (d *TurnBasedEventHandlerDelegate) HasHandleTurnEventForMatchDidBecomeActive() bool {
	return d._HandleTurnEventForMatchDidBecomeActive != nil
}
