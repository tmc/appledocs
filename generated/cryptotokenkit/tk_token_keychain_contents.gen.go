// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class TKTokenKeychainContents */


/* debug [class_header]: Header for TKTokenKeychainContents */
// The class instance for the [TKTokenKeychainContents] class.
var (
	TKTokenKeychainContentsClass     _TKTokenKeychainContentsClass
	TKTokenKeychainContentsClassOnce sync.Once
)

func getTKTokenKeychainContentsClass() _TKTokenKeychainContentsClass {
	TKTokenKeychainContentsClassOnce.Do(func() {
		TKTokenKeychainContentsClass = _TKTokenKeychainContentsClass{objc.GetClass("TKTokenKeychainContents")}
	})
	return TKTokenKeychainContentsClass
}

type _TKTokenKeychainContentsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKTokenKeychainContents */
// An interface definition for the [TKTokenKeychainContents] class.
type ITKTokenKeychainContents interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TKTokenKeychainContents */
	// properties:
	Items() []TKTokenKeychainItem
	KeychainContents() ITKTokenKeychainContents
	SetKeychainContents(value ITKTokenKeychainContents)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKTokenKeychainContents */
	// methods:
	CertificateForObjectIDError(objectID TKTokenObjectID /* typedef */, error_ unsafe.Pointer) ITKTokenKeychainCertificate
	FillWithItems(items []TKTokenKeychainItem)
	KeyForObjectIDError(objectID TKTokenObjectID /* typedef */, error_ unsafe.Pointer) ITKTokenKeychainKey
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKTokenKeychainContents */
// Alloc allocates a new instance without initialization.
func (tc _TKTokenKeychainContentsClass) Alloc() TKTokenKeychainContents {
	rv := objc.Send[TKTokenKeychainContents](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKTokenKeychainContentsClass) New() TKTokenKeychainContents {
	rv := objc.Send[TKTokenKeychainContents](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKTokenKeychainContents) Init() TKTokenKeychainContents {
	rv := objc.Send[TKTokenKeychainContents](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKTokenKeychainContents) Autorelease() TKTokenKeychainContents {
	rv := objc.Send[TKTokenKeychainContents](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenKeychainContents creates a new TKTokenKeychainContents instance.
func NewTKTokenKeychainContents() TKTokenKeychainContents {
	return getTKTokenKeychainContentsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKTokenKeychainContents */
// A representation of the state of the keychain for a particular token.


// A representation of the state of the keychain for a particular token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainContents
type TKTokenKeychainContents struct {
	objectivec.Object
}

// TKTokenKeychainContentsFrom constructs a [TKTokenKeychainContents] from an unsafe.Pointer.
//
// A representation of the state of the keychain for a particular token.
func TKTokenKeychainContentsFrom(ptr unsafe.Pointer) TKTokenKeychainContents {
	return TKTokenKeychainContents{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKTokenKeychainContents *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKTokenKeychainContents */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKTokenKeychainContents */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKTokenKeychainContents */

// Returns the key for a specified object identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainContents/certificate(forObjectID:)
func (t_ TKTokenKeychainContents) CertificateForObjectIDError(objectID TKTokenObjectID /* typedef */, error_ unsafe.Pointer) ITKTokenKeychainCertificate {
	rv := objc.Send[TKTokenKeychainCertificate](t_.ID, objc.Sel("certificateForObjectID:error:"), objectID, error_)
	return rv
}/* debug [instance_methods/method]: CertificateForObjectIDError */


// Fills the keychain with the specified items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainContents/fill(with:)
func (t_ TKTokenKeychainContents) FillWithItems(items []TKTokenKeychainItem) {
	objc.Send[objc.ID](t_.ID, objc.Sel("fillWithItems:"), items)
}/* debug [instance_methods/method]: FillWithItems */


// Returns the key for a specified object identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainContents/key(forObjectID:)
func (t_ TKTokenKeychainContents) KeyForObjectIDError(objectID TKTokenObjectID /* typedef */, error_ unsafe.Pointer) ITKTokenKeychainKey {
	rv := objc.Send[TKTokenKeychainKey](t_.ID, objc.Sel("keyForObjectID:error:"), objectID, error_)
	return rv
}/* debug [instance_methods/method]: KeyForObjectIDError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKTokenKeychainContents */

// Returns all items for token in the keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeychainContents/items
func (t_ TKTokenKeychainContents) Items() []TKTokenKeychainItem {
	rv := objc.Send[[]TKTokenKeychainItem](t_.ID, objc.Sel("items"))
	return rv
}/* debug [instance_properties/getter]: items */


// The contents of the keychain for this token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktoken/keychaincontents
func (t_ TKTokenKeychainContents) KeychainContents() ITKTokenKeychainContents {
	rv := objc.Send[TKTokenKeychainContents](t_.ID, objc.Sel("keychainContents"))
	return rv
}/* debug [instance_properties/getter]: keychainContents */


// The contents of the keychain for this token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktoken/keychaincontents
func (t_ TKTokenKeychainContents) SetKeychainContents(value ITKTokenKeychainContents) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setKeychainContents:"), value)
}/* debug [instance_properties/setter]: keychainContents */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKTokenKeychainContents */



