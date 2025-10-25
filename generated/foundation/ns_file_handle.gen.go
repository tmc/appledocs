// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSFileHandle */


/* debug [class_header]: Header for NSFileHandle */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FileHandle */
// An interface definition for the [FileHandle] class.
type IFileHandle interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FileHandle */
	// properties:
	AvailableData() IData
	OffsetInFile() uint64
	Bytes() objectivec.IObject
	SetBytes(value objectivec.IObject)
	FileDescriptor() objectivec.IObject
	SetFileDescriptor(value objectivec.IObject)
	ReadabilityHandler() objectivec.IObject
	SetReadabilityHandler(value objectivec.IObject)
	WriteabilityHandler() objectivec.IObject
	SetWriteabilityHandler(value objectivec.IObject)
	NSFileHandleNotificationMonitorModes() IString
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FileHandle */
	// methods:
	AcceptConnectionInBackgroundAndNotify()
	AcceptConnectionInBackgroundAndNotifyForModes(modes []string)
	CloseAndReturnError(error_ IError) bool
	ReadInBackgroundAndNotify()
	ReadInBackgroundAndNotifyForModes(modes []string)
	ReadToEndOfFileInBackgroundAndNotify()
	ReadToEndOfFileInBackgroundAndNotifyForModes(modes []string)
	SeekToOffsetError(offset uint64, error_ IError) bool
	SynchronizeAndReturnError(error_ IError) bool
	TruncateAtOffsetError(offset uint64, error_ IError) bool
	WaitForDataInBackgroundAndNotify()
	WaitForDataInBackgroundAndNotifyForModes(modes []string)
	GetOffsetError(offsetInFile objectivec.IObject, error_ IError) bool
	ReadDataToEndOfFileAndReturnError(error_ IError) IData
	ReadDataUpToLengthError(length uint, error_ IError) IData
	SeekToEndReturningOffsetError(offsetInFile objectivec.IObject, error_ IError) bool
	WriteDataError(data IData, error_ IError) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FileHandle */
// Alloc allocates a new instance without initialization.
func (fc _FileHandleClass) Alloc() FileHandle {
	rv := objc.Send[FileHandle](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FileHandle */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FileHandle */

// Returns a file handle initialized for reading the file, device, or named socket at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(forReadingAtPath:)
func NewFileHandleForReadingAtPath(path IString) FileHandle {
	rv := objc.Send[FileHandle](objc.ID(getFileHandleClass().class), objc.Sel("fileHandleForReadingAtPath:"), path)
	return rv
}/* debug [class_init_methods/constructor]: NewFileHandleForReadingAtPath */


// Returns a file handle initialized for reading the file, device, or named socket at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(forReadingFromURL:)
func NewFileHandleForReadingFromURLError(url IURL, error_ IError) FileHandle {
	rv := objc.Send[FileHandle](objc.ID(getFileHandleClass().class), objc.Sel("fileHandleForReadingFromURL:error:"), url, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewFileHandleForReadingFromURLError */


// Returns a file handle initialized for reading and writing to the file, device, or named socket at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(forUpdatingAtPath:)
func NewFileHandleForUpdatingAtPath(path IString) FileHandle {
	rv := objc.Send[FileHandle](objc.ID(getFileHandleClass().class), objc.Sel("fileHandleForUpdatingAtPath:"), path)
	return rv
}/* debug [class_init_methods/constructor]: NewFileHandleForUpdatingAtPath */


// Returns a file handle initialized for reading and writing to the file, device, or named socket at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(forUpdatingURL:)
func NewFileHandleForUpdatingURLError(url IURL, error_ IError) FileHandle {
	rv := objc.Send[FileHandle](objc.ID(getFileHandleClass().class), objc.Sel("fileHandleForUpdatingURL:error:"), url, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewFileHandleForUpdatingURLError */


// Returns a file handle initialized for writing to the file, device, or named socket at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(forWritingAtPath:)
func NewFileHandleForWritingAtPath(path IString) FileHandle {
	rv := objc.Send[FileHandle](objc.ID(getFileHandleClass().class), objc.Sel("fileHandleForWritingAtPath:"), path)
	return rv
}/* debug [class_init_methods/constructor]: NewFileHandleForWritingAtPath */


// Returns a file handle initialized for writing to the file, device, or named socket at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(forWritingToURL:)
func NewFileHandleForWritingToURLError(url IURL, error_ IError) FileHandle {
	rv := objc.Send[FileHandle](objc.ID(getFileHandleClass().class), objc.Sel("fileHandleForWritingToURL:error:"), url, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewFileHandleForWritingToURLError */


// Returns a file handle initialized from data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(coder:)
func NewFileHandleWithCoder(coder ICoder) FileHandle {
	instance := getFileHandleClass().Alloc()
	rv := objc.Send[FileHandle](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFileHandleWithCoder */


// Creates and returns a file handle object associated with the specified file descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(fileDescriptor:)
func NewFileHandleWithFileDescriptor(fd int) FileHandle {
	instance := getFileHandleClass().Alloc()
	rv := objc.Send[FileHandle](instance.ID, objc.Sel("initWithFileDescriptor:"), fd)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFileHandleWithFileDescriptor */


// Creates and returns a file handle object associated with the specified file descriptor and deallocation policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(fileDescriptor:closeOnDealloc:)
func NewFileHandleWithFileDescriptorCloseOnDealloc(fd int, closeopt bool) FileHandle {
	instance := getFileHandleClass().Alloc()
	rv := objc.Send[FileHandle](instance.ID, objc.Sel("initWithFileDescriptor:closeOnDealloc:"), fd, closeopt)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFileHandleWithFileDescriptorCloseOnDealloc */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FileHandle */

// Returns a file handle initialized for reading the file, device, or named socket at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(forReadingAtPath:)
func (fc _FileHandleClass) FileHandleForReadingAtPath(path IString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("fileHandleForReadingAtPath:"), path)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FileHandleForReadingAtPath) */


// Returns a file handle initialized for reading the file, device, or named socket at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(forReadingFromURL:)
func (fc _FileHandleClass) FileHandleForReadingFromURLError(url IURL, error_ IError) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("fileHandleForReadingFromURL:error:"), url, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FileHandleForReadingFromURLError) */


// Returns a file handle initialized for reading and writing to the file, device, or named socket at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(forUpdatingAtPath:)
func (fc _FileHandleClass) FileHandleForUpdatingAtPath(path IString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("fileHandleForUpdatingAtPath:"), path)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FileHandleForUpdatingAtPath) */


// Returns a file handle initialized for reading and writing to the file, device, or named socket at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(forUpdatingURL:)
func (fc _FileHandleClass) FileHandleForUpdatingURLError(url IURL, error_ IError) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("fileHandleForUpdatingURL:error:"), url, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FileHandleForUpdatingURLError) */


// Returns a file handle initialized for writing to the file, device, or named socket at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(forWritingAtPath:)
func (fc _FileHandleClass) FileHandleForWritingAtPath(path IString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("fileHandleForWritingAtPath:"), path)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FileHandleForWritingAtPath) */


// Returns a file handle initialized for writing to the file, device, or named socket at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(forWritingToURL:)
func (fc _FileHandleClass) FileHandleForWritingToURLError(url IURL, error_ IError) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("fileHandleForWritingToURL:error:"), url, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FileHandleForWritingToURLError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FileHandle */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FileHandle */

// Accepts a socket connection (for stream-type sockets only) in the background and creates a file handle for the “near” (client) end of the communications channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/acceptConnectionInBackgroundAndNotify()
func (f_ FileHandle) AcceptConnectionInBackgroundAndNotify() {
	objc.Send[objc.ID](f_.ID, objc.Sel("acceptConnectionInBackgroundAndNotify"))
}/* debug [instance_methods/method]: AcceptConnectionInBackgroundAndNotify */


// Accepts a socket connection (for stream-type sockets only) in the background and creates a file handle for the “near” (client) end of the communications channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/acceptConnectionInBackgroundAndNotify(forModes:)
func (f_ FileHandle) AcceptConnectionInBackgroundAndNotifyForModes(modes []string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("acceptConnectionInBackgroundAndNotifyForModes:"), modes)
}/* debug [instance_methods/method]: AcceptConnectionInBackgroundAndNotifyForModes */


// Disallows further access to the represented file or communications channel and signals end of file on communications channels that permit writing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/close()
func (f_ FileHandle) CloseAndReturnError(error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("closeAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: CloseAndReturnError */


// Reads from the file or communications channel in the background and posts a notification when finished.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/readInBackgroundAndNotify()
func (f_ FileHandle) ReadInBackgroundAndNotify() {
	objc.Send[objc.ID](f_.ID, objc.Sel("readInBackgroundAndNotify"))
}/* debug [instance_methods/method]: ReadInBackgroundAndNotify */


// Reads from the file or communications channel in the background and posts a notification when finished.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/readInBackgroundAndNotify(forModes:)
func (f_ FileHandle) ReadInBackgroundAndNotifyForModes(modes []string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("readInBackgroundAndNotifyForModes:"), modes)
}/* debug [instance_methods/method]: ReadInBackgroundAndNotifyForModes */


// Reads to the end of file from the file or communications channel in the background and posts a notification when finished.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/readToEndOfFileInBackgroundAndNotify()
func (f_ FileHandle) ReadToEndOfFileInBackgroundAndNotify() {
	objc.Send[objc.ID](f_.ID, objc.Sel("readToEndOfFileInBackgroundAndNotify"))
}/* debug [instance_methods/method]: ReadToEndOfFileInBackgroundAndNotify */


// Reads to the end of file from the file or communications channel in the background and posts a notification when finished.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/readToEndOfFileInBackgroundAndNotify(forModes:)
func (f_ FileHandle) ReadToEndOfFileInBackgroundAndNotifyForModes(modes []string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("readToEndOfFileInBackgroundAndNotifyForModes:"), modes)
}/* debug [instance_methods/method]: ReadToEndOfFileInBackgroundAndNotifyForModes */


// Moves the file pointer to the specified offset within the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/seek(toOffset:)
func (f_ FileHandle) SeekToOffsetError(offset uint64, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("seekToOffset:error:"), offset, error_)
	return rv
}/* debug [instance_methods/method]: SeekToOffsetError */


// Causes all in-memory data and attributes of the file represented by the file handle to write to permanent storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/synchronize()
func (f_ FileHandle) SynchronizeAndReturnError(error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("synchronizeAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: SynchronizeAndReturnError */


// Truncates or extends the file represented by the file handle to a specified offset within the file and puts the file pointer at that position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/truncate(atOffset:)
func (f_ FileHandle) TruncateAtOffsetError(offset uint64, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("truncateAtOffset:error:"), offset, error_)
	return rv
}/* debug [instance_methods/method]: TruncateAtOffsetError */


// Asynchronously checks to see if data is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/waitForDataInBackgroundAndNotify()
func (f_ FileHandle) WaitForDataInBackgroundAndNotify() {
	objc.Send[objc.ID](f_.ID, objc.Sel("waitForDataInBackgroundAndNotify"))
}/* debug [instance_methods/method]: WaitForDataInBackgroundAndNotify */


// Asynchronously checks to see if data is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/waitForDataInBackgroundAndNotify(forModes:)
func (f_ FileHandle) WaitForDataInBackgroundAndNotifyForModes(modes []string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("waitForDataInBackgroundAndNotifyForModes:"), modes)
}/* debug [instance_methods/method]: WaitForDataInBackgroundAndNotifyForModes */


// Get the current position of the file pointer within the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileHandle/getOffset:error:
func (f_ FileHandle) GetOffsetError(offsetInFile objectivec.IObject, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("getOffset:error:"), offsetInFile, error_)
	return rv
}/* debug [instance_methods/method]: GetOffsetError */


// Reads the available data synchronously up to the end of file or maximum number of bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileHandle/readDataToEndOfFileAndReturnError:
func (f_ FileHandle) ReadDataToEndOfFileAndReturnError(error_ IError) IData {
	rv := objc.Send[Data](f_.ID, objc.Sel("readDataToEndOfFileAndReturnError:"), error_)
	return rv
}/* debug [instance_methods/method]: ReadDataToEndOfFileAndReturnError */


// Reads data synchronously up to the specified number of bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileHandle/readDataUpToLength:error:
func (f_ FileHandle) ReadDataUpToLengthError(length uint, error_ IError) IData {
	rv := objc.Send[Data](f_.ID, objc.Sel("readDataUpToLength:error:"), length, error_)
	return rv
}/* debug [instance_methods/method]: ReadDataUpToLengthError */


// Places the file pointer at the end of the file referenced by the file handle and returns the new file offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileHandle/seekToEndReturningOffset:error:
func (f_ FileHandle) SeekToEndReturningOffsetError(offsetInFile objectivec.IObject, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("seekToEndReturningOffset:error:"), offsetInFile, error_)
	return rv
}/* debug [instance_methods/method]: SeekToEndReturningOffsetError */


// Writes the specified data synchronously to the file handle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileHandle/writeData:error:
func (f_ FileHandle) WriteDataError(data IData, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("writeData:error:"), data, error_)
	return rv
}/* debug [instance_methods/method]: WriteDataError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FileHandle */

// The data currently available in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/availableData
func (f_ FileHandle) AvailableData() IData {
	rv := objc.Send[Data](f_.ID, objc.Sel("availableData"))
	return rv
}/* debug [instance_properties/getter]: availableData */


// The position of the file pointer within the file represented by the file handle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/offsetInFile
func (f_ FileHandle) OffsetInFile() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("offsetInFile"))
	return rv
}/* debug [instance_properties/getter]: offsetInFile */


// The file’s contents, as an asynchronous sequence of bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/bytes
func (f_ FileHandle) Bytes() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("bytes"))
	return rv
}/* debug [instance_properties/getter]: bytes */


// The file’s contents, as an asynchronous sequence of bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/bytes
func (f_ FileHandle) SetBytes(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setBytes:"), value)
}/* debug [instance_properties/setter]: bytes */


// The POSIX file descriptor associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/filedescriptor
func (f_ FileHandle) FileDescriptor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("fileDescriptor"))
	return rv
}/* debug [instance_properties/getter]: fileDescriptor */


// The POSIX file descriptor associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/filedescriptor
func (f_ FileHandle) SetFileDescriptor(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFileDescriptor:"), value)
}/* debug [instance_properties/setter]: fileDescriptor */


// The block to use for reading the contents of the file handle asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/readabilityhandler
func (f_ FileHandle) ReadabilityHandler() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("readabilityHandler"))
	return rv
}/* debug [instance_properties/getter]: readabilityHandler */


// The block to use for reading the contents of the file handle asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/readabilityhandler
func (f_ FileHandle) SetReadabilityHandler(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setReadabilityHandler:"), value)
}/* debug [instance_properties/setter]: readabilityHandler */


// The block to use for writing the contents of the file handle asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/writeabilityhandler
func (f_ FileHandle) WriteabilityHandler() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](f_.ID, objc.Sel("writeabilityHandler"))
	return rv
}/* debug [instance_properties/getter]: writeabilityHandler */


// The block to use for writing the contents of the file handle asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/writeabilityhandler
func (f_ FileHandle) SetWriteabilityHandler(value objectivec.IObject) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setWriteabilityHandler:"), value)
}/* debug [instance_properties/setter]: writeabilityHandler */


// Currently unused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandlenotificationmonitormodes
func (f_ FileHandle) NSFileHandleNotificationMonitorModes() IString {
	rv := objc.Send[String](f_.ID, objc.Sel("NSFileHandleNotificationMonitorModes"))
	return rv
}/* debug [instance_properties/getter]: NSFileHandleNotificationMonitorModes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSFileHandle */


