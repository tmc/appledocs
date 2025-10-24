// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PChallengeEventHandlerDelegate is the GKChallengeEventHandlerDelegate protocol interface.
//
// You implement the   delegate to control how challenges are displayed in your game.
//
// Availability:
//   - macOS 10.8+ (Deprecated in 10.10)
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// See: doc://com.apple.gamekit/documentation/GameKit/GKChallengeEventHandlerDelegate
type PChallengeEventHandlerDelegate interface {
	// Optional methods
	LocalPlayerDidCompleteChallenge(challenge IGKChallenge)
	HasLocalPlayerDidCompleteChallenge() bool
	LocalPlayerDidReceiveChallenge(challenge IGKChallenge)
	HasLocalPlayerDidReceiveChallenge() bool
	LocalPlayerDidSelectChallenge(challenge IGKChallenge)
	HasLocalPlayerDidSelectChallenge() bool
	RemotePlayerDidCompleteChallenge(challenge IGKChallenge)
	HasRemotePlayerDidCompleteChallenge() bool
	ShouldShowBannerForLocallyCompletedChallenge(challenge IGKChallenge) bool
	HasShouldShowBannerForLocallyCompletedChallenge() bool
	ShouldShowBannerForLocallyReceivedChallenge(challenge IGKChallenge) bool
	HasShouldShowBannerForLocallyReceivedChallenge() bool
	ShouldShowBannerForRemotelyCompletedChallenge(challenge IGKChallenge) bool
	HasShouldShowBannerForRemotelyCompletedChallenge() bool
}

// ChallengeEventHandlerDelegate is a delegate implementation builder for the PChallengeEventHandlerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ChallengeEventHandlerDelegate struct {
	_LocalPlayerDidCompleteChallenge func(challenge IGKChallenge)
	_LocalPlayerDidReceiveChallenge func(challenge IGKChallenge)
	_LocalPlayerDidSelectChallenge func(challenge IGKChallenge)
	_RemotePlayerDidCompleteChallenge func(challenge IGKChallenge)
	_ShouldShowBannerForLocallyCompletedChallenge func(challenge IGKChallenge) bool
	_ShouldShowBannerForLocallyReceivedChallenge func(challenge IGKChallenge) bool
	_ShouldShowBannerForRemotelyCompletedChallenge func(challenge IGKChallenge) bool
}

// SetLocalPlayerDidCompleteChallenge sets the handler for the LocalPlayerDidCompleteChallenge delegate method.
//
// Called when the local player completes a challenge.
func (d *ChallengeEventHandlerDelegate) SetLocalPlayerDidCompleteChallenge(f func(challenge IGKChallenge)) {
	d._LocalPlayerDidCompleteChallenge = f
}

// SetLocalPlayerDidReceiveChallenge sets the handler for the LocalPlayerDidReceiveChallenge delegate method.
//
// Called when the local player receives a new challenge.
func (d *ChallengeEventHandlerDelegate) SetLocalPlayerDidReceiveChallenge(f func(challenge IGKChallenge)) {
	d._LocalPlayerDidReceiveChallenge = f
}

// SetLocalPlayerDidSelectChallenge sets the handler for the LocalPlayerDidSelectChallenge delegate method.
//
// Called when the local player selects a challenge banner displayed by GameKit.
func (d *ChallengeEventHandlerDelegate) SetLocalPlayerDidSelectChallenge(f func(challenge IGKChallenge)) {
	d._LocalPlayerDidSelectChallenge = f
}

// SetRemotePlayerDidCompleteChallenge sets the handler for the RemotePlayerDidCompleteChallenge delegate method.
//
// Called when a remote player completes a challenge issued by the local player.
func (d *ChallengeEventHandlerDelegate) SetRemotePlayerDidCompleteChallenge(f func(challenge IGKChallenge)) {
	d._RemotePlayerDidCompleteChallenge = f
}

// SetShouldShowBannerForLocallyCompletedChallenge sets the handler for the ShouldShowBannerForLocallyCompletedChallenge delegate method.
//
// Called to determine whether a banner should be shown when the local player completes a challenge.
func (d *ChallengeEventHandlerDelegate) SetShouldShowBannerForLocallyCompletedChallenge(f func(challenge IGKChallenge) bool) {
	d._ShouldShowBannerForLocallyCompletedChallenge = f
}

// SetShouldShowBannerForLocallyReceivedChallenge sets the handler for the ShouldShowBannerForLocallyReceivedChallenge delegate method.
//
// Called to determine whether a banner should be shown when the local player receives a challenge.
func (d *ChallengeEventHandlerDelegate) SetShouldShowBannerForLocallyReceivedChallenge(f func(challenge IGKChallenge) bool) {
	d._ShouldShowBannerForLocallyReceivedChallenge = f
}

// SetShouldShowBannerForRemotelyCompletedChallenge sets the handler for the ShouldShowBannerForRemotelyCompletedChallenge delegate method.
//
// Called to determine whether a banner should be shown when a remote player completes a challenge.
func (d *ChallengeEventHandlerDelegate) SetShouldShowBannerForRemotelyCompletedChallenge(f func(challenge IGKChallenge) bool) {
	d._ShouldShowBannerForRemotelyCompletedChallenge = f
}

// LocalPlayerDidCompleteChallenge implements the PChallengeEventHandlerDelegate interface.
func (d *ChallengeEventHandlerDelegate) LocalPlayerDidCompleteChallenge(challenge IGKChallenge) {
	if d._LocalPlayerDidCompleteChallenge != nil {
		d._LocalPlayerDidCompleteChallenge(challenge)
	}
}

// HasLocalPlayerDidCompleteChallenge returns true if a handler for LocalPlayerDidCompleteChallenge has been set.
func (d *ChallengeEventHandlerDelegate) HasLocalPlayerDidCompleteChallenge() bool {
	return d._LocalPlayerDidCompleteChallenge != nil
}

// LocalPlayerDidReceiveChallenge implements the PChallengeEventHandlerDelegate interface.
func (d *ChallengeEventHandlerDelegate) LocalPlayerDidReceiveChallenge(challenge IGKChallenge) {
	if d._LocalPlayerDidReceiveChallenge != nil {
		d._LocalPlayerDidReceiveChallenge(challenge)
	}
}

// HasLocalPlayerDidReceiveChallenge returns true if a handler for LocalPlayerDidReceiveChallenge has been set.
func (d *ChallengeEventHandlerDelegate) HasLocalPlayerDidReceiveChallenge() bool {
	return d._LocalPlayerDidReceiveChallenge != nil
}

// LocalPlayerDidSelectChallenge implements the PChallengeEventHandlerDelegate interface.
func (d *ChallengeEventHandlerDelegate) LocalPlayerDidSelectChallenge(challenge IGKChallenge) {
	if d._LocalPlayerDidSelectChallenge != nil {
		d._LocalPlayerDidSelectChallenge(challenge)
	}
}

// HasLocalPlayerDidSelectChallenge returns true if a handler for LocalPlayerDidSelectChallenge has been set.
func (d *ChallengeEventHandlerDelegate) HasLocalPlayerDidSelectChallenge() bool {
	return d._LocalPlayerDidSelectChallenge != nil
}

// RemotePlayerDidCompleteChallenge implements the PChallengeEventHandlerDelegate interface.
func (d *ChallengeEventHandlerDelegate) RemotePlayerDidCompleteChallenge(challenge IGKChallenge) {
	if d._RemotePlayerDidCompleteChallenge != nil {
		d._RemotePlayerDidCompleteChallenge(challenge)
	}
}

// HasRemotePlayerDidCompleteChallenge returns true if a handler for RemotePlayerDidCompleteChallenge has been set.
func (d *ChallengeEventHandlerDelegate) HasRemotePlayerDidCompleteChallenge() bool {
	return d._RemotePlayerDidCompleteChallenge != nil
}

// ShouldShowBannerForLocallyCompletedChallenge implements the PChallengeEventHandlerDelegate interface.
func (d *ChallengeEventHandlerDelegate) ShouldShowBannerForLocallyCompletedChallenge(challenge IGKChallenge) bool {
	if d._ShouldShowBannerForLocallyCompletedChallenge != nil {
		return d._ShouldShowBannerForLocallyCompletedChallenge(challenge)
	}
	var zero bool
	return zero
}

// HasShouldShowBannerForLocallyCompletedChallenge returns true if a handler for ShouldShowBannerForLocallyCompletedChallenge has been set.
func (d *ChallengeEventHandlerDelegate) HasShouldShowBannerForLocallyCompletedChallenge() bool {
	return d._ShouldShowBannerForLocallyCompletedChallenge != nil
}

// ShouldShowBannerForLocallyReceivedChallenge implements the PChallengeEventHandlerDelegate interface.
func (d *ChallengeEventHandlerDelegate) ShouldShowBannerForLocallyReceivedChallenge(challenge IGKChallenge) bool {
	if d._ShouldShowBannerForLocallyReceivedChallenge != nil {
		return d._ShouldShowBannerForLocallyReceivedChallenge(challenge)
	}
	var zero bool
	return zero
}

// HasShouldShowBannerForLocallyReceivedChallenge returns true if a handler for ShouldShowBannerForLocallyReceivedChallenge has been set.
func (d *ChallengeEventHandlerDelegate) HasShouldShowBannerForLocallyReceivedChallenge() bool {
	return d._ShouldShowBannerForLocallyReceivedChallenge != nil
}

// ShouldShowBannerForRemotelyCompletedChallenge implements the PChallengeEventHandlerDelegate interface.
func (d *ChallengeEventHandlerDelegate) ShouldShowBannerForRemotelyCompletedChallenge(challenge IGKChallenge) bool {
	if d._ShouldShowBannerForRemotelyCompletedChallenge != nil {
		return d._ShouldShowBannerForRemotelyCompletedChallenge(challenge)
	}
	var zero bool
	return zero
}

// HasShouldShowBannerForRemotelyCompletedChallenge returns true if a handler for ShouldShowBannerForRemotelyCompletedChallenge has been set.
func (d *ChallengeEventHandlerDelegate) HasShouldShowBannerForRemotelyCompletedChallenge() bool {
	return d._ShouldShowBannerForRemotelyCompletedChallenge != nil
}
