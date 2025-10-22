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
	CloseAndReturnError(error_ IError) bool
	CloseFile()
	ReadDataOfLength(length uint) Data
	ReadDataToEndOfFile() Data
	ReadInBackgroundAndNotifyForModes(modes []string)
	SeekToFileOffset(offset uint64)
	SeekToOffsetError(offset uint64, error_ IError) bool
	SeekToEndOfFile() uint64
	SynchronizeAndReturnError(error_ IError) bool
	SynchronizeFile()
	TruncateAtOffsetError(offset uint64, error_ IError) bool
	TruncateFileAtOffset(offset uint64)
	WriteData(data IData)
	GetOffsetError(offsetInFile unsafe.Pointer, error_ IError) bool
	ReadDataToEndOfFileAndReturnError(error_ IError) Data
	ReadDataUpToLengthError(length uint, error_ IError) Data
	SeekToEndReturningOffsetError(offsetInFile unsafe.Pointer, error_ IError) bool
	OffsetInFile() uint64
	ReadabilityHandler() unsafe.Pointer
	SetReadabilityHandler(value unsafe.Pointer)
	WriteabilityHandler() unsafe.Pointer
	SetWriteabilityHandler(value unsafe.Pointer)
	AvailableData() Data
	SetAvailableData(value IData)
	Bytes() unsafe.Pointer
	SetBytes(value unsafe.Pointer)
	FileDescriptor() unsafe.Pointer
	SetFileDescriptor(value unsafe.Pointer)
	NSFileHandleNotificationMonitorModes() string
}

// An object-oriented wrapper for a file descriptor.
//
// You use file handle objects to access data associated with files, sockets, pipes, and devices. For files, you can read, write, and seek within the file. For sockets, pipes, and devices, you can use a file handle object to monitor the device and process data asynchronously. Most creation methods for cause the file handle object to take ownership of the associated file descriptor. This means that the file handle object both creates the file descriptor and is responsible for closing it later, usually when the system deallocates the file handle object. If you want to use a file handle object with a file descriptor that you created, use the method or use the method and pass for the parameter.
//
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




// Returns a file handle initialized for reading the file, device, or named socket at the specified path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(forReadingAtPath:)
func NewFileHandleForReadingAtPath(path string) FileHandle {
	rv := objc.Send[FileHandle](objc.ID(getFileHandleClass().class), objc.Sel("fileHandleForReadingAtPath:"), objc.String(path))
	return rv
}



// Returns a file handle initialized for reading the file, device, or named socket at the specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(forReadingFromURL:)
func NewFileHandleForReadingFromURLError(url IURL, error_ IError) FileHandle {
	rv := objc.Send[FileHandle](objc.ID(getFileHandleClass().class), objc.Sel("fileHandleForReadingFromURL:error:"), url, error_)
	return rv
}



// Returns a file handle initialized for reading and writing to the file, device, or named socket at the specified path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(forUpdatingAtPath:)
func NewFileHandleForUpdatingAtPath(path string) FileHandle {
	rv := objc.Send[FileHandle](objc.ID(getFileHandleClass().class), objc.Sel("fileHandleForUpdatingAtPath:"), objc.String(path))
	return rv
}



// Returns a file handle initialized for reading and writing to the file, device, or named socket at the specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(forUpdatingURL:)
func NewFileHandleForUpdatingURLError(url IURL, error_ IError) FileHandle {
	rv := objc.Send[FileHandle](objc.ID(getFileHandleClass().class), objc.Sel("fileHandleForUpdatingURL:error:"), url, error_)
	return rv
}



// Returns a file handle initialized for writing to the file, device, or named socket at the specified path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(forWritingAtPath:)
func NewFileHandleForWritingAtPath(path string) FileHandle {
	rv := objc.Send[FileHandle](objc.ID(getFileHandleClass().class), objc.Sel("fileHandleForWritingAtPath:"), objc.String(path))
	return rv
}



// Returns a file handle initialized for writing to the file, device, or named socket at the specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(forWritingToURL:)
func NewFileHandleForWritingToURLError(url IURL, error_ IError) FileHandle {
	rv := objc.Send[FileHandle](objc.ID(getFileHandleClass().class), objc.Sel("fileHandleForWritingToURL:error:"), url, error_)
	return rv
}



// Returns a file handle initialized from data in an unarchiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(coder:)
func NewFileHandleWithCoder(coder ICoder) FileHandle {
	instance := getFileHandleClass().Alloc()
	rv := objc.Send[FileHandle](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}



// Creates and returns a file handle object associated with the specified file descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(fileDescriptor:)
func NewFileHandleWithFileDescriptor(fd int) FileHandle {
	instance := getFileHandleClass().Alloc()
	rv := objc.Send[FileHandle](instance.ID, objc.Sel("initWithFileDescriptor:"), fd)
	rv.Autorelease()
	return rv
}



// Creates and returns a file handle object associated with the specified file descriptor and deallocation policy.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(fileDescriptor:closeOnDealloc:)
func NewFileHandleWithFileDescriptorCloseOnDealloc(fd int, closeopt bool) FileHandle {
	instance := getFileHandleClass().Alloc()
	rv := objc.Send[FileHandle](instance.ID, objc.Sel("initWithFileDescriptor:closeOnDealloc:"), fd, closeopt)
	rv.Autorelease()
	return rv
}


// Returns a file handle initialized for reading the file, device, or named socket at the specified path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(forReadingAtPath:)
func (fc _FileHandleClass) FileHandleForReadingAtPath(path string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("fileHandleForReadingAtPath:"), objc.String(path))
	return rv
}

// Returns a file handle initialized for reading the file, device, or named socket at the specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(forReadingFromURL:)
func (fc _FileHandleClass) FileHandleForReadingFromURLError(url IURL, error_ IError) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("fileHandleForReadingFromURL:error:"), url, error_)
	return rv
}

// Returns a file handle initialized for reading and writing to the file, device, or named socket at the specified path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(forUpdatingAtPath:)
func (fc _FileHandleClass) FileHandleForUpdatingAtPath(path string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("fileHandleForUpdatingAtPath:"), objc.String(path))
	return rv
}

// Returns a file handle initialized for reading and writing to the file, device, or named socket at the specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(forUpdatingURL:)
func (fc _FileHandleClass) FileHandleForUpdatingURLError(url IURL, error_ IError) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("fileHandleForUpdatingURL:error:"), url, error_)
	return rv
}

// Returns a file handle initialized for writing to the file, device, or named socket at the specified path.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(forWritingAtPath:)
func (fc _FileHandleClass) FileHandleForWritingAtPath(path string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("fileHandleForWritingAtPath:"), objc.String(path))
	return rv
}

// Returns a file handle initialized for writing to the file, device, or named socket at the specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/init(forWritingToURL:)
func (fc _FileHandleClass) FileHandleForWritingToURLError(url IURL, error_ IError) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("fileHandleForWritingToURL:error:"), url, error_)
	return rv
}

// The file handle associated with a null device.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/nullDevice
func (fc _FileHandleClass) FileHandleWithNullDevice() FileHandle {
	rv := objc.Send[NSFileHandle](objc.ID(fc.class), objc.Sel("fileHandleWithNullDevice"))
	return rv
}
// The file handle associated with the standard error file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/standardError
func (fc _FileHandleClass) FileHandleWithStandardError() FileHandle {
	rv := objc.Send[NSFileHandle](objc.ID(fc.class), objc.Sel("fileHandleWithStandardError"))
	return rv
}
// The file handle associated with the standard input file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/standardInput
func (fc _FileHandleClass) FileHandleWithStandardInput() FileHandle {
	rv := objc.Send[NSFileHandle](objc.ID(fc.class), objc.Sel("fileHandleWithStandardInput"))
	return rv
}
// Disallows further access to the represented file or communications channel and signals end of file on communications channels that permit writing.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/close()
func (f_ FileHandle) CloseAndReturnError(error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("closeAndReturnError:"), error_)
	return rv
}

// Disallows further access to the represented file or communications channel and signals end of file on communications channels that permit writing.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/closeFile()
func (f_ FileHandle) CloseFile() {
	objc.Send[objc.ID](f_.ID, objc.Sel("closeFile"))
}

// Reads data synchronously up to the specified number of bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/readData(ofLength:)
func (f_ FileHandle) ReadDataOfLength(length uint) Data {
	rv := objc.Send[Data](f_.ID, objc.Sel("readDataOfLength:"), length)
	return rv
}

// Reads the available data synchronously up to the end of file or maximum number of bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/readDataToEndOfFile()
func (f_ FileHandle) ReadDataToEndOfFile() Data {
	rv := objc.Send[Data](f_.ID, objc.Sel("readDataToEndOfFile"))
	return rv
}

// Reads from the file or communications channel in the background and posts a notification when finished.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/readInBackgroundAndNotify(forModes:)
func (f_ FileHandle) ReadInBackgroundAndNotifyForModes(modes []string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("readInBackgroundAndNotifyForModes:"), modes)
}

// Moves the file pointer to the specified offset within the file represented by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/seek(toFileOffset:)
func (f_ FileHandle) SeekToFileOffset(offset uint64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("seekToFileOffset:"), offset)
}

// Moves the file pointer to the specified offset within the file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/seek(toOffset:)
func (f_ FileHandle) SeekToOffsetError(offset uint64, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("seekToOffset:error:"), offset, error_)
	return rv
}

// Places the file pointer at the end of the file referenced by the file handle and returns the new file offset.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/seekToEndOfFile()
func (f_ FileHandle) SeekToEndOfFile() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("seekToEndOfFile"))
	return rv
}

// Causes all in-memory data and attributes of the file represented by the file handle to write to permanent storage.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/synchronize()
func (f_ FileHandle) SynchronizeAndReturnError(error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("synchronizeAndReturnError:"), error_)
	return rv
}

// Causes all in-memory data and attributes of the file represented by the handle to write to permanent storage.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/synchronizeFile()
func (f_ FileHandle) SynchronizeFile() {
	objc.Send[objc.ID](f_.ID, objc.Sel("synchronizeFile"))
}

// Truncates or extends the file represented by the file handle to a specified offset within the file and puts the file pointer at that position.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/truncate(atOffset:)
func (f_ FileHandle) TruncateAtOffsetError(offset uint64, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("truncateAtOffset:error:"), offset, error_)
	return rv
}

// Truncates or extends the file represented by the file handle to a specified offset within the file and puts the file pointer at that position.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/truncateFile(atOffset:)
func (f_ FileHandle) TruncateFileAtOffset(offset uint64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("truncateFileAtOffset:"), offset)
}

// Writes the specified data synchronously to the file handle.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/write(_:)
func (f_ FileHandle) WriteData(data IData) {
	objc.Send[objc.ID](f_.ID, objc.Sel("writeData:"), data)
}

// Get the current position of the file pointer within the file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileHandle/getOffset:error:
func (f_ FileHandle) GetOffsetError(offsetInFile unsafe.Pointer, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("getOffset:error:"), offsetInFile, error_)
	return rv
}

// Reads the available data synchronously up to the end of file or maximum number of bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileHandle/readDataToEndOfFileAndReturnError:
func (f_ FileHandle) ReadDataToEndOfFileAndReturnError(error_ IError) Data {
	rv := objc.Send[Data](f_.ID, objc.Sel("readDataToEndOfFileAndReturnError:"), error_)
	return rv
}

// Reads data synchronously up to the specified number of bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileHandle/readDataUpToLength:error:
func (f_ FileHandle) ReadDataUpToLengthError(length uint, error_ IError) Data {
	rv := objc.Send[Data](f_.ID, objc.Sel("readDataUpToLength:error:"), length, error_)
	return rv
}

// Places the file pointer at the end of the file referenced by the file handle and returns the new file offset.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFileHandle/seekToEndReturningOffset:error:
func (f_ FileHandle) SeekToEndReturningOffsetError(offsetInFile unsafe.Pointer, error_ IError) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("seekToEndReturningOffset:error:"), offsetInFile, error_)
	return rv
}

// The file handle associated with a null device.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/nullDevice
func (f_ FileHandle) FileHandleWithNullDevice() NSFileHandle {
	rv := objc.Send[NSFileHandle](f_.ID, objc.Sel("fileHandleWithNullDevice"))
	return rv
}

// The position of the file pointer within the file represented by the file handle.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/offsetInFile
func (f_ FileHandle) OffsetInFile() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("offsetInFile"))
	return rv
}

// The block to use for reading the contents of the file handle asynchronously.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/readabilityHandler
func (f_ FileHandle) ReadabilityHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("readabilityHandler"))
	return rv
}


// SetReadabilityHandler sets the value of the readabilityHandler property.
// The block to use for reading the contents of the file handle asynchronously.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/readabilityHandler
func (f_ FileHandle) SetReadabilityHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setReadabilityHandler:"), value)
}

// The file handle associated with the standard error file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/standardError
func (f_ FileHandle) FileHandleWithStandardError() NSFileHandle {
	rv := objc.Send[NSFileHandle](f_.ID, objc.Sel("fileHandleWithStandardError"))
	return rv
}

// The file handle associated with the standard input file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/standardInput
func (f_ FileHandle) FileHandleWithStandardInput() NSFileHandle {
	rv := objc.Send[NSFileHandle](f_.ID, objc.Sel("fileHandleWithStandardInput"))
	return rv
}

// The block to use for writing the contents of the file handle asynchronously.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/writeabilityHandler
func (f_ FileHandle) WriteabilityHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("writeabilityHandler"))
	return rv
}


// SetWriteabilityHandler sets the value of the writeabilityHandler property.
// The block to use for writing the contents of the file handle asynchronously.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/writeabilityHandler
func (f_ FileHandle) SetWriteabilityHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setWriteabilityHandler:"), value)
}

// The data currently available in the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/availabledata
func (f_ FileHandle) AvailableData() Data {
	rv := objc.Send[Data](f_.ID, objc.Sel("availableData"))
	return rv
}


// SetAvailableData sets the value of the availableData property.
// The data currently available in the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/availabledata
func (f_ FileHandle) SetAvailableData(value IData) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAvailableData:"), value)
}

// The file’s contents, as an asynchronous sequence of bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/bytes
func (f_ FileHandle) Bytes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("bytes"))
	return rv
}


// SetBytes sets the value of the bytes property.
// The file’s contents, as an asynchronous sequence of bytes.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/bytes
func (f_ FileHandle) SetBytes(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setBytes:"), value)
}

// The POSIX file descriptor associated with the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/filedescriptor
func (f_ FileHandle) FileDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("fileDescriptor"))
	return rv
}


// SetFileDescriptor sets the value of the fileDescriptor property.
// The POSIX file descriptor associated with the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/filehandle/filedescriptor
func (f_ FileHandle) SetFileDescriptor(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFileDescriptor:"), value)
}

// Currently unused.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilehandlenotificationmonitormodes
func (f_ FileHandle) NSFileHandleNotificationMonitorModes() string {
	rv := objc.Send[string](f_.ID, objc.Sel("NSFileHandleNotificationMonitorModes"))
	return rv
}


