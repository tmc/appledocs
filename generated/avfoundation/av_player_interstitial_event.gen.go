// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVPlayerInterstitialEvent] class.
var (
	aVPlayerInterstitialEventClass     _AVPlayerInterstitialEventClass
	aVPlayerInterstitialEventClassOnce sync.Once
)

func getAVPlayerInterstitialEventClass() _AVPlayerInterstitialEventClass {
	aVPlayerInterstitialEventClassOnce.Do(func() {
		aVPlayerInterstitialEventClass = _AVPlayerInterstitialEventClass{objc.GetClass("AVPlayerInterstitialEvent")}
	})
	return aVPlayerInterstitialEventClass
}

type _AVPlayerInterstitialEventClass struct {
	class objc.Class
}

// An interface definition for the [AVPlayerInterstitialEvent] class.
type IAVPlayerInterstitialEvent interface {
	objectivec.IObject
}

// An object that provides instructions for how a player presents interstitial content.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent
type AVPlayerInterstitialEvent struct {
	objectivec.Object
}

// AVPlayerInterstitialEventFrom constructs a [AVPlayerInterstitialEvent] from an unsafe.Pointer.
//
// An object that provides instructions for how a player presents interstitial content.
func AVPlayerInterstitialEventFrom(ptr unsafe.Pointer) AVPlayerInterstitialEvent {
	return AVPlayerInterstitialEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AVPlayerInterstitialEventClass) Alloc() AVPlayerInterstitialEvent {
	rv := objc.Send[AVPlayerInterstitialEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVPlayerInterstitialEventClass) New() AVPlayerInterstitialEvent {
	rv := objc.Send[AVPlayerInterstitialEvent](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVPlayerInterstitialEvent) Init() AVPlayerInterstitialEvent {
	rv := objc.Send[AVPlayerInterstitialEvent](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVPlayerInterstitialEvent) Autorelease() AVPlayerInterstitialEvent {
	rv := objc.Send[AVPlayerInterstitialEvent](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerInterstitialEvent creates a new AVPlayerInterstitialEvent instance.
func NewAVPlayerInterstitialEvent() AVPlayerInterstitialEvent {
	return getAVPlayerInterstitialEventClass().New()
}




