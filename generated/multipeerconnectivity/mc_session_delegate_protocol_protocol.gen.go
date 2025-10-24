// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PMCSessionDelegate is the MCSessionDelegate protocol interface.
//
// The   protocol defines methods that a delegate of the   class can implement to handle session-related events. For more information, see  .
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.10+
//   - tvOS 10.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.multipeerconnectivity/documentation/MultipeerConnectivity/MCSessionDelegate
type PMCSessionDelegate interface {
	// Required methods
	SessionDidFinishReceivingResourceWithNameFromPeerAtURLWithError(session IMCSession, resourceName objc.IObject /* cross-framework: NSString */, peerID IMCPeerID, localURL objc.IObject /* cross-framework: NSURL */, error_ objc.IObject /* cross-framework: Error */)/* debug [protocol_interface/required_method]: SessionDidFinishReceivingResourceWithNameFromPeerAtURLWithError */
	SessionDidReceiveDataFromPeer(session IMCSession, data objc.IObject /* cross-framework: NSData */, peerID IMCPeerID)/* debug [protocol_interface/required_method]: SessionDidReceiveDataFromPeer */
	SessionDidReceiveStreamWithNameFromPeer(session IMCSession, stream foundation.InputStream, streamName objc.IObject /* cross-framework: NSString */, peerID IMCPeerID)/* debug [protocol_interface/required_method]: SessionDidReceiveStreamWithNameFromPeer */
	SessionDidStartReceivingResourceWithNameFromPeerWithProgress(session IMCSession, resourceName objc.IObject /* cross-framework: NSString */, peerID IMCPeerID, progress foundation.Progress)/* debug [protocol_interface/required_method]: SessionDidStartReceivingResourceWithNameFromPeerWithProgress */
	SessionPeerDidChangeState(session IMCSession, peerID IMCPeerID, state MCSessionState)/* debug [protocol_interface/required_method]: SessionPeerDidChangeState */
	// Optional methods
	SessionDidReceiveCertificateFromPeerCertificateHandler(session IMCSession, certificate objc.IObject /* cross-framework: NSArray */, peerID IMCPeerID, certificateHandler unsafe.Pointer)
	HasSessionDidReceiveCertificateFromPeerCertificateHandler() bool
}

// MCSessionDelegate is a delegate implementation builder for the PMCSessionDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type MCSessionDelegate struct {
	_SessionDidReceiveCertificateFromPeerCertificateHandler func(session IMCSession, certificate objc.IObject /* cross-framework: NSArray */, peerID IMCPeerID, certificateHandler unsafe.Pointer)
	_SessionDidFinishReceivingResourceWithNameFromPeerAtURLWithError func(session IMCSession, resourceName objc.IObject /* cross-framework: NSString */, peerID IMCPeerID, localURL objc.IObject /* cross-framework: NSURL */, error_ objc.IObject /* cross-framework: Error */)
	_SessionDidReceiveDataFromPeer func(session IMCSession, data objc.IObject /* cross-framework: NSData */, peerID IMCPeerID)
	_SessionDidReceiveStreamWithNameFromPeer func(session IMCSession, stream foundation.InputStream, streamName objc.IObject /* cross-framework: NSString */, peerID IMCPeerID)
	_SessionDidStartReceivingResourceWithNameFromPeerWithProgress func(session IMCSession, resourceName objc.IObject /* cross-framework: NSString */, peerID IMCPeerID, progress foundation.Progress)
	_SessionPeerDidChangeState func(session IMCSession, peerID IMCPeerID, state MCSessionState)
}

// SetSessionDidReceiveCertificateFromPeerCertificateHandler sets the handler for the SessionDidReceiveCertificateFromPeerCertificateHandler delegate method.
//
// Called to validate the client certificate provided by a peer when the connection is first established.
func (d *MCSessionDelegate) SetSessionDidReceiveCertificateFromPeerCertificateHandler(f func(session IMCSession, certificate objc.IObject /* cross-framework: NSArray */, peerID IMCPeerID, certificateHandler unsafe.Pointer)) {
	d._SessionDidReceiveCertificateFromPeerCertificateHandler = f
}

// SetSessionDidFinishReceivingResourceWithNameFromPeerAtURLWithError sets the handler for the SessionDidFinishReceivingResourceWithNameFromPeerAtURLWithError delegate method.
//
// Indicates that the local peer finished receiving a resource from a nearby peer.
func (d *MCSessionDelegate) SetSessionDidFinishReceivingResourceWithNameFromPeerAtURLWithError(f func(session IMCSession, resourceName objc.IObject /* cross-framework: NSString */, peerID IMCPeerID, localURL objc.IObject /* cross-framework: NSURL */, error_ objc.IObject /* cross-framework: Error */)) {
	d._SessionDidFinishReceivingResourceWithNameFromPeerAtURLWithError = f
}

// SetSessionDidReceiveDataFromPeer sets the handler for the SessionDidReceiveDataFromPeer delegate method.
//
// Indicates that an   object has been received from a nearby peer.
func (d *MCSessionDelegate) SetSessionDidReceiveDataFromPeer(f func(session IMCSession, data objc.IObject /* cross-framework: NSData */, peerID IMCPeerID)) {
	d._SessionDidReceiveDataFromPeer = f
}

// SetSessionDidReceiveStreamWithNameFromPeer sets the handler for the SessionDidReceiveStreamWithNameFromPeer delegate method.
//
// Called when a nearby peer opens a byte stream connection to the local peer.
func (d *MCSessionDelegate) SetSessionDidReceiveStreamWithNameFromPeer(f func(session IMCSession, stream foundation.InputStream, streamName objc.IObject /* cross-framework: NSString */, peerID IMCPeerID)) {
	d._SessionDidReceiveStreamWithNameFromPeer = f
}

// SetSessionDidStartReceivingResourceWithNameFromPeerWithProgress sets the handler for the SessionDidStartReceivingResourceWithNameFromPeerWithProgress delegate method.
//
// Indicates that the local peer began receiving a resource from a nearby peer.
func (d *MCSessionDelegate) SetSessionDidStartReceivingResourceWithNameFromPeerWithProgress(f func(session IMCSession, resourceName objc.IObject /* cross-framework: NSString */, peerID IMCPeerID, progress foundation.Progress)) {
	d._SessionDidStartReceivingResourceWithNameFromPeerWithProgress = f
}

// SetSessionPeerDidChangeState sets the handler for the SessionPeerDidChangeState delegate method.
//
// Called when the state of a nearby peer changes.
func (d *MCSessionDelegate) SetSessionPeerDidChangeState(f func(session IMCSession, peerID IMCPeerID, state MCSessionState)) {
	d._SessionPeerDidChangeState = f
}

// SessionDidReceiveCertificateFromPeerCertificateHandler implements the PMCSessionDelegate interface.
func (d *MCSessionDelegate) SessionDidReceiveCertificateFromPeerCertificateHandler(session IMCSession, certificate objc.IObject /* cross-framework: NSArray */, peerID IMCPeerID, certificateHandler unsafe.Pointer) {
	if d._SessionDidReceiveCertificateFromPeerCertificateHandler != nil {
		d._SessionDidReceiveCertificateFromPeerCertificateHandler(session, certificate, peerID, certificateHandler)
	}
}

// HasSessionDidReceiveCertificateFromPeerCertificateHandler returns true if a handler for SessionDidReceiveCertificateFromPeerCertificateHandler has been set.
func (d *MCSessionDelegate) HasSessionDidReceiveCertificateFromPeerCertificateHandler() bool {
	return d._SessionDidReceiveCertificateFromPeerCertificateHandler != nil
}

// SessionDidFinishReceivingResourceWithNameFromPeerAtURLWithError implements the PMCSessionDelegate interface.
func (d *MCSessionDelegate) SessionDidFinishReceivingResourceWithNameFromPeerAtURLWithError(session IMCSession, resourceName objc.IObject /* cross-framework: NSString */, peerID IMCPeerID, localURL objc.IObject /* cross-framework: NSURL */, error_ objc.IObject /* cross-framework: Error */) {
	if d._SessionDidFinishReceivingResourceWithNameFromPeerAtURLWithError != nil {
		d._SessionDidFinishReceivingResourceWithNameFromPeerAtURLWithError(session, resourceName, peerID, localURL, error_)
	}
}

// HasSessionDidFinishReceivingResourceWithNameFromPeerAtURLWithError returns true if a handler for SessionDidFinishReceivingResourceWithNameFromPeerAtURLWithError has been set.
func (d *MCSessionDelegate) HasSessionDidFinishReceivingResourceWithNameFromPeerAtURLWithError() bool {
	return d._SessionDidFinishReceivingResourceWithNameFromPeerAtURLWithError != nil
}

// SessionDidReceiveDataFromPeer implements the PMCSessionDelegate interface.
func (d *MCSessionDelegate) SessionDidReceiveDataFromPeer(session IMCSession, data objc.IObject /* cross-framework: NSData */, peerID IMCPeerID) {
	if d._SessionDidReceiveDataFromPeer != nil {
		d._SessionDidReceiveDataFromPeer(session, data, peerID)
	}
}

// HasSessionDidReceiveDataFromPeer returns true if a handler for SessionDidReceiveDataFromPeer has been set.
func (d *MCSessionDelegate) HasSessionDidReceiveDataFromPeer() bool {
	return d._SessionDidReceiveDataFromPeer != nil
}

// SessionDidReceiveStreamWithNameFromPeer implements the PMCSessionDelegate interface.
func (d *MCSessionDelegate) SessionDidReceiveStreamWithNameFromPeer(session IMCSession, stream foundation.InputStream, streamName objc.IObject /* cross-framework: NSString */, peerID IMCPeerID) {
	if d._SessionDidReceiveStreamWithNameFromPeer != nil {
		d._SessionDidReceiveStreamWithNameFromPeer(session, stream, streamName, peerID)
	}
}

// HasSessionDidReceiveStreamWithNameFromPeer returns true if a handler for SessionDidReceiveStreamWithNameFromPeer has been set.
func (d *MCSessionDelegate) HasSessionDidReceiveStreamWithNameFromPeer() bool {
	return d._SessionDidReceiveStreamWithNameFromPeer != nil
}

// SessionDidStartReceivingResourceWithNameFromPeerWithProgress implements the PMCSessionDelegate interface.
func (d *MCSessionDelegate) SessionDidStartReceivingResourceWithNameFromPeerWithProgress(session IMCSession, resourceName objc.IObject /* cross-framework: NSString */, peerID IMCPeerID, progress foundation.Progress) {
	if d._SessionDidStartReceivingResourceWithNameFromPeerWithProgress != nil {
		d._SessionDidStartReceivingResourceWithNameFromPeerWithProgress(session, resourceName, peerID, progress)
	}
}

// HasSessionDidStartReceivingResourceWithNameFromPeerWithProgress returns true if a handler for SessionDidStartReceivingResourceWithNameFromPeerWithProgress has been set.
func (d *MCSessionDelegate) HasSessionDidStartReceivingResourceWithNameFromPeerWithProgress() bool {
	return d._SessionDidStartReceivingResourceWithNameFromPeerWithProgress != nil
}

// SessionPeerDidChangeState implements the PMCSessionDelegate interface.
func (d *MCSessionDelegate) SessionPeerDidChangeState(session IMCSession, peerID IMCPeerID, state MCSessionState) {
	if d._SessionPeerDidChangeState != nil {
		d._SessionPeerDidChangeState(session, peerID, state)
	}
}

// HasSessionPeerDidChangeState returns true if a handler for SessionPeerDidChangeState has been set.
func (d *MCSessionDelegate) HasSessionPeerDidChangeState() bool {
	return d._SessionPeerDidChangeState != nil
}
