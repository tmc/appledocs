// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	ReasonForWaitingToPlay() unsafe.Pointer
	SetReasonForWaitingToPlay(value unsafe.Pointer)
	TimeControlStatus() unsafe.Pointer
	SetTimeControlStatus(value unsafe.Pointer)
	TemplateItems() AVPlayerItem
	SetTemplateItems(value IAVPlayerItem)
	CurrentEvent() AVPlayerInterstitialEvent
	SetCurrentEvent(value IAVPlayerInterstitialEvent)
	CurrentEventSkipControlLabel() string
	SetCurrentEventSkipControlLabel(value string)
	CurrentEventSkippableState() unsafe.Pointer
	SetCurrentEventSkippableState(value unsafe.Pointer)
	Events() AVPlayerInterstitialEvent
	SetEvents(value IAVPlayerInterstitialEvent)
	InterstitialPlayer() AVQueuePlayer
	SetInterstitialPlayer(value IAVQueuePlayer)
	PrimaryPlayer() AVPlayer
	SetPrimaryPlayer(value IAVPlayer)
}

// An object that monitors the scheduling and progress of interstitial events.
//
// This object monitors interstitial events that exist within the content of the primary items, such as events defined by an HLS media playlist, and also events managed by an object. You can access the schedule of interstitial events through the property. When it’s time to present an interstitial event, the system suspends playback of the primary item and changes its player’s to with a value of . When the system suspends primary playback, it creates player items based on the event’s to play interstitial content. The interstitial player temporarily assumes the primary player’s output configuration, such as routing its visual output to player layers that reference the primary player. After the interstitial player finishes playback, or its current item otherwise becomes , playback of primary content resumes.


// An object that monitors the scheduling and progress of interstitial events.
//
// [Full Topic]
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



// The reason the player is currently waiting for playback to begin or resume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/reasonforwaitingtoplay
func (p_ PlayerInterstitialEventMonitor) ReasonForWaitingToPlay() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("reasonForWaitingToPlay"))
	return rv
}


// The reason the player is currently waiting for playback to begin or resume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/reasonforwaitingtoplay
func (p_ PlayerInterstitialEventMonitor) SetReasonForWaitingToPlay(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setReasonForWaitingToPlay:"), value)
}


// A value that indicates whether playback is in progress, paused indefinitely, or waiting for network conditions to improve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/timecontrolstatus-swift.property
func (p_ PlayerInterstitialEventMonitor) TimeControlStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("timeControlStatus"))
	return rv
}


// A value that indicates whether playback is in progress, paused indefinitely, or waiting for network conditions to improve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/timecontrolstatus-swift.property
func (p_ PlayerInterstitialEventMonitor) SetTimeControlStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTimeControlStatus:"), value)
}


// An array of player item configurations to use as templates for player items that play interstitial content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/templateitems
func (p_ PlayerInterstitialEventMonitor) TemplateItems() AVPlayerItem {
	rv := objc.Send[AVPlayerItem](p_.ID, objc.Sel("templateItems"))
	return rv
}


// An array of player item configurations to use as templates for player items that play interstitial content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/templateitems
func (p_ PlayerInterstitialEventMonitor) SetTemplateItems(value IAVPlayerItem) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTemplateItems:"), value)
}


// The current interstitial event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventmonitor/currentevent
func (p_ PlayerInterstitialEventMonitor) CurrentEvent() AVPlayerInterstitialEvent {
	rv := objc.Send[AVPlayerInterstitialEvent](p_.ID, objc.Sel("currentEvent"))
	return rv
}


// The current interstitial event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventmonitor/currentevent
func (p_ PlayerInterstitialEventMonitor) SetCurrentEvent(value IAVPlayerInterstitialEvent) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentEvent:"), value)
}


// The skip control label for the currentEvent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventmonitor/currenteventskipcontrollabel
func (p_ PlayerInterstitialEventMonitor) CurrentEventSkipControlLabel() string {
	rv := objc.Send[string](p_.ID, objc.Sel("currentEventSkipControlLabel"))
	return rv
}


// The skip control label for the currentEvent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventmonitor/currenteventskipcontrollabel
func (p_ PlayerInterstitialEventMonitor) SetCurrentEventSkipControlLabel(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentEventSkipControlLabel:"), objc.String(value))
}


// The skippable event state for the currentEvent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventmonitor/currenteventskippablestate
func (p_ PlayerInterstitialEventMonitor) CurrentEventSkippableState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("currentEventSkippableState"))
	return rv
}


// The skippable event state for the currentEvent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventmonitor/currenteventskippablestate
func (p_ PlayerInterstitialEventMonitor) SetCurrentEventSkippableState(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentEventSkippableState:"), value)
}


// The schedule of interstitial events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventmonitor/events
func (p_ PlayerInterstitialEventMonitor) Events() AVPlayerInterstitialEvent {
	rv := objc.Send[AVPlayerInterstitialEvent](p_.ID, objc.Sel("events"))
	return rv
}


// The schedule of interstitial events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventmonitor/events
func (p_ PlayerInterstitialEventMonitor) SetEvents(value IAVPlayerInterstitialEvent) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEvents:"), value)
}


// An object that plays interstitial content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventmonitor/interstitialplayer
func (p_ PlayerInterstitialEventMonitor) InterstitialPlayer() AVQueuePlayer {
	rv := objc.Send[AVQueuePlayer](p_.ID, objc.Sel("interstitialPlayer"))
	return rv
}


// An object that plays interstitial content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventmonitor/interstitialplayer
func (p_ PlayerInterstitialEventMonitor) SetInterstitialPlayer(value IAVQueuePlayer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInterstitialPlayer:"), value)
}


// An object that plays primary content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventmonitor/primaryplayer
func (p_ PlayerInterstitialEventMonitor) PrimaryPlayer() AVPlayer {
	rv := objc.Send[AVPlayer](p_.ID, objc.Sel("primaryPlayer"))
	return rv
}


// An object that plays primary content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialeventmonitor/primaryplayer
func (p_ PlayerInterstitialEventMonitor) SetPrimaryPlayer(value IAVPlayer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPrimaryPlayer:"), value)
}



