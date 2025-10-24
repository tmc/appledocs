// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPChangeLanguageOptionCommandEvent */


/* debug [class_header]: Header for MPChangeLanguageOptionCommandEvent */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ChangeLanguageOptionCommandEvent */
// An interface definition for the [ChangeLanguageOptionCommandEvent] class.
type IChangeLanguageOptionCommandEvent interface {
	IRemoteCommandEvent
	
/* debug [class_interface_properties]: Properties for ChangeLanguageOptionCommandEvent */
	// properties:
	LanguageOption() IMPNowPlayingInfoLanguageOption
	Setting() ChangeLanguageOptionSetting
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ChangeLanguageOptionCommandEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ChangeLanguageOptionCommandEvent */
// Alloc allocates a new instance without initialization.
func (cc _ChangeLanguageOptionCommandEventClass) Alloc() ChangeLanguageOptionCommandEvent {
	rv := objc.Send[ChangeLanguageOptionCommandEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ChangeLanguageOptionCommandEvent */
// An event requesting a change in the language option.


// An event requesting a change in the language option.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ChangeLanguageOptionCommandEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ChangeLanguageOptionCommandEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ChangeLanguageOptionCommandEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ChangeLanguageOptionCommandEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ChangeLanguageOptionCommandEvent */

// The requested language option to change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangeLanguageOptionCommandEvent/languageOption
func (c_ ChangeLanguageOptionCommandEvent) LanguageOption() IMPNowPlayingInfoLanguageOption {
	rv := objc.Send[NowPlayingInfoLanguageOption](c_.ID, objc.Sel("languageOption"))
	return rv
}/* debug [instance_properties/getter]: languageOption */


// The extent of the language setting change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangeLanguageOptionCommandEvent/setting
func (c_ ChangeLanguageOptionCommandEvent) Setting() ChangeLanguageOptionSetting {
	rv := objc.Send[ChangeLanguageOptionSetting](c_.ID, objc.Sel("setting"))
	return rv
}/* debug [instance_properties/getter]: setting */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPChangeLanguageOptionCommandEvent */



