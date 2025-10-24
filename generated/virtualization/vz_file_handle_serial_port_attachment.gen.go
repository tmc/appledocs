// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [VZFileHandleSerialPortAttachment] class.
type IVZFileHandleSerialPortAttachment interface {
	IVZSerialPortAttachment
	// properties:
	FileHandleForReading() objc.IObject /* cross-framework: FileHandle */
	FileHandleForWriting() objc.IObject /* cross-framework: FileHandle */
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (vc _VZFileHandleSerialPortAttachmentClass) Alloc() VZFileHandleSerialPortAttachment {
	rv := objc.Send[VZFileHandleSerialPortAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates a serial port attachment object from the specified file handles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZFileHandleSerialPortAttachment/init(fileHandleForReading:fileHandleForWriting:)
func NewVZFileHandleSerialPortAttachmentWithFileHandleForReadingFileHandleForWriting(fileHandleForReading objc.IObject /* cross-framework: FileHandle */, fileHandleForWriting objc.IObject /* cross-framework: FileHandle */) VZFileHandleSerialPortAttachment {
	instance := getVZFileHandleSerialPortAttachmentClass().Alloc()
	rv := objc.Send[VZFileHandleSerialPortAttachment](instance.ID, objc.Sel("initWithFileHandleForReading:fileHandleForWriting:"), fileHandleForReading, fileHandleForWriting)
	rv.Autorelease()
	return rv
}



// The file handle that the guest operating system uses to read data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZFileHandleSerialPortAttachment/fileHandleForReading
func (v_ VZFileHandleSerialPortAttachment) FileHandleForReading() objc.IObject /* cross-framework: FileHandle */ {
	rv := objc.Send[foundation.FileHandle](v_.ID, objc.Sel("fileHandleForReading"))
	return rv
}


// The file handle that the guest operating system uses to write data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZFileHandleSerialPortAttachment/fileHandleForWriting
func (v_ VZFileHandleSerialPortAttachment) FileHandleForWriting() objc.IObject /* cross-framework: FileHandle */ {
	rv := objc.Send[foundation.FileHandle](v_.ID, objc.Sel("fileHandleForWriting"))
	return rv
}


