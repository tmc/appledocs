// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FileHandle] class.
var (
	FileHandleClass     _FileHandleClass
	FileHandleClassOnce sync.Once
)

func getFileHandleClass() _FileHandleClass {
	FileHandleClassOnce.Do(func() {
		FileHandleClass = _FileHandleClass{objc.GetClass("NSFileHandle")}
	})
	return FileHandleClass
}

type _FileHandleClass struct {
	class objc.Class
}

// An interface definition for the [FileHandle] class.
type IFileHandle interface {
	objectivec.IObject
	// properties:
	AvailableData() IData
	SetAvailableData(value IData)
	Bytes() unsafe.Pointer
	SetBytes(value unsafe.Pointer)
	FileDescriptor() unsafe.Pointer
	SetFileDescriptor(value unsafe.Pointer)
	OffsetInFile() uint64
	SetOffsetInFile(value uint64)
	ReadabilityHandler() unsafe.Pointer
	SetReadabilityHandler(value unsafe.Pointer)
	WriteabilityHandler() unsafe.Pointer
	SetWriteabilityHandler(value unsafe.Pointer)
	NSFileHandleNotificationMonitorModes() IString
	// methods:
}

// An object-oriented wrapper for a file descriptor.
//
// You use file handle objects to access data associated with files, sockets, pipes, and devices. For files, you can read, write, and seek within the file. For sockets, pipes, and devices, you can use a file handle object to monitor the device and process data asynchronously. Most creation methods for cause the file handle object to take ownership of the associated file descriptor. This means that the file handle object both creates the file descriptor and is responsible for closing it later, usually when the system deallocates the file handle object. If you want to use a file handle object with a file descriptor that you created, use the method or use the method and pass for the parameter.


// An object-oriented wrapper for a file descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle
type FileHandle struct {
	objectivec.Object
}

// FileHandleFrom constructs a [FileHandle] from an unsafe.Pointer.
//
// An object-oriented wrapper for a file descriptor.
func FileHandleFrom(ptr unsafe.Pointer) FileHandle {
	return FileHandle{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FileHandleClass) Alloc() FileHandle {
	rv := objc.Send[FileHandle](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FileHandleClass) New() FileHandle {
	rv := objc.Send[FileHandle](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FileHandle) Init() FileHandle {
	rv := objc.Send[FileHandle](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FileHandle) Autorelease() FileHandle {
	rv := objc.Send[FileHandle](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileHandle creates a new FileHandle instance.
func NewFileHandle() FileHandle {
	return getFileHandleClass().New()
}



// The data currently available in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/availabledata
func (f_ FileHandle) AvailableData() IData {
	rv := objc.Send[Data](f_.ID, objc.Sel("availableData"))
	return rv
}


// The data currently available in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/availabledata
func (f_ FileHandle) SetAvailableData(value IData) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAvailableData:"), value)
}


// The file’s contents, as an asynchronous sequence of bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/bytes
func (f_ FileHandle) Bytes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("bytes"))
	return rv
}


// The file’s contents, as an asynchronous sequence of bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/bytes
func (f_ FileHandle) SetBytes(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setBytes:"), value)
}


// The POSIX file descriptor associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/filedescriptor
func (f_ FileHandle) FileDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("fileDescriptor"))
	return rv
}


// The POSIX file descriptor associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/filedescriptor
func (f_ FileHandle) SetFileDescriptor(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFileDescriptor:"), value)
}


// The position of the file pointer within the file represented by the file handle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/offsetinfile
func (f_ FileHandle) OffsetInFile() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("offsetInFile"))
	return rv
}


// The position of the file pointer within the file represented by the file handle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/offsetinfile
func (f_ FileHandle) SetOffsetInFile(value uint64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOffsetInFile:"), value)
}


// The block to use for reading the contents of the file handle asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/readabilityhandler
func (f_ FileHandle) ReadabilityHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("readabilityHandler"))
	return rv
}


// The block to use for reading the contents of the file handle asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/readabilityhandler
func (f_ FileHandle) SetReadabilityHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setReadabilityHandler:"), value)
}


// The block to use for writing the contents of the file handle asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/writeabilityhandler
func (f_ FileHandle) WriteabilityHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("writeabilityHandler"))
	return rv
}


// The block to use for writing the contents of the file handle asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/writeabilityhandler
func (f_ FileHandle) SetWriteabilityHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setWriteabilityHandler:"), value)
}


// Currently unused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandlenotificationmonitormodes
func (f_ FileHandle) NSFileHandleNotificationMonitorModes() IString {
	rv := objc.Send[String](f_.ID, objc.Sel("NSFileHandleNotificationMonitorModes"))
	return rv
}



