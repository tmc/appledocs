// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVFragmentedMovieTrack */


/* debug [class_header]: Header for AVFragmentedMovieTrack */
// The class instance for the [FragmentedMovieTrack] class.
var (
	FragmentedMovieTrackClass     _FragmentedMovieTrackClass
	FragmentedMovieTrackClassOnce sync.Once
)

func getFragmentedMovieTrackClass() _FragmentedMovieTrackClass {
	FragmentedMovieTrackClassOnce.Do(func() {
		FragmentedMovieTrackClass = _FragmentedMovieTrackClass{objc.GetClass("AVFragmentedMovieTrack")}
	})
	return FragmentedMovieTrackClass
}

type _FragmentedMovieTrackClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FragmentedMovieTrack */
// An interface definition for the [FragmentedMovieTrack] class.
type IFragmentedMovieTrack interface {
	IMovieTrack
	
/* debug [class_interface_properties]: Properties for FragmentedMovieTrack */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FragmentedMovieTrack */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FragmentedMovieTrack */
// Alloc allocates a new instance without initialization.
func (fc _FragmentedMovieTrackClass) Alloc() FragmentedMovieTrack {
	rv := objc.Send[FragmentedMovieTrack](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FragmentedMovieTrackClass) New() FragmentedMovieTrack {
	rv := objc.Send[FragmentedMovieTrack](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FragmentedMovieTrack) Init() FragmentedMovieTrack {
	rv := objc.Send[FragmentedMovieTrack](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FragmentedMovieTrack) Autorelease() FragmentedMovieTrack {
	rv := objc.Send[FragmentedMovieTrack](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFragmentedMovieTrack creates a new FragmentedMovieTrack instance.
func NewFragmentedMovieTrack() FragmentedMovieTrack {
	return getFragmentedMovieTrackClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FragmentedMovieTrack */
// An object that represents a track in a fragmented movie.


// An object that represents a track in a fragmented movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieTrack
type FragmentedMovieTrack struct {
	MovieTrack
}

// FragmentedMovieTrackFrom constructs a [FragmentedMovieTrack] from an unsafe.Pointer.
//
// An object that represents a track in a fragmented movie.
func FragmentedMovieTrackFrom(ptr unsafe.Pointer) FragmentedMovieTrack {
	return FragmentedMovieTrack{
		MovieTrack: MovieTrackFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FragmentedMovieTrack *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FragmentedMovieTrack */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FragmentedMovieTrack */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FragmentedMovieTrack */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FragmentedMovieTrack */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVFragmentedMovieTrack */



