// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class TKToken */


/* debug [class_header]: Header for TKToken */
// The class instance for the [TKToken] class.
var (
	TKTokenClass     _TKTokenClass
	TKTokenClassOnce sync.Once
)

func getTKTokenClass() _TKTokenClass {
	TKTokenClassOnce.Do(func() {
		TKTokenClass = _TKTokenClass{objc.GetClass("TKToken")}
	})
	return TKTokenClass
}

type _TKTokenClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKToken */
// An interface definition for the [TKToken] class.
type ITKToken interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TKToken */
	// properties:
	Configuration() ITKTokenConfiguration
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	KeychainContents() ITKTokenKeychainContents
	TokenDriver() ITKTokenDriver
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKToken */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKToken */
// Alloc allocates a new instance without initialization.
func (tc _TKTokenClass) Alloc() TKToken {
	rv := objc.Send[TKToken](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKTokenClass) New() TKToken {
	rv := objc.Send[TKToken](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKToken) Init() TKToken {
	rv := objc.Send[TKToken](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKToken) Autorelease() TKToken {
	rv := objc.Send[TKToken](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKToken creates a new TKToken instance.
func NewTKToken() TKToken {
	return getTKTokenClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKToken */
// A representation of a hardware-based cryptographic token.


// A representation of a hardware-based cryptographic token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKToken
type TKToken struct {
	objectivec.Object
}

// TKTokenFrom constructs a [TKToken] from an unsafe.Pointer.
//
// A representation of a hardware-based cryptographic token.
func TKTokenFrom(ptr unsafe.Pointer) TKToken {
	return TKToken{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKToken */

// Initializes a token with the driver you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/init(tokenDriver:instanceID:)
func NewTKTokenWithTokenDriverInstanceID(tokenDriver ITKTokenDriver, instanceID TKTokenInstanceID /* typedef */) TKToken {
	instance := getTKTokenClass().Alloc()
	rv := objc.Send[TKToken](instance.ID, objc.Sel("initWithTokenDriver:instanceID:"), tokenDriver, instanceID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTKTokenWithTokenDriverInstanceID */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKToken */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKToken */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKToken */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKToken */

// The current configuration for a token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/configuration-swift.property
func (t_ TKToken) Configuration() ITKTokenConfiguration {
	rv := objc.Send[TKTokenConfiguration](t_.ID, objc.Sel("configuration"))
	return rv
}/* debug [instance_properties/getter]: configuration */


// The token delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/delegate
func (t_ TKToken) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The token delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/delegate
func (t_ TKToken) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The contents of the keychain for this token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/keychainContents
func (t_ TKToken) KeychainContents() ITKTokenKeychainContents {
	rv := objc.Send[TKTokenKeychainContents](t_.ID, objc.Sel("keychainContents"))
	return rv
}/* debug [instance_properties/getter]: keychainContents */


// The token driver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/tokenDriver
func (t_ TKToken) TokenDriver() ITKTokenDriver {
	rv := objc.Send[TKTokenDriver](t_.ID, objc.Sel("tokenDriver"))
	return rv
}/* debug [instance_properties/getter]: tokenDriver */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKToken */


