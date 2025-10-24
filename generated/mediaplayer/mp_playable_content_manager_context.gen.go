// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPPlayableContentManagerContext */


/* debug [class_header]: Header for MPPlayableContentManagerContext */
// The class instance for the [PlayableContentManagerContext] class.
var (
	PlayableContentManagerContextClass     _PlayableContentManagerContextClass
	PlayableContentManagerContextClassOnce sync.Once
)

func getPlayableContentManagerContextClass() _PlayableContentManagerContextClass {
	PlayableContentManagerContextClassOnce.Do(func() {
		PlayableContentManagerContextClass = _PlayableContentManagerContextClass{objc.GetClass("MPPlayableContentManagerContext")}
	})
	return PlayableContentManagerContextClass
}

type _PlayableContentManagerContextClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PlayableContentManagerContext */
// An interface definition for the [PlayableContentManagerContext] class.
type IPlayableContentManagerContext interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PlayableContentManagerContext */
	// properties:
	ImageCropRect() corefoundation.CGRect
	SetImageCropRect(value corefoundation.CGRect)
	ShowsRouteButton() bool
	SetShowsRouteButton(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PlayableContentManagerContext */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PlayableContentManagerContext */
// Alloc allocates a new instance without initialization.
func (pc _PlayableContentManagerContextClass) Alloc() PlayableContentManagerContext {
	rv := objc.Send[PlayableContentManagerContext](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayableContentManagerContextClass) New() PlayableContentManagerContext {
	rv := objc.Send[PlayableContentManagerContext](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayableContentManagerContext) Init() PlayableContentManagerContext {
	rv := objc.Send[PlayableContentManagerContext](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayableContentManagerContext) Autorelease() PlayableContentManagerContext {
	rv := objc.Send[PlayableContentManagerContext](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayableContentManagerContext creates a new PlayableContentManagerContext instance.
func NewPlayableContentManagerContext() PlayableContentManagerContext {
	return getPlayableContentManagerContextClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PlayableContentManagerContext */
// An object representing the current state of the playable endpoint.


// An object representing the current state of the playable endpoint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPPlayableContentManagerContext
type PlayableContentManagerContext struct {
	objectivec.Object
}

// PlayableContentManagerContextFrom constructs a [PlayableContentManagerContext] from an unsafe.Pointer.
//
// An object representing the current state of the playable endpoint.
func PlayableContentManagerContextFrom(ptr unsafe.Pointer) PlayableContentManagerContext {
	return PlayableContentManagerContext{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PlayableContentManagerContext *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PlayableContentManagerContext */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PlayableContentManagerContext */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PlayableContentManagerContext */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PlayableContentManagerContext */

// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (p_ PlayableContentManagerContext) ImageCropRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p_.ID, objc.Sel("imageCropRect"))
	return rv
}/* debug [instance_properties/getter]: imageCropRect */


// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (p_ PlayableContentManagerContext) SetImageCropRect(value corefoundation.CGRect) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImageCropRect:"), value)
}/* debug [instance_properties/setter]: imageCropRect */


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (p_ PlayableContentManagerContext) ShowsRouteButton() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsRouteButton"))
	return rv
}/* debug [instance_properties/getter]: showsRouteButton */


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (p_ PlayableContentManagerContext) SetShowsRouteButton(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsRouteButton:"), value)
}/* debug [instance_properties/setter]: showsRouteButton */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPPlayableContentManagerContext */


