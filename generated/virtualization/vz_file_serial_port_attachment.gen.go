// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class VZFileSerialPortAttachment */

/* debug [class_header]: Header for VZFileSerialPortAttachment */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for VZFileSerialPortAttachment */
// An interface definition for the [VZFileSerialPortAttachment] class.
type IVZFileSerialPortAttachment interface {
	IVZSerialPortAttachment

	/* debug [class_interface_properties]: Properties for VZFileSerialPortAttachment */
	// properties:
	Append() bool
	URL() objc.IObject /* cross-framework: NSURL */
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for VZFileSerialPortAttachment */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for VZFileSerialPortAttachment */
// Alloc allocates a new instance without initialization.
func (vc _VZFileSerialPortAttachmentClass) Alloc() VZFileSerialPortAttachment {
	rv := objc.Send[VZFileSerialPortAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for VZFileSerialPortAttachment */
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for VZFileSerialPortAttachment */

// Creates a file-based serial port attachment object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZFileSerialPortAttachment/init(url:append:)
func NewVZFileSerialPortAttachmentWithURLAppendError(url objc.IObject /* cross-framework: NSURL */, shouldAppend bool, error_ unsafe.Pointer) VZFileSerialPortAttachment {
	instance := getVZFileSerialPortAttachmentClass().Alloc()
	rv := objc.Send[VZFileSerialPortAttachment](instance.ID, objc.Sel("initWithURL:append:error:"), url, shouldAppend, error_)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewVZFileSerialPortAttachmentWithURLAppendError */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for VZFileSerialPortAttachment */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for VZFileSerialPortAttachment */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for VZFileSerialPortAttachment */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for VZFileSerialPortAttachment */

// A Boolean that indicates whether the virtual machine appends data to the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZFileSerialPortAttachment/append
func (v_ VZFileSerialPortAttachment) Append() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("append"))
	return rv
} /* debug [instance_properties/getter]: append */

// The URL of a file on the local file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZFileSerialPortAttachment/url
func (v_ VZFileSerialPortAttachment) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](v_.ID, objc.Sel("URL"))
	return rv
} /* debug [instance_properties/getter]: URL */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VZFileSerialPortAttachment */
