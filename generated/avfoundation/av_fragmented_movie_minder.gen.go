// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVFragmentedMovieMinder */


/* debug [class_header]: Header for AVFragmentedMovieMinder */
// The class instance for the [FragmentedMovieMinder] class.
var (
	FragmentedMovieMinderClass     _FragmentedMovieMinderClass
	FragmentedMovieMinderClassOnce sync.Once
)

func getFragmentedMovieMinderClass() _FragmentedMovieMinderClass {
	FragmentedMovieMinderClassOnce.Do(func() {
		FragmentedMovieMinderClass = _FragmentedMovieMinderClass{objc.GetClass("AVFragmentedMovieMinder")}
	})
	return FragmentedMovieMinderClass
}

type _FragmentedMovieMinderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FragmentedMovieMinder */
// An interface definition for the [FragmentedMovieMinder] class.
type IFragmentedMovieMinder interface {
	IFragmentedAssetMinder
	
/* debug [class_interface_properties]: Properties for FragmentedMovieMinder */
	// properties:
	MindingInterval() float64
	SetMindingInterval(value float64)
	Movies() []FragmentedMovie
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FragmentedMovieMinder */
	// methods:
	AddFragmentedMovie(movie IAVFragmentedMovie)
	RemoveFragmentedMovie(movie IAVFragmentedMovie)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FragmentedMovieMinder */
// Alloc allocates a new instance without initialization.
func (fc _FragmentedMovieMinderClass) Alloc() FragmentedMovieMinder {
	rv := objc.Send[FragmentedMovieMinder](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FragmentedMovieMinderClass) New() FragmentedMovieMinder {
	rv := objc.Send[FragmentedMovieMinder](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FragmentedMovieMinder) Init() FragmentedMovieMinder {
	rv := objc.Send[FragmentedMovieMinder](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FragmentedMovieMinder) Autorelease() FragmentedMovieMinder {
	rv := objc.Send[FragmentedMovieMinder](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFragmentedMovieMinder creates a new FragmentedMovieMinder instance.
func NewFragmentedMovieMinder() FragmentedMovieMinder {
	return getFragmentedMovieMinderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FragmentedMovieMinder */
// An object that checks whether a fragmented movie appends additional movie fragments.
//
// This class is identical to except that it’s capable of minding only assets of type .


// An object that checks whether a fragmented movie appends additional movie fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieMinder
type FragmentedMovieMinder struct {
	FragmentedAssetMinder
}

// FragmentedMovieMinderFrom constructs a [FragmentedMovieMinder] from an unsafe.Pointer.
//
// An object that checks whether a fragmented movie appends additional movie fragments.
func FragmentedMovieMinderFrom(ptr unsafe.Pointer) FragmentedMovieMinder {
	return FragmentedMovieMinder{
		FragmentedAssetMinder: FragmentedAssetMinderFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FragmentedMovieMinder */

// Creates a movie minder and adds a movie with a minding interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieMinder/init(movie:mindingInterval:)
func NewFragmentedMovieMinderWithMovieMindingInterval(movie IAVFragmentedMovie, mindingInterval float64) FragmentedMovieMinder {
	instance := getFragmentedMovieMinderClass().Alloc()
	rv := objc.Send[FragmentedMovieMinder](instance.ID, objc.Sel("initWithMovie:mindingInterval:"), movie, mindingInterval)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewFragmentedMovieMinderWithMovieMindingInterval */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FragmentedMovieMinder */

// Creates a movie minder and adds a movie with a minding interval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieMinder/fragmentedMovieMinderWithMovie:mindingInterval:
func (fc _FragmentedMovieMinderClass) FragmentedMovieMinderWithMovieMindingInterval(movie IAVFragmentedMovie, mindingInterval float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(fc.class), objc.Sel("fragmentedMovieMinderWithMovie:mindingInterval:"), movie, mindingInterval)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FragmentedMovieMinderWithMovieMindingInterval) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FragmentedMovieMinder */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FragmentedMovieMinder */

// Adds a fragmented movie to the array of movies being minded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieMinder/add(_:)
func (f_ FragmentedMovieMinder) AddFragmentedMovie(movie IAVFragmentedMovie) {
	objc.Send[objc.ID](f_.ID, objc.Sel("addFragmentedMovie:"), movie)
}/* debug [instance_methods/method]: AddFragmentedMovie */


// Removes a fragmented movie from the array of movies being minded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieMinder/remove(_:)
func (f_ FragmentedMovieMinder) RemoveFragmentedMovie(movie IAVFragmentedMovie) {
	objc.Send[objc.ID](f_.ID, objc.Sel("removeFragmentedMovie:"), movie)
}/* debug [instance_methods/method]: RemoveFragmentedMovie */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FragmentedMovieMinder */

// The amount of time between checks for additional movie fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieMinder/mindingInterval
func (f_ FragmentedMovieMinder) MindingInterval() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("mindingInterval"))
	return rv
}/* debug [instance_properties/getter]: mindingInterval */


// The amount of time between checks for additional movie fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieMinder/mindingInterval
func (f_ FragmentedMovieMinder) SetMindingInterval(value float64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setMindingInterval:"), value)
}/* debug [instance_properties/setter]: mindingInterval */


// An array containing the fragmented movie objects being minded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieMinder/movies
func (f_ FragmentedMovieMinder) Movies() []FragmentedMovie {
	rv := objc.Send[[]FragmentedMovie](f_.ID, objc.Sel("movies"))
	return rv
}/* debug [instance_properties/getter]: movies */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVFragmentedMovieMinder */


