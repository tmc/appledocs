// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVPlayerInterstitialEventController] class.
var aVPlayerInterstitialEventControllerClass = _AVPlayerInterstitialEventControllerClass{objc.GetClass("AVPlayerInterstitialEventController")}

type _AVPlayerInterstitialEventControllerClass struct {
	class objc.Class
}

// An object that schedules interstitial events for items played by the primary player. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventController

type AVPlayerInterstitialEventController struct {
	AVPlayerInterstitialEventMonitor
}

// AVPlayerInterstitialEventControllerFrom constructs a [AVPlayerInterstitialEventController] from an unsafe.Pointer.
//
// An object that schedules interstitial events for items played by the primary player.
func AVPlayerInterstitialEventControllerFrom(ptr unsafe.Pointer) AVPlayerInterstitialEventController {
	return AVPlayerInterstitialEventController{
		AVPlayerInterstitialEventMonitor: AVPlayerInterstitialEventMonitorFrom(ptr),
	}
}



