// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [MediaPickerController] class.
var (
	MediaPickerControllerClass     _MediaPickerControllerClass
	MediaPickerControllerClassOnce sync.Once
)

func getMediaPickerControllerClass() _MediaPickerControllerClass {
	MediaPickerControllerClassOnce.Do(func() {
		MediaPickerControllerClass = _MediaPickerControllerClass{objc.GetClass("MPMediaPickerController")}
	})
	return MediaPickerControllerClass
}

type _MediaPickerControllerClass struct {
	class objc.Class
}

// An interface definition for the [MediaPickerController] class.
type IMediaPickerController interface {
	appkit.IViewController
}

// A specialized view controller that provides a graphical interface for selecting media items.
//
// An object, or media item picker, is a specialized view controller that you employ to provide a graphical interface for selecting media items. To display a media item picker, present it modally on an existing view controller. Presenting an in non-modal mode; for example, pushing a onto an existing stack causes your app to crash. describes media items. To respond to user selections and to dismiss a media item picker, use the protocol.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPickerController
type MediaPickerController struct {
	appkit.ViewController
}

// MediaPickerControllerFrom constructs a [MediaPickerController] from an unsafe.Pointer.
//
// A specialized view controller that provides a graphical interface for selecting media items.
func MediaPickerControllerFrom(ptr unsafe.Pointer) MediaPickerController {
	return MediaPickerController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MediaPickerControllerClass) Alloc() MediaPickerController {
	rv := objc.Send[MediaPickerController](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MediaPickerControllerClass) New() MediaPickerController {
	rv := objc.Send[MediaPickerController](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaPickerController) Init() MediaPickerController {
	rv := objc.Send[MediaPickerController](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaPickerController) Autorelease() MediaPickerController {
	rv := objc.Send[MediaPickerController](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaPickerController creates a new MediaPickerController instance.
func NewMediaPickerController() MediaPickerController {
	return getMediaPickerControllerClass().New()
}




// Initializes a media item picker for specified media types.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPickerController/init(mediaTypes:)
func NewMediaPickerControllerWithMediaTypes(mediaTypes unsafe.Pointer) MediaPickerController {
	instance := getMediaPickerControllerClass().Alloc()
	rv := objc.Send[MediaPickerController](instance.ID, objc.Sel("initWithMediaTypes:"), mediaTypes)
	rv.Autorelease()
	return rv
}


// A Boolean value specifying the default selection behavior for a media item picker.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPickerController/allowsPickingMultipleItems
func (m_ MediaPickerController) AllowsPickingMultipleItems() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsPickingMultipleItems"))
	return rv
}


// SetAllowsPickingMultipleItems sets the value of the allowsPickingMultipleItems property.
// A Boolean value specifying the default selection behavior for a media item picker.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPickerController/allowsPickingMultipleItems
func (m_ MediaPickerController) SetAllowsPickingMultipleItems(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllowsPickingMultipleItems:"), value)
}

// The delegate for a media item picker.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPickerController/delegate
func (m_ MediaPickerController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate for a media item picker.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPickerController/delegate
func (m_ MediaPickerController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}

// The media types that media item picker presents.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPickerController/mediaTypes
func (m_ MediaPickerController) MediaTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mediaTypes"))
	return rv
}

// A prompt, for the user, that appears above the navigation bar buttons.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPickerController/prompt
func (m_ MediaPickerController) Prompt() string {
	rv := objc.Send[string](m_.ID, objc.Sel("prompt"))
	return rv
}


// SetPrompt sets the value of the prompt property.
// A prompt, for the user, that appears above the navigation bar buttons.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPickerController/prompt
func (m_ MediaPickerController) SetPrompt(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrompt:"), objc.String(value))
}

// A Boolean value specifying whether to display iCloud Media Library items for a media picker.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPickerController/showsCloudItems
func (m_ MediaPickerController) ShowsCloudItems() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsCloudItems"))
	return rv
}


// SetShowsCloudItems sets the value of the showsCloudItems property.
// A Boolean value specifying whether to display iCloud Media Library items for a media picker.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPickerController/showsCloudItems
func (m_ MediaPickerController) SetShowsCloudItems(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsCloudItems:"), value)
}

// A Boolean value that specifies whether the media item picker displays protected assets.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPickerController/showsItemsWithProtectedAssets
func (m_ MediaPickerController) ShowsItemsWithProtectedAssets() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("showsItemsWithProtectedAssets"))
	return rv
}


// SetShowsItemsWithProtectedAssets sets the value of the showsItemsWithProtectedAssets property.
// A Boolean value that specifies whether the media item picker displays protected assets.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPickerController/showsItemsWithProtectedAssets
func (m_ MediaPickerController) SetShowsItemsWithProtectedAssets(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShowsItemsWithProtectedAssets:"), value)
}


