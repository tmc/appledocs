// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	Cancel()
	PrepareLivePhotoForPlaybackWithTargetSizeOptionsCompletionHandler(targetSize coregraphics.CGSize, options unsafe.Pointer, handler unsafe.Pointer)
	SaveLivePhotoToOutputOptionsCompletionHandler(output IPHContentEditingOutput, options unsafe.Pointer, handler unsafe.Pointer)
}

// An editing session for modifying the photo, video, and audio content of a Live Photo.
//
// A Live Photo is a picture, captured by a supported iOS device, that includes motion and sound from the moments just before and after it was taken. Editing the content of a Live Photo works much like editing other asset types: In an app using the Photos framework, fetch a object that represents the Live Photo to edit, and use that object’s method to retrieve a object. In a photo editing extension that runs within the Photos app, your extension’s main view controller (which adopts the protocol) receives a object when the user chooses to edit a Live Photo with your extension. 2. Create a Live Photo editing context with the initializer. You can create a Live Photo editing context only from object that represents a Live Photo. Use the property of the editing input to verify that it has live Photo content. 3. Use the property to define a block to be used in processing the Live Photo’s visual content. Photos will call this block repeatedly to process each frame of the Live Photo’s video and still photo content. 4. Create a object to store the results of your edit, then call the to process the Live Photo and save it to your editing output object. This method applies your to each frame. To allow a user to continue working with the edit later (for example, to adjust the parameters of a filter), create a object describing your changes, and store it in the property of your editing output. In an app using the Photos framework, use a photo library change block to commit the edit. (For details, see .) In the block, create a object and set its property to the editing output that you created. In a photo editing extension, provide the object that you created in your main view controller’s method. When you use either of the methods listed in Processing an Editing Context’s Live Photo, Photos calls your block repeatedly to process each frame of the Live Photo’s video and still photo content. In that block, a object provides the Live Photo’s existing content as a object. You use Core Image to modify the image, then provide the result of your edits by returning a object representing the result of processing the input image.
//
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




// Creates a Live Photo editing context for the specified editing input.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLivePhotoEditingContext/init(livePhotoEditingInput:)
func NewPHLivePhotoEditingContextWithLivePhotoEditingInput(livePhotoInput IPHContentEditingInput) PHLivePhotoEditingContext {
	instance := getPHLivePhotoEditingContextClass().Alloc()
	rv := objc.Send[PHLivePhotoEditingContext](instance.ID, objc.Sel("initWithLivePhotoEditingInput:"), livePhotoInput)
	rv.Autorelease()
	return rv
}


// Aborts any Live Photo processing in progress.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLivePhotoEditingContext/cancel()
func (p_ PHLivePhotoEditingContext) Cancel() {
	objc.Send[objc.ID](p_.ID, objc.Sel("cancel"))
}

// Processes a Live Photo with your edits for viewing.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLivePhotoEditingContext/prepareLivePhotoForPlayback(withTargetSize:options:completionHandler:)
func (p_ PHLivePhotoEditingContext) PrepareLivePhotoForPlaybackWithTargetSizeOptionsCompletionHandler(targetSize coregraphics.CGSize, options unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("prepareLivePhotoForPlaybackWithTargetSize:options:completionHandler:"), targetSize, options, handler)
}

// Processes and saves a full-quality Live Photo as the output of your editing session.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLivePhotoEditingContext/saveLivePhoto(to:options:completionHandler:)
func (p_ PHLivePhotoEditingContext) SaveLivePhotoToOutputOptionsCompletionHandler(output IPHContentEditingOutput, options unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("saveLivePhotoToOutput:options:completionHandler:"), output, options, handler)
}

// The audio gain to apply to the processed Live Photo.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLivePhotoEditingContext/audioVolume
func (p_ PHLivePhotoEditingContext) AudioVolume() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("audioVolume"))
	return rv
}


// SetAudioVolume sets the value of the audioVolume property.
// The audio gain to apply to the processed Live Photo.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLivePhotoEditingContext/audioVolume
func (p_ PHLivePhotoEditingContext) SetAudioVolume(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAudioVolume:"), value)
}

// The duration, in seconds, of the Live Photo.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLivePhotoEditingContext/duration
func (p_ PHLivePhotoEditingContext) Duration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("duration"))
	return rv
}

// A block to be called by Photos for processing each frame of the Live Photo’s visual content.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLivePhotoEditingContext/frameProcessor
func (p_ PHLivePhotoEditingContext) FrameProcessor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("frameProcessor"))
	return rv
}


// SetFrameProcessor sets the value of the frameProcessor property.
// A block to be called by Photos for processing each frame of the Live Photo’s visual content.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLivePhotoEditingContext/frameProcessor
func (p_ PHLivePhotoEditingContext) SetFrameProcessor(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFrameProcessor:"), value)
}

// The unedited still photo content of the Live Photo.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLivePhotoEditingContext/fullSizeImage
func (p_ PHLivePhotoEditingContext) FullSizeImage() appkit.Image {
	rv := objc.Send[appkit.Image](p_.ID, objc.Sel("fullSizeImage"))
	return rv
}

// The image orientation of the Live Photo.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLivePhotoEditingContext/orientation
func (p_ PHLivePhotoEditingContext) Orientation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("orientation"))
	return rv
}

// The offset, in seconds, from the beginning of the Live Photo’s duration to the time corresponding to its still photo.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLivePhotoEditingContext/photoTime
func (p_ PHLivePhotoEditingContext) PhotoTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("photoTime"))
	return rv
}

// The output of an asset content editing session.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetchangerequest/contenteditingoutput
func (p_ PHLivePhotoEditingContext) ContentEditingOutput() PHContentEditingOutput {
	rv := objc.Send[PHContentEditingOutput](p_.ID, objc.Sel("contentEditingOutput"))
	return rv
}


// SetContentEditingOutput sets the value of the contentEditingOutput property.
// The output of an asset content editing session.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetchangerequest/contenteditingoutput
func (p_ PHLivePhotoEditingContext) SetContentEditingOutput(value IPHContentEditingOutput) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentEditingOutput:"), value)
}

// The unedited Live Photo content of the editing input.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/livephoto
func (p_ PHLivePhotoEditingContext) LivePhoto() PHLivePhoto {
	rv := objc.Send[PHLivePhoto](p_.ID, objc.Sel("livePhoto"))
	return rv
}


// SetLivePhoto sets the value of the livePhoto property.
// The unedited Live Photo content of the editing input.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditinginput/livephoto
func (p_ PHLivePhotoEditingContext) SetLivePhoto(value IPHLivePhoto) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLivePhoto:"), value)
}

// An object describing the changes made to the asset.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditingoutput/adjustmentdata
func (p_ PHLivePhotoEditingContext) AdjustmentData() PHAdjustmentData {
	rv := objc.Send[PHAdjustmentData](p_.ID, objc.Sel("adjustmentData"))
	return rv
}


// SetAdjustmentData sets the value of the adjustmentData property.
// An object describing the changes made to the asset.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phcontenteditingoutput/adjustmentdata
func (p_ PHLivePhotoEditingContext) SetAdjustmentData(value IPHAdjustmentData) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAdjustmentData:"), value)
}

// The domain value for error objects produced by a Live Photo editing context.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phlivephotoeditingerrordomain
func (p_ PHLivePhotoEditingContext) PHLivePhotoEditingErrorDomain() appkit.string {
	rv := objc.Send[appkit.string](p_.ID, objc.Sel("PHLivePhotoEditingErrorDomain"))
	return rv
}


