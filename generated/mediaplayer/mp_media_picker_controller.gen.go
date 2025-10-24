// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MPMediaPickerController */


/* debug [class_header]: Header for MPMediaPickerController */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MediaPickerController */
// An interface definition for the [MediaPickerController] class.
type IMediaPickerController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for MediaPickerController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MediaPickerController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MediaPickerController */
// Alloc allocates a new instance without initialization.
func (mc _MediaPickerControllerClass) Alloc() MediaPickerController {
	rv := objc.Send[MediaPickerController](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MediaPickerController */
// A specialized view controller that provides a graphical interface for selecting media items.
//
// An object, or media item picker, is a specialized view controller that you employ to provide a graphical interface for selecting media items. To display a media item picker, present it modally on an existing view controller. Presenting an in non-modal mode; for example, pushing a onto an existing stack causes your app to crash. describes media items. To respond to user selections and to dismiss a media item picker, use the protocol.


// A specialized view controller that provides a graphical interface for selecting media items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPickerController
type MediaPickerController struct {
	ViewController
}

// MediaPickerControllerFrom constructs a [MediaPickerController] from an unsafe.Pointer.
//
// A specialized view controller that provides a graphical interface for selecting media items.
func MediaPickerControllerFrom(ptr unsafe.Pointer) MediaPickerController {
	return MediaPickerController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MediaPickerController */

// Initializes a media item picker for specified media types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPickerController/init(mediaTypes:)
func NewMediaPickerControllerWithMediaTypes(mediaTypes MediaType) MediaPickerController {
	instance := getMediaPickerControllerClass().Alloc()
	rv := objc.Send[MediaPickerController](instance.ID, objc.Sel("initWithMediaTypes:"), mediaTypes)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMediaPickerControllerWithMediaTypes */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MediaPickerController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MediaPickerController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MediaPickerController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MediaPickerController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPMediaPickerController */


