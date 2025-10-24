// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CIDetector */


/* debug [class_header]: Header for CIDetector */
// The class instance for the [Detector] class.
var (
	DetectorClass     _DetectorClass
	DetectorClassOnce sync.Once
)

func getDetectorClass() _DetectorClass {
	DetectorClassOnce.Do(func() {
		DetectorClass = _DetectorClass{objc.GetClass("CIDetector")}
	})
	return DetectorClass
}

type _DetectorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Detector */
// An interface definition for the [Detector] class.
type IDetector interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Detector */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Detector */
	// methods:
	FeaturesInImage(image ICIImage) []Feature
	FeaturesInImageOptions(image ICIImage, options foundation.IDictionary) []Feature
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Detector */
// Alloc allocates a new instance without initialization.
func (dc _DetectorClass) Alloc() Detector {
	rv := objc.Send[Detector](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DetectorClass) New() Detector {
	rv := objc.Send[Detector](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ Detector) Init() Detector {
	rv := objc.Send[Detector](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ Detector) Autorelease() Detector {
	rv := objc.Send[Detector](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDetector creates a new Detector instance.
func NewDetector() Detector {
	return getDetectorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Detector */
// An image processor that identifies notable features, such as faces and barcodes, in a still image or video.
//
// A object uses image processing to search for and identify notable features (faces, rectangles, and barcodes) in a still image or video. Detected features are represented by objects that provide more information about each feature. This class can maintain many state variables that can impact performance. So for best performance, reuse instances instead of creating new ones.


// An image processor that identifies notable features, such as faces and barcodes, in a still image or video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDetector
type Detector struct {
	objectivec.Object
}

// DetectorFrom constructs a [Detector] from an unsafe.Pointer.
//
// An image processor that identifies notable features, such as faces and barcodes, in a still image or video.
func DetectorFrom(ptr unsafe.Pointer) Detector {
	return Detector{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Detector */

// Creates and returns a configured detector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDetector/init(ofType:context:options:)
func NewDetectorOfTypeContextOptions(type_ objc.IObject /* cross-framework: NSString */, context ICIContext, options foundation.IDictionary) Detector {
	rv := objc.Send[Detector](objc.ID(getDetectorClass().class), objc.Sel("detectorOfType:context:options:"), type_, context, options)
	return rv
}/* debug [class_init_methods/constructor]: NewDetectorOfTypeContextOptions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Detector */

// Creates and returns a configured detector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDetector/init(ofType:context:options:)
func (dc _DetectorClass) DetectorOfTypeContextOptions(type_ objc.IObject /* cross-framework: NSString */, context ICIContext, options foundation.IDictionary) IDetector {
	rv := objc.Send[Detector](objc.ID(dc.class), objc.Sel("detectorOfType:context:options:"), type_, context, options)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DetectorOfTypeContextOptions) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Detector */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Detector */

// Searches for features in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDetector/features(in:)
func (d_ Detector) FeaturesInImage(image ICIImage) []Feature {
	rv := objc.Send[[]Feature](d_.ID, objc.Sel("featuresInImage:"), image)
	return rv
}/* debug [instance_methods/method]: FeaturesInImage */


// Searches for features in an image based on the specified image orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDetector/features(in:options:)
func (d_ Detector) FeaturesInImageOptions(image ICIImage, options foundation.IDictionary) []Feature {
	rv := objc.Send[[]Feature](d_.ID, objc.Sel("featuresInImage:options:"), image, options)
	return rv
}/* debug [instance_methods/method]: FeaturesInImageOptions */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Detector */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CIDetector */


