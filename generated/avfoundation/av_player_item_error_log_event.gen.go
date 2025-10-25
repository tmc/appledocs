// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVPlayerItemErrorLogEvent */


/* debug [class_header]: Header for AVPlayerItemErrorLogEvent */
// The class instance for the [PlayerItemErrorLogEvent] class.
var (
	PlayerItemErrorLogEventClass     _PlayerItemErrorLogEventClass
	PlayerItemErrorLogEventClassOnce sync.Once
)

func getPlayerItemErrorLogEventClass() _PlayerItemErrorLogEventClass {
	PlayerItemErrorLogEventClassOnce.Do(func() {
		PlayerItemErrorLogEventClass = _PlayerItemErrorLogEventClass{objc.GetClass("AVPlayerItemErrorLogEvent")}
	})
	return PlayerItemErrorLogEventClass
}

type _PlayerItemErrorLogEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PlayerItemErrorLogEvent */
// An interface definition for the [PlayerItemErrorLogEvent] class.
type IPlayerItemErrorLogEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PlayerItemErrorLogEvent */
	// properties:
	AllHTTPResponseHeaderFields() foundation.IDictionary
	Date() objc.IObject /* cross-framework: NSDate */
	ErrorComment() objc.IObject /* cross-framework: NSString */
	ErrorDomain() objc.IObject /* cross-framework: NSString */
	ErrorStatusCode() int
	PlaybackSessionID() objc.IObject /* cross-framework: NSString */
	ServerAddress() objc.IObject /* cross-framework: NSString */
	URI() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PlayerItemErrorLogEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PlayerItemErrorLogEvent */
// Alloc allocates a new instance without initialization.
func (pc _PlayerItemErrorLogEventClass) Alloc() PlayerItemErrorLogEvent {
	rv := objc.Send[PlayerItemErrorLogEvent](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayerItemErrorLogEventClass) New() PlayerItemErrorLogEvent {
	rv := objc.Send[PlayerItemErrorLogEvent](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerItemErrorLogEvent) Init() PlayerItemErrorLogEvent {
	rv := objc.Send[PlayerItemErrorLogEvent](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerItemErrorLogEvent) Autorelease() PlayerItemErrorLogEvent {
	rv := objc.Send[PlayerItemErrorLogEvent](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerItemErrorLogEvent creates a new PlayerItemErrorLogEvent instance.
func NewPlayerItemErrorLogEvent() PlayerItemErrorLogEvent {
	return getPlayerItemErrorLogEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PlayerItemErrorLogEvent */
// A single item in a player item’s error log.
//
// This object provides properties for accessing the data fields of each log event. Each event is a single entry in an object’s error log. These properties aren’t observable. For more information about key-value observing, see .


// A single item in a player item’s error log.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent
type PlayerItemErrorLogEvent struct {
	objectivec.Object
}

// PlayerItemErrorLogEventFrom constructs a [PlayerItemErrorLogEvent] from an unsafe.Pointer.
//
// A single item in a player item’s error log.
func PlayerItemErrorLogEventFrom(ptr unsafe.Pointer) PlayerItemErrorLogEvent {
	return PlayerItemErrorLogEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PlayerItemErrorLogEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PlayerItemErrorLogEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PlayerItemErrorLogEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PlayerItemErrorLogEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PlayerItemErrorLogEvent */

// The HTTP header fields the server returns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent/allHTTPResponseHeaderFields
func (p_ PlayerItemErrorLogEvent) AllHTTPResponseHeaderFields() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("allHTTPResponseHeaderFields"))
	return rv
}/* debug [instance_properties/getter]: allHTTPResponseHeaderFields */


// The date and time when the error occurred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent/date
func (p_ PlayerItemErrorLogEvent) Date() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("date"))
	return rv
}/* debug [instance_properties/getter]: date */


// A description of the error encountered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent/errorComment
func (p_ PlayerItemErrorLogEvent) ErrorComment() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("errorComment"))
	return rv
}/* debug [instance_properties/getter]: errorComment */


// The domain of the error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent/errorDomain
func (p_ PlayerItemErrorLogEvent) ErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("errorDomain"))
	return rv
}/* debug [instance_properties/getter]: errorDomain */


// A unique error code identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent/errorStatusCode
func (p_ PlayerItemErrorLogEvent) ErrorStatusCode() int {
	rv := objc.Send[int](p_.ID, objc.Sel("errorStatusCode"))
	return rv
}/* debug [instance_properties/getter]: errorStatusCode */


// A GUID that identifies the playback session that had an error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent/playbackSessionID
func (p_ PlayerItemErrorLogEvent) PlaybackSessionID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("playbackSessionID"))
	return rv
}/* debug [instance_properties/getter]: playbackSessionID */


// The IP address of the server that was the source of the error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent/serverAddress
func (p_ PlayerItemErrorLogEvent) ServerAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("serverAddress"))
	return rv
}/* debug [instance_properties/getter]: serverAddress */


// The URI of the playback item that had an error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemErrorLogEvent/uri
func (p_ PlayerItemErrorLogEvent) URI() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("URI"))
	return rv
}/* debug [instance_properties/getter]: URI */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVPlayerItemErrorLogEvent */



