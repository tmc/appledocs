// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [VZBridgedNetworkInterface] class.
var (
	VZBridgedNetworkInterfaceClass     _VZBridgedNetworkInterfaceClass
	VZBridgedNetworkInterfaceClassOnce sync.Once
)

func getVZBridgedNetworkInterfaceClass() _VZBridgedNetworkInterfaceClass {
	VZBridgedNetworkInterfaceClassOnce.Do(func() {
		VZBridgedNetworkInterfaceClass = _VZBridgedNetworkInterfaceClass{objc.GetClass("VZBridgedNetworkInterface")}
	})
	return VZBridgedNetworkInterfaceClass
}

type _VZBridgedNetworkInterfaceClass struct {
	class objc.Class
}

// An interface definition for the [VZBridgedNetworkInterface] class.
type IVZBridgedNetworkInterface interface {
	objectivec.IObject
}

// An object that identifies the supported network interfaces of the host computer.
//
// Use a object to retrieve the physical interfaces on the host computer. Use a bridged network interface to create a object, which maps that interface to one of your virtual machine’s network devices. The host computer and your virtual machine share access to the physical network interface, but communicate over it using distinct network layers. You don’t create objects directly. Instead, the system creates one object for each physical interface of the host computer and stores those objects in the property. Iterate over the objects in that property to retrieve the network interfaces you need.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkInterface
type VZBridgedNetworkInterface struct {
	objectivec.Object
}

// VZBridgedNetworkInterfaceFrom constructs a [VZBridgedNetworkInterface] from an unsafe.Pointer.
//
// An object that identifies the supported network interfaces of the host computer.
func VZBridgedNetworkInterfaceFrom(ptr unsafe.Pointer) VZBridgedNetworkInterface {
	return VZBridgedNetworkInterface{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZBridgedNetworkInterfaceClass) Alloc() VZBridgedNetworkInterface {
	rv := objc.Send[VZBridgedNetworkInterface](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZBridgedNetworkInterfaceClass) New() VZBridgedNetworkInterface {
	rv := objc.Send[VZBridgedNetworkInterface](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZBridgedNetworkInterface) Init() VZBridgedNetworkInterface {
	rv := objc.Send[VZBridgedNetworkInterface](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZBridgedNetworkInterface) Autorelease() VZBridgedNetworkInterface {
	rv := objc.Send[VZBridgedNetworkInterface](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZBridgedNetworkInterface creates a new VZBridgedNetworkInterface instance.
func NewVZBridgedNetworkInterface() VZBridgedNetworkInterface {
	return getVZBridgedNetworkInterfaceClass().New()
}




