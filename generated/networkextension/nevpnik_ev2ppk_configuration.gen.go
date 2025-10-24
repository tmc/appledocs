// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEVPNIKEv2PPKConfiguration */


/* debug [class_header]: Header for NEVPNIKEv2PPKConfiguration */
// The class instance for the [NEVPNIKEv2PPKConfiguration] class.
var (
	NEVPNIKEv2PPKConfigurationClass     _NEVPNIKEv2PPKConfigurationClass
	NEVPNIKEv2PPKConfigurationClassOnce sync.Once
)

func getNEVPNIKEv2PPKConfigurationClass() _NEVPNIKEv2PPKConfigurationClass {
	NEVPNIKEv2PPKConfigurationClassOnce.Do(func() {
		NEVPNIKEv2PPKConfigurationClass = _NEVPNIKEv2PPKConfigurationClass{objc.GetClass("NEVPNIKEv2PPKConfiguration")}
	})
	return NEVPNIKEv2PPKConfigurationClass
}

type _NEVPNIKEv2PPKConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEVPNIKEv2PPKConfiguration */
// An interface definition for the [NEVPNIKEv2PPKConfiguration] class.
type INEVPNIKEv2PPKConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEVPNIKEv2PPKConfiguration */
	// properties:
	Identifier() objc.IObject /* cross-framework: NSString */
	IsMandatory() bool
	SetIsMandatory(value bool)
	KeychainReference() objc.IObject /* cross-framework: NSData */
	AllowPostQuantumKeyExchangeFallback() bool
	SetAllowPostQuantumKeyExchangeFallback(value bool)
	PpkConfiguration() INEVPNIKEv2PPKConfiguration
	SetPpkConfiguration(value INEVPNIKEv2PPKConfiguration)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEVPNIKEv2PPKConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEVPNIKEv2PPKConfiguration */
// Alloc allocates a new instance without initialization.
func (nc _NEVPNIKEv2PPKConfigurationClass) Alloc() NEVPNIKEv2PPKConfiguration {
	rv := objc.Send[NEVPNIKEv2PPKConfiguration](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEVPNIKEv2PPKConfigurationClass) New() NEVPNIKEv2PPKConfiguration {
	rv := objc.Send[NEVPNIKEv2PPKConfiguration](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEVPNIKEv2PPKConfiguration) Init() NEVPNIKEv2PPKConfiguration {
	rv := objc.Send[NEVPNIKEv2PPKConfiguration](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEVPNIKEv2PPKConfiguration) Autorelease() NEVPNIKEv2PPKConfiguration {
	rv := objc.Send[NEVPNIKEv2PPKConfiguration](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEVPNIKEv2PPKConfiguration creates a new NEVPNIKEv2PPKConfiguration instance.
func NewNEVPNIKEv2PPKConfiguration() NEVPNIKEv2PPKConfiguration {
	return getNEVPNIKEv2PPKConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEVPNIKEv2PPKConfiguration */
// A class that manages parameters of a post-quantum pre-shared key (PPK).
//
// Instances of this class are thread safe. The class conforms to RFC 8784.


// A class that manages parameters of a post-quantum pre-shared key (PPK).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2PPKConfiguration
type NEVPNIKEv2PPKConfiguration struct {
	objectivec.Object
}

// NEVPNIKEv2PPKConfigurationFrom constructs a [NEVPNIKEv2PPKConfiguration] from an unsafe.Pointer.
//
// A class that manages parameters of a post-quantum pre-shared key (PPK).
func NEVPNIKEv2PPKConfigurationFrom(ptr unsafe.Pointer) NEVPNIKEv2PPKConfiguration {
	return NEVPNIKEv2PPKConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEVPNIKEv2PPKConfiguration */

// Initializes a quantum-secure pre-shared key (PPK) configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2PPKConfiguration/init(identifier:keychainReference:)
func NewNEVPNIKEv2PPKConfigurationWithIdentifierKeychainReference(identifier objc.IObject /* cross-framework: NSString */, keychainReference objc.IObject /* cross-framework: NSData */) NEVPNIKEv2PPKConfiguration {
	instance := getNEVPNIKEv2PPKConfigurationClass().Alloc()
	rv := objc.Send[NEVPNIKEv2PPKConfiguration](instance.ID, objc.Sel("initWithIdentifier:keychainReference:"), identifier, keychainReference)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNEVPNIKEv2PPKConfigurationWithIdentifierKeychainReference */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEVPNIKEv2PPKConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEVPNIKEv2PPKConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEVPNIKEv2PPKConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEVPNIKEv2PPKConfiguration */

// The identifier for the PPK.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2PPKConfiguration/identifier
func (n_ NEVPNIKEv2PPKConfiguration) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// A Boolean value that indicates whether it’s mandatory for the VPN server to use this PPK.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2PPKConfiguration/isMandatory
func (n_ NEVPNIKEv2PPKConfiguration) IsMandatory() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isMandatory"))
	return rv
}/* debug [instance_properties/getter]: isMandatory */


// A Boolean value that indicates whether it’s mandatory for the VPN server to use this PPK.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2PPKConfiguration/isMandatory
func (n_ NEVPNIKEv2PPKConfiguration) SetIsMandatory(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsMandatory:"), value)
}/* debug [instance_properties/setter]: isMandatory */


// A persistent reference to the key in the keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2PPKConfiguration/keychainReference
func (n_ NEVPNIKEv2PPKConfiguration) KeychainReference() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("keychainReference"))
	return rv
}/* debug [instance_properties/getter]: keychainReference */


// A Boolean value that indicates whether servers that don’t support post-quantum key exchanges can skip them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnprotocolikev2/allowpostquantumkeyexchangefallback
func (n_ NEVPNIKEv2PPKConfiguration) AllowPostQuantumKeyExchangeFallback() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("allowPostQuantumKeyExchangeFallback"))
	return rv
}/* debug [instance_properties/getter]: allowPostQuantumKeyExchangeFallback */


// A Boolean value that indicates whether servers that don’t support post-quantum key exchanges can skip them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnprotocolikev2/allowpostquantumkeyexchangefallback
func (n_ NEVPNIKEv2PPKConfiguration) SetAllowPostQuantumKeyExchangeFallback(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAllowPostQuantumKeyExchangeFallback:"), value)
}/* debug [instance_properties/setter]: allowPostQuantumKeyExchangeFallback */


// The configuration for a post-quantum pre-shared key (PPK).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnprotocolikev2/ppkconfiguration
func (n_ NEVPNIKEv2PPKConfiguration) PpkConfiguration() INEVPNIKEv2PPKConfiguration {
	rv := objc.Send[NEVPNIKEv2PPKConfiguration](n_.ID, objc.Sel("ppkConfiguration"))
	return rv
}/* debug [instance_properties/getter]: ppkConfiguration */


// The configuration for a post-quantum pre-shared key (PPK).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnprotocolikev2/ppkconfiguration
func (n_ NEVPNIKEv2PPKConfiguration) SetPpkConfiguration(value INEVPNIKEv2PPKConfiguration) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPpkConfiguration:"), value)
}/* debug [instance_properties/setter]: ppkConfiguration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEVPNIKEv2PPKConfiguration */


