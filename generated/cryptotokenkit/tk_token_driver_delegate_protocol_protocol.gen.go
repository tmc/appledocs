// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PTKTokenDriverDelegate is the TKTokenDriverDelegate protocol interface.
//
// The interface that a token driver delegate implements to respond to token creation events.
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
// See: doc://com.apple.cryptotokenkit/documentation/CryptoTokenKit/TKTokenDriverDelegate
type PTKTokenDriverDelegate interface {
	// Optional methods
	TokenDriverTerminateToken(driver ITKTokenDriver, token ITKToken)
	HasTokenDriverTerminateToken() bool
	TokenDriverTokenForConfigurationError(driver ITKTokenDriver, configuration ITKTokenConfiguration, error_ unsafe.Pointer) TKToken
	HasTokenDriverTokenForConfigurationError() bool
}

// TKTokenDriverDelegate is a delegate implementation builder for the PTKTokenDriverDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TKTokenDriverDelegate struct {
	_TokenDriverTerminateToken func(driver ITKTokenDriver, token ITKToken)
	_TokenDriverTokenForConfigurationError func(driver ITKTokenDriver, configuration ITKTokenConfiguration, error_ unsafe.Pointer) TKToken
}

// SetTokenDriverTerminateToken sets the handler for the TokenDriverTerminateToken delegate method.
//
// Tells the delegate to terminate the token you specify.
func (d *TKTokenDriverDelegate) SetTokenDriverTerminateToken(f func(driver ITKTokenDriver, token ITKToken)) {
	d._TokenDriverTerminateToken = f
}

// SetTokenDriverTokenForConfigurationError sets the handler for the TokenDriverTokenForConfigurationError delegate method.
//
// Creates a new token for the configuration you specify.
func (d *TKTokenDriverDelegate) SetTokenDriverTokenForConfigurationError(f func(driver ITKTokenDriver, configuration ITKTokenConfiguration, error_ unsafe.Pointer) TKToken) {
	d._TokenDriverTokenForConfigurationError = f
}

// TokenDriverTerminateToken implements the PTKTokenDriverDelegate interface.
func (d *TKTokenDriverDelegate) TokenDriverTerminateToken(driver ITKTokenDriver, token ITKToken) {
	if d._TokenDriverTerminateToken != nil {
		d._TokenDriverTerminateToken(driver, token)
	}
}

// HasTokenDriverTerminateToken returns true if a handler for TokenDriverTerminateToken has been set.
func (d *TKTokenDriverDelegate) HasTokenDriverTerminateToken() bool {
	return d._TokenDriverTerminateToken != nil
}

// TokenDriverTokenForConfigurationError implements the PTKTokenDriverDelegate interface.
func (d *TKTokenDriverDelegate) TokenDriverTokenForConfigurationError(driver ITKTokenDriver, configuration ITKTokenConfiguration, error_ unsafe.Pointer) TKToken {
	if d._TokenDriverTokenForConfigurationError != nil {
		return d._TokenDriverTokenForConfigurationError(driver, configuration, error_)
	}
	var zero TKToken
	return zero
}

// HasTokenDriverTokenForConfigurationError returns true if a handler for TokenDriverTokenForConfigurationError has been set.
func (d *TKTokenDriverDelegate) HasTokenDriverTokenForConfigurationError() bool {
	return d._TokenDriverTokenForConfigurationError != nil
}
