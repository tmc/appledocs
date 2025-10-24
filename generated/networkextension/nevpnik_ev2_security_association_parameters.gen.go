// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEVPNIKEv2SecurityAssociationParameters */


/* debug [class_header]: Header for NEVPNIKEv2SecurityAssociationParameters */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEVPNIKEv2SecurityAssociationParameters */
// An interface definition for the [NEVPNIKEv2SecurityAssociationParameters] class.
type INEVPNIKEv2SecurityAssociationParameters interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEVPNIKEv2SecurityAssociationParameters */
	// properties:
	DiffieHellmanGroup() NEVPNIKEv2DiffieHellmanGroup
	SetDiffieHellmanGroup(value NEVPNIKEv2DiffieHellmanGroup)
	EncryptionAlgorithm() NEVPNIKEv2EncryptionAlgorithm
	SetEncryptionAlgorithm(value NEVPNIKEv2EncryptionAlgorithm)
	IntegrityAlgorithm() NEVPNIKEv2IntegrityAlgorithm
	SetIntegrityAlgorithm(value NEVPNIKEv2IntegrityAlgorithm)
	LifetimeMinutes() int32 /* not a class type */
	SetLifetimeMinutes(value int32 /* not a class type */)
	PostQuantumKeyExchangeMethods() []foundation.Number
	SetPostQuantumKeyExchangeMethods(value []foundation.Number)
	ChildSecurityAssociationParameters() INEVPNIKEv2SecurityAssociationParameters
	SetChildSecurityAssociationParameters(value INEVPNIKEv2SecurityAssociationParameters)
	IkeSecurityAssociationParameters() INEVPNIKEv2SecurityAssociationParameters
	SetIkeSecurityAssociationParameters(value INEVPNIKEv2SecurityAssociationParameters)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEVPNIKEv2SecurityAssociationParameters */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEVPNIKEv2SecurityAssociationParameters */
// Alloc allocates a new instance without initialization.
func (nc _NEVPNIKEv2SecurityAssociationParametersClass) Alloc() NEVPNIKEv2SecurityAssociationParameters {
	rv := objc.Send[NEVPNIKEv2SecurityAssociationParameters](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEVPNIKEv2SecurityAssociationParameters */
// Parameters for an IKEv2 Security Association.


// Parameters for an IKEv2 Security Association.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEVPNIKEv2SecurityAssociationParameters *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEVPNIKEv2SecurityAssociationParameters */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEVPNIKEv2SecurityAssociationParameters */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEVPNIKEv2SecurityAssociationParameters */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEVPNIKEv2SecurityAssociationParameters */

// The Diffie Hellman group used by the Security Association.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2SecurityAssociationParameters/diffieHellmanGroup
func (n_ NEVPNIKEv2SecurityAssociationParameters) DiffieHellmanGroup() NEVPNIKEv2DiffieHellmanGroup {
	rv := objc.Send[NEVPNIKEv2DiffieHellmanGroup](n_.ID, objc.Sel("diffieHellmanGroup"))
	return rv
}/* debug [instance_properties/getter]: diffieHellmanGroup */


// The Diffie Hellman group used by the Security Association.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2SecurityAssociationParameters/diffieHellmanGroup
func (n_ NEVPNIKEv2SecurityAssociationParameters) SetDiffieHellmanGroup(value NEVPNIKEv2DiffieHellmanGroup) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDiffieHellmanGroup:"), value)
}/* debug [instance_properties/setter]: diffieHellmanGroup */


// The algorithm used by the Security Association to encrypt and decrypt data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2SecurityAssociationParameters/encryptionAlgorithm
func (n_ NEVPNIKEv2SecurityAssociationParameters) EncryptionAlgorithm() NEVPNIKEv2EncryptionAlgorithm {
	rv := objc.Send[NEVPNIKEv2EncryptionAlgorithm](n_.ID, objc.Sel("encryptionAlgorithm"))
	return rv
}/* debug [instance_properties/getter]: encryptionAlgorithm */


// The algorithm used by the Security Association to encrypt and decrypt data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2SecurityAssociationParameters/encryptionAlgorithm
func (n_ NEVPNIKEv2SecurityAssociationParameters) SetEncryptionAlgorithm(value NEVPNIKEv2EncryptionAlgorithm) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setEncryptionAlgorithm:"), value)
}/* debug [instance_properties/setter]: encryptionAlgorithm */


// The algorithm used by the Security Association to verify the integrity of data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2SecurityAssociationParameters/integrityAlgorithm
func (n_ NEVPNIKEv2SecurityAssociationParameters) IntegrityAlgorithm() NEVPNIKEv2IntegrityAlgorithm {
	rv := objc.Send[NEVPNIKEv2IntegrityAlgorithm](n_.ID, objc.Sel("integrityAlgorithm"))
	return rv
}/* debug [instance_properties/getter]: integrityAlgorithm */


// The algorithm used by the Security Association to verify the integrity of data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2SecurityAssociationParameters/integrityAlgorithm
func (n_ NEVPNIKEv2SecurityAssociationParameters) SetIntegrityAlgorithm(value NEVPNIKEv2IntegrityAlgorithm) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIntegrityAlgorithm:"), value)
}/* debug [instance_properties/setter]: integrityAlgorithm */


// The duration of the lifetime of the Security Association, in minutes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2SecurityAssociationParameters/lifetimeMinutes
func (n_ NEVPNIKEv2SecurityAssociationParameters) LifetimeMinutes() int32 /* not a class type */ {
	rv := objc.Send[int32](n_.ID, objc.Sel("lifetimeMinutes"))
	return rv
}/* debug [instance_properties/getter]: lifetimeMinutes */


// The duration of the lifetime of the Security Association, in minutes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2SecurityAssociationParameters/lifetimeMinutes
func (n_ NEVPNIKEv2SecurityAssociationParameters) SetLifetimeMinutes(value int32 /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLifetimeMinutes:"), value)
}/* debug [instance_properties/setter]: lifetimeMinutes */


// A list of the quantum-secure key exchange methods the Security Association uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2SecurityAssociationParameters/postQuantumKeyExchangeMethods-56672
func (n_ NEVPNIKEv2SecurityAssociationParameters) PostQuantumKeyExchangeMethods() []foundation.Number {
	rv := objc.Send[[]foundation.Number](n_.ID, objc.Sel("postQuantumKeyExchangeMethods"))
	return rv
}/* debug [instance_properties/getter]: postQuantumKeyExchangeMethods */


// A list of the quantum-secure key exchange methods the Security Association uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2SecurityAssociationParameters/postQuantumKeyExchangeMethods-56672
func (n_ NEVPNIKEv2SecurityAssociationParameters) SetPostQuantumKeyExchangeMethods(value []foundation.Number) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](n_.ID, objc.Sel("setPostQuantumKeyExchangeMethods:"), nsArray)
}/* debug [instance_properties/setter]: postQuantumKeyExchangeMethods */


// An
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnprotocolikev2/childsecurityassociationparameters
func (n_ NEVPNIKEv2SecurityAssociationParameters) ChildSecurityAssociationParameters() INEVPNIKEv2SecurityAssociationParameters {
	rv := objc.Send[NEVPNIKEv2SecurityAssociationParameters](n_.ID, objc.Sel("childSecurityAssociationParameters"))
	return rv
}/* debug [instance_properties/getter]: childSecurityAssociationParameters */


// An
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnprotocolikev2/childsecurityassociationparameters
func (n_ NEVPNIKEv2SecurityAssociationParameters) SetChildSecurityAssociationParameters(value INEVPNIKEv2SecurityAssociationParameters) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setChildSecurityAssociationParameters:"), value)
}/* debug [instance_properties/setter]: childSecurityAssociationParameters */


// An
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnprotocolikev2/ikesecurityassociationparameters
func (n_ NEVPNIKEv2SecurityAssociationParameters) IkeSecurityAssociationParameters() INEVPNIKEv2SecurityAssociationParameters {
	rv := objc.Send[NEVPNIKEv2SecurityAssociationParameters](n_.ID, objc.Sel("ikeSecurityAssociationParameters"))
	return rv
}/* debug [instance_properties/getter]: ikeSecurityAssociationParameters */


// An
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nevpnprotocolikev2/ikesecurityassociationparameters
func (n_ NEVPNIKEv2SecurityAssociationParameters) SetIkeSecurityAssociationParameters(value INEVPNIKEv2SecurityAssociationParameters) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIkeSecurityAssociationParameters:"), value)
}/* debug [instance_properties/setter]: ikeSecurityAssociationParameters */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEVPNIKEv2SecurityAssociationParameters */



