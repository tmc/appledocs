// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSMovie */


/* debug [class_header]: Header for NSMovie */
// The class instance for the [Movie] class.
var (
	MovieClass     _MovieClass
	MovieClassOnce sync.Once
)

func getMovieClass() _MovieClass {
	MovieClassOnce.Do(func() {
		MovieClass = _MovieClass{objc.GetClass("NSMovie")}
	})
	return MovieClass
}

type _MovieClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Movie */
// An interface definition for the [Movie] class.
type IMovie interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Movie */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Movie */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Movie */
// Alloc allocates a new instance without initialization.
func (mc _MovieClass) Alloc() Movie {
	rv := objc.Send[Movie](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MovieClass) New() Movie {
	rv := objc.Send[Movie](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ Movie) Init() Movie {
	rv := objc.Send[Movie](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ Movie) Autorelease() Movie {
	rv := objc.Send[Movie](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMovie creates a new Movie instance.
func NewMovie() Movie {
	return getMovieClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Movie */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMovie
type Movie struct {
	objectivec.Object
}

// MovieFrom constructs a [Movie] from an unsafe.Pointer.
func MovieFrom(ptr unsafe.Pointer) Movie {
	return Movie{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Movie */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMovie/initWithCoder:
func NewMovieWithCoder(coder foundation.Coder) Movie {
	instance := getMovieClass().Alloc()
	rv := objc.Send[Movie](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMovieWithCoder */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMovie/initWithMovie:
func NewMovieWithMovie(movie objectivec.IObject) Movie {
	instance := getMovieClass().Alloc()
	rv := objc.Send[Movie](instance.ID, objc.Sel("initWithMovie:"), movie)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMovieWithMovie */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Movie */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Movie */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Movie */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Movie */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMovie */


