// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZSpiceAgentPortAttachment */

/* debug [class_header]: Header for VZSpiceAgentPortAttachment */
// The class instance for the [VZSpiceAgentPortAttachment] class.
var (
	VZSpiceAgentPortAttachmentClass     _VZSpiceAgentPortAttachmentClass
	VZSpiceAgentPortAttachmentClassOnce sync.Once
)

func getVZSpiceAgentPortAttachmentClass() _VZSpiceAgentPortAttachmentClass {
	VZSpiceAgentPortAttachmentClassOnce.Do(func() {
		VZSpiceAgentPortAttachmentClass = _VZSpiceAgentPortAttachmentClass{objc.GetClass("VZSpiceAgentPortAttachment")}
	})
	return VZSpiceAgentPortAttachmentClass
}

type _VZSpiceAgentPortAttachmentClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZSpiceAgentPortAttachment */
// An interface definition for the [VZSpiceAgentPortAttachment] class.
type IVZSpiceAgentPortAttachment interface {
	IVZSerialPortAttachment

	/* debug [class_interface_properties]: Properties for VZSpiceAgentPortAttachment */
	// properties:
	SharesClipboard() bool
	SetSharesClipboard(value bool)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZSpiceAgentPortAttachment */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZSpiceAgentPortAttachment */
// Alloc allocates a new instance without initialization.
func (vc _VZSpiceAgentPortAttachmentClass) Alloc() VZSpiceAgentPortAttachment {
	rv := objc.Send[VZSpiceAgentPortAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZSpiceAgentPortAttachmentClass) New() VZSpiceAgentPortAttachment {
	rv := objc.Send[VZSpiceAgentPortAttachment](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZSpiceAgentPortAttachment) Init() VZSpiceAgentPortAttachment {
	rv := objc.Send[VZSpiceAgentPortAttachment](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZSpiceAgentPortAttachment) Autorelease() VZSpiceAgentPortAttachment {
	rv := objc.Send[VZSpiceAgentPortAttachment](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSpiceAgentPortAttachment creates a new VZSpiceAgentPortAttachment instance.
func NewVZSpiceAgentPortAttachment() VZSpiceAgentPortAttachment {
	return getVZSpiceAgentPortAttachmentClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZSpiceAgentPortAttachment */
// An attachment point that enables the Spice clipboard sharing capability.

// An attachment point that enables the Spice clipboard sharing capability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSpiceAgentPortAttachment
type VZSpiceAgentPortAttachment struct {
	VZSerialPortAttachment
}

// VZSpiceAgentPortAttachmentFrom constructs a [VZSpiceAgentPortAttachment] from an unsafe.Pointer.
//
// An attachment point that enables the Spice clipboard sharing capability.
func VZSpiceAgentPortAttachmentFrom(ptr unsafe.Pointer) VZSpiceAgentPortAttachment {
	return VZSpiceAgentPortAttachment{
		VZSerialPortAttachment: VZSerialPortAttachmentFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZSpiceAgentPortAttachment */
/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZSpiceAgentPortAttachment */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZSpiceAgentPortAttachment */

// The name of the Virtio console port for the Spice guest agent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSpiceAgentPortAttachment/spiceAgentPortName
func (vc _VZSpiceAgentPortAttachmentClass) SpiceAgentPortName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](objc.ID(vc.class), objc.Sel("spiceAgentPortName"))
	return rv
} /* debug [class_properties_class/property]: spiceAgentPortName */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZSpiceAgentPortAttachment */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZSpiceAgentPortAttachment */

// A Boolean value that indicates whether the framework needs to share the clipboard between the host and the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSpiceAgentPortAttachment/sharesClipboard
func (v_ VZSpiceAgentPortAttachment) SharesClipboard() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("sharesClipboard"))
	return rv
} /* debug [instance_properties/getter]: sharesClipboard */

// A Boolean value that indicates whether the framework needs to share the clipboard between the host and the VM.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSpiceAgentPortAttachment/sharesClipboard
func (v_ VZSpiceAgentPortAttachment) SetSharesClipboard(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSharesClipboard:"), value)
} /* debug [instance_properties/setter]: sharesClipboard */

// The name of the Virtio console port for the Spice guest agent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSpiceAgentPortAttachment/spiceAgentPortName
func (v_ VZSpiceAgentPortAttachment) SpiceAgentPortName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("spiceAgentPortName"))
	return rv
} /* debug [instance_properties/getter]: spiceAgentPortName */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZSpiceAgentPortAttachment */
