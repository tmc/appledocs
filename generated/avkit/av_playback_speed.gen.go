// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [PlaybackSpeed] class.
var (
	PlaybackSpeedClass     _PlaybackSpeedClass
	PlaybackSpeedClassOnce sync.Once
)

func getPlaybackSpeedClass() _PlaybackSpeedClass {
	PlaybackSpeedClassOnce.Do(func() {
		PlaybackSpeedClass = _PlaybackSpeedClass{objc.GetClass("AVPlaybackSpeed")}
	})
	return PlaybackSpeedClass
}

type _PlaybackSpeedClass struct {
	class objc.Class
}





// An interface definition for the [PlaybackSpeed] class.
type IPlaybackSpeed interface {
	objectivec.IObject
	

	// properties:
	LocalizedName() objc.IObject /* cross-framework: NSString */
	LocalizedNumericName() objc.IObject /* cross-framework: NSString */
	Rate() float32
	SelectedSpeed() IAVPlaybackSpeed
	SetSelectedSpeed(value IAVPlaybackSpeed)
	Speeds() IAVPlaybackSpeed
	SetSpeeds(value IAVPlaybackSpeed)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (pc _PlaybackSpeedClass) Alloc() PlaybackSpeed {
	rv := objc.Send[PlaybackSpeed](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlaybackSpeedClass) New() PlaybackSpeed {
	rv := objc.Send[PlaybackSpeed](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlaybackSpeed) Init() PlaybackSpeed {
	rv := objc.Send[PlaybackSpeed](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlaybackSpeed) Autorelease() PlaybackSpeed {
	rv := objc.Send[PlaybackSpeed](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlaybackSpeed creates a new PlaybackSpeed instance.
func NewPlaybackSpeed() PlaybackSpeed {
	return getPlaybackSpeedClass().New()
}





// An object that represents a user-selectable playback speed in a playback user interface.


// An object that represents a user-selectable playback speed in a playback user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlaybackSpeed
type PlaybackSpeed struct {
	objectivec.Object
}

// PlaybackSpeedFrom constructs a [PlaybackSpeed] from an unsafe.Pointer.
//
// An object that represents a user-selectable playback speed in a playback user interface.
func PlaybackSpeedFrom(ptr unsafe.Pointer) PlaybackSpeed {
	return PlaybackSpeed{objectivec.Object{objc.ID(ptr)}}
}






// Creates a playback speed with a rate and localized name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlaybackSpeed/init(rate:localizedName:)
func NewPlaybackSpeedWithRateLocalizedName(rate float32, localizedName objc.IObject /* cross-framework: NSString */) PlaybackSpeed {
	instance := getPlaybackSpeedClass().Alloc()
	rv := objc.Send[PlaybackSpeed](instance.ID, objc.Sel("initWithRate:localizedName:"), rate, localizedName)
	rv.Autorelease()
	return rv
}












// A list of playback speeds the system uses by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlaybackSpeed/systemDefaultSpeeds
func (pc _PlaybackSpeedClass) SystemDefaultSpeeds() []PlaybackSpeed {
	rv := objc.Send[[]PlaybackSpeed](objc.ID(pc.class), objc.Sel("systemDefaultSpeeds"))
	return rv
}











// A localized name for a speed that’s suitable for display in a user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlaybackSpeed/localizedName
func (p_ PlaybackSpeed) LocalizedName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("localizedName"))
	return rv
}


// A localized numeric name for a speed that’s suitable for display in a user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlaybackSpeed/localizedNumericName
func (p_ PlaybackSpeed) LocalizedNumericName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("localizedNumericName"))
	return rv
}


// The playback rate to use when you select this speed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlaybackSpeed/rate
func (p_ PlaybackSpeed) Rate() float32 {
	rv := objc.Send[float32](p_.ID, objc.Sel("rate"))
	return rv
}


// A list of playback speeds the system uses by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVPlaybackSpeed/systemDefaultSpeeds
func (p_ PlaybackSpeed) SystemDefaultSpeeds() []PlaybackSpeed {
	rv := objc.Send[[]PlaybackSpeed](p_.ID, objc.Sel("systemDefaultSpeeds"))
	return rv
}


// The currently selected playback speed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/selectedspeed
func (p_ PlaybackSpeed) SelectedSpeed() IAVPlaybackSpeed {
	rv := objc.Send[PlaybackSpeed](p_.ID, objc.Sel("selectedSpeed"))
	return rv
}


// The currently selected playback speed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/selectedspeed
func (p_ PlaybackSpeed) SetSelectedSpeed(value IAVPlaybackSpeed) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSelectedSpeed:"), value)
}


// A list of user-selectable playback speeds to show in the playback speed control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/speeds
func (p_ PlaybackSpeed) Speeds() IAVPlaybackSpeed {
	rv := objc.Send[PlaybackSpeed](p_.ID, objc.Sel("speeds"))
	return rv
}


// A list of user-selectable playback speeds to show in the playback speed control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/speeds
func (p_ PlaybackSpeed) SetSpeeds(value IAVPlaybackSpeed) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSpeeds:"), value)
}







