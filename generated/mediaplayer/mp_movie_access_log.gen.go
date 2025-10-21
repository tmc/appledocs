// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MovieAccessLog] class.
var (
	MovieAccessLogClass     _MovieAccessLogClass
	MovieAccessLogClassOnce sync.Once
)

func getMovieAccessLogClass() _MovieAccessLogClass {
	MovieAccessLogClassOnce.Do(func() {
		MovieAccessLogClass = _MovieAccessLogClass{objc.GetClass("MPMovieAccessLog")}
	})
	return MovieAccessLogClass
}

type _MovieAccessLogClass struct {
	class objc.Class
}

// An interface definition for the [MovieAccessLog] class.
type IMovieAccessLog interface {
	objectivec.IObject
}

// Key metrics about network playback for an associated movie player that’s playing streamed content.
//
// The log presents these metrics as a collection of instances and also makes it available in a textual format. A movie access log describes one uninterrupted period of playback. A movie player (an instance of the class) can access this log from its property. All movie access log properties are read-only.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieAccessLog
type MovieAccessLog struct {
	objectivec.Object
}

// MovieAccessLogFrom constructs a [MovieAccessLog] from an unsafe.Pointer.
//
// Key metrics about network playback for an associated movie player that’s playing streamed content.
func MovieAccessLogFrom(ptr unsafe.Pointer) MovieAccessLog {
	return MovieAccessLog{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MovieAccessLogClass) Alloc() MovieAccessLog {
	rv := objc.Send[MovieAccessLog](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MovieAccessLogClass) New() MovieAccessLog {
	rv := objc.Send[MovieAccessLog](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MovieAccessLog) Init() MovieAccessLog {
	rv := objc.Send[MovieAccessLog](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MovieAccessLog) Autorelease() MovieAccessLog {
	rv := objc.Send[MovieAccessLog](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMovieAccessLog creates a new MovieAccessLog instance.
func NewMovieAccessLog() MovieAccessLog {
	return getMovieAccessLogClass().New()
}




