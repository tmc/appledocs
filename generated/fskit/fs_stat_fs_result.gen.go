// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FSStatFSResult] class.
var (
	FSStatFSResultClass     _FSStatFSResultClass
	FSStatFSResultClassOnce sync.Once
)

func getFSStatFSResultClass() _FSStatFSResultClass {
	FSStatFSResultClassOnce.Do(func() {
		FSStatFSResultClass = _FSStatFSResultClass{objc.GetClass("FSStatFSResult")}
	})
	return FSStatFSResultClass
}

type _FSStatFSResultClass struct {
	class objc.Class
}

// An interface definition for the [FSStatFSResult] class.
type IFSStatFSResult interface {
	objectivec.IObject
	AvailableBlocks() uint64
	SetAvailableBlocks(value uint64)
	AvailableBytes() uint64
	SetAvailableBytes(value uint64)
	BlockSize() int
	SetBlockSize(value int)
	FileSystemSubType() int
	SetFileSystemSubType(value int)
	FileSystemTypeName() string
	SetFileSystemTypeName(value string)
	FreeBlocks() uint64
	SetFreeBlocks(value uint64)
	FreeBytes() uint64
	SetFreeBytes(value uint64)
	FreeFiles() uint64
	SetFreeFiles(value uint64)
	IoSize() int
	SetIoSize(value int)
	TotalBlocks() uint64
	SetTotalBlocks(value uint64)
	TotalBytes() uint64
	SetTotalBytes(value uint64)
	TotalFiles() uint64
	SetTotalFiles(value uint64)
	UsedBlocks() uint64
	SetUsedBlocks(value uint64)
	UsedBytes() uint64
	SetUsedBytes(value uint64)
	SupportedVolumeCapabilities() FSVolumeSupportedCapabilities
	SetSupportedVolumeCapabilities(value IFSVolumeSupportedCapabilities)
	VolumeStatistics() FSStatFSResult
	SetVolumeStatistics(value IFSStatFSResult)
}

// A type used to report a volume’s statistics.
//
// The names of this type’s properties match those in the structure in , which reports these values for an FSKit file system. All numeric properties default to . Override these values, unless a given property has no meaningful value to provide. For the read-only , set this value with the designated initializer.


// A type used to report a volume’s statistics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSStatFSResult

type FSStatFSResult struct {
	objectivec.Object
}

// FSStatFSResultFrom constructs a [FSStatFSResult] from an unsafe.Pointer.
//
// A type used to report a volume’s statistics.
func FSStatFSResultFrom(ptr unsafe.Pointer) FSStatFSResult {
	return FSStatFSResult{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FSStatFSResultClass) Alloc() FSStatFSResult {
	rv := objc.Send[FSStatFSResult](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FSStatFSResultClass) New() FSStatFSResult {
	rv := objc.Send[FSStatFSResult](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSStatFSResult) Init() FSStatFSResult {
	rv := objc.Send[FSStatFSResult](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSStatFSResult) Autorelease() FSStatFSResult {
	rv := objc.Send[FSStatFSResult](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSStatFSResult creates a new FSStatFSResult instance.
func NewFSStatFSResult() FSStatFSResult {
	return getFSStatFSResultClass().New()
}



// A property for the number of free blocks available to a non-superuser on the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/availableblocks

func (f_ FSStatFSResult) AvailableBlocks() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("availableBlocks"))
	return rv
}


// A property for the number of free blocks available to a non-superuser on the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/availableblocks

func (f_ FSStatFSResult) SetAvailableBlocks(value uint64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAvailableBlocks:"), value)
}


// A property for the amount of space available to users, in bytes, in the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/availablebytes

func (f_ FSStatFSResult) AvailableBytes() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("availableBytes"))
	return rv
}


// A property for the amount of space available to users, in bytes, in the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/availablebytes

func (f_ FSStatFSResult) SetAvailableBytes(value uint64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAvailableBytes:"), value)
}


// A property for the volume’s block size, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/blocksize

func (f_ FSStatFSResult) BlockSize() int {
	rv := objc.Send[int](f_.ID, objc.Sel("blockSize"))
	return rv
}


// A property for the volume’s block size, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/blocksize

func (f_ FSStatFSResult) SetBlockSize(value int) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setBlockSize:"), value)
}


// A property for the file system’s subtype or flavor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/filesystemsubtype

func (f_ FSStatFSResult) FileSystemSubType() int {
	rv := objc.Send[int](f_.ID, objc.Sel("fileSystemSubType"))
	return rv
}


// A property for the file system’s subtype or flavor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/filesystemsubtype

func (f_ FSStatFSResult) SetFileSystemSubType(value int) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFileSystemSubType:"), value)
}


// A property for the file system type name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/filesystemtypename

func (f_ FSStatFSResult) FileSystemTypeName() string {
	rv := objc.Send[string](f_.ID, objc.Sel("fileSystemTypeName"))
	return rv
}


// A property for the file system type name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/filesystemtypename

func (f_ FSStatFSResult) SetFileSystemTypeName(value string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFileSystemTypeName:"), objc.String(value))
}


// A property for the number of free blocks in the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/freeblocks

func (f_ FSStatFSResult) FreeBlocks() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("freeBlocks"))
	return rv
}


// A property for the number of free blocks in the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/freeblocks

func (f_ FSStatFSResult) SetFreeBlocks(value uint64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFreeBlocks:"), value)
}


// A property for the amount of free space, in bytes, in the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/freebytes

func (f_ FSStatFSResult) FreeBytes() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("freeBytes"))
	return rv
}


// A property for the amount of free space, in bytes, in the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/freebytes

func (f_ FSStatFSResult) SetFreeBytes(value uint64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFreeBytes:"), value)
}


// A property for the total number of free file slots in the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/freefiles

func (f_ FSStatFSResult) FreeFiles() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("freeFiles"))
	return rv
}


// A property for the total number of free file slots in the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/freefiles

func (f_ FSStatFSResult) SetFreeFiles(value uint64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFreeFiles:"), value)
}


// A property for the optimal block size with which to perform I/O.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/iosize

func (f_ FSStatFSResult) IoSize() int {
	rv := objc.Send[int](f_.ID, objc.Sel("ioSize"))
	return rv
}


// A property for the optimal block size with which to perform I/O.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/iosize

func (f_ FSStatFSResult) SetIoSize(value int) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setIoSize:"), value)
}


// A property for the volume’s total data block count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/totalblocks

func (f_ FSStatFSResult) TotalBlocks() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("totalBlocks"))
	return rv
}


// A property for the volume’s total data block count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/totalblocks

func (f_ FSStatFSResult) SetTotalBlocks(value uint64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTotalBlocks:"), value)
}


// A property for the total size, in bytes, of the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/totalbytes

func (f_ FSStatFSResult) TotalBytes() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("totalBytes"))
	return rv
}


// A property for the total size, in bytes, of the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/totalbytes

func (f_ FSStatFSResult) SetTotalBytes(value uint64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTotalBytes:"), value)
}


// A property for the total number of file slots in the volume,
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/totalfiles

func (f_ FSStatFSResult) TotalFiles() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("totalFiles"))
	return rv
}


// A property for the total number of file slots in the volume,
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/totalfiles

func (f_ FSStatFSResult) SetTotalFiles(value uint64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setTotalFiles:"), value)
}


// A property for the number of used blocks in the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/usedblocks

func (f_ FSStatFSResult) UsedBlocks() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("usedBlocks"))
	return rv
}


// A property for the number of used blocks in the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/usedblocks

func (f_ FSStatFSResult) SetUsedBlocks(value uint64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setUsedBlocks:"), value)
}


// A property for the amount of used space, in bytes, in the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/usedbytes

func (f_ FSStatFSResult) UsedBytes() uint64 {
	rv := objc.Send[uint64](f_.ID, objc.Sel("usedBytes"))
	return rv
}


// A property for the amount of used space, in bytes, in the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsstatfsresult/usedbytes

func (f_ FSStatFSResult) SetUsedBytes(value uint64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setUsedBytes:"), value)
}


// A property that provides the supported capabilities of the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsvolume/operations/supportedvolumecapabilities

func (f_ FSStatFSResult) SupportedVolumeCapabilities() FSVolumeSupportedCapabilities {
	rv := objc.Send[FSVolumeSupportedCapabilities](f_.ID, objc.Sel("supportedVolumeCapabilities"))
	return rv
}


// A property that provides the supported capabilities of the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsvolume/operations/supportedvolumecapabilities

func (f_ FSStatFSResult) SetSupportedVolumeCapabilities(value IFSVolumeSupportedCapabilities) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSupportedVolumeCapabilities:"), value)
}


// A property that provides up-to-date statistics of the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsvolume/operations/volumestatistics

func (f_ FSStatFSResult) VolumeStatistics() FSStatFSResult {
	rv := objc.Send[FSStatFSResult](f_.ID, objc.Sel("volumeStatistics"))
	return rv
}


// A property that provides up-to-date statistics of the volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/fskit/fsvolume/operations/volumestatistics

func (f_ FSStatFSResult) SetVolumeStatistics(value IFSStatFSResult) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setVolumeStatistics:"), value)
}



