// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PMatchmakerViewControllerDelegate is the GKMatchmakerViewControllerDelegate protocol interface.
//
// An object that handles when the status of matchmaking changes.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.gamekit/documentation/GameKit/GKMatchmakerViewControllerDelegate
type PMatchmakerViewControllerDelegate interface {
	// Required methods
	MatchmakerViewControllerDidFailWithError(viewController IGKMatchmakerViewController, error_ objc.IObject /* cross-framework: Error */)/* debug [protocol_interface/required_method]: MatchmakerViewControllerDidFailWithError */
	MatchmakerViewControllerWasCancelled(viewController IGKMatchmakerViewController)/* debug [protocol_interface/required_method]: MatchmakerViewControllerWasCancelled */
	// Optional methods
	MatchmakerViewControllerDidFindMatch(viewController IGKMatchmakerViewController, match IGKMatch)
	HasMatchmakerViewControllerDidFindMatch() bool
	MatchmakerViewControllerDidFindHostedPlayers(viewController IGKMatchmakerViewController, players []Player)
	HasMatchmakerViewControllerDidFindHostedPlayers() bool
	MatchmakerViewControllerDidFindPlayers(viewController IGKMatchmakerViewController, playerIDs []string)
	HasMatchmakerViewControllerDidFindPlayers() bool
	MatchmakerViewControllerDidReceiveAcceptFromHostedPlayer(viewController IGKMatchmakerViewController, playerID objc.IObject /* cross-framework: NSString */)
	HasMatchmakerViewControllerDidReceiveAcceptFromHostedPlayer() bool
	MatchmakerViewControllerGetMatchPropertiesForRecipientWithCompletionHandler(viewController IGKMatchmakerViewController, recipient IGKPlayer, completionHandler unsafe.Pointer)
	HasMatchmakerViewControllerGetMatchPropertiesForRecipientWithCompletionHandler() bool
	MatchmakerViewControllerHostedPlayerDidAccept(viewController IGKMatchmakerViewController, player IGKPlayer)
	HasMatchmakerViewControllerHostedPlayerDidAccept() bool
}

// MatchmakerViewControllerDelegate is a delegate implementation builder for the PMatchmakerViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MatchmakerViewControllerDelegate struct {
	_MatchmakerViewControllerDidFindMatch func(viewController IGKMatchmakerViewController, match IGKMatch)
	_MatchmakerViewControllerDidFindHostedPlayers func(viewController IGKMatchmakerViewController, players []Player)
	_MatchmakerViewControllerDidFindPlayers func(viewController IGKMatchmakerViewController, playerIDs []string)
	_MatchmakerViewControllerDidReceiveAcceptFromHostedPlayer func(viewController IGKMatchmakerViewController, playerID objc.IObject /* cross-framework: NSString */)
	_MatchmakerViewControllerGetMatchPropertiesForRecipientWithCompletionHandler func(viewController IGKMatchmakerViewController, recipient IGKPlayer, completionHandler unsafe.Pointer)
	_MatchmakerViewControllerHostedPlayerDidAccept func(viewController IGKMatchmakerViewController, player IGKPlayer)
	_MatchmakerViewControllerDidFailWithError func(viewController IGKMatchmakerViewController, error_ objc.IObject /* cross-framework: Error */)
	_MatchmakerViewControllerWasCancelled func(viewController IGKMatchmakerViewController)
}

// SetMatchmakerViewControllerDidFindMatch sets the handler for the MatchmakerViewControllerDidFindMatch delegate method.
//
// Handles when the view controller finds players for a peer-to-peer match.
func (d *MatchmakerViewControllerDelegate) SetMatchmakerViewControllerDidFindMatch(f func(viewController IGKMatchmakerViewController, match IGKMatch)) {
	d._MatchmakerViewControllerDidFindMatch = f
}

// SetMatchmakerViewControllerDidFindHostedPlayers sets the handler for the MatchmakerViewControllerDidFindHostedPlayers delegate method.
//
// Handles when the view controller finds all requested players for a hosted match.
func (d *MatchmakerViewControllerDelegate) SetMatchmakerViewControllerDidFindHostedPlayers(f func(viewController IGKMatchmakerViewController, players []Player)) {
	d._MatchmakerViewControllerDidFindHostedPlayers = f
}

// SetMatchmakerViewControllerDidFindPlayers sets the handler for the MatchmakerViewControllerDidFindPlayers delegate method.
//
// Called when a hosted match is found.
func (d *MatchmakerViewControllerDelegate) SetMatchmakerViewControllerDidFindPlayers(f func(viewController IGKMatchmakerViewController, playerIDs []string)) {
	d._MatchmakerViewControllerDidFindPlayers = f
}

// SetMatchmakerViewControllerDidReceiveAcceptFromHostedPlayer sets the handler for the MatchmakerViewControllerDidReceiveAcceptFromHostedPlayer delegate method.
//
// Called when a player in a hosted match accepts the invitation.
func (d *MatchmakerViewControllerDelegate) SetMatchmakerViewControllerDidReceiveAcceptFromHostedPlayer(f func(viewController IGKMatchmakerViewController, playerID objc.IObject /* cross-framework: NSString */)) {
	d._MatchmakerViewControllerDidReceiveAcceptFromHostedPlayer = f
}

// SetMatchmakerViewControllerGetMatchPropertiesForRecipientWithCompletionHandler sets the handler for the MatchmakerViewControllerGetMatchPropertiesForRecipientWithCompletionHandler delegate method.
//
// Returns the properties for another player that the local player invites using the view controller interface.
func (d *MatchmakerViewControllerDelegate) SetMatchmakerViewControllerGetMatchPropertiesForRecipientWithCompletionHandler(f func(viewController IGKMatchmakerViewController, recipient IGKPlayer, completionHandler unsafe.Pointer)) {
	d._MatchmakerViewControllerGetMatchPropertiesForRecipientWithCompletionHandler = f
}

// SetMatchmakerViewControllerHostedPlayerDidAccept sets the handler for the MatchmakerViewControllerHostedPlayerDidAccept delegate method.
//
// Handles when a player in a hosted match accepts the invitation.
func (d *MatchmakerViewControllerDelegate) SetMatchmakerViewControllerHostedPlayerDidAccept(f func(viewController IGKMatchmakerViewController, player IGKPlayer)) {
	d._MatchmakerViewControllerHostedPlayerDidAccept = f
}

// SetMatchmakerViewControllerDidFailWithError sets the handler for the MatchmakerViewControllerDidFailWithError delegate method.
//
// Handles when a view controller encounters an error while finding players for a match.
func (d *MatchmakerViewControllerDelegate) SetMatchmakerViewControllerDidFailWithError(f func(viewController IGKMatchmakerViewController, error_ objc.IObject /* cross-framework: Error */)) {
	d._MatchmakerViewControllerDidFailWithError = f
}

// SetMatchmakerViewControllerWasCancelled sets the handler for the MatchmakerViewControllerWasCancelled delegate method.
//
// Handles when a player cancels a request to find players for a match.
func (d *MatchmakerViewControllerDelegate) SetMatchmakerViewControllerWasCancelled(f func(viewController IGKMatchmakerViewController)) {
	d._MatchmakerViewControllerWasCancelled = f
}

// MatchmakerViewControllerDidFindMatch implements the PMatchmakerViewControllerDelegate interface.
func (d *MatchmakerViewControllerDelegate) MatchmakerViewControllerDidFindMatch(viewController IGKMatchmakerViewController, match IGKMatch) {
	if d._MatchmakerViewControllerDidFindMatch != nil {
		d._MatchmakerViewControllerDidFindMatch(viewController, match)
	}
}

// HasMatchmakerViewControllerDidFindMatch returns true if a handler for MatchmakerViewControllerDidFindMatch has been set.
func (d *MatchmakerViewControllerDelegate) HasMatchmakerViewControllerDidFindMatch() bool {
	return d._MatchmakerViewControllerDidFindMatch != nil
}

// MatchmakerViewControllerDidFindHostedPlayers implements the PMatchmakerViewControllerDelegate interface.
func (d *MatchmakerViewControllerDelegate) MatchmakerViewControllerDidFindHostedPlayers(viewController IGKMatchmakerViewController, players []Player) {
	if d._MatchmakerViewControllerDidFindHostedPlayers != nil {
		d._MatchmakerViewControllerDidFindHostedPlayers(viewController, players)
	}
}

// HasMatchmakerViewControllerDidFindHostedPlayers returns true if a handler for MatchmakerViewControllerDidFindHostedPlayers has been set.
func (d *MatchmakerViewControllerDelegate) HasMatchmakerViewControllerDidFindHostedPlayers() bool {
	return d._MatchmakerViewControllerDidFindHostedPlayers != nil
}

// MatchmakerViewControllerDidFindPlayers implements the PMatchmakerViewControllerDelegate interface.
func (d *MatchmakerViewControllerDelegate) MatchmakerViewControllerDidFindPlayers(viewController IGKMatchmakerViewController, playerIDs []string) {
	if d._MatchmakerViewControllerDidFindPlayers != nil {
		d._MatchmakerViewControllerDidFindPlayers(viewController, playerIDs)
	}
}

// HasMatchmakerViewControllerDidFindPlayers returns true if a handler for MatchmakerViewControllerDidFindPlayers has been set.
func (d *MatchmakerViewControllerDelegate) HasMatchmakerViewControllerDidFindPlayers() bool {
	return d._MatchmakerViewControllerDidFindPlayers != nil
}

// MatchmakerViewControllerDidReceiveAcceptFromHostedPlayer implements the PMatchmakerViewControllerDelegate interface.
func (d *MatchmakerViewControllerDelegate) MatchmakerViewControllerDidReceiveAcceptFromHostedPlayer(viewController IGKMatchmakerViewController, playerID objc.IObject /* cross-framework: NSString */) {
	if d._MatchmakerViewControllerDidReceiveAcceptFromHostedPlayer != nil {
		d._MatchmakerViewControllerDidReceiveAcceptFromHostedPlayer(viewController, playerID)
	}
}

// HasMatchmakerViewControllerDidReceiveAcceptFromHostedPlayer returns true if a handler for MatchmakerViewControllerDidReceiveAcceptFromHostedPlayer has been set.
func (d *MatchmakerViewControllerDelegate) HasMatchmakerViewControllerDidReceiveAcceptFromHostedPlayer() bool {
	return d._MatchmakerViewControllerDidReceiveAcceptFromHostedPlayer != nil
}

// MatchmakerViewControllerGetMatchPropertiesForRecipientWithCompletionHandler implements the PMatchmakerViewControllerDelegate interface.
func (d *MatchmakerViewControllerDelegate) MatchmakerViewControllerGetMatchPropertiesForRecipientWithCompletionHandler(viewController IGKMatchmakerViewController, recipient IGKPlayer, completionHandler unsafe.Pointer) {
	if d._MatchmakerViewControllerGetMatchPropertiesForRecipientWithCompletionHandler != nil {
		d._MatchmakerViewControllerGetMatchPropertiesForRecipientWithCompletionHandler(viewController, recipient, completionHandler)
	}
}

// HasMatchmakerViewControllerGetMatchPropertiesForRecipientWithCompletionHandler returns true if a handler for MatchmakerViewControllerGetMatchPropertiesForRecipientWithCompletionHandler has been set.
func (d *MatchmakerViewControllerDelegate) HasMatchmakerViewControllerGetMatchPropertiesForRecipientWithCompletionHandler() bool {
	return d._MatchmakerViewControllerGetMatchPropertiesForRecipientWithCompletionHandler != nil
}

// MatchmakerViewControllerHostedPlayerDidAccept implements the PMatchmakerViewControllerDelegate interface.
func (d *MatchmakerViewControllerDelegate) MatchmakerViewControllerHostedPlayerDidAccept(viewController IGKMatchmakerViewController, player IGKPlayer) {
	if d._MatchmakerViewControllerHostedPlayerDidAccept != nil {
		d._MatchmakerViewControllerHostedPlayerDidAccept(viewController, player)
	}
}

// HasMatchmakerViewControllerHostedPlayerDidAccept returns true if a handler for MatchmakerViewControllerHostedPlayerDidAccept has been set.
func (d *MatchmakerViewControllerDelegate) HasMatchmakerViewControllerHostedPlayerDidAccept() bool {
	return d._MatchmakerViewControllerHostedPlayerDidAccept != nil
}

// MatchmakerViewControllerDidFailWithError implements the PMatchmakerViewControllerDelegate interface.
func (d *MatchmakerViewControllerDelegate) MatchmakerViewControllerDidFailWithError(viewController IGKMatchmakerViewController, error_ objc.IObject /* cross-framework: Error */) {
	if d._MatchmakerViewControllerDidFailWithError != nil {
		d._MatchmakerViewControllerDidFailWithError(viewController, error_)
	}
}

// HasMatchmakerViewControllerDidFailWithError returns true if a handler for MatchmakerViewControllerDidFailWithError has been set.
func (d *MatchmakerViewControllerDelegate) HasMatchmakerViewControllerDidFailWithError() bool {
	return d._MatchmakerViewControllerDidFailWithError != nil
}

// MatchmakerViewControllerWasCancelled implements the PMatchmakerViewControllerDelegate interface.
func (d *MatchmakerViewControllerDelegate) MatchmakerViewControllerWasCancelled(viewController IGKMatchmakerViewController) {
	if d._MatchmakerViewControllerWasCancelled != nil {
		d._MatchmakerViewControllerWasCancelled(viewController)
	}
}

// HasMatchmakerViewControllerWasCancelled returns true if a handler for MatchmakerViewControllerWasCancelled has been set.
func (d *MatchmakerViewControllerDelegate) HasMatchmakerViewControllerWasCancelled() bool {
	return d._MatchmakerViewControllerWasCancelled != nil
}
