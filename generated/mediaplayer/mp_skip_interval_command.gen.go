// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MPSkipIntervalCommand */


/* debug [class_header]: Header for MPSkipIntervalCommand */
// The class instance for the [SkipIntervalCommand] class.
var (
	SkipIntervalCommandClass     _SkipIntervalCommandClass
	SkipIntervalCommandClassOnce sync.Once
)

func getSkipIntervalCommandClass() _SkipIntervalCommandClass {
	SkipIntervalCommandClassOnce.Do(func() {
		SkipIntervalCommandClass = _SkipIntervalCommandClass{objc.GetClass("MPSkipIntervalCommand")}
	})
	return SkipIntervalCommandClass
}

type _SkipIntervalCommandClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SkipIntervalCommand */
// An interface definition for the [SkipIntervalCommand] class.
type ISkipIntervalCommand interface {
	IRemoteCommand
	
/* debug [class_interface_properties]: Properties for SkipIntervalCommand */
	// properties:
	PreferredIntervals() []foundation.Number
	SetPreferredIntervals(value []foundation.Number)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SkipIntervalCommand */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SkipIntervalCommand */
// Alloc allocates a new instance without initialization.
func (sc _SkipIntervalCommandClass) Alloc() SkipIntervalCommand {
	rv := objc.Send[SkipIntervalCommand](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SkipIntervalCommandClass) New() SkipIntervalCommand {
	rv := objc.Send[SkipIntervalCommand](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SkipIntervalCommand) Init() SkipIntervalCommand {
	rv := objc.Send[SkipIntervalCommand](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SkipIntervalCommand) Autorelease() SkipIntervalCommand {
	rv := objc.Send[SkipIntervalCommand](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkipIntervalCommand creates a new SkipIntervalCommand instance.
func NewSkipIntervalCommand() SkipIntervalCommand {
	return getSkipIntervalCommandClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SkipIntervalCommand */
// An object that defines the skip intervals for the player.
//
// You use a skip interval to move the playback of a media item, forward or backward, the indicated number of seconds. For example, a forward skip interval of 30 seconds at 2 minutes and 30 seconds into a song would immediately jump to 3 minutes into the song and continue playing. The skipped content isn’t played.


// An object that defines the skip intervals for the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPSkipIntervalCommand
type SkipIntervalCommand struct {
	RemoteCommand
}

// SkipIntervalCommandFrom constructs a [SkipIntervalCommand] from an unsafe.Pointer.
//
// An object that defines the skip intervals for the player.
func SkipIntervalCommandFrom(ptr unsafe.Pointer) SkipIntervalCommand {
	return SkipIntervalCommand{
		RemoteCommand: RemoteCommandFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SkipIntervalCommand *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SkipIntervalCommand */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SkipIntervalCommand */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SkipIntervalCommand */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SkipIntervalCommand */

// The available skip intervals, in seconds, for a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPSkipIntervalCommand/preferredIntervals
func (s_ SkipIntervalCommand) PreferredIntervals() []foundation.Number {
	rv := objc.Send[[]foundation.Number](s_.ID, objc.Sel("preferredIntervals"))
	return rv
}/* debug [instance_properties/getter]: preferredIntervals */


// The available skip intervals, in seconds, for a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPSkipIntervalCommand/preferredIntervals
func (s_ SkipIntervalCommand) SetPreferredIntervals(value []foundation.Number) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreferredIntervals:"), nsArray)
}/* debug [instance_properties/setter]: preferredIntervals */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSkipIntervalCommand */



