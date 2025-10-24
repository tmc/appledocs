// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PTKTokenDelegate is the TKTokenDelegate protocol interface.
//
// The interface that a token delegate implements to respond to session creation events.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.12+
//   - tvOS 11.0+
//   - visionOS 1.0+
//   - watchOS 4.0+
//
// See: doc://com.apple.cryptotokenkit/documentation/CryptoTokenKit/TKTokenDelegate
type PTKTokenDelegate interface {
	// Required methods
	TokenCreateSessionWithError(token ITKToken, error_ unsafe.Pointer) TKTokenSession/* debug [protocol_interface/required_method]: TokenCreateSessionWithError */
	// Optional methods
	TokenTerminateSession(token ITKToken, session ITKTokenSession)
	HasTokenTerminateSession() bool
}

// TKTokenDelegate is a delegate implementation builder for the PTKTokenDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TKTokenDelegate struct {
	_TokenTerminateSession func(token ITKToken, session ITKTokenSession)
	_TokenCreateSessionWithError func(token ITKToken, error_ unsafe.Pointer) TKTokenSession
}

// SetTokenTerminateSession sets the handler for the TokenTerminateSession delegate method.
//
// Tells the delegate to terminate the specified token session.
func (d *TKTokenDelegate) SetTokenTerminateSession(f func(token ITKToken, session ITKTokenSession)) {
	d._TokenTerminateSession = f
}

// SetTokenCreateSessionWithError sets the handler for the TokenCreateSessionWithError delegate method.
//
// Tells the delegate to create a session for the specified token.
func (d *TKTokenDelegate) SetTokenCreateSessionWithError(f func(token ITKToken, error_ unsafe.Pointer) TKTokenSession) {
	d._TokenCreateSessionWithError = f
}

// TokenTerminateSession implements the PTKTokenDelegate interface.
func (d *TKTokenDelegate) TokenTerminateSession(token ITKToken, session ITKTokenSession) {
	if d._TokenTerminateSession != nil {
		d._TokenTerminateSession(token, session)
	}
}

// HasTokenTerminateSession returns true if a handler for TokenTerminateSession has been set.
func (d *TKTokenDelegate) HasTokenTerminateSession() bool {
	return d._TokenTerminateSession != nil
}

// TokenCreateSessionWithError implements the PTKTokenDelegate interface.
func (d *TKTokenDelegate) TokenCreateSessionWithError(token ITKToken, error_ unsafe.Pointer) TKTokenSession {
	if d._TokenCreateSessionWithError != nil {
		return d._TokenCreateSessionWithError(token, error_)
	}
	var zero TKTokenSession
	return zero
}

// HasTokenCreateSessionWithError returns true if a handler for TokenCreateSessionWithError has been set.
func (d *TKTokenDelegate) HasTokenCreateSessionWithError() bool {
	return d._TokenCreateSessionWithError != nil
}
