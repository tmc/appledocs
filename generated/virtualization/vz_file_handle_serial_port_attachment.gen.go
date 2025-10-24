// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZFileHandleSerialPortAttachment */


/* debug [class_header]: Header for VZFileHandleSerialPortAttachment */
// The class instance for the [VZFileHandleSerialPortAttachment] class.
var (
	VZFileHandleSerialPortAttachmentClass     _VZFileHandleSerialPortAttachmentClass
	VZFileHandleSerialPortAttachmentClassOnce sync.Once
)

func getVZFileHandleSerialPortAttachmentClass() _VZFileHandleSerialPortAttachmentClass {
	VZFileHandleSerialPortAttachmentClassOnce.Do(func() {
		VZFileHandleSerialPortAttachmentClass = _VZFileHandleSerialPortAttachmentClass{objc.GetClass("VZFileHandleSerialPortAttachment")}
	})
	return VZFileHandleSerialPortAttachmentClass
}

type _VZFileHandleSerialPortAttachmentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZFileHandleSerialPortAttachment */
// An interface definition for the [VZFileHandleSerialPortAttachment] class.
type IVZFileHandleSerialPortAttachment interface {
	IVZSerialPortAttachment
	
/* debug [class_interface_properties]: Properties for VZFileHandleSerialPortAttachment */
	// properties:
	FileHandleForReading() foundation.FileHandle
	FileHandleForWriting() foundation.FileHandle
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZFileHandleSerialPortAttachment */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZFileHandleSerialPortAttachment */
// Alloc allocates a new instance without initialization.
func (vc _VZFileHandleSerialPortAttachmentClass) Alloc() VZFileHandleSerialPortAttachment {
	rv := objc.Send[VZFileHandleSerialPortAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZFileHandleSerialPortAttachmentClass) New() VZFileHandleSerialPortAttachment {
	rv := objc.Send[VZFileHandleSerialPortAttachment](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZFileHandleSerialPortAttachment) Init() VZFileHandleSerialPortAttachment {
	rv := objc.Send[VZFileHandleSerialPortAttachment](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZFileHandleSerialPortAttachment) Autorelease() VZFileHandleSerialPortAttachment {
	rv := objc.Send[VZFileHandleSerialPortAttachment](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZFileHandleSerialPortAttachment creates a new VZFileHandleSerialPortAttachment instance.
func NewVZFileHandleSerialPortAttachment() VZFileHandleSerialPortAttachment {
	return getVZFileHandleSerialPortAttachmentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZFileHandleSerialPortAttachment */
// An attachment point that allows bidirectional communication using file handles.
//
// Use a object to configure a serial port using separate file handles for reading and writing data. In your virtual machine, use the file handles in this object in the following way: To send data to the guest operating system, write data to the file handle in the property. To receive data from the guest operating system, read data from the file handle in the property.


// An attachment point that allows bidirectional communication using file handles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZFileHandleSerialPortAttachment
type VZFileHandleSerialPortAttachment struct {
	VZSerialPortAttachment
}

// VZFileHandleSerialPortAttachmentFrom constructs a [VZFileHandleSerialPortAttachment] from an unsafe.Pointer.
//
// An attachment point that allows bidirectional communication using file handles.
func VZFileHandleSerialPortAttachmentFrom(ptr unsafe.Pointer) VZFileHandleSerialPortAttachment {
	return VZFileHandleSerialPortAttachment{
		VZSerialPortAttachment: VZSerialPortAttachmentFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZFileHandleSerialPortAttachment */

// Creates a serial port attachment object from the specified file handles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZFileHandleSerialPortAttachment/init(fileHandleForReading:fileHandleForWriting:)
func NewVZFileHandleSerialPortAttachmentWithFileHandleForReadingFileHandleForWriting(fileHandleForReading foundation.FileHandle, fileHandleForWriting foundation.FileHandle) VZFileHandleSerialPortAttachment {
	instance := getVZFileHandleSerialPortAttachmentClass().Alloc()
	rv := objc.Send[VZFileHandleSerialPortAttachment](instance.ID, objc.Sel("initWithFileHandleForReading:fileHandleForWriting:"), fileHandleForReading, fileHandleForWriting)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewVZFileHandleSerialPortAttachmentWithFileHandleForReadingFileHandleForWriting */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZFileHandleSerialPortAttachment */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZFileHandleSerialPortAttachment */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZFileHandleSerialPortAttachment */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZFileHandleSerialPortAttachment */

// The file handle that the guest operating system uses to read data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZFileHandleSerialPortAttachment/fileHandleForReading
func (v_ VZFileHandleSerialPortAttachment) FileHandleForReading() foundation.FileHandle {
	rv := objc.Send[foundation.FileHandle](v_.ID, objc.Sel("fileHandleForReading"))
	return rv
}/* debug [instance_properties/getter]: fileHandleForReading */


// The file handle that the guest operating system uses to write data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZFileHandleSerialPortAttachment/fileHandleForWriting
func (v_ VZFileHandleSerialPortAttachment) FileHandleForWriting() foundation.FileHandle {
	rv := objc.Send[foundation.FileHandle](v_.ID, objc.Sel("fileHandleForWriting"))
	return rv
}/* debug [instance_properties/getter]: fileHandleForWriting */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZFileHandleSerialPortAttachment */


