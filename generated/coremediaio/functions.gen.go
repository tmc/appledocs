// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

package coremediaio

/* debug [functions.gen.go]: Generating 35 functions for CoreMediaIO */
import (
	"unsafe"

	"github.com/ebitengine/purego"
	corevideo "github.com/tmc/appledocs/generated/corevideo"
)


// CoreMediaIO Functions (35 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CMIODeviceProcessAVCCommand func(IODeviceID, unsafe.Pointer) unsafe.Pointer
	_CMIODeviceProcessRS422Command func(IODeviceID, unsafe.Pointer) unsafe.Pointer
	_CMIODeviceStartStream func(IODeviceID, IOStreamID) unsafe.Pointer
	_CMIODeviceStopStream func(IODeviceID, IOStreamID) unsafe.Pointer
	_CMIOObjectAddPropertyListener func(IOObjectID, unsafe.Pointer, IOObjectPropertyListenerProc, unsafe.Pointer) unsafe.Pointer
	_CMIOObjectAddPropertyListenerBlock func(IOObjectID, unsafe.Pointer, unsafe.Pointer, IOObjectPropertyListenerBlock) unsafe.Pointer
	_CMIOObjectCreate func(IOHardwarePlugInRef, IOObjectID, IOClassID, unsafe.Pointer) unsafe.Pointer
	_CMIOObjectGetPropertyData func(IOObjectID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMIOObjectGetPropertyDataSize func(IOObjectID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMIOObjectHasProperty func(IOObjectID, unsafe.Pointer) unsafe.Pointer
	_CMIOObjectIsPropertySettable func(IOObjectID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMIOObjectPropertiesChanged func(IOHardwarePlugInRef, IOObjectID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMIOObjectRemovePropertyListener func(IOObjectID, unsafe.Pointer, IOObjectPropertyListenerProc, unsafe.Pointer) unsafe.Pointer
	_CMIOObjectRemovePropertyListenerBlock func(IOObjectID, unsafe.Pointer, unsafe.Pointer, IOObjectPropertyListenerBlock) unsafe.Pointer
	_CMIOObjectSetPropertyData func(IOObjectID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMIOObjectShow func(IOObjectID)
	_CMIOObjectsPublishedAndDied func(IOHardwarePlugInRef, IOObjectID, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMIOSampleBufferCopyNonRequiredAttachments func(SampleBufferRef, SampleBufferRef, AttachmentMode) unsafe.Pointer
	_CMIOSampleBufferCopySampleAttachments func(SampleBufferRef, SampleBufferRef) unsafe.Pointer
	_CMIOSampleBufferCreate func(AllocatorRef, BlockBufferRef, FormatDescriptionRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMIOSampleBufferCreateForImageBuffer func(AllocatorRef, ImageBufferRef, VideoFormatDescriptionRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMIOSampleBufferCreateNoDataMarker func(AllocatorRef, unsafe.Pointer, FormatDescriptionRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMIOSampleBufferGetDiscontinuityFlags func(SampleBufferRef) unsafe.Pointer
	_CMIOSampleBufferGetSequenceNumber func(SampleBufferRef) unsafe.Pointer
	_CMIOSampleBufferSetDiscontinuityFlags func(AllocatorRef, SampleBufferRef, unsafe.Pointer)
	_CMIOSampleBufferSetSequenceNumber func(AllocatorRef, SampleBufferRef, unsafe.Pointer)
	_CMIOStreamClockConvertHostTimeToDeviceTime func(unsafe.Pointer, TypeRef) corevideo.Time
	_CMIOStreamClockCreate func(AllocatorRef, StringRef, unsafe.Pointer, corevideo.Time, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMIOStreamClockInvalidate func(TypeRef) unsafe.Pointer
	_CMIOStreamClockPostTimingEvent func(corevideo.Time, unsafe.Pointer, unsafe.Pointer, TypeRef) unsafe.Pointer
	_CMIOStreamCopyBufferQueue func(IOStreamID, IODeviceStreamQueueAlteredProc, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMIOStreamDeckCueTo func(IOStreamID, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMIOStreamDeckJog func(IOStreamID, unsafe.Pointer) unsafe.Pointer
	_CMIOStreamDeckPlay func(IOStreamID) unsafe.Pointer
	_CMIOStreamDeckStop func(IOStreamID) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_CMIODeviceProcessAVCCommand, lib, "CMIODeviceProcessAVCCommand")
	tryRegister(&_CMIODeviceProcessRS422Command, lib, "CMIODeviceProcessRS422Command")
	tryRegister(&_CMIODeviceStartStream, lib, "CMIODeviceStartStream")
	tryRegister(&_CMIODeviceStopStream, lib, "CMIODeviceStopStream")
	tryRegister(&_CMIOObjectAddPropertyListener, lib, "CMIOObjectAddPropertyListener")
	tryRegister(&_CMIOObjectAddPropertyListenerBlock, lib, "CMIOObjectAddPropertyListenerBlock")
	tryRegister(&_CMIOObjectCreate, lib, "CMIOObjectCreate")
	tryRegister(&_CMIOObjectGetPropertyData, lib, "CMIOObjectGetPropertyData")
	tryRegister(&_CMIOObjectGetPropertyDataSize, lib, "CMIOObjectGetPropertyDataSize")
	tryRegister(&_CMIOObjectHasProperty, lib, "CMIOObjectHasProperty")
	tryRegister(&_CMIOObjectIsPropertySettable, lib, "CMIOObjectIsPropertySettable")
	tryRegister(&_CMIOObjectPropertiesChanged, lib, "CMIOObjectPropertiesChanged")
	tryRegister(&_CMIOObjectRemovePropertyListener, lib, "CMIOObjectRemovePropertyListener")
	tryRegister(&_CMIOObjectRemovePropertyListenerBlock, lib, "CMIOObjectRemovePropertyListenerBlock")
	tryRegister(&_CMIOObjectSetPropertyData, lib, "CMIOObjectSetPropertyData")
	tryRegister(&_CMIOObjectShow, lib, "CMIOObjectShow")
	tryRegister(&_CMIOObjectsPublishedAndDied, lib, "CMIOObjectsPublishedAndDied")
	tryRegister(&_CMIOSampleBufferCopyNonRequiredAttachments, lib, "CMIOSampleBufferCopyNonRequiredAttachments")
	tryRegister(&_CMIOSampleBufferCopySampleAttachments, lib, "CMIOSampleBufferCopySampleAttachments")
	tryRegister(&_CMIOSampleBufferCreate, lib, "CMIOSampleBufferCreate")
	tryRegister(&_CMIOSampleBufferCreateForImageBuffer, lib, "CMIOSampleBufferCreateForImageBuffer")
	tryRegister(&_CMIOSampleBufferCreateNoDataMarker, lib, "CMIOSampleBufferCreateNoDataMarker")
	tryRegister(&_CMIOSampleBufferGetDiscontinuityFlags, lib, "CMIOSampleBufferGetDiscontinuityFlags")
	tryRegister(&_CMIOSampleBufferGetSequenceNumber, lib, "CMIOSampleBufferGetSequenceNumber")
	tryRegister(&_CMIOSampleBufferSetDiscontinuityFlags, lib, "CMIOSampleBufferSetDiscontinuityFlags")
	tryRegister(&_CMIOSampleBufferSetSequenceNumber, lib, "CMIOSampleBufferSetSequenceNumber")
	tryRegister(&_CMIOStreamClockConvertHostTimeToDeviceTime, lib, "CMIOStreamClockConvertHostTimeToDeviceTime")
	tryRegister(&_CMIOStreamClockCreate, lib, "CMIOStreamClockCreate")
	tryRegister(&_CMIOStreamClockInvalidate, lib, "CMIOStreamClockInvalidate")
	tryRegister(&_CMIOStreamClockPostTimingEvent, lib, "CMIOStreamClockPostTimingEvent")
	tryRegister(&_CMIOStreamCopyBufferQueue, lib, "CMIOStreamCopyBufferQueue")
	tryRegister(&_CMIOStreamDeckCueTo, lib, "CMIOStreamDeckCueTo")
	tryRegister(&_CMIOStreamDeckJog, lib, "CMIOStreamDeckJog")
	tryRegister(&_CMIOStreamDeckPlay, lib, "CMIOStreamDeckPlay")
	tryRegister(&_CMIOStreamDeckStop, lib, "CMIOStreamDeckStop")
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



// CMIODeviceProcessAVCCommand is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceProcessAVCCommand(_:_:)
func CMIODeviceProcessAVCCommand(deviceID IODeviceID, ioAVCCommand unsafe.Pointer) unsafe.Pointer {
	return _CMIODeviceProcessAVCCommand(deviceID, ioAVCCommand)
}/* debug [functions.gen.go/function]: CMIODeviceProcessAVCCommand */

// CMIODeviceProcessRS422Command is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceProcessRS422Command(_:_:)
func CMIODeviceProcessRS422Command(deviceID IODeviceID, ioRS422Command unsafe.Pointer) unsafe.Pointer {
	return _CMIODeviceProcessRS422Command(deviceID, ioRS422Command)
}/* debug [functions.gen.go/function]: CMIODeviceProcessRS422Command */

// CMIODeviceStartStream is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceStartStream(_:_:)
func CMIODeviceStartStream(deviceID IODeviceID, streamID IOStreamID) unsafe.Pointer {
	return _CMIODeviceStartStream(deviceID, streamID)
}/* debug [functions.gen.go/function]: CMIODeviceStartStream */

// CMIODeviceStopStream is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceStopStream(_:_:)
func CMIODeviceStopStream(deviceID IODeviceID, streamID IOStreamID) unsafe.Pointer {
	return _CMIODeviceStopStream(deviceID, streamID)
}/* debug [functions.gen.go/function]: CMIODeviceStopStream */

// CMIOObjectAddPropertyListener is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectAddPropertyListener(_:_:_:_:)
func CMIOObjectAddPropertyListener(objectID IOObjectID, address unsafe.Pointer, listener IOObjectPropertyListenerProc, clientData unsafe.Pointer) unsafe.Pointer {
	return _CMIOObjectAddPropertyListener(objectID, address, listener, clientData)
}/* debug [functions.gen.go/function]: CMIOObjectAddPropertyListener */

// CMIOObjectAddPropertyListenerBlock is a CoreMediaIO function.
//
// Added in macOS 10.8.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectAddPropertyListenerBlock(_:_:_:_:)
func CMIOObjectAddPropertyListenerBlock(objectID IOObjectID, address unsafe.Pointer, dispatchQueue unsafe.Pointer, listener IOObjectPropertyListenerBlock) unsafe.Pointer {
	return _CMIOObjectAddPropertyListenerBlock(objectID, address, dispatchQueue, listener)
}/* debug [functions.gen.go/function]: CMIOObjectAddPropertyListenerBlock */

// CMIOObjectCreate is a CoreMediaIO function.
//
// Deprecated: This function was deprecated in macOS 12.3.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectCreate
func CMIOObjectCreate(owningPlugIn IOHardwarePlugInRef, owningObjectID IOObjectID, classID IOClassID, objectID unsafe.Pointer) unsafe.Pointer {
	return _CMIOObjectCreate(owningPlugIn, owningObjectID, classID, objectID)
}/* debug [functions.gen.go/function]: CMIOObjectCreate */

// CMIOObjectGetPropertyData is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectGetPropertyData(_:_:_:_:_:_:_:)
func CMIOObjectGetPropertyData(objectID IOObjectID, address unsafe.Pointer, qualifierDataSize unsafe.Pointer, qualifierData unsafe.Pointer, dataSize unsafe.Pointer, dataUsed unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _CMIOObjectGetPropertyData(objectID, address, qualifierDataSize, qualifierData, dataSize, dataUsed, data)
}/* debug [functions.gen.go/function]: CMIOObjectGetPropertyData */

// CMIOObjectGetPropertyDataSize is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectGetPropertyDataSize(_:_:_:_:_:)
func CMIOObjectGetPropertyDataSize(objectID IOObjectID, address unsafe.Pointer, qualifierDataSize unsafe.Pointer, qualifierData unsafe.Pointer, dataSize unsafe.Pointer) unsafe.Pointer {
	return _CMIOObjectGetPropertyDataSize(objectID, address, qualifierDataSize, qualifierData, dataSize)
}/* debug [functions.gen.go/function]: CMIOObjectGetPropertyDataSize */

// CMIOObjectHasProperty is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectHasProperty(_:_:)
func CMIOObjectHasProperty(objectID IOObjectID, address unsafe.Pointer) unsafe.Pointer {
	return _CMIOObjectHasProperty(objectID, address)
}/* debug [functions.gen.go/function]: CMIOObjectHasProperty */

// CMIOObjectIsPropertySettable is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectIsPropertySettable(_:_:_:)
func CMIOObjectIsPropertySettable(objectID IOObjectID, address unsafe.Pointer, isSettable unsafe.Pointer) unsafe.Pointer {
	return _CMIOObjectIsPropertySettable(objectID, address, isSettable)
}/* debug [functions.gen.go/function]: CMIOObjectIsPropertySettable */

// CMIOObjectPropertiesChanged is a CoreMediaIO function.
//
// Deprecated: This function was deprecated in macOS 12.3.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectPropertiesChanged
func CMIOObjectPropertiesChanged(owningPlugIn IOHardwarePlugInRef, objectID IOObjectID, numberAddresses unsafe.Pointer, addresses unsafe.Pointer) unsafe.Pointer {
	return _CMIOObjectPropertiesChanged(owningPlugIn, objectID, numberAddresses, addresses)
}/* debug [functions.gen.go/function]: CMIOObjectPropertiesChanged */

// CMIOObjectRemovePropertyListener is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectRemovePropertyListener(_:_:_:_:)
func CMIOObjectRemovePropertyListener(objectID IOObjectID, address unsafe.Pointer, listener IOObjectPropertyListenerProc, clientData unsafe.Pointer) unsafe.Pointer {
	return _CMIOObjectRemovePropertyListener(objectID, address, listener, clientData)
}/* debug [functions.gen.go/function]: CMIOObjectRemovePropertyListener */

// CMIOObjectRemovePropertyListenerBlock is a CoreMediaIO function.
//
// Added in macOS 10.8.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectRemovePropertyListenerBlock(_:_:_:_:)
func CMIOObjectRemovePropertyListenerBlock(objectID IOObjectID, address unsafe.Pointer, dispatchQueue unsafe.Pointer, listener IOObjectPropertyListenerBlock) unsafe.Pointer {
	return _CMIOObjectRemovePropertyListenerBlock(objectID, address, dispatchQueue, listener)
}/* debug [functions.gen.go/function]: CMIOObjectRemovePropertyListenerBlock */

// CMIOObjectSetPropertyData is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectSetPropertyData(_:_:_:_:_:_:)
func CMIOObjectSetPropertyData(objectID IOObjectID, address unsafe.Pointer, qualifierDataSize unsafe.Pointer, qualifierData unsafe.Pointer, dataSize unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _CMIOObjectSetPropertyData(objectID, address, qualifierDataSize, qualifierData, dataSize, data)
}/* debug [functions.gen.go/function]: CMIOObjectSetPropertyData */

// CMIOObjectShow is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectShow(_:)
func CMIOObjectShow(objectID IOObjectID) {
	_CMIOObjectShow(objectID)
}/* debug [functions.gen.go/function]: CMIOObjectShow */

// CMIOObjectsPublishedAndDied is a CoreMediaIO function.
//
// Deprecated: This function was deprecated in macOS 12.3.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectsPublishedAndDied
func CMIOObjectsPublishedAndDied(owningPlugIn IOHardwarePlugInRef, owningObjectID IOObjectID, numberPublishedCMIOObjects unsafe.Pointer, publishedCMIOObjects unsafe.Pointer, numberDeadCMIOObjects unsafe.Pointer, deadCMIOObjects unsafe.Pointer) unsafe.Pointer {
	return _CMIOObjectsPublishedAndDied(owningPlugIn, owningObjectID, numberPublishedCMIOObjects, publishedCMIOObjects, numberDeadCMIOObjects, deadCMIOObjects)
}/* debug [functions.gen.go/function]: CMIOObjectsPublishedAndDied */

// CMIOSampleBufferCopyNonRequiredAttachments is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferCopyNonRequiredAttachments
func CMIOSampleBufferCopyNonRequiredAttachments(sourceSBuf SampleBufferRef, destSBuf SampleBufferRef, attachmentMode AttachmentMode) unsafe.Pointer {
	return _CMIOSampleBufferCopyNonRequiredAttachments(sourceSBuf, destSBuf, attachmentMode)
}/* debug [functions.gen.go/function]: CMIOSampleBufferCopyNonRequiredAttachments */

// CMIOSampleBufferCopySampleAttachments is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferCopySampleAttachments
func CMIOSampleBufferCopySampleAttachments(sourceSBuf SampleBufferRef, destSBuf SampleBufferRef) unsafe.Pointer {
	return _CMIOSampleBufferCopySampleAttachments(sourceSBuf, destSBuf)
}/* debug [functions.gen.go/function]: CMIOSampleBufferCopySampleAttachments */

// CMIOSampleBufferCreate is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferCreate
func CMIOSampleBufferCreate(allocator AllocatorRef, dataBuffer BlockBufferRef, formatDescription FormatDescriptionRef, numSamples unsafe.Pointer, numSampleTimingEntries unsafe.Pointer, sampleTimingArray unsafe.Pointer, numSampleSizeEntries unsafe.Pointer, sampleSizeArray unsafe.Pointer, sequenceNumber unsafe.Pointer, discontinuityFlags unsafe.Pointer, sBufOut unsafe.Pointer) unsafe.Pointer {
	return _CMIOSampleBufferCreate(allocator, dataBuffer, formatDescription, numSamples, numSampleTimingEntries, sampleTimingArray, numSampleSizeEntries, sampleSizeArray, sequenceNumber, discontinuityFlags, sBufOut)
}/* debug [functions.gen.go/function]: CMIOSampleBufferCreate */

// CMIOSampleBufferCreateForImageBuffer is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferCreateForImageBuffer
func CMIOSampleBufferCreateForImageBuffer(allocator AllocatorRef, imageBuffer ImageBufferRef, formatDescription VideoFormatDescriptionRef, sampleTiming unsafe.Pointer, sequenceNumber unsafe.Pointer, discontinuityFlags unsafe.Pointer, sBufOut unsafe.Pointer) unsafe.Pointer {
	return _CMIOSampleBufferCreateForImageBuffer(allocator, imageBuffer, formatDescription, sampleTiming, sequenceNumber, discontinuityFlags, sBufOut)
}/* debug [functions.gen.go/function]: CMIOSampleBufferCreateForImageBuffer */

// CMIOSampleBufferCreateNoDataMarker is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferCreateNoDataMarker
func CMIOSampleBufferCreateNoDataMarker(allocator AllocatorRef, noDataEvent unsafe.Pointer, formatDescription FormatDescriptionRef, sequenceNumber unsafe.Pointer, discontinuityFlags unsafe.Pointer, sBufOut unsafe.Pointer) unsafe.Pointer {
	return _CMIOSampleBufferCreateNoDataMarker(allocator, noDataEvent, formatDescription, sequenceNumber, discontinuityFlags, sBufOut)
}/* debug [functions.gen.go/function]: CMIOSampleBufferCreateNoDataMarker */

// CMIOSampleBufferGetDiscontinuityFlags is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferGetDiscontinuityFlags
func CMIOSampleBufferGetDiscontinuityFlags(sbuf SampleBufferRef) unsafe.Pointer {
	return _CMIOSampleBufferGetDiscontinuityFlags(sbuf)
}/* debug [functions.gen.go/function]: CMIOSampleBufferGetDiscontinuityFlags */

// CMIOSampleBufferGetSequenceNumber is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferGetSequenceNumber
func CMIOSampleBufferGetSequenceNumber(sbuf SampleBufferRef) unsafe.Pointer {
	return _CMIOSampleBufferGetSequenceNumber(sbuf)
}/* debug [functions.gen.go/function]: CMIOSampleBufferGetSequenceNumber */

// CMIOSampleBufferSetDiscontinuityFlags is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferSetDiscontinuityFlags
func CMIOSampleBufferSetDiscontinuityFlags(allocator AllocatorRef, sbuf SampleBufferRef, discontinuityFlags unsafe.Pointer) {
	_CMIOSampleBufferSetDiscontinuityFlags(allocator, sbuf, discontinuityFlags)
}/* debug [functions.gen.go/function]: CMIOSampleBufferSetDiscontinuityFlags */

// CMIOSampleBufferSetSequenceNumber is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferSetSequenceNumber
func CMIOSampleBufferSetSequenceNumber(allocator AllocatorRef, sbuf SampleBufferRef, sequenceNumber unsafe.Pointer) {
	_CMIOSampleBufferSetSequenceNumber(allocator, sbuf, sequenceNumber)
}/* debug [functions.gen.go/function]: CMIOSampleBufferSetSequenceNumber */

// CMIOStreamClockConvertHostTimeToDeviceTime is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamClockConvertHostTimeToDeviceTime(_:_:)
func CMIOStreamClockConvertHostTimeToDeviceTime(hostTime unsafe.Pointer, clock TypeRef) corevideo.Time {
	return _CMIOStreamClockConvertHostTimeToDeviceTime(hostTime, clock)
}/* debug [functions.gen.go/function]: CMIOStreamClockConvertHostTimeToDeviceTime */

// CMIOStreamClockCreate is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamClockCreate(_:_:_:_:_:_:_:)
func CMIOStreamClockCreate(allocator AllocatorRef, clockName StringRef, sourceIdentifier unsafe.Pointer, getTimeCallMinimumInterval corevideo.Time, numberOfEventsForRateSmoothing unsafe.Pointer, numberOfAveragesForRateSmoothing unsafe.Pointer, clock unsafe.Pointer) unsafe.Pointer {
	return _CMIOStreamClockCreate(allocator, clockName, sourceIdentifier, getTimeCallMinimumInterval, numberOfEventsForRateSmoothing, numberOfAveragesForRateSmoothing, clock)
}/* debug [functions.gen.go/function]: CMIOStreamClockCreate */

// CMIOStreamClockInvalidate is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamClockInvalidate(_:)
func CMIOStreamClockInvalidate(clock TypeRef) unsafe.Pointer {
	return _CMIOStreamClockInvalidate(clock)
}/* debug [functions.gen.go/function]: CMIOStreamClockInvalidate */

// CMIOStreamClockPostTimingEvent is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamClockPostTimingEvent(_:_:_:_:)
func CMIOStreamClockPostTimingEvent(eventTime corevideo.Time, hostTime unsafe.Pointer, resynchronize unsafe.Pointer, clock TypeRef) unsafe.Pointer {
	return _CMIOStreamClockPostTimingEvent(eventTime, hostTime, resynchronize, clock)
}/* debug [functions.gen.go/function]: CMIOStreamClockPostTimingEvent */

// CMIOStreamCopyBufferQueue is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamCopyBufferQueue(_:_:_:_:)
func CMIOStreamCopyBufferQueue(streamID IOStreamID, queueAlteredProc IODeviceStreamQueueAlteredProc, queueAlteredRefCon unsafe.Pointer, queue unsafe.Pointer) unsafe.Pointer {
	return _CMIOStreamCopyBufferQueue(streamID, queueAlteredProc, queueAlteredRefCon, queue)
}/* debug [functions.gen.go/function]: CMIOStreamCopyBufferQueue */

// CMIOStreamDeckCueTo is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamDeckCueTo(_:_:_:)
func CMIOStreamDeckCueTo(streamID IOStreamID, frameNumber unsafe.Pointer, playOnCue unsafe.Pointer) unsafe.Pointer {
	return _CMIOStreamDeckCueTo(streamID, frameNumber, playOnCue)
}/* debug [functions.gen.go/function]: CMIOStreamDeckCueTo */

// CMIOStreamDeckJog is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamDeckJog(_:_:)
func CMIOStreamDeckJog(streamID IOStreamID, speed unsafe.Pointer) unsafe.Pointer {
	return _CMIOStreamDeckJog(streamID, speed)
}/* debug [functions.gen.go/function]: CMIOStreamDeckJog */

// CMIOStreamDeckPlay is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamDeckPlay(_:)
func CMIOStreamDeckPlay(streamID IOStreamID) unsafe.Pointer {
	return _CMIOStreamDeckPlay(streamID)
}/* debug [functions.gen.go/function]: CMIOStreamDeckPlay */

// CMIOStreamDeckStop is a CoreMediaIO function.
//
// Added in macOS 10.7.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamDeckStop(_:)
func CMIOStreamDeckStop(streamID IOStreamID) unsafe.Pointer {
	return _CMIOStreamDeckStop(streamID)
}/* debug [functions.gen.go/function]: CMIOStreamDeckStop */




