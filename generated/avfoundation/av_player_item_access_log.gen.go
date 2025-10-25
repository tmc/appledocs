// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVPlayerItemAccessLog */


/* debug [class_header]: Header for AVPlayerItemAccessLog */
// The class instance for the [PlayerItemAccessLog] class.
var (
	PlayerItemAccessLogClass     _PlayerItemAccessLogClass
	PlayerItemAccessLogClassOnce sync.Once
)

func getPlayerItemAccessLogClass() _PlayerItemAccessLogClass {
	PlayerItemAccessLogClassOnce.Do(func() {
		PlayerItemAccessLogClass = _PlayerItemAccessLogClass{objc.GetClass("AVPlayerItemAccessLog")}
	})
	return PlayerItemAccessLogClass
}

type _PlayerItemAccessLogClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PlayerItemAccessLog */
// An interface definition for the [PlayerItemAccessLog] class.
type IPlayerItemAccessLog interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PlayerItemAccessLog */
	// properties:
	Events() []PlayerItemAccessLogEvent
	ExtendedLogDataStringEncoding() StringEncoding /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PlayerItemAccessLog */
	// methods:
	ExtendedLogData() foundation.Data
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PlayerItemAccessLog */
// Alloc allocates a new instance without initialization.
func (pc _PlayerItemAccessLogClass) Alloc() PlayerItemAccessLog {
	rv := objc.Send[PlayerItemAccessLog](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayerItemAccessLogClass) New() PlayerItemAccessLog {
	rv := objc.Send[PlayerItemAccessLog](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerItemAccessLog) Init() PlayerItemAccessLog {
	rv := objc.Send[PlayerItemAccessLog](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerItemAccessLog) Autorelease() PlayerItemAccessLog {
	rv := objc.Send[PlayerItemAccessLog](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerItemAccessLog creates a new PlayerItemAccessLog instance.
func NewPlayerItemAccessLog() PlayerItemAccessLog {
	return getPlayerItemAccessLogClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PlayerItemAccessLog */
// An object used to retrieve the access log associated with a player item.
//
// An object accumulates key metrics about network playback and presents them as a collection of instances. Each event instance collates the data that relates to each uninterrupted period of playback.


// An object used to retrieve the access log associated with a player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLog
type PlayerItemAccessLog struct {
	objectivec.Object
}

// PlayerItemAccessLogFrom constructs a [PlayerItemAccessLog] from an unsafe.Pointer.
//
// An object used to retrieve the access log associated with a player item.
func PlayerItemAccessLogFrom(ptr unsafe.Pointer) PlayerItemAccessLog {
	return PlayerItemAccessLog{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PlayerItemAccessLog *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PlayerItemAccessLog */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PlayerItemAccessLog */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PlayerItemAccessLog */

// Returns a serialized representation of the access log in the Extended Log File Format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLog/extendedLogData()
func (p_ PlayerItemAccessLog) ExtendedLogData() foundation.Data {
	rv := objc.Send[foundation.Data](p_.ID, objc.Sel("extendedLogData"))
	return rv
}/* debug [instance_methods/method]: ExtendedLogData */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PlayerItemAccessLog */

// A chronologically ordered array of player item access log events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLog/events
func (p_ PlayerItemAccessLog) Events() []PlayerItemAccessLogEvent {
	rv := objc.Send[[]PlayerItemAccessLogEvent](p_.ID, objc.Sel("events"))
	return rv
}/* debug [instance_properties/getter]: events */


// The string encoding of the extended log data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLog/extendedLogDataStringEncoding
func (p_ PlayerItemAccessLog) ExtendedLogDataStringEncoding() StringEncoding /* not a class type */ {
	rv := objc.Send[StringEncoding](p_.ID, objc.Sel("extendedLogDataStringEncoding"))
	return rv
}/* debug [instance_properties/getter]: extendedLogDataStringEncoding */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVPlayerItemAccessLog */



