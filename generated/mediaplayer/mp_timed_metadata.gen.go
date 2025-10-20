// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [TimedMetadata] class.
type ITimedMetadata interface {
	objectivec.IObject
}

// A carries time-based information within HTTP streamed media.
//
// Content providers can embed these objects when creating a stream. The properties and constants in this class let you extract the metadata as you play the stream using an object. For example, the provider of a live sports video stream could use instances to embed game scores, with timestamps, in the stream. On the client side—that is, on the user’s device—their application could employ the properties of this class to update their app’s user interface in real time during the game. A Javascript implementation of this class is also available for use by web-based applications.
//
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

// Alloc allocates a new instance without initialization.
func (tc _TimedMetadataClass) Alloc() TimedMetadata {
	rv := objc.Send[TimedMetadata](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




