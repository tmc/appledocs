// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PNowPlayingSessionDelegate is the MPNowPlayingSessionDelegate protocol interface.
//
// A protocol that defines the delegate interface for a Now Playing session.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.mediaplayer/documentation/MediaPlayer/MPNowPlayingSessionDelegate
type PNowPlayingSessionDelegate interface {
	// Optional methods
	NowPlayingSessionDidChangeActive(nowPlayingSession IMPNowPlayingSession)
	HasNowPlayingSessionDidChangeActive() bool
	NowPlayingSessionDidChangeCanBecomeActive(nowPlayingSession IMPNowPlayingSession)
	HasNowPlayingSessionDidChangeCanBecomeActive() bool
}

// NowPlayingSessionDelegate is a delegate implementation builder for the PNowPlayingSessionDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type NowPlayingSessionDelegate struct {
	_NowPlayingSessionDidChangeActive func(nowPlayingSession IMPNowPlayingSession)
	_NowPlayingSessionDidChangeCanBecomeActive func(nowPlayingSession IMPNowPlayingSession)
}

// SetNowPlayingSessionDidChangeActive sets the handler for the NowPlayingSessionDidChangeActive delegate method.
//
// Tells the delegate that the session changed its active status.
func (d *NowPlayingSessionDelegate) SetNowPlayingSessionDidChangeActive(f func(nowPlayingSession IMPNowPlayingSession)) {
	d._NowPlayingSessionDidChangeActive = f
}

// SetNowPlayingSessionDidChangeCanBecomeActive sets the handler for the NowPlayingSessionDidChangeCanBecomeActive delegate method.
//
// Tells the delegate that the session is eligible to become active.
func (d *NowPlayingSessionDelegate) SetNowPlayingSessionDidChangeCanBecomeActive(f func(nowPlayingSession IMPNowPlayingSession)) {
	d._NowPlayingSessionDidChangeCanBecomeActive = f
}

// NowPlayingSessionDidChangeActive implements the PNowPlayingSessionDelegate interface.
func (d *NowPlayingSessionDelegate) NowPlayingSessionDidChangeActive(nowPlayingSession IMPNowPlayingSession) {
	if d._NowPlayingSessionDidChangeActive != nil {
		d._NowPlayingSessionDidChangeActive(nowPlayingSession)
	}
}

// HasNowPlayingSessionDidChangeActive returns true if a handler for NowPlayingSessionDidChangeActive has been set.
func (d *NowPlayingSessionDelegate) HasNowPlayingSessionDidChangeActive() bool {
	return d._NowPlayingSessionDidChangeActive != nil
}

// NowPlayingSessionDidChangeCanBecomeActive implements the PNowPlayingSessionDelegate interface.
func (d *NowPlayingSessionDelegate) NowPlayingSessionDidChangeCanBecomeActive(nowPlayingSession IMPNowPlayingSession) {
	if d._NowPlayingSessionDidChangeCanBecomeActive != nil {
		d._NowPlayingSessionDidChangeCanBecomeActive(nowPlayingSession)
	}
}

// HasNowPlayingSessionDidChangeCanBecomeActive returns true if a handler for NowPlayingSessionDidChangeCanBecomeActive has been set.
func (d *NowPlayingSessionDelegate) HasNowPlayingSessionDidChangeCanBecomeActive() bool {
	return d._NowPlayingSessionDidChangeCanBecomeActive != nil
}
