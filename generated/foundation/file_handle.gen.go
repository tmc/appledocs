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
	fileHandleClass     _FileHandleClass
	fileHandleClassOnce sync.Once
)

func getFileHandleClass() _FileHandleClass {
	fileHandleClassOnce.Do(func() {
		fileHandleClass = _FileHandleClass{objc.GetClass("NSFileHandle")}
	})
	return fileHandleClass
}

type _FileHandleClass struct {
	class objc.Class
}

// An interface definition for the [FileHandle] class.
type IFileHandle interface {
	objectivec.IObject
	CloseFile()
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


// Disallows further access to the represented file or communications channel and signals end of file on communications channels that permit writing.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileHandle/closeFile()
func (f_ FileHandle) CloseFile() {
	objc.Send[objc.ID](f_.ID, objc.Sel("closeFile"))
}



