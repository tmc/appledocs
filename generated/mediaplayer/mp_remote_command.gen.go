// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [RemoteCommand] class.
type IRemoteCommand interface {
	objectivec.IObject
	// properties:
	IsEnabled() bool
	SetIsEnabled(value bool)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (rc _RemoteCommandClass) Alloc() RemoteCommand {
	rv := objc.Send[RemoteCommand](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A Boolean value that indicates whether a user can interact with the displayed element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommand/isenabled
func (r_ RemoteCommand) IsEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean value that indicates whether a user can interact with the displayed element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommand/isenabled
func (r_ RemoteCommand) SetIsEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsEnabled:"), value)
}



