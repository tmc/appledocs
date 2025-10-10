// Code generated from Apple documentation for ImageIO. DO NOT EDIT.

package imageio

// ImageIO Functions
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

// Discovered functions (20 total):

// CGAnimateImageAtURLWithBlock(url _, options :  CFURL,  _, block :  CFDictionary?,  _, :  @escaping  CGImageSourceAnimationBlock) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CGAnimateImageDataWithBlock(data _, options :  CFData,  _, block :  CFDictionary?,  _, :  @escaping  CGImageSourceAnimationBlock) ->  OSStatus) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// CGImageDestinationAddImageAndMetadata(idst _, image :  CGImageDestination,  _, metadata :  CGImage,  _, options :  CGImageMetadata?,  _, :  CFDictionary?)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGImageDestinationCopyImageSource(idst _, isrc :  CGImageDestination,  _, options :  CGImageSource,  _, err :  CFDictionary?,  _, :  UnsafeMutablePointer< Unmanaged< CFError>?>?) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGImageDestinationFinalize(idst _, :  CGImageDestination) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.4+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGImageDestinationGetTypeID() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.4+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGImageMetadataCopyTagWithPath(metadata _, parent :  CGImageMetadata,  _, path :  CGImageMetadataTag?,  _, :  CFString) ->  CGImageMetadataTag?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGImageMetadataCreateMutableCopy(metadata _, :  CGImageMetadata) ->  CGMutableImageMetadata?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGImageMetadataGetTypeID() func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGImageMetadataSetTagWithPath(metadata _, parent :  CGMutableImageMetadata,  _, path :  CGImageMetadataTag?,  _, tag :  CFString,  _, :  CGImageMetadataTag) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGImageMetadataSetValueWithPath(metadata _, parent :  CGMutableImageMetadata,  _, path :  CGImageMetadataTag?,  _, value :  CFString,  _, :  CFTypeRef) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGImageMetadataTagCopyName(tag _, :  CGImageMetadataTag) ->  CFString?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGImageMetadataTagCopyQualifiers(tag _, :  CGImageMetadataTag) ->  CFArray?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGImageMetadataTagCopyValue(tag _, :  CGImageMetadataTag) ->  CFTypeRef?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGImageMetadataTagCreate(xmlns _, prefix :  CFString,  _, name :  CFString?,  _, type :  CFString,  _, value :  CGImageMetadataType,  _, :  CFTypeRef) ->  CGImageMetadataTag?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGImageSourceCopyProperties(isrc _, options :  CGImageSource,  _, :  CFDictionary?) ->  CFDictionary?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.4+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGImageSourceCopyPropertiesAtIndex(isrc _, index :  CGImageSource,  _, options :  Int,  _, :  CFDictionary?) ->  CFDictionary?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.4+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGImageSourceCopyTypeIdentifiers() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.4+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CGImageSourceCreateWithURL(url _, options :  CFURL,  _, :  CFDictionary?) ->  CGImageSource?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.4+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CGImageSourceRemoveCacheAtIndex(isrc _, index :  CGImageSource,  _, :  Int) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

