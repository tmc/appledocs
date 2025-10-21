// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [VZNetworkDevice] class.
var (
	VZNetworkDeviceClass     _VZNetworkDeviceClass
	VZNetworkDeviceClassOnce sync.Once
)

func getVZNetworkDeviceClass() _VZNetworkDeviceClass {
	VZNetworkDeviceClassOnce.Do(func() {
		VZNetworkDeviceClass = _VZNetworkDeviceClass{objc.GetClass("VZNetworkDevice")}
	})
	return VZNetworkDeviceClass
}

type _VZNetworkDeviceClass struct {
	class objc.Class
}

// An interface definition for the [VZNetworkDevice] class.
type IVZNetworkDevice interface {
	objectivec.IObject
}

// A base class that represents a network device in a virtual machine.
//
// Don’t instantiate a   directly. When you create a   instance with a   the system creates the number of network devices based on the number of   objects you specify in the VM configuration. Before initializing the virtual machine (VM), validate the configuration using   to ensure the user’s computer supports the number of network and other devices you’ve specified.   For many purposes, a single network that uses a Network Address Translation (NAT) attachment and connects the VM to the host computer’s network is sufficient. You can use additional network interfaces for purposes of your own design, such as: Bridging several physical interfaces to connect to multiple networks. Using the file descriptor attachment to create specialized connections for different purposes. You access the network devices through the  . property. The network devices map to their respective configurations in a one to one relationship, where index   of   corresponds to the network device configuration at index   set on  . .
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZNetworkDevice
type VZNetworkDevice struct {
	objectivec.Object
}

// VZNetworkDeviceFrom constructs a [VZNetworkDevice] from an unsafe.Pointer.
//
// A base class that represents a network device in a virtual machine.
func VZNetworkDeviceFrom(ptr unsafe.Pointer) VZNetworkDevice {
	return VZNetworkDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZNetworkDeviceClass) Alloc() VZNetworkDevice {
	rv := objc.Send[VZNetworkDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZNetworkDeviceClass) New() VZNetworkDevice {
	rv := objc.Send[VZNetworkDevice](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZNetworkDevice) Init() VZNetworkDevice {
	rv := objc.Send[VZNetworkDevice](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZNetworkDevice) Autorelease() VZNetworkDevice {
	rv := objc.Send[VZNetworkDevice](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZNetworkDevice creates a new VZNetworkDevice instance.
func NewVZNetworkDevice() VZNetworkDevice {
	return getVZNetworkDeviceClass().New()
}




