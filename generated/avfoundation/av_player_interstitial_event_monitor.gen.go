// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PlayerInterstitialEventMonitor] class.
var (
	PlayerInterstitialEventMonitorClass     _PlayerInterstitialEventMonitorClass
	PlayerInterstitialEventMonitorClassOnce sync.Once
)

func getPlayerInterstitialEventMonitorClass() _PlayerInterstitialEventMonitorClass {
	PlayerInterstitialEventMonitorClassOnce.Do(func() {
		PlayerInterstitialEventMonitorClass = _PlayerInterstitialEventMonitorClass{objc.GetClass("AVPlayerInterstitialEventMonitor")}
	})
	return PlayerInterstitialEventMonitorClass
}

type _PlayerInterstitialEventMonitorClass struct {
	class objc.Class
}

// An interface definition for the [PlayerInterstitialEventMonitor] class.
type IPlayerInterstitialEventMonitor interface {
	objectivec.IObject
}

// An object that monitors the scheduling and progress of interstitial events.
//
// This object monitors interstitial events that exist within the content of the primary items, such as events defined by an HLS media playlist, and also events managed by an object. You can access the schedule of interstitial events through the property. When it’s time to present an interstitial event, the system suspends playback of the primary item and changes its player’s to with a value of . When the system suspends primary playback, it creates player items based on the event’s to play interstitial content. The interstitial player temporarily assumes the primary player’s output configuration, such as routing its visual output to player layers that reference the primary player. After the interstitial player finishes playback, or its current item otherwise becomes , playback of primary content resumes.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor
type PlayerInterstitialEventMonitor struct {
	objectivec.Object
}

// PlayerInterstitialEventMonitorFrom constructs a [PlayerInterstitialEventMonitor] from an unsafe.Pointer.
//
// An object that monitors the scheduling and progress of interstitial events.
func PlayerInterstitialEventMonitorFrom(ptr unsafe.Pointer) PlayerInterstitialEventMonitor {
	return PlayerInterstitialEventMonitor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PlayerInterstitialEventMonitorClass) Alloc() PlayerInterstitialEventMonitor {
	rv := objc.Send[PlayerInterstitialEventMonitor](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PlayerInterstitialEventMonitorClass) New() PlayerInterstitialEventMonitor {
	rv := objc.Send[PlayerInterstitialEventMonitor](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerInterstitialEventMonitor) Init() PlayerInterstitialEventMonitor {
	rv := objc.Send[PlayerInterstitialEventMonitor](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerInterstitialEventMonitor) Autorelease() PlayerInterstitialEventMonitor {
	rv := objc.Send[PlayerInterstitialEventMonitor](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerInterstitialEventMonitor creates a new PlayerInterstitialEventMonitor instance.
func NewPlayerInterstitialEventMonitor() PlayerInterstitialEventMonitor {
	return getPlayerInterstitialEventMonitorClass().New()
}




