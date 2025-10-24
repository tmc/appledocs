// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CITextFeature */


/* debug [class_header]: Header for CITextFeature */
// The class instance for the [TextFeature] class.
var (
	TextFeatureClass     _TextFeatureClass
	TextFeatureClassOnce sync.Once
)

func getTextFeatureClass() _TextFeatureClass {
	TextFeatureClassOnce.Do(func() {
		TextFeatureClass = _TextFeatureClass{objc.GetClass("CITextFeature")}
	})
	return TextFeatureClass
}

type _TextFeatureClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextFeature */
// An interface definition for the [TextFeature] class.
type ITextFeature interface {
	IFeature
	
/* debug [class_interface_properties]: Properties for TextFeature */
	// properties:
	BottomLeft() corefoundation.CGPoint
	BottomRight() corefoundation.CGPoint
	Bounds() corefoundation.CGRect
	SubFeatures() objc.IObject /* cross-framework: NSArray */
	TopLeft() corefoundation.CGPoint
	TopRight() corefoundation.CGPoint
	CIDetectorTypeText() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextFeature */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextFeature */
// Alloc allocates a new instance without initialization.
func (tc _TextFeatureClass) Alloc() TextFeature {
	rv := objc.Send[TextFeature](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TextFeatureClass) New() TextFeature {
	rv := objc.Send[TextFeature](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextFeature) Init() TextFeature {
	rv := objc.Send[TextFeature](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextFeature) Autorelease() TextFeature {
	rv := objc.Send[TextFeature](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextFeature creates a new TextFeature instance.
func NewTextFeature() TextFeature {
	return getTextFeatureClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextFeature */
// Information about a text that was detected in a still or video image.
//
// A detected text feature is not necessarily rectangular in the plane of the image; rather, the feature identifies a shape that may be rectangular in space (for example a text on a sign) but which appears as a four-sided polygon in the image. The properties of a object identify its four corners in image coordinates. To detect text in an image or video, choose the type when initializing a object, and use the option to specify the desired orientation for finding upright text.


// Information about a text that was detected in a still or video image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CITextFeature
type TextFeature struct {
	Feature
}

// TextFeatureFrom constructs a [TextFeature] from an unsafe.Pointer.
//
// Information about a text that was detected in a still or video image.
func TextFeatureFrom(ptr unsafe.Pointer) TextFeature {
	return TextFeature{
		Feature: FeatureFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextFeature *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextFeature */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextFeature */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextFeature */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextFeature */

// The image coordinate of the lower-left corner of the detected text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CITextFeature/bottomLeft
func (t_ TextFeature) BottomLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](t_.ID, objc.Sel("bottomLeft"))
	return rv
}/* debug [instance_properties/getter]: bottomLeft */


// The image coordinate of the lower-right corner of the detected text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CITextFeature/bottomRight
func (t_ TextFeature) BottomRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](t_.ID, objc.Sel("bottomRight"))
	return rv
}/* debug [instance_properties/getter]: bottomRight */


// A rectangle that indicates the position and extent of the text feature in image coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CITextFeature/bounds
func (t_ TextFeature) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t_.ID, objc.Sel("bounds"))
	return rv
}/* debug [instance_properties/getter]: bounds */


// An array containing additional features detected within the feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CITextFeature/subFeatures
func (t_ TextFeature) SubFeatures() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](t_.ID, objc.Sel("subFeatures"))
	return rv
}/* debug [instance_properties/getter]: subFeatures */


// The image coordinate of the upper-left corner of the detected text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CITextFeature/topLeft
func (t_ TextFeature) TopLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](t_.ID, objc.Sel("topLeft"))
	return rv
}/* debug [instance_properties/getter]: topLeft */


// The image coordinate of the upper-right corner of the detected text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CITextFeature/topRight
func (t_ TextFeature) TopRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](t_.ID, objc.Sel("topRight"))
	return rv
}/* debug [instance_properties/getter]: topRight */


// A detector that searches for text in a still image or video, returning
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidetectortypetext
func (t_ TextFeature) CIDetectorTypeText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("CIDetectorTypeText"))
	return rv
}/* debug [instance_properties/getter]: CIDetectorTypeText */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CITextFeature */



