// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMovieTrack */


/* debug [class_header]: Header for AVMovieTrack */
// The class instance for the [MovieTrack] class.
var (
	MovieTrackClass     _MovieTrackClass
	MovieTrackClassOnce sync.Once
)

func getMovieTrackClass() _MovieTrackClass {
	MovieTrackClassOnce.Do(func() {
		MovieTrackClass = _MovieTrackClass{objc.GetClass("AVMovieTrack")}
	})
	return MovieTrackClass
}

type _MovieTrackClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MovieTrack */
// An interface definition for the [MovieTrack] class.
type IMovieTrack interface {
	IAssetTrack
	
/* debug [class_interface_properties]: Properties for MovieTrack */
	// properties:
	AlternateGroupID() int
	MediaDataStorage() IAVMediaDataStorage
	MediaDecodeTimeRange() TimeRange /* not a class type */
	MediaPresentationTimeRange() TimeRange /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MovieTrack */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MovieTrack */
// Alloc allocates a new instance without initialization.
func (mc _MovieTrackClass) Alloc() MovieTrack {
	rv := objc.Send[MovieTrack](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MovieTrackClass) New() MovieTrack {
	rv := objc.Send[MovieTrack](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MovieTrack) Init() MovieTrack {
	rv := objc.Send[MovieTrack](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MovieTrack) Autorelease() MovieTrack {
	rv := objc.Send[MovieTrack](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMovieTrack creates a new MovieTrack instance.
func NewMovieTrack() MovieTrack {
	return getMovieTrackClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MovieTrack */
// A track in a movie that conforms to the QuickTime or ISO base media file format.


// A track in a movie that conforms to the QuickTime or ISO base media file format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovieTrack
type MovieTrack struct {
	AssetTrack
}

// MovieTrackFrom constructs a [MovieTrack] from an unsafe.Pointer.
//
// A track in a movie that conforms to the QuickTime or ISO base media file format.
func MovieTrackFrom(ptr unsafe.Pointer) MovieTrack {
	return MovieTrack{
		AssetTrack: AssetTrackFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MovieTrack *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MovieTrack */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MovieTrack */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MovieTrack */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MovieTrack */

// A value that identifies the track as a member of a particular alternate group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovieTrack/alternateGroupID
func (m_ MovieTrack) AlternateGroupID() int {
	rv := objc.Send[int](m_.ID, objc.Sel("alternateGroupID"))
	return rv
}/* debug [instance_properties/getter]: alternateGroupID */


// The storage container for media data added to a track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovieTrack/mediaDataStorage
func (m_ MovieTrack) MediaDataStorage() IAVMediaDataStorage {
	rv := objc.Send[MediaDataStorage](m_.ID, objc.Sel("mediaDataStorage"))
	return rv
}/* debug [instance_properties/getter]: mediaDataStorage */


// A range of decode times for the track’s media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovieTrack/mediaDecodeTimeRange
func (m_ MovieTrack) MediaDecodeTimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](m_.ID, objc.Sel("mediaDecodeTimeRange"))
	return rv
}/* debug [instance_properties/getter]: mediaDecodeTimeRange */


// A range of presentation times for the track’s media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovieTrack/mediaPresentationTimeRange
func (m_ MovieTrack) MediaPresentationTimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](m_.ID, objc.Sel("mediaPresentationTimeRange"))
	return rv
}/* debug [instance_properties/getter]: mediaPresentationTimeRange */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMovieTrack */



