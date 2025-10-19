// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVPlayerInterstitialEvent] class.
var aVPlayerInterstitialEventClass = _AVPlayerInterstitialEventClass{objc.GetClass("AVPlayerInterstitialEvent")}

type _AVPlayerInterstitialEventClass struct {
	class objc.Class
}

// An object that provides instructions for how a player presents interstitial content. [Full Topic]
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



