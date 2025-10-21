// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [SampleCursor] class.
var (
	SampleCursorClass     _SampleCursorClass
	SampleCursorClassOnce sync.Once
)

func getSampleCursorClass() _SampleCursorClass {
	SampleCursorClassOnce.Do(func() {
		SampleCursorClass = _SampleCursorClass{objc.GetClass("AVSampleCursor")}
	})
	return SampleCursorClass
}

type _SampleCursorClass struct {
	class objc.Class
}

// An interface definition for the [SampleCursor] class.
type ISampleCursor interface {
	objectivec.IObject
}

// An object that provides information about the media sample at the cursor’s current position.
//
// You position a sample cursor at a specific media sample in a sequence of samples contained in a higher-level object, like an . You can move it to a new position in that sequence either backwards or forwards, either in decode order or in presentation order. You can also request moving it according to a count of samples or a delta in time. Use a sample cursor to get information about the media sample such as its duration, timestamps, dependency information, and so on. You can also use them to synchronously to perform I/O in order to load media data of one or more media samples into memory.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor
type SampleCursor struct {
	objectivec.Object
}

// SampleCursorFrom constructs a [SampleCursor] from an unsafe.Pointer.
//
// An object that provides information about the media sample at the cursor’s current position.
func SampleCursorFrom(ptr unsafe.Pointer) SampleCursor {
	return SampleCursor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SampleCursorClass) Alloc() SampleCursor {
	rv := objc.Send[SampleCursor](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SampleCursorClass) New() SampleCursor {
	rv := objc.Send[SampleCursor](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SampleCursor) Init() SampleCursor {
	rv := objc.Send[SampleCursor](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SampleCursor) Autorelease() SampleCursor {
	rv := objc.Send[SampleCursor](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSampleCursor creates a new SampleCursor instance.
func NewSampleCursor() SampleCursor {
	return getSampleCursorClass().New()
}




