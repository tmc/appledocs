// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZSerialPortAttachment */

/* debug [class_header]: Header for VZSerialPortAttachment */
// The class instance for the [VZSerialPortAttachment] class.
var (
	VZSerialPortAttachmentClass     _VZSerialPortAttachmentClass
	VZSerialPortAttachmentClassOnce sync.Once
)

func getVZSerialPortAttachmentClass() _VZSerialPortAttachmentClass {
	VZSerialPortAttachmentClassOnce.Do(func() {
		VZSerialPortAttachmentClass = _VZSerialPortAttachmentClass{objc.GetClass("VZSerialPortAttachment")}
	})
	return VZSerialPortAttachmentClass
}

type _VZSerialPortAttachmentClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZSerialPortAttachment */
// An interface definition for the [VZSerialPortAttachment] class.
type IVZSerialPortAttachment interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for VZSerialPortAttachment */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZSerialPortAttachment */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZSerialPortAttachment */
// Alloc allocates a new instance without initialization.
func (vc _VZSerialPortAttachmentClass) Alloc() VZSerialPortAttachment {
	rv := objc.Send[VZSerialPortAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZSerialPortAttachmentClass) New() VZSerialPortAttachment {
	rv := objc.Send[VZSerialPortAttachment](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZSerialPortAttachment) Init() VZSerialPortAttachment {
	rv := objc.Send[VZSerialPortAttachment](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZSerialPortAttachment) Autorelease() VZSerialPortAttachment {
	rv := objc.Send[VZSerialPortAttachment](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSerialPortAttachment creates a new VZSerialPortAttachment instance.
func NewVZSerialPortAttachment() VZSerialPortAttachment {
	return getVZSerialPortAttachmentClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZSerialPortAttachment */
// The common behaviors for the serial attachment points of your virtual machine.
//
// Don’t create a object directly. Instead, instantiate a concrete subclass such as to configure how the virtual machine’s serial port connects with the host computer.

// The common behaviors for the serial attachment points of your virtual machine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSerialPortAttachment
type VZSerialPortAttachment struct {
	objectivec.Object
}

// VZSerialPortAttachmentFrom constructs a [VZSerialPortAttachment] from an unsafe.Pointer.
//
// The common behaviors for the serial attachment points of your virtual machine.
func VZSerialPortAttachmentFrom(ptr unsafe.Pointer) VZSerialPortAttachment {
	return VZSerialPortAttachment{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZSerialPortAttachment */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZSerialPortAttachment */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZSerialPortAttachment */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZSerialPortAttachment */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZSerialPortAttachment */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZSerialPortAttachment */
