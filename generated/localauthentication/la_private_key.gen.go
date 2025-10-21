// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PrivateKey] class.
var (
	PrivateKeyClass     _PrivateKeyClass
	PrivateKeyClassOnce sync.Once
)

func getPrivateKeyClass() _PrivateKeyClass {
	PrivateKeyClassOnce.Do(func() {
		PrivateKeyClass = _PrivateKeyClass{objc.GetClass("LAPrivateKey")}
	})
	return PrivateKeyClass
}

type _PrivateKeyClass struct {
	class objc.Class
}

// An interface definition for the [PrivateKey] class.
type IPrivateKey interface {
	objectivec.IObject
	CanDecryptUsingSecKeyAlgorithm(algorithm unsafe.Pointer) bool
	CanExchangeKeysUsingSecKeyAlgorithm(algorithm unsafe.Pointer) bool
	CanSignUsingSecKeyAlgorithm(algorithm unsafe.Pointer) bool
	DecryptDataSecKeyAlgorithmCompletion(data unsafe.Pointer, algorithm unsafe.Pointer, handler unsafe.Pointer)
	ExchangeKeysWithPublicKeySecKeyAlgorithmSecKeyParametersCompletion(publicKey unsafe.Pointer, algorithm unsafe.Pointer, parameters objc.ID, handler unsafe.Pointer)
	SignDataSecKeyAlgorithmCompletion(data unsafe.Pointer, algorithm unsafe.Pointer, handler unsafe.Pointer)
}

// The private portion of an asymmetric key pair.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPrivateKey
type PrivateKey struct {
	objectivec.Object
}

// PrivateKeyFrom constructs a [PrivateKey] from an unsafe.Pointer.
//
// The private portion of an asymmetric key pair.
func PrivateKeyFrom(ptr unsafe.Pointer) PrivateKey {
	return PrivateKey{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PrivateKeyClass) Alloc() PrivateKey {
	rv := objc.Send[PrivateKey](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PrivateKeyClass) New() PrivateKey {
	rv := objc.Send[PrivateKey](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PrivateKey) Init() PrivateKey {
	rv := objc.Send[PrivateKey](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PrivateKey) Autorelease() PrivateKey {
	rv := objc.Send[PrivateKey](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPrivateKey creates a new PrivateKey instance.
func NewPrivateKey() PrivateKey {
	return getPrivateKeyClass().New()
}


// Checks whether the algorithm you supply is valid for decrypting data with the key.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPrivateKey/canDecrypt(using:)
func (p_ PrivateKey) CanDecryptUsingSecKeyAlgorithm(algorithm unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canDecryptUsingSecKeyAlgorithm:"), algorithm)
	return rv
}

// Checks whether the algorithm you supply is valid for performing key exchanges.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPrivateKey/canExchangeKeys(using:)
func (p_ PrivateKey) CanExchangeKeysUsingSecKeyAlgorithm(algorithm unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canExchangeKeysUsingSecKeyAlgorithm:"), algorithm)
	return rv
}

// Checks whether the algorithm you supply is valid for signing data with the key.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPrivateKey/canSign(using:)
func (p_ PrivateKey) CanSignUsingSecKeyAlgorithm(algorithm unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canSignUsingSecKeyAlgorithm:"), algorithm)
	return rv
}

// Decrypts the data you supply with a given algorithm.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPrivateKey/decrypt(_:algorithm:completion:)
func (p_ PrivateKey) DecryptDataSecKeyAlgorithmCompletion(data unsafe.Pointer, algorithm unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("decryptData:secKeyAlgorithm:completion:"), data, algorithm, handler)
}

// Performs a Diffie-Hellman style key exchange operation.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPrivateKey/exchangeKeys(publicKey:algorithm:parameters:completion:)
func (p_ PrivateKey) ExchangeKeysWithPublicKeySecKeyAlgorithmSecKeyParametersCompletion(publicKey unsafe.Pointer, algorithm unsafe.Pointer, parameters objc.ID, handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("exchangeKeysWithPublicKey:secKeyAlgorithm:secKeyParameters:completion:"), publicKey, algorithm, parameters, handler)
}

// Generates a digital signature for the data you supply.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPrivateKey/sign(_:algorithm:completion:)
func (p_ PrivateKey) SignDataSecKeyAlgorithmCompletion(data unsafe.Pointer, algorithm unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("signData:secKeyAlgorithm:completion:"), data, algorithm, handler)
}

// The public key that corresponds with the private key in a key pair.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPrivateKey/publicKey
func (p_ PrivateKey) PublicKey() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("publicKey"))
	return rv
}



