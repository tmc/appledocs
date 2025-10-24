// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	// methods:
}

// A specialized view controller that provides a graphical interface for selecting media items.
//
// An object, or media item picker, is a specialized view controller that you employ to provide a graphical interface for selecting media items. To display a media item picker, present it modally on an existing view controller. Presenting an in non-modal mode; for example, pushing a onto an existing stack causes your app to crash. describes media items. To respond to user selections and to dismiss a media item picker, use the protocol.


// A specialized view controller that provides a graphical interface for selecting media items.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPickerController/init(mediaTypes:)
func NewMediaPickerControllerWithMediaTypes(mediaTypes MediaType) MediaPickerController {
	instance := getMediaPickerControllerClass().Alloc()
	rv := objc.Send[MediaPickerController](instance.ID, objc.Sel("initWithMediaTypes:"), mediaTypes)
	rv.Autorelease()
	return rv
}



