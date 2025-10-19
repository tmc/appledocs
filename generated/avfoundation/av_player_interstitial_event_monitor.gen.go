// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVPlayerInterstitialEventMonitor] class.
var aVPlayerInterstitialEventMonitorClass = _AVPlayerInterstitialEventMonitorClass{objc.GetClass("AVPlayerInterstitialEventMonitor")}

type _AVPlayerInterstitialEventMonitorClass struct {
	class objc.Class
}

// An object that monitors the scheduling and progress of interstitial events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor

type AVPlayerInterstitialEventMonitor struct {
	objectivec.Object
}

// AVPlayerInterstitialEventMonitorFrom constructs a [AVPlayerInterstitialEventMonitor] from an unsafe.Pointer.
//
// An object that monitors the scheduling and progress of interstitial events.
func AVPlayerInterstitialEventMonitorFrom(ptr unsafe.Pointer) AVPlayerInterstitialEventMonitor {
	return AVPlayerInterstitialEventMonitor{objectivec.Object{objc.ID(ptr)}}
}



