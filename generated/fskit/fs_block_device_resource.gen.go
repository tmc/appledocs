// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [FSBlockDeviceResource] class.
var (
	FSBlockDeviceResourceClass     _FSBlockDeviceResourceClass
	FSBlockDeviceResourceClassOnce sync.Once
)

func getFSBlockDeviceResourceClass() _FSBlockDeviceResourceClass {
	FSBlockDeviceResourceClassOnce.Do(func() {
		FSBlockDeviceResourceClass = _FSBlockDeviceResourceClass{objc.GetClass("FSBlockDeviceResource")}
	})
	return FSBlockDeviceResourceClass
}

type _FSBlockDeviceResourceClass struct {
	class objc.Class
}

// An interface definition for the [FSBlockDeviceResource] class.
type IFSBlockDeviceResource interface {
	IFSResource
	AsynchronousMetadataFlushWithError(error_ unsafe.Pointer) bool
	DelayedMetadataWriteFromStartingAtLengthError(buffer unsafe.Pointer, offset unsafe.Pointer, length unsafe.Pointer, error_ unsafe.Pointer) bool
	MetadataClearWithDelayedWritesError(rangesToClear unsafe.Pointer, withDelayedWrites bool, error_ unsafe.Pointer) bool
	MetadataFlushWithError(error_ unsafe.Pointer) bool
	MetadataPurgeError(rangesToPurge unsafe.Pointer, error_ unsafe.Pointer) bool
	MetadataReadIntoStartingAtLengthError(buffer unsafe.Pointer, offset unsafe.Pointer, length unsafe.Pointer, error_ unsafe.Pointer) bool
	MetadataWriteFromStartingAtLengthError(buffer unsafe.Pointer, offset unsafe.Pointer, length unsafe.Pointer, error_ unsafe.Pointer) bool
	ReadIntoStartingAtLengthCompletionHandler(buffer unsafe.Pointer, offset unsafe.Pointer, length unsafe.Pointer, completionHandler unsafe.Pointer)
	ReadIntoStartingAtLengthError(buffer unsafe.Pointer, offset unsafe.Pointer, length unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
	WriteFromStartingAtLengthCompletionHandler(buffer unsafe.Pointer, offset unsafe.Pointer, length unsafe.Pointer, completionHandler unsafe.Pointer)
	WriteFromStartingAtLengthError(buffer unsafe.Pointer, offset unsafe.Pointer, length unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
}

// A resource that represents a block storage disk partition.
//
// A can exist in either a proxied or nonproxied version. Only the daemon creates “real” (nonproxied) instances of this class. Client applications and daemons create proxy objects for requests, and opens the underlying device during the processing of the request. This class wraps a file descriptor for a disk device or partition. Its fundamental identifier is the BSD disk name ( ) for the underlying IOMedia object. However, doesn’t expose the underlying file descriptor. Instead, it provides accessor methods that can read from and write to the partition, either directly or using the kernel buffer cache. When you use a , your file system implementation also conforms to a maintenance operation protocol. These protocols add support for checking, repairing, and optionally formatting file systems. The system doesn’t mount block device file systems until they pass a file system check. For an that uses , conform to .
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource
type FSBlockDeviceResource struct {
	FSResource
}

// FSBlockDeviceResourceFrom constructs a [FSBlockDeviceResource] from an unsafe.Pointer.
//
// A resource that represents a block storage disk partition.
func FSBlockDeviceResourceFrom(ptr unsafe.Pointer) FSBlockDeviceResource {
	return FSBlockDeviceResource{
		FSResource: FSResourceFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (fc _FSBlockDeviceResourceClass) Alloc() FSBlockDeviceResource {
	rv := objc.Send[FSBlockDeviceResource](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FSBlockDeviceResourceClass) New() FSBlockDeviceResource {
	rv := objc.Send[FSBlockDeviceResource](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSBlockDeviceResource) Init() FSBlockDeviceResource {
	rv := objc.Send[FSBlockDeviceResource](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSBlockDeviceResource) Autorelease() FSBlockDeviceResource {
	rv := objc.Send[FSBlockDeviceResource](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSBlockDeviceResource creates a new FSBlockDeviceResource instance.
func NewFSBlockDeviceResource() FSBlockDeviceResource {
	return getFSBlockDeviceResourceClass().New()
}


// Asynchronously flushes the resource’s buffer cache.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/asynchronousMetadataFlush()
func (f_ FSBlockDeviceResource) AsynchronousMetadataFlushWithError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("asynchronousMetadataFlushWithError:"), error_)
	return rv
}

// Writes file system metadata from a buffer to a cache, prior to flushing it to the resource.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/delayedMetadataWriteFrom:startingAt:length:error:
func (f_ FSBlockDeviceResource) DelayedMetadataWriteFromStartingAtLengthError(buffer unsafe.Pointer, offset unsafe.Pointer, length unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("delayedMetadataWriteFrom:startingAt:length:error:"), buffer, offset, length, error_)
	return rv
}

// Clears the given ranges within the buffer cache.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/metadataClear(_:withDelayedWrites:)
func (f_ FSBlockDeviceResource) MetadataClearWithDelayedWritesError(rangesToClear unsafe.Pointer, withDelayedWrites bool, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("metadataClear:withDelayedWrites:error:"), rangesToClear, withDelayedWrites, error_)
	return rv
}

// Synchronously flushes the resource’s buffer cache.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/metadataFlush()
func (f_ FSBlockDeviceResource) MetadataFlushWithError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("metadataFlushWithError:"), error_)
	return rv
}

// Synchronously purges the given ranges from the buffer cache.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/metadataPurge(_:)
func (f_ FSBlockDeviceResource) MetadataPurgeError(rangesToPurge unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("metadataPurge:error:"), rangesToPurge, error_)
	return rv
}

// Synchronously reads file system metadata from the resource into a buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/metadataReadInto:startingAt:length:error:
func (f_ FSBlockDeviceResource) MetadataReadIntoStartingAtLengthError(buffer unsafe.Pointer, offset unsafe.Pointer, length unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("metadataReadInto:startingAt:length:error:"), buffer, offset, length, error_)
	return rv
}

// Synchronously writes file system metadata from a buffer to the resource.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/metadataWriteFrom:startingAt:length:error:
func (f_ FSBlockDeviceResource) MetadataWriteFromStartingAtLengthError(buffer unsafe.Pointer, offset unsafe.Pointer, length unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("metadataWriteFrom:startingAt:length:error:"), buffer, offset, length, error_)
	return rv
}

// Reads data from the resource into a buffer and executes a block afterwards.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/readInto:startingAt:length:completionHandler:
func (f_ FSBlockDeviceResource) ReadIntoStartingAtLengthCompletionHandler(buffer unsafe.Pointer, offset unsafe.Pointer, length unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("readInto:startingAt:length:completionHandler:"), buffer, offset, length, completionHandler)
}

// Synchronously reads data from the resource into a buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/readInto:startingAt:length:error:
func (f_ FSBlockDeviceResource) ReadIntoStartingAtLengthError(buffer unsafe.Pointer, offset unsafe.Pointer, length unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("readInto:startingAt:length:error:"), buffer, offset, length, error_)
	return rv
}

// Writes data from from a buffer to the resource and executes a block afterwards.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/writeFrom:startingAt:length:completionHandler:
func (f_ FSBlockDeviceResource) WriteFromStartingAtLengthCompletionHandler(buffer unsafe.Pointer, offset unsafe.Pointer, length unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("writeFrom:startingAt:length:completionHandler:"), buffer, offset, length, completionHandler)
}

// Synchronously writes data from from a buffer to the resource and executes a block afterwards.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/writeFrom:startingAt:length:error:
func (f_ FSBlockDeviceResource) WriteFromStartingAtLengthError(buffer unsafe.Pointer, offset unsafe.Pointer, length unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("writeFrom:startingAt:length:error:"), buffer, offset, length, error_)
	return rv
}

// The block count on this resource.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/blockCount
func (f_ FSBlockDeviceResource) BlockCount() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("blockCount"))
	return rv
}

// The logical block size, the size of data blocks used by the file system.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/blockSize
func (f_ FSBlockDeviceResource) BlockSize() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("blockSize"))
	return rv
}

// The device name of the resource.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/bsdName
func (f_ FSBlockDeviceResource) BSDName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("BSDName"))
	return rv
}

// A Boolean property that indicates whether the resource can write data to the device.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/isWritable
func (f_ FSBlockDeviceResource) Writable() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("writable"))
	return rv
}

// The sector size of the device.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/physicalBlockSize
func (f_ FSBlockDeviceResource) PhysicalBlockSize() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("physicalBlockSize"))
	return rv
}



