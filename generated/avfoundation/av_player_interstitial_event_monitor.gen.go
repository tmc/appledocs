// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVPlayerInterstitialEventMonitor */


/* debug [class_header]: Header for AVPlayerInterstitialEventMonitor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PlayerInterstitialEventMonitor */
// An interface definition for the [PlayerInterstitialEventMonitor] class.
type IPlayerInterstitialEventMonitor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PlayerInterstitialEventMonitor */
	// properties:
	CurrentEvent() IAVPlayerInterstitialEvent
	CurrentEventSkipControlLabel() objc.IObject /* cross-framework: NSString */
	CurrentEventSkippableState() PlayerInterstitialEventSkippableEventState
	Events() []PlayerInterstitialEvent
	InterstitialPlayer() IAVQueuePlayer
	PrimaryPlayer() IAVPlayer
	ReasonForWaitingToPlay() objectivec.IObject
	SetReasonForWaitingToPlay(value objectivec.IObject)
	TimeControlStatus() objectivec.IObject
	SetTimeControlStatus(value objectivec.IObject)
	TemplateItems() IAVPlayerItem
	SetTemplateItems(value IAVPlayerItem)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PlayerInterstitialEventMonitor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PlayerInterstitialEventMonitor */
// Alloc allocates a new instance without initialization.
func (pc _PlayerInterstitialEventMonitorClass) Alloc() PlayerInterstitialEventMonitor {
	rv := objc.Send[PlayerInterstitialEventMonitor](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PlayerInterstitialEventMonitor */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PlayerInterstitialEventMonitor */

// Creates an observer with a player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/init(primaryPlayer:)
func NewPlayerInterstitialEventMonitorWithPrimaryPlayer(primaryPlayer IAVPlayer) PlayerInterstitialEventMonitor {
	instance := getPlayerInterstitialEventMonitorClass().Alloc()
	rv := objc.Send[PlayerInterstitialEventMonitor](instance.ID, objc.Sel("initWithPrimaryPlayer:"), primaryPlayer)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPlayerInterstitialEventMonitorWithPrimaryPlayer */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PlayerInterstitialEventMonitor */

// A convenience initializer that creates an observer with a player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/interstitialEventMonitorWithPrimaryPlayer:
func (pc _PlayerInterstitialEventMonitorClass) InterstitialEventMonitorWithPrimaryPlayer(primaryPlayer IAVPlayer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("interstitialEventMonitorWithPrimaryPlayer:"), primaryPlayer)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=InterstitialEventMonitorWithPrimaryPlayer) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PlayerInterstitialEventMonitor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PlayerInterstitialEventMonitor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PlayerInterstitialEventMonitor */

// The current interstitial event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/currentEvent
func (p_ PlayerInterstitialEventMonitor) CurrentEvent() IAVPlayerInterstitialEvent {
	rv := objc.Send[PlayerInterstitialEvent](p_.ID, objc.Sel("currentEvent"))
	return rv
}/* debug [instance_properties/getter]: currentEvent */


// The skip control label for the currentEvent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/currentEventSkipControlLabel
func (p_ PlayerInterstitialEventMonitor) CurrentEventSkipControlLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("currentEventSkipControlLabel"))
	return rv
}/* debug [instance_properties/getter]: currentEventSkipControlLabel */


// The skippable event state for the currentEvent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/currentEventSkippableState
func (p_ PlayerInterstitialEventMonitor) CurrentEventSkippableState() PlayerInterstitialEventSkippableEventState {
	rv := objc.Send[PlayerInterstitialEventSkippableEventState](p_.ID, objc.Sel("currentEventSkippableState"))
	return rv
}/* debug [instance_properties/getter]: currentEventSkippableState */


// The schedule of interstitial events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/events
func (p_ PlayerInterstitialEventMonitor) Events() []PlayerInterstitialEvent {
	rv := objc.Send[[]PlayerInterstitialEvent](p_.ID, objc.Sel("events"))
	return rv
}/* debug [instance_properties/getter]: events */


// An object that plays interstitial content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/interstitialPlayer
func (p_ PlayerInterstitialEventMonitor) InterstitialPlayer() IAVQueuePlayer {
	rv := objc.Send[QueuePlayer](p_.ID, objc.Sel("interstitialPlayer"))
	return rv
}/* debug [instance_properties/getter]: interstitialPlayer */


// An object that plays primary content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/primaryPlayer
func (p_ PlayerInterstitialEventMonitor) PrimaryPlayer() IAVPlayer {
	rv := objc.Send[Player](p_.ID, objc.Sel("primaryPlayer"))
	return rv
}/* debug [instance_properties/getter]: primaryPlayer */


// The reason the player is currently waiting for playback to begin or resume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/reasonforwaitingtoplay
func (p_ PlayerInterstitialEventMonitor) ReasonForWaitingToPlay() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("reasonForWaitingToPlay"))
	return rv
}/* debug [instance_properties/getter]: reasonForWaitingToPlay */


// The reason the player is currently waiting for playback to begin or resume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/reasonforwaitingtoplay
func (p_ PlayerInterstitialEventMonitor) SetReasonForWaitingToPlay(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setReasonForWaitingToPlay:"), value)
}/* debug [instance_properties/setter]: reasonForWaitingToPlay */


// A value that indicates whether playback is in progress, paused indefinitely, or waiting for network conditions to improve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/timecontrolstatus-swift.property
func (p_ PlayerInterstitialEventMonitor) TimeControlStatus() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("timeControlStatus"))
	return rv
}/* debug [instance_properties/getter]: timeControlStatus */


// A value that indicates whether playback is in progress, paused indefinitely, or waiting for network conditions to improve.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayer/timecontrolstatus-swift.property
func (p_ PlayerInterstitialEventMonitor) SetTimeControlStatus(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTimeControlStatus:"), value)
}/* debug [instance_properties/setter]: timeControlStatus */


// An array of player item configurations to use as templates for player items that play interstitial content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/templateitems
func (p_ PlayerInterstitialEventMonitor) TemplateItems() IAVPlayerItem {
	rv := objc.Send[PlayerItem](p_.ID, objc.Sel("templateItems"))
	return rv
}/* debug [instance_properties/getter]: templateItems */


// An array of player item configurations to use as templates for player items that play interstitial content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/templateitems
func (p_ PlayerInterstitialEventMonitor) SetTemplateItems(value IAVPlayerItem) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTemplateItems:"), value)
}/* debug [instance_properties/setter]: templateItems */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVPlayerInterstitialEventMonitor */


