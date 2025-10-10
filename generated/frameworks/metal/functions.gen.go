// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

// Metal Functions
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

// Discovered functions (13 total):

// MTL4BufferRangeMake(bufferAddress _, length :  MTLGPUAddress,  _, :  UInt64) ->  MTL4BufferRange) func

// MTLClearColorMake(red _, green :  Double,  _, blue :  Double,  _, alpha :  Double,  _, :  Double) ->  MTLClearColor) func

// MTLCoordinate2DMake(x _, y :  Float,  _, :  Float) ->  MTLCoordinate2D) func


// MTLCopyAllDevices() func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 10.11+
//   - tvOS 18.0+
//   - visionOS 2.0+

// MTLCopyAllDevicesWithObserver(observer id<NSObject> *, handler ,  MTLDeviceNotificationHandler, );) extern   NSArray<id<MTLDevice>>  *
//
// Availability:
//   - macOS 10.13+

// MTLCreateSystemDefaultDevice() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.11+
//   - visionOS 1.0+


// MTLIOCompressionContextAppendData(context _, data :  MTLIOCompressionContext,  _, size :  UnsafeRawPointer,  _, :  Int) func
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+

// MTLIOCompressionContextDefaultChunkSize() func

// MTLIOCreateCompressionContext(path const  char *, type ,  MTLIOCompressionMethod, chunkSize ,  size_t, );) extern   MTLIOCompressionContext
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+


// MTLIOFlushAndDestroyCompressionContext(context _, :  MTLIOCompressionContext) ->  MTLIOCompressionStatus) func
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+

// MTLIndirectCommandBufferExecutionRangeMake(location _, length :  UInt32,  _, :  UInt32) ->  MTLIndirectCommandBufferExecutionRange) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.14+
//   - tvOS 13.0+
//   - visionOS 1.0+

// MTLPackedFloat3Make(x _, y :  Float,  _, z :  Float,  _, :  Float) ->  MTLPackedFloat3) func


// MTLPackedFloatQuaternionMake(x float, y ,  float, z ,  float, w ,  float, );) static   MTLPackedFloatQuaternion

