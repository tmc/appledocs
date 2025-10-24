// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PTKSmartCardTokenDriverDelegate is the TKSmartCardTokenDriverDelegate protocol interface.
//
// The interface that a smart card token driver delegate implements to respond to token creation events.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//   - watchOS +
//
// See: doc://com.apple.cryptotokenkit/documentation/CryptoTokenKit/TKSmartCardTokenDriverDelegate
type PTKSmartCardTokenDriverDelegate interface {
	// Required methods
	TokenDriverCreateTokenForSmartCardAIDError(driver ITKSmartCardTokenDriver, smartCard ITKSmartCard, AID objc.IObject /* cross-framework: NSData */, error_ unsafe.Pointer) TKSmartCardToken/* debug [protocol_interface/required_method]: TokenDriverCreateTokenForSmartCardAIDError */
}

// TKSmartCardTokenDriverDelegate is a delegate implementation builder for the PTKSmartCardTokenDriverDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TKSmartCardTokenDriverDelegate struct {
	_TokenDriverCreateTokenForSmartCardAIDError func(driver ITKSmartCardTokenDriver, smartCard ITKSmartCard, AID objc.IObject /* cross-framework: NSData */, error_ unsafe.Pointer) TKSmartCardToken
}

// SetTokenDriverCreateTokenForSmartCardAIDError sets the handler for the TokenDriverCreateTokenForSmartCardAIDError delegate method.
//
// Tells the delegate that a new Smart Card is detected.
func (d *TKSmartCardTokenDriverDelegate) SetTokenDriverCreateTokenForSmartCardAIDError(f func(driver ITKSmartCardTokenDriver, smartCard ITKSmartCard, AID objc.IObject /* cross-framework: NSData */, error_ unsafe.Pointer) TKSmartCardToken) {
	d._TokenDriverCreateTokenForSmartCardAIDError = f
}

// TokenDriverCreateTokenForSmartCardAIDError implements the PTKSmartCardTokenDriverDelegate interface.
func (d *TKSmartCardTokenDriverDelegate) TokenDriverCreateTokenForSmartCardAIDError(driver ITKSmartCardTokenDriver, smartCard ITKSmartCard, AID objc.IObject /* cross-framework: NSData */, error_ unsafe.Pointer) TKSmartCardToken {
	if d._TokenDriverCreateTokenForSmartCardAIDError != nil {
		return d._TokenDriverCreateTokenForSmartCardAIDError(driver, smartCard, AID, error_)
	}
	var zero TKSmartCardToken
	return zero
}

// HasTokenDriverCreateTokenForSmartCardAIDError returns true if a handler for TokenDriverCreateTokenForSmartCardAIDError has been set.
func (d *TKSmartCardTokenDriverDelegate) HasTokenDriverCreateTokenForSmartCardAIDError() bool {
	return d._TokenDriverCreateTokenForSmartCardAIDError != nil
}
