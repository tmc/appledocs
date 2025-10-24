// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PNWTCPConnectionAuthenticationDelegate is the NWTCPConnectionAuthenticationDelegate protocol interface.
//
// A delegate protocol to customize the TLS authentication done by a connection.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 18.0)
//   - iOS 9.0+ (Deprecated in 18.0)
//   - iPadOS 9.0+ (Deprecated in 18.0)
//   - macOS 10.11+ (Deprecated in 15.0)
//   - tvOS 17.0+ (Deprecated in 18.0)
//   - visionOS 1.0+ (Deprecated in 2.0)
//
// See: doc://com.apple.networkextension/documentation/NetworkExtension/NWTCPConnectionAuthenticationDelegate
type PNWTCPConnectionAuthenticationDelegate interface {
	// Optional methods
	EvaluateTrustForConnectionPeerCertificateChainCompletionHandler(connection INWTCPConnection, peerCertificateChain []objc.ID, completion unsafe.Pointer)
	HasEvaluateTrustForConnectionPeerCertificateChainCompletionHandler() bool
	ProvideIdentityForConnectionCompletionHandler(connection INWTCPConnection, completion unsafe.Pointer)
	HasProvideIdentityForConnectionCompletionHandler() bool
	ShouldEvaluateTrustForConnection(connection INWTCPConnection) bool
	HasShouldEvaluateTrustForConnection() bool
	ShouldProvideIdentityForConnection(connection INWTCPConnection) bool
	HasShouldProvideIdentityForConnection() bool
}

// NWTCPConnectionAuthenticationDelegate is a delegate implementation builder for the PNWTCPConnectionAuthenticationDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type NWTCPConnectionAuthenticationDelegate struct {
	_EvaluateTrustForConnectionPeerCertificateChainCompletionHandler func(connection INWTCPConnection, peerCertificateChain []objc.ID, completion unsafe.Pointer)
	_ProvideIdentityForConnectionCompletionHandler func(connection INWTCPConnection, completion unsafe.Pointer)
	_ShouldEvaluateTrustForConnection func(connection INWTCPConnection) bool
	_ShouldProvideIdentityForConnection func(connection INWTCPConnection) bool
}

// SetEvaluateTrustForConnectionPeerCertificateChainCompletionHandler sets the handler for the EvaluateTrustForConnectionPeerCertificateChainCompletionHandler delegate method.
//
// Override the default trust evaluation for the connection.
func (d *NWTCPConnectionAuthenticationDelegate) SetEvaluateTrustForConnectionPeerCertificateChainCompletionHandler(f func(connection INWTCPConnection, peerCertificateChain []objc.ID, completion unsafe.Pointer)) {
	d._EvaluateTrustForConnectionPeerCertificateChainCompletionHandler = f
}

// SetProvideIdentityForConnectionCompletionHandler sets the handler for the ProvideIdentityForConnectionCompletionHandler delegate method.
//
// Provide the identity and an optional certificate chain to be used for authentication.
func (d *NWTCPConnectionAuthenticationDelegate) SetProvideIdentityForConnectionCompletionHandler(f func(connection INWTCPConnection, completion unsafe.Pointer)) {
	d._ProvideIdentityForConnectionCompletionHandler = f
}

// SetShouldEvaluateTrustForConnection sets the handler for the ShouldEvaluateTrustForConnection delegate method.
//
// Indicate that the delegate should override the default trust evaluation for the connection.
func (d *NWTCPConnectionAuthenticationDelegate) SetShouldEvaluateTrustForConnection(f func(connection INWTCPConnection) bool) {
	d._ShouldEvaluateTrustForConnection = f
}

// SetShouldProvideIdentityForConnection sets the handler for the ShouldProvideIdentityForConnection delegate method.
//
// Indicate that the delegate can provide an identity for the connection authentication.
func (d *NWTCPConnectionAuthenticationDelegate) SetShouldProvideIdentityForConnection(f func(connection INWTCPConnection) bool) {
	d._ShouldProvideIdentityForConnection = f
}

// EvaluateTrustForConnectionPeerCertificateChainCompletionHandler implements the PNWTCPConnectionAuthenticationDelegate interface.
func (d *NWTCPConnectionAuthenticationDelegate) EvaluateTrustForConnectionPeerCertificateChainCompletionHandler(connection INWTCPConnection, peerCertificateChain []objc.ID, completion unsafe.Pointer) {
	if d._EvaluateTrustForConnectionPeerCertificateChainCompletionHandler != nil {
		d._EvaluateTrustForConnectionPeerCertificateChainCompletionHandler(connection, peerCertificateChain, completion)
	}
}

// HasEvaluateTrustForConnectionPeerCertificateChainCompletionHandler returns true if a handler for EvaluateTrustForConnectionPeerCertificateChainCompletionHandler has been set.
func (d *NWTCPConnectionAuthenticationDelegate) HasEvaluateTrustForConnectionPeerCertificateChainCompletionHandler() bool {
	return d._EvaluateTrustForConnectionPeerCertificateChainCompletionHandler != nil
}

// ProvideIdentityForConnectionCompletionHandler implements the PNWTCPConnectionAuthenticationDelegate interface.
func (d *NWTCPConnectionAuthenticationDelegate) ProvideIdentityForConnectionCompletionHandler(connection INWTCPConnection, completion unsafe.Pointer) {
	if d._ProvideIdentityForConnectionCompletionHandler != nil {
		d._ProvideIdentityForConnectionCompletionHandler(connection, completion)
	}
}

// HasProvideIdentityForConnectionCompletionHandler returns true if a handler for ProvideIdentityForConnectionCompletionHandler has been set.
func (d *NWTCPConnectionAuthenticationDelegate) HasProvideIdentityForConnectionCompletionHandler() bool {
	return d._ProvideIdentityForConnectionCompletionHandler != nil
}

// ShouldEvaluateTrustForConnection implements the PNWTCPConnectionAuthenticationDelegate interface.
func (d *NWTCPConnectionAuthenticationDelegate) ShouldEvaluateTrustForConnection(connection INWTCPConnection) bool {
	if d._ShouldEvaluateTrustForConnection != nil {
		return d._ShouldEvaluateTrustForConnection(connection)
	}
	var zero bool
	return zero
}

// HasShouldEvaluateTrustForConnection returns true if a handler for ShouldEvaluateTrustForConnection has been set.
func (d *NWTCPConnectionAuthenticationDelegate) HasShouldEvaluateTrustForConnection() bool {
	return d._ShouldEvaluateTrustForConnection != nil
}

// ShouldProvideIdentityForConnection implements the PNWTCPConnectionAuthenticationDelegate interface.
func (d *NWTCPConnectionAuthenticationDelegate) ShouldProvideIdentityForConnection(connection INWTCPConnection) bool {
	if d._ShouldProvideIdentityForConnection != nil {
		return d._ShouldProvideIdentityForConnection(connection)
	}
	var zero bool
	return zero
}

// HasShouldProvideIdentityForConnection returns true if a handler for ShouldProvideIdentityForConnection has been set.
func (d *NWTCPConnectionAuthenticationDelegate) HasShouldProvideIdentityForConnection() bool {
	return d._ShouldProvideIdentityForConnection != nil
}
