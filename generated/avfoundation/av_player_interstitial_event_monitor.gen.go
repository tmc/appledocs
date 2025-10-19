// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVPlayerInterstitialEventMonitor] class.
var (
	aVPlayerInterstitialEventMonitorClass     _AVPlayerInterstitialEventMonitorClass
	aVPlayerInterstitialEventMonitorClassOnce sync.Once
)

func getAVPlayerInterstitialEventMonitorClass() _AVPlayerInterstitialEventMonitorClass {
	aVPlayerInterstitialEventMonitorClassOnce.Do(func() {
		aVPlayerInterstitialEventMonitorClass = _AVPlayerInterstitialEventMonitorClass{objc.GetClass("AVPlayerInterstitialEventMonitor")}
	})
	return aVPlayerInterstitialEventMonitorClass
}

type _AVPlayerInterstitialEventMonitorClass struct {
	class objc.Class
}

// An interface definition for the [AVPlayerInterstitialEventMonitor] class.
type IAVPlayerInterstitialEventMonitor interface {
	objectivec.IObject
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

// Alloc allocates a new instance without initialization.
func (ac _AVPlayerInterstitialEventMonitorClass) Alloc() AVPlayerInterstitialEventMonitor {
	rv := objc.Send[AVPlayerInterstitialEventMonitor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVPlayerInterstitialEventMonitorClass) New() AVPlayerInterstitialEventMonitor {
	rv := objc.Send[AVPlayerInterstitialEventMonitor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVPlayerInterstitialEventMonitor) Init() AVPlayerInterstitialEventMonitor {
	rv := objc.Send[AVPlayerInterstitialEventMonitor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVPlayerInterstitialEventMonitor) Autorelease() AVPlayerInterstitialEventMonitor {
	rv := objc.Send[AVPlayerInterstitialEventMonitor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerInterstitialEventMonitor creates a new AVPlayerInterstitialEventMonitor instance.
func NewAVPlayerInterstitialEventMonitor() AVPlayerInterstitialEventMonitor {
	return getAVPlayerInterstitialEventMonitorClass().New()
}




