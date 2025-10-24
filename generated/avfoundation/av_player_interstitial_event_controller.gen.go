// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corevideo"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVPlayerInterstitialEventController */


/* debug [class_header]: Header for AVPlayerInterstitialEventController */
// The class instance for the [PlayerInterstitialEventController] class.
var (
	PlayerInterstitialEventControllerClass     _PlayerInterstitialEventControllerClass
	PlayerInterstitialEventControllerClassOnce sync.Once
)

func getPlayerInterstitialEventControllerClass() _PlayerInterstitialEventControllerClass {
	PlayerInterstitialEventControllerClassOnce.Do(func() {
		PlayerInterstitialEventControllerClass = _PlayerInterstitialEventControllerClass{objc.GetClass("AVPlayerInterstitialEventController")}
	})
	return PlayerInterstitialEventControllerClass
}

type _PlayerInterstitialEventControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PlayerInterstitialEventController */
// An interface definition for the [PlayerInterstitialEventController] class.
type IPlayerInterstitialEventController interface {
	IPlayerInterstitialEventMonitor
	
/* debug [class_interface_properties]: Properties for PlayerInterstitialEventController */
	// properties:
	Events() []PlayerInterstitialEvent
	SetEvents(value []PlayerInterstitialEvent)
	LocalizedStringsBundle() foundation.Bundle
	SetLocalizedStringsBundle(value foundation.Bundle)
	LocalizedStringsTableName() objc.IObject /* cross-framework: NSString */
	SetLocalizedStringsTableName(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PlayerInterstitialEventController */
	// methods:
	CancelCurrentEventWithResumptionOffset(resumptionOffset objc.IObject /* cross-framework: Time */)
	SkipCurrentEvent()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PlayerInterstitialEventController */
// Alloc allocates a new instance without initialization.
func (pc _PlayerInterstitialEventControllerClass) Alloc() PlayerInterstitialEventController {
	rv := objc.Send[PlayerInterstitialEventController](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayerInterstitialEventControllerClass) New() PlayerInterstitialEventController {
	rv := objc.Send[PlayerInterstitialEventController](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerInterstitialEventController) Init() PlayerInterstitialEventController {
	rv := objc.Send[PlayerInterstitialEventController](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerInterstitialEventController) Autorelease() PlayerInterstitialEventController {
	rv := objc.Send[PlayerInterstitialEventController](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerInterstitialEventController creates a new PlayerInterstitialEventController instance.
func NewPlayerInterstitialEventController() PlayerInterstitialEventController {
	return getPlayerInterstitialEventControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PlayerInterstitialEventController */
// An object that schedules interstitial events for items played by the primary player.
//
// This class is a subclass of that you use to manage the schedule of interstitial events to present during playback of primary content.


// An object that schedules interstitial events for items played by the primary player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventController
type PlayerInterstitialEventController struct {
	PlayerInterstitialEventMonitor
}

// PlayerInterstitialEventControllerFrom constructs a [PlayerInterstitialEventController] from an unsafe.Pointer.
//
// An object that schedules interstitial events for items played by the primary player.
func PlayerInterstitialEventControllerFrom(ptr unsafe.Pointer) PlayerInterstitialEventController {
	return PlayerInterstitialEventController{
		PlayerInterstitialEventMonitor: PlayerInterstitialEventMonitorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PlayerInterstitialEventController */

// Creates an event controller with a player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventController/init(primaryPlayer:)
func NewPlayerInterstitialEventControllerWithPrimaryPlayer(primaryPlayer IAVPlayer) PlayerInterstitialEventController {
	instance := getPlayerInterstitialEventControllerClass().Alloc()
	rv := objc.Send[PlayerInterstitialEventController](instance.ID, objc.Sel("initWithPrimaryPlayer:"), primaryPlayer)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPlayerInterstitialEventControllerWithPrimaryPlayer */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PlayerInterstitialEventController */

// A convenience initializer that creates an event controller with a player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventController/interstitialEventControllerWithPrimaryPlayer:
func (pc _PlayerInterstitialEventControllerClass) InterstitialEventControllerWithPrimaryPlayer(primaryPlayer IAVPlayer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("interstitialEventControllerWithPrimaryPlayer:"), primaryPlayer)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=InterstitialEventControllerWithPrimaryPlayer) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PlayerInterstitialEventController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PlayerInterstitialEventController */

// Cancels the playback of all currently playing and scheduled interstitial events, and resumes playback of primary content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventController/cancelCurrentEvent(withResumptionOffset:)
func (p_ PlayerInterstitialEventController) CancelCurrentEventWithResumptionOffset(resumptionOffset objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("cancelCurrentEventWithResumptionOffset:"), resumptionOffset)
}/* debug [instance_methods/method]: CancelCurrentEventWithResumptionOffset */


// Causes the playback of the currently playing interstital event to be abandoned.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventController/skipCurrentEvent()
func (p_ PlayerInterstitialEventController) SkipCurrentEvent() {
	objc.Send[objc.ID](p_.ID, objc.Sel("skipCurrentEvent"))
}/* debug [instance_methods/method]: SkipCurrentEvent */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PlayerInterstitialEventController */

// The current schedule of interstitial events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventController/events
func (p_ PlayerInterstitialEventController) Events() []PlayerInterstitialEvent {
	rv := objc.Send[[]PlayerInterstitialEvent](p_.ID, objc.Sel("events"))
	return rv
}/* debug [instance_properties/getter]: events */


// The current schedule of interstitial events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventController/events
func (p_ PlayerInterstitialEventController) SetEvents(value []PlayerInterstitialEvent) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setEvents:"), nsArray)
}/* debug [instance_properties/setter]: events */


// The bundle that contains the localized strings to be used by the AVPlayerInterstitialEventController.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventController/localizedStringsBundle
func (p_ PlayerInterstitialEventController) LocalizedStringsBundle() foundation.Bundle {
	rv := objc.Send[foundation.Bundle](p_.ID, objc.Sel("localizedStringsBundle"))
	return rv
}/* debug [instance_properties/getter]: localizedStringsBundle */


// The bundle that contains the localized strings to be used by the AVPlayerInterstitialEventController.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventController/localizedStringsBundle
func (p_ PlayerInterstitialEventController) SetLocalizedStringsBundle(value foundation.Bundle) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLocalizedStringsBundle:"), value)
}/* debug [instance_properties/setter]: localizedStringsBundle */


// The name of the table in the bundle that contains the localized strings to be used by the AVPlayerInterstitialEventController.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventController/localizedStringsTableName
func (p_ PlayerInterstitialEventController) LocalizedStringsTableName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("localizedStringsTableName"))
	return rv
}/* debug [instance_properties/getter]: localizedStringsTableName */


// The name of the table in the bundle that contains the localized strings to be used by the AVPlayerInterstitialEventController.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventController/localizedStringsTableName
func (p_ PlayerInterstitialEventController) SetLocalizedStringsTableName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLocalizedStringsTableName:"), value)
}/* debug [instance_properties/setter]: localizedStringsTableName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVPlayerInterstitialEventController */


