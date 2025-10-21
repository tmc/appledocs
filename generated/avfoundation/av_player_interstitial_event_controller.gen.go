// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [PlayerInterstitialEventController] class.
type IPlayerInterstitialEventController interface {
	IPlayerInterstitialEventMonitor
}

// An object that schedules interstitial events for items played by the primary player.
//
// This class is a subclass of that you use to manage the schedule of interstitial events to present during playback of primary content.
//
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

// Alloc allocates a new instance without initialization.
func (pc _PlayerInterstitialEventControllerClass) Alloc() PlayerInterstitialEventController {
	rv := objc.Send[PlayerInterstitialEventController](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The name of the table in the bundle that contains the localized strings to be used by the AVPlayerInterstitialEventController.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventController/localizedStringsTableName
func (p_ PlayerInterstitialEventController) LocalizedStringsTableName() string {
	rv := objc.Send[string](p_.ID, objc.Sel("localizedStringsTableName"))
	return rv
}


// SetLocalizedStringsTableName sets the value of the localizedStringsTableName property.
// The name of the table in the bundle that contains the localized strings to be used by the AVPlayerInterstitialEventController.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventController/localizedStringsTableName
func (p_ PlayerInterstitialEventController) SetLocalizedStringsTableName(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLocalizedStringsTableName:"), objc.String(value))
}



