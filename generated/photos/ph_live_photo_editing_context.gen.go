// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coremedia"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHLivePhotoEditingContext] class.
var (
	PHLivePhotoEditingContextClass     _PHLivePhotoEditingContextClass
	PHLivePhotoEditingContextClassOnce sync.Once
)

func getPHLivePhotoEditingContextClass() _PHLivePhotoEditingContextClass {
	PHLivePhotoEditingContextClassOnce.Do(func() {
		PHLivePhotoEditingContextClass = _PHLivePhotoEditingContextClass{objc.GetClass("PHLivePhotoEditingContext")}
	})
	return PHLivePhotoEditingContextClass
}

type _PHLivePhotoEditingContextClass struct {
	class objc.Class
}

// An interface definition for the [PHLivePhotoEditingContext] class.
type IPHLivePhotoEditingContext interface {
	objectivec.IObject
	// properties:
	ContentEditingOutput() IPHContentEditingOutput
	SetContentEditingOutput(value IPHContentEditingOutput)
	LivePhoto() IPHLivePhoto
	SetLivePhoto(value IPHLivePhoto)
	AdjustmentData() IPHAdjustmentData
	SetAdjustmentData(value IPHAdjustmentData)
	AudioVolume() float32
	SetAudioVolume(value float32)
	Duration() objc.IObject /* cross-framework: Time */
	SetDuration(value objc.IObject /* cross-framework: Time */)
	FrameProcessor() unsafe.Pointer
	SetFrameProcessor(value unsafe.Pointer)
	FullSizeImage() objc.IObject /* cross-framework: Image */
	SetFullSizeImage(value objc.IObject /* cross-framework: Image */)
	Orientation() ImagePropertyOrientation /* not a class type */
	SetOrientation(value ImagePropertyOrientation /* not a class type */)
	PhotoTime() objc.IObject /* cross-framework: Time */
	SetPhotoTime(value objc.IObject /* cross-framework: Time */)
	PHLivePhotoEditingErrorDomain() objc.IObject /* cross-framework: NSString */
	// methods:
}

// An editing session for modifying the photo, video, and audio content of a Live Photo.
//
// A Live Photo is a picture, captured by a supported iOS device, that includes motion and sound from the moments just before and after it was taken. Editing the content of a Live Photo works much like editing other asset types: In an app using the Photos framework, fetch a object that represents the Live Photo to edit, and use that object’s method to retrieve a object. In a photo editing extension that runs within the Photos app, your extension’s main view controller (which adopts the protocol) receives a object when the user chooses to edit a Live Photo with your extension. 2. Create a Live Photo editing context with the initializer. You can create a Live Photo editing context only from object that represents a Live Photo. Use the property of the editing input to verify that it has live Photo content. 3. Use the property to define a block to be used in processing the Live Photo’s visual content. Photos will call this block repeatedly to process each frame of the Live Photo’s video and still photo content. 4. Create a object to store the results of your edit, then call the to process the Live Photo and save it to your editing output object. This method applies your to each frame. To allow a user to continue working with the edit later (for example, to adjust the parameters of a filter), create a object describing your changes, and store it in the property of your editing output. In an app using the Photos framework, use a photo library change block to commit the edit. (For details, see .) In the block, create a object and set its property to the editing output that you created. In a photo editing extension, provide the object that you created in your main view controller’s method. When you use either of the methods listed in Processing an Editing Context’s Live Photo, Photos calls your block repeatedly to process each frame of the Live Photo’s video and still photo content. In that block, a object provides the Live Photo’s existing content as a object. You use Core Image to modify the image, then provide the result of your edits by returning a object representing the result of processing the input image.


// An editing session for modifying the photo, video, and audio content of a Live Photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLivePhotoEditingContext
type PHLivePhotoEditingContext struct {
	objectivec.Object
}

// PHLivePhotoEditingContextFrom constructs a [PHLivePhotoEditingContext] from an unsafe.Pointer.
//
// An editing session for modifying the photo, video, and audio content of a Live Photo.
func PHLivePhotoEditingContextFrom(ptr unsafe.Pointer) PHLivePhotoEditingContext {
	return PHLivePhotoEditingContext{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHLivePhotoEditingContextClass) Alloc() PHLivePhotoEditingContext {
	rv := objc.Send[PHLivePhotoEditingContext](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHLivePhotoEditingContextClass) New() PHLivePhotoEditingContext {
	rv := objc.Send[PHLivePhotoEditingContext](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHLivePhotoEditingContext) Init() PHLivePhotoEditingContext {
	rv := objc.Send[PHLivePhotoEditingContext](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHLivePhotoEditingContext) Autorelease() PHLivePhotoEditingContext {
	rv := objc.Send[PHLivePhotoEditingContext](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHLivePhotoEditingContext creates a new PHLivePhotoEditingContext instance.
func NewPHLivePhotoEditingContext() PHLivePhotoEditingContext {
	return getPHLivePhotoEditingContextClass().New()
}



// The output of an asset content editing session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetchangerequest/contenteditingoutput
func (p_ PHLivePhotoEditingContext) ContentEditingOutput() IPHContentEditingOutput {
	rv := objc.Send[PHContentEditingOutput](p_.ID, objc.Sel("contentEditingOutput"))
	return rv
}


// The output of an asset content editing session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetchangerequest/contenteditingoutput
func (p_ PHLivePhotoEditingContext) SetContentEditingOutput(value IPHContentEditingOutput) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentEditingOutput:"), value)
}


// The unedited Live Photo content of the editing input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/livephoto
func (p_ PHLivePhotoEditingContext) LivePhoto() IPHLivePhoto {
	rv := objc.Send[PHLivePhoto](p_.ID, objc.Sel("livePhoto"))
	return rv
}


// The unedited Live Photo content of the editing input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/livephoto
func (p_ PHLivePhotoEditingContext) SetLivePhoto(value IPHLivePhoto) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLivePhoto:"), value)
}


// An object describing the changes made to the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditingoutput/adjustmentdata
func (p_ PHLivePhotoEditingContext) AdjustmentData() IPHAdjustmentData {
	rv := objc.Send[PHAdjustmentData](p_.ID, objc.Sel("adjustmentData"))
	return rv
}


// An object describing the changes made to the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditingoutput/adjustmentdata
func (p_ PHLivePhotoEditingContext) SetAdjustmentData(value IPHAdjustmentData) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAdjustmentData:"), value)
}


// The audio gain to apply to the processed Live Photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlivephotoeditingcontext/audiovolume
func (p_ PHLivePhotoEditingContext) AudioVolume() float32 {
	rv := objc.Send[float32](p_.ID, objc.Sel("audioVolume"))
	return rv
}


// The audio gain to apply to the processed Live Photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlivephotoeditingcontext/audiovolume
func (p_ PHLivePhotoEditingContext) SetAudioVolume(value float32) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAudioVolume:"), value)
}


// The duration, in seconds, of the Live Photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlivephotoeditingcontext/duration
func (p_ PHLivePhotoEditingContext) Duration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[coremedia.Time](p_.ID, objc.Sel("duration"))
	return rv
}


// The duration, in seconds, of the Live Photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlivephotoeditingcontext/duration
func (p_ PHLivePhotoEditingContext) SetDuration(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDuration:"), value)
}


// A block to be called by Photos for processing each frame of the Live Photo’s visual content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlivephotoeditingcontext/frameprocessor
func (p_ PHLivePhotoEditingContext) FrameProcessor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("frameProcessor"))
	return rv
}


// A block to be called by Photos for processing each frame of the Live Photo’s visual content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlivephotoeditingcontext/frameprocessor
func (p_ PHLivePhotoEditingContext) SetFrameProcessor(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFrameProcessor:"), value)
}


// The unedited still photo content of the Live Photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlivephotoeditingcontext/fullsizeimage
func (p_ PHLivePhotoEditingContext) FullSizeImage() objc.IObject /* cross-framework: Image */ {
	rv := objc.Send[appkit.Image](p_.ID, objc.Sel("fullSizeImage"))
	return rv
}


// The unedited still photo content of the Live Photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlivephotoeditingcontext/fullsizeimage
func (p_ PHLivePhotoEditingContext) SetFullSizeImage(value objc.IObject /* cross-framework: Image */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFullSizeImage:"), value)
}


// The image orientation of the Live Photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlivephotoeditingcontext/orientation
func (p_ PHLivePhotoEditingContext) Orientation() ImagePropertyOrientation /* not a class type */ {
	rv := objc.Send[ImagePropertyOrientation](p_.ID, objc.Sel("orientation"))
	return rv
}


// The image orientation of the Live Photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlivephotoeditingcontext/orientation
func (p_ PHLivePhotoEditingContext) SetOrientation(value ImagePropertyOrientation /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOrientation:"), value)
}


// The offset, in seconds, from the beginning of the Live Photo’s duration to the time corresponding to its still photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlivephotoeditingcontext/phototime
func (p_ PHLivePhotoEditingContext) PhotoTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[coremedia.Time](p_.ID, objc.Sel("photoTime"))
	return rv
}


// The offset, in seconds, from the beginning of the Live Photo’s duration to the time corresponding to its still photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlivephotoeditingcontext/phototime
func (p_ PHLivePhotoEditingContext) SetPhotoTime(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPhotoTime:"), value)
}


// The domain value for error objects produced by a Live Photo editing context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlivephotoeditingerrordomain
func (p_ PHLivePhotoEditingContext) PHLivePhotoEditingErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("PHLivePhotoEditingErrorDomain"))
	return rv
}



