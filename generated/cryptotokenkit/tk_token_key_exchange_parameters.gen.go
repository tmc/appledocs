// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class TKTokenKeyExchangeParameters */


/* debug [class_header]: Header for TKTokenKeyExchangeParameters */
// The class instance for the [TKTokenKeyExchangeParameters] class.
var (
	TKTokenKeyExchangeParametersClass     _TKTokenKeyExchangeParametersClass
	TKTokenKeyExchangeParametersClassOnce sync.Once
)

func getTKTokenKeyExchangeParametersClass() _TKTokenKeyExchangeParametersClass {
	TKTokenKeyExchangeParametersClassOnce.Do(func() {
		TKTokenKeyExchangeParametersClass = _TKTokenKeyExchangeParametersClass{objc.GetClass("TKTokenKeyExchangeParameters")}
	})
	return TKTokenKeyExchangeParametersClass
}

type _TKTokenKeyExchangeParametersClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKTokenKeyExchangeParameters */
// An interface definition for the [TKTokenKeyExchangeParameters] class.
type ITKTokenKeyExchangeParameters interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TKTokenKeyExchangeParameters */
	// properties:
	RequestedSize() int
	SharedInfo() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKTokenKeyExchangeParameters */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKTokenKeyExchangeParameters */
// Alloc allocates a new instance without initialization.
func (tc _TKTokenKeyExchangeParametersClass) Alloc() TKTokenKeyExchangeParameters {
	rv := objc.Send[TKTokenKeyExchangeParameters](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKTokenKeyExchangeParametersClass) New() TKTokenKeyExchangeParameters {
	rv := objc.Send[TKTokenKeyExchangeParameters](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKTokenKeyExchangeParameters) Init() TKTokenKeyExchangeParameters {
	rv := objc.Send[TKTokenKeyExchangeParameters](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKTokenKeyExchangeParameters) Autorelease() TKTokenKeyExchangeParameters {
	rv := objc.Send[TKTokenKeyExchangeParameters](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenKeyExchangeParameters creates a new TKTokenKeyExchangeParameters instance.
func NewTKTokenKeyExchangeParameters() TKTokenKeyExchangeParameters {
	return getTKTokenKeyExchangeParametersClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKTokenKeyExchangeParameters */
// Parameters used to perform specific key exchange operations.


// Parameters used to perform specific key exchange operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeyExchangeParameters
type TKTokenKeyExchangeParameters struct {
	objectivec.Object
}

// TKTokenKeyExchangeParametersFrom constructs a [TKTokenKeyExchangeParameters] from an unsafe.Pointer.
//
// Parameters used to perform specific key exchange operations.
func TKTokenKeyExchangeParametersFrom(ptr unsafe.Pointer) TKTokenKeyExchangeParameters {
	return TKTokenKeyExchangeParameters{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKTokenKeyExchangeParameters *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKTokenKeyExchangeParameters */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKTokenKeyExchangeParameters */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKTokenKeyExchangeParameters */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKTokenKeyExchangeParameters */

// Returns the requested output size, in bytes, of key exchange result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeyExchangeParameters/requestedSize
func (t_ TKTokenKeyExchangeParameters) RequestedSize() int {
	rv := objc.Send[int](t_.ID, objc.Sel("requestedSize"))
	return rv
}/* debug [instance_properties/getter]: requestedSize */


// Returns shared information typically used during the key derivation (KDF) step of a key exchange algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenKeyExchangeParameters/sharedInfo
func (t_ TKTokenKeyExchangeParameters) SharedInfo() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("sharedInfo"))
	return rv
}/* debug [instance_properties/getter]: sharedInfo */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKTokenKeyExchangeParameters */



