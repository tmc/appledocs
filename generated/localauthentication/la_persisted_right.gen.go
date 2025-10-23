// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PersistedRight] class.
var (
	PersistedRightClass     _PersistedRightClass
	PersistedRightClassOnce sync.Once
)

func getPersistedRightClass() _PersistedRightClass {
	PersistedRightClassOnce.Do(func() {
		PersistedRightClass = _PersistedRightClass{objc.GetClass("LAPersistedRight")}
	})
	return PersistedRightClass
}

type _PersistedRightClass struct {
	class objc.Class
}

// An interface definition for the [PersistedRight] class.
type IPersistedRight interface {
	IRight
	// properties:
	Key() ILAPrivateKey
	Secret() ILASecret
	// methods:
}

// A right that gates access to a key and a secret.
//
// An is a right that’s backed by a unique key in the Secure Enclave with an access control list that matches the authorization requirements of the right. You can access the key that backs a right to perform cryptographic operations like encryption, decryption, signing, and verification. You can use the key that backs an to perform both public key and private key operations, but private key operations — like decryption, signing, and key exchange — are only available after you authorize the right. Public key operations like encryption and verification are always available. The following generates a right with the default authorization requirements, stores it in the , and exports the public key so that you can use it to verify signatures that the corresponding private key produces: The following uses the private key associated with the right from the previous example to sign a challenge issued by a server: The signature operation occurs after verifying that the user has the proper authorization and confirming that the private key supports the given signing algorithm.


// A right that gates access to a key and a secret.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPersistedRight
type PersistedRight struct {
	Right
}

// PersistedRightFrom constructs a [PersistedRight] from an unsafe.Pointer.
//
// A right that gates access to a key and a secret.
func PersistedRightFrom(ptr unsafe.Pointer) PersistedRight {
	return PersistedRight{
		Right: RightFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PersistedRightClass) Alloc() PersistedRight {
	rv := objc.Send[PersistedRight](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PersistedRightClass) New() PersistedRight {
	rv := objc.Send[PersistedRight](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersistedRight) Init() PersistedRight {
	rv := objc.Send[PersistedRight](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersistedRight) Autorelease() PersistedRight {
	rv := objc.Send[PersistedRight](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersistedRight creates a new PersistedRight instance.
func NewPersistedRight() PersistedRight {
	return getPersistedRightClass().New()
}



// The private key that’s persisted by the right.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPersistedRight/key
func (p_ PersistedRight) Key() ILAPrivateKey {
	rv := objc.Send[PrivateKey](p_.ID, objc.Sel("key"))
	return rv
}


// The data kept secret by the right.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAPersistedRight/secret
func (p_ PersistedRight) Secret() ILASecret {
	rv := objc.Send[Secret](p_.ID, objc.Sel("secret"))
	return rv
}



