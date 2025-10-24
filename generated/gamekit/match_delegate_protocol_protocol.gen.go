// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PMatchDelegate is the GKMatchDelegate protocol interface.
//
// An object that receives connection status and data transmitted in a multiplayer game.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.gamekit/documentation/GameKit/GKMatchDelegate
type PMatchDelegate interface {
	// Optional methods
	MatchDidFailWithError(match IGKMatch, error_ objc.IObject /* cross-framework: Error */)
	HasMatchDidFailWithError() bool
	MatchDidReceiveDataForRecipientFromRemotePlayer(match IGKMatch, data objc.IObject /* cross-framework: NSData */, recipient IGKPlayer, player IGKPlayer)
	HasMatchDidReceiveDataForRecipientFromRemotePlayer() bool
	MatchDidReceiveDataFromPlayer(match IGKMatch, data objc.IObject /* cross-framework: NSData */, playerID objc.IObject /* cross-framework: NSString */)
	HasMatchDidReceiveDataFromPlayer() bool
	MatchDidReceiveDataFromRemotePlayer(match IGKMatch, data objc.IObject /* cross-framework: NSData */, player IGKPlayer)
	HasMatchDidReceiveDataFromRemotePlayer() bool
	MatchPlayerDidChangeState(match IGKMatch, playerID objc.IObject /* cross-framework: NSString */, state PlayerConnectionState)
	HasMatchPlayerDidChangeState() bool
	MatchPlayerDidChangeConnectionState(match IGKMatch, player IGKPlayer, state PlayerConnectionState)
	HasMatchPlayerDidChangeConnectionState() bool
	MatchShouldReinviteDisconnectedPlayer(match IGKMatch, player IGKPlayer) bool
	HasMatchShouldReinviteDisconnectedPlayer() bool
	MatchShouldReinvitePlayer(match IGKMatch, playerID objc.IObject /* cross-framework: NSString */) bool
	HasMatchShouldReinvitePlayer() bool
}

// MatchDelegate is a delegate implementation builder for the PMatchDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MatchDelegate struct {
	_MatchDidFailWithError func(match IGKMatch, error_ objc.IObject /* cross-framework: Error */)
	_MatchDidReceiveDataForRecipientFromRemotePlayer func(match IGKMatch, data objc.IObject /* cross-framework: NSData */, recipient IGKPlayer, player IGKPlayer)
	_MatchDidReceiveDataFromPlayer func(match IGKMatch, data objc.IObject /* cross-framework: NSData */, playerID objc.IObject /* cross-framework: NSString */)
	_MatchDidReceiveDataFromRemotePlayer func(match IGKMatch, data objc.IObject /* cross-framework: NSData */, player IGKPlayer)
	_MatchPlayerDidChangeState func(match IGKMatch, playerID objc.IObject /* cross-framework: NSString */, state PlayerConnectionState)
	_MatchPlayerDidChangeConnectionState func(match IGKMatch, player IGKPlayer, state PlayerConnectionState)
	_MatchShouldReinviteDisconnectedPlayer func(match IGKMatch, player IGKPlayer) bool
	_MatchShouldReinvitePlayer func(match IGKMatch, playerID objc.IObject /* cross-framework: NSString */) bool
}

// SetMatchDidFailWithError sets the handler for the MatchDidFailWithError delegate method.
//
// Handles the local player’s connection errors to a match.
func (d *MatchDelegate) SetMatchDidFailWithError(f func(match IGKMatch, error_ objc.IObject /* cross-framework: Error */)) {
	d._MatchDidFailWithError = f
}

// SetMatchDidReceiveDataForRecipientFromRemotePlayer sets the handler for the MatchDidReceiveDataForRecipientFromRemotePlayer delegate method.
//
// Processes the data sent from one player to another.
func (d *MatchDelegate) SetMatchDidReceiveDataForRecipientFromRemotePlayer(f func(match IGKMatch, data objc.IObject /* cross-framework: NSData */, recipient IGKPlayer, player IGKPlayer)) {
	d._MatchDidReceiveDataForRecipientFromRemotePlayer = f
}

// SetMatchDidReceiveDataFromPlayer sets the handler for the MatchDidReceiveDataFromPlayer delegate method.
//
// Handles when a player receives data in a match.
func (d *MatchDelegate) SetMatchDidReceiveDataFromPlayer(f func(match IGKMatch, data objc.IObject /* cross-framework: NSData */, playerID objc.IObject /* cross-framework: NSString */)) {
	d._MatchDidReceiveDataFromPlayer = f
}

// SetMatchDidReceiveDataFromRemotePlayer sets the handler for the MatchDidReceiveDataFromRemotePlayer delegate method.
//
// Processes the data sent from another player to the local player.
func (d *MatchDelegate) SetMatchDidReceiveDataFromRemotePlayer(f func(match IGKMatch, data objc.IObject /* cross-framework: NSData */, player IGKPlayer)) {
	d._MatchDidReceiveDataFromRemotePlayer = f
}

// SetMatchPlayerDidChangeState sets the handler for the MatchPlayerDidChangeState delegate method.
//
// Handles when a player connects or disconnects from a match.
func (d *MatchDelegate) SetMatchPlayerDidChangeState(f func(match IGKMatch, playerID objc.IObject /* cross-framework: NSString */, state PlayerConnectionState)) {
	d._MatchPlayerDidChangeState = f
}

// SetMatchPlayerDidChangeConnectionState sets the handler for the MatchPlayerDidChangeConnectionState delegate method.
//
// Handles when players connect or disconnect from a match.
func (d *MatchDelegate) SetMatchPlayerDidChangeConnectionState(f func(match IGKMatch, player IGKPlayer, state PlayerConnectionState)) {
	d._MatchPlayerDidChangeConnectionState = f
}

// SetMatchShouldReinviteDisconnectedPlayer sets the handler for the MatchShouldReinviteDisconnectedPlayer delegate method.
//
// Determines whether the local player should reinvite another player who disconnected from a two-player match.
func (d *MatchDelegate) SetMatchShouldReinviteDisconnectedPlayer(f func(match IGKMatch, player IGKPlayer) bool) {
	d._MatchShouldReinviteDisconnectedPlayer = f
}

// SetMatchShouldReinvitePlayer sets the handler for the MatchShouldReinvitePlayer delegate method.
//
// Handles when a player disconnects from a two-player match.
func (d *MatchDelegate) SetMatchShouldReinvitePlayer(f func(match IGKMatch, playerID objc.IObject /* cross-framework: NSString */) bool) {
	d._MatchShouldReinvitePlayer = f
}

// MatchDidFailWithError implements the PMatchDelegate interface.
func (d *MatchDelegate) MatchDidFailWithError(match IGKMatch, error_ objc.IObject /* cross-framework: Error */) {
	if d._MatchDidFailWithError != nil {
		d._MatchDidFailWithError(match, error_)
	}
}

// HasMatchDidFailWithError returns true if a handler for MatchDidFailWithError has been set.
func (d *MatchDelegate) HasMatchDidFailWithError() bool {
	return d._MatchDidFailWithError != nil
}

// MatchDidReceiveDataForRecipientFromRemotePlayer implements the PMatchDelegate interface.
func (d *MatchDelegate) MatchDidReceiveDataForRecipientFromRemotePlayer(match IGKMatch, data objc.IObject /* cross-framework: NSData */, recipient IGKPlayer, player IGKPlayer) {
	if d._MatchDidReceiveDataForRecipientFromRemotePlayer != nil {
		d._MatchDidReceiveDataForRecipientFromRemotePlayer(match, data, recipient, player)
	}
}

// HasMatchDidReceiveDataForRecipientFromRemotePlayer returns true if a handler for MatchDidReceiveDataForRecipientFromRemotePlayer has been set.
func (d *MatchDelegate) HasMatchDidReceiveDataForRecipientFromRemotePlayer() bool {
	return d._MatchDidReceiveDataForRecipientFromRemotePlayer != nil
}

// MatchDidReceiveDataFromPlayer implements the PMatchDelegate interface.
func (d *MatchDelegate) MatchDidReceiveDataFromPlayer(match IGKMatch, data objc.IObject /* cross-framework: NSData */, playerID objc.IObject /* cross-framework: NSString */) {
	if d._MatchDidReceiveDataFromPlayer != nil {
		d._MatchDidReceiveDataFromPlayer(match, data, playerID)
	}
}

// HasMatchDidReceiveDataFromPlayer returns true if a handler for MatchDidReceiveDataFromPlayer has been set.
func (d *MatchDelegate) HasMatchDidReceiveDataFromPlayer() bool {
	return d._MatchDidReceiveDataFromPlayer != nil
}

// MatchDidReceiveDataFromRemotePlayer implements the PMatchDelegate interface.
func (d *MatchDelegate) MatchDidReceiveDataFromRemotePlayer(match IGKMatch, data objc.IObject /* cross-framework: NSData */, player IGKPlayer) {
	if d._MatchDidReceiveDataFromRemotePlayer != nil {
		d._MatchDidReceiveDataFromRemotePlayer(match, data, player)
	}
}

// HasMatchDidReceiveDataFromRemotePlayer returns true if a handler for MatchDidReceiveDataFromRemotePlayer has been set.
func (d *MatchDelegate) HasMatchDidReceiveDataFromRemotePlayer() bool {
	return d._MatchDidReceiveDataFromRemotePlayer != nil
}

// MatchPlayerDidChangeState implements the PMatchDelegate interface.
func (d *MatchDelegate) MatchPlayerDidChangeState(match IGKMatch, playerID objc.IObject /* cross-framework: NSString */, state PlayerConnectionState) {
	if d._MatchPlayerDidChangeState != nil {
		d._MatchPlayerDidChangeState(match, playerID, state)
	}
}

// HasMatchPlayerDidChangeState returns true if a handler for MatchPlayerDidChangeState has been set.
func (d *MatchDelegate) HasMatchPlayerDidChangeState() bool {
	return d._MatchPlayerDidChangeState != nil
}

// MatchPlayerDidChangeConnectionState implements the PMatchDelegate interface.
func (d *MatchDelegate) MatchPlayerDidChangeConnectionState(match IGKMatch, player IGKPlayer, state PlayerConnectionState) {
	if d._MatchPlayerDidChangeConnectionState != nil {
		d._MatchPlayerDidChangeConnectionState(match, player, state)
	}
}

// HasMatchPlayerDidChangeConnectionState returns true if a handler for MatchPlayerDidChangeConnectionState has been set.
func (d *MatchDelegate) HasMatchPlayerDidChangeConnectionState() bool {
	return d._MatchPlayerDidChangeConnectionState != nil
}

// MatchShouldReinviteDisconnectedPlayer implements the PMatchDelegate interface.
func (d *MatchDelegate) MatchShouldReinviteDisconnectedPlayer(match IGKMatch, player IGKPlayer) bool {
	if d._MatchShouldReinviteDisconnectedPlayer != nil {
		return d._MatchShouldReinviteDisconnectedPlayer(match, player)
	}
	var zero bool
	return zero
}

// HasMatchShouldReinviteDisconnectedPlayer returns true if a handler for MatchShouldReinviteDisconnectedPlayer has been set.
func (d *MatchDelegate) HasMatchShouldReinviteDisconnectedPlayer() bool {
	return d._MatchShouldReinviteDisconnectedPlayer != nil
}

// MatchShouldReinvitePlayer implements the PMatchDelegate interface.
func (d *MatchDelegate) MatchShouldReinvitePlayer(match IGKMatch, playerID objc.IObject /* cross-framework: NSString */) bool {
	if d._MatchShouldReinvitePlayer != nil {
		return d._MatchShouldReinvitePlayer(match, playerID)
	}
	var zero bool
	return zero
}

// HasMatchShouldReinvitePlayer returns true if a handler for MatchShouldReinvitePlayer has been set.
func (d *MatchDelegate) HasMatchShouldReinvitePlayer() bool {
	return d._MatchShouldReinvitePlayer != nil
}
