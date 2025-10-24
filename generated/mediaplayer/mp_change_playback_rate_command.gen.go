// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MPChangePlaybackRateCommand */


/* debug [class_header]: Header for MPChangePlaybackRateCommand */
// The class instance for the [ChangePlaybackRateCommand] class.
var (
	ChangePlaybackRateCommandClass     _ChangePlaybackRateCommandClass
	ChangePlaybackRateCommandClassOnce sync.Once
)

func getChangePlaybackRateCommandClass() _ChangePlaybackRateCommandClass {
	ChangePlaybackRateCommandClassOnce.Do(func() {
		ChangePlaybackRateCommandClass = _ChangePlaybackRateCommandClass{objc.GetClass("MPChangePlaybackRateCommand")}
	})
	return ChangePlaybackRateCommandClass
}

type _ChangePlaybackRateCommandClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ChangePlaybackRateCommand */
// An interface definition for the [ChangePlaybackRateCommand] class.
type IChangePlaybackRateCommand interface {
	IRemoteCommand
	
/* debug [class_interface_properties]: Properties for ChangePlaybackRateCommand */
	// properties:
	SupportedPlaybackRates() []foundation.Number
	SetSupportedPlaybackRates(value []foundation.Number)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ChangePlaybackRateCommand */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ChangePlaybackRateCommand */
// Alloc allocates a new instance without initialization.
func (cc _ChangePlaybackRateCommandClass) Alloc() ChangePlaybackRateCommand {
	rv := objc.Send[ChangePlaybackRateCommand](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ChangePlaybackRateCommandClass) New() ChangePlaybackRateCommand {
	rv := objc.Send[ChangePlaybackRateCommand](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ChangePlaybackRateCommand) Init() ChangePlaybackRateCommand {
	rv := objc.Send[ChangePlaybackRateCommand](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ChangePlaybackRateCommand) Autorelease() ChangePlaybackRateCommand {
	rv := objc.Send[ChangePlaybackRateCommand](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewChangePlaybackRateCommand creates a new ChangePlaybackRateCommand instance.
func NewChangePlaybackRateCommand() ChangePlaybackRateCommand {
	return getChangePlaybackRateCommandClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ChangePlaybackRateCommand */
// An object that responds to requests to change the playback rate of the playing item.
//
// Apps can change the current playback rate of a media item to one of the supported rates defined by the property.


// An object that responds to requests to change the playback rate of the playing item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangePlaybackRateCommand
type ChangePlaybackRateCommand struct {
	RemoteCommand
}

// ChangePlaybackRateCommandFrom constructs a [ChangePlaybackRateCommand] from an unsafe.Pointer.
//
// An object that responds to requests to change the playback rate of the playing item.
func ChangePlaybackRateCommandFrom(ptr unsafe.Pointer) ChangePlaybackRateCommand {
	return ChangePlaybackRateCommand{
		RemoteCommand: RemoteCommandFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ChangePlaybackRateCommand *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ChangePlaybackRateCommand */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ChangePlaybackRateCommand */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ChangePlaybackRateCommand */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ChangePlaybackRateCommand */

// The supported playback rates for a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangePlaybackRateCommand/supportedPlaybackRates
func (c_ ChangePlaybackRateCommand) SupportedPlaybackRates() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("supportedPlaybackRates"))
	return rv
}/* debug [instance_properties/getter]: supportedPlaybackRates */


// The supported playback rates for a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangePlaybackRateCommand/supportedPlaybackRates
func (c_ ChangePlaybackRateCommand) SetSupportedPlaybackRates(value []foundation.Number) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportedPlaybackRates:"), nsArray)
}/* debug [instance_properties/setter]: supportedPlaybackRates */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPChangePlaybackRateCommand */



