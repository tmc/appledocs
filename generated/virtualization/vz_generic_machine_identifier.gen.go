// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [VZGenericMachineIdentifier] class.
type IVZGenericMachineIdentifier interface {
	objectivec.IObject
	// properties:
	DataRepresentation() objc.IObject /* cross-framework: Data */
	SetDataRepresentation(value objc.IObject /* cross-framework: Data */)
	IsNestedVirtualizationEnabled() bool
	SetIsNestedVirtualizationEnabled(value bool)
	MachineIdentifier() IVZGenericMachineIdentifier
	SetMachineIdentifier(value IVZGenericMachineIdentifier)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (vc _VZGenericMachineIdentifierClass) Alloc() VZGenericMachineIdentifier {
	rv := objc.Send[VZGenericMachineIdentifier](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// An opaque data representation of the VM’s identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzgenericmachineidentifier/datarepresentation
func (v_ VZGenericMachineIdentifier) DataRepresentation() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](v_.ID, objc.Sel("dataRepresentation"))
	return rv
}


// An opaque data representation of the VM’s identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzgenericmachineidentifier/datarepresentation
func (v_ VZGenericMachineIdentifier) SetDataRepresentation(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDataRepresentation:"), value)
}


// A Boolean value that indicates whether nested virtualization is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzgenericplatformconfiguration/isnestedvirtualizationenabled
func (v_ VZGenericMachineIdentifier) IsNestedVirtualizationEnabled() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isNestedVirtualizationEnabled"))
	return rv
}


// A Boolean value that indicates whether nested virtualization is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzgenericplatformconfiguration/isnestedvirtualizationenabled
func (v_ VZGenericMachineIdentifier) SetIsNestedVirtualizationEnabled(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsNestedVirtualizationEnabled:"), value)
}


// A value that represents a unique identifier for the virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzgenericplatformconfiguration/machineidentifier
func (v_ VZGenericMachineIdentifier) MachineIdentifier() IVZGenericMachineIdentifier {
	rv := objc.Send[VZGenericMachineIdentifier](v_.ID, objc.Sel("machineIdentifier"))
	return rv
}


// A value that represents a unique identifier for the virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzgenericplatformconfiguration/machineidentifier
func (v_ VZGenericMachineIdentifier) SetMachineIdentifier(value IVZGenericMachineIdentifier) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMachineIdentifier:"), value)
}



