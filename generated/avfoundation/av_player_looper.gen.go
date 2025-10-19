// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVPlayerLooper] class.
var aVPlayerLooperClass = _AVPlayerLooperClass{objc.GetClass("AVPlayerLooper")}

type _AVPlayerLooperClass struct {
	class objc.Class
}

// An interface definition for the [AVPlayerLooper] class.
type IAVPlayerLooper interface {
	objectivec.IObject
}

// An object that loops media content using a queue player. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper

type AVPlayerLooper struct {
	objectivec.Object
}

// AVPlayerLooperFrom constructs a [AVPlayerLooper] from an unsafe.Pointer.
//
// An object that loops media content using a queue player.
func AVPlayerLooperFrom(ptr unsafe.Pointer) AVPlayerLooper {
	return AVPlayerLooper{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (ac _AVPlayerLooperClass) Alloc() AVPlayerLooper {
	rv := objc.Send[AVPlayerLooper](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AVPlayerLooperClass) New() AVPlayerLooper {
	rv := objc.Send[AVPlayerLooper](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVPlayerLooper) Init() AVPlayerLooper {
	rv := objc.Send[AVPlayerLooper](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVPlayerLooper) Autorelease() AVPlayerLooper {
	rv := objc.Send[AVPlayerLooper](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerLooper creates a new AVPlayerLooper instance.
func NewAVPlayerLooper() AVPlayerLooper {
	return aVPlayerLooperClass.New()
}




