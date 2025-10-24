// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class TKTokenDriverConfiguration */


/* debug [class_header]: Header for TKTokenDriverConfiguration */
// The class instance for the [TKTokenDriverConfiguration] class.
var (
	TKTokenDriverConfigurationClass     _TKTokenDriverConfigurationClass
	TKTokenDriverConfigurationClassOnce sync.Once
)

func getTKTokenDriverConfigurationClass() _TKTokenDriverConfigurationClass {
	TKTokenDriverConfigurationClassOnce.Do(func() {
		TKTokenDriverConfigurationClass = _TKTokenDriverConfigurationClass{objc.GetClass("TKTokenDriverConfiguration")}
	})
	return TKTokenDriverConfigurationClass
}

type _TKTokenDriverConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKTokenDriverConfiguration */
// An interface definition for the [TKTokenDriverConfiguration] class.
type ITKTokenDriverConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TKTokenDriverConfiguration */
	// properties:
	ClassID() TKTokenDriverClassID /* typedef */
	TokenConfigurations() foundation.IDictionary
	Delegate() objc.IObject /* cross-framework: TKTokenDriverDelegate */
	SetDelegate(value objc.IObject /* cross-framework: TKTokenDriverDelegate */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKTokenDriverConfiguration */
	// methods:
	AddTokenConfigurationForTokenInstanceID(instanceID TKTokenInstanceID /* typedef */) ITKTokenConfiguration
	RemoveTokenConfigurationForTokenInstanceID(instanceID TKTokenInstanceID /* typedef */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKTokenDriverConfiguration */
// Alloc allocates a new instance without initialization.
func (tc _TKTokenDriverConfigurationClass) Alloc() TKTokenDriverConfiguration {
	rv := objc.Send[TKTokenDriverConfiguration](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKTokenDriverConfigurationClass) New() TKTokenDriverConfiguration {
	rv := objc.Send[TKTokenDriverConfiguration](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKTokenDriverConfiguration) Init() TKTokenDriverConfiguration {
	rv := objc.Send[TKTokenDriverConfiguration](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKTokenDriverConfiguration) Autorelease() TKTokenDriverConfiguration {
	rv := objc.Send[TKTokenDriverConfiguration](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenDriverConfiguration creates a new TKTokenDriverConfiguration instance.
func NewTKTokenDriverConfiguration() TKTokenDriverConfiguration {
	return getTKTokenDriverConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKTokenDriverConfiguration */
// A configuration for one class of token.


// A configuration for one class of token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriver/Configuration
type TKTokenDriverConfiguration struct {
	objectivec.Object
}

// TKTokenDriverConfigurationFrom constructs a [TKTokenDriverConfiguration] from an unsafe.Pointer.
//
// A configuration for one class of token.
func TKTokenDriverConfigurationFrom(ptr unsafe.Pointer) TKTokenDriverConfiguration {
	return TKTokenDriverConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKTokenDriverConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKTokenDriverConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKTokenDriverConfiguration */

// A dictionary of token class configurations which the class identifier of the token driver keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriver/Configuration/driverConfigurations
func (tc _TKTokenDriverConfigurationClass) DriverConfigurations() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](objc.ID(tc.class), objc.Sel("driverConfigurations"))
	return rv
}/* debug [class_properties_class/property]: driverConfigurations */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKTokenDriverConfiguration */

// Creates a configuration object for a token with the token instance identifier you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriver/Configuration/addTokenConfiguration(for:)
func (t_ TKTokenDriverConfiguration) AddTokenConfigurationForTokenInstanceID(instanceID TKTokenInstanceID /* typedef */) ITKTokenConfiguration {
	rv := objc.Send[TKTokenConfiguration](t_.ID, objc.Sel("addTokenConfigurationForTokenInstanceID:"), instanceID)
	return rv
}/* debug [instance_methods/method]: AddTokenConfigurationForTokenInstanceID */


// Removes a configuration for a token with the token instance identifier you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriver/Configuration/removeTokenConfiguration(for:)
func (t_ TKTokenDriverConfiguration) RemoveTokenConfigurationForTokenInstanceID(instanceID TKTokenInstanceID /* typedef */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeTokenConfigurationForTokenInstanceID:"), instanceID)
}/* debug [instance_methods/method]: RemoveTokenConfigurationForTokenInstanceID */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKTokenDriverConfiguration */

// The class identifier of the token driver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriver/Configuration/classID
func (t_ TKTokenDriverConfiguration) ClassID() TKTokenDriverClassID /* typedef */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("classID"))
	return rv
}/* debug [instance_properties/getter]: classID */


// A dictionary of token class configurations which the class identifier of the token driver keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriver/Configuration/driverConfigurations
func (t_ TKTokenDriverConfiguration) DriverConfigurations() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("driverConfigurations"))
	return rv
}/* debug [instance_properties/getter]: driverConfigurations */


// A dictionary of all currently configured tokens for this token class, which the token instance identifier keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriver/Configuration/tokenConfigurations
func (t_ TKTokenDriverConfiguration) TokenConfigurations() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](t_.ID, objc.Sel("tokenConfigurations"))
	return rv
}/* debug [instance_properties/getter]: tokenConfigurations */


// The token driver delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktokendriver/delegate
func (t_ TKTokenDriverConfiguration) Delegate() objc.IObject /* cross-framework: TKTokenDriverDelegate */ {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The token driver delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktokendriver/delegate
func (t_ TKTokenDriverConfiguration) SetDelegate(value objc.IObject /* cross-framework: TKTokenDriverDelegate */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKTokenDriverConfiguration */



