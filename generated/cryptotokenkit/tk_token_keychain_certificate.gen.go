// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class TKTokenKeychainCertificate */


/* debug [class_header]: Header for TKTokenKeychainCertificate */
// The class instance for the [TKTokenKeychainCertificate] class.
var (
	TKTokenKeychainCertificateClass     _TKTokenKeychainCertificateClass
	TKTokenKeychainCertificateClassOnce sync.Once
)

func getTKTokenKeychainCertificateClass() _TKTokenKeychainCertificateClass {
	TKTokenKeychainCertificateClassOnce.Do(func() {
		TKTokenKeychainCertificateClass = _TKTokenKeychainCertificateClass{objc.GetClass("TKTokenKeychainCertificate")}
	})
	return TKTokenKeychainCertificateClass
}

type _TKTokenKeychainCertificateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKTokenKeychainCertificate */
// An interface definition for the [TKTokenKeychainCertificate] class.
type ITKTokenKeychainCertificate interface {
	ITKTokenKeychainItem
	
/* debug [class_interface_properties]: Properties for TKTokenKeychainCertificate */
	// properties:
	Data() objc.IObject /* cross-framework: NSData */
	KeychainContents() ITKTokenKeychainContents
	SetKeychainContents(value ITKTokenKeychainContents)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKTokenKeychainCertificate */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKTokenKeychainCertificate */
// Alloc allocates a new instance without initialization.
func (tc _TKTokenKeychainCertificateClass) Alloc() TKTokenKeychainCertificate {
	rv := objc.Send[TKTokenKeychainCertificate](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKTokenKeychainCertificateClass) New() TKTokenKeychainCertificate {
	rv := objc.Send[TKTokenKeychainCertificate](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKTokenKeychainCertificate) Init() TKTokenKeychainCertificate {
	rv := objc.Send[TKTokenKeychainCertificate](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKTokenKeychainCertificate) Autorelease() TKTokenKeychainCertificate {
	rv := objc.Send[TKTokenKeychainCertificate](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenKeychainCertificate creates a new TKTokenKeychainCertificate instance.
func NewTKTokenKeychainCertificate() TKTokenKeychainCertificate {
	return getTKTokenKeychainCertificateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKTokenKeychainCertificate */
// A token’s certificate as stored in the keychain.


// A token’s certificate as stored in the keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainCertificate
type TKTokenKeychainCertificate struct {
	TKTokenKeychainItem
}

// TKTokenKeychainCertificateFrom constructs a [TKTokenKeychainCertificate] from an unsafe.Pointer.
//
// A token’s certificate as stored in the keychain.
func TKTokenKeychainCertificateFrom(ptr unsafe.Pointer) TKTokenKeychainCertificate {
	return TKTokenKeychainCertificate{
		TKTokenKeychainItem: TKTokenKeychainItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKTokenKeychainCertificate */

// Initializes a token keychain certificate with data from the specified certificate reference and a given object ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainCertificate/init(certificate:objectID:)
func NewTKTokenKeychainCertificateWithCertificateObjectID(certificateRef unsafe.Pointer, objectID TKTokenObjectID /* typedef */) TKTokenKeychainCertificate {
	instance := getTKTokenKeychainCertificateClass().Alloc()
	rv := objc.Send[TKTokenKeychainCertificate](instance.ID, objc.Sel("initWithCertificate:objectID:"), certificateRef, objectID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTKTokenKeychainCertificateWithCertificateObjectID */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKTokenKeychainCertificate */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKTokenKeychainCertificate */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKTokenKeychainCertificate */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKTokenKeychainCertificate */

// Returns a DER-encoded representation of an X.509 certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainCertificate/data
func (t_ TKTokenKeychainCertificate) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// The contents of the keychain for this token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktoken/keychaincontents
func (t_ TKTokenKeychainCertificate) KeychainContents() ITKTokenKeychainContents {
	rv := objc.Send[TKTokenKeychainContents](t_.ID, objc.Sel("keychainContents"))
	return rv
}/* debug [instance_properties/getter]: keychainContents */


// The contents of the keychain for this token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktoken/keychaincontents
func (t_ TKTokenKeychainCertificate) SetKeychainContents(value ITKTokenKeychainContents) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setKeychainContents:"), value)
}/* debug [instance_properties/setter]: keychainContents */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKTokenKeychainCertificate */


