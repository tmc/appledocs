// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZVirtioConsolePort */


/* debug [class_header]: Header for VZVirtioConsolePort */
// The class instance for the [VZVirtioConsolePort] class.
var (
	VZVirtioConsolePortClass     _VZVirtioConsolePortClass
	VZVirtioConsolePortClassOnce sync.Once
)

func getVZVirtioConsolePortClass() _VZVirtioConsolePortClass {
	VZVirtioConsolePortClassOnce.Do(func() {
		VZVirtioConsolePortClass = _VZVirtioConsolePortClass{objc.GetClass("VZVirtioConsolePort")}
	})
	return VZVirtioConsolePortClass
}

type _VZVirtioConsolePortClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZVirtioConsolePort */
// An interface definition for the [VZVirtioConsolePort] class.
type IVZVirtioConsolePort interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZVirtioConsolePort */
	// properties:
	Attachment() IVZSerialPortAttachment
	SetAttachment(value IVZSerialPortAttachment)
	Name() objc.IObject /* cross-framework: NSString */
	Ports() IVZVirtioConsolePortArray
	SetPorts(value IVZVirtioConsolePortArray)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZVirtioConsolePort */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZVirtioConsolePort */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtioConsolePortClass) Alloc() VZVirtioConsolePort {
	rv := objc.Send[VZVirtioConsolePort](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZVirtioConsolePortClass) New() VZVirtioConsolePort {
	rv := objc.Send[VZVirtioConsolePort](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioConsolePort) Init() VZVirtioConsolePort {
	rv := objc.Send[VZVirtioConsolePort](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioConsolePort) Autorelease() VZVirtioConsolePort {
	rv := objc.Send[VZVirtioConsolePort](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioConsolePort creates a new VZVirtioConsolePort instance.
func NewVZVirtioConsolePort() VZVirtioConsolePort {
	return getVZVirtioConsolePortClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZVirtioConsolePort */
// A class that represents a Virtio console port in a VM.
//
// Don’t instantiate a directly. You retrieve this object from the property.


// A class that represents a Virtio console port in a VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePort
type VZVirtioConsolePort struct {
	objectivec.Object
}

// VZVirtioConsolePortFrom constructs a [VZVirtioConsolePort] from an unsafe.Pointer.
//
// A class that represents a Virtio console port in a VM.
func VZVirtioConsolePortFrom(ptr unsafe.Pointer) VZVirtioConsolePort {
	return VZVirtioConsolePort{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZVirtioConsolePort *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZVirtioConsolePort */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZVirtioConsolePort */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZVirtioConsolePort */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZVirtioConsolePort */

// An array of serial port attachments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePort/attachment
func (v_ VZVirtioConsolePort) Attachment() IVZSerialPortAttachment {
	rv := objc.Send[VZSerialPortAttachment](v_.ID, objc.Sel("attachment"))
	return rv
}/* debug [instance_properties/getter]: attachment */


// An array of serial port attachments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePort/attachment
func (v_ VZVirtioConsolePort) SetAttachment(value IVZSerialPortAttachment) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAttachment:"), value)
}/* debug [instance_properties/setter]: attachment */


// The name of the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePort/name
func (v_ VZVirtioConsolePort) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The array of console ports that a specific device uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtioconsoledevice/ports
func (v_ VZVirtioConsolePort) Ports() IVZVirtioConsolePortArray {
	rv := objc.Send[VZVirtioConsolePortArray](v_.ID, objc.Sel("ports"))
	return rv
}/* debug [instance_properties/getter]: ports */


// The array of console ports that a specific device uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzvirtioconsoledevice/ports
func (v_ VZVirtioConsolePort) SetPorts(value IVZVirtioConsolePortArray) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPorts:"), value)
}/* debug [instance_properties/setter]: ports */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZVirtioConsolePort */



