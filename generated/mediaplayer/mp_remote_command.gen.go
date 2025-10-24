// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPRemoteCommand */


/* debug [class_header]: Header for MPRemoteCommand */
// The class instance for the [RemoteCommand] class.
var (
	RemoteCommandClass     _RemoteCommandClass
	RemoteCommandClassOnce sync.Once
)

func getRemoteCommandClass() _RemoteCommandClass {
	RemoteCommandClassOnce.Do(func() {
		RemoteCommandClass = _RemoteCommandClass{objc.GetClass("MPRemoteCommand")}
	})
	return RemoteCommandClass
}

type _RemoteCommandClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RemoteCommand */
// An interface definition for the [RemoteCommand] class.
type IRemoteCommand interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RemoteCommand */
	// properties:
	Enabled() bool
	SetEnabled(value bool)
	IsEnabled() bool
	SetIsEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RemoteCommand */
	// methods:
	AddTargetAction(target objc.IObject, action objc.SEL)
	AddTargetWithHandler(handler unsafe.Pointer) objc.ID
	RemoveTarget(target objc.IObject)
	RemoveTargetAction(target objc.IObject, action objc.SEL)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RemoteCommand */
// Alloc allocates a new instance without initialization.
func (rc _RemoteCommandClass) Alloc() RemoteCommand {
	rv := objc.Send[RemoteCommand](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RemoteCommandClass) New() RemoteCommand {
	rv := objc.Send[RemoteCommand](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RemoteCommand) Init() RemoteCommand {
	rv := objc.Send[RemoteCommand](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RemoteCommand) Autorelease() RemoteCommand {
	rv := objc.Send[RemoteCommand](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRemoteCommand creates a new RemoteCommand instance.
func NewRemoteCommand() RemoteCommand {
	return getRemoteCommandClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RemoteCommand */
// An object that responds to remote command events.
//
// The Media Player framework defines a standard set of remote command objects for handling media-related events. When an accessory or iOS user interface generates a remote control event, the system notifies the corresponding command object on the shared instance. That command object executes any attached handlers. To respond to a particular event, register a handler with the appropriate object. Listing 1. Registering a remote control event handler If you explicitly don’t want to enable a given command, fetch the command object and set its enabled property to . Disabling a remote command lets the system know that it shouldn’t display any related UI for that command when your app is the Now Playing app. The framework defines many subclasses to handle specific kinds of commands. Sometimes, these subclasses let you specify other information related to the command. For example, feedback commands let you specify a localized string that describes the meaning of the feedback. When supporting a particular command, be sure to look up the specific class used to handle those events.


// An object that responds to remote command events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommand
type RemoteCommand struct {
	objectivec.Object
}

// RemoteCommandFrom constructs a [RemoteCommand] from an unsafe.Pointer.
//
// An object that responds to remote command events.
func RemoteCommandFrom(ptr unsafe.Pointer) RemoteCommand {
	return RemoteCommand{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RemoteCommand *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RemoteCommand */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RemoteCommand */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RemoteCommand */

// Adds a target object to be called when an event is received.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommand/addTarget(_:action:)
func (r_ RemoteCommand) AddTargetAction(target objc.IObject, action objc.SEL) {
	objc.Send[objc.ID](r_.ID, objc.Sel("addTarget:action:"), target, action)
}/* debug [instance_methods/method]: AddTargetAction */


// Adds a block to be called when an event is received.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommand/addTarget(handler:)
func (r_ RemoteCommand) AddTargetWithHandler(handler unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("addTargetWithHandler:"), handler)
	return rv
}/* debug [instance_methods/method]: AddTargetWithHandler */


// Removes a target from the remote command object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommand/removeTarget(_:)
func (r_ RemoteCommand) RemoveTarget(target objc.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("removeTarget:"), target)
}/* debug [instance_methods/method]: RemoveTarget */


// Removes a target and action from a remote command object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommand/removeTarget(_:action:)
func (r_ RemoteCommand) RemoveTargetAction(target objc.IObject, action objc.SEL) {
	objc.Send[objc.ID](r_.ID, objc.Sel("removeTarget:action:"), target, action)
}/* debug [instance_methods/method]: RemoveTargetAction */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RemoteCommand */

// A Boolean value that indicates whether a user can interact with the displayed element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommand/isEnabled
func (r_ RemoteCommand) Enabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// A Boolean value that indicates whether a user can interact with the displayed element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPRemoteCommand/isEnabled
func (r_ RemoteCommand) SetEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setEnabled:"), value)
}/* debug [instance_properties/setter]: enabled */


// A Boolean value that indicates whether a user can interact with the displayed element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommand/isenabled
func (r_ RemoteCommand) IsEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A Boolean value that indicates whether a user can interact with the displayed element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommand/isenabled
func (r_ RemoteCommand) SetIsEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPRemoteCommand */



