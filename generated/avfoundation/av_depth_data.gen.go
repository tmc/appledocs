// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVDepthData */


/* debug [class_header]: Header for AVDepthData */
// The class instance for the [DepthData] class.
var (
	DepthDataClass     _DepthDataClass
	DepthDataClassOnce sync.Once
)

func getDepthDataClass() _DepthDataClass {
	DepthDataClassOnce.Do(func() {
		DepthDataClass = _DepthDataClass{objc.GetClass("AVDepthData")}
	})
	return DepthDataClass
}

type _DepthDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DepthData */
// An interface definition for the [DepthData] class.
type IDepthData interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DepthData */
	// properties:
	AvailableDepthDataTypes() []foundation.Number
	CameraCalibrationData() IAVCameraCalibrationData
	DepthDataAccuracy() DepthDataAccuracy
	DepthDataMap() PixelBufferRef /* not a class type */
	DepthDataQuality() DepthDataQuality
	DepthDataType() uint32 /* not a class type */
	DepthDataFiltered() bool
	IsDepthDataFiltered() bool
	SetIsDepthDataFiltered(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DepthData */
	// methods:
	DepthDataByApplyingExifOrientation(exifOrientation ImagePropertyOrientation /* not a class type */) objectivec.IObject
	DepthDataByConvertingToDepthDataType(depthDataType uint32 /* not a class type */) objectivec.IObject
	DictionaryRepresentationForAuxiliaryDataType(outAuxDataType objc.IObject /* cross-framework: NSString */) foundation.Dictionary
	DepthDataByReplacingDepthDataMapWithPixelBufferError(pixelBuffer PixelBufferRef /* not a class type */, outError objectivec.IObject) objectivec.IObject
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DepthData */
// Alloc allocates a new instance without initialization.
func (dc _DepthDataClass) Alloc() DepthData {
	rv := objc.Send[DepthData](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DepthDataClass) New() DepthData {
	rv := objc.Send[DepthData](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DepthData) Init() DepthData {
	rv := objc.Send[DepthData](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DepthData) Autorelease() DepthData {
	rv := objc.Send[DepthData](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDepthData creates a new DepthData instance.
func NewDepthData() DepthData {
	return getDepthDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DepthData */
// A container for per-pixel distance or disparity information captured by compatible camera devices.
//
// is a generic term for a map of per-pixel data containing depth-related information. A depth data object wraps a disparity or depth map and provides conversion methods, focus information, and camera calibration data to aid in using the map for rendering or computer vision tasks. A depth map describes at each pixel the distance to an object, in meters. A disparity map describes normalized shift values for use in comparing two images. The value for each pixel in the map is in units of 1/meters: ( ). The capture pipeline generates disparity or depth maps from camera images containing nonrectilinear data. Camera lenses have small imperfections that cause small distortions in their resultant images compared to an ideal pinhole camera model, so maps contain nonrectilinear (nondistortion-corrected) data as well. The maps’ values are warped to match the lens distortion characteristics present in the YUV image pixel buffers captured at the same time. Because a depth data map is nonrectilinear, you can use an map as a proxy for depth when rendering effects to its accompanying image, but not to correlate points in 3D space. To use depth data for computer vision tasks, use the data in the property to rectify the depth data. There are two ways to capture depth data: The class captures and delivers depth data in a stream (similar to how the delivers video data). The class delivers depth data as a property of an object containing the captured image. You can also create objects using information obtained from image files with the framework. When editing images containing depth information, use the methods listed in Transforming and Processing to generate derivative objects reflecting the edits that have been performed.


// A container for per-pixel distance or disparity information captured by compatible camera devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDepthData
type DepthData struct {
	objectivec.Object
}

// DepthDataFrom constructs a [DepthData] from an unsafe.Pointer.
//
// A container for per-pixel distance or disparity information captured by compatible camera devices.
func DepthDataFrom(ptr unsafe.Pointer) DepthData {
	return DepthData{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DepthData */

// Creates a depth data object from depth information such as that found in an image file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDepthData/init(fromDictionaryRepresentation:)
func NewDepthDataFromDictionaryRepresentationError(imageSourceAuxDataInfoDictionary objc.IObject /* cross-framework: NSDictionary */, outError objectivec.IObject) DepthData {
	rv := objc.Send[DepthData](objc.ID(getDepthDataClass().class), objc.Sel("depthDataFromDictionaryRepresentation:error:"), imageSourceAuxDataInfoDictionary, outError)
	return rv
}/* debug [class_init_methods/constructor]: NewDepthDataFromDictionaryRepresentationError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DepthData */

// Creates a depth data object from depth information such as that found in an image file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDepthData/init(fromDictionaryRepresentation:)
func (dc _DepthDataClass) DepthDataFromDictionaryRepresentationError(imageSourceAuxDataInfoDictionary objc.IObject /* cross-framework: NSDictionary */, outError objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(dc.class), objc.Sel("depthDataFromDictionaryRepresentation:error:"), imageSourceAuxDataInfoDictionary, outError)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DepthDataFromDictionaryRepresentationError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DepthData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DepthData */

// Returns a derivative depth data object by mirroring or rotating it to the specified orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDepthData/applyingExifOrientation(_:)
func (d_ DepthData) DepthDataByApplyingExifOrientation(exifOrientation ImagePropertyOrientation /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("depthDataByApplyingExifOrientation:"), exifOrientation)
	return rv
}/* debug [instance_methods/method]: DepthDataByApplyingExifOrientation */


// Returns a derivative depth data object by converting the depth data map to the specified data type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDepthData/converting(toDepthDataType:)
func (d_ DepthData) DepthDataByConvertingToDepthDataType(depthDataType uint32 /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("depthDataByConvertingToDepthDataType:"), depthDataType)
	return rv
}/* debug [instance_methods/method]: DepthDataByConvertingToDepthDataType */


// Returns a dictionary representation of the depth data suitable for writing into an image file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDepthData/dictionaryRepresentation(forAuxiliaryDataType:)
func (d_ DepthData) DictionaryRepresentationForAuxiliaryDataType(outAuxDataType objc.IObject /* cross-framework: NSString */) foundation.Dictionary {
	rv := objc.Send[foundation.Dictionary](d_.ID, objc.Sel("dictionaryRepresentationForAuxiliaryDataType:"), outAuxDataType)
	return rv
}/* debug [instance_methods/method]: DictionaryRepresentationForAuxiliaryDataType */


// Returns a derivative depth data object by replacing the depth data map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDepthData/replacingDepthDataMap(with:)
func (d_ DepthData) DepthDataByReplacingDepthDataMapWithPixelBufferError(pixelBuffer PixelBufferRef /* not a class type */, outError objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("depthDataByReplacingDepthDataMapWithPixelBuffer:error:"), pixelBuffer, outError)
	return rv
}/* debug [instance_methods/method]: DepthDataByReplacingDepthDataMapWithPixelBufferError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DepthData */

// The list of depth data formats to which you can convert this depth data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDepthData/availableDepthDataTypes-472g0
func (d_ DepthData) AvailableDepthDataTypes() []foundation.Number {
	rv := objc.Send[[]foundation.Number](d_.ID, objc.Sel("availableDepthDataTypes"))
	return rv
}/* debug [instance_properties/getter]: availableDepthDataTypes */


// The imaging parameters with which this depth data was captured.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDepthData/cameraCalibrationData
func (d_ DepthData) CameraCalibrationData() IAVCameraCalibrationData {
	rv := objc.Send[CameraCalibrationData](d_.ID, objc.Sel("cameraCalibrationData"))
	return rv
}/* debug [instance_properties/getter]: cameraCalibrationData */


// The general accuracy of depth data map values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDepthData/depthDataAccuracy
func (d_ DepthData) DepthDataAccuracy() DepthDataAccuracy {
	rv := objc.Send[DepthDataAccuracy](d_.ID, objc.Sel("depthDataAccuracy"))
	return rv
}/* debug [instance_properties/getter]: depthDataAccuracy */


// A pixel buffer containing the depth data’s per-pixel depth or disparity data map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDepthData/depthDataMap
func (d_ DepthData) DepthDataMap() PixelBufferRef /* not a class type */ {
	rv := objc.Send[PixelBufferRef](d_.ID, objc.Sel("depthDataMap"))
	return rv
}/* debug [instance_properties/getter]: depthDataMap */


// The overall quality of the depth map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDepthData/depthDataQuality
func (d_ DepthData) DepthDataQuality() DepthDataQuality {
	rv := objc.Send[DepthDataQuality](d_.ID, objc.Sel("depthDataQuality"))
	return rv
}/* debug [instance_properties/getter]: depthDataQuality */


// The pixel format of the depth data map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDepthData/depthDataType
func (d_ DepthData) DepthDataType() uint32 /* not a class type */ {
	rv := objc.Send[uint32](d_.ID, objc.Sel("depthDataType"))
	return rv
}/* debug [instance_properties/getter]: depthDataType */


// A Boolean value indicating whether the depth map contains temporally smoothed data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDepthData/isDepthDataFiltered
func (d_ DepthData) DepthDataFiltered() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("depthDataFiltered"))
	return rv
}/* debug [instance_properties/getter]: depthDataFiltered */


// A Boolean value indicating whether the depth map contains temporally smoothed data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdepthdata/isdepthdatafiltered
func (d_ DepthData) IsDepthDataFiltered() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isDepthDataFiltered"))
	return rv
}/* debug [instance_properties/getter]: isDepthDataFiltered */


// A Boolean value indicating whether the depth map contains temporally smoothed data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avdepthdata/isdepthdatafiltered
func (d_ DepthData) SetIsDepthDataFiltered(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsDepthDataFiltered:"), value)
}/* debug [instance_properties/setter]: isDepthDataFiltered */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVDepthData */


