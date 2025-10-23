// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [RectangleFeature] class.
var (
	RectangleFeatureClass     _RectangleFeatureClass
	RectangleFeatureClassOnce sync.Once
)

func getRectangleFeatureClass() _RectangleFeatureClass {
	RectangleFeatureClassOnce.Do(func() {
		RectangleFeatureClass = _RectangleFeatureClass{objc.GetClass("CIRectangleFeature")}
	})
	return RectangleFeatureClass
}

type _RectangleFeatureClass struct {
	class objc.Class
}

// An interface definition for the [RectangleFeature] class.
type IRectangleFeature interface {
	IFeature
	// properties:
	BottomLeft() coregraphics.CGPoint
	BottomRight() coregraphics.CGPoint
	Bounds() coregraphics.CGRect
	TopLeft() coregraphics.CGPoint
	TopRight() coregraphics.CGPoint
	CIDetectorTypeRectangle() string /* primitive/slice/pointer. */
	// methods:
}

// Information about a rectangular region detected in a still or video image.
//
// A detected rectangle feature is not necessarily rectangular in the plane of the image; rather, the feature identifies a shape that may be rectangular in space (for example a book on a desk) but which appears as a four-sided polygon in the image. The properties of a object identify its four corners in image coordinates. You can use rectangle feature detection together with the filter to transform the feature to a normal orientation. To detect rectangles in an image or video, choose when initializing a object, and use the and options to specify the approximate shape of rectangular features to search for. The detector returns at most one rectangle feature, the most prominent found in the image.


// Information about a rectangular region detected in a still or video image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature
type RectangleFeature struct {
	Feature
}

// RectangleFeatureFrom constructs a [RectangleFeature] from an unsafe.Pointer.
//
// Information about a rectangular region detected in a still or video image.
func RectangleFeatureFrom(ptr unsafe.Pointer) RectangleFeature {
	return RectangleFeature{
		Feature: FeatureFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RectangleFeatureClass) Alloc() RectangleFeature {
	rv := objc.Send[RectangleFeature](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RectangleFeatureClass) New() RectangleFeature {
	rv := objc.Send[RectangleFeature](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RectangleFeature) Init() RectangleFeature {
	rv := objc.Send[RectangleFeature](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RectangleFeature) Autorelease() RectangleFeature {
	rv := objc.Send[RectangleFeature](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRectangleFeature creates a new RectangleFeature instance.
func NewRectangleFeature() RectangleFeature {
	return getRectangleFeatureClass().New()
}



// The lower-left corner of the detected rectangle, in image coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature/bottomLeft-swift.property
func (r_ RectangleFeature) BottomLeft() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](r_.ID, objc.Sel("bottomLeft"))
	return rv
}


// The lower-right corner of the detected rectangle, in image coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature/bottomRight-swift.property
func (r_ RectangleFeature) BottomRight() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](r_.ID, objc.Sel("bottomRight"))
	return rv
}


// A rectangle indicating the position and extent of the feature in image coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature/bounds-swift.property
func (r_ RectangleFeature) Bounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](r_.ID, objc.Sel("bounds"))
	return rv
}


// The upper-left corner of the detected rectangle, in image coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature/topLeft-swift.property
func (r_ RectangleFeature) TopLeft() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](r_.ID, objc.Sel("topLeft"))
	return rv
}


// The upper-right corner of the detected rectangle, in image coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature/topRight-swift.property
func (r_ RectangleFeature) TopRight() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](r_.ID, objc.Sel("topRight"))
	return rv
}


// A detector that searches for rectangular areas in a still image or video, returning
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidetectortyperectangle
func (r_ RectangleFeature) CIDetectorTypeRectangle() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](r_.ID, objc.Sel("CIDetectorTypeRectangle"))
	return rv
}



