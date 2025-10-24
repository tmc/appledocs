// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class TKTokenDriver */


/* debug [class_header]: Header for TKTokenDriver */
// The class instance for the [TKTokenDriver] class.
var (
	TKTokenDriverClass     _TKTokenDriverClass
	TKTokenDriverClassOnce sync.Once
)

func getTKTokenDriverClass() _TKTokenDriverClass {
	TKTokenDriverClassOnce.Do(func() {
		TKTokenDriverClass = _TKTokenDriverClass{objc.GetClass("TKTokenDriver")}
	})
	return TKTokenDriverClass
}

type _TKTokenDriverClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKTokenDriver */
// An interface definition for the [TKTokenDriver] class.
type ITKTokenDriver interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TKTokenDriver */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	ConfigurationData() foundation.Data
	SetConfigurationData(value foundation.Data)
	KeychainItems() ITKTokenKeychainItem
	SetKeychainItems(value ITKTokenKeychainItem)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKTokenDriver */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKTokenDriver */
// Alloc allocates a new instance without initialization.
func (tc _TKTokenDriverClass) Alloc() TKTokenDriver {
	rv := objc.Send[TKTokenDriver](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKTokenDriverClass) New() TKTokenDriver {
	rv := objc.Send[TKTokenDriver](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKTokenDriver) Init() TKTokenDriver {
	rv := objc.Send[TKTokenDriver](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKTokenDriver) Autorelease() TKTokenDriver {
	rv := objc.Send[TKTokenDriver](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenDriver creates a new TKTokenDriver instance.
func NewTKTokenDriver() TKTokenDriver {
	return getTKTokenDriverClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKTokenDriver */
// A base class for building token drivers.
//
// When using the class, implement the protocol with the method, which the system invokes when it requests the creation of a token instance. After you create the token driver, it can examine and to implement your desired functionality. An implementation can also access its associated token configuration using the property.


// A base class for building token drivers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriver
type TKTokenDriver struct {
	objectivec.Object
}

// TKTokenDriverFrom constructs a [TKTokenDriver] from an unsafe.Pointer.
//
// A base class for building token drivers.
func TKTokenDriverFrom(ptr unsafe.Pointer) TKTokenDriver {
	return TKTokenDriver{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKTokenDriver *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKTokenDriver */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKTokenDriver */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKTokenDriver */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKTokenDriver */

// The token driver delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriver/delegate
func (t_ TKTokenDriver) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The token driver delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKTokenDriver/delegate
func (t_ TKTokenDriver) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// Additional configuration information for the token instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktoken/configuration-swift.class/configurationdata
func (t_ TKTokenDriver) ConfigurationData() foundation.Data {
	rv := objc.Send[foundation.Data](t_.ID, objc.Sel("configurationData"))
	return rv
}/* debug [instance_properties/getter]: configurationData */


// Additional configuration information for the token instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktoken/configuration-swift.class/configurationdata
func (t_ TKTokenDriver) SetConfigurationData(value foundation.Data) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setConfigurationData:"), value)
}/* debug [instance_properties/setter]: configurationData */


// The keychain items associated with this token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktoken/configuration-swift.class/keychainitems
func (t_ TKTokenDriver) KeychainItems() ITKTokenKeychainItem {
	rv := objc.Send[TKTokenKeychainItem](t_.ID, objc.Sel("keychainItems"))
	return rv
}/* debug [instance_properties/getter]: keychainItems */


// The keychain items associated with this token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktoken/configuration-swift.class/keychainitems
func (t_ TKTokenDriver) SetKeychainItems(value ITKTokenKeychainItem) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setKeychainItems:"), value)
}/* debug [instance_properties/setter]: keychainItems */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKTokenDriver */



