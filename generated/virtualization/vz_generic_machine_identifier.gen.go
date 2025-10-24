// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZGenericMachineIdentifier */


/* debug [class_header]: Header for VZGenericMachineIdentifier */
// The class instance for the [VZGenericMachineIdentifier] class.
var (
	VZGenericMachineIdentifierClass     _VZGenericMachineIdentifierClass
	VZGenericMachineIdentifierClassOnce sync.Once
)

func getVZGenericMachineIdentifierClass() _VZGenericMachineIdentifierClass {
	VZGenericMachineIdentifierClassOnce.Do(func() {
		VZGenericMachineIdentifierClass = _VZGenericMachineIdentifierClass{objc.GetClass("VZGenericMachineIdentifier")}
	})
	return VZGenericMachineIdentifierClass
}

type _VZGenericMachineIdentifierClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZGenericMachineIdentifier */
// An interface definition for the [VZGenericMachineIdentifier] class.
type IVZGenericMachineIdentifier interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZGenericMachineIdentifier */
	// properties:
	DataRepresentation() objc.IObject /* cross-framework: NSData */
	IsNestedVirtualizationEnabled() bool
	SetIsNestedVirtualizationEnabled(value bool)
	MachineIdentifier() IVZGenericMachineIdentifier
	SetMachineIdentifier(value IVZGenericMachineIdentifier)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZGenericMachineIdentifier */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZGenericMachineIdentifier */
// Alloc allocates a new instance without initialization.
func (vc _VZGenericMachineIdentifierClass) Alloc() VZGenericMachineIdentifier {
	rv := objc.Send[VZGenericMachineIdentifier](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZGenericMachineIdentifierClass) New() VZGenericMachineIdentifier {
	rv := objc.Send[VZGenericMachineIdentifier](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZGenericMachineIdentifier) Init() VZGenericMachineIdentifier {
	rv := objc.Send[VZGenericMachineIdentifier](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZGenericMachineIdentifier) Autorelease() VZGenericMachineIdentifier {
	rv := objc.Send[VZGenericMachineIdentifier](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGenericMachineIdentifier creates a new VZGenericMachineIdentifier instance.
func NewVZGenericMachineIdentifier() VZGenericMachineIdentifier {
	return getVZGenericMachineIdentifierClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZGenericMachineIdentifier */
// An object that represents a unique identifier for a virtual machine.
//
// Use the data representation in to save the VM’s identifier. To restore a previously saved identifier use .


// An object that represents a unique identifier for a virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGenericMachineIdentifier
type VZGenericMachineIdentifier struct {
	objectivec.Object
}

// VZGenericMachineIdentifierFrom constructs a [VZGenericMachineIdentifier] from an unsafe.Pointer.
//
// An object that represents a unique identifier for a virtual machine.
func VZGenericMachineIdentifierFrom(ptr unsafe.Pointer) VZGenericMachineIdentifier {
	return VZGenericMachineIdentifier{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZGenericMachineIdentifier */

// Creates a new unique identifier for a VM with the provided data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGenericMachineIdentifier/init(dataRepresentation:)
func NewVZGenericMachineIdentifierWithDataRepresentation(dataRepresentation objc.IObject /* cross-framework: NSData */) VZGenericMachineIdentifier {
	instance := getVZGenericMachineIdentifierClass().Alloc()
	rv := objc.Send[VZGenericMachineIdentifier](instance.ID, objc.Sel("initWithDataRepresentation:"), dataRepresentation)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVZGenericMachineIdentifierWithDataRepresentation */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZGenericMachineIdentifier */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZGenericMachineIdentifier */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZGenericMachineIdentifier */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZGenericMachineIdentifier */

// An opaque data representation of the VM’s identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZGenericMachineIdentifier/dataRepresentation
func (v_ VZGenericMachineIdentifier) DataRepresentation() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](v_.ID, objc.Sel("dataRepresentation"))
	return rv
}/* debug [instance_properties/getter]: dataRepresentation */


// A Boolean value that indicates whether nested virtualization is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzgenericplatformconfiguration/isnestedvirtualizationenabled
func (v_ VZGenericMachineIdentifier) IsNestedVirtualizationEnabled() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isNestedVirtualizationEnabled"))
	return rv
}/* debug [instance_properties/getter]: isNestedVirtualizationEnabled */


// A Boolean value that indicates whether nested virtualization is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzgenericplatformconfiguration/isnestedvirtualizationenabled
func (v_ VZGenericMachineIdentifier) SetIsNestedVirtualizationEnabled(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsNestedVirtualizationEnabled:"), value)
}/* debug [instance_properties/setter]: isNestedVirtualizationEnabled */


// A value that represents a unique identifier for the virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzgenericplatformconfiguration/machineidentifier
func (v_ VZGenericMachineIdentifier) MachineIdentifier() IVZGenericMachineIdentifier {
	rv := objc.Send[VZGenericMachineIdentifier](v_.ID, objc.Sel("machineIdentifier"))
	return rv
}/* debug [instance_properties/getter]: machineIdentifier */


// A value that represents a unique identifier for the virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzgenericplatformconfiguration/machineidentifier
func (v_ VZGenericMachineIdentifier) SetMachineIdentifier(value IVZGenericMachineIdentifier) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMachineIdentifier:"), value)
}/* debug [instance_properties/setter]: machineIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZGenericMachineIdentifier */


