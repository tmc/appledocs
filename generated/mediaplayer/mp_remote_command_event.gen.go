// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPRemoteCommandEvent */


/* debug [class_header]: Header for MPRemoteCommandEvent */
// The class instance for the [RemoteCommandEvent] class.
var (
	RemoteCommandEventClass     _RemoteCommandEventClass
	RemoteCommandEventClassOnce sync.Once
)

func getRemoteCommandEventClass() _RemoteCommandEventClass {
	RemoteCommandEventClassOnce.Do(func() {
		RemoteCommandEventClass = _RemoteCommandEventClass{objc.GetClass("MPRemoteCommandEvent")}
	})
	return RemoteCommandEventClass
}

type _RemoteCommandEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RemoteCommandEvent */
// An interface definition for the [RemoteCommandEvent] class.
type IRemoteCommandEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RemoteCommandEvent */
	// properties:
	Command() IMPRemoteCommand
	Timestamp() float64
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RemoteCommandEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RemoteCommandEvent */
// Alloc allocates a new instance without initialization.
func (rc _RemoteCommandEventClass) Alloc() RemoteCommandEvent {
	rv := objc.Send[RemoteCommandEvent](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RemoteCommandEventClass) New() RemoteCommandEvent {
	rv := objc.Send[RemoteCommandEvent](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RemoteCommandEvent) Init() RemoteCommandEvent {
	rv := objc.Send[RemoteCommandEvent](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RemoteCommandEvent) Autorelease() RemoteCommandEvent {
	rv := objc.Send[RemoteCommandEvent](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRemoteCommandEvent creates a new RemoteCommandEvent instance.
func NewRemoteCommandEvent() RemoteCommandEvent {
	return getRemoteCommandEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RemoteCommandEvent */
// A description of a command sent by an external media player.


// A description of a command sent by an external media player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandEvent
type RemoteCommandEvent struct {
	objectivec.Object
}

// RemoteCommandEventFrom constructs a [RemoteCommandEvent] from an unsafe.Pointer.
//
// A description of a command sent by an external media player.
func RemoteCommandEventFrom(ptr unsafe.Pointer) RemoteCommandEvent {
	return RemoteCommandEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RemoteCommandEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RemoteCommandEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RemoteCommandEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RemoteCommandEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RemoteCommandEvent */

// The command that sent the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandEvent/command
func (r_ RemoteCommandEvent) Command() IMPRemoteCommand {
	rv := objc.Send[RemoteCommand](r_.ID, objc.Sel("command"))
	return rv
}/* debug [instance_properties/getter]: command */


// The time the event occurred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommandEvent/timestamp
func (r_ RemoteCommandEvent) Timestamp() float64 {
	rv := objc.Send[float64](r_.ID, objc.Sel("timestamp"))
	return rv
}/* debug [instance_properties/getter]: timestamp */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPRemoteCommandEvent */



