// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PlayerInterstitialEvent] class.
var (
	PlayerInterstitialEventClass     _PlayerInterstitialEventClass
	PlayerInterstitialEventClassOnce sync.Once
)

func getPlayerInterstitialEventClass() _PlayerInterstitialEventClass {
	PlayerInterstitialEventClassOnce.Do(func() {
		PlayerInterstitialEventClass = _PlayerInterstitialEventClass{objc.GetClass("AVPlayerInterstitialEvent")}
	})
	return PlayerInterstitialEventClass
}

type _PlayerInterstitialEventClass struct {
	class objc.Class
}

// An interface definition for the [PlayerInterstitialEvent] class.
type IPlayerInterstitialEvent interface {
	objectivec.IObject
}

// An object that provides instructions for how a player presents interstitial content.
//
// An interstitial event defines a or , on the timeline of its , at which playback of interstitial content begins. It specifies the alternative interstitial content to play as an array of one or more template player items. The system uses the configuration of the event’s to build new player item instances to present the interstitial content. Use to observe the scheduling and progress of interstitial events. If your app requires specifying the schedule of interstitial events, use instead.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent
type PlayerInterstitialEvent struct {
	objectivec.Object
}

// PlayerInterstitialEventFrom constructs a [PlayerInterstitialEvent] from an unsafe.Pointer.
//
// An object that provides instructions for how a player presents interstitial content.
func PlayerInterstitialEventFrom(ptr unsafe.Pointer) PlayerInterstitialEvent {
	return PlayerInterstitialEvent{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PlayerInterstitialEventClass) Alloc() PlayerInterstitialEvent {
	rv := objc.Send[PlayerInterstitialEvent](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PlayerInterstitialEventClass) New() PlayerInterstitialEvent {
	rv := objc.Send[PlayerInterstitialEvent](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerInterstitialEvent) Init() PlayerInterstitialEvent {
	rv := objc.Send[PlayerInterstitialEvent](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerInterstitialEvent) Autorelease() PlayerInterstitialEvent {
	rv := objc.Send[PlayerInterstitialEvent](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerInterstitialEvent creates a new PlayerInterstitialEvent instance.
func NewPlayerInterstitialEvent() PlayerInterstitialEvent {
	return getPlayerInterstitialEventClass().New()
}


// Attributes of the event that the vendor or app defines.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/userDefinedAttributes
func (p_ PlayerInterstitialEvent) UserDefinedAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("userDefinedAttributes"))
	return rv
}



