// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	Identifier() string
	SetIdentifier(value string)
	IsMandatory() bool
	SetIsMandatory(value bool)
	KeychainReference() foundation.Data
	SetKeychainReference(value foundation.IData)
	AllowPostQuantumKeyExchangeFallback() bool
	SetAllowPostQuantumKeyExchangeFallback(value bool)
	PpkConfiguration() NEVPNIKEv2PPKConfiguration
	SetPpkConfiguration(value INEVPNIKEv2PPKConfiguration)
}

// A class that manages parameters of a post-quantum pre-shared key (PPK).
//
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

// Alloc allocates a new instance without initialization.
func (nc _NEVPNIKEv2PPKConfigurationClass) Alloc() NEVPNIKEv2PPKConfiguration {
	rv := objc.Send[NEVPNIKEv2PPKConfiguration](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The identifier for the PPK.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnikev2ppkconfiguration/identifier
func (n_ NEVPNIKEv2PPKConfiguration) Identifier() string {
	rv := objc.Send[string](n_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// The identifier for the PPK.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnikev2ppkconfiguration/identifier
func (n_ NEVPNIKEv2PPKConfiguration) SetIdentifier(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}

// A Boolean value that indicates whether it’s mandatory for the VPN server to use this PPK.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnikev2ppkconfiguration/ismandatory
func (n_ NEVPNIKEv2PPKConfiguration) IsMandatory() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isMandatory"))
	return rv
}


// SetIsMandatory sets the value of the isMandatory property.
// A Boolean value that indicates whether it’s mandatory for the VPN server to use this PPK.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnikev2ppkconfiguration/ismandatory
func (n_ NEVPNIKEv2PPKConfiguration) SetIsMandatory(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsMandatory:"), value)
}

// A persistent reference to the key in the keychain.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnikev2ppkconfiguration/keychainreference
func (n_ NEVPNIKEv2PPKConfiguration) KeychainReference() foundation.Data {
	rv := objc.Send[foundation.Data](n_.ID, objc.Sel("keychainReference"))
	return rv
}


// SetKeychainReference sets the value of the keychainReference property.
// A persistent reference to the key in the keychain.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnikev2ppkconfiguration/keychainreference
func (n_ NEVPNIKEv2PPKConfiguration) SetKeychainReference(value foundation.IData) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setKeychainReference:"), value)
}

// A Boolean value that indicates whether servers that don’t support post-quantum key exchanges can skip them.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnprotocolikev2/allowpostquantumkeyexchangefallback
func (n_ NEVPNIKEv2PPKConfiguration) AllowPostQuantumKeyExchangeFallback() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("allowPostQuantumKeyExchangeFallback"))
	return rv
}


// SetAllowPostQuantumKeyExchangeFallback sets the value of the allowPostQuantumKeyExchangeFallback property.
// A Boolean value that indicates whether servers that don’t support post-quantum key exchanges can skip them.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnprotocolikev2/allowpostquantumkeyexchangefallback
func (n_ NEVPNIKEv2PPKConfiguration) SetAllowPostQuantumKeyExchangeFallback(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAllowPostQuantumKeyExchangeFallback:"), value)
}

// The configuration for a post-quantum pre-shared key (PPK).
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnprotocolikev2/ppkconfiguration
func (n_ NEVPNIKEv2PPKConfiguration) PpkConfiguration() NEVPNIKEv2PPKConfiguration {
	rv := objc.Send[NEVPNIKEv2PPKConfiguration](n_.ID, objc.Sel("ppkConfiguration"))
	return rv
}


// SetPpkConfiguration sets the value of the ppkConfiguration property.
// The configuration for a post-quantum pre-shared key (PPK).

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnprotocolikev2/ppkconfiguration
func (n_ NEVPNIKEv2PPKConfiguration) SetPpkConfiguration(value INEVPNIKEv2PPKConfiguration) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPpkConfiguration:"), value)
}



