// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVFragmentedMovie */


/* debug [class_header]: Header for AVFragmentedMovie */
// The class instance for the [FragmentedMovie] class.
var (
	FragmentedMovieClass     _FragmentedMovieClass
	FragmentedMovieClassOnce sync.Once
)

func getFragmentedMovieClass() _FragmentedMovieClass {
	FragmentedMovieClassOnce.Do(func() {
		FragmentedMovieClass = _FragmentedMovieClass{objc.GetClass("AVFragmentedMovie")}
	})
	return FragmentedMovieClass
}

type _FragmentedMovieClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FragmentedMovie */
// An interface definition for the [FragmentedMovie] class.
type IFragmentedMovie interface {
	IMovie
	
/* debug [class_interface_properties]: Properties for FragmentedMovie */
	// properties:
	Tracks() []FragmentedMovieTrack
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FragmentedMovie */
	// methods:
	LoadTrackWithTrackIDCompletionHandler(trackID PersistentTrackID /* not a class type */, completionHandler unsafe.Pointer)
	LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic MediaCharacteristic /* typedef */, completionHandler unsafe.Pointer)
	LoadTracksWithMediaTypeCompletionHandler(mediaType MediaType /* typedef */, completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FragmentedMovie */
// Alloc allocates a new instance without initialization.
func (fc _FragmentedMovieClass) Alloc() FragmentedMovie {
	rv := objc.Send[FragmentedMovie](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FragmentedMovieClass) New() FragmentedMovie {
	rv := objc.Send[FragmentedMovie](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FragmentedMovie) Init() FragmentedMovie {
	rv := objc.Send[FragmentedMovie](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FragmentedMovie) Autorelease() FragmentedMovie {
	rv := objc.Send[FragmentedMovie](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFragmentedMovie creates a new FragmentedMovie instance.
func NewFragmentedMovie() FragmentedMovie {
	return getFragmentedMovieClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FragmentedMovie */
// An object that represents a fragmented movie file.


// An object that represents a fragmented movie file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovie
type FragmentedMovie struct {
	Movie
}

// FragmentedMovieFrom constructs a [FragmentedMovie] from an unsafe.Pointer.
//
// An object that represents a fragmented movie file.
func FragmentedMovieFrom(ptr unsafe.Pointer) FragmentedMovie {
	return FragmentedMovie{
		Movie: MovieFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FragmentedMovie *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FragmentedMovie */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FragmentedMovie */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FragmentedMovie */

// Loads a track that contains the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovie/loadTrack(withTrackID:completionHandler:)
func (f_ FragmentedMovie) LoadTrackWithTrackIDCompletionHandler(trackID PersistentTrackID /* not a class type */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("loadTrackWithTrackID:completionHandler:"), trackID, completionHandler)
}/* debug [instance_methods/method]: LoadTrackWithTrackIDCompletionHandler */


// Loads tracks that contain media of a specified characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovie/loadTracks(withMediaCharacteristic:completionHandler:)
func (f_ FragmentedMovie) LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic MediaCharacteristic /* typedef */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("loadTracksWithMediaCharacteristic:completionHandler:"), mediaCharacteristic, completionHandler)
}/* debug [instance_methods/method]: LoadTracksWithMediaCharacteristicCompletionHandler */


// Loads tracks that contain media of a specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovie/loadTracks(withMediaType:completionHandler:)
func (f_ FragmentedMovie) LoadTracksWithMediaTypeCompletionHandler(mediaType MediaType /* typedef */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("loadTracksWithMediaType:completionHandler:"), mediaType, completionHandler)
}/* debug [instance_methods/method]: LoadTracksWithMediaTypeCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FragmentedMovie */

// The tracks that a movie contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovie/tracks
func (f_ FragmentedMovie) Tracks() []FragmentedMovieTrack {
	rv := objc.Send[[]FragmentedMovieTrack](f_.ID, objc.Sel("tracks"))
	return rv
}/* debug [instance_properties/getter]: tracks */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVFragmentedMovie */



