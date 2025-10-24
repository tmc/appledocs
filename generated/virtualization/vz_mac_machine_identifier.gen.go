// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZMacMachineIdentifier */

/* debug [class_header]: Header for VZMacMachineIdentifier */
// The class instance for the [VZMacMachineIdentifier] class.
var (
	VZMacMachineIdentifierClass     _VZMacMachineIdentifierClass
	VZMacMachineIdentifierClassOnce sync.Once
)

func getVZMacMachineIdentifierClass() _VZMacMachineIdentifierClass {
	VZMacMachineIdentifierClassOnce.Do(func() {
		VZMacMachineIdentifierClass = _VZMacMachineIdentifierClass{objc.GetClass("VZMacMachineIdentifier")}
	})
	return VZMacMachineIdentifierClass
}

type _VZMacMachineIdentifierClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZMacMachineIdentifier */
// An interface definition for the [VZMacMachineIdentifier] class.
type IVZMacMachineIdentifier interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for VZMacMachineIdentifier */
	// properties:
	DataRepresentation() objc.IObject /* cross-framework: NSData */
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZMacMachineIdentifier */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZMacMachineIdentifier */
// Alloc allocates a new instance without initialization.
func (vc _VZMacMachineIdentifierClass) Alloc() VZMacMachineIdentifier {
	rv := objc.Send[VZMacMachineIdentifier](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZMacMachineIdentifierClass) New() VZMacMachineIdentifier {
	rv := objc.Send[VZMacMachineIdentifier](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZMacMachineIdentifier) Init() VZMacMachineIdentifier {
	rv := objc.Send[VZMacMachineIdentifier](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZMacMachineIdentifier) Autorelease() VZMacMachineIdentifier {
	rv := objc.Send[VZMacMachineIdentifier](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacMachineIdentifier creates a new VZMacMachineIdentifier instance.
func NewVZMacMachineIdentifier() VZMacMachineIdentifier {
	return getVZMacMachineIdentifierClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZMacMachineIdentifier */
// A unique identifier for a VM.
//
// This value uniquely identifies a virtual Mac hardware instance. Two VMs running concurrently shouldn’t use the same identifier. When serializing the VM to disk, you can preserve the identifier in a binary representation by serializing the data in the . property. Conversely, you can recreate the identifier with from the binary representation. You can compare the contents of two identifiers with .

// A unique identifier for a VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacMachineIdentifier
type VZMacMachineIdentifier struct {
	objectivec.Object
}

// VZMacMachineIdentifierFrom constructs a [VZMacMachineIdentifier] from an unsafe.Pointer.
//
// A unique identifier for a VM.
func VZMacMachineIdentifierFrom(ptr unsafe.Pointer) VZMacMachineIdentifier {
	return VZMacMachineIdentifier{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZMacMachineIdentifier */

// Create a machine identifier described by the specified data representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacMachineIdentifier/init(dataRepresentation:)
func NewVZMacMachineIdentifierWithDataRepresentation(dataRepresentation objc.IObject /* cross-framework: NSData */) VZMacMachineIdentifier {
	instance := getVZMacMachineIdentifierClass().Alloc()
	rv := objc.Send[VZMacMachineIdentifier](instance.ID, objc.Sel("initWithDataRepresentation:"), dataRepresentation)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewVZMacMachineIdentifierWithDataRepresentation */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZMacMachineIdentifier */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZMacMachineIdentifier */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZMacMachineIdentifier */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZMacMachineIdentifier */

// Returns the opaque data representation of the machine identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacMachineIdentifier/dataRepresentation
func (v_ VZMacMachineIdentifier) DataRepresentation() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](v_.ID, objc.Sel("dataRepresentation"))
	return rv
} /* debug [instance_properties/getter]: dataRepresentation */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZMacMachineIdentifier */
