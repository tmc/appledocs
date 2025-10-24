// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class FSBlockDeviceResource */


/* debug [class_header]: Header for FSBlockDeviceResource */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FSBlockDeviceResource */
// An interface definition for the [FSBlockDeviceResource] class.
type IFSBlockDeviceResource interface {
	IFSResource
	
/* debug [class_interface_properties]: Properties for FSBlockDeviceResource */
	// properties:
	BlockCount() uint64
	BlockSize() uint64
	BSDName() objc.IObject /* cross-framework: NSString */
	Writable() bool
	PhysicalBlockSize() uint64
	IsWritable() bool
	SetIsWritable(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FSBlockDeviceResource */
	// methods:
	AsynchronousMetadataFlushWithError(error_ unsafe.Pointer) bool
	DelayedMetadataWriteFromStartingAtLengthError(buffer unsafe.Pointer, offset unsafe.Pointer, length uintptr /* not a class type */, error_ unsafe.Pointer) bool
	MetadataClearWithDelayedWritesError(rangesToClear []FSMetadataRange, withDelayedWrites bool, error_ unsafe.Pointer) bool
	MetadataFlushWithError(error_ unsafe.Pointer) bool
	MetadataPurgeError(rangesToPurge []FSMetadataRange, error_ unsafe.Pointer) bool
	MetadataReadIntoStartingAtLengthError(buffer unsafe.Pointer, offset unsafe.Pointer, length uintptr /* not a class type */, error_ unsafe.Pointer) bool
	MetadataWriteFromStartingAtLengthError(buffer unsafe.Pointer, offset unsafe.Pointer, length uintptr /* not a class type */, error_ unsafe.Pointer) bool
	ReadIntoStartingAtLengthCompletionHandler(buffer unsafe.Pointer, offset unsafe.Pointer, length uintptr /* not a class type */, completionHandler unsafe.Pointer)
	ReadIntoStartingAtLengthError(buffer unsafe.Pointer, offset unsafe.Pointer, length uintptr /* not a class type */, error_ unsafe.Pointer) uintptr /* not a class type */
	WriteFromStartingAtLengthCompletionHandler(buffer unsafe.Pointer, offset unsafe.Pointer, length uintptr /* not a class type */, completionHandler unsafe.Pointer)
	WriteFromStartingAtLengthError(buffer unsafe.Pointer, offset unsafe.Pointer, length uintptr /* not a class type */, error_ unsafe.Pointer) uintptr /* not a class type */
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FSBlockDeviceResource */
// Alloc allocates a new instance without initialization.
func (fc _FSBlockDeviceResourceClass) Alloc() FSBlockDeviceResource {
	rv := objc.Send[FSBlockDeviceResource](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FSBlockDeviceResource */
// A resource that represents a block storage disk partition.
//
// A can exist in either a proxied or nonproxied version. Only the daemon creates “real” (nonproxied) instances of this class. Client applications and daemons create proxy objects for requests, and opens the underlying device during the processing of the request. This class wraps a file descriptor for a disk device or partition. Its fundamental identifier is the BSD disk name ( ) for the underlying IOMedia object. However, doesn’t expose the underlying file descriptor. Instead, it provides accessor methods that can read from and write to the partition, either directly or using the kernel buffer cache. When you use a , your file system implementation also conforms to a maintenance operation protocol. These protocols add support for checking, repairing, and optionally formatting file systems. The system doesn’t mount block device file systems until they pass a file system check. For an that uses , conform to .


// A resource that represents a block storage disk partition.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FSBlockDeviceResource *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FSBlockDeviceResource */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FSBlockDeviceResource */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FSBlockDeviceResource */

// Asynchronously flushes the resource’s buffer cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/asynchronousMetadataFlush()
func (f_ FSBlockDeviceResource) AsynchronousMetadataFlushWithError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("asynchronousMetadataFlushWithError:"), error_)
	return rv
}/* debug [instance_methods/method]: AsynchronousMetadataFlushWithError */


// Writes file system metadata from a buffer to a cache, prior to flushing it to the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/delayedMetadataWriteFrom:startingAt:length:error:
func (f_ FSBlockDeviceResource) DelayedMetadataWriteFromStartingAtLengthError(buffer unsafe.Pointer, offset unsafe.Pointer, length uintptr /* not a class type */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("delayedMetadataWriteFrom:startingAt:length:error:"), buffer, offset, length, error_)
	return rv
}/* debug [instance_methods/method]: DelayedMetadataWriteFromStartingAtLengthError */


// Clears the given ranges within the buffer cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/metadataClear(_:withDelayedWrites:)
func (f_ FSBlockDeviceResource) MetadataClearWithDelayedWritesError(rangesToClear []FSMetadataRange, withDelayedWrites bool, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("metadataClear:withDelayedWrites:error:"), rangesToClear, withDelayedWrites, error_)
	return rv
}/* debug [instance_methods/method]: MetadataClearWithDelayedWritesError */


// Synchronously flushes the resource’s buffer cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/metadataFlush()
func (f_ FSBlockDeviceResource) MetadataFlushWithError(error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("metadataFlushWithError:"), error_)
	return rv
}/* debug [instance_methods/method]: MetadataFlushWithError */


// Synchronously purges the given ranges from the buffer cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/metadataPurge(_:)
func (f_ FSBlockDeviceResource) MetadataPurgeError(rangesToPurge []FSMetadataRange, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("metadataPurge:error:"), rangesToPurge, error_)
	return rv
}/* debug [instance_methods/method]: MetadataPurgeError */


// Synchronously reads file system metadata from the resource into a buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/metadataReadInto:startingAt:length:error:
func (f_ FSBlockDeviceResource) MetadataReadIntoStartingAtLengthError(buffer unsafe.Pointer, offset unsafe.Pointer, length uintptr /* not a class type */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("metadataReadInto:startingAt:length:error:"), buffer, offset, length, error_)
	return rv
}/* debug [instance_methods/method]: MetadataReadIntoStartingAtLengthError */


// Synchronously writes file system metadata from a buffer to the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/metadataWriteFrom:startingAt:length:error:
func (f_ FSBlockDeviceResource) MetadataWriteFromStartingAtLengthError(buffer unsafe.Pointer, offset unsafe.Pointer, length uintptr /* not a class type */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("metadataWriteFrom:startingAt:length:error:"), buffer, offset, length, error_)
	return rv
}/* debug [instance_methods/method]: MetadataWriteFromStartingAtLengthError */


// Reads data from the resource into a buffer and executes a block afterwards.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/readInto:startingAt:length:completionHandler:
func (f_ FSBlockDeviceResource) ReadIntoStartingAtLengthCompletionHandler(buffer unsafe.Pointer, offset unsafe.Pointer, length uintptr /* not a class type */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("readInto:startingAt:length:completionHandler:"), buffer, offset, length, completionHandler)
}/* debug [instance_methods/method]: ReadIntoStartingAtLengthCompletionHandler */


// Synchronously reads data from the resource into a buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/readInto:startingAt:length:error:
func (f_ FSBlockDeviceResource) ReadIntoStartingAtLengthError(buffer unsafe.Pointer, offset unsafe.Pointer, length uintptr /* not a class type */, error_ unsafe.Pointer) uintptr /* not a class type */ {
	rv := objc.Send[uintptr](f_.ID, objc.Sel("readInto:startingAt:length:error:"), buffer, offset, length, error_)
	return rv
}/* debug [instance_methods/method]: ReadIntoStartingAtLengthError */


// Writes data from from a buffer to the resource and executes a block afterwards.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/writeFrom:startingAt:length:completionHandler:
func (f_ FSBlockDeviceResource) WriteFromStartingAtLengthCompletionHandler(buffer unsafe.Pointer, offset unsafe.Pointer, length uintptr /* not a class type */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("writeFrom:startingAt:length:completionHandler:"), buffer, offset, length, completionHandler)
}/* debug [instance_methods/method]: WriteFromStartingAtLengthCompletionHandler */


// Synchronously writes data from from a buffer to the resource and executes a block afterwards.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/writeFrom:startingAt:length:error:
func (f_ FSBlockDeviceResource) WriteFromStartingAtLengthError(buffer unsafe.Pointer, offset unsafe.Pointer, length uintptr /* not a class type */, error_ unsafe.Pointer) uintptr /* not a class type */ {
	rv := objc.Send[uintptr](f_.ID, objc.Sel("writeFrom:startingAt:length:error:"), buffer, offset, length, error_)
	return rv
}/* debug [instance_methods/method]: WriteFromStartingAtLengthError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FSBlockDeviceResource */

// The block count on this resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/blockCount
func (f_ FSBlockDeviceResource) BlockCount() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("blockCount"))
	return rv
}/* debug [instance_properties/getter]: blockCount */


// The logical block size, the size of data blocks used by the file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/blockSize
func (f_ FSBlockDeviceResource) BlockSize() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("blockSize"))
	return rv
}/* debug [instance_properties/getter]: blockSize */


// The device name of the resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/bsdName
func (f_ FSBlockDeviceResource) BSDName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("BSDName"))
	return rv
}/* debug [instance_properties/getter]: BSDName */


// A Boolean property that indicates whether the resource can write data to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/isWritable
func (f_ FSBlockDeviceResource) Writable() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("writable"))
	return rv
}/* debug [instance_properties/getter]: writable */


// The sector size of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSBlockDeviceResource/physicalBlockSize
func (f_ FSBlockDeviceResource) PhysicalBlockSize() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("physicalBlockSize"))
	return rv
}/* debug [instance_properties/getter]: physicalBlockSize */


// A Boolean property that indicates whether the resource can write data to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsblockdeviceresource/iswritable
func (f_ FSBlockDeviceResource) IsWritable() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isWritable"))
	return rv
}/* debug [instance_properties/getter]: isWritable */


// A Boolean property that indicates whether the resource can write data to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsblockdeviceresource/iswritable
func (f_ FSBlockDeviceResource) SetIsWritable(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIsWritable:"), value)
}/* debug [instance_properties/setter]: isWritable */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class FSBlockDeviceResource */



