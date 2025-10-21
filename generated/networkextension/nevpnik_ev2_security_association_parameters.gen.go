// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEVPNIKEv2SecurityAssociationParameters] class.
var (
	NEVPNIKEv2SecurityAssociationParametersClass     _NEVPNIKEv2SecurityAssociationParametersClass
	NEVPNIKEv2SecurityAssociationParametersClassOnce sync.Once
)

func getNEVPNIKEv2SecurityAssociationParametersClass() _NEVPNIKEv2SecurityAssociationParametersClass {
	NEVPNIKEv2SecurityAssociationParametersClassOnce.Do(func() {
		NEVPNIKEv2SecurityAssociationParametersClass = _NEVPNIKEv2SecurityAssociationParametersClass{objc.GetClass("NEVPNIKEv2SecurityAssociationParameters")}
	})
	return NEVPNIKEv2SecurityAssociationParametersClass
}

type _NEVPNIKEv2SecurityAssociationParametersClass struct {
	class objc.Class
}

// An interface definition for the [NEVPNIKEv2SecurityAssociationParameters] class.
type INEVPNIKEv2SecurityAssociationParameters interface {
	objectivec.IObject
}

// Parameters for an IKEv2 Security Association.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2SecurityAssociationParameters
type NEVPNIKEv2SecurityAssociationParameters struct {
	objectivec.Object
}

// NEVPNIKEv2SecurityAssociationParametersFrom constructs a [NEVPNIKEv2SecurityAssociationParameters] from an unsafe.Pointer.
//
// Parameters for an IKEv2 Security Association.
func NEVPNIKEv2SecurityAssociationParametersFrom(ptr unsafe.Pointer) NEVPNIKEv2SecurityAssociationParameters {
	return NEVPNIKEv2SecurityAssociationParameters{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEVPNIKEv2SecurityAssociationParametersClass) Alloc() NEVPNIKEv2SecurityAssociationParameters {
	rv := objc.Send[NEVPNIKEv2SecurityAssociationParameters](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEVPNIKEv2SecurityAssociationParametersClass) New() NEVPNIKEv2SecurityAssociationParameters {
	rv := objc.Send[NEVPNIKEv2SecurityAssociationParameters](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEVPNIKEv2SecurityAssociationParameters) Init() NEVPNIKEv2SecurityAssociationParameters {
	rv := objc.Send[NEVPNIKEv2SecurityAssociationParameters](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEVPNIKEv2SecurityAssociationParameters) Autorelease() NEVPNIKEv2SecurityAssociationParameters {
	rv := objc.Send[NEVPNIKEv2SecurityAssociationParameters](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEVPNIKEv2SecurityAssociationParameters creates a new NEVPNIKEv2SecurityAssociationParameters instance.
func NewNEVPNIKEv2SecurityAssociationParameters() NEVPNIKEv2SecurityAssociationParameters {
	return getNEVPNIKEv2SecurityAssociationParametersClass().New()
}


// The Diffie Hellman group used by the Security Association.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnikev2securityassociationparameters/diffiehellmangroup
func (n_ NEVPNIKEv2SecurityAssociationParameters) DiffieHellmanGroup() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("diffieHellmanGroup"))
	return rv
}


// SetDiffieHellmanGroup sets the value of the diffieHellmanGroup property.
// The Diffie Hellman group used by the Security Association.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnikev2securityassociationparameters/diffiehellmangroup
func (n_ NEVPNIKEv2SecurityAssociationParameters) SetDiffieHellmanGroup(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDiffieHellmanGroup:"), value)
}

// The algorithm used by the Security Association to encrypt and decrypt data.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnikev2securityassociationparameters/encryptionalgorithm
func (n_ NEVPNIKEv2SecurityAssociationParameters) EncryptionAlgorithm() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("encryptionAlgorithm"))
	return rv
}


// SetEncryptionAlgorithm sets the value of the encryptionAlgorithm property.
// The algorithm used by the Security Association to encrypt and decrypt data.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnikev2securityassociationparameters/encryptionalgorithm
func (n_ NEVPNIKEv2SecurityAssociationParameters) SetEncryptionAlgorithm(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setEncryptionAlgorithm:"), value)
}

// The algorithm used by the Security Association to verify the integrity of data.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnikev2securityassociationparameters/integrityalgorithm
func (n_ NEVPNIKEv2SecurityAssociationParameters) IntegrityAlgorithm() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("integrityAlgorithm"))
	return rv
}


// SetIntegrityAlgorithm sets the value of the integrityAlgorithm property.
// The algorithm used by the Security Association to verify the integrity of data.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnikev2securityassociationparameters/integrityalgorithm
func (n_ NEVPNIKEv2SecurityAssociationParameters) SetIntegrityAlgorithm(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIntegrityAlgorithm:"), value)
}

// The duration of the lifetime of the Security Association, in minutes.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnikev2securityassociationparameters/lifetimeminutes
func (n_ NEVPNIKEv2SecurityAssociationParameters) LifetimeMinutes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("lifetimeMinutes"))
	return rv
}


// SetLifetimeMinutes sets the value of the lifetimeMinutes property.
// The duration of the lifetime of the Security Association, in minutes.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnikev2securityassociationparameters/lifetimeminutes
func (n_ NEVPNIKEv2SecurityAssociationParameters) SetLifetimeMinutes(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLifetimeMinutes:"), value)
}

// A list of the quantum-secure key exchange methods the Security Association uses.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnikev2securityassociationparameters/postquantumkeyexchangemethods-3173s
func (n_ NEVPNIKEv2SecurityAssociationParameters) PostQuantumKeyExchangeMethods() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("postQuantumKeyExchangeMethods"))
	return rv
}


// SetPostQuantumKeyExchangeMethods sets the value of the postQuantumKeyExchangeMethods property.
// A list of the quantum-secure key exchange methods the Security Association uses.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnikev2securityassociationparameters/postquantumkeyexchangemethods-3173s
func (n_ NEVPNIKEv2SecurityAssociationParameters) SetPostQuantumKeyExchangeMethods(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPostQuantumKeyExchangeMethods:"), value)
}

// An
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnprotocolikev2/childsecurityassociationparameters
func (n_ NEVPNIKEv2SecurityAssociationParameters) ChildSecurityAssociationParameters() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("childSecurityAssociationParameters"))
	return rv
}


// SetChildSecurityAssociationParameters sets the value of the childSecurityAssociationParameters property.
// An

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnprotocolikev2/childsecurityassociationparameters
func (n_ NEVPNIKEv2SecurityAssociationParameters) SetChildSecurityAssociationParameters(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setChildSecurityAssociationParameters:"), value)
}

// An
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnprotocolikev2/ikesecurityassociationparameters
func (n_ NEVPNIKEv2SecurityAssociationParameters) IkeSecurityAssociationParameters() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("ikeSecurityAssociationParameters"))
	return rv
}


// SetIkeSecurityAssociationParameters sets the value of the ikeSecurityAssociationParameters property.
// An

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnprotocolikev2/ikesecurityassociationparameters
func (n_ NEVPNIKEv2SecurityAssociationParameters) SetIkeSecurityAssociationParameters(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIkeSecurityAssociationParameters:"), value)
}



