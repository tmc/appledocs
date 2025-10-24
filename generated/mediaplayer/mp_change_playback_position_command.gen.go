// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPChangePlaybackPositionCommand */


/* debug [class_header]: Header for MPChangePlaybackPositionCommand */
// The class instance for the [ChangePlaybackPositionCommand] class.
var (
	ChangePlaybackPositionCommandClass     _ChangePlaybackPositionCommandClass
	ChangePlaybackPositionCommandClassOnce sync.Once
)

func getChangePlaybackPositionCommandClass() _ChangePlaybackPositionCommandClass {
	ChangePlaybackPositionCommandClassOnce.Do(func() {
		ChangePlaybackPositionCommandClass = _ChangePlaybackPositionCommandClass{objc.GetClass("MPChangePlaybackPositionCommand")}
	})
	return ChangePlaybackPositionCommandClass
}

type _ChangePlaybackPositionCommandClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ChangePlaybackPositionCommand */
// An interface definition for the [ChangePlaybackPositionCommand] class.
type IChangePlaybackPositionCommand interface {
	IRemoteCommand
	
/* debug [class_interface_properties]: Properties for ChangePlaybackPositionCommand */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ChangePlaybackPositionCommand */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ChangePlaybackPositionCommand */
// Alloc allocates a new instance without initialization.
func (cc _ChangePlaybackPositionCommandClass) Alloc() ChangePlaybackPositionCommand {
	rv := objc.Send[ChangePlaybackPositionCommand](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ChangePlaybackPositionCommandClass) New() ChangePlaybackPositionCommand {
	rv := objc.Send[ChangePlaybackPositionCommand](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ChangePlaybackPositionCommand) Init() ChangePlaybackPositionCommand {
	rv := objc.Send[ChangePlaybackPositionCommand](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ChangePlaybackPositionCommand) Autorelease() ChangePlaybackPositionCommand {
	rv := objc.Send[ChangePlaybackPositionCommand](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewChangePlaybackPositionCommand creates a new ChangePlaybackPositionCommand instance.
func NewChangePlaybackPositionCommand() ChangePlaybackPositionCommand {
	return getChangePlaybackPositionCommandClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ChangePlaybackPositionCommand */
// An object that responds to requests to change the current playback position of the playing item.


// An object that responds to requests to change the current playback position of the playing item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangePlaybackPositionCommand
type ChangePlaybackPositionCommand struct {
	RemoteCommand
}

// ChangePlaybackPositionCommandFrom constructs a [ChangePlaybackPositionCommand] from an unsafe.Pointer.
//
// An object that responds to requests to change the current playback position of the playing item.
func ChangePlaybackPositionCommandFrom(ptr unsafe.Pointer) ChangePlaybackPositionCommand {
	return ChangePlaybackPositionCommand{
		RemoteCommand: RemoteCommandFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ChangePlaybackPositionCommand *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ChangePlaybackPositionCommand */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ChangePlaybackPositionCommand */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ChangePlaybackPositionCommand */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ChangePlaybackPositionCommand */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPChangePlaybackPositionCommand */



