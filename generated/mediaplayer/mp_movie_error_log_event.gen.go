// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MovieErrorLogEvent] class.
var (
	MovieErrorLogEventClass     _MovieErrorLogEventClass
	MovieErrorLogEventClassOnce sync.Once
)

func getMovieErrorLogEventClass() _MovieErrorLogEventClass {
	MovieErrorLogEventClassOnce.Do(func() {
		MovieErrorLogEventClass = _MovieErrorLogEventClass{objc.GetClass("MPMovieErrorLogEvent")}
	})
	return MovieErrorLogEventClass
}

type _MovieErrorLogEventClass struct {
	class objc.Class
}

// An interface definition for the [MovieErrorLogEvent] class.
type IMovieErrorLogEvent interface {
	objectivec.IObject
}

// A single piece of information for a movie error log.
//
// All movie error log event properties are read-only. For a description of movie error logs, see .
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMovieErrorLogEvent
type MovieErrorLogEvent struct {
	objectivec.Object
}

// MovieErrorLogEventFrom constructs a [MovieErrorLogEvent] from an unsafe.Pointer.
//
// A single piece of information for a movie error log.
func MovieErrorLogEventFrom(ptr unsafe.Pointer) MovieErrorLogEvent {
	return MovieErrorLogEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MovieErrorLogEventClass) Alloc() MovieErrorLogEvent {
	rv := objc.Send[MovieErrorLogEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MovieErrorLogEventClass) New() MovieErrorLogEvent {
	rv := objc.Send[MovieErrorLogEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MovieErrorLogEvent) Init() MovieErrorLogEvent {
	rv := objc.Send[MovieErrorLogEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MovieErrorLogEvent) Autorelease() MovieErrorLogEvent {
	rv := objc.Send[MovieErrorLogEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMovieErrorLogEvent creates a new MovieErrorLogEvent instance.
func NewMovieErrorLogEvent() MovieErrorLogEvent {
	return getMovieErrorLogEventClass().New()
}




