// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"unsafe"

	"github.com/ebitengine/purego"
	coregraphics "github.com/tmc/appledocs/generated/coregraphics"
)


// Vision Functions (12 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_VNElementTypeSize func(unsafe.Pointer) unsafe.Pointer
	_VNImagePointForFaceLandmarkPoint func(unsafe.Pointer, coregraphics.CGRect, unsafe.Pointer, unsafe.Pointer) coregraphics.CGPoint
	_VNImagePointForNormalizedPoint func(coregraphics.CGPoint, unsafe.Pointer, unsafe.Pointer) coregraphics.CGPoint
	_VNImagePointForNormalizedPointUsingRegionOfInterest func(coregraphics.CGPoint, unsafe.Pointer, unsafe.Pointer, coregraphics.CGRect) coregraphics.CGPoint
	_VNImageRectForNormalizedRect func(coregraphics.CGRect, unsafe.Pointer, unsafe.Pointer) coregraphics.CGRect
	_VNImageRectForNormalizedRectUsingRegionOfInterest func(coregraphics.CGRect, unsafe.Pointer, unsafe.Pointer, coregraphics.CGRect) coregraphics.CGRect
	_VNNormalizedFaceBoundingBoxPointForLandmarkPoint func(unsafe.Pointer, coregraphics.CGRect, unsafe.Pointer, unsafe.Pointer) coregraphics.CGPoint
	_VNNormalizedPointForImagePoint func(coregraphics.CGPoint, unsafe.Pointer, unsafe.Pointer) coregraphics.CGPoint
	_VNNormalizedPointForImagePointUsingRegionOfInterest func(coregraphics.CGPoint, unsafe.Pointer, unsafe.Pointer, coregraphics.CGRect) coregraphics.CGPoint
	_VNNormalizedRectForImageRect func(coregraphics.CGRect, unsafe.Pointer, unsafe.Pointer) coregraphics.CGRect
	_VNNormalizedRectForImageRectUsingRegionOfInterest func(coregraphics.CGRect, unsafe.Pointer, unsafe.Pointer, coregraphics.CGRect) coregraphics.CGRect
	_VNNormalizedRectIsIdentityRect func(coregraphics.CGRect) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_VNElementTypeSize, lib, "VNElementTypeSize")
	tryRegister(&_VNImagePointForFaceLandmarkPoint, lib, "VNImagePointForFaceLandmarkPoint")
	tryRegister(&_VNImagePointForNormalizedPoint, lib, "VNImagePointForNormalizedPoint")
	tryRegister(&_VNImagePointForNormalizedPointUsingRegionOfInterest, lib, "VNImagePointForNormalizedPointUsingRegionOfInterest")
	tryRegister(&_VNImageRectForNormalizedRect, lib, "VNImageRectForNormalizedRect")
	tryRegister(&_VNImageRectForNormalizedRectUsingRegionOfInterest, lib, "VNImageRectForNormalizedRectUsingRegionOfInterest")
	tryRegister(&_VNNormalizedFaceBoundingBoxPointForLandmarkPoint, lib, "VNNormalizedFaceBoundingBoxPointForLandmarkPoint")
	tryRegister(&_VNNormalizedPointForImagePoint, lib, "VNNormalizedPointForImagePoint")
	tryRegister(&_VNNormalizedPointForImagePointUsingRegionOfInterest, lib, "VNNormalizedPointForImagePointUsingRegionOfInterest")
	tryRegister(&_VNNormalizedRectForImageRect, lib, "VNNormalizedRectForImageRect")
	tryRegister(&_VNNormalizedRectForImageRectUsingRegionOfInterest, lib, "VNNormalizedRectForImageRectUsingRegionOfInterest")
	tryRegister(&_VNNormalizedRectIsIdentityRect, lib, "VNNormalizedRectIsIdentityRect")
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



// Returns the size of a feature print element. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNElementTypeSize(_:)
func VNElementTypeSize(elementType unsafe.Pointer) unsafe.Pointer {
	return _VNElementTypeSize(elementType)
	}


// Returns the image coordinates of a specified face landmark point. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImagePointForFaceLandmarkPoint(_:_:_:_:)
func VNImagePointForFaceLandmarkPoint(faceLandmarkPoint unsafe.Pointer, faceBoundingBox coregraphics.CGRect, imageWidth unsafe.Pointer, imageHeight unsafe.Pointer) coregraphics.CGPoint {
	return _VNImagePointForFaceLandmarkPoint(faceLandmarkPoint, faceBoundingBox, imageWidth, imageHeight)
	}


// Projects a point in normalized coordinates into image coordinates. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImagePointForNormalizedPoint(_:_:_:)
func VNImagePointForNormalizedPoint(normalizedPoint coregraphics.CGPoint, imageWidth unsafe.Pointer, imageHeight unsafe.Pointer) coregraphics.CGPoint {
	return _VNImagePointForNormalizedPoint(normalizedPoint, imageWidth, imageHeight)
	}


// Projects a point from a region of interest within the normalized coordinates into image coordinates. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImagePointForNormalizedPointUsingRegionOfInterest(_:_:_:_:)
func VNImagePointForNormalizedPointUsingRegionOfInterest(normalizedPoint coregraphics.CGPoint, imageWidth unsafe.Pointer, imageHeight unsafe.Pointer, roi coregraphics.CGRect) coregraphics.CGPoint {
	return _VNImagePointForNormalizedPointUsingRegionOfInterest(normalizedPoint, imageWidth, imageHeight, roi)
	}


// Projects a rectangle from normalized coordinates into image coordinates. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageRectForNormalizedRect(_:_:_:)
func VNImageRectForNormalizedRect(normalizedRect coregraphics.CGRect, imageWidth unsafe.Pointer, imageHeight unsafe.Pointer) coregraphics.CGRect {
	return _VNImageRectForNormalizedRect(normalizedRect, imageWidth, imageHeight)
	}


// Projects a rectangle from a region of interest within the normalized coordinates into image coordinates. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageRectForNormalizedRectUsingRegionOfInterest(_:_:_:_:)
func VNImageRectForNormalizedRectUsingRegionOfInterest(normalizedRect coregraphics.CGRect, imageWidth unsafe.Pointer, imageHeight unsafe.Pointer, roi coregraphics.CGRect) coregraphics.CGRect {
	return _VNImageRectForNormalizedRectUsingRegionOfInterest(normalizedRect, imageWidth, imageHeight, roi)
	}


// Returns the coordinates of a specified face landmark point, in bounding box coordinates. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNNormalizedFaceBoundingBoxPointForLandmarkPoint(_:_:_:_:)
func VNNormalizedFaceBoundingBoxPointForLandmarkPoint(faceLandmarkPoint unsafe.Pointer, faceBoundingBox coregraphics.CGRect, imageWidth unsafe.Pointer, imageHeight unsafe.Pointer) coregraphics.CGPoint {
	return _VNNormalizedFaceBoundingBoxPointForLandmarkPoint(faceLandmarkPoint, faceBoundingBox, imageWidth, imageHeight)
	}


// Projects a point from image coordinates into normalized coordinates. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNNormalizedPointForImagePoint(_:_:_:)
func VNNormalizedPointForImagePoint(imagePoint coregraphics.CGPoint, imageWidth unsafe.Pointer, imageHeight unsafe.Pointer) coregraphics.CGPoint {
	return _VNNormalizedPointForImagePoint(imagePoint, imageWidth, imageHeight)
	}


// Projects a point from a region of interest within the image coordinates into normalized coordinates. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNNormalizedPointForImagePointUsingRegionOfInterest(_:_:_:_:)
func VNNormalizedPointForImagePointUsingRegionOfInterest(imagePoint coregraphics.CGPoint, imageWidth unsafe.Pointer, imageHeight unsafe.Pointer, roi coregraphics.CGRect) coregraphics.CGPoint {
	return _VNNormalizedPointForImagePointUsingRegionOfInterest(imagePoint, imageWidth, imageHeight, roi)
	}


// Projects a rectangle from image coordinates into normalized coordinates. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNNormalizedRectForImageRect(_:_:_:)
func VNNormalizedRectForImageRect(imageRect coregraphics.CGRect, imageWidth unsafe.Pointer, imageHeight unsafe.Pointer) coregraphics.CGRect {
	return _VNNormalizedRectForImageRect(imageRect, imageWidth, imageHeight)
	}


// Projects a rectangle from a region of interest within the image coordinates space into normalized coordinates. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNNormalizedRectForImageRectUsingRegionOfInterest(_:_:_:_:)
func VNNormalizedRectForImageRectUsingRegionOfInterest(imageRect coregraphics.CGRect, imageWidth unsafe.Pointer, imageHeight unsafe.Pointer, roi coregraphics.CGRect) coregraphics.CGRect {
	return _VNNormalizedRectForImageRectUsingRegionOfInterest(imageRect, imageWidth, imageHeight, roi)
	}


// Returns a Boolean value that indicates whether the rectangle has an origin of zero and unit length and width. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNNormalizedRectIsIdentityRect(_:)
func VNNormalizedRectIsIdentityRect(normalizedRect coregraphics.CGRect) unsafe.Pointer {
	return _VNNormalizedRectIsIdentityRect(normalizedRect)
	}




