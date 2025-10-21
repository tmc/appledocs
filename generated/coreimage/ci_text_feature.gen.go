// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

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

// An interface definition for the [TextFeature] class.
type ITextFeature interface {
	IFeature
}

// Information about a text that was detected in a still or video image.
//
// A detected text feature is not necessarily rectangular in the plane of the image; rather, the feature identifies a shape that may be rectangular in space (for example a text on a sign) but which appears as a four-sided polygon in the image. The properties of a object identify its four corners in image coordinates. To detect text in an image or video, choose the type when initializing a object, and use the option to specify the desired orientation for finding upright text.
//
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

// Alloc allocates a new instance without initialization.
func (tc _TextFeatureClass) Alloc() TextFeature {
	rv := objc.Send[TextFeature](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A detector that searches for text in a still image or video, returning
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidetectortypetext
func (t_ TextFeature) CIDetectorTypeText() string {
	rv := objc.Send[string](t_.ID, objc.Sel("CIDetectorTypeText"))
	return rv
}

// The image coordinate of the lower-left corner of the detected text.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CITextFeature/bottomLeft
func (t_ TextFeature) BottomLeft() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](t_.ID, objc.Sel("bottomLeft"))
	return rv
}

// The image coordinate of the lower-right corner of the detected text.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CITextFeature/bottomRight
func (t_ TextFeature) BottomRight() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](t_.ID, objc.Sel("bottomRight"))
	return rv
}

// A rectangle that indicates the position and extent of the text feature in image coordinates.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CITextFeature/bounds
func (t_ TextFeature) Bounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](t_.ID, objc.Sel("bounds"))
	return rv
}

// An array containing additional features detected within the feature.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CITextFeature/subFeatures
func (t_ TextFeature) SubFeatures() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("subFeatures"))
	return rv
}

// The image coordinate of the upper-left corner of the detected text.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CITextFeature/topLeft
func (t_ TextFeature) TopLeft() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](t_.ID, objc.Sel("topLeft"))
	return rv
}

// The image coordinate of the upper-right corner of the detected text.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CITextFeature/topRight
func (t_ TextFeature) TopRight() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](t_.ID, objc.Sel("topRight"))
	return rv
}



