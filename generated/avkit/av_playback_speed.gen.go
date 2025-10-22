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
	LocalizedName() string
	SetLocalizedName(value string)
	LocalizedNumericName() string
	SetLocalizedNumericName(value string)
	Rate() float32
	SetRate(value float32)
	SelectedSpeed() AVPlaybackSpeed
	SetSelectedSpeed(value IAVPlaybackSpeed)
	Speeds() AVPlaybackSpeed
	SetSpeeds(value IAVPlaybackSpeed)
}

// An object that represents a user-selectable playback speed in a playback user interface.
//
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

// Alloc allocates a new instance without initialization.
func (pc _PlaybackSpeedClass) Alloc() PlaybackSpeed {
	rv := objc.Send[PlaybackSpeed](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A localized name for a speed that’s suitable for display in a user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplaybackspeed/localizedname
func (p_ PlaybackSpeed) LocalizedName() string {
	rv := objc.Send[string](p_.ID, objc.Sel("localizedName"))
	return rv
}


// SetLocalizedName sets the value of the localizedName property.
// A localized name for a speed that’s suitable for display in a user interface.

//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplaybackspeed/localizedname
func (p_ PlaybackSpeed) SetLocalizedName(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLocalizedName:"), objc.String(value))
}

// A localized numeric name for a speed that’s suitable for display in a user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplaybackspeed/localizednumericname
func (p_ PlaybackSpeed) LocalizedNumericName() string {
	rv := objc.Send[string](p_.ID, objc.Sel("localizedNumericName"))
	return rv
}


// SetLocalizedNumericName sets the value of the localizedNumericName property.
// A localized numeric name for a speed that’s suitable for display in a user interface.

//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplaybackspeed/localizednumericname
func (p_ PlaybackSpeed) SetLocalizedNumericName(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLocalizedNumericName:"), objc.String(value))
}

// The playback rate to use when you select this speed.
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplaybackspeed/rate
func (p_ PlaybackSpeed) Rate() float32 {
	rv := objc.Send[float32](p_.ID, objc.Sel("rate"))
	return rv
}


// SetRate sets the value of the rate property.
// The playback rate to use when you select this speed.

//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplaybackspeed/rate
func (p_ PlaybackSpeed) SetRate(value float32) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRate:"), value)
}

// The currently selected playback speed.
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/selectedspeed
func (p_ PlaybackSpeed) SelectedSpeed() AVPlaybackSpeed {
	rv := objc.Send[AVPlaybackSpeed](p_.ID, objc.Sel("selectedSpeed"))
	return rv
}


// SetSelectedSpeed sets the value of the selectedSpeed property.
// The currently selected playback speed.

//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/selectedspeed
func (p_ PlaybackSpeed) SetSelectedSpeed(value IAVPlaybackSpeed) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSelectedSpeed:"), value)
}

// A list of user-selectable playback speeds to show in the playback speed control.
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/speeds
func (p_ PlaybackSpeed) Speeds() AVPlaybackSpeed {
	rv := objc.Send[AVPlaybackSpeed](p_.ID, objc.Sel("speeds"))
	return rv
}


// SetSpeeds sets the value of the speeds property.
// A list of user-selectable playback speeds to show in the playback speed control.

//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerview/speeds
func (p_ PlaybackSpeed) SetSpeeds(value IAVPlaybackSpeed) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSpeeds:"), value)
}



