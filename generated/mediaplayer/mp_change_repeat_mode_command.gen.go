// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPChangeRepeatModeCommand */


/* debug [class_header]: Header for MPChangeRepeatModeCommand */
// The class instance for the [ChangeRepeatModeCommand] class.
var (
	ChangeRepeatModeCommandClass     _ChangeRepeatModeCommandClass
	ChangeRepeatModeCommandClassOnce sync.Once
)

func getChangeRepeatModeCommandClass() _ChangeRepeatModeCommandClass {
	ChangeRepeatModeCommandClassOnce.Do(func() {
		ChangeRepeatModeCommandClass = _ChangeRepeatModeCommandClass{objc.GetClass("MPChangeRepeatModeCommand")}
	})
	return ChangeRepeatModeCommandClass
}

type _ChangeRepeatModeCommandClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ChangeRepeatModeCommand */
// An interface definition for the [ChangeRepeatModeCommand] class.
type IChangeRepeatModeCommand interface {
	IRemoteCommand
	
/* debug [class_interface_properties]: Properties for ChangeRepeatModeCommand */
	// properties:
	CurrentRepeatType() RepeatType
	SetCurrentRepeatType(value RepeatType)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ChangeRepeatModeCommand */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ChangeRepeatModeCommand */
// Alloc allocates a new instance without initialization.
func (cc _ChangeRepeatModeCommandClass) Alloc() ChangeRepeatModeCommand {
	rv := objc.Send[ChangeRepeatModeCommand](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ChangeRepeatModeCommandClass) New() ChangeRepeatModeCommand {
	rv := objc.Send[ChangeRepeatModeCommand](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ChangeRepeatModeCommand) Init() ChangeRepeatModeCommand {
	rv := objc.Send[ChangeRepeatModeCommand](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ChangeRepeatModeCommand) Autorelease() ChangeRepeatModeCommand {
	rv := objc.Send[ChangeRepeatModeCommand](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewChangeRepeatModeCommand creates a new ChangeRepeatModeCommand instance.
func NewChangeRepeatModeCommand() ChangeRepeatModeCommand {
	return getChangeRepeatModeCommandClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ChangeRepeatModeCommand */
// An object that responds to requests to change the current repeat mode used during playback.


// An object that responds to requests to change the current repeat mode used during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangeRepeatModeCommand
type ChangeRepeatModeCommand struct {
	RemoteCommand
}

// ChangeRepeatModeCommandFrom constructs a [ChangeRepeatModeCommand] from an unsafe.Pointer.
//
// An object that responds to requests to change the current repeat mode used during playback.
func ChangeRepeatModeCommandFrom(ptr unsafe.Pointer) ChangeRepeatModeCommand {
	return ChangeRepeatModeCommand{
		RemoteCommand: RemoteCommandFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ChangeRepeatModeCommand *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ChangeRepeatModeCommand */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ChangeRepeatModeCommand */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ChangeRepeatModeCommand */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ChangeRepeatModeCommand */

// The current repeat option for a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangeRepeatModeCommand/currentRepeatType
func (c_ ChangeRepeatModeCommand) CurrentRepeatType() RepeatType {
	rv := objc.Send[RepeatType](c_.ID, objc.Sel("currentRepeatType"))
	return rv
}/* debug [instance_properties/getter]: currentRepeatType */


// The current repeat option for a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPChangeRepeatModeCommand/currentRepeatType
func (c_ ChangeRepeatModeCommand) SetCurrentRepeatType(value RepeatType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCurrentRepeatType:"), value)
}/* debug [instance_properties/setter]: currentRepeatType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPChangeRepeatModeCommand */



