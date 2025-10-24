// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPTimedMetadata */


/* debug [class_header]: Header for MPTimedMetadata */
// The class instance for the [TimedMetadata] class.
var (
	TimedMetadataClass     _TimedMetadataClass
	TimedMetadataClassOnce sync.Once
)

func getTimedMetadataClass() _TimedMetadataClass {
	TimedMetadataClassOnce.Do(func() {
		TimedMetadataClass = _TimedMetadataClass{objc.GetClass("MPTimedMetadata")}
	})
	return TimedMetadataClass
}

type _TimedMetadataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TimedMetadata */
// An interface definition for the [TimedMetadata] class.
type ITimedMetadata interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TimedMetadata */
	// properties:
	ImageCropRect() corefoundation.CGRect
	SetImageCropRect(value corefoundation.CGRect)
	MPMoviePlayerTimedMetadataUserInfoKey() objc.IObject /* cross-framework: NSString */
	ShowsRouteButton() bool
	SetShowsRouteButton(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TimedMetadata */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TimedMetadata */
// Alloc allocates a new instance without initialization.
func (tc _TimedMetadataClass) Alloc() TimedMetadata {
	rv := objc.Send[TimedMetadata](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TimedMetadataClass) New() TimedMetadata {
	rv := objc.Send[TimedMetadata](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TimedMetadata) Init() TimedMetadata {
	rv := objc.Send[TimedMetadata](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TimedMetadata) Autorelease() TimedMetadata {
	rv := objc.Send[TimedMetadata](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTimedMetadata creates a new TimedMetadata instance.
func NewTimedMetadata() TimedMetadata {
	return getTimedMetadataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TimedMetadata */
// A carries time-based information within HTTP streamed media.
//
// Content providers can embed these objects when creating a stream. The properties and constants in this class let you extract the metadata as you play the stream using an object. For example, the provider of a live sports video stream could use instances to embed game scores, with timestamps, in the stream. On the client side—that is, on the user’s device—their application could employ the properties of this class to update their app’s user interface in real time during the game. A Javascript implementation of this class is also available for use by web-based applications.


// A carries time-based information within HTTP streamed media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPTimedMetadata
type TimedMetadata struct {
	objectivec.Object
}

// TimedMetadataFrom constructs a [TimedMetadata] from an unsafe.Pointer.
//
// A carries time-based information within HTTP streamed media.
func TimedMetadataFrom(ptr unsafe.Pointer) TimedMetadata {
	return TimedMetadata{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TimedMetadata *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TimedMetadata */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TimedMetadata */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TimedMetadata */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TimedMetadata */

// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (t_ TimedMetadata) ImageCropRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t_.ID, objc.Sel("imageCropRect"))
	return rv
}/* debug [instance_properties/getter]: imageCropRect */


// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (t_ TimedMetadata) SetImageCropRect(value corefoundation.CGRect) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setImageCropRect:"), value)
}/* debug [instance_properties/setter]: imageCropRect */


// An NSDictionary object containing the most recent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmovieplayertimedmetadatauserinfokey
func (t_ TimedMetadata) MPMoviePlayerTimedMetadataUserInfoKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("MPMoviePlayerTimedMetadataUserInfoKey"))
	return rv
}/* debug [instance_properties/getter]: MPMoviePlayerTimedMetadataUserInfoKey */


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (t_ TimedMetadata) ShowsRouteButton() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("showsRouteButton"))
	return rv
}/* debug [instance_properties/getter]: showsRouteButton */


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (t_ TimedMetadata) SetShowsRouteButton(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setShowsRouteButton:"), value)
}/* debug [instance_properties/setter]: showsRouteButton */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPTimedMetadata */


