// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [Detector] class.
type IDetector interface {
	objectivec.IObject
	

	// properties:


	

	// methods:
	FeaturesInImage(image ICIImage) []Feature
	FeaturesInImageOptions(image ICIImage, options foundation.IDictionary) []Feature


}





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






// Creates and returns a configured detector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDetector/init(ofType:context:options:)
func NewDetectorOfTypeContextOptions(type_ objc.IObject /* cross-framework: NSString */, context ICIContext, options foundation.IDictionary) Detector {
	rv := objc.Send[Detector](objc.ID(getDetectorClass().class), objc.Sel("detectorOfType:context:options:"), type_, context, options)
	return rv
}







// Creates and returns a configured detector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDetector/init(ofType:context:options:)
func (dc _DetectorClass) DetectorOfTypeContextOptions(type_ objc.IObject /* cross-framework: NSString */, context ICIContext, options foundation.IDictionary) IDetector {
	rv := objc.Send[Detector](objc.ID(dc.class), objc.Sel("detectorOfType:context:options:"), type_, context, options)
	return rv
}












// Searches for features in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDetector/features(in:)
func (d_ Detector) FeaturesInImage(image ICIImage) []Feature {
	rv := objc.Send[[]Feature](d_.ID, objc.Sel("featuresInImage:"), image)
	return rv
}


// Searches for features in an image based on the specified image orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDetector/features(in:options:)
func (d_ Detector) FeaturesInImageOptions(image ICIImage, options foundation.IDictionary) []Feature {
	rv := objc.Send[[]Feature](d_.ID, objc.Sel("featuresInImage:options:"), image, options)
	return rv
}












