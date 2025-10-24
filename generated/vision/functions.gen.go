// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision


import (
	"unsafe"

	"github.com/ebitengine/purego"
	corefoundation "github.com/tmc/appledocs/generated/corefoundation"
)


// Vision Functions (12 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_VNElementTypeSize func(ElementType) uint
	_VNImagePointForFaceLandmarkPoint func(unsafe.Pointer, corefoundation.CGRect, uintptr, uintptr) corefoundation.CGPoint
	_VNImagePointForNormalizedPoint func(corefoundation.CGPoint, uintptr, uintptr) corefoundation.CGPoint
	_VNImagePointForNormalizedPointUsingRegionOfInterest func(corefoundation.CGPoint, uintptr, uintptr, corefoundation.CGRect) corefoundation.CGPoint
	_VNImageRectForNormalizedRect func(corefoundation.CGRect, uintptr, uintptr) corefoundation.CGRect
	_VNImageRectForNormalizedRectUsingRegionOfInterest func(corefoundation.CGRect, uintptr, uintptr, corefoundation.CGRect) corefoundation.CGRect
	_VNNormalizedFaceBoundingBoxPointForLandmarkPoint func(unsafe.Pointer, corefoundation.CGRect, uintptr, uintptr) corefoundation.CGPoint
	_VNNormalizedPointForImagePoint func(corefoundation.CGPoint, uintptr, uintptr) corefoundation.CGPoint
	_VNNormalizedPointForImagePointUsingRegionOfInterest func(corefoundation.CGPoint, uintptr, uintptr, corefoundation.CGRect) corefoundation.CGPoint
	_VNNormalizedRectForImageRect func(corefoundation.CGRect, uintptr, uintptr) corefoundation.CGRect
	_VNNormalizedRectForImageRectUsingRegionOfInterest func(corefoundation.CGRect, uintptr, uintptr, corefoundation.CGRect) corefoundation.CGRect
	_VNNormalizedRectIsIdentityRect func(corefoundation.CGRect) bool
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



// Returns the size of a feature print element.
//
// Added in macOS 10.15.
// Returns the size of a feature print element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNElementTypeSize(_:)
func VNElementTypeSize(elementType ElementType) uint {
	return _VNElementTypeSize(elementType)
}

// Returns the image coordinates of a specified face landmark point.
//
// Added in macOS 10.13.
// Returns the image coordinates of a specified face landmark point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImagePointForFaceLandmarkPoint(_:_:_:_:)
func VNImagePointForFaceLandmarkPoint(faceLandmarkPoint unsafe.Pointer, faceBoundingBox corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr) corefoundation.CGPoint {
	return _VNImagePointForFaceLandmarkPoint(faceLandmarkPoint, faceBoundingBox, imageWidth, imageHeight)
}

// Projects a point in normalized coordinates into image coordinates.
//
// Added in macOS 10.13.
// Projects a point in normalized coordinates into image coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImagePointForNormalizedPoint(_:_:_:)
func VNImagePointForNormalizedPoint(normalizedPoint corefoundation.CGPoint, imageWidth uintptr, imageHeight uintptr) corefoundation.CGPoint {
	return _VNImagePointForNormalizedPoint(normalizedPoint, imageWidth, imageHeight)
}

// Projects a point from a region of interest within the normalized coordinates into image coordinates.
//
// Added in macOS 12.0.
// Projects a point from a region of interest within the normalized coordinates into image coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImagePointForNormalizedPointUsingRegionOfInterest(_:_:_:_:)
func VNImagePointForNormalizedPointUsingRegionOfInterest(normalizedPoint corefoundation.CGPoint, imageWidth uintptr, imageHeight uintptr, roi corefoundation.CGRect) corefoundation.CGPoint {
	return _VNImagePointForNormalizedPointUsingRegionOfInterest(normalizedPoint, imageWidth, imageHeight, roi)
}

// Projects a rectangle from normalized coordinates into image coordinates.
//
// Added in macOS 10.13.
// Projects a rectangle from normalized coordinates into image coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageRectForNormalizedRect(_:_:_:)
func VNImageRectForNormalizedRect(normalizedRect corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr) corefoundation.CGRect {
	return _VNImageRectForNormalizedRect(normalizedRect, imageWidth, imageHeight)
}

// Projects a rectangle from a region of interest within the normalized coordinates into image coordinates.
//
// Added in macOS 12.0.
// Projects a rectangle from a region of interest within the normalized coordinates into image coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNImageRectForNormalizedRectUsingRegionOfInterest(_:_:_:_:)
func VNImageRectForNormalizedRectUsingRegionOfInterest(normalizedRect corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr, roi corefoundation.CGRect) corefoundation.CGRect {
	return _VNImageRectForNormalizedRectUsingRegionOfInterest(normalizedRect, imageWidth, imageHeight, roi)
}

// Returns the coordinates of a specified face landmark point, in bounding box coordinates.
//
// Added in macOS 10.13.
// Returns the coordinates of a specified face landmark point, in bounding box coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNNormalizedFaceBoundingBoxPointForLandmarkPoint(_:_:_:_:)
func VNNormalizedFaceBoundingBoxPointForLandmarkPoint(faceLandmarkPoint unsafe.Pointer, faceBoundingBox corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr) corefoundation.CGPoint {
	return _VNNormalizedFaceBoundingBoxPointForLandmarkPoint(faceLandmarkPoint, faceBoundingBox, imageWidth, imageHeight)
}

// Projects a point from image coordinates into normalized coordinates.
//
// Added in macOS 11.0.
// Projects a point from image coordinates into normalized coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNNormalizedPointForImagePoint(_:_:_:)
func VNNormalizedPointForImagePoint(imagePoint corefoundation.CGPoint, imageWidth uintptr, imageHeight uintptr) corefoundation.CGPoint {
	return _VNNormalizedPointForImagePoint(imagePoint, imageWidth, imageHeight)
}

// Projects a point from a region of interest within the image coordinates into normalized coordinates.
//
// Added in macOS 12.0.
// Projects a point from a region of interest within the image coordinates into normalized coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNNormalizedPointForImagePointUsingRegionOfInterest(_:_:_:_:)
func VNNormalizedPointForImagePointUsingRegionOfInterest(imagePoint corefoundation.CGPoint, imageWidth uintptr, imageHeight uintptr, roi corefoundation.CGRect) corefoundation.CGPoint {
	return _VNNormalizedPointForImagePointUsingRegionOfInterest(imagePoint, imageWidth, imageHeight, roi)
}

// Projects a rectangle from image coordinates into normalized coordinates.
//
// Added in macOS 10.13.
// Projects a rectangle from image coordinates into normalized coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNNormalizedRectForImageRect(_:_:_:)
func VNNormalizedRectForImageRect(imageRect corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr) corefoundation.CGRect {
	return _VNNormalizedRectForImageRect(imageRect, imageWidth, imageHeight)
}

// Projects a rectangle from a region of interest within the image coordinates space into normalized coordinates.
//
// Added in macOS 12.0.
// Projects a rectangle from a region of interest within the image coordinates space into normalized coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNNormalizedRectForImageRectUsingRegionOfInterest(_:_:_:_:)
func VNNormalizedRectForImageRectUsingRegionOfInterest(imageRect corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr, roi corefoundation.CGRect) corefoundation.CGRect {
	return _VNNormalizedRectForImageRectUsingRegionOfInterest(imageRect, imageWidth, imageHeight, roi)
}

// Returns a Boolean value that indicates whether the rectangle has an origin of zero and unit length and width.
//
// Added in macOS 10.13.
// Returns a Boolean value that indicates whether the rectangle has an origin of zero and unit length and width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNNormalizedRectIsIdentityRect(_:)
func VNNormalizedRectIsIdentityRect(normalizedRect corefoundation.CGRect) bool {
	return _VNNormalizedRectIsIdentityRect(normalizedRect)
}




