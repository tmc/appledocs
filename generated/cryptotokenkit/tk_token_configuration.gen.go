// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class TKTokenConfiguration */


/* debug [class_header]: Header for TKTokenConfiguration */
// The class instance for the [TKTokenConfiguration] class.
var (
	TKTokenConfigurationClass     _TKTokenConfigurationClass
	TKTokenConfigurationClassOnce sync.Once
)

func getTKTokenConfigurationClass() _TKTokenConfigurationClass {
	TKTokenConfigurationClassOnce.Do(func() {
		TKTokenConfigurationClass = _TKTokenConfigurationClass{objc.GetClass("TKTokenConfiguration")}
	})
	return TKTokenConfigurationClass
}

type _TKTokenConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKTokenConfiguration */
// An interface definition for the [TKTokenConfiguration] class.
type ITKTokenConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TKTokenConfiguration */
	// properties:
	ConfigurationData() objc.IObject /* cross-framework: NSData */
	SetConfigurationData(value objc.IObject /* cross-framework: NSData */)
	InstanceID() TKTokenInstanceID /* typedef */
	KeychainItems() []TKTokenKeychainItem
	SetKeychainItems(value []TKTokenKeychainItem)
	Configuration() ITKTokenConfiguration
	SetConfiguration(value ITKTokenConfiguration)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKTokenConfiguration */
	// methods:
	CertificateForObjectIDError(objectID TKTokenObjectID /* typedef */, error_ unsafe.Pointer) ITKTokenKeychainCertificate
	KeyForObjectIDError(objectID TKTokenObjectID /* typedef */, error_ unsafe.Pointer) ITKTokenKeychainKey
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKTokenConfiguration */
// Alloc allocates a new instance without initialization.
func (tc _TKTokenConfigurationClass) Alloc() TKTokenConfiguration {
	rv := objc.Send[TKTokenConfiguration](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKTokenConfigurationClass) New() TKTokenConfiguration {
	rv := objc.Send[TKTokenConfiguration](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKTokenConfiguration) Init() TKTokenConfiguration {
	rv := objc.Send[TKTokenConfiguration](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKTokenConfiguration) Autorelease() TKTokenConfiguration {
	rv := objc.Send[TKTokenConfiguration](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKTokenConfiguration creates a new TKTokenConfiguration instance.
func NewTKTokenConfiguration() TKTokenConfiguration {
	return getTKTokenConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKTokenConfiguration */
// A token’s configuration.
//
// When you introduce a new into the system, it can inform the system about its identities, consisting of both private keys and certificates, which the property provides. Use the property to set additional configuration data. You configure always-available tokens on a per-user basis. Although the token driver and the app hosting the token extension are shared across the system, the configuration for a token is stored individually for each user.


// A token’s configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/Configuration-swift.class
type TKTokenConfiguration struct {
	objectivec.Object
}

// TKTokenConfigurationFrom constructs a [TKTokenConfiguration] from an unsafe.Pointer.
//
// A token’s configuration.
func TKTokenConfigurationFrom(ptr unsafe.Pointer) TKTokenConfiguration {
	return TKTokenConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKTokenConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKTokenConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKTokenConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKTokenConfiguration */

// Returns a certificate from the keychain with the object identifier you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/Configuration-swift.class/certificate(for:)
func (t_ TKTokenConfiguration) CertificateForObjectIDError(objectID TKTokenObjectID /* typedef */, error_ unsafe.Pointer) ITKTokenKeychainCertificate {
	rv := objc.Send[TKTokenKeychainCertificate](t_.ID, objc.Sel("certificateForObjectID:error:"), objectID, error_)
	return rv
}/* debug [instance_methods/method]: CertificateForObjectIDError */


// Returns a key from the keychain with the object identifier you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/Configuration-swift.class/key(for:)
func (t_ TKTokenConfiguration) KeyForObjectIDError(objectID TKTokenObjectID /* typedef */, error_ unsafe.Pointer) ITKTokenKeychainKey {
	rv := objc.Send[TKTokenKeychainKey](t_.ID, objc.Sel("keyForObjectID:error:"), objectID, error_)
	return rv
}/* debug [instance_methods/method]: KeyForObjectIDError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKTokenConfiguration */

// Additional configuration information for the token instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/Configuration-swift.class/configurationData
func (t_ TKTokenConfiguration) ConfigurationData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("configurationData"))
	return rv
}/* debug [instance_properties/getter]: configurationData */


// Additional configuration information for the token instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/Configuration-swift.class/configurationData
func (t_ TKTokenConfiguration) SetConfigurationData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setConfigurationData:"), value)
}/* debug [instance_properties/setter]: configurationData */


// The unique, persistent identifier of this token that the token implementation creates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/Configuration-swift.class/instanceID
func (t_ TKTokenConfiguration) InstanceID() TKTokenInstanceID /* typedef */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("instanceID"))
	return rv
}/* debug [instance_properties/getter]: instanceID */


// The keychain items associated with this token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/Configuration-swift.class/keychainItems
func (t_ TKTokenConfiguration) KeychainItems() []TKTokenKeychainItem {
	rv := objc.Send[[]TKTokenKeychainItem](t_.ID, objc.Sel("keychainItems"))
	return rv
}/* debug [instance_properties/getter]: keychainItems */


// The keychain items associated with this token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKToken/Configuration-swift.class/keychainItems
func (t_ TKTokenConfiguration) SetKeychainItems(value []TKTokenKeychainItem) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](t_.ID, objc.Sel("setKeychainItems:"), nsArray)
}/* debug [instance_properties/setter]: keychainItems */


// The current configuration for a token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktoken/configuration-swift.property
func (t_ TKTokenConfiguration) Configuration() ITKTokenConfiguration {
	rv := objc.Send[TKTokenConfiguration](t_.ID, objc.Sel("configuration"))
	return rv
}/* debug [instance_properties/getter]: configuration */


// The current configuration for a token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cryptotokenkit/tktoken/configuration-swift.property
func (t_ TKTokenConfiguration) SetConfiguration(value ITKTokenConfiguration) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setConfiguration:"), value)
}/* debug [instance_properties/setter]: configuration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKTokenConfiguration */



