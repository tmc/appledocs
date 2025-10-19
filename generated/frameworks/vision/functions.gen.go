// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

// Vision Functions (11 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_VNImagePointForFaceLandmarkPoint func(unsafe.Pointer, unsafe.Pointer, uintptr, uintptr) unsafe.Pointer
	_VNImagePointForNormalizedPoint func(unsafe.Pointer, uintptr, uintptr) unsafe.Pointer
	_VNImagePointForNormalizedPointUsingRegionOfInterest func(unsafe.Pointer, uintptr, uintptr, unsafe.Pointer) unsafe.Pointer
	_VNImageRectForNormalizedRect func(unsafe.Pointer, uintptr, uintptr) unsafe.Pointer
	_VNImageRectForNormalizedRectUsingRegionOfInterest func(unsafe.Pointer, uintptr, uintptr, unsafe.Pointer) unsafe.Pointer
	_VNNormalizedFaceBoundingBoxPointForLandmarkPoint func(unsafe.Pointer, unsafe.Pointer, uintptr, uintptr) unsafe.Pointer
	_VNNormalizedPointForImagePoint func(unsafe.Pointer, uintptr, uintptr) unsafe.Pointer
	_VNNormalizedPointForImagePointUsingRegionOfInterest func(unsafe.Pointer, uintptr, uintptr, unsafe.Pointer) unsafe.Pointer
	_VNNormalizedRectForImageRect func(unsafe.Pointer, uintptr, uintptr) unsafe.Pointer
	_VNNormalizedRectForImageRectUsingRegionOfInterest func(unsafe.Pointer, uintptr, uintptr, unsafe.Pointer) unsafe.Pointer
	_VNNormalizedRectIsIdentityRect func(unsafe.Pointer) bool
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
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


// Returns the image coordinates of a specified face landmark point. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: doc://com.apple.vision/documentation/Vision/VNImagePointForFaceLandmarkPoint(_:_:_:_:)
func VNImagePointForFaceLandmarkPoint(faceLandmarkPoint unsafe.Pointer, faceBoundingBox unsafe.Pointer, imageWidth uintptr, imageHeight uintptr) unsafe.Pointer {
	return _VNImagePointForFaceLandmarkPoint(faceLandmarkPoint, faceBoundingBox, imageWidth, imageHeight)
	}


// Projects a point in normalized coordinates into image coordinates. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: doc://com.apple.vision/documentation/Vision/VNImagePointForNormalizedPoint(_:_:_:)
func VNImagePointForNormalizedPoint(normalizedPoint unsafe.Pointer, imageWidth uintptr, imageHeight uintptr) unsafe.Pointer {
	return _VNImagePointForNormalizedPoint(normalizedPoint, imageWidth, imageHeight)
	}


// Projects a point from a region of interest within the normalized coordinates into image coordinates. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: doc://com.apple.vision/documentation/Vision/VNImagePointForNormalizedPointUsingRegionOfInterest(_:_:_:_:)
func VNImagePointForNormalizedPointUsingRegionOfInterest(normalizedPoint unsafe.Pointer, imageWidth uintptr, imageHeight uintptr, roi unsafe.Pointer) unsafe.Pointer {
	return _VNImagePointForNormalizedPointUsingRegionOfInterest(normalizedPoint, imageWidth, imageHeight, roi)
	}


// Projects a rectangle from normalized coordinates into image coordinates. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: doc://com.apple.vision/documentation/Vision/VNImageRectForNormalizedRect(_:_:_:)
func VNImageRectForNormalizedRect(normalizedRect unsafe.Pointer, imageWidth uintptr, imageHeight uintptr) unsafe.Pointer {
	return _VNImageRectForNormalizedRect(normalizedRect, imageWidth, imageHeight)
	}


// Projects a rectangle from a region of interest within the normalized coordinates into image coordinates. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: doc://com.apple.vision/documentation/Vision/VNImageRectForNormalizedRectUsingRegionOfInterest(_:_:_:_:)
func VNImageRectForNormalizedRectUsingRegionOfInterest(normalizedRect unsafe.Pointer, imageWidth uintptr, imageHeight uintptr, roi unsafe.Pointer) unsafe.Pointer {
	return _VNImageRectForNormalizedRectUsingRegionOfInterest(normalizedRect, imageWidth, imageHeight, roi)
	}


// Returns the coordinates of a specified face landmark point, in bounding box coordinates. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: doc://com.apple.vision/documentation/Vision/VNNormalizedFaceBoundingBoxPointForLandmarkPoint(_:_:_:_:)
func VNNormalizedFaceBoundingBoxPointForLandmarkPoint(faceLandmarkPoint unsafe.Pointer, faceBoundingBox unsafe.Pointer, imageWidth uintptr, imageHeight uintptr) unsafe.Pointer {
	return _VNNormalizedFaceBoundingBoxPointForLandmarkPoint(faceLandmarkPoint, faceBoundingBox, imageWidth, imageHeight)
	}


// Projects a point from image coordinates into normalized coordinates. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: doc://com.apple.vision/documentation/Vision/VNNormalizedPointForImagePoint(_:_:_:)
func VNNormalizedPointForImagePoint(imagePoint unsafe.Pointer, imageWidth uintptr, imageHeight uintptr) unsafe.Pointer {
	return _VNNormalizedPointForImagePoint(imagePoint, imageWidth, imageHeight)
	}


// Projects a point from a region of interest within the image coordinates into normalized coordinates. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: doc://com.apple.vision/documentation/Vision/VNNormalizedPointForImagePointUsingRegionOfInterest(_:_:_:_:)
func VNNormalizedPointForImagePointUsingRegionOfInterest(imagePoint unsafe.Pointer, imageWidth uintptr, imageHeight uintptr, roi unsafe.Pointer) unsafe.Pointer {
	return _VNNormalizedPointForImagePointUsingRegionOfInterest(imagePoint, imageWidth, imageHeight, roi)
	}


// Projects a rectangle from image coordinates into normalized coordinates. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: doc://com.apple.vision/documentation/Vision/VNNormalizedRectForImageRect(_:_:_:)
func VNNormalizedRectForImageRect(imageRect unsafe.Pointer, imageWidth uintptr, imageHeight uintptr) unsafe.Pointer {
	return _VNNormalizedRectForImageRect(imageRect, imageWidth, imageHeight)
	}


// Projects a rectangle from a region of interest within the image coordinates space into normalized coordinates. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: doc://com.apple.vision/documentation/Vision/VNNormalizedRectForImageRectUsingRegionOfInterest(_:_:_:_:)
func VNNormalizedRectForImageRectUsingRegionOfInterest(imageRect unsafe.Pointer, imageWidth uintptr, imageHeight uintptr, roi unsafe.Pointer) unsafe.Pointer {
	return _VNNormalizedRectForImageRectUsingRegionOfInterest(imageRect, imageWidth, imageHeight, roi)
	}


// Returns a Boolean value that indicates whether the rectangle has an origin of zero and unit length and width. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: doc://com.apple.vision/documentation/Vision/VNNormalizedRectIsIdentityRect(_:)
func VNNormalizedRectIsIdentityRect(normalizedRect unsafe.Pointer) bool {
	return _VNNormalizedRectIsIdentityRect(normalizedRect)
	}



