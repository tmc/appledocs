// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CIRectangleFeature */


/* debug [class_header]: Header for CIRectangleFeature */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RectangleFeature */
// An interface definition for the [RectangleFeature] class.
type IRectangleFeature interface {
	IFeature
	
/* debug [class_interface_properties]: Properties for RectangleFeature */
	// properties:
	BottomLeft() corefoundation.CGPoint
	BottomRight() corefoundation.CGPoint
	Bounds() corefoundation.CGRect
	TopLeft() corefoundation.CGPoint
	TopRight() corefoundation.CGPoint
	CIDetectorTypeRectangle() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RectangleFeature */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RectangleFeature */
// Alloc allocates a new instance without initialization.
func (rc _RectangleFeatureClass) Alloc() RectangleFeature {
	rv := objc.Send[RectangleFeature](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RectangleFeature */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RectangleFeature *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RectangleFeature */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RectangleFeature */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RectangleFeature */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RectangleFeature */

// The lower-left corner of the detected rectangle, in image coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature/bottomLeft-swift.property
func (r_ RectangleFeature) BottomLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](r_.ID, objc.Sel("bottomLeft"))
	return rv
}/* debug [instance_properties/getter]: bottomLeft */


// The lower-right corner of the detected rectangle, in image coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature/bottomRight-swift.property
func (r_ RectangleFeature) BottomRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](r_.ID, objc.Sel("bottomRight"))
	return rv
}/* debug [instance_properties/getter]: bottomRight */


// A rectangle indicating the position and extent of the feature in image coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature/bounds-swift.property
func (r_ RectangleFeature) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](r_.ID, objc.Sel("bounds"))
	return rv
}/* debug [instance_properties/getter]: bounds */


// The upper-left corner of the detected rectangle, in image coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature/topLeft-swift.property
func (r_ RectangleFeature) TopLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](r_.ID, objc.Sel("topLeft"))
	return rv
}/* debug [instance_properties/getter]: topLeft */


// The upper-right corner of the detected rectangle, in image coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIRectangleFeature/topRight-swift.property
func (r_ RectangleFeature) TopRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](r_.ID, objc.Sel("topRight"))
	return rv
}/* debug [instance_properties/getter]: topRight */


// A detector that searches for rectangular areas in a still image or video, returning
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidetectortyperectangle
func (r_ RectangleFeature) CIDetectorTypeRectangle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("CIDetectorTypeRectangle"))
	return rv
}/* debug [instance_properties/getter]: CIDetectorTypeRectangle */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CIRectangleFeature */



