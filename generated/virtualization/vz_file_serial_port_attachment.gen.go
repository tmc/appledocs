// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [VZFileSerialPortAttachment] class.
var (
	VZFileSerialPortAttachmentClass     _VZFileSerialPortAttachmentClass
	VZFileSerialPortAttachmentClassOnce sync.Once
)

func getVZFileSerialPortAttachmentClass() _VZFileSerialPortAttachmentClass {
	VZFileSerialPortAttachmentClassOnce.Do(func() {
		VZFileSerialPortAttachmentClass = _VZFileSerialPortAttachmentClass{objc.GetClass("VZFileSerialPortAttachment")}
	})
	return VZFileSerialPortAttachmentClass
}

type _VZFileSerialPortAttachmentClass struct {
	class objc.Class
}

// An interface definition for the [VZFileSerialPortAttachment] class.
type IVZFileSerialPortAttachment interface {
	IVZSerialPortAttachment
	// properties:
	Append() bool
	SetAppend(value bool)
	Url() objc.IObject /* cross-framework: URL */
	SetUrl(value objc.IObject /* cross-framework: URL */)
	// methods:
}

// An attachment point that writes data from the guest system to a file.
//
// Use a object to configure a one-way serial port from the guest operating system to the virtual machine. When the guest sends data to the serial port, the virtual machine writes that data to the specified file. You can’t use this serial port to send data back to the guest. Create a object and assign it to an appropriate subclass of object, such as . The file you use to create this object must be writable.


// An attachment point that writes data from the guest system to a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZFileSerialPortAttachment
type VZFileSerialPortAttachment struct {
	VZSerialPortAttachment
}

// VZFileSerialPortAttachmentFrom constructs a [VZFileSerialPortAttachment] from an unsafe.Pointer.
//
// An attachment point that writes data from the guest system to a file.
func VZFileSerialPortAttachmentFrom(ptr unsafe.Pointer) VZFileSerialPortAttachment {
	return VZFileSerialPortAttachment{
		VZSerialPortAttachment: VZSerialPortAttachmentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _VZFileSerialPortAttachmentClass) Alloc() VZFileSerialPortAttachment {
	rv := objc.Send[VZFileSerialPortAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZFileSerialPortAttachmentClass) New() VZFileSerialPortAttachment {
	rv := objc.Send[VZFileSerialPortAttachment](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZFileSerialPortAttachment) Init() VZFileSerialPortAttachment {
	rv := objc.Send[VZFileSerialPortAttachment](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZFileSerialPortAttachment) Autorelease() VZFileSerialPortAttachment {
	rv := objc.Send[VZFileSerialPortAttachment](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZFileSerialPortAttachment creates a new VZFileSerialPortAttachment instance.
func NewVZFileSerialPortAttachment() VZFileSerialPortAttachment {
	return getVZFileSerialPortAttachmentClass().New()
}



// A Boolean that indicates whether the virtual machine appends data to the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzfileserialportattachment/append
func (v_ VZFileSerialPortAttachment) Append() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("append"))
	return rv
}


// A Boolean that indicates whether the virtual machine appends data to the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzfileserialportattachment/append
func (v_ VZFileSerialPortAttachment) SetAppend(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAppend:"), value)
}


// The URL of a file on the local file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzfileserialportattachment/url
func (v_ VZFileSerialPortAttachment) Url() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](v_.ID, objc.Sel("url"))
	return rv
}


// The URL of a file on the local file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/virtualization/vzfileserialportattachment/url
func (v_ VZFileSerialPortAttachment) SetUrl(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setUrl:"), value)
}



