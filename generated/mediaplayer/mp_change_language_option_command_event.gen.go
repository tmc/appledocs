// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ChangeLanguageOptionCommandEvent] class.
var (
	ChangeLanguageOptionCommandEventClass     _ChangeLanguageOptionCommandEventClass
	ChangeLanguageOptionCommandEventClassOnce sync.Once
)

func getChangeLanguageOptionCommandEventClass() _ChangeLanguageOptionCommandEventClass {
	ChangeLanguageOptionCommandEventClassOnce.Do(func() {
		ChangeLanguageOptionCommandEventClass = _ChangeLanguageOptionCommandEventClass{objc.GetClass("MPChangeLanguageOptionCommandEvent")}
	})
	return ChangeLanguageOptionCommandEventClass
}

type _ChangeLanguageOptionCommandEventClass struct {
	class objc.Class
}

// An interface definition for the [ChangeLanguageOptionCommandEvent] class.
type IChangeLanguageOptionCommandEvent interface {
	IRemoteCommandEvent
}

// An event requesting a change in the language option.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangeLanguageOptionCommandEvent
type ChangeLanguageOptionCommandEvent struct {
	RemoteCommandEvent
}

// ChangeLanguageOptionCommandEventFrom constructs a [ChangeLanguageOptionCommandEvent] from an unsafe.Pointer.
//
// An event requesting a change in the language option.
func ChangeLanguageOptionCommandEventFrom(ptr unsafe.Pointer) ChangeLanguageOptionCommandEvent {
	return ChangeLanguageOptionCommandEvent{
		RemoteCommandEvent: RemoteCommandEventFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _ChangeLanguageOptionCommandEventClass) Alloc() ChangeLanguageOptionCommandEvent {
	rv := objc.Send[ChangeLanguageOptionCommandEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ChangeLanguageOptionCommandEventClass) New() ChangeLanguageOptionCommandEvent {
	rv := objc.Send[ChangeLanguageOptionCommandEvent](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ChangeLanguageOptionCommandEvent) Init() ChangeLanguageOptionCommandEvent {
	rv := objc.Send[ChangeLanguageOptionCommandEvent](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ChangeLanguageOptionCommandEvent) Autorelease() ChangeLanguageOptionCommandEvent {
	rv := objc.Send[ChangeLanguageOptionCommandEvent](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewChangeLanguageOptionCommandEvent creates a new ChangeLanguageOptionCommandEvent instance.
func NewChangeLanguageOptionCommandEvent() ChangeLanguageOptionCommandEvent {
	return getChangeLanguageOptionCommandEventClass().New()
}


// The requested language option to change.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangelanguageoptioncommandevent/languageoption
func (c_ ChangeLanguageOptionCommandEvent) LanguageOption() MPNowPlayingInfoLanguageOption {
	rv := objc.Send[MPNowPlayingInfoLanguageOption](c_.ID, objc.Sel("languageOption"))
	return rv
}


// SetLanguageOption sets the value of the languageOption property.
// The requested language option to change.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangelanguageoptioncommandevent/languageoption
func (c_ ChangeLanguageOptionCommandEvent) SetLanguageOption(value IMPNowPlayingInfoLanguageOption) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLanguageOption:"), value)
}

// The extent of the language setting change.
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangelanguageoptioncommandevent/setting
func (c_ ChangeLanguageOptionCommandEvent) Setting() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("setting"))
	return rv
}


// SetSetting sets the value of the setting property.
// The extent of the language setting change.

//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpchangelanguageoptioncommandevent/setting
func (c_ ChangeLanguageOptionCommandEvent) SetSetting(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSetting:"), value)
}



