// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class TKSmartCardToken */


/* debug [class_header]: Header for TKSmartCardToken */
// The class instance for the [TKSmartCardToken] class.
var (
	TKSmartCardTokenClass     _TKSmartCardTokenClass
	TKSmartCardTokenClassOnce sync.Once
)

func getTKSmartCardTokenClass() _TKSmartCardTokenClass {
	TKSmartCardTokenClassOnce.Do(func() {
		TKSmartCardTokenClass = _TKSmartCardTokenClass{objc.GetClass("TKSmartCardToken")}
	})
	return TKSmartCardTokenClass
}

type _TKSmartCardTokenClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKSmartCardToken */
// An interface definition for the [TKSmartCardToken] class.
type ITKSmartCardToken interface {
	ITKToken
	
/* debug [class_interface_properties]: Properties for TKSmartCardToken */
	// properties:
	AID() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKSmartCardToken */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKSmartCardToken */
// Alloc allocates a new instance without initialization.
func (tc _TKSmartCardTokenClass) Alloc() TKSmartCardToken {
	rv := objc.Send[TKSmartCardToken](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKSmartCardTokenClass) New() TKSmartCardToken {
	rv := objc.Send[TKSmartCardToken](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKSmartCardToken) Init() TKSmartCardToken {
	rv := objc.Send[TKSmartCardToken](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKSmartCardToken) Autorelease() TKSmartCardToken {
	rv := objc.Send[TKSmartCardToken](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardToken creates a new TKSmartCardToken instance.
func NewTKSmartCardToken() TKSmartCardToken {
	return getTKSmartCardTokenClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKSmartCardToken */
// A representation of a smart card based cryptographic token.


// A representation of a smart card based cryptographic token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardToken
type TKSmartCardToken struct {
	TKToken
}

// TKSmartCardTokenFrom constructs a [TKSmartCardToken] from an unsafe.Pointer.
//
// A representation of a smart card based cryptographic token.
func TKSmartCardTokenFrom(ptr unsafe.Pointer) TKSmartCardToken {
	return TKSmartCardToken{
		TKToken: TKTokenFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKSmartCardToken */

// Initializes a smart card token with the specified smart card, application identifier, and token driver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardToken/init(smartCard:aid:instanceID:tokenDriver:)
func NewTKSmartCardTokenWithSmartCardAIDInstanceIDTokenDriver(smartCard ITKSmartCard, AID objc.IObject /* cross-framework: NSData */, instanceID objc.IObject /* cross-framework: NSString */, tokenDriver ITKSmartCardTokenDriver) TKSmartCardToken {
	instance := getTKSmartCardTokenClass().Alloc()
	rv := objc.Send[TKSmartCardToken](instance.ID, objc.Sel("initWithSmartCard:AID:instanceID:tokenDriver:"), smartCard, AID, instanceID, tokenDriver)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTKSmartCardTokenWithSmartCardAIDInstanceIDTokenDriver */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKSmartCardToken */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKSmartCardToken */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKSmartCardToken */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKSmartCardToken */

// The ISO 7816-4 application identifiers of the Smart Card.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardToken/aid
func (t_ TKSmartCardToken) AID() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("AID"))
	return rv
}/* debug [instance_properties/getter]: AID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKSmartCardToken */


