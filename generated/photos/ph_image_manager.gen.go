// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHImageManager] class.
var (
	PHImageManagerClass     _PHImageManagerClass
	PHImageManagerClassOnce sync.Once
)

func getPHImageManagerClass() _PHImageManagerClass {
	PHImageManagerClassOnce.Do(func() {
		PHImageManagerClass = _PHImageManagerClass{objc.GetClass("PHImageManager")}
	})
	return PHImageManagerClass
}

type _PHImageManagerClass struct {
	class objc.Class
}

// An interface definition for the [PHImageManager] class.
type IPHImageManager interface {
	objectivec.IObject
	// properties:
	PHImageManagerMaximumSize() objc.IObject   /* cross-framework: Size */
	PHInvalidImageRequestID() PHImageRequestID /* typedef */
	// methods:
	CancelImageRequest(requestID PHImageRequestID /* typedef */)
	RequestAVAssetForVideoOptionsResultHandler(asset IPHAsset, options IPHVideoRequestOptions, resultHandler unsafe.Pointer) PHImageRequestID                                                                                                                 /* typedef */
	RequestExportSessionForVideoOptionsExportPresetResultHandler(asset IPHAsset, options IPHVideoRequestOptions, exportPreset objc.IObject /* cross-framework: NSString */, resultHandler unsafe.Pointer) PHImageRequestID                                    /* typedef */
	RequestImageForAssetTargetSizeContentModeOptionsResultHandler(asset IPHAsset, targetSize objc.IObject /* cross-framework: Size */, contentMode PHImageContentMode, options IPHImageRequestOptions, resultHandler unsafe.Pointer) PHImageRequestID         /* typedef */
	RequestImageDataAndOrientationForAssetOptionsResultHandler(asset IPHAsset, options IPHImageRequestOptions, resultHandler unsafe.Pointer) PHImageRequestID                                                                                                 /* typedef */
	RequestLivePhotoForAssetTargetSizeContentModeOptionsResultHandler(asset IPHAsset, targetSize objc.IObject /* cross-framework: Size */, contentMode PHImageContentMode, options IPHLivePhotoRequestOptions, resultHandler unsafe.Pointer) PHImageRequestID /* typedef */
	RequestPlayerItemForVideoOptionsResultHandler(asset IPHAsset, options IPHVideoRequestOptions, resultHandler unsafe.Pointer) PHImageRequestID                                                                                                              /* typedef */
}

// An object that facilitates retrieving or generating preview thumbnails and asset data.
//
// Use these methods to fetch full-size photo assets or thumbnail images, or to retrieve AVFoundation objects for playing, exporting, and manipulating video assets. To load image or video data: Use the class to fetch the asset you’re interested in. Call the method to retrieve the shared image manager object. Use one of the methods listed in the Requesting groups below to load the asset’s image or video data. The image manager caches the asset images and data it provides, so later requests for the same assets with similar parameters will return results more quickly. If you need to load image data for many assets together, use the class to “preheat” the cache by loading images you expect to need soon. For example, when populating a collection view with photo asset thumbnails, you can cache images ahead of the current scroll position.

// An object that facilitates retrieving or generating preview thumbnails and asset data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageManager
type PHImageManager struct {
	objectivec.Object
}

// PHImageManagerFrom constructs a [PHImageManager] from an unsafe.Pointer.
//
// An object that facilitates retrieving or generating preview thumbnails and asset data.
func PHImageManagerFrom(ptr unsafe.Pointer) PHImageManager {
	return PHImageManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHImageManagerClass) Alloc() PHImageManager {
	rv := objc.Send[PHImageManager](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHImageManagerClass) New() PHImageManager {
	rv := objc.Send[PHImageManager](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHImageManager) Init() PHImageManager {
	rv := objc.Send[PHImageManager](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHImageManager) Autorelease() PHImageManager {
	rv := objc.Send[PHImageManager](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHImageManager creates a new PHImageManager instance.
func NewPHImageManager() PHImageManager {
	return getPHImageManagerClass().New()
}

// Returns the shared image manager object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageManager/default()
func (pc _PHImageManagerClass) DefaultManager() PHImageManager {
	rv := objc.Send[PHImageManager](objc.ID(pc.class), objc.Sel("defaultManager"))
	return rv
}

// Cancels an asynchronous request
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageManager/cancelImageRequest(_:)
func (p_ PHImageManager) CancelImageRequest(requestID PHImageRequestID /* typedef */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("cancelImageRequest:"), requestID)
}

// Requests AVFoundation objects representing the video asset’s content and state, to be loaded asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageManager/requestAVAsset(forVideo:options:resultHandler:)
func (p_ PHImageManager) RequestAVAssetForVideoOptionsResultHandler(asset IPHAsset, options IPHVideoRequestOptions, resultHandler unsafe.Pointer) PHImageRequestID /* typedef */ {
	rv := objc.Send[PHImageRequestID](p_.ID, objc.Sel("requestAVAssetForVideo:options:resultHandler:"), asset, options, resultHandler)
	return rv
}

// Requests an export session for writing the video asset’s data to a file, to be loaded asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageManager/requestExportSession(forVideo:options:exportPreset:resultHandler:)
func (p_ PHImageManager) RequestExportSessionForVideoOptionsExportPresetResultHandler(asset IPHAsset, options IPHVideoRequestOptions, exportPreset objc.IObject /* cross-framework: NSString */, resultHandler unsafe.Pointer) PHImageRequestID /* typedef */ {
	rv := objc.Send[PHImageRequestID](p_.ID, objc.Sel("requestExportSessionForVideo:options:exportPreset:resultHandler:"), asset, options, exportPreset, resultHandler)
	return rv
}

// Requests an image representation for the specified asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageManager/requestImage(for:targetSize:contentMode:options:resultHandler:)
func (p_ PHImageManager) RequestImageForAssetTargetSizeContentModeOptionsResultHandler(asset IPHAsset, targetSize objc.IObject /* cross-framework: Size */, contentMode PHImageContentMode, options IPHImageRequestOptions, resultHandler unsafe.Pointer) PHImageRequestID /* typedef */ {
	rv := objc.Send[PHImageRequestID](p_.ID, objc.Sel("requestImageForAsset:targetSize:contentMode:options:resultHandler:"), asset, targetSize, contentMode, options, resultHandler)
	return rv
}

// Requests the largest represented image as data bytes and EXIF orientation for the specified asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageManager/requestImageDataAndOrientation(for:options:resultHandler:)
func (p_ PHImageManager) RequestImageDataAndOrientationForAssetOptionsResultHandler(asset IPHAsset, options IPHImageRequestOptions, resultHandler unsafe.Pointer) PHImageRequestID /* typedef */ {
	rv := objc.Send[PHImageRequestID](p_.ID, objc.Sel("requestImageDataAndOrientationForAsset:options:resultHandler:"), asset, options, resultHandler)
	return rv
}

// Requests a Live Photo representation for the specified asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageManager/requestLivePhoto(for:targetSize:contentMode:options:resultHandler:)
func (p_ PHImageManager) RequestLivePhotoForAssetTargetSizeContentModeOptionsResultHandler(asset IPHAsset, targetSize objc.IObject /* cross-framework: Size */, contentMode PHImageContentMode, options IPHLivePhotoRequestOptions, resultHandler unsafe.Pointer) PHImageRequestID /* typedef */ {
	rv := objc.Send[PHImageRequestID](p_.ID, objc.Sel("requestLivePhotoForAsset:targetSize:contentMode:options:resultHandler:"), asset, targetSize, contentMode, options, resultHandler)
	return rv
}

// Requests a representation of the video asset for playback, to be loaded asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHImageManager/requestPlayerItem(forVideo:options:resultHandler:)
func (p_ PHImageManager) RequestPlayerItemForVideoOptionsResultHandler(asset IPHAsset, options IPHVideoRequestOptions, resultHandler unsafe.Pointer) PHImageRequestID /* typedef */ {
	rv := objc.Send[PHImageRequestID](p_.ID, objc.Sel("requestPlayerItemForVideo:options:resultHandler:"), asset, options, resultHandler)
	return rv
}

// A special value for requesting original image data or the largest rendered image available. .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phimagemanagermaximumsize
func (p_ PHImageManager) PHImageManagerMaximumSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](p_.ID, objc.Sel("PHImageManagerMaximumSize"))
	return rv
}

// A special value provided for asynchronous image requests that cannot be canceled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phinvalidimagerequestid
func (p_ PHImageManager) PHInvalidImageRequestID() PHImageRequestID /* typedef */ {
	rv := objc.Send[PHImageRequestID](p_.ID, objc.Sel("PHInvalidImageRequestID"))
	return rv
}
