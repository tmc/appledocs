// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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




