// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVPlayerInterstitialEventController] class.
var (
	aVPlayerInterstitialEventControllerClass     _AVPlayerInterstitialEventControllerClass
	aVPlayerInterstitialEventControllerClassOnce sync.Once
)

func getAVPlayerInterstitialEventControllerClass() _AVPlayerInterstitialEventControllerClass {
	aVPlayerInterstitialEventControllerClassOnce.Do(func() {
		aVPlayerInterstitialEventControllerClass = _AVPlayerInterstitialEventControllerClass{objc.GetClass("AVPlayerInterstitialEventController")}
	})
	return aVPlayerInterstitialEventControllerClass
}

type _AVPlayerInterstitialEventControllerClass struct {
	class objc.Class
}

// An interface definition for the [AVPlayerInterstitialEventController] class.
type IAVPlayerInterstitialEventController interface {
	IAVPlayerInterstitialEventMonitor
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

// Alloc allocates a new instance without initialization.
func (ac _AVPlayerInterstitialEventControllerClass) Alloc() AVPlayerInterstitialEventController {
	rv := objc.Send[AVPlayerInterstitialEventController](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVPlayerInterstitialEventControllerClass) New() AVPlayerInterstitialEventController {
	rv := objc.Send[AVPlayerInterstitialEventController](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVPlayerInterstitialEventController) Init() AVPlayerInterstitialEventController {
	rv := objc.Send[AVPlayerInterstitialEventController](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVPlayerInterstitialEventController) Autorelease() AVPlayerInterstitialEventController {
	rv := objc.Send[AVPlayerInterstitialEventController](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerInterstitialEventController creates a new AVPlayerInterstitialEventController instance.
func NewAVPlayerInterstitialEventController() AVPlayerInterstitialEventController {
	return getAVPlayerInterstitialEventControllerClass().New()
}




