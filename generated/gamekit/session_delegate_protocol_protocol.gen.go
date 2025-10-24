// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PSessionDelegate is the GKSessionDelegate protocol interface.
//
// An object implements the   protocol to control the behavior of a   object. The delegate is called when other visible peers change their state relative to the session. It is also called to determine whether another peer is allowed to connect to the session.
//
// Availability:
//   - macOS 10.8+ (Deprecated in 10.10)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// See: doc://com.apple.gamekit/documentation/GameKit/GKSessionDelegate
type PSessionDelegate interface {
	// Optional methods
	SessionConnectionWithPeerFailedWithError(session IGKSession, peerID objc.IObject /* cross-framework: NSString */, error_ objc.IObject /* cross-framework: Error */)
	HasSessionConnectionWithPeerFailedWithError() bool
	SessionDidFailWithError(session IGKSession, error_ objc.IObject /* cross-framework: Error */)
	HasSessionDidFailWithError() bool
	SessionDidReceiveConnectionRequestFromPeer(session IGKSession, peerID objc.IObject /* cross-framework: NSString */)
	HasSessionDidReceiveConnectionRequestFromPeer() bool
	SessionPeerDidChangeState(session IGKSession, peerID objc.IObject /* cross-framework: NSString */, state PeerConnectionState)
	HasSessionPeerDidChangeState() bool
}

// SessionDelegate is a delegate implementation builder for the PSessionDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type SessionDelegate struct {
	_SessionConnectionWithPeerFailedWithError func(session IGKSession, peerID objc.IObject /* cross-framework: NSString */, error_ objc.IObject /* cross-framework: Error */)
	_SessionDidFailWithError func(session IGKSession, error_ objc.IObject /* cross-framework: Error */)
	_SessionDidReceiveConnectionRequestFromPeer func(session IGKSession, peerID objc.IObject /* cross-framework: NSString */)
	_SessionPeerDidChangeState func(session IGKSession, peerID objc.IObject /* cross-framework: NSString */, state PeerConnectionState)
}

// SetSessionConnectionWithPeerFailedWithError sets the handler for the SessionConnectionWithPeerFailedWithError delegate method.
//
// Received by the delegate when an attempt to connect to another peer failed.
func (d *SessionDelegate) SetSessionConnectionWithPeerFailedWithError(f func(session IGKSession, peerID objc.IObject /* cross-framework: NSString */, error_ objc.IObject /* cross-framework: Error */)) {
	d._SessionConnectionWithPeerFailedWithError = f
}

// SetSessionDidFailWithError sets the handler for the SessionDidFailWithError delegate method.
//
// Sent to the delegate when a serious error has occurred in the session.
func (d *SessionDelegate) SetSessionDidFailWithError(f func(session IGKSession, error_ objc.IObject /* cross-framework: Error */)) {
	d._SessionDidFailWithError = f
}

// SetSessionDidReceiveConnectionRequestFromPeer sets the handler for the SessionDidReceiveConnectionRequestFromPeer delegate method.
//
// Received by the delegate when a remote peer wants to create a connection to the session.
func (d *SessionDelegate) SetSessionDidReceiveConnectionRequestFromPeer(f func(session IGKSession, peerID objc.IObject /* cross-framework: NSString */)) {
	d._SessionDidReceiveConnectionRequestFromPeer = f
}

// SetSessionPeerDidChangeState sets the handler for the SessionPeerDidChangeState delegate method.
//
// Received by the delegate when a peer changes state.
func (d *SessionDelegate) SetSessionPeerDidChangeState(f func(session IGKSession, peerID objc.IObject /* cross-framework: NSString */, state PeerConnectionState)) {
	d._SessionPeerDidChangeState = f
}

// SessionConnectionWithPeerFailedWithError implements the PSessionDelegate interface.
func (d *SessionDelegate) SessionConnectionWithPeerFailedWithError(session IGKSession, peerID objc.IObject /* cross-framework: NSString */, error_ objc.IObject /* cross-framework: Error */) {
	if d._SessionConnectionWithPeerFailedWithError != nil {
		d._SessionConnectionWithPeerFailedWithError(session, peerID, error_)
	}
}

// HasSessionConnectionWithPeerFailedWithError returns true if a handler for SessionConnectionWithPeerFailedWithError has been set.
func (d *SessionDelegate) HasSessionConnectionWithPeerFailedWithError() bool {
	return d._SessionConnectionWithPeerFailedWithError != nil
}

// SessionDidFailWithError implements the PSessionDelegate interface.
func (d *SessionDelegate) SessionDidFailWithError(session IGKSession, error_ objc.IObject /* cross-framework: Error */) {
	if d._SessionDidFailWithError != nil {
		d._SessionDidFailWithError(session, error_)
	}
}

// HasSessionDidFailWithError returns true if a handler for SessionDidFailWithError has been set.
func (d *SessionDelegate) HasSessionDidFailWithError() bool {
	return d._SessionDidFailWithError != nil
}

// SessionDidReceiveConnectionRequestFromPeer implements the PSessionDelegate interface.
func (d *SessionDelegate) SessionDidReceiveConnectionRequestFromPeer(session IGKSession, peerID objc.IObject /* cross-framework: NSString */) {
	if d._SessionDidReceiveConnectionRequestFromPeer != nil {
		d._SessionDidReceiveConnectionRequestFromPeer(session, peerID)
	}
}

// HasSessionDidReceiveConnectionRequestFromPeer returns true if a handler for SessionDidReceiveConnectionRequestFromPeer has been set.
func (d *SessionDelegate) HasSessionDidReceiveConnectionRequestFromPeer() bool {
	return d._SessionDidReceiveConnectionRequestFromPeer != nil
}

// SessionPeerDidChangeState implements the PSessionDelegate interface.
func (d *SessionDelegate) SessionPeerDidChangeState(session IGKSession, peerID objc.IObject /* cross-framework: NSString */, state PeerConnectionState) {
	if d._SessionPeerDidChangeState != nil {
		d._SessionPeerDidChangeState(session, peerID, state)
	}
}

// HasSessionPeerDidChangeState returns true if a handler for SessionPeerDidChangeState has been set.
func (d *SessionDelegate) HasSessionPeerDidChangeState() bool {
	return d._SessionPeerDidChangeState != nil
}
