// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

// QuartzCore Functions
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

// Discovered functions (16 total):

// CACurrentMediaTime() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - visionOS 1.0+

// CAFrameRateRangeIsEqualToRange(range CAFrameRateRange, other ,  CAFrameRateRange, );) extern   bool
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+

// CAFrameRateRangeMake(minimum float, maximum ,  float, preferred ,  float, );) extern   CAFrameRateRange
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+

// CATransform3DConcat(a _, b :  CATransform3D,  _, :  CATransform3D) ->  CATransform3D) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+

// CATransform3DEqualToTransform(a _, b :  CATransform3D,  _, :  CATransform3D) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+

// CATransform3DGetAffineTransform(t _, :  CATransform3D) ->  CGAffineTransform) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+

// CATransform3DInvert(t _, :  CATransform3D) ->  CATransform3D) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+

// CATransform3DIsAffine(t _, :  CATransform3D) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+

// CATransform3DIsIdentity(t _, :  CATransform3D) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+

// CATransform3DMakeAffineTransform(m _, :  CGAffineTransform) ->  CATransform3D) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+

// CATransform3DMakeRotation(angle _, x :  CGFloat,  _, y :  CGFloat,  _, z :  CGFloat,  _, :  CGFloat) ->  CATransform3D) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+

// CATransform3DMakeScale(sx _, sy :  CGFloat,  _, sz :  CGFloat,  _, :  CGFloat) ->  CATransform3D) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+

// CATransform3DMakeTranslation(tx _, ty :  CGFloat,  _, tz :  CGFloat,  _, :  CGFloat) ->  CATransform3D) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+

// CATransform3DRotate(t _, angle :  CATransform3D,  _, x :  CGFloat,  _, y :  CGFloat,  _, z :  CGFloat,  _, :  CGFloat) ->  CATransform3D) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+

// CATransform3DScale(t _, sx :  CATransform3D,  _, sy :  CGFloat,  _, sz :  CGFloat,  _, :  CGFloat) ->  CATransform3D) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+

// CATransform3DTranslate(t _, tx :  CATransform3D,  _, ty :  CGFloat,  _, tz :  CGFloat,  _, :  CGFloat) ->  CATransform3D) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
