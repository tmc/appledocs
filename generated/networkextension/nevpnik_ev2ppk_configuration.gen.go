// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [NEVPNIKEv2PPKConfiguration] class.
type INEVPNIKEv2PPKConfiguration interface {
	objectivec.IObject
	

	// properties:
	Identifier() foundation.foundation.INSString
	IsMandatory() bool
	SetIsMandatory(value bool)
	KeychainReference() foundation.foundation.INSData
	AllowPostQuantumKeyExchangeFallback() bool
	SetAllowPostQuantumKeyExchangeFallback(value bool)
	PpkConfiguration() INEVPNIKEv2PPKConfiguration
	SetPpkConfiguration(value INEVPNIKEv2PPKConfiguration)


	

	// methods:


}





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






// Initializes a quantum-secure pre-shared key (PPK) configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2PPKConfiguration/init(identifier:keychainReference:)
func NewNEVPNIKEv2PPKConfigurationWithIdentifierKeychainReference(identifier foundation.foundation.INSString, keychainReference foundation.foundation.INSData) NEVPNIKEv2PPKConfiguration {
	instance := getNEVPNIKEv2PPKConfigurationClass().Alloc()
	rv := objc.Send[NEVPNIKEv2PPKConfiguration](instance.ID, objc.Sel("initWithIdentifier:keychainReference:"), identifier, keychainReference)
	rv.Autorelease()
	return rv
}






















// The identifier for the PPK.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2PPKConfiguration/identifier
func (n_ NEVPNIKEv2PPKConfiguration) Identifier() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("identifier"))
	return rv
}


// A Boolean value that indicates whether it’s mandatory for the VPN server to use this PPK.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2PPKConfiguration/isMandatory
func (n_ NEVPNIKEv2PPKConfiguration) IsMandatory() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isMandatory"))
	return rv
}


// A Boolean value that indicates whether it’s mandatory for the VPN server to use this PPK.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2PPKConfiguration/isMandatory
func (n_ NEVPNIKEv2PPKConfiguration) SetIsMandatory(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsMandatory:"), value)
}


// A persistent reference to the key in the keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2PPKConfiguration/keychainReference
func (n_ NEVPNIKEv2PPKConfiguration) KeychainReference() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](n_.ID, objc.Sel("keychainReference"))
	return rv
}


// A Boolean value that indicates whether servers that don’t support post-quantum key exchanges can skip them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnprotocolikev2/allowpostquantumkeyexchangefallback
func (n_ NEVPNIKEv2PPKConfiguration) AllowPostQuantumKeyExchangeFallback() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("allowPostQuantumKeyExchangeFallback"))
	return rv
}


// A Boolean value that indicates whether servers that don’t support post-quantum key exchanges can skip them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnprotocolikev2/allowpostquantumkeyexchangefallback
func (n_ NEVPNIKEv2PPKConfiguration) SetAllowPostQuantumKeyExchangeFallback(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAllowPostQuantumKeyExchangeFallback:"), value)
}


// The configuration for a post-quantum pre-shared key (PPK).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnprotocolikev2/ppkconfiguration
func (n_ NEVPNIKEv2PPKConfiguration) PpkConfiguration() INEVPNIKEv2PPKConfiguration {
	rv := objc.Send[NEVPNIKEv2PPKConfiguration](n_.ID, objc.Sel("ppkConfiguration"))
	return rv
}


// The configuration for a post-quantum pre-shared key (PPK).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnprotocolikev2/ppkconfiguration
func (n_ NEVPNIKEv2PPKConfiguration) SetPpkConfiguration(value INEVPNIKEv2PPKConfiguration) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPpkConfiguration:"), value)
}







