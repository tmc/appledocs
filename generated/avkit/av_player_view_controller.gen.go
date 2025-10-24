// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [PlayerViewController] class.
var (
	PlayerViewControllerClass     _PlayerViewControllerClass
	PlayerViewControllerClassOnce sync.Once
)

func getPlayerViewControllerClass() _PlayerViewControllerClass {
	PlayerViewControllerClassOnce.Do(func() {
		PlayerViewControllerClass = _PlayerViewControllerClass{objc.GetClass("AVPlayerViewController")}
	})
	return PlayerViewControllerClass
}

type _PlayerViewControllerClass struct {
	class objc.Class
}





// An interface definition for the [PlayerViewController] class.
type IPlayerViewController interface {
	IViewController
	

	// properties:
	ExperienceController() ExperienceController /* not a class type */
	SetExperienceController(value ExperienceController /* not a class type */)
	IsReadyForDisplay() bool
	SetIsReadyForDisplay(value bool)
	IsSkipBackwardEnabled() bool
	SetIsSkipBackwardEnabled(value bool)
	IsSkipForwardEnabled() bool
	SetIsSkipForwardEnabled(value bool)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (pc _PlayerViewControllerClass) Alloc() PlayerViewController {
	rv := objc.Send[PlayerViewController](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayerViewControllerClass) New() PlayerViewController {
	rv := objc.Send[PlayerViewController](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerViewController) Init() PlayerViewController {
	rv := objc.Send[PlayerViewController](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerViewController) Autorelease() PlayerViewController {
	rv := objc.Send[PlayerViewController](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerViewController creates a new PlayerViewController instance.
func NewPlayerViewController() PlayerViewController {
	return getPlayerViewControllerClass().New()
}





// A view controller that displays content from a player and presents a native user interface to control playback.
//
// A player view controller makes it simple to add media playback capabilities to your app that match the styling and features of the native system players. Using this object also means that your app automatically adopts the new features and styling of future operating system releases.


// A view controller that displays content from a player and presents a native user interface to control playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController
type PlayerViewController struct {
	ViewController
}

// PlayerViewControllerFrom constructs a [PlayerViewController] from an unsafe.Pointer.
//
// A view controller that displays content from a player and presents a native user interface to control playback.
func PlayerViewControllerFrom(ptr unsafe.Pointer) PlayerViewController {
	return PlayerViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController/mediaCharacteristicsForSupportedCustomMediaSelectionSchemes
func (pc _PlayerViewControllerClass) MediaCharacteristicsForSupportedCustomMediaSelectionSchemes() []string {
	rv := objc.Send[[]string](objc.ID(pc.class), objc.Sel("mediaCharacteristicsForSupportedCustomMediaSelectionSchemes"))
	return rv
}











// The experience controller for this view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewcontroller/experiencecontroller
func (p_ PlayerViewController) ExperienceController() ExperienceController /* not a class type */ {
	rv := objc.Send[ExperienceController](p_.ID, objc.Sel("experienceController"))
	return rv
}


// The experience controller for this view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewcontroller/experiencecontroller
func (p_ PlayerViewController) SetExperienceController(value ExperienceController /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setExperienceController:"), value)
}


// A Boolean value that indicates whether the player item’s first video frame is ready for display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewcontroller/isreadyfordisplay
func (p_ PlayerViewController) IsReadyForDisplay() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isReadyForDisplay"))
	return rv
}


// A Boolean value that indicates whether the player item’s first video frame is ready for display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewcontroller/isreadyfordisplay
func (p_ PlayerViewController) SetIsReadyForDisplay(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsReadyForDisplay:"), value)
}


// A Boolean value that indicates whether backward-skipping is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewcontroller/isskipbackwardenabled
func (p_ PlayerViewController) IsSkipBackwardEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isSkipBackwardEnabled"))
	return rv
}


// A Boolean value that indicates whether backward-skipping is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewcontroller/isskipbackwardenabled
func (p_ PlayerViewController) SetIsSkipBackwardEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsSkipBackwardEnabled:"), value)
}


// A Boolean value that indicates whether forward-skipping is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewcontroller/isskipforwardenabled
func (p_ PlayerViewController) IsSkipForwardEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isSkipForwardEnabled"))
	return rv
}


// A Boolean value that indicates whether forward-skipping is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewcontroller/isskipforwardenabled
func (p_ PlayerViewController) SetIsSkipForwardEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsSkipForwardEnabled:"), value)
}







