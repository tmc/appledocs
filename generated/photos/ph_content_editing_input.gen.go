// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
)

// The class instance for the [PHContentEditingInput] class.
var (
	PHContentEditingInputClass     _PHContentEditingInputClass
	PHContentEditingInputClassOnce sync.Once
)

func getPHContentEditingInputClass() _PHContentEditingInputClass {
	PHContentEditingInputClassOnce.Do(func() {
		PHContentEditingInputClass = _PHContentEditingInputClass{objc.GetClass("PHContentEditingInput")}
	})
	return PHContentEditingInputClass
}

type _PHContentEditingInputClass struct {
	class objc.Class
}

// An interface definition for the [PHContentEditingInput] class.
type IPHContentEditingInput interface {
	objectivec.IObject
	// properties:
	ContentEditingOutput() IPHContentEditingOutput
	SetContentEditingOutput(value IPHContentEditingOutput)
	AdjustmentData() IPHAdjustmentData
	SetAdjustmentData(value IPHAdjustmentData)
	AudiovisualAsset() objc.IObject /* cross-framework: Asset */
	SetAudiovisualAsset(value objc.IObject /* cross-framework: Asset */)
	AvAsset() objc.IObject /* cross-framework: Asset */
	SetAvAsset(value objc.IObject /* cross-framework: Asset */)
	ContentType() objc.IObject /* cross-framework: UTType */
	SetContentType(value objc.IObject /* cross-framework: UTType */)
	CreationDate() objc.IObject /* cross-framework: Date */
	SetCreationDate(value objc.IObject /* cross-framework: Date */)
	DisplaySizeImage() objc.IObject /* cross-framework: Image */
	SetDisplaySizeImage(value objc.IObject /* cross-framework: Image */)
	FullSizeImageOrientation() unsafe.Pointer
	SetFullSizeImageOrientation(value unsafe.Pointer)
	FullSizeImageURL() objc.IObject /* cross-framework: URL */
	SetFullSizeImageURL(value objc.IObject /* cross-framework: URL */)
	LivePhoto() IPHLivePhoto
	SetLivePhoto(value IPHLivePhoto)
	Location() objc.IObject /* cross-framework: Location */
	SetLocation(value objc.IObject /* cross-framework: Location */)
	MediaSubtypes() PHAssetMediaSubtype
	SetMediaSubtypes(value PHAssetMediaSubtype)
	MediaType() PHAssetMediaType
	SetMediaType(value PHAssetMediaType)
	PlaybackStyle() unsafe.Pointer
	SetPlaybackStyle(value unsafe.Pointer)
	UniformTypeIdentifier() objc.IObject /* cross-framework: NSString */
	SetUniformTypeIdentifier(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// A container that provides information about and access to the image, video, or Live Photo content of an asset to be edited.
//
// To edit an asset’s photo or video content: Fetch a object that represents the photo or video to be edited. Call the asset’s method to retrieve a object. Apply your edits to the asset. To allow a user to continue working with the edit later (for example, to adjust the parameters of a photo filter), create a object describing the changes. Initialize a object. For photo- or video-only assets, use the editing output’s properties to provide edited asset data. For Live Photo assets, create a object to edit the Live Photo content. Use a photo library change block to commit the edit. In the block, create a object and set its property to the editing output that you created. For more details, see . You can also edit assets from photo editing extensions. In this case, instead of working with a object, you implement methods in the protocol. Photos provides a object when your extension begins editing. When editing is complete, Photos requests a object that contains the edited asset content.

// A container that provides information about and access to the image, video, or Live Photo content of an asset to be edited.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHContentEditingInput
type PHContentEditingInput struct {
	objectivec.Object
}

// PHContentEditingInputFrom constructs a [PHContentEditingInput] from an unsafe.Pointer.
//
// A container that provides information about and access to the image, video, or Live Photo content of an asset to be edited.
func PHContentEditingInputFrom(ptr unsafe.Pointer) PHContentEditingInput {
	return PHContentEditingInput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHContentEditingInputClass) Alloc() PHContentEditingInput {
	rv := objc.Send[PHContentEditingInput](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHContentEditingInputClass) New() PHContentEditingInput {
	rv := objc.Send[PHContentEditingInput](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHContentEditingInput) Init() PHContentEditingInput {
	rv := objc.Send[PHContentEditingInput](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHContentEditingInput) Autorelease() PHContentEditingInput {
	rv := objc.Send[PHContentEditingInput](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHContentEditingInput creates a new PHContentEditingInput instance.
func NewPHContentEditingInput() PHContentEditingInput {
	return getPHContentEditingInputClass().New()
}

// The output of an asset content editing session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetchangerequest/contenteditingoutput
func (p_ PHContentEditingInput) ContentEditingOutput() IPHContentEditingOutput {
	rv := objc.Send[PHContentEditingOutput](p_.ID, objc.Sel("contentEditingOutput"))
	return rv
}

// The output of an asset content editing session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetchangerequest/contenteditingoutput
func (p_ PHContentEditingInput) SetContentEditingOutput(value IPHContentEditingOutput) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentEditingOutput:"), value)
}

// An object that describes the most recent edit to the asset’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/adjustmentdata
func (p_ PHContentEditingInput) AdjustmentData() IPHAdjustmentData {
	rv := objc.Send[PHAdjustmentData](p_.ID, objc.Sel("adjustmentData"))
	return rv
}

// An object that describes the most recent edit to the asset’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/adjustmentdata
func (p_ PHContentEditingInput) SetAdjustmentData(value IPHAdjustmentData) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAdjustmentData:"), value)
}

// The video asset, as an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/audiovisualasset
func (p_ PHContentEditingInput) AudiovisualAsset() objc.IObject /* cross-framework: Asset */ {
	rv := objc.Send[avfoundation.Asset](p_.ID, objc.Sel("audiovisualAsset"))
	return rv
}

// The video asset, as an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/audiovisualasset
func (p_ PHContentEditingInput) SetAudiovisualAsset(value objc.IObject /* cross-framework: Asset */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAudiovisualAsset:"), value)
}

// The video asset, as an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/avasset
func (p_ PHContentEditingInput) AvAsset() objc.IObject /* cross-framework: Asset */ {
	rv := objc.Send[avfoundation.Asset](p_.ID, objc.Sel("avAsset"))
	return rv
}

// The video asset, as an
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/avasset
func (p_ PHContentEditingInput) SetAvAsset(value objc.IObject /* cross-framework: Asset */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAvAsset:"), value)
}

// The type of data provided as the asset’s content editing input image or video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/contenttype
func (p_ PHContentEditingInput) ContentType() objc.IObject /* cross-framework: UTType */ {
	rv := objc.Send[uniformtypeidentifiers.UTType](p_.ID, objc.Sel("contentType"))
	return rv
}

// The type of data provided as the asset’s content editing input image or video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/contenttype
func (p_ PHContentEditingInput) SetContentType(value objc.IObject /* cross-framework: UTType */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentType:"), value)
}

// The date and time when the asset was originally created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/creationdate
func (p_ PHContentEditingInput) CreationDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("creationDate"))
	return rv
}

// The date and time when the asset was originally created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/creationdate
func (p_ PHContentEditingInput) SetCreationDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCreationDate:"), value)
}

// An image of the asset’s contents, appropriately sized for display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/displaysizeimage
func (p_ PHContentEditingInput) DisplaySizeImage() objc.IObject /* cross-framework: Image */ {
	rv := objc.Send[appkit.Image](p_.ID, objc.Sel("displaySizeImage"))
	return rv
}

// An image of the asset’s contents, appropriately sized for display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/displaysizeimage
func (p_ PHContentEditingInput) SetDisplaySizeImage(value objc.IObject /* cross-framework: Image */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDisplaySizeImage:"), value)
}

// The Exif display orientation of the full-size image file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/fullsizeimageorientation
func (p_ PHContentEditingInput) FullSizeImageOrientation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("fullSizeImageOrientation"))
	return rv
}

// The Exif display orientation of the full-size image file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/fullsizeimageorientation
func (p_ PHContentEditingInput) SetFullSizeImageOrientation(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFullSizeImageOrientation:"), value)
}

// The URL to a file that contains the full-sized image data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/fullsizeimageurl
func (p_ PHContentEditingInput) FullSizeImageURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](p_.ID, objc.Sel("fullSizeImageURL"))
	return rv
}

// The URL to a file that contains the full-sized image data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/fullsizeimageurl
func (p_ PHContentEditingInput) SetFullSizeImageURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFullSizeImageURL:"), value)
}

// The unedited Live Photo content of the editing input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/livephoto
func (p_ PHContentEditingInput) LivePhoto() IPHLivePhoto {
	rv := objc.Send[PHLivePhoto](p_.ID, objc.Sel("livePhoto"))
	return rv
}

// The unedited Live Photo content of the editing input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/livephoto
func (p_ PHContentEditingInput) SetLivePhoto(value IPHLivePhoto) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLivePhoto:"), value)
}

// The location information that was saved with the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/location
func (p_ PHContentEditingInput) Location() objc.IObject /* cross-framework: Location */ {
	rv := objc.Send[corelocation.Location](p_.ID, objc.Sel("location"))
	return rv
}

// The location information that was saved with the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/location
func (p_ PHContentEditingInput) SetLocation(value objc.IObject /* cross-framework: Location */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLocation:"), value)
}

// The subtypes of the asset, identifying special kinds of assets such as a panoramic photo or a high-frame-rate video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/mediasubtypes
func (p_ PHContentEditingInput) MediaSubtypes() PHAssetMediaSubtype {
	rv := objc.Send[PHAssetMediaSubtype](p_.ID, objc.Sel("mediaSubtypes"))
	return rv
}

// The subtypes of the asset, identifying special kinds of assets such as a panoramic photo or a high-frame-rate video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/mediasubtypes
func (p_ PHContentEditingInput) SetMediaSubtypes(value PHAssetMediaSubtype) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMediaSubtypes:"), value)
}

// The type of the asset, such as video or audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/mediatype
func (p_ PHContentEditingInput) MediaType() PHAssetMediaType {
	rv := objc.Send[PHAssetMediaType](p_.ID, objc.Sel("mediaType"))
	return rv
}

// The type of the asset, such as video or audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/mediatype
func (p_ PHContentEditingInput) SetMediaType(value PHAssetMediaType) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMediaType:"), value)
}

// The style in which to present this content to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/playbackstyle
func (p_ PHContentEditingInput) PlaybackStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("playbackStyle"))
	return rv
}

// The style in which to present this content to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/playbackstyle
func (p_ PHContentEditingInput) SetPlaybackStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlaybackStyle:"), value)
}

// The uniform type identifier for the asset’s image or video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/uniformtypeidentifier
func (p_ PHContentEditingInput) UniformTypeIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("uniformTypeIdentifier"))
	return rv
}

// The uniform type identifier for the asset’s image or video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/uniformtypeidentifier
func (p_ PHContentEditingInput) SetUniformTypeIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUniformTypeIdentifier:"), value)
}
