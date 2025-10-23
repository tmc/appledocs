// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [Movie] class.
type IMovie interface {
	objectivec.IObject
	QTMovie() unsafe.Pointer
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMovie
type Movie struct {
	objectivec.Object
}

// MovieFrom constructs a [Movie] from an unsafe.Pointer.
func MovieFrom(ptr unsafe.Pointer) Movie {
	return Movie{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MovieClass) Alloc() Movie {
	rv := objc.Send[Movie](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMovie/QTMovie
func (m_ Movie) QTMovie() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("QTMovie"))
	return rv
}



