// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
}

// A container that provides information about and access to the image, video, or Live Photo content of an asset to be edited.
//
// To edit an asset’s photo or video content: Fetch a object that represents the photo or video to be edited. Call the asset’s method to retrieve a object. Apply your edits to the asset. To allow a user to continue working with the edit later (for example, to adjust the parameters of a photo filter), create a object describing the changes. Initialize a object. For photo- or video-only assets, use the editing output’s properties to provide edited asset data. For Live Photo assets, create a object to edit the Live Photo content. Use a photo library change block to commit the edit. In the block, create a object and set its property to the editing output that you created. For more details, see . You can also edit assets from photo editing extensions. In this case, instead of working with a object, you implement methods in the protocol. Photos provides a object when your extension begins editing. When editing is complete, Photos requests a object that contains the edited asset content.
//
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
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetchangerequest/contenteditingoutput
func (p_ PHContentEditingInput) ContentEditingOutput() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("contentEditingOutput"))
	return rv
}


// SetContentEditingOutput sets the value of the contentEditingOutput property.
// The output of an asset content editing session.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetchangerequest/contenteditingoutput
func (p_ PHContentEditingInput) SetContentEditingOutput(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentEditingOutput:"), value)
}

// An object that describes the most recent edit to the asset’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHContentEditingInput/adjustmentData
func (p_ PHContentEditingInput) AdjustmentData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("adjustmentData"))
	return rv
}

// The video asset, as an object.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHContentEditingInput/audiovisualAsset
func (p_ PHContentEditingInput) AudiovisualAsset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("audiovisualAsset"))
	return rv
}

// The video asset, as an object.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHContentEditingInput/avAsset
func (p_ PHContentEditingInput) AvAsset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("avAsset"))
	return rv
}

// The type of data provided as the asset’s content editing input image or video.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHContentEditingInput/contentType
func (p_ PHContentEditingInput) ContentType() UTType {
	rv := objc.Send[UTType](p_.ID, objc.Sel("contentType"))
	return rv
}

// The date and time when the asset was originally created.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHContentEditingInput/creationDate
func (p_ PHContentEditingInput) CreationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("creationDate"))
	return rv
}

// An image of the asset’s contents, appropriately sized for display.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHContentEditingInput/displaySizeImage
func (p_ PHContentEditingInput) DisplaySizeImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("displaySizeImage"))
	return rv
}

// The Exif display orientation of the full-size image file.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHContentEditingInput/fullSizeImageOrientation
func (p_ PHContentEditingInput) FullSizeImageOrientation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("fullSizeImageOrientation"))
	return rv
}

// The URL to a file that contains the full-sized image data.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHContentEditingInput/fullSizeImageURL
func (p_ PHContentEditingInput) FullSizeImageURL() foundation.URL {
	rv := objc.Send[foundation.URL](p_.ID, objc.Sel("fullSizeImageURL"))
	return rv
}

// The unedited Live Photo content of the editing input.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHContentEditingInput/livePhoto
func (p_ PHContentEditingInput) LivePhoto() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("livePhoto"))
	return rv
}

// The location information that was saved with the asset.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHContentEditingInput/location
func (p_ PHContentEditingInput) Location() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("location"))
	return rv
}

// The subtypes of the asset, identifying special kinds of assets such as a panoramic photo or a high-frame-rate video.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHContentEditingInput/mediaSubtypes
func (p_ PHContentEditingInput) MediaSubtypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("mediaSubtypes"))
	return rv
}

// The type of the asset, such as video or audio.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHContentEditingInput/mediaType
func (p_ PHContentEditingInput) MediaType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("mediaType"))
	return rv
}

// The style in which to present this content to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHContentEditingInput/playbackStyle
func (p_ PHContentEditingInput) PlaybackStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("playbackStyle"))
	return rv
}

// The uniform type identifier for the asset’s image or video data.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHContentEditingInput/uniformTypeIdentifier
func (p_ PHContentEditingInput) UniformTypeIdentifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("uniformTypeIdentifier"))
	return rv
}



