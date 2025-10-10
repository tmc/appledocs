// Code generated from Apple documentation for CoreMedia. DO NOT EDIT.

package coremedia

// CoreMedia Functions
//
// This file contains function declarations discovered from Apple's documentation.
// To use these functions, you need to:
//   1. Map C types to Go types
//   2. Create function variables
//   3. Register them with purego.RegisterLibFunc
//
// Example:
//   var CGContextSetRGBFillColor func(c CGContextRef, red, green, blue, alpha CGFloat)
//   purego.RegisterLibFunc(&CGContextSetRGBFillColor, lib, "CGContextSetRGBFillColor")

// Discovered functions (382 total):

// CMAudioClockCreate(allocator:  CFAllocator?,  clockOut:  UnsafeMutablePointer< CMClock?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMAudioDeviceClockCreate(allocator:  CFAllocator?,  deviceUID:  CFString?,  clockOut:  UnsafeMutablePointer< CMClock?>) ->  OSStatus) func
//
// Availability:
//   - macOS 10.8+

// CMAudioDeviceClockCreateFromAudioDeviceID(allocator:  CFAllocator?,  deviceID:  AudioDeviceID,  clockOut:  UnsafeMutablePointer< CMClock?>) ->  OSStatus) func
//
// Availability:
//   - macOS 10.8+


// CMAudioDeviceClockGetAudioDevice(clock _, :  CMClock,  deviceUIDOut:  AutoreleasingUnsafeMutablePointer< CFString?>?,  deviceIDOut:  UnsafeMutablePointer< AudioDeviceID>?,  trackingDefaultDeviceOut:  UnsafeMutablePointer< DarwinBoolean>?) ->  OSStatus) func
//
// Availability:
//   - macOS 10.8+

// CMAudioDeviceClockSetAudioDeviceID(clock _, :  CMClock,  deviceID:  AudioDeviceID) ->  OSStatus) func
//
// Availability:
//   - macOS 10.8+

// CMAudioDeviceClockSetAudioDeviceUID(clock _, :  CMClock,  deviceUID:  CFString?) ->  OSStatus) func
//
// Availability:
//   - macOS 10.8+


// CMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer(allocator:  CFAllocator?,  audioFormatDescription:  CMAudioFormatDescription,  flavor:  CMSoundDescriptionFlavor?,  blockBufferOut:  UnsafeMutablePointer< CMBlockBuffer?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMAudioFormatDescriptionCreate(allocator:  CFAllocator?,  asbd:  UnsafePointer< AudioStreamBasicDescription>,  layoutSize:  Int,  layout:  UnsafePointer< AudioChannelLayout>?,  magicCookieSize:  Int,  magicCookie:  UnsafeRawPointer?,  extensions:  CFDictionary?,  formatDescriptionOut:  UnsafeMutablePointer< CMAudioFormatDescription?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer(soundDescriptionBlockBuffer allocator:  CFAllocator?,  bigEndianSoundDescriptionBlockBuffer, :  CMBlockBuffer,  flavor:  CMSoundDescriptionFlavor?,  formatDescriptionOut:  UnsafeMutablePointer< CMAudioFormatDescription?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData(soundDescriptionData allocator:  CFAllocator?,  bigEndianSoundDescriptionData, :  UnsafePointer< UInt8>,  size:  Int,  flavor:  CMSoundDescriptionFlavor?,  formatDescriptionOut:  UnsafeMutablePointer< CMAudioFormatDescription?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMAudioFormatDescriptionCreateSummary(allocator:  CFAllocator?,  formatDescriptionArray:  CFArray,  flags:  UInt32,  formatDescriptionOut:  UnsafeMutablePointer< CMAudioFormatDescription?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMAudioFormatDescriptionEqual(formatDescription _, :  CMAudioFormatDescription,  otherFormatDescription:  CMAudioFormatDescription,  equalityMask:  CMAudioFormatDescriptionMask,  equalityMaskOut:  UnsafeMutablePointer< CMAudioFormatDescriptionMask>?) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMAudioFormatDescriptionGetChannelLayout(desc _, :  CMAudioFormatDescription,  sizeOut:  UnsafeMutablePointer< Int>?) ->  UnsafePointer< AudioChannelLayout>?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMAudioFormatDescriptionGetFormatList(desc _, :  CMAudioFormatDescription,  sizeOut:  UnsafeMutablePointer< Int>?) ->  UnsafePointer< AudioFormatListItem>?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMAudioFormatDescriptionGetMagicCookie(desc _, :  CMAudioFormatDescription,  sizeOut:  UnsafeMutablePointer< Int>?) ->  UnsafeRawPointer?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMAudioFormatDescriptionGetMostCompatibleFormat(desc _, :  CMAudioFormatDescription) ->  UnsafePointer< AudioFormatListItem>?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMAudioFormatDescriptionGetRichestDecodableFormat(desc _, :  CMAudioFormatDescription) ->  UnsafePointer< AudioFormatListItem>?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMAudioFormatDescriptionGetStreamBasicDescription(desc _, :  CMAudioFormatDescription) ->  UnsafePointer< AudioStreamBasicDescription>?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMAudioSampleBufferCreateReadyWithPacketDescriptions(numSamples allocator:  CFAllocator?,  dataBuffer:  CMBlockBuffer,  formatDescription:  CMFormatDescription,  sampleCount, :  CMItemCount,  presentationTimeStamp:  CMTime,  packetDescriptions:  UnsafePointer< AudioStreamPacketDescription>?,  sampleBufferOut:  UnsafeMutablePointer< CMSampleBuffer?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMAudioSampleBufferCreateWithPacketDescriptions(makeDataReadyRefcon allocator:  CFAllocator?,  dataBuffer:  CMBlockBuffer?,  dataReady:  Bool,  makeDataReadyCallback:  CMSampleBufferMakeDataReadyCallback?,  refcon, numSamples :  UnsafeMutableRawPointer?,  formatDescription:  CMFormatDescription,  sampleCount, :  CMItemCount,  presentationTimeStamp:  CMTime,  packetDescriptions:  UnsafePointer< AudioStreamPacketDescription>?,  sampleBufferOut:  UnsafeMutablePointer< CMSampleBuffer?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler(allocator _, dataBuffer :  CFAllocator?,  _, dataReady :  CMBlockBuffer?,  _, formatDescription :  Bool,  _, numSamples :  CMFormatDescription,  _, presentationTimeStamp :  CMItemCount,  _, packetDescriptions :  CMTime,  _, sampleBufferOut :  UnsafePointer< AudioStreamPacketDescription>?,  _, makeDataReadyHandler :  UnsafeMutablePointer< CMSampleBuffer?>,  _, :  CMSampleBufferMakeDataReadyHandler?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.2+
//   - iPadOS 12.2+
//   - macOS 10.14.4+
//   - tvOS 12.2+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMBlockBufferAccessDataBytes(theBuffer _, offset :  CMBlockBuffer,  atOffset, :  Int,  length:  Int,  temporaryBlock:  UnsafeMutableRawPointer,  returnedPointerOut:  UnsafeMutablePointer< UnsafeMutablePointer< CChar>?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBlockBufferAppendBufferReference(theBuffer _, :  CMBlockBuffer,  targetBBuf:  CMBlockBuffer,  offsetToData:  Int,  dataLength:  Int,  flags:  CMBlockBufferFlags) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBlockBufferAppendMemoryBlock(theBuffer _, blockLength :  CMBlockBuffer,  memoryBlock:  UnsafeMutableRawPointer?,  length, :  Int,  blockAllocator:  CFAllocator?,  customBlockSource:  UnsafePointer< CMBlockBufferCustomBlockSource>?,  offsetToData:  Int,  dataLength:  Int,  flags:  CMBlockBufferFlags) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMBlockBufferAssureBlockMemory(theBuffer _, :  CMBlockBuffer) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBlockBufferCopyDataBytes(theSourceBuffer _, offsetToData :  CMBlockBuffer,  atOffset, :  Int,  dataLength:  Int,  destination:  UnsafeMutableRawPointer) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBlockBufferCreateContiguous(structureAllocator allocator, :  CFAllocator?,  sourceBuffer:  CMBlockBuffer,  blockAllocator:  CFAllocator?,  customBlockSource:  UnsafePointer< CMBlockBufferCustomBlockSource>?,  offsetToData:  Int,  dataLength:  Int,  flags:  CMBlockBufferFlags,  blockBufferOut:  UnsafeMutablePointer< CMBlockBuffer?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMBlockBufferCreateEmpty(structureAllocator allocator, subBlockCapacity :  CFAllocator?,  capacity, :  UInt32,  flags:  CMBlockBufferFlags,  blockBufferOut:  UnsafeMutablePointer< CMBlockBuffer?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBlockBufferCreateWithBufferReference(structureAllocator allocator, bufferReference :  CFAllocator?,  referenceBuffer, :  CMBlockBuffer,  offsetToData:  Int,  dataLength:  Int,  flags:  CMBlockBufferFlags,  blockBufferOut:  UnsafeMutablePointer< CMBlockBuffer?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBlockBufferCreateWithMemoryBlock(structureAllocator allocator, :  CFAllocator?,  memoryBlock:  UnsafeMutableRawPointer?,  blockLength:  Int,  blockAllocator:  CFAllocator?,  customBlockSource:  UnsafePointer< CMBlockBufferCustomBlockSource>?,  offsetToData:  Int,  dataLength:  Int,  flags:  CMBlockBufferFlags,  blockBufferOut:  UnsafeMutablePointer< CMBlockBuffer?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMBlockBufferFillDataBytes(fillByte with, destinationBuffer :  CChar,  blockBuffer, :  CMBlockBuffer,  offsetIntoDestination:  Int,  dataLength:  Int) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBlockBufferGetDataLength(theBuffer _, :  CMBlockBuffer) ->  Int) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBlockBufferGetDataPointer(theBuffer _, offset :  CMBlockBuffer,  atOffset, :  Int,  lengthAtOffsetOut:  UnsafeMutablePointer< Int>?,  totalLengthOut:  UnsafeMutablePointer< Int>?,  dataPointerOut:  UnsafeMutablePointer< UnsafeMutablePointer< CChar>?>?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMBlockBufferGetTypeID() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBlockBufferIsEmpty(theBuffer _, :  CMBlockBuffer) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBlockBufferIsRangeContiguous(theBuffer _, offset :  CMBlockBuffer,  atOffset, :  Int,  length:  Int) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMBlockBufferReplaceDataBytes(sourceBytes with, destinationBuffer :  UnsafeRawPointer,  blockBuffer, :  CMBlockBuffer,  offsetIntoDestination:  Int,  dataLength:  Int) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBufferQueueCallForEachBuffer(queue _, :  CMBufferQueue,  callback: ( CMBuffer,  UnsafeMutableRawPointer?) ->  OSStatus,  refcon:  UnsafeMutableRawPointer?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBufferQueueContainsEndOfData(queue _, :  CMBufferQueue) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMBufferQueueCopyHead(queue _, :  CMBufferQueue) ->  CMBuffer?) func
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMBufferQueueCreate(allocator:  CFAllocator?,  capacity:  CMItemCount,  callbacks:  UnsafePointer< CMBufferCallbacks>,  queueOut:  UnsafeMutablePointer< CMBufferQueue?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBufferQueueCreateWithHandlers(allocator _, capacity :  CFAllocator?,  _, handlers :  CMItemCount,  _, queueOut :  OpaquePointer,  _, :  UnsafeMutablePointer< CMBufferQueue?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.2+
//   - iPadOS 12.2+
//   - macOS 10.14.4+
//   - tvOS 12.2+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMBufferQueueDequeue(queue _, :  CMBufferQueue) ->  CMBuffer?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBufferQueueDequeueIfDataReady(queue _, :  CMBufferQueue) ->  CMBuffer?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBufferQueueEnqueue(queue _, buf :  CMBufferQueue,  buffer, :  CMBuffer) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMBufferQueueGetBufferCount(queue _, :  CMBufferQueue) ->  CMItemCount) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.3+
//   - iPadOS 4.3+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBufferQueueGetCallbacksForUnsortedSampleBuffers() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMBufferQueueGetDuration(queue _, :  CMBufferQueue) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBufferQueueGetEndPresentationTimeStamp(queue _, :  CMBufferQueue) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBufferQueueGetFirstDecodeTimeStamp(queue _, :  CMBufferQueue) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMBufferQueueGetFirstPresentationTimeStamp(queue _, :  CMBufferQueue) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBufferQueueGetHead(queue _, :  CMBufferQueue) ->  CMBuffer?) func
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 18.0)
//   - iOS 4.0+ (Deprecated in 18.0)
//   - iPadOS 4.0+ (Deprecated in 18.0)
//   - macOS 10.7+ (Deprecated in 15.0)
//   - tvOS 9.0+ (Deprecated in 18.0)
//   - visionOS 1.0+ (Deprecated in 2.0)
//   - watchOS 6.0+ (Deprecated in 11.0)
//
// Deprecated: This function is deprecated.

// CMBufferQueueGetMaxPresentationTimeStamp(queue _, :  CMBufferQueue) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMBufferQueueGetMinDecodeTimeStamp(queue _, :  CMBufferQueue) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBufferQueueGetMinPresentationTimeStamp(queue _, :  CMBufferQueue) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBufferQueueGetTotalSize(queue _, :  CMBufferQueue) ->  Int) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.1+
//   - iPadOS 7.1+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMBufferQueueGetTypeID() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBufferQueueInstallTrigger(queue _, :  CMBufferQueue,  callback:  CMBufferQueueTriggerCallback?,  refcon:  UnsafeMutableRawPointer?,  condition:  CMBufferQueueTriggerCondition,  time:  CMTime,  triggerTokenOut:  UnsafeMutablePointer< CMBufferQueueTriggerToken?>?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBufferQueueInstallTriggerHandler(queue _, condition :  CMBufferQueue,  _, time :  CMBufferQueueTriggerCondition,  _, triggerTokenOut :  CMTime,  _, handler :  UnsafeMutablePointer< CMBufferQueueTriggerToken?>?,  _, :  CMBufferQueueTriggerHandler?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.2+
//   - iPadOS 12.2+
//   - macOS 10.14.4+
//   - tvOS 12.2+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMBufferQueueInstallTriggerHandlerWithIntegerThreshold(queue _, condition :  CMBufferQueue,  _, threshold :  CMBufferQueueTriggerCondition,  _, triggerTokenOut :  CMItemCount,  _, handler :  UnsafeMutablePointer< CMBufferQueueTriggerToken?>?,  _, :  CMBufferQueueTriggerHandler?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.2+
//   - iPadOS 12.2+
//   - macOS 10.14.4+
//   - tvOS 12.2+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBufferQueueInstallTriggerWithIntegerThreshold(queue _, :  CMBufferQueue,  callback:  CMBufferQueueTriggerCallback?,  refcon:  UnsafeMutableRawPointer?,  condition:  CMBufferQueueTriggerCondition,  threshold:  CMItemCount,  triggerTokenOut:  UnsafeMutablePointer< CMBufferQueueTriggerToken?>?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBufferQueueIsAtEndOfData(queue _, :  CMBufferQueue) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMBufferQueueIsEmpty(queue _, :  CMBufferQueue) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBufferQueueMarkEndOfData(queue _, :  CMBufferQueue) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBufferQueueRemoveTrigger(queue _, :  CMBufferQueue,  triggerToken:  CMBufferQueueTriggerToken) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMBufferQueueReset(queue _, :  CMBufferQueue) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBufferQueueResetWithCallback(queue _, :  CMBufferQueue,  callback: ( CMBuffer,  UnsafeMutableRawPointer?) ->  Void,  refcon:  UnsafeMutableRawPointer?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBufferQueueSetValidationCallback(queue _, :  CMBufferQueue,  callback:  CMBufferValidationCallback,  refcon:  UnsafeMutableRawPointer?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMBufferQueueSetValidationHandler(queue _, handler :  CMBufferQueue,  _, :  @escaping  CMBufferValidationHandler) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.2+
//   - iPadOS 12.2+
//   - macOS 10.14.4+
//   - tvOS 12.2+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMBufferQueueTestTrigger(queue _, :  CMBufferQueue,  triggerToken:  CMBufferQueueTriggerToken) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMClockConvertHostTimeToSystemUnits(hostTime _, :  CMTime) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMClockGetAnchorTime(clock _, :  CMClock,  clockTimeOut:  UnsafeMutablePointer< CMTime>,  referenceClockTimeOut:  UnsafeMutablePointer< CMTime>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMClockGetHostTimeClock() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMClockGetTime(clock _, :  CMClock) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMClockGetTypeID() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMClockInvalidate(clock _, :  CMClock) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMClockMakeHostTimeFromSystemUnits(hostTime _, :  UInt64) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMClockMightDrift(clock _, :  CMClock,  otherClock:  CMClock) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer(allocator:  CFAllocator?,  closedCaptionFormatDescription:  CMClosedCaptionFormatDescription,  flavor:  CMClosedCaptionDescriptionFlavor?,  blockBufferOut:  UnsafeMutablePointer< CMBlockBuffer?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer(closedCaptionDescriptionBlockBuffer allocator:  CFAllocator?,  bigEndianClosedCaptionDescriptionBlockBuffer, :  CMBlockBuffer,  flavor:  CMClosedCaptionDescriptionFlavor?,  formatDescriptionOut:  UnsafeMutablePointer< CMClosedCaptionFormatDescription?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData(closedCaptionDescriptionData allocator:  CFAllocator?,  bigEndianClosedCaptionDescriptionData, :  UnsafePointer< UInt8>,  size:  Int,  flavor:  CMClosedCaptionDescriptionFlavor?,  formatDescriptionOut:  UnsafeMutablePointer< CMClosedCaptionFormatDescription?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMCopyDictionaryOfAttachments(allocator:  CFAllocator?,  target:  CMAttachmentBearer,  attachmentMode:  CMAttachmentMode) -> sending  CFDictionary?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout(soundDescriptionBlockBuffer _, :  CMBlockBuffer,  flavor:  CMSoundDescriptionFlavor?) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMFormatDescriptionCreate(allocator:  CFAllocator?,  mediaType:  CMMediaType,  mediaSubType:  FourCharCode,  extensions:  CFDictionary?,  formatDescriptionOut:  UnsafeMutablePointer< CMFormatDescription?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMFormatDescriptionEqual(formatDescription _, :  CMFormatDescription?,  otherFormatDescription:  CMFormatDescription?) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMFormatDescriptionEqualIgnoringExtensionKeys(formatDescription _, formatDescriptionExtensionKeysToIgnore :  CMFormatDescription?,  otherFormatDescription:  CMFormatDescription?,  extensionKeysToIgnore, :  CFTypeRef?,  sampleDescriptionExtensionAtomKeysToIgnore:  CFTypeRef?) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.3+
//   - iPadOS 4.3+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMFormatDescriptionGetExtension(desc _, :  CMFormatDescription,  extensionKey:  CFString) ->  CFPropertyList?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMFormatDescriptionGetExtensions(desc _, :  CMFormatDescription) ->  CFDictionary?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMFormatDescriptionGetMediaSubType(desc _, :  CMFormatDescription) ->  FourCharCode) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMFormatDescriptionGetMediaType(desc _, :  CMFormatDescription) ->  CMMediaType) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMFormatDescriptionGetTypeID() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMGetAttachment(target _, :  CMAttachmentBearer,  key:  CFString,  attachmentModeOut:  UnsafeMutablePointer< CMAttachmentMode>?) ->  CFTypeRef?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMMemoryPoolCreate(options:  CFDictionary?) ->  CMMemoryPool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMMemoryPoolFlush(pool _, :  CMMemoryPool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMMemoryPoolGetAllocator(pool _, :  CMMemoryPool) ->  CFAllocator) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMMemoryPoolGetTypeID() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMMemoryPoolInvalidate(pool _, :  CMMemoryPool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMMetadataCreateIdentifierForKeyAndKeySpace(allocator:  CFAllocator?,  key:  CFTypeRef,  keySpace:  CFString,  identifierOut:  UnsafeMutablePointer< CFString?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMMetadataCreateKeyFromIdentifier(allocator:  CFAllocator?,  identifier:  CFString,  keyOut:  UnsafeMutablePointer< CFTypeRef?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMMetadataCreateKeyFromIdentifierAsCFData(allocator:  CFAllocator?,  identifier:  CFString,  keyOut:  UnsafeMutablePointer< CFData?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMMetadataCreateKeySpaceFromIdentifier(allocator:  CFAllocator?,  identifier:  CFString,  keySpaceOut:  UnsafeMutablePointer< CFString?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMMetadataDataTypeRegistryDataTypeConformsToDataType(dataType _, conformsToDataType :  CFString,  conformsTo, :  CFString) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMMetadataDataTypeRegistryDataTypeIsBaseDataType(dataType _, :  CFString) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMMetadataDataTypeRegistryDataTypeIsRegistered(dataType _, :  CFString) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType(dataType _, :  CFString) ->  CFString) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMMetadataDataTypeRegistryGetBaseDataTypes() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMMetadataDataTypeRegistryGetConformingDataTypes(dataType _, :  CFString) ->  CFArray) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMMetadataDataTypeRegistryGetDataTypeDescription(dataType _, :  CFString) ->  CFString) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMMetadataDataTypeRegistryRegisterDataType(dataType _, :  CFString,  description:  CFString,  conformingDataTypes:  CFArray) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer(allocator:  CFAllocator?,  metadataFormatDescription:  CMMetadataFormatDescription,  flavor:  CMMetadataDescriptionFlavor?,  blockBufferOut:  UnsafeMutablePointer< CMBlockBuffer?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions(allocator:  CFAllocator?,  sourceDescription:  CMMetadataFormatDescription,  otherSourceDescription:  CMMetadataFormatDescription,  formatDescriptionOut:  UnsafeMutablePointer< CMMetadataFormatDescription?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer(metadataDescriptionBlockBuffer allocator:  CFAllocator?,  bigEndianMetadataDescriptionBlockBuffer, :  CMBlockBuffer,  flavor:  CMMetadataDescriptionFlavor?,  formatDescriptionOut:  UnsafeMutablePointer< CMMetadataFormatDescription?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData(metadataDescriptionData allocator:  CFAllocator?,  bigEndianMetadataDescriptionData, :  UnsafePointer< UInt8>,  size:  Int,  flavor:  CMMetadataDescriptionFlavor?,  formatDescriptionOut:  UnsafeMutablePointer< CMMetadataFormatDescription?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMMetadataFormatDescriptionCreateWithKeys(allocator:  CFAllocator?,  metadataType:  CMMetadataFormatType,  keys:  CFArray?,  formatDescriptionOut:  UnsafeMutablePointer< CMMetadataFormatDescription?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications(allocator:  CFAllocator?,  sourceDescription:  CMMetadataFormatDescription,  metadataSpecifications:  CFArray,  formatDescriptionOut:  UnsafeMutablePointer< CMMetadataFormatDescription?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMMetadataFormatDescriptionCreateWithMetadataSpecifications(allocator:  CFAllocator?,  metadataType:  CMMetadataFormatType,  metadataSpecifications:  CFArray,  formatDescriptionOut:  UnsafeMutablePointer< CMMetadataFormatDescription?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMMetadataFormatDescriptionGetIdentifiers(desc _, :  CMMetadataFormatDescription) ->  CFArray?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMMetadataFormatDescriptionGetKeyWithLocalID(desc _, :  CMMetadataFormatDescription,  localKeyID:  OSType) ->  CFDictionary?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMMuxedFormatDescriptionCreate(allocator:  CFAllocator?,  muxType:  CMMuxedStreamType,  extensions:  CFDictionary?,  formatDescriptionOut:  UnsafeMutablePointer< CMMuxedFormatDescription?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMPropagateAttachments(source _, :  CMAttachmentBearer,  destination:  CMAttachmentBearer) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMRemoveAllAttachments(target _, :  CMAttachmentBearer) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMRemoveAttachment(target _, :  CMAttachmentBearer,  key:  CFString) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSampleBufferCallBlockForEachSample(sbuf _, handler :  CMSampleBuffer,  _, : ( CMSampleBuffer,  CMItemCount) ->  OSStatus) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferCallForEachSample(sbuf _, :  CMSampleBuffer,  callback: ( CMSampleBuffer,  CMItemCount,  UnsafeMutableRawPointer?) ->  OSStatus,  refcon:  UnsafeMutableRawPointer?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSampleBufferCopyPCMDataIntoAudioBufferList(sbuf _, frameOffset :  CMSampleBuffer,  at, numFrames :  Int32,  frameCount, bufferList :  Int32,  into, :  UnsafeMutablePointer< AudioBufferList>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferCopySampleBufferForRange(sbuf allocator:  CFAllocator?,  sampleBuffer, :  CMSampleBuffer,  sampleRange:  CFRange,  sampleBufferOut:  UnsafeMutablePointer< CMSampleBuffer?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferCreate(makeDataReadyRefcon allocator:  CFAllocator?,  dataBuffer:  CMBlockBuffer?,  dataReady:  Bool,  makeDataReadyCallback:  CMSampleBufferMakeDataReadyCallback?,  refcon, numSamples :  UnsafeMutableRawPointer?,  formatDescription:  CMFormatDescription?,  sampleCount, numSampleTimingEntries :  CMItemCount,  sampleTimingEntryCount, numSampleSizeEntries :  CMItemCount,  sampleTimingArray:  UnsafePointer< CMSampleTimingInfo>?,  sampleSizeEntryCount, :  CMItemCount,  sampleSizeArray:  UnsafePointer< Int>?,  sampleBufferOut:  UnsafeMutablePointer< CMSampleBuffer?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSampleBufferCreateCopy(sbuf allocator:  CFAllocator?,  sampleBuffer, :  CMSampleBuffer,  sampleBufferOut:  UnsafeMutablePointer< CMSampleBuffer?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferCreateCopyWithNewTiming(originalSBuf allocator:  CFAllocator?,  sampleBuffer, numSampleTimingEntries :  CMSampleBuffer,  sampleTimingEntryCount, :  CMItemCount,  sampleTimingArray:  UnsafePointer< CMSampleTimingInfo>?,  sampleBufferOut:  UnsafeMutablePointer< CMSampleBuffer?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferCreateForImageBuffer(makeDataReadyRefcon allocator:  CFAllocator?,  imageBuffer:  CVImageBuffer,  dataReady:  Bool,  makeDataReadyCallback:  CMSampleBufferMakeDataReadyCallback?,  refcon, :  UnsafeMutableRawPointer?,  formatDescription:  CMVideoFormatDescription,  sampleTiming:  UnsafePointer< CMSampleTimingInfo>,  sampleBufferOut:  UnsafeMutablePointer< CMSampleBuffer?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSampleBufferCreateForImageBufferWithMakeDataReadyHandler(allocator _, imageBuffer :  CFAllocator?,  _, dataReady :  CVImageBuffer,  _, formatDescription :  Bool,  _, sampleTiming :  CMVideoFormatDescription,  _, sampleBufferOut :  UnsafePointer< CMSampleTimingInfo>,  _, makeDataReadyHandler :  UnsafeMutablePointer< CMSampleBuffer?>,  _, :  CMSampleBufferMakeDataReadyHandler?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.2+
//   - iPadOS 12.2+
//   - macOS 10.14.4+
//   - tvOS 12.2+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferCreateForTaggedBufferGroup(allocator CFAllocatorRef, taggedBufferGroup ,  CMTaggedBufferGroupRef, sbufPTS ,  CMTime, sbufDuration ,  CMTime, formatDescription ,  CMTaggedBufferGroupFormatDescriptionRef, sBufOut ,  CMSampleBufferRef *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMSampleBufferCreateReady(numSamples allocator:  CFAllocator?,  dataBuffer:  CMBlockBuffer?,  formatDescription:  CMFormatDescription?,  sampleCount, numSampleTimingEntries :  CMItemCount,  sampleTimingEntryCount, numSampleSizeEntries :  CMItemCount,  sampleTimingArray:  UnsafePointer< CMSampleTimingInfo>?,  sampleSizeEntryCount, :  CMItemCount,  sampleSizeArray:  UnsafePointer< Int>?,  sampleBufferOut:  UnsafeMutablePointer< CMSampleBuffer?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSampleBufferCreateReadyWithImageBuffer(allocator:  CFAllocator?,  imageBuffer:  CVImageBuffer,  formatDescription:  CMVideoFormatDescription,  sampleTiming:  UnsafePointer< CMSampleTimingInfo>,  sampleBufferOut:  UnsafeMutablePointer< CMSampleBuffer?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferCreateWithMakeDataReadyHandler(allocator _, dataBuffer :  CFAllocator?,  _, dataReady :  CMBlockBuffer?,  _, formatDescription :  Bool,  _, numSamples :  CMFormatDescription?,  _, numSampleTimingEntries :  CMItemCount,  _, sampleTimingArray :  CMItemCount,  _, numSampleSizeEntries :  UnsafePointer< CMSampleTimingInfo>?,  _, sampleSizeArray :  CMItemCount,  _, sampleBufferOut :  UnsafePointer< Int>?,  _, makeDataReadyHandler :  UnsafeMutablePointer< CMSampleBuffer?>,  _, :  CMSampleBufferMakeDataReadyHandler?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.2+
//   - iPadOS 12.2+
//   - macOS 10.14.4+
//   - tvOS 12.2+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferDataIsReady(sbuf _, :  CMSampleBuffer) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSampleBufferGetAudioBufferListWithRetainedBlockBuffer(sbuf _, blockBufferStructureAllocator :  CMSampleBuffer,  bufferListSizeNeededOut:  UnsafeMutablePointer< Int>?,  bufferListOut:  UnsafeMutablePointer< AudioBufferList>?,  bufferListSize:  Int,  blockBufferAllocator, blockBufferBlockAllocator :  CFAllocator?,  blockBufferMemoryAllocator, :  CFAllocator?,  flags:  UInt32,  blockBufferOut:  UnsafeMutablePointer< CMBlockBuffer?>?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferGetAudioStreamPacketDescriptions(sbuf _, packetDescriptionsSize :  CMSampleBuffer,  allocatedSize, :  Int,  packetDescriptionsOut:  UnsafeMutablePointer< AudioStreamPacketDescription>?,  packetDescriptionsSizeNeededOut:  UnsafeMutablePointer< Int>?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferGetAudioStreamPacketDescriptionsPtr(sbuf _, packetDescriptionsSizeOut :  CMSampleBuffer,  packetDescriptionsPointerOut:  UnsafeMutablePointer< UnsafePointer< AudioStreamPacketDescription>?>?,  sizeOut, :  UnsafeMutablePointer< Int>?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSampleBufferGetDataBuffer(sbuf _, :  CMSampleBuffer) ->  CMBlockBuffer?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferGetDecodeTimeStamp(sbuf _, :  CMSampleBuffer) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferGetDuration(sbuf _, :  CMSampleBuffer) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSampleBufferGetFormatDescription(sbuf _, :  CMSampleBuffer) ->  CMFormatDescription?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferGetImageBuffer(sbuf _, :  CMSampleBuffer) ->  CVImageBuffer?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferGetNumSamples(sbuf _, :  CMSampleBuffer) ->  CMItemCount) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSampleBufferGetOutputDecodeTimeStamp(sbuf _, :  CMSampleBuffer) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferGetOutputDuration(sbuf _, :  CMSampleBuffer) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferGetOutputPresentationTimeStamp(sbuf _, :  CMSampleBuffer) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSampleBufferGetOutputSampleTimingInfoArray(sbuf _, timingArrayEntries :  CMSampleBuffer,  entryCount, timingArrayOut :  CMItemCount,  arrayToFill, timingArrayEntriesNeededOut :  UnsafeMutablePointer< CMSampleTimingInfo>?,  entriesNeededOut, :  UnsafeMutablePointer< CMItemCount>?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferGetPresentationTimeStamp(sbuf _, :  CMSampleBuffer) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferGetSampleAttachmentsArray(sbuf _, :  CMSampleBuffer,  createIfNecessary:  Bool) ->  CFArray?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSampleBufferGetSampleSize(sbuf _, sampleIndex :  CMSampleBuffer,  at, :  CMItemIndex) ->  Int) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferGetSampleSizeArray(sbuf _, sizeArrayEntries :  CMSampleBuffer,  entryCount, sizeArrayOut :  CMItemCount,  arrayToFill, sizeArrayEntriesNeededOut :  UnsafeMutablePointer< Int>?,  entriesNeededOut, :  UnsafeMutablePointer< CMItemCount>?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferGetSampleTimingInfo(sbuf _, sampleIndex :  CMSampleBuffer,  at, :  CMItemIndex,  timingInfoOut:  UnsafeMutablePointer< CMSampleTimingInfo>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSampleBufferGetSampleTimingInfoArray(sbuf _, numSampleTimingEntries :  CMSampleBuffer,  entryCount, timingArrayOut :  CMItemCount,  arrayToFill, timingArrayEntriesNeededOut :  UnsafeMutablePointer< CMSampleTimingInfo>?,  entriesNeededOut, :  UnsafeMutablePointer< CMItemCount>?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferGetTaggedBufferGroup(sbuf CMSampleBufferRef, );) extern   CMTaggedBufferGroupRef
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMSampleBufferGetTotalSampleSize(sbuf _, :  CMSampleBuffer) ->  Int) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSampleBufferGetTypeID() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferHasDataFailed(sbuf _, :  CMSampleBuffer,  statusOut:  UnsafeMutablePointer< OSStatus>?) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferInvalidate(sbuf _, :  CMSampleBuffer) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSampleBufferIsValid(sbuf _, :  CMSampleBuffer) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferMakeDataReady(sbuf _, :  CMSampleBuffer) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferSetDataBuffer(sbuf _, dataBuffer :  CMSampleBuffer,  newValue, :  CMBlockBuffer) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSampleBufferSetDataBufferFromAudioBufferList(sbuf _, blockBufferStructureAllocator :  CMSampleBuffer,  blockBufferAllocator, blockBufferBlockAllocator :  CFAllocator?,  blockBufferMemoryAllocator, :  CFAllocator?,  flags:  UInt32,  bufferList:  UnsafePointer< AudioBufferList>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferSetDataFailed(sbuf _, :  CMSampleBuffer,  status:  OSStatus) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferSetDataReady(sbuf _, :  CMSampleBuffer) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSampleBufferSetInvalidateCallback(sbuf _, invalidateCallback :  CMSampleBuffer,  callback, invalidateRefCon :  CMSampleBufferInvalidateCallback,  refcon, :  UInt64) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferSetInvalidateHandler(sbuf _, :  CMSampleBuffer,  invalidateHandler:  @escaping  CMSampleBufferInvalidateHandler) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSampleBufferSetOutputPresentationTimeStamp(sbuf _, outputPresentationTimeStamp :  CMSampleBuffer,  newValue, :  CMTime) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSampleBufferTrackDataReadiness(sbuf _, :  CMSampleBuffer,  sampleBufferToTrack:  CMSampleBuffer) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSetAttachment(target _, :  CMAttachmentBearer,  key:  CFString,  value:  CFTypeRef?,  attachmentMode:  CMAttachmentMode) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSetAttachments(target _, theAttachments :  CMAttachmentBearer,  attachments, :  CFDictionary,  attachmentMode:  CMAttachmentMode) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSimpleQueueCreate(allocator:  CFAllocator?,  capacity:  Int32,  queueOut:  UnsafeMutablePointer< CMSimpleQueue?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSimpleQueueDequeue(queue _, :  CMSimpleQueue) ->  UnsafeRawPointer?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSimpleQueueEnqueue(queue _, :  CMSimpleQueue,  element:  UnsafeRawPointer) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSimpleQueueGetCapacity(queue _, :  CMSimpleQueue) ->  Int32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSimpleQueueGetCount(queue _, :  CMSimpleQueue) ->  Int32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSimpleQueueGetHead(queue _, :  CMSimpleQueue) ->  UnsafeRawPointer?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSimpleQueueGetTypeID() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSimpleQueueReset(queue _, :  CMSimpleQueue) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSwapBigEndianClosedCaptionDescriptionToHost(closedCaptionDescriptionData _, closedCaptionDescriptionSize :  UnsafeMutablePointer< UInt8>,  _, :  Int) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSwapBigEndianImageDescriptionToHost(imageDescriptionData _, imageDescriptionSize :  UnsafeMutablePointer< UInt8>,  _, :  Int) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSwapBigEndianMetadataDescriptionToHost(metadataDescriptionData _, metadataDescriptionSize :  UnsafeMutablePointer< UInt8>,  _, :  Int) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSwapBigEndianSoundDescriptionToHost(soundDescriptionData _, soundDescriptionSize :  UnsafeMutablePointer< UInt8>,  _, :  Int) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSwapBigEndianTextDescriptionToHost(textDescriptionData _, textDescriptionSize :  UnsafeMutablePointer< UInt8>,  _, :  Int) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSwapBigEndianTimeCodeDescriptionToHost(timeCodeDescriptionData _, timeCodeDescriptionSize :  UnsafeMutablePointer< UInt8>,  _, :  Int) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSwapHostEndianClosedCaptionDescriptionToBig(closedCaptionDescriptionData _, closedCaptionDescriptionSize :  UnsafeMutablePointer< UInt8>,  _, :  Int) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSwapHostEndianImageDescriptionToBig(imageDescriptionData _, imageDescriptionSize :  UnsafeMutablePointer< UInt8>,  _, :  Int) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSwapHostEndianMetadataDescriptionToBig(metadataDescriptionData _, metadataDescriptionSize :  UnsafeMutablePointer< UInt8>,  _, :  Int) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSwapHostEndianSoundDescriptionToBig(soundDescriptionData _, soundDescriptionSize :  UnsafeMutablePointer< UInt8>,  _, :  Int) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSwapHostEndianTextDescriptionToBig(textDescriptionData _, textDescriptionSize :  UnsafeMutablePointer< UInt8>,  _, :  Int) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSwapHostEndianTimeCodeDescriptionToBig(timeCodeDescriptionData _, timeCodeDescriptionSize :  UnsafeMutablePointer< UInt8>,  _, :  Int) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSyncConvertTime(time _, fromClockOrTimebase :  CMTime,  from, toClockOrTimebase :  CMClockOrTimebase,  to, :  CMClockOrTimebase) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSyncGetRelativeRate(ofClockOrTimebase _, relativeToClockOrTimebase :  CMClockOrTimebase,  relativeTo, :  CMClockOrTimebase) ->  Float64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSyncGetRelativeRateAndAnchorTime(ofClockOrTimebase _, relativeToClockOrTimebase :  CMClockOrTimebase,  relativeTo, outRelativeRate :  CMClockOrTimebase,  relativeRateOut, outOfClockOrTimebaseAnchorTime :  UnsafeMutablePointer< Float64>?,  anchorTimeOut, outRelativeToClockOrTimebaseAnchorTime :  UnsafeMutablePointer< CMTime>?,  relativeToAnchorTimeOut, :  UnsafeMutablePointer< CMTime>?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMSyncGetTime(clockOrTimebase _, :  CMClockOrTimebase) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMSyncMightDrift(clockOrTimebase1 _, clockOrTimebase2 :  CMClockOrTimebase,  _, :  CMClockOrTimebase) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTagCategoryEqualToTagCategory(tag1 CMTag, tag2 ,  CMTag, );) static   Boolean
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagCategoryValueEqualToValue(tag1 CMTag, tag2 ,  CMTag, );) static   Boolean
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMTagCollectionAddTag(tagCollection CMMutableTagCollectionRef, tagToAdd ,  CMTag, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagCollectionAddTagsFromArray(tagCollection CMMutableTagCollectionRef, tags ,  CMTag *, tagCount ,  CMItemCount, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagCollectionAddTagsFromCollection(tagCollection CMMutableTagCollectionRef, collectionWithTagsToAdd ,  CMTagCollectionRef, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMTagCollectionApply(tagCollection CMTagCollectionRef, applier ,  CMTagCollectionApplierFunction, context ,  void *, );) extern   void
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagCollectionApplyUntil(tagCollection CMTagCollectionRef, applier ,  CMTagCollectionTagFilterFunction, context ,  void *, );) extern   CMTag
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagCollectionContainsCategory(tagCollection CMTagCollectionRef, category ,  CMTagCategory, );) extern   Boolean
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMTagCollectionContainsSpecifiedTags(tagCollection CMTagCollectionRef, containedTags ,  const  CMTag *, containedTagCount ,  CMItemCount, );) extern   Boolean
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagCollectionContainsTag(tagCollection CMTagCollectionRef, tag ,  CMTag, );) extern   Boolean
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagCollectionContainsTagsOfCollection(tagCollection CMTagCollectionRef, containedTagCollection ,  CMTagCollectionRef, );) extern   Boolean
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMTagCollectionCopyAsData(tagCollection CMTagCollectionRef, allocator ,  CFAllocatorRef, );) extern   CFDataRef
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagCollectionCopyAsDictionary(tagCollection CMTagCollectionRef, allocator ,  CFAllocatorRef, );) extern   CFDictionaryRef
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagCollectionCopyDescription(allocator CFAllocatorRef, tagCollection ,  CMTagCollectionRef, );) extern   CFStringRef
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMTagCollectionCopyTagsOfCategories(allocator CFAllocatorRef, tagCollection ,  CMTagCollectionRef, categories ,  const  CMTagCategory *, categoriesCount ,  CMItemCount, collectionWithTagsOfCategories ,  CMTagCollectionRef *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagCollectionCountTagsWithFilterFunction(tagCollection CMTagCollectionRef, filterApplier ,  CMTagCollectionTagFilterFunction, context ,  void *, );) extern   CMItemCount
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagCollectionCreate(allocator CFAllocatorRef, tags ,  const  CMTag *, tagCount ,  CMItemCount, newCollectionOut ,  CMTagCollectionRef *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMTagCollectionCreateCopy(tagCollection CMTagCollectionRef, allocator ,  CFAllocatorRef, newCollectionCopyOut ,  CMTagCollectionRef *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagCollectionCreateDifference(tagCollectionMinuend CMTagCollectionRef, tagCollectionSubtrahend ,  CMTagCollectionRef, tagCollectionOut ,  CMTagCollectionRef *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagCollectionCreateExclusiveOr(tagCollection1 CMTagCollectionRef, tagCollection2 ,  CMTagCollectionRef, tagCollectionOut ,  CMTagCollectionRef *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMTagCollectionCreateFromData(data CFDataRef, allocator ,  CFAllocatorRef, newCollectionOut ,  CMTagCollectionRef *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagCollectionCreateFromDictionary(dict CFDictionaryRef, allocator ,  CFAllocatorRef, newCollectionOut ,  CMTagCollectionRef *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagCollectionCreateIntersection(tagCollection1 CMTagCollectionRef, tagCollection2 ,  CMTagCollectionRef, tagCollectionOut ,  CMTagCollectionRef *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMTagCollectionCreateMutable(allocator CFAllocatorRef, capacity ,  CFIndex, newMutableCollectionOut ,  CMMutableTagCollectionRef *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagCollectionCreateMutableCopy(tagCollection CMTagCollectionRef, allocator ,  CFAllocatorRef, newMutableCollectionCopyOut ,  CMMutableTagCollectionRef *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagCollectionCreateUnion(tagCollection1 CMTagCollectionRef, tagCollection2 ,  CMTagCollectionRef, tagCollectionOut ,  CMTagCollectionRef *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMTagCollectionGetCount(tagCollection CMTagCollectionRef, );) extern   CMItemCount
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagCollectionGetCountOfCategory(tagCollection CMTagCollectionRef, category ,  CMTagCategory, );) extern   CMItemCount
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagCollectionGetTags(tagCollection CMTagCollectionRef, tagBuffer ,  CMTag *, tagBufferCount ,  CMItemCount, numberOfTagsCopied ,  CMItemCount *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMTagCollectionGetTagsWithCategory(tagCollection CMTagCollectionRef, category ,  CMTagCategory, tagBuffer ,  CMTag *, tagBufferCount ,  CMItemCount, numberOfTagsCopied ,  CMItemCount *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagCollectionGetTagsWithFilterFunction(tagCollection CMTagCollectionRef, tagBuffer ,  CMTag *, tagBufferCount ,  CMItemCount, numberOfTagsCopied ,  CMItemCount *, filter ,  CMTagCollectionTagFilterFunction, context ,  void *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagCollectionGetTypeID() extern   CFTypeID
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMTagCollectionIsEmpty(tagCollection CMTagCollectionRef, );) extern   Boolean
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagCollectionRemoveAllTags(tagCollection CMMutableTagCollectionRef, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagCollectionRemoveAllTagsOfCategory(tagCollection CMMutableTagCollectionRef, category ,  CMTagCategory, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMTagCollectionRemoveTag(tagCollection CMMutableTagCollectionRef, tagToRemove ,  CMTag, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagCompare(tag1 CMTag, tag2 ,  CMTag, );) extern   CFComparisonResult
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagCopyAsDictionary(tag CMTag, allocator ,  CFAllocatorRef, );) extern   CFDictionaryRef
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMTagCopyDescription(allocator CFAllocatorRef, tag ,  CMTag, );) extern   CFStringRef
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagEqualToTag(tag1 CMTag, tag2 ,  CMTag, );) extern   Boolean
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagGetCategory(tag CMTag, );) static   CMTagCategory
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMTagGetFlagsValue(tag CMTag, );) extern   uint64_t
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagGetFloat64Value(tag CMTag, );) extern   Float64
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagGetOSTypeValue(tag CMTag, );) extern   OSType
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMTagGetSInt64Value(tag CMTag, );) extern   int64_t
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagGetValue(tag CMTag, );) static   CMTagValue
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagGetValueDataType(tag CMTag, );) extern   CMTagDataType
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMTagHasCategory(tag CMTag, category ,  CMTagCategory, );) static   Boolean
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagHasFlagsValue(tag CMTag, );) extern   Boolean
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagHasFloat64Value(tag CMTag, );) extern   Boolean
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMTagHasOSTypeValue(tag CMTag, );) extern   Boolean
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagHasSInt64Value(tag CMTag, );) extern   Boolean
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagHash(tag CMTag, );) extern   CFHashCode
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMTagIsValid(tag CMTag, );) static   Boolean
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagMakeFromDictionary(dict CFDictionaryRef, );) extern   CMTag
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagMakeWithFlagsValue(category CMTagCategory, flagsForTag ,  uint64_t, );) extern   CMTag
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMTagMakeWithFloat64Value(category CMTagCategory, value ,  Float64, );) extern   CMTag
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagMakeWithOSTypeValue(category CMTagCategory, value ,  OSType, );) extern   CMTag
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTagMakeWithSInt64Value(category CMTagCategory, value ,  int64_t, );) extern   CMTag
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMTaggedBufferGroupCreate(allocator CFAllocatorRef, tagCollections ,  CFArrayRef, buffers ,  CFArrayRef, groupOut ,  CMTaggedBufferGroupRef *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTaggedBufferGroupCreateCombined(allocator CFAllocatorRef, taggedBufferGroups ,  CFArrayRef, groupOut ,  CMTaggedBufferGroupRef *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup(allocator CFAllocatorRef, taggedBufferGroup ,  CMTaggedBufferGroupRef, formatDescriptionOut ,  CMTaggedBufferGroupFormatDescriptionRef *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions(allocator CFAllocatorRef, taggedBufferGroup ,  CMTaggedBufferGroupRef, extensions ,  CFDictionaryRef, formatDescriptionOut ,  CMTaggedBufferGroupFormatDescriptionRef *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+

// CMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup(desc CMTaggedBufferGroupFormatDescriptionRef, taggedBufferGroup ,  CMTaggedBufferGroupRef, );) extern   Boolean
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTaggedBufferGroupGetCMSampleBufferAtIndex(group CMTaggedBufferGroupRef, index ,  CFIndex, );) extern   CMSampleBufferRef
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMTaggedBufferGroupGetCMSampleBufferForTag(group CMTaggedBufferGroupRef, tag ,  CMTag, indexOut ,  CFIndex *, );) extern   CMSampleBufferRef
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTaggedBufferGroupGetCMSampleBufferForTagCollection(group CMTaggedBufferGroupRef, tagCollection ,  CMTagCollectionRef, indexOut ,  CFIndex *, );) extern   CMSampleBufferRef
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTaggedBufferGroupGetCVPixelBufferAtIndex(group CMTaggedBufferGroupRef, index ,  CFIndex, );) extern   CVPixelBufferRef
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMTaggedBufferGroupGetCVPixelBufferForTag(group CMTaggedBufferGroupRef, tag ,  CMTag, indexOut ,  CFIndex *, );) extern   CVPixelBufferRef
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTaggedBufferGroupGetCVPixelBufferForTagCollection(group CMTaggedBufferGroupRef, tagCollection ,  CMTagCollectionRef, indexOut ,  CFIndex *, );) extern   CVPixelBufferRef
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTaggedBufferGroupGetCount(group CMTaggedBufferGroupRef, );) extern   CMItemCount
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMTaggedBufferGroupGetNumberOfMatchesForTagCollection(group CMTaggedBufferGroupRef, tagCollection ,  CMTagCollectionRef, );) extern   CMItemCount
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTaggedBufferGroupGetTagCollectionAtIndex(group CMTaggedBufferGroupRef, index ,  CFIndex, );) extern   CMTagCollectionRef
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// CMTaggedBufferGroupGetTypeID() extern   CFTypeID
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer(allocator:  CFAllocator?,  textFormatDescription:  CMTextFormatDescription,  flavor:  CMTextDescriptionFlavor?,  blockBufferOut:  UnsafeMutablePointer< CMBlockBuffer?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer(textDescriptionBlockBuffer allocator:  CFAllocator?,  bigEndianTextDescriptionBlockBuffer, :  CMBlockBuffer,  flavor:  CMTextDescriptionFlavor?,  mediaType:  CMMediaType,  formatDescriptionOut:  UnsafeMutablePointer< CMTextFormatDescription?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTextFormatDescriptionCreateFromBigEndianTextDescriptionData(textDescriptionData allocator:  CFAllocator?,  bigEndianTextDescriptionData, :  UnsafePointer< UInt8>,  size:  Int,  flavor:  CMTextDescriptionFlavor?,  mediaType:  CMMediaType,  formatDescriptionOut:  UnsafeMutablePointer< CMTextFormatDescription?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTextFormatDescriptionGetDefaultStyle(desc _, :  CMFormatDescription,  localFontIDOut:  UnsafeMutablePointer< UInt16>?,  boldOut:  UnsafeMutablePointer< DarwinBoolean>?,  italicOut:  UnsafeMutablePointer< DarwinBoolean>?,  underlineOut:  UnsafeMutablePointer< DarwinBoolean>?,  fontSizeOut:  UnsafeMutablePointer< CGFloat>?,  colorComponentsOut:  UnsafeMutablePointer< CGFloat>?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTextFormatDescriptionGetDefaultTextBox(desc _, :  CMFormatDescription,  originIsAtTopLeft:  Bool,  heightOfTextTrack:  CGFloat,  defaultTextBoxOut:  UnsafeMutablePointer< CGRect>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTextFormatDescriptionGetDisplayFlags(desc _, :  CMFormatDescription,  displayFlagsOut:  UnsafeMutablePointer< CMTextDisplayFlags>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTextFormatDescriptionGetFontName(desc _, :  CMFormatDescription,  localFontID:  UInt16,  fontNameOut:  AutoreleasingUnsafeMutablePointer< CFString?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTextFormatDescriptionGetJustification(desc _, horizontaJustificationlOut :  CMFormatDescription,  horizontalOut, verticalJustificationOut :  UnsafeMutablePointer< CMTextJustificationValue>?,  verticalOut, :  UnsafeMutablePointer< CMTextJustificationValue>?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeAbsoluteValue(time _, :  CMTime) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTimeAdd(lhs _, rhs :  CMTime,  _, :  CMTime) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeClampToRange(time _, :  CMTime,  range:  CMTimeRange) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer(allocator:  CFAllocator?,  timeCodeFormatDescription:  CMTimeCodeFormatDescription,  flavor:  CMTimeCodeDescriptionFlavor?,  blockBufferOut:  UnsafeMutablePointer< CMBlockBuffer?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTimeCodeFormatDescriptionCreate(allocator:  CFAllocator?,  timeCodeFormatType:  CMTimeCodeFormatType,  frameDuration:  CMTime,  frameQuanta:  UInt32,  flags:  UInt32,  extensions:  CFDictionary?,  formatDescriptionOut:  UnsafeMutablePointer< CMTimeCodeFormatDescription?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer(timeCodeDescriptionBlockBuffer allocator:  CFAllocator?,  bigEndianTimeCodeDescriptionBlockBuffer, :  CMBlockBuffer,  flavor:  CMTimeCodeDescriptionFlavor?,  formatDescriptionOut:  UnsafeMutablePointer< CMTimeCodeFormatDescription?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData(timeCodeDescriptionData allocator:  CFAllocator?,  bigEndianTimeCodeDescriptionData, :  UnsafePointer< UInt8>,  size:  Int,  flavor:  CMTimeCodeDescriptionFlavor?,  formatDescriptionOut:  UnsafeMutablePointer< CMTimeCodeFormatDescription?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTimeCodeFormatDescriptionGetFrameDuration(timeCodeFormatDescription _, :  CMTimeCodeFormatDescription) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeCodeFormatDescriptionGetFrameQuanta(timeCodeFormatDescription _, :  CMTimeCodeFormatDescription) ->  UInt32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeCodeFormatDescriptionGetTimeCodeFlags(desc _, :  CMTimeCodeFormatDescription) ->  UInt32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTimeCompare(time1 _, time2 :  CMTime,  _, :  CMTime) ->  Int32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeConvertScale(time _, newTimescale :  CMTime,  timescale, :  Int32,  method:  CMTimeRoundingMethod) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeCopyAsDictionary(time _, :  CMTime,  allocator:  CFAllocator?) ->  CFDictionary?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTimeCopyDescription(allocator:  CFAllocator?,  time:  CMTime) ->  CFString?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeFoldIntoRange(time _, :  CMTime,  foldRange:  CMTimeRange) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeGetSeconds(time _, :  CMTime) ->  Float64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTimeMake(value:  Int64,  timescale:  Int32) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeMakeFromDictionary(dictionaryRepresentation _, :  CFDictionary?) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeMakeWithEpoch(value:  Int64,  timescale:  Int32,  epoch:  Int64) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTimeMakeWithSeconds(seconds _, :  Float64,  preferredTimescale:  Int32) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeMapDurationFromRangeToRange(dur _, :  CMTime,  fromRange:  CMTimeRange,  toRange:  CMTimeRange) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeMapTimeFromRangeToRange(t _, :  CMTime,  fromRange:  CMTimeRange,  toRange:  CMTimeRange) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTimeMappingCopyAsDictionary(mapping _, :  CMTimeMapping,  allocator:  CFAllocator?) ->  CFDictionary?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeMappingCopyDescription(allocator:  CFAllocator?,  mapping:  CMTimeMapping) ->  CFString?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeMappingMake(source:  CMTimeRange,  target:  CMTimeRange) ->  CMTimeMapping) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTimeMappingMakeEmpty(target:  CMTimeRange) ->  CMTimeMapping) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeMappingMakeFromDictionary(dictionaryRepresentation _, :  CFDictionary) ->  CMTimeMapping) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeMappingShow(mapping _, :  CMTimeMapping) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTimeMaximum(time1 _, time2 :  CMTime,  _, :  CMTime) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeMinimum(time1 _, time2 :  CMTime,  _, :  CMTime) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeMultiply(time _, :  CMTime,  multiplier:  Int32) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTimeMultiplyByFloat64(time _, :  CMTime,  multiplier:  Float64) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeMultiplyByRatio(time _, :  CMTime,  multiplier:  Int32,  divisor:  Int32) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.1+
//   - iPadOS 7.1+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeRangeContainsTime(range _, :  CMTimeRange,  time:  CMTime) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTimeRangeContainsTimeRange(range _, :  CMTimeRange,  otherRange:  CMTimeRange) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeRangeCopyAsDictionary(range _, :  CMTimeRange,  allocator:  CFAllocator?) ->  CFDictionary?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeRangeCopyDescription(allocator:  CFAllocator?,  range:  CMTimeRange) ->  CFString?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTimeRangeEqual(range1 _, range2 :  CMTimeRange,  _, :  CMTimeRange) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeRangeFromTimeToTime(start:  CMTime,  end:  CMTime) ->  CMTimeRange) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeRangeGetEnd(range _, :  CMTimeRange) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTimeRangeGetIntersection(range _, :  CMTimeRange,  otherRange:  CMTimeRange) ->  CMTimeRange) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeRangeGetUnion(range _, :  CMTimeRange,  otherRange:  CMTimeRange) ->  CMTimeRange) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeRangeMake(start:  CMTime,  duration:  CMTime) ->  CMTimeRange) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTimeRangeMakeFromDictionary(dictionaryRepresentation _, :  CFDictionary) ->  CMTimeRange) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeRangeShow(range _, :  CMTimeRange) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimeShow(time _, :  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTimeSubtract(lhs _, rhs :  CMTime,  _, :  CMTime) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimebaseAddTimer(timebase _, :  CMTimebase,  timer:  CFRunLoopTimer,  runloop:  CFRunLoop) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimebaseAddTimerDispatchSource(timebase _, :  CMTimebase,  timerSource:  dispatch_source_t) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTimebaseCopyMaster(timebase _, :  CMTimebase) ->  CMClockOrTimebase) func
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 9.0+ (Deprecated in 9.0)
//   - iPadOS 9.0+ (Deprecated in 9.0)
//   - macOS 10.11+ (Deprecated in 10.11)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 6.0+ (Deprecated in 6.0)
//
// Deprecated: This function is deprecated.

// CMTimebaseCopyMasterClock(timebase _, :  CMTimebase) ->  CMClock?) func
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 9.0+ (Deprecated in 9.0)
//   - iPadOS 9.0+ (Deprecated in 9.0)
//   - macOS 10.11+ (Deprecated in 10.11)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 6.0+ (Deprecated in 6.0)
//
// Deprecated: This function is deprecated.

// CMTimebaseCopyMasterTimebase(timebase _, :  CMTimebase) ->  CMTimebase?) func
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 9.0+ (Deprecated in 9.0)
//   - iPadOS 9.0+ (Deprecated in 9.0)
//   - macOS 10.11+ (Deprecated in 10.11)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 6.0+ (Deprecated in 6.0)
//
// Deprecated: This function is deprecated.


// CMTimebaseCopySource(timebase _, :  CMTimebase) ->  CMClockOrTimebase) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimebaseCopySourceClock(timebase _, :  CMTimebase) ->  CMClock?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimebaseCopySourceTimebase(timebase _, :  CMTimebase) ->  CMTimebase?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTimebaseCopyUltimateMasterClock(timebase _, :  CMTimebase) ->  CMClock) func
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 9.0+ (Deprecated in 9.0)
//   - iPadOS 9.0+ (Deprecated in 9.0)
//   - macOS 10.11+ (Deprecated in 10.11)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 6.0+ (Deprecated in 6.0)
//
// Deprecated: This function is deprecated.

// CMTimebaseCopyUltimateSourceClock(timebase _, :  CMTimebase) ->  CMClock) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimebaseCreateWithMasterClock(allocator:  CFAllocator?,  masterClock:  CMClock,  timebaseOut:  UnsafeMutablePointer< CMTimebase?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 6.0+ (Deprecated in 8.0)
//   - iPadOS 6.0+ (Deprecated in 8.0)
//   - macOS 10.8+ (Deprecated in 10.10)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 6.0+ (Deprecated in 6.0)
//
// Deprecated: This function is deprecated.


// CMTimebaseCreateWithMasterTimebase(allocator:  CFAllocator?,  masterTimebase:  CMTimebase,  timebaseOut:  UnsafeMutablePointer< CMTimebase?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 6.0+ (Deprecated in 8.0)
//   - iPadOS 6.0+ (Deprecated in 8.0)
//   - macOS 10.8+ (Deprecated in 10.10)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 6.0+ (Deprecated in 6.0)
//
// Deprecated: This function is deprecated.

// CMTimebaseCreateWithSourceClock(allocator:  CFAllocator?,  sourceClock:  CMClock,  timebaseOut:  UnsafeMutablePointer< CMTimebase?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimebaseCreateWithSourceTimebase(allocator:  CFAllocator?,  sourceTimebase:  CMTimebase,  timebaseOut:  UnsafeMutablePointer< CMTimebase?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTimebaseGetEffectiveRate(timebase _, :  CMTimebase) ->  Float64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimebaseGetMaster(timebase _, :  CMTimebase) ->  CMClockOrTimebase?) func
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 6.0+ (Deprecated in 9.0)
//   - iPadOS 6.0+ (Deprecated in 9.0)
//   - macOS 10.8+ (Deprecated in 10.11)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// Deprecated: This function is deprecated.

// CMTimebaseGetMasterClock(timebase _, :  CMTimebase) ->  CMClock?) func
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 6.0+ (Deprecated in 9.0)
//   - iPadOS 6.0+ (Deprecated in 9.0)
//   - macOS 10.8+ (Deprecated in 10.11)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// Deprecated: This function is deprecated.


// CMTimebaseGetMasterTimebase(timebase _, :  CMTimebase) ->  CMTimebase?) func
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 6.0+ (Deprecated in 9.0)
//   - iPadOS 6.0+ (Deprecated in 9.0)
//   - macOS 10.8+ (Deprecated in 10.11)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// Deprecated: This function is deprecated.

// CMTimebaseGetRate(timebase _, :  CMTimebase) ->  Float64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimebaseGetTime(timebase _, :  CMTimebase) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTimebaseGetTimeAndRate(timebase _, :  CMTimebase,  timeOut:  UnsafeMutablePointer< CMTime>?,  rateOut:  UnsafeMutablePointer< Float64>?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimebaseGetTimeWithTimeScale(timebase _, :  CMTimebase,  timescale:  CMTimeScale,  method:  CMTimeRoundingMethod) ->  CMTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimebaseGetTypeID() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTimebaseGetUltimateMasterClock(timebase _, :  CMTimebase) ->  CMClock?) func
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 6.0+ (Deprecated in 9.0)
//   - iPadOS 6.0+ (Deprecated in 9.0)
//   - macOS 10.8+ (Deprecated in 10.11)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// Deprecated: This function is deprecated.

// CMTimebaseNotificationBarrier(timebase _, :  CMTimebase) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimebaseRemoveTimer(timebase _, :  CMTimebase,  timer:  CFRunLoopTimer) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTimebaseRemoveTimerDispatchSource(timebase _, :  CMTimebase,  timerSource:  dispatch_source_t) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimebaseSetAnchorTime(timebase _, :  CMTimebase,  timebaseTime:  CMTime,  immediateSourceTime:  CMTime) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimebaseSetMasterClock(timebase _, newMasterClock :  CMTimebase,  _, :  CMClock) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 6.0+ (Deprecated in 8.0)
//   - iPadOS 6.0+ (Deprecated in 8.0)
//   - macOS 10.8+ (Deprecated in 10.10)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 6.0+ (Deprecated in 6.0)
//
// Deprecated: This function is deprecated.


// CMTimebaseSetMasterTimebase(timebase _, newMasterTimebase :  CMTimebase,  _, :  CMTimebase) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 6.0+ (Deprecated in 8.0)
//   - iPadOS 6.0+ (Deprecated in 8.0)
//   - macOS 10.8+ (Deprecated in 10.10)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 6.0+ (Deprecated in 6.0)
//
// Deprecated: This function is deprecated.

// CMTimebaseSetRate(timebase _, :  CMTimebase,  rate:  Float64) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimebaseSetRateAndAnchorTime(timebase _, timebaseTime :  CMTimebase,  rate:  Float64,  anchorTime, :  CMTime,  immediateSourceTime:  CMTime) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTimebaseSetSourceClock(timebase _, newSourceClock :  CMTimebase,  _, :  CMClock) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimebaseSetSourceTimebase(timebase _, newSourceTimebase :  CMTimebase,  _, :  CMTimebase) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimebaseSetTime(timebase _, :  CMTimebase,  time:  CMTime) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTimebaseSetTimerDispatchSourceNextFireTime(timebase _, :  CMTimebase,  timerSource:  dispatch_source_t,  fireTime:  CMTime,  flags:  UInt32) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimebaseSetTimerDispatchSourceToFireImmediately(timebase _, :  CMTimebase,  timerSource:  dispatch_source_t) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMTimebaseSetTimerNextFireTime(timebase _, :  CMTimebase,  timer:  CFRunLoopTimer,  fireTime:  CMTime,  flags:  UInt32) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMTimebaseSetTimerToFireImmediately(timebase _, :  CMTimebase,  timer:  CFRunLoopTimer) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer(allocator:  CFAllocator?,  videoFormatDescription:  CMVideoFormatDescription,  stringEncoding:  CFStringEncoding,  flavor:  CMImageDescriptionFlavor?,  blockBufferOut:  UnsafeMutablePointer< CMBlockBuffer?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMVideoFormatDescriptionCopyTagCollectionArray(formatDescription CMVideoFormatDescriptionRef, tagCollectionsOut ,  CFArrayRef *, );) extern   OSStatus
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+


// CMVideoFormatDescriptionCreate(allocator:  CFAllocator?,  codecType:  CMVideoCodecType,  width:  Int32,  height:  Int32,  extensions:  CFDictionary?,  formatDescriptionOut:  UnsafeMutablePointer< CMVideoFormatDescription?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMVideoFormatDescriptionCreateForImageBuffer(allocator:  CFAllocator?,  imageBuffer:  CVImageBuffer,  formatDescriptionOut:  UnsafeMutablePointer< CMVideoFormatDescription?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer(imageDescriptionBlockBuffer allocator:  CFAllocator?,  bigEndianImageDescriptionBlockBuffer, :  CMBlockBuffer,  stringEncoding:  CFStringEncoding,  flavor:  CMImageDescriptionFlavor?,  formatDescriptionOut:  UnsafeMutablePointer< CMVideoFormatDescription?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData(imageDescriptionData allocator:  CFAllocator?,  bigEndianImageDescriptionData, :  UnsafePointer< UInt8>,  size:  Int,  stringEncoding:  CFStringEncoding,  flavor:  CMImageDescriptionFlavor?,  formatDescriptionOut:  UnsafeMutablePointer< CMVideoFormatDescription?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMVideoFormatDescriptionCreateFromH264ParameterSets(NALUnitHeaderLength allocator:  CFAllocator?,  parameterSetCount:  Int,  parameterSetPointers:  UnsafePointer< UnsafePointer< UInt8>>,  parameterSetSizes:  UnsafePointer< Int>,  nalUnitHeaderLength, :  Int32,  formatDescriptionOut:  UnsafeMutablePointer< CMFormatDescription?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMVideoFormatDescriptionCreateFromHEVCParameterSets(NALUnitHeaderLength allocator:  CFAllocator?,  parameterSetCount:  Int,  parameterSetPointers:  UnsafePointer< UnsafePointer< UInt8>>,  parameterSetSizes:  UnsafePointer< Int>,  nalUnitHeaderLength, :  Int32,  extensions:  CFDictionary?,  formatDescriptionOut:  UnsafeMutablePointer< CMFormatDescription?>) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.13+
//   - tvOS 11.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMVideoFormatDescriptionGetCleanAperture(videoDesc _, :  CMVideoFormatDescription,  originIsAtTopLeft:  Bool) ->  CGRect) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMVideoFormatDescriptionGetDimensions(videoDesc _, :  CMVideoFormatDescription) ->  CMVideoDimensions) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMVideoFormatDescriptionGetH264ParameterSetAtIndex(videoDesc _, NALUnitHeaderLengthOut :  CMFormatDescription,  parameterSetIndex:  Int,  parameterSetPointerOut:  UnsafeMutablePointer< UnsafePointer< UInt8>?>?,  parameterSetSizeOut:  UnsafeMutablePointer< Int>?,  parameterSetCountOut:  UnsafeMutablePointer< Int>?,  nalUnitHeaderLengthOut, :  UnsafeMutablePointer< Int32>?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMVideoFormatDescriptionGetHEVCParameterSetAtIndex(videoDesc _, NALUnitHeaderLengthOut :  CMFormatDescription,  parameterSetIndex:  Int,  parameterSetPointerOut:  UnsafeMutablePointer< UnsafePointer< UInt8>?>?,  parameterSetSizeOut:  UnsafeMutablePointer< Int>?,  parameterSetCountOut:  UnsafeMutablePointer< Int>?,  nalUnitHeaderLengthOut, :  UnsafeMutablePointer< Int32>?) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.13+
//   - tvOS 11.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CMVideoFormatDescriptionGetPresentationDimensions(videoDesc _, :  CMVideoFormatDescription,  usePixelAspectRatio:  Bool,  useCleanAperture:  Bool) ->  CGSize) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+


// CMVideoFormatDescriptionMatchesImageBuffer(desc _, :  CMVideoFormatDescription,  imageBuffer:  CVImageBuffer) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

