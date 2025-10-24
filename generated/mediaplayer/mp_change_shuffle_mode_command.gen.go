// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPChangeShuffleModeCommand */


/* debug [class_header]: Header for MPChangeShuffleModeCommand */
// The class instance for the [ChangeShuffleModeCommand] class.
var (
	ChangeShuffleModeCommandClass     _ChangeShuffleModeCommandClass
	ChangeShuffleModeCommandClassOnce sync.Once
)

func getChangeShuffleModeCommandClass() _ChangeShuffleModeCommandClass {
	ChangeShuffleModeCommandClassOnce.Do(func() {
		ChangeShuffleModeCommandClass = _ChangeShuffleModeCommandClass{objc.GetClass("MPChangeShuffleModeCommand")}
	})
	return ChangeShuffleModeCommandClass
}

type _ChangeShuffleModeCommandClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ChangeShuffleModeCommand */
// An interface definition for the [ChangeShuffleModeCommand] class.
type IChangeShuffleModeCommand interface {
	IRemoteCommand
	
/* debug [class_interface_properties]: Properties for ChangeShuffleModeCommand */
	// properties:
	CurrentShuffleType() ShuffleType
	SetCurrentShuffleType(value ShuffleType)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ChangeShuffleModeCommand */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ChangeShuffleModeCommand */
// Alloc allocates a new instance without initialization.
func (cc _ChangeShuffleModeCommandClass) Alloc() ChangeShuffleModeCommand {
	rv := objc.Send[ChangeShuffleModeCommand](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ChangeShuffleModeCommandClass) New() ChangeShuffleModeCommand {
	rv := objc.Send[ChangeShuffleModeCommand](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ChangeShuffleModeCommand) Init() ChangeShuffleModeCommand {
	rv := objc.Send[ChangeShuffleModeCommand](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ChangeShuffleModeCommand) Autorelease() ChangeShuffleModeCommand {
	rv := objc.Send[ChangeShuffleModeCommand](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewChangeShuffleModeCommand creates a new ChangeShuffleModeCommand instance.
func NewChangeShuffleModeCommand() ChangeShuffleModeCommand {
	return getChangeShuffleModeCommandClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ChangeShuffleModeCommand */
// An object that responds to requests to change the current shuffle mode used during playback.


// An object that responds to requests to change the current shuffle mode used during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangeShuffleModeCommand
type ChangeShuffleModeCommand struct {
	RemoteCommand
}

// ChangeShuffleModeCommandFrom constructs a [ChangeShuffleModeCommand] from an unsafe.Pointer.
//
// An object that responds to requests to change the current shuffle mode used during playback.
func ChangeShuffleModeCommandFrom(ptr unsafe.Pointer) ChangeShuffleModeCommand {
	return ChangeShuffleModeCommand{
		RemoteCommand: RemoteCommandFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ChangeShuffleModeCommand *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ChangeShuffleModeCommand */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ChangeShuffleModeCommand */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ChangeShuffleModeCommand */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ChangeShuffleModeCommand */

// The current shuffle mode for a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangeShuffleModeCommand/currentShuffleType
func (c_ ChangeShuffleModeCommand) CurrentShuffleType() ShuffleType {
	rv := objc.Send[ShuffleType](c_.ID, objc.Sel("currentShuffleType"))
	return rv
}/* debug [instance_properties/getter]: currentShuffleType */


// The current shuffle mode for a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangeShuffleModeCommand/currentShuffleType
func (c_ ChangeShuffleModeCommand) SetCurrentShuffleType(value ShuffleType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCurrentShuffleType:"), value)
}/* debug [instance_properties/setter]: currentShuffleType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPChangeShuffleModeCommand */



