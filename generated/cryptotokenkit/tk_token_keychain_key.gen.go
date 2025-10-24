// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class TKTokenKeychainKey */


/* debug [class_header]: Header for TKTokenKeychainKey */
// The class instance for the [TKTokenKeychainKey] class.
var (
	TKTokenKeychainKeyClass     _TKTokenKeychainKeyClass
	TKTokenKeychainKeyClassOnce sync.Once
)

func getTKTokenKeychainKeyClass() _TKTokenKeychainKeyClass {
	TKTokenKeychainKeyClassOnce.Do(func() {
		TKTokenKeychainKeyClass = _TKTokenKeychainKeyClass{objc.GetClass("TKTokenKeychainKey")}
	})
	return TKTokenKeychainKeyClass
}

type _TKTokenKeychainKeyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKTokenKeychainKey */
// An interface definition for the [TKTokenKeychainKey] class.
type ITKTokenKeychainKey interface {
	ITKTokenKeychainItem
	
/* debug [class_interface_properties]: Properties for TKTokenKeychainKey */
	// properties:
	ApplicationTag() objc.IObject /* cross-framework: NSData */
	SetApplicationTag(value objc.IObject /* cross-framework: NSData */)
	CanDecrypt() bool
	SetCanDecrypt(value bool)
	CanPerformKeyExchange() bool
	SetCanPerformKeyExchange(value bool)
	CanSign() bool
	SetCanSign(value bool)
	SuitableForLogin() bool
	SetSuitableForLogin(value bool)
	KeySizeInBits() int
	SetKeySizeInBits(value int)
	KeyType() objc.IObject /* cross-framework: NSString */
	SetKeyType(value objc.IObject /* cross-framework: NSString */)
	PublicKeyData() objc.IObject /* cross-framework: NSData */
	SetPublicKeyData(value objc.IObject /* cross-framework: NSData */)
	PublicKeyHash() objc.IObject /* cross-framework: NSData */
	SetPublicKeyHash(value objc.IObject /* cross-framework: NSData */)
	KeychainContents() ITKTokenKeychainContents
	SetKeychainContents(value ITKTokenKeychainContents)
	IsSuitableForLogin() bool
	SetIsSuitableForLogin(value bool)
	KSecAttrKeyTypeRSA() foundation.String
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKTokenKeychainKey */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKTokenKeychainKey */
// Alloc allocates a new instance without initialization.
func (tc _TKTokenKeychainKeyClass) Alloc() TKTokenKeychainKey {
	rv := objc.Send[TKTokenKeychainKey](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKTokenKeychainKeyClass) New() TKTokenKeychainKey {
	rv := objc.Send[TKTokenKeychainKey](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKTokenKeychainKey) Init() TKTokenKeychainKey {
	rv := objc.Send[TKTokenKeychainKey](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKTokenKeychainKey) Autorelease() TKTokenKeychainKey {
	rv := objc.Send[TKTokenKeychainKey](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenKeychainKey creates a new TKTokenKeychainKey instance.
func NewTKTokenKeychainKey() TKTokenKeychainKey {
	return getTKTokenKeychainKeyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKTokenKeychainKey */
// A token’s key as stored in the keychain.


// A token’s key as stored in the keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey
type TKTokenKeychainKey struct {
	TKTokenKeychainItem
}

// TKTokenKeychainKeyFrom constructs a [TKTokenKeychainKey] from an unsafe.Pointer.
//
// A token’s key as stored in the keychain.
func TKTokenKeychainKeyFrom(ptr unsafe.Pointer) TKTokenKeychainKey {
	return TKTokenKeychainKey{
		TKTokenKeychainItem: TKTokenKeychainItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKTokenKeychainKey */

// Initializes a token keychain key with data from the specified certificate reference and a given object ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/init(certificate:objectID:)
func NewTKTokenKeychainKeyWithCertificateObjectID(certificateRef unsafe.Pointer, objectID TKTokenObjectID /* typedef */) TKTokenKeychainKey {
	instance := getTKTokenKeychainKeyClass().Alloc()
	rv := objc.Send[TKTokenKeychainKey](instance.ID, objc.Sel("initWithCertificate:objectID:"), certificateRef, objectID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTKTokenKeychainKeyWithCertificateObjectID */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKTokenKeychainKey */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKTokenKeychainKey */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKTokenKeychainKey */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKTokenKeychainKey */

// The private tag data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/applicationTag
func (t_ TKTokenKeychainKey) ApplicationTag() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("applicationTag"))
	return rv
}/* debug [instance_properties/getter]: applicationTag */


// The private tag data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/applicationTag
func (t_ TKTokenKeychainKey) SetApplicationTag(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setApplicationTag:"), value)
}/* debug [instance_properties/setter]: applicationTag */


// Whether the key can be used to decrypt data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/canDecrypt
func (t_ TKTokenKeychainKey) CanDecrypt() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("canDecrypt"))
	return rv
}/* debug [instance_properties/getter]: canDecrypt */


// Whether the key can be used to decrypt data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/canDecrypt
func (t_ TKTokenKeychainKey) SetCanDecrypt(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCanDecrypt:"), value)
}/* debug [instance_properties/setter]: canDecrypt */


// Whether the key can be used to perform Diffie-Hellman style cryptographic key exchange.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/canPerformKeyExchange
func (t_ TKTokenKeychainKey) CanPerformKeyExchange() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("canPerformKeyExchange"))
	return rv
}/* debug [instance_properties/getter]: canPerformKeyExchange */


// Whether the key can be used to perform Diffie-Hellman style cryptographic key exchange.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/canPerformKeyExchange
func (t_ TKTokenKeychainKey) SetCanPerformKeyExchange(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCanPerformKeyExchange:"), value)
}/* debug [instance_properties/setter]: canPerformKeyExchange */


// Whether the key can be used to sign data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/canSign
func (t_ TKTokenKeychainKey) CanSign() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("canSign"))
	return rv
}/* debug [instance_properties/getter]: canSign */


// Whether the key can be used to sign data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/canSign
func (t_ TKTokenKeychainKey) SetCanSign(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCanSign:"), value)
}/* debug [instance_properties/setter]: canSign */


// Whether the key can be used for system login.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/isSuitableForLogin
func (t_ TKTokenKeychainKey) SuitableForLogin() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("suitableForLogin"))
	return rv
}/* debug [instance_properties/getter]: suitableForLogin */


// Whether the key can be used for system login.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/isSuitableForLogin
func (t_ TKTokenKeychainKey) SetSuitableForLogin(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSuitableForLogin:"), value)
}/* debug [instance_properties/setter]: suitableForLogin */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/keySizeInBits
func (t_ TKTokenKeychainKey) KeySizeInBits() int {
	rv := objc.Send[int](t_.ID, objc.Sel("keySizeInBits"))
	return rv
}/* debug [instance_properties/getter]: keySizeInBits */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/keySizeInBits
func (t_ TKTokenKeychainKey) SetKeySizeInBits(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setKeySizeInBits:"), value)
}/* debug [instance_properties/setter]: keySizeInBits */


// The type of the key. Currently, only and are supported values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/keyType
func (t_ TKTokenKeychainKey) KeyType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("keyType"))
	return rv
}/* debug [instance_properties/getter]: keyType */


// The type of the key. Currently, only and are supported values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/keyType
func (t_ TKTokenKeychainKey) SetKeyType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setKeyType:"), value)
}/* debug [instance_properties/setter]: keyType */


// The public key data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/publicKeyData
func (t_ TKTokenKeychainKey) PublicKeyData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("publicKeyData"))
	return rv
}/* debug [instance_properties/getter]: publicKeyData */


// The public key data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/publicKeyData
func (t_ TKTokenKeychainKey) SetPublicKeyData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPublicKeyData:"), value)
}/* debug [instance_properties/setter]: publicKeyData */


// The SHA1 hash of the raw public key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/publicKeyHash
func (t_ TKTokenKeychainKey) PublicKeyHash() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("publicKeyHash"))
	return rv
}/* debug [instance_properties/getter]: publicKeyHash */


// The SHA1 hash of the raw public key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainKey/publicKeyHash
func (t_ TKTokenKeychainKey) SetPublicKeyHash(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPublicKeyHash:"), value)
}/* debug [instance_properties/setter]: publicKeyHash */


// The contents of the keychain for this token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktoken/keychaincontents
func (t_ TKTokenKeychainKey) KeychainContents() ITKTokenKeychainContents {
	rv := objc.Send[TKTokenKeychainContents](t_.ID, objc.Sel("keychainContents"))
	return rv
}/* debug [instance_properties/getter]: keychainContents */


// The contents of the keychain for this token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktoken/keychaincontents
func (t_ TKTokenKeychainKey) SetKeychainContents(value ITKTokenKeychainContents) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setKeychainContents:"), value)
}/* debug [instance_properties/setter]: keychainContents */


// Whether the key can be used for system login.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktokenkeychainkey/issuitableforlogin
func (t_ TKTokenKeychainKey) IsSuitableForLogin() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isSuitableForLogin"))
	return rv
}/* debug [instance_properties/getter]: isSuitableForLogin */


// Whether the key can be used for system login.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktokenkeychainkey/issuitableforlogin
func (t_ TKTokenKeychainKey) SetIsSuitableForLogin(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsSuitableForLogin:"), value)
}/* debug [instance_properties/setter]: isSuitableForLogin */


// RSA algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/kSecAttrKeyTypeRSA
func (t_ TKTokenKeychainKey) KSecAttrKeyTypeRSA() foundation.String {
	rv := objc.Send[foundation.String](t_.ID, objc.Sel("kSecAttrKeyTypeRSA"))
	return rv
}/* debug [instance_properties/getter]: kSecAttrKeyTypeRSA */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKTokenKeychainKey */


