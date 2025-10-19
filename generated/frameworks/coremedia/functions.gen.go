// Code generated from Apple documentation for CoreMedia. DO NOT EDIT.

package coremedia

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

// CoreMedia Functions (15 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CMBufferQueueCopyHead func(unsafe.Pointer) unsafe.Pointer
	_CMMemoryPoolGetTypeID func() unsafe.Pointer
	_CMMetadataDataTypeRegistryDataTypeConformsToDataType func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSampleBufferGetSampleAttachmentsArray func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSimpleQueueGetTypeID func() unsafe.Pointer
	_CMSwapBigEndianClosedCaptionDescriptionToHost func(unsafe.Pointer, uintptr) unsafe.Pointer
	_CMSyncConvertTime func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTagCollectionRemoveAllTags func(unsafe.Pointer) unsafe.Pointer
	_CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTaggedBufferGroupGetCMSampleBufferAtIndex func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTimeGetSeconds func(unsafe.Pointer) unsafe.Pointer
	_CMTimeMappingCopyDescription func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTimeRangeCopyAsDictionary func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTimebaseGetTimeAndRate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMTimebaseGetTypeID func() unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_CMBufferQueueCopyHead, lib, "CMBufferQueueCopyHead")
	tryRegister(&_CMMemoryPoolGetTypeID, lib, "CMMemoryPoolGetTypeID")
	tryRegister(&_CMMetadataDataTypeRegistryDataTypeConformsToDataType, lib, "CMMetadataDataTypeRegistryDataTypeConformsToDataType")
	tryRegister(&_CMSampleBufferGetSampleAttachmentsArray, lib, "CMSampleBufferGetSampleAttachmentsArray")
	tryRegister(&_CMSimpleQueueGetTypeID, lib, "CMSimpleQueueGetTypeID")
	tryRegister(&_CMSwapBigEndianClosedCaptionDescriptionToHost, lib, "CMSwapBigEndianClosedCaptionDescriptionToHost")
	tryRegister(&_CMSyncConvertTime, lib, "CMSyncConvertTime")
	tryRegister(&_CMTagCollectionRemoveAllTags, lib, "CMTagCollectionRemoveAllTags")
	tryRegister(&_CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions, lib, "CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions")
	tryRegister(&_CMTaggedBufferGroupGetCMSampleBufferAtIndex, lib, "CMTaggedBufferGroupGetCMSampleBufferAtIndex")
	tryRegister(&_CMTimeGetSeconds, lib, "CMTimeGetSeconds")
	tryRegister(&_CMTimeMappingCopyDescription, lib, "CMTimeMappingCopyDescription")
	tryRegister(&_CMTimeRangeCopyAsDictionary, lib, "CMTimeRangeCopyAsDictionary")
	tryRegister(&_CMTimebaseGetTimeAndRate, lib, "CMTimebaseGetTimeAndRate")
	tryRegister(&_CMTimebaseGetTypeID, lib, "CMTimebaseGetTypeID")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}


// CMBufferQueueCopyHead is a CoreMedia function. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: doc://com.apple.coremedia/documentation/CoreMedia/CMBufferQueueCopyHead(_:)
func CMBufferQueueCopyHead(queue unsafe.Pointer) unsafe.Pointer {
	return _CMBufferQueueCopyHead(queue)
	}


// Returns the type identifier of memory pool objects. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: doc://com.apple.coremedia/documentation/CoreMedia/CMMemoryPoolGetTypeID()
func CMMemoryPoolGetTypeID() unsafe.Pointer {
	return _CMMemoryPoolGetTypeID()
	}


// Returns a Boolean value that indicates whether a data type conforms to another data type. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: doc://com.apple.coremedia/documentation/CoreMedia/CMMetadataDataTypeRegistryDataTypeConformsToDataType(_:conformsTo:)
func CMMetadataDataTypeRegistryDataTypeConformsToDataType(dataType unsafe.Pointer, conformsToDataType unsafe.Pointer) unsafe.Pointer {
	return _CMMetadataDataTypeRegistryDataTypeConformsToDataType(dataType, conformsToDataType)
	}


// Retrieves an array of sample attachment dictionaries that represents each sample in a sample buffer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.coremedia/documentation/CoreMedia/CMSampleBufferGetSampleAttachmentsArray(_:createIfNecessary:)
func CMSampleBufferGetSampleAttachmentsArray(sbuf unsafe.Pointer, createIfNecessary unsafe.Pointer) unsafe.Pointer {
	return _CMSampleBufferGetSampleAttachmentsArray(sbuf, createIfNecessary)
	}


// Returns the type identifier of sample buffer objects. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.coremedia/documentation/CoreMedia/CMSimpleQueueGetTypeID()
func CMSimpleQueueGetTypeID() unsafe.Pointer {
	return _CMSimpleQueueGetTypeID()
	}


// Converts a closed caption description structure from big-endian to host-endian, in place. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: doc://com.apple.coremedia/documentation/CoreMedia/CMSwapBigEndianClosedCaptionDescriptionToHost(_:_:)
func CMSwapBigEndianClosedCaptionDescriptionToHost(closedCaptionDescriptionData unsafe.Pointer, closedCaptionDescriptionSize uintptr) unsafe.Pointer {
	return _CMSwapBigEndianClosedCaptionDescriptionToHost(closedCaptionDescriptionData, closedCaptionDescriptionSize)
	}


// Converts a time from one timebase or clock to another timebase or clock. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: doc://com.apple.coremedia/documentation/CoreMedia/CMSyncConvertTime(_:from:to:)
func CMSyncConvertTime(time unsafe.Pointer, fromClockOrTimebase unsafe.Pointer, toClockOrTimebase unsafe.Pointer) unsafe.Pointer {
	return _CMSyncConvertTime(time, fromClockOrTimebase, toClockOrTimebase)
	}


// Removes all tags from a collection. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: doc://com.apple.coremedia/documentation/CoreMedia/CMTagCollectionRemoveAllTags
func CMTagCollectionRemoveAllTags(tagCollection unsafe.Pointer) unsafe.Pointer {
	return _CMTagCollectionRemoveAllTags(tagCollection)
	}


// CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions is a CoreMedia function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: doc://com.apple.coremedia/documentation/CoreMedia/CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions
func CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions(allocator unsafe.Pointer, taggedBufferGroup unsafe.Pointer, extensions unsafe.Pointer, formatDescriptionOut unsafe.Pointer) unsafe.Pointer {
	return _CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions(allocator, taggedBufferGroup, extensions, formatDescriptionOut)
	}


// Gets the sample buffer at a given index in the buffer group. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: doc://com.apple.coremedia/documentation/CoreMedia/CMTaggedBufferGroupGetCMSampleBufferAtIndex
func CMTaggedBufferGroupGetCMSampleBufferAtIndex(group unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	return _CMTaggedBufferGroupGetCMSampleBufferAtIndex(group, index)
	}


// Returns a representation of the time in seconds. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.coremedia/documentation/CoreMedia/CMTimeGetSeconds(_:)
func CMTimeGetSeconds(time unsafe.Pointer) unsafe.Pointer {
	return _CMTimeGetSeconds(time)
	}


// Copies a string description of a time mapping. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: doc://com.apple.coremedia/documentation/CoreMedia/CMTimeMappingCopyDescription(allocator:mapping:)
func CMTimeMappingCopyDescription(allocator unsafe.Pointer, mapping unsafe.Pointer) unsafe.Pointer {
	return _CMTimeMappingCopyDescription(allocator, mapping)
	}


// Returns a dictionary representation of a time range. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.coremedia/documentation/CoreMedia/CMTimeRangeCopyAsDictionary(_:allocator:)
func CMTimeRangeCopyAsDictionary(range_ unsafe.Pointer, allocator unsafe.Pointer) unsafe.Pointer {
	return _CMTimeRangeCopyAsDictionary(range_, allocator)
	}


// Returns the current time and rate of a timebase. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: doc://com.apple.coremedia/documentation/CoreMedia/CMTimebaseGetTimeAndRate(_:timeOut:rateOut:)
func CMTimebaseGetTimeAndRate(timebase unsafe.Pointer, timeOut unsafe.Pointer, rateOut unsafe.Pointer) unsafe.Pointer {
	return _CMTimebaseGetTimeAndRate(timebase, timeOut, rateOut)
	}


// Returns the Core Foundation type identifier that identifies a timebase object. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: doc://com.apple.coremedia/documentation/CoreMedia/CMTimebaseGetTypeID()
func CMTimebaseGetTypeID() unsafe.Pointer {
	return _CMTimebaseGetTypeID()
	}



