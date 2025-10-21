// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

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

// An interface definition for the [VZSpiceAgentPortAttachment] class.
type IVZSpiceAgentPortAttachment interface {
	IVZSerialPortAttachment
}

// An attachment point that enables the Spice clipboard sharing capability.
//
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

// Alloc allocates a new instance without initialization.
func (vc _VZSpiceAgentPortAttachmentClass) Alloc() VZSpiceAgentPortAttachment {
	rv := objc.Send[VZSpiceAgentPortAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The name of the Virtio console port for the Spice guest agent.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSpiceAgentPortAttachment/spiceAgentPortName
func (vc _VZSpiceAgentPortAttachmentClass) SpiceAgentPortName() appkit.string {
	rv := objc.Send[appkit.string](objc.ID(vc.class), objc.Sel("spiceAgentPortName"))
	return rv
}
// A Boolean value that indicates whether the framework needs to share the clipboard between the host and the VM.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSpiceAgentPortAttachment/sharesClipboard
func (v_ VZSpiceAgentPortAttachment) SharesClipboard() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("sharesClipboard"))
	return rv
}


// SetSharesClipboard sets the value of the sharesClipboard property.
// A Boolean value that indicates whether the framework needs to share the clipboard between the host and the VM.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSpiceAgentPortAttachment/sharesClipboard
func (v_ VZSpiceAgentPortAttachment) SetSharesClipboard(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSharesClipboard:"), value)
}

// The name of the Virtio console port for the Spice guest agent.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZSpiceAgentPortAttachment/spiceAgentPortName
func (v_ VZSpiceAgentPortAttachment) SpiceAgentPortName() appkit.string {
	rv := objc.Send[appkit.string](v_.ID, objc.Sel("spiceAgentPortName"))
	return rv
}


