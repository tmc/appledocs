// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PublicKey] class.
var (
	PublicKeyClass     _PublicKeyClass
	PublicKeyClassOnce sync.Once
)

func getPublicKeyClass() _PublicKeyClass {
	PublicKeyClassOnce.Do(func() {
		PublicKeyClass = _PublicKeyClass{objc.GetClass("LAPublicKey")}
	})
	return PublicKeyClass
}

type _PublicKeyClass struct {
	class objc.Class
}

// An interface definition for the [PublicKey] class.
type IPublicKey interface {
	objectivec.IObject
	// properties:
	// methods:
	CanEncryptUsingSecKeyAlgorithm(algorithm unsafe.Pointer) bool
	CanVerifyUsingSecKeyAlgorithm(algorithm unsafe.Pointer) bool
	EncryptDataSecKeyAlgorithmCompletion(data foundation.NSData, algorithm unsafe.Pointer, handler unsafe.Pointer)
	ExportBytesWithCompletion(handler unsafe.Pointer)
	VerifyDataSignatureSecKeyAlgorithmCompletion(signedData foundation.NSData, signature foundation.NSData, algorithm unsafe.Pointer, handler unsafe.Pointer)
}

// The public portion of an asymmetric key pair.


// The public portion of an asymmetric key pair.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPublicKey
type PublicKey struct {
	objectivec.Object
}

// PublicKeyFrom constructs a [PublicKey] from an unsafe.Pointer.
//
// The public portion of an asymmetric key pair.
func PublicKeyFrom(ptr unsafe.Pointer) PublicKey {
	return PublicKey{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PublicKeyClass) Alloc() PublicKey {
	rv := objc.Send[PublicKey](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PublicKeyClass) New() PublicKey {
	rv := objc.Send[PublicKey](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PublicKey) Init() PublicKey {
	rv := objc.Send[PublicKey](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PublicKey) Autorelease() PublicKey {
	rv := objc.Send[PublicKey](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPublicKey creates a new PublicKey instance.
func NewPublicKey() PublicKey {
	return getPublicKeyClass().New()
}



// Checks whether the algorithm you supply is valid for encrypting data with the key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPublicKey/canEncrypt(using:)
func (p_ PublicKey) CanEncryptUsingSecKeyAlgorithm(algorithm unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canEncryptUsingSecKeyAlgorithm:"), algorithm)
	return rv
}


// Checks whether the algorithm you supply is valid for verifying signatures with the key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPublicKey/canVerify(using:)
func (p_ PublicKey) CanVerifyUsingSecKeyAlgorithm(algorithm unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canVerifyUsingSecKeyAlgorithm:"), algorithm)
	return rv
}


// Encrypts the data you supply with a given algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPublicKey/encrypt(_:algorithm:completion:)
func (p_ PublicKey) EncryptDataSecKeyAlgorithmCompletion(data foundation.NSData, algorithm unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("encryptData:secKeyAlgorithm:completion:"), data, algorithm, handler)
}


// Exports the data that represents a public key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPublicKey/exportBytes(completion:)
func (p_ PublicKey) ExportBytesWithCompletion(handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("exportBytesWithCompletion:"), handler)
}


// Verifies a digital signature for the data you supply.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPublicKey/verify(_:signature:algorithm:completion:)
func (p_ PublicKey) VerifyDataSignatureSecKeyAlgorithmCompletion(signedData foundation.NSData, signature foundation.NSData, algorithm unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("verifyData:signature:secKeyAlgorithm:completion:"), signedData, signature, algorithm, handler)
}



