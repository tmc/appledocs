// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class LADomainStateBiometry */


/* debug [class_header]: Header for LADomainStateBiometry */
// The class instance for the [DomainStateBiometry] class.
var (
	DomainStateBiometryClass     _DomainStateBiometryClass
	DomainStateBiometryClassOnce sync.Once
)

func getDomainStateBiometryClass() _DomainStateBiometryClass {
	DomainStateBiometryClassOnce.Do(func() {
		DomainStateBiometryClass = _DomainStateBiometryClass{objc.GetClass("LADomainStateBiometry")}
	})
	return DomainStateBiometryClass
}

type _DomainStateBiometryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DomainStateBiometry */
// An interface definition for the [DomainStateBiometry] class.
type IDomainStateBiometry interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DomainStateBiometry */
	// properties:
	BiometryType() BiometryType
	StateHash() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DomainStateBiometry */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DomainStateBiometry */
// Alloc allocates a new instance without initialization.
func (dc _DomainStateBiometryClass) Alloc() DomainStateBiometry {
	rv := objc.Send[DomainStateBiometry](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DomainStateBiometryClass) New() DomainStateBiometry {
	rv := objc.Send[DomainStateBiometry](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DomainStateBiometry) Init() DomainStateBiometry {
	rv := objc.Send[DomainStateBiometry](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DomainStateBiometry) Autorelease() DomainStateBiometry {
	rv := objc.Send[DomainStateBiometry](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDomainStateBiometry creates a new DomainStateBiometry instance.
func NewDomainStateBiometry() DomainStateBiometry {
	return getDomainStateBiometryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DomainStateBiometry */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LADomainStateBiometry
type DomainStateBiometry struct {
	objectivec.Object
}

// DomainStateBiometryFrom constructs a [DomainStateBiometry] from an unsafe.Pointer.
func DomainStateBiometryFrom(ptr unsafe.Pointer) DomainStateBiometry {
	return DomainStateBiometry{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DomainStateBiometry *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DomainStateBiometry */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DomainStateBiometry */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DomainStateBiometry */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DomainStateBiometry */

// Indicates biometry type available on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LADomainStateBiometry/biometryType
func (d_ DomainStateBiometry) BiometryType() BiometryType {
	rv := objc.Send[BiometryType](d_.ID, objc.Sel("biometryType"))
	return rv
}/* debug [instance_properties/getter]: biometryType */


// Contains state hash data for the available biometry type. Returns if no biometry entities are enrolled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LADomainStateBiometry/stateHash
func (d_ DomainStateBiometry) StateHash() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](d_.ID, objc.Sel("stateHash"))
	return rv
}/* debug [instance_properties/getter]: stateHash */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class LADomainStateBiometry */



