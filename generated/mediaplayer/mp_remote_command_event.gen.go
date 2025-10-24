// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [RemoteCommandEvent] class.
type IRemoteCommandEvent interface {
	objectivec.IObject
	// properties:
	Command() IMPRemoteCommand
	SetCommand(value IMPRemoteCommand)
	Timestamp() float64
	SetTimestamp(value float64)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (rc _RemoteCommandEventClass) Alloc() RemoteCommandEvent {
	rv := objc.Send[RemoteCommandEvent](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The command that sent the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandevent/command
func (r_ RemoteCommandEvent) Command() IMPRemoteCommand {
	rv := objc.Send[RemoteCommand](r_.ID, objc.Sel("command"))
	return rv
}


// The command that sent the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandevent/command
func (r_ RemoteCommandEvent) SetCommand(value IMPRemoteCommand) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setCommand:"), value)
}


// The time the event occurred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandevent/timestamp
func (r_ RemoteCommandEvent) Timestamp() float64 {
	rv := objc.Send[float64](r_.ID, objc.Sel("timestamp"))
	return rv
}


// The time the event occurred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommandevent/timestamp
func (r_ RemoteCommandEvent) SetTimestamp(value float64) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setTimestamp:"), value)
}



