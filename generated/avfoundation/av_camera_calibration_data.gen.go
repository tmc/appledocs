// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CameraCalibrationData] class.
var (
	CameraCalibrationDataClass     _CameraCalibrationDataClass
	CameraCalibrationDataClassOnce sync.Once
)

func getCameraCalibrationDataClass() _CameraCalibrationDataClass {
	CameraCalibrationDataClassOnce.Do(func() {
		CameraCalibrationDataClass = _CameraCalibrationDataClass{objc.GetClass("AVCameraCalibrationData")}
	})
	return CameraCalibrationDataClass
}

type _CameraCalibrationDataClass struct {
	class objc.Class
}





// An interface definition for the [CameraCalibrationData] class.
type ICameraCalibrationData interface {
	objectivec.IObject
	

	// properties:
	ExtrinsicMatrix() objectivec.IObject
	IntrinsicMatrix() objectivec.IObject
	IntrinsicMatrixReferenceDimensions() corefoundation.CGSize
	InverseLensDistortionLookupTable() objc.IObject /* cross-framework: NSData */
	LensDistortionCenter() corefoundation.CGPoint
	LensDistortionLookupTable() objc.IObject /* cross-framework: NSData */
	PixelSize() float32


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CameraCalibrationDataClass) Alloc() CameraCalibrationData {
	rv := objc.Send[CameraCalibrationData](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CameraCalibrationDataClass) New() CameraCalibrationData {
	rv := objc.Send[CameraCalibrationData](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CameraCalibrationData) Init() CameraCalibrationData {
	rv := objc.Send[CameraCalibrationData](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CameraCalibrationData) Autorelease() CameraCalibrationData {
	rv := objc.Send[CameraCalibrationData](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCameraCalibrationData creates a new CameraCalibrationData instance.
func NewCameraCalibrationData() CameraCalibrationData {
	return getCameraCalibrationDataClass().New()
}





// Information about the camera characteristics used to capture images and depth data.
//
// Information about the calibration of a camera—such as its pixel focal length, principal point, and lens distortion characteristics—helps to determine the geometric relationships between the camera device and the images it captures. You can use this information to accurately render visual effects into images produced by a camera or perform computer vision tasks such as correcting images for geometric distortions.


// Information about the camera characteristics used to capture images and depth data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCameraCalibrationData
type CameraCalibrationData struct {
	objectivec.Object
}

// CameraCalibrationDataFrom constructs a [CameraCalibrationData] from an unsafe.Pointer.
//
// Information about the camera characteristics used to capture images and depth data.
func CameraCalibrationDataFrom(ptr unsafe.Pointer) CameraCalibrationData {
	return CameraCalibrationData{objectivec.Object{objc.ID(ptr)}}
}

























// A matrix relating a camera’s position and orientation to a world or scene coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCameraCalibrationData/extrinsicMatrix
func (c_ CameraCalibrationData) ExtrinsicMatrix() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("extrinsicMatrix"))
	return rv
}


// A matrix that relates a camera’s internal properties to an ideal pinhole-camera model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCameraCalibrationData/intrinsicMatrix
func (c_ CameraCalibrationData) IntrinsicMatrix() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("intrinsicMatrix"))
	return rv
}


// The image dimensions to which the camera’s intrinsic matrix values are relative.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCameraCalibrationData/intrinsicMatrixReferenceDimensions
func (c_ CameraCalibrationData) IntrinsicMatrixReferenceDimensions() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c_.ID, objc.Sel("intrinsicMatrixReferenceDimensions"))
	return rv
}


// A map of floating-point values describing radial distortions for use in reapplying camera geometry to a rectified image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCameraCalibrationData/inverseLensDistortionLookupTable
func (c_ CameraCalibrationData) InverseLensDistortionLookupTable() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("inverseLensDistortionLookupTable"))
	return rv
}


// The offset of the distortion center of the camera lens from the top-left corner of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCameraCalibrationData/lensDistortionCenter
func (c_ CameraCalibrationData) LensDistortionCenter() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](c_.ID, objc.Sel("lensDistortionCenter"))
	return rv
}


// A map of floating-point values describing radial distortions imparted by the camera lens, for use in rectifying camera images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCameraCalibrationData/lensDistortionLookupTable
func (c_ CameraCalibrationData) LensDistortionLookupTable() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("lensDistortionLookupTable"))
	return rv
}


// The size, in millimeters, of one image pixel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCameraCalibrationData/pixelSize
func (c_ CameraCalibrationData) PixelSize() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("pixelSize"))
	return rv
}








