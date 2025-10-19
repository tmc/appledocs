// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVPlayerItem] class.
var aVPlayerItemClass = _AVPlayerItemClass{objc.GetClass("AVPlayerItem")}

type _AVPlayerItemClass struct {
	class objc.Class
}

// An interface definition for the [AVPlayerItem] class.
type IAVPlayerItem interface {
	objectivec.IObject
	CurrentTime() unsafe.Pointer
	RequestPlaybackRestrictionsAuthorization(completion unsafe.Pointer)
}

// An object that models the timing and presentation state of an asset during playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem

type AVPlayerItem struct {
	objectivec.Object
}

// AVPlayerItemFrom constructs a [AVPlayerItem] from an unsafe.Pointer.
//
// An object that models the timing and presentation state of an asset during playback.
func AVPlayerItemFrom(ptr unsafe.Pointer) AVPlayerItem {
	return AVPlayerItem{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (ac _AVPlayerItemClass) Alloc() AVPlayerItem {
	rv := objc.Send[AVPlayerItem](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AVPlayerItemClass) New() AVPlayerItem {
	rv := objc.Send[AVPlayerItem](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVPlayerItem) Init() AVPlayerItem {
	rv := objc.Send[AVPlayerItem](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVPlayerItem) Autorelease() AVPlayerItem {
	rv := objc.Send[AVPlayerItem](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerItem creates a new AVPlayerItem instance.
func NewAVPlayerItem() AVPlayerItem {
	return aVPlayerItemClass.New()
}


// Returns the current time of the item. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/currentTime()
func (a_ AVPlayerItem) CurrentTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("currentTime"))
	return rv
}
// Determines whether this item is subject to parental restrictions, and, if so, prompts the user to enter the restrictions passcode. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/requestPlaybackRestrictionsAuthorization(_:)
func (a_ AVPlayerItem) RequestPlaybackRestrictionsAuthorization(completion unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("requestPlaybackRestrictionsAuthorization:"), completion)
}


