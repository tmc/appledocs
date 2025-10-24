// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PTKTokenSessionDelegate is the TKTokenSessionDelegate protocol interface.
//
// The interface that a session instance delegate implements to respond to token session authentication events.
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
// See: doc://com.apple.cryptotokenkit/documentation/CryptoTokenKit/TKTokenSessionDelegate
type PTKTokenSessionDelegate interface {
	// Optional methods
	TokenSessionBeginAuthForOperationConstraintError(session ITKTokenSession, operation TKTokenOperation, constraint TKTokenOperationConstraint /* typedef */, error_ unsafe.Pointer) TKTokenAuthOperation
	HasTokenSessionBeginAuthForOperationConstraintError() bool
	TokenSessionDecryptDataUsingKeyAlgorithmError(session ITKTokenSession, ciphertext objc.IObject /* cross-framework: NSData */, keyObjectID TKTokenObjectID /* typedef */, algorithm ITKTokenKeyAlgorithm, error_ unsafe.Pointer) foundation.Data
	HasTokenSessionDecryptDataUsingKeyAlgorithmError() bool
	TokenSessionPerformKeyExchangeWithPublicKeyUsingKeyAlgorithmParametersError(session ITKTokenSession, otherPartyPublicKeyData objc.IObject /* cross-framework: NSData */, objectID TKTokenObjectID /* typedef */, algorithm ITKTokenKeyAlgorithm, parameters ITKTokenKeyExchangeParameters, error_ unsafe.Pointer) foundation.Data
	HasTokenSessionPerformKeyExchangeWithPublicKeyUsingKeyAlgorithmParametersError() bool
	TokenSessionSignDataUsingKeyAlgorithmError(session ITKTokenSession, dataToSign objc.IObject /* cross-framework: NSData */, keyObjectID TKTokenObjectID /* typedef */, algorithm ITKTokenKeyAlgorithm, error_ unsafe.Pointer) foundation.Data
	HasTokenSessionSignDataUsingKeyAlgorithmError() bool
	TokenSessionSupportsOperationUsingKeyAlgorithm(session ITKTokenSession, operation TKTokenOperation, keyObjectID TKTokenObjectID /* typedef */, algorithm ITKTokenKeyAlgorithm) bool
	HasTokenSessionSupportsOperationUsingKeyAlgorithm() bool
}

// TKTokenSessionDelegate is a delegate implementation builder for the PTKTokenSessionDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type TKTokenSessionDelegate struct {
	_TokenSessionBeginAuthForOperationConstraintError func(session ITKTokenSession, operation TKTokenOperation, constraint TKTokenOperationConstraint /* typedef */, error_ unsafe.Pointer) TKTokenAuthOperation
	_TokenSessionDecryptDataUsingKeyAlgorithmError func(session ITKTokenSession, ciphertext objc.IObject /* cross-framework: NSData */, keyObjectID TKTokenObjectID /* typedef */, algorithm ITKTokenKeyAlgorithm, error_ unsafe.Pointer) foundation.Data
	_TokenSessionPerformKeyExchangeWithPublicKeyUsingKeyAlgorithmParametersError func(session ITKTokenSession, otherPartyPublicKeyData objc.IObject /* cross-framework: NSData */, objectID TKTokenObjectID /* typedef */, algorithm ITKTokenKeyAlgorithm, parameters ITKTokenKeyExchangeParameters, error_ unsafe.Pointer) foundation.Data
	_TokenSessionSignDataUsingKeyAlgorithmError func(session ITKTokenSession, dataToSign objc.IObject /* cross-framework: NSData */, keyObjectID TKTokenObjectID /* typedef */, algorithm ITKTokenKeyAlgorithm, error_ unsafe.Pointer) foundation.Data
	_TokenSessionSupportsOperationUsingKeyAlgorithm func(session ITKTokenSession, operation TKTokenOperation, keyObjectID TKTokenObjectID /* typedef */, algorithm ITKTokenKeyAlgorithm) bool
}

// SetTokenSessionBeginAuthForOperationConstraintError sets the handler for the TokenSessionBeginAuthForOperationConstraintError delegate method.
//
// Tells the delegate that authentication has begun for the specified operation and constraint.
func (d *TKTokenSessionDelegate) SetTokenSessionBeginAuthForOperationConstraintError(f func(session ITKTokenSession, operation TKTokenOperation, constraint TKTokenOperationConstraint /* typedef */, error_ unsafe.Pointer) TKTokenAuthOperation) {
	d._TokenSessionBeginAuthForOperationConstraintError = f
}

// SetTokenSessionDecryptDataUsingKeyAlgorithmError sets the handler for the TokenSessionDecryptDataUsingKeyAlgorithmError delegate method.
//
// Tells the delegate to decrypt a data object using the specified key and algorithm.
func (d *TKTokenSessionDelegate) SetTokenSessionDecryptDataUsingKeyAlgorithmError(f func(session ITKTokenSession, ciphertext objc.IObject /* cross-framework: NSData */, keyObjectID TKTokenObjectID /* typedef */, algorithm ITKTokenKeyAlgorithm, error_ unsafe.Pointer) foundation.Data) {
	d._TokenSessionDecryptDataUsingKeyAlgorithmError = f
}

// SetTokenSessionPerformKeyExchangeWithPublicKeyUsingKeyAlgorithmParametersError sets the handler for the TokenSessionPerformKeyExchangeWithPublicKeyUsingKeyAlgorithmParametersError delegate method.
//
// Tells the delegate to perform a key exchange using the specified key and algorithm.
func (d *TKTokenSessionDelegate) SetTokenSessionPerformKeyExchangeWithPublicKeyUsingKeyAlgorithmParametersError(f func(session ITKTokenSession, otherPartyPublicKeyData objc.IObject /* cross-framework: NSData */, objectID TKTokenObjectID /* typedef */, algorithm ITKTokenKeyAlgorithm, parameters ITKTokenKeyExchangeParameters, error_ unsafe.Pointer) foundation.Data) {
	d._TokenSessionPerformKeyExchangeWithPublicKeyUsingKeyAlgorithmParametersError = f
}

// SetTokenSessionSignDataUsingKeyAlgorithmError sets the handler for the TokenSessionSignDataUsingKeyAlgorithmError delegate method.
//
// Tells the delegate to sign a data object using the specified key and algorithm.
func (d *TKTokenSessionDelegate) SetTokenSessionSignDataUsingKeyAlgorithmError(f func(session ITKTokenSession, dataToSign objc.IObject /* cross-framework: NSData */, keyObjectID TKTokenObjectID /* typedef */, algorithm ITKTokenKeyAlgorithm, error_ unsafe.Pointer) foundation.Data) {
	d._TokenSessionSignDataUsingKeyAlgorithmError = f
}

// SetTokenSessionSupportsOperationUsingKeyAlgorithm sets the handler for the TokenSessionSupportsOperationUsingKeyAlgorithm delegate method.
//
// Asks the delegate whether the token session supports a given operation using the specified key and algorithm.
func (d *TKTokenSessionDelegate) SetTokenSessionSupportsOperationUsingKeyAlgorithm(f func(session ITKTokenSession, operation TKTokenOperation, keyObjectID TKTokenObjectID /* typedef */, algorithm ITKTokenKeyAlgorithm) bool) {
	d._TokenSessionSupportsOperationUsingKeyAlgorithm = f
}

// TokenSessionBeginAuthForOperationConstraintError implements the PTKTokenSessionDelegate interface.
func (d *TKTokenSessionDelegate) TokenSessionBeginAuthForOperationConstraintError(session ITKTokenSession, operation TKTokenOperation, constraint TKTokenOperationConstraint /* typedef */, error_ unsafe.Pointer) TKTokenAuthOperation {
	if d._TokenSessionBeginAuthForOperationConstraintError != nil {
		return d._TokenSessionBeginAuthForOperationConstraintError(session, operation, constraint, error_)
	}
	var zero TKTokenAuthOperation
	return zero
}

// HasTokenSessionBeginAuthForOperationConstraintError returns true if a handler for TokenSessionBeginAuthForOperationConstraintError has been set.
func (d *TKTokenSessionDelegate) HasTokenSessionBeginAuthForOperationConstraintError() bool {
	return d._TokenSessionBeginAuthForOperationConstraintError != nil
}

// TokenSessionDecryptDataUsingKeyAlgorithmError implements the PTKTokenSessionDelegate interface.
func (d *TKTokenSessionDelegate) TokenSessionDecryptDataUsingKeyAlgorithmError(session ITKTokenSession, ciphertext objc.IObject /* cross-framework: NSData */, keyObjectID TKTokenObjectID /* typedef */, algorithm ITKTokenKeyAlgorithm, error_ unsafe.Pointer) foundation.Data {
	if d._TokenSessionDecryptDataUsingKeyAlgorithmError != nil {
		return d._TokenSessionDecryptDataUsingKeyAlgorithmError(session, ciphertext, keyObjectID, algorithm, error_)
	}
	var zero foundation.Data
	return zero
}

// HasTokenSessionDecryptDataUsingKeyAlgorithmError returns true if a handler for TokenSessionDecryptDataUsingKeyAlgorithmError has been set.
func (d *TKTokenSessionDelegate) HasTokenSessionDecryptDataUsingKeyAlgorithmError() bool {
	return d._TokenSessionDecryptDataUsingKeyAlgorithmError != nil
}

// TokenSessionPerformKeyExchangeWithPublicKeyUsingKeyAlgorithmParametersError implements the PTKTokenSessionDelegate interface.
func (d *TKTokenSessionDelegate) TokenSessionPerformKeyExchangeWithPublicKeyUsingKeyAlgorithmParametersError(session ITKTokenSession, otherPartyPublicKeyData objc.IObject /* cross-framework: NSData */, objectID TKTokenObjectID /* typedef */, algorithm ITKTokenKeyAlgorithm, parameters ITKTokenKeyExchangeParameters, error_ unsafe.Pointer) foundation.Data {
	if d._TokenSessionPerformKeyExchangeWithPublicKeyUsingKeyAlgorithmParametersError != nil {
		return d._TokenSessionPerformKeyExchangeWithPublicKeyUsingKeyAlgorithmParametersError(session, otherPartyPublicKeyData, objectID, algorithm, parameters, error_)
	}
	var zero foundation.Data
	return zero
}

// HasTokenSessionPerformKeyExchangeWithPublicKeyUsingKeyAlgorithmParametersError returns true if a handler for TokenSessionPerformKeyExchangeWithPublicKeyUsingKeyAlgorithmParametersError has been set.
func (d *TKTokenSessionDelegate) HasTokenSessionPerformKeyExchangeWithPublicKeyUsingKeyAlgorithmParametersError() bool {
	return d._TokenSessionPerformKeyExchangeWithPublicKeyUsingKeyAlgorithmParametersError != nil
}

// TokenSessionSignDataUsingKeyAlgorithmError implements the PTKTokenSessionDelegate interface.
func (d *TKTokenSessionDelegate) TokenSessionSignDataUsingKeyAlgorithmError(session ITKTokenSession, dataToSign objc.IObject /* cross-framework: NSData */, keyObjectID TKTokenObjectID /* typedef */, algorithm ITKTokenKeyAlgorithm, error_ unsafe.Pointer) foundation.Data {
	if d._TokenSessionSignDataUsingKeyAlgorithmError != nil {
		return d._TokenSessionSignDataUsingKeyAlgorithmError(session, dataToSign, keyObjectID, algorithm, error_)
	}
	var zero foundation.Data
	return zero
}

// HasTokenSessionSignDataUsingKeyAlgorithmError returns true if a handler for TokenSessionSignDataUsingKeyAlgorithmError has been set.
func (d *TKTokenSessionDelegate) HasTokenSessionSignDataUsingKeyAlgorithmError() bool {
	return d._TokenSessionSignDataUsingKeyAlgorithmError != nil
}

// TokenSessionSupportsOperationUsingKeyAlgorithm implements the PTKTokenSessionDelegate interface.
func (d *TKTokenSessionDelegate) TokenSessionSupportsOperationUsingKeyAlgorithm(session ITKTokenSession, operation TKTokenOperation, keyObjectID TKTokenObjectID /* typedef */, algorithm ITKTokenKeyAlgorithm) bool {
	if d._TokenSessionSupportsOperationUsingKeyAlgorithm != nil {
		return d._TokenSessionSupportsOperationUsingKeyAlgorithm(session, operation, keyObjectID, algorithm)
	}
	var zero bool
	return zero
}

// HasTokenSessionSupportsOperationUsingKeyAlgorithm returns true if a handler for TokenSessionSupportsOperationUsingKeyAlgorithm has been set.
func (d *TKTokenSessionDelegate) HasTokenSessionSupportsOperationUsingKeyAlgorithm() bool {
	return d._TokenSessionSupportsOperationUsingKeyAlgorithm != nil
}
