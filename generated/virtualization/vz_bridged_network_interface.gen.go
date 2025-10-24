// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZBridgedNetworkInterface */


/* debug [class_header]: Header for VZBridgedNetworkInterface */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZBridgedNetworkInterface */
// An interface definition for the [VZBridgedNetworkInterface] class.
type IVZBridgedNetworkInterface interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZBridgedNetworkInterface */
	// properties:
	Identifier() objc.IObject /* cross-framework: NSString */
	LocalizedDisplayName() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZBridgedNetworkInterface */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZBridgedNetworkInterface */
// Alloc allocates a new instance without initialization.
func (vc _VZBridgedNetworkInterfaceClass) Alloc() VZBridgedNetworkInterface {
	rv := objc.Send[VZBridgedNetworkInterface](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZBridgedNetworkInterface */
// An object that identifies the supported network interfaces of the host computer.
//
// Use a object to retrieve the physical interfaces on the host computer. Use a bridged network interface to create a object, which maps that interface to one of your virtual machine’s network devices. The host computer and your virtual machine share access to the physical network interface, but communicate over it using distinct network layers. You don’t create objects directly. Instead, the system creates one object for each physical interface of the host computer and stores those objects in the property. Iterate over the objects in that property to retrieve the network interfaces you need.


// An object that identifies the supported network interfaces of the host computer.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZBridgedNetworkInterface *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZBridgedNetworkInterface */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZBridgedNetworkInterface */

// The bridged network interfaces that you may use in your virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkInterface/networkInterfaces
func (vc _VZBridgedNetworkInterfaceClass) NetworkInterfaces() []VZBridgedNetworkInterface {
	rv := objc.Send[[]VZBridgedNetworkInterface](objc.ID(vc.class), objc.Sel("networkInterfaces"))
	return rv
}/* debug [class_properties_class/property]: networkInterfaces */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZBridgedNetworkInterface */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZBridgedNetworkInterface */

// The unique BSD name of this network interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkInterface/identifier
func (v_ VZBridgedNetworkInterface) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// A user-visible name for the network interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkInterface/localizedDisplayName
func (v_ VZBridgedNetworkInterface) LocalizedDisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("localizedDisplayName"))
	return rv
}/* debug [instance_properties/getter]: localizedDisplayName */


// The bridged network interfaces that you may use in your virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZBridgedNetworkInterface/networkInterfaces
func (v_ VZBridgedNetworkInterface) NetworkInterfaces() []VZBridgedNetworkInterface {
	rv := objc.Send[[]VZBridgedNetworkInterface](v_.ID, objc.Sel("networkInterfaces"))
	return rv
}/* debug [instance_properties/getter]: networkInterfaces */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZBridgedNetworkInterface */



