// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
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
}

// Information about the camera characteristics used to capture images and depth data.
//
// Information about the calibration of a camera—such as its pixel focal length, principal point, and lens distortion characteristics—helps to determine the geometric relationships between the camera device and the images it captures. You can use this information to accurately render visual effects into images produced by a camera or perform computer vision tasks such as correcting images for geometric distortions.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CameraCalibrationDataClass) Alloc() CameraCalibrationData {
	rv := objc.Send[CameraCalibrationData](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A matrix relating a camera’s position and orientation to a world or scene coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcameracalibrationdata/extrinsicmatrix
func (c_ CameraCalibrationData) ExtrinsicMatrix() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("extrinsicMatrix"))
	return rv
}


// SetExtrinsicMatrix sets the value of the extrinsicMatrix property.
// A matrix relating a camera’s position and orientation to a world or scene coordinate system.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcameracalibrationdata/extrinsicmatrix
func (c_ CameraCalibrationData) SetExtrinsicMatrix(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExtrinsicMatrix:"), value)
}

// A matrix that relates a camera’s internal properties to an ideal pinhole-camera model.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcameracalibrationdata/intrinsicmatrix
func (c_ CameraCalibrationData) IntrinsicMatrix() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("intrinsicMatrix"))
	return rv
}


// SetIntrinsicMatrix sets the value of the intrinsicMatrix property.
// A matrix that relates a camera’s internal properties to an ideal pinhole-camera model.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcameracalibrationdata/intrinsicmatrix
func (c_ CameraCalibrationData) SetIntrinsicMatrix(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIntrinsicMatrix:"), value)
}

// The image dimensions to which the camera’s intrinsic matrix values are relative.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcameracalibrationdata/intrinsicmatrixreferencedimensions
func (c_ CameraCalibrationData) IntrinsicMatrixReferenceDimensions() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](c_.ID, objc.Sel("intrinsicMatrixReferenceDimensions"))
	return rv
}


// SetIntrinsicMatrixReferenceDimensions sets the value of the intrinsicMatrixReferenceDimensions property.
// The image dimensions to which the camera’s intrinsic matrix values are relative.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcameracalibrationdata/intrinsicmatrixreferencedimensions
func (c_ CameraCalibrationData) SetIntrinsicMatrixReferenceDimensions(value coregraphics.CGSize) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIntrinsicMatrixReferenceDimensions:"), value)
}

// A map of floating-point values describing radial distortions for use in reapplying camera geometry to a rectified image.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcameracalibrationdata/inverselensdistortionlookuptable
func (c_ CameraCalibrationData) InverseLensDistortionLookupTable() foundation.Data {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("inverseLensDistortionLookupTable"))
	return rv
}


// SetInverseLensDistortionLookupTable sets the value of the inverseLensDistortionLookupTable property.
// A map of floating-point values describing radial distortions for use in reapplying camera geometry to a rectified image.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcameracalibrationdata/inverselensdistortionlookuptable
func (c_ CameraCalibrationData) SetInverseLensDistortionLookupTable(value foundation.IData) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInverseLensDistortionLookupTable:"), value)
}

// The offset of the distortion center of the camera lens from the top-left corner of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcameracalibrationdata/lensdistortioncenter
func (c_ CameraCalibrationData) LensDistortionCenter() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](c_.ID, objc.Sel("lensDistortionCenter"))
	return rv
}


// SetLensDistortionCenter sets the value of the lensDistortionCenter property.
// The offset of the distortion center of the camera lens from the top-left corner of the image.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcameracalibrationdata/lensdistortioncenter
func (c_ CameraCalibrationData) SetLensDistortionCenter(value coregraphics.CGPoint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLensDistortionCenter:"), value)
}

// A map of floating-point values describing radial distortions imparted by the camera lens, for use in rectifying camera images.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcameracalibrationdata/lensdistortionlookuptable
func (c_ CameraCalibrationData) LensDistortionLookupTable() foundation.Data {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("lensDistortionLookupTable"))
	return rv
}


// SetLensDistortionLookupTable sets the value of the lensDistortionLookupTable property.
// A map of floating-point values describing radial distortions imparted by the camera lens, for use in rectifying camera images.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcameracalibrationdata/lensdistortionlookuptable
func (c_ CameraCalibrationData) SetLensDistortionLookupTable(value foundation.IData) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLensDistortionLookupTable:"), value)
}

// The size, in millimeters, of one image pixel.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcameracalibrationdata/pixelsize
func (c_ CameraCalibrationData) PixelSize() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("pixelSize"))
	return rv
}


// SetPixelSize sets the value of the pixelSize property.
// The size, in millimeters, of one image pixel.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcameracalibrationdata/pixelsize
func (c_ CameraCalibrationData) SetPixelSize(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPixelSize:"), value)
}



