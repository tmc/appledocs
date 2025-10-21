// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [VZMacMachineIdentifier] class.
type IVZMacMachineIdentifier interface {
	objectivec.IObject
}

// A unique identifier for a VM.
//
// This value uniquely identifies a virtual Mac hardware instance. Two VMs running concurrently shouldn’t use the same identifier. When serializing the VM to disk, you can preserve the identifier in a binary representation by serializing the data in the . property. Conversely, you can recreate the identifier with from the binary representation. You can compare the contents of two identifiers with .
//
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

// Alloc allocates a new instance without initialization.
func (vc _VZMacMachineIdentifierClass) Alloc() VZMacMachineIdentifier {
	rv := objc.Send[VZMacMachineIdentifier](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Create a machine identifier described by the specified data representation.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacMachineIdentifier/init(dataRepresentation:)
func NewVZMacMachineIdentifierWithDataRepresentation(dataRepresentation unsafe.Pointer) VZMacMachineIdentifier {
	instance := getVZMacMachineIdentifierClass().Alloc()
	rv := objc.Send[VZMacMachineIdentifier](instance.ID, objc.Sel("initWithDataRepresentation:"), dataRepresentation)
	rv.Autorelease()
	return rv
}


// Returns the opaque data representation of the machine identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZMacMachineIdentifier/dataRepresentation
func (v_ VZMacMachineIdentifier) DataRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("dataRepresentation"))
	return rv
}


