// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MovieErrorLog] class.
var (
	MovieErrorLogClass     _MovieErrorLogClass
	MovieErrorLogClassOnce sync.Once
)

func getMovieErrorLogClass() _MovieErrorLogClass {
	MovieErrorLogClassOnce.Do(func() {
		MovieErrorLogClass = _MovieErrorLogClass{objc.GetClass("MPMovieErrorLog")}
	})
	return MovieErrorLogClass
}

type _MovieErrorLogClass struct {
	class objc.Class
}

// An interface definition for the [MovieErrorLog] class.
type IMovieErrorLog interface {
	objectivec.IObject
}

// Data describing network resource playback failures for the associated movie player, including timestamps indicating when each failure occurred.
//
// All movie error log properties are read-only.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieErrorLog
type MovieErrorLog struct {
	objectivec.Object
}

// MovieErrorLogFrom constructs a [MovieErrorLog] from an unsafe.Pointer.
//
// Data describing network resource playback failures for the associated movie player, including timestamps indicating when each failure occurred.
func MovieErrorLogFrom(ptr unsafe.Pointer) MovieErrorLog {
	return MovieErrorLog{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MovieErrorLogClass) Alloc() MovieErrorLog {
	rv := objc.Send[MovieErrorLog](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MovieErrorLogClass) New() MovieErrorLog {
	rv := objc.Send[MovieErrorLog](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MovieErrorLog) Init() MovieErrorLog {
	rv := objc.Send[MovieErrorLog](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MovieErrorLog) Autorelease() MovieErrorLog {
	rv := objc.Send[MovieErrorLog](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMovieErrorLog creates a new MovieErrorLog instance.
func NewMovieErrorLog() MovieErrorLog {
	return getMovieErrorLogClass().New()
}




