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
