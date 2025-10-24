// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSMediaLibraryBrowserController */


/* debug [class_header]: Header for NSMediaLibraryBrowserController */
// The class instance for the [MediaLibraryBrowserController] class.
var (
	MediaLibraryBrowserControllerClass     _MediaLibraryBrowserControllerClass
	MediaLibraryBrowserControllerClassOnce sync.Once
)

func getMediaLibraryBrowserControllerClass() _MediaLibraryBrowserControllerClass {
	MediaLibraryBrowserControllerClassOnce.Do(func() {
		MediaLibraryBrowserControllerClass = _MediaLibraryBrowserControllerClass{objc.GetClass("NSMediaLibraryBrowserController")}
	})
	return MediaLibraryBrowserControllerClass
}

type _MediaLibraryBrowserControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MediaLibraryBrowserController */
// An interface definition for the [MediaLibraryBrowserController] class.
type IMediaLibraryBrowserController interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MediaLibraryBrowserController */
	// properties:
	Frame() Rect /* not a class type */
	SetFrame(value Rect /* not a class type */)
	Visible() bool
	SetVisible(value bool)
	MediaLibraries() MediaLibrary
	SetMediaLibraries(value MediaLibrary)
	IsVisible() bool
	SetIsVisible(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MediaLibraryBrowserController */
	// methods:
	TogglePanel(sender objc.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MediaLibraryBrowserController */
// Alloc allocates a new instance without initialization.
func (mc _MediaLibraryBrowserControllerClass) Alloc() MediaLibraryBrowserController {
	rv := objc.Send[MediaLibraryBrowserController](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MediaLibraryBrowserControllerClass) New() MediaLibraryBrowserController {
	rv := objc.Send[MediaLibraryBrowserController](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaLibraryBrowserController) Init() MediaLibraryBrowserController {
	rv := objc.Send[MediaLibraryBrowserController](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaLibraryBrowserController) Autorelease() MediaLibraryBrowserController {
	rv := objc.Send[MediaLibraryBrowserController](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaLibraryBrowserController creates a new MediaLibraryBrowserController instance.
func NewMediaLibraryBrowserController() MediaLibraryBrowserController {
	return getMediaLibraryBrowserControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MediaLibraryBrowserController */
// An object that configures and displays a Media Library Browser panel.
//
// From this panel a user can drag media into views in their app. The class provides a standard interface to the MediaLibrary framework content. For more information see , , , and in .


// An object that configures and displays a Media Library Browser panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController
type MediaLibraryBrowserController struct {
	objectivec.Object
}

// MediaLibraryBrowserControllerFrom constructs a [MediaLibraryBrowserController] from an unsafe.Pointer.
//
// An object that configures and displays a Media Library Browser panel.
func MediaLibraryBrowserControllerFrom(ptr unsafe.Pointer) MediaLibraryBrowserController {
	return MediaLibraryBrowserController{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MediaLibraryBrowserController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MediaLibraryBrowserController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MediaLibraryBrowserController */

// Returns the shared Media Library Browser instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/shared
func (mc _MediaLibraryBrowserControllerClass) SharedMediaLibraryBrowserController() MediaLibraryBrowserController {
	rv := objc.Send[MediaLibraryBrowserController](objc.ID(mc.class), objc.Sel("sharedMediaLibraryBrowserController"))
	return rv
}/* debug [class_properties_class/property]: sharedMediaLibraryBrowserController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MediaLibraryBrowserController */

// Toggles the visibility of the Media Library Browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/togglePanel(_:)
func (m_ MediaLibraryBrowserController) TogglePanel(sender objc.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("togglePanel:"), sender)
}/* debug [instance_methods/method]: TogglePanel */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MediaLibraryBrowserController */

// The frame, in global coordinates, used to display the Media Library Browser panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/frame
func (m_ MediaLibraryBrowserController) Frame() Rect /* not a class type */ {
	rv := objc.Send[Rect](m_.ID, objc.Sel("frame"))
	return rv
}/* debug [instance_properties/getter]: frame */


// The frame, in global coordinates, used to display the Media Library Browser panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/frame
func (m_ MediaLibraryBrowserController) SetFrame(value Rect /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFrame:"), value)
}/* debug [instance_properties/setter]: frame */


// A Boolean value that determines whether the Media Library Browser panel is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/isVisible
func (m_ MediaLibraryBrowserController) Visible() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("visible"))
	return rv
}/* debug [instance_properties/getter]: visible */


// A Boolean value that determines whether the Media Library Browser panel is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/isVisible
func (m_ MediaLibraryBrowserController) SetVisible(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVisible:"), value)
}/* debug [instance_properties/setter]: visible */


// The media library that is in use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/mediaLibraries
func (m_ MediaLibraryBrowserController) MediaLibraries() MediaLibrary {
	rv := objc.Send[MediaLibrary](m_.ID, objc.Sel("mediaLibraries"))
	return rv
}/* debug [instance_properties/getter]: mediaLibraries */


// The media library that is in use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/mediaLibraries
func (m_ MediaLibraryBrowserController) SetMediaLibraries(value MediaLibrary) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMediaLibraries:"), value)
}/* debug [instance_properties/setter]: mediaLibraries */


// Returns the shared Media Library Browser instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/shared
func (m_ MediaLibraryBrowserController) SharedMediaLibraryBrowserController() IMediaLibraryBrowserController {
	rv := objc.Send[MediaLibraryBrowserController](m_.ID, objc.Sel("sharedMediaLibraryBrowserController"))
	return rv
}/* debug [instance_properties/getter]: sharedMediaLibraryBrowserController */


// A Boolean value that determines whether the Media Library Browser panel is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmedialibrarybrowsercontroller/isvisible
func (m_ MediaLibraryBrowserController) IsVisible() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isVisible"))
	return rv
}/* debug [instance_properties/getter]: isVisible */


// A Boolean value that determines whether the Media Library Browser panel is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmedialibrarybrowsercontroller/isvisible
func (m_ MediaLibraryBrowserController) SetIsVisible(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsVisible:"), value)
}/* debug [instance_properties/setter]: isVisible */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMediaLibraryBrowserController */



