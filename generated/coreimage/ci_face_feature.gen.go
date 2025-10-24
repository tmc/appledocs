// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
)

// The class instance for the [FaceFeature] class.
var (
	FaceFeatureClass     _FaceFeatureClass
	FaceFeatureClassOnce sync.Once
)

func getFaceFeatureClass() _FaceFeatureClass {
	FaceFeatureClassOnce.Do(func() {
		FaceFeatureClass = _FaceFeatureClass{objc.GetClass("CIFaceFeature")}
	})
	return FaceFeatureClass
}

type _FaceFeatureClass struct {
	class objc.Class
}

// An interface definition for the [FaceFeature] class.
type IFaceFeature interface {
	IFeature
	// properties:
	Bounds() objc.IObject /* cross-framework: Rect */
	FaceAngle() float32
	HasFaceAngle() bool
	HasLeftEyePosition() bool
	HasMouthPosition() bool
	HasRightEyePosition() bool
	HasSmile() bool
	HasTrackingFrameCount() bool
	HasTrackingID() bool
	LeftEyeClosed() bool
	LeftEyePosition() objc.IObject /* cross-framework: Point */
	MouthPosition() objc.IObject /* cross-framework: Point */
	RightEyeClosed() bool
	RightEyePosition() objc.IObject /* cross-framework: Point */
	TrackingFrameCount() int
	TrackingID() int
	// methods:
}

// Information about a face detected in a still or video image.
//
// The properties of a object provide information about the face’s eyes and mouth. A face object in a video can also have properties that track its location over time, tracking ID and frame count.


// Information about a face detected in a still or video image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature
type FaceFeature struct {
	Feature
}

// FaceFeatureFrom constructs a [FaceFeature] from an unsafe.Pointer.
//
// Information about a face detected in a still or video image.
func FaceFeatureFrom(ptr unsafe.Pointer) FaceFeature {
	return FaceFeature{
		Feature: FeatureFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (fc _FaceFeatureClass) Alloc() FaceFeature {
	rv := objc.Send[FaceFeature](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FaceFeatureClass) New() FaceFeature {
	rv := objc.Send[FaceFeature](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FaceFeature) Init() FaceFeature {
	rv := objc.Send[FaceFeature](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FaceFeature) Autorelease() FaceFeature {
	rv := objc.Send[FaceFeature](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFaceFeature creates a new FaceFeature instance.
func NewFaceFeature() FaceFeature {
	return getFaceFeatureClass().New()
}



// A rectangle indicating the position and extent of the face feature in image coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/bounds-swift.property
func (f_ FaceFeature) Bounds() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](f_.ID, objc.Sel("bounds"))
	return rv
}


// The rotation of the face.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/faceAngle-swift.property
func (f_ FaceFeature) FaceAngle() float32 {
	rv := objc.Send[float32](f_.ID, objc.Sel("faceAngle"))
	return rv
}


// A Boolean value that indicates whether information about face rotation is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasFaceAngle-swift.property
func (f_ FaceFeature) HasFaceAngle() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("hasFaceAngle"))
	return rv
}


// A Boolean value that indicates whether the detector found the face’s left eye.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasLeftEyePosition-swift.property
func (f_ FaceFeature) HasLeftEyePosition() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("hasLeftEyePosition"))
	return rv
}


// A Boolean value that indicates whether the detector found the face’s mouth.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasMouthPosition-swift.property
func (f_ FaceFeature) HasMouthPosition() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("hasMouthPosition"))
	return rv
}


// A Boolean value that indicates whether the detector found the face’s right eye.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasRightEyePosition-swift.property
func (f_ FaceFeature) HasRightEyePosition() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("hasRightEyePosition"))
	return rv
}


// A Boolean value that indicates whether a smile is detected in the face.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasSmile-swift.property
func (f_ FaceFeature) HasSmile() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("hasSmile"))
	return rv
}


// A Boolean value that indicates the face object has a tracking frame count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasTrackingFrameCount-swift.property
func (f_ FaceFeature) HasTrackingFrameCount() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("hasTrackingFrameCount"))
	return rv
}


// A Boolean value that indicates whether the face object has a tracking ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasTrackingID-swift.property
func (f_ FaceFeature) HasTrackingID() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("hasTrackingID"))
	return rv
}


// A Boolean value that indicates whether a closed left eye is detected in the face.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/leftEyeClosed-swift.property
func (f_ FaceFeature) LeftEyeClosed() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("leftEyeClosed"))
	return rv
}


// The image coordinate of the center of the left eye.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/leftEyePosition-swift.property
func (f_ FaceFeature) LeftEyePosition() objc.IObject /* cross-framework: Point */ {
	rv := objc.Send[corefoundation.Point](f_.ID, objc.Sel("leftEyePosition"))
	return rv
}


// The image coordinate of the center of the mouth.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/mouthPosition-swift.property
func (f_ FaceFeature) MouthPosition() objc.IObject /* cross-framework: Point */ {
	rv := objc.Send[corefoundation.Point](f_.ID, objc.Sel("mouthPosition"))
	return rv
}


// A Boolean value that indicates whether a closed right eye is detected in the face.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/rightEyeClosed-swift.property
func (f_ FaceFeature) RightEyeClosed() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("rightEyeClosed"))
	return rv
}


// The image coordinate of the center of the right eye.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/rightEyePosition-swift.property
func (f_ FaceFeature) RightEyePosition() objc.IObject /* cross-framework: Point */ {
	rv := objc.Send[corefoundation.Point](f_.ID, objc.Sel("rightEyePosition"))
	return rv
}


// The tracking frame count of the face.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/trackingFrameCount-swift.property
func (f_ FaceFeature) TrackingFrameCount() int {
	rv := objc.Send[int](f_.ID, objc.Sel("trackingFrameCount"))
	return rv
}


// The tracking identifier of the face object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/trackingID-swift.property
func (f_ FaceFeature) TrackingID() int {
	rv := objc.Send[int](f_.ID, objc.Sel("trackingID"))
	return rv
}



