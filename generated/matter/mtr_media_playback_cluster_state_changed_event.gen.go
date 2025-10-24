// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRMediaPlaybackClusterStateChangedEvent */


/* debug [class_header]: Header for MTRMediaPlaybackClusterStateChangedEvent */
// The class instance for the [MTRMediaPlaybackClusterStateChangedEvent] class.
var (
	MTRMediaPlaybackClusterStateChangedEventClass     _MTRMediaPlaybackClusterStateChangedEventClass
	MTRMediaPlaybackClusterStateChangedEventClassOnce sync.Once
)

func getMTRMediaPlaybackClusterStateChangedEventClass() _MTRMediaPlaybackClusterStateChangedEventClass {
	MTRMediaPlaybackClusterStateChangedEventClassOnce.Do(func() {
		MTRMediaPlaybackClusterStateChangedEventClass = _MTRMediaPlaybackClusterStateChangedEventClass{objc.GetClass("MTRMediaPlaybackClusterStateChangedEvent")}
	})
	return MTRMediaPlaybackClusterStateChangedEventClass
}

type _MTRMediaPlaybackClusterStateChangedEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRMediaPlaybackClusterStateChangedEvent */
// An interface definition for the [MTRMediaPlaybackClusterStateChangedEvent] class.
type IMTRMediaPlaybackClusterStateChangedEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRMediaPlaybackClusterStateChangedEvent */
	// properties:
	AudioAdvanceUnmuted() objc.IObject /* cross-framework: NSNumber */
	SetAudioAdvanceUnmuted(value objc.IObject /* cross-framework: NSNumber */)
	CurrentState() objc.IObject /* cross-framework: NSNumber */
	SetCurrentState(value objc.IObject /* cross-framework: NSNumber */)
	Data() foundation.Data
	SetData(value foundation.Data)
	Duration() objc.IObject /* cross-framework: NSNumber */
	SetDuration(value objc.IObject /* cross-framework: NSNumber */)
	PlaybackSpeed() objc.IObject /* cross-framework: NSNumber */
	SetPlaybackSpeed(value objc.IObject /* cross-framework: NSNumber */)
	SampledPosition() objc.IObject /* cross-framework: MTRMediaPlaybackClusterPlaybackPositionStruct */
	SetSampledPosition(value objc.IObject /* cross-framework: MTRMediaPlaybackClusterPlaybackPositionStruct */)
	SeekRangeEnd() objc.IObject /* cross-framework: NSNumber */
	SetSeekRangeEnd(value objc.IObject /* cross-framework: NSNumber */)
	SeekRangeStart() objc.IObject /* cross-framework: NSNumber */
	SetSeekRangeStart(value objc.IObject /* cross-framework: NSNumber */)
	StartTime() objc.IObject /* cross-framework: NSNumber */
	SetStartTime(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRMediaPlaybackClusterStateChangedEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRMediaPlaybackClusterStateChangedEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRMediaPlaybackClusterStateChangedEventClass) Alloc() MTRMediaPlaybackClusterStateChangedEvent {
	rv := objc.Send[MTRMediaPlaybackClusterStateChangedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRMediaPlaybackClusterStateChangedEventClass) New() MTRMediaPlaybackClusterStateChangedEvent {
	rv := objc.Send[MTRMediaPlaybackClusterStateChangedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaPlaybackClusterStateChangedEvent) Init() MTRMediaPlaybackClusterStateChangedEvent {
	rv := objc.Send[MTRMediaPlaybackClusterStateChangedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaPlaybackClusterStateChangedEvent) Autorelease() MTRMediaPlaybackClusterStateChangedEvent {
	rv := objc.Send[MTRMediaPlaybackClusterStateChangedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaPlaybackClusterStateChangedEvent creates a new MTRMediaPlaybackClusterStateChangedEvent instance.
func NewMTRMediaPlaybackClusterStateChangedEvent() MTRMediaPlaybackClusterStateChangedEvent {
	return getMTRMediaPlaybackClusterStateChangedEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRMediaPlaybackClusterStateChangedEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent
type MTRMediaPlaybackClusterStateChangedEvent struct {
	objectivec.Object
}

// MTRMediaPlaybackClusterStateChangedEventFrom constructs a [MTRMediaPlaybackClusterStateChangedEvent] from an unsafe.Pointer.
func MTRMediaPlaybackClusterStateChangedEventFrom(ptr unsafe.Pointer) MTRMediaPlaybackClusterStateChangedEvent {
	return MTRMediaPlaybackClusterStateChangedEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRMediaPlaybackClusterStateChangedEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRMediaPlaybackClusterStateChangedEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRMediaPlaybackClusterStateChangedEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRMediaPlaybackClusterStateChangedEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRMediaPlaybackClusterStateChangedEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/audioAdvanceUnmuted
func (m_ MTRMediaPlaybackClusterStateChangedEvent) AudioAdvanceUnmuted() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("audioAdvanceUnmuted"))
	return rv
}/* debug [instance_properties/getter]: audioAdvanceUnmuted */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/audioAdvanceUnmuted
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetAudioAdvanceUnmuted(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAudioAdvanceUnmuted:"), value)
}/* debug [instance_properties/setter]: audioAdvanceUnmuted */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/currentState
func (m_ MTRMediaPlaybackClusterStateChangedEvent) CurrentState() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("currentState"))
	return rv
}/* debug [instance_properties/getter]: currentState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStateChangedEvent/currentState
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetCurrentState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrentState:"), value)
}/* debug [instance_properties/setter]: currentState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterstatechangedevent/data
func (m_ MTRMediaPlaybackClusterStateChangedEvent) Data() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterstatechangedevent/data
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetData(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}/* debug [instance_properties/setter]: data */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterstatechangedevent/duration
func (m_ MTRMediaPlaybackClusterStateChangedEvent) Duration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterstatechangedevent/duration
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}/* debug [instance_properties/setter]: duration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterstatechangedevent/playbackspeed
func (m_ MTRMediaPlaybackClusterStateChangedEvent) PlaybackSpeed() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("playbackSpeed"))
	return rv
}/* debug [instance_properties/getter]: playbackSpeed */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterstatechangedevent/playbackspeed
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetPlaybackSpeed(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlaybackSpeed:"), value)
}/* debug [instance_properties/setter]: playbackSpeed */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterstatechangedevent/sampledposition
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SampledPosition() objc.IObject /* cross-framework: MTRMediaPlaybackClusterPlaybackPositionStruct */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("sampledPosition"))
	return rv
}/* debug [instance_properties/getter]: sampledPosition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterstatechangedevent/sampledposition
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetSampledPosition(value objc.IObject /* cross-framework: MTRMediaPlaybackClusterPlaybackPositionStruct */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSampledPosition:"), value)
}/* debug [instance_properties/setter]: sampledPosition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterstatechangedevent/seekrangeend
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SeekRangeEnd() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("seekRangeEnd"))
	return rv
}/* debug [instance_properties/getter]: seekRangeEnd */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterstatechangedevent/seekrangeend
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetSeekRangeEnd(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSeekRangeEnd:"), value)
}/* debug [instance_properties/setter]: seekRangeEnd */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterstatechangedevent/seekrangestart
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SeekRangeStart() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("seekRangeStart"))
	return rv
}/* debug [instance_properties/getter]: seekRangeStart */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterstatechangedevent/seekrangestart
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetSeekRangeStart(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSeekRangeStart:"), value)
}/* debug [instance_properties/setter]: seekRangeStart */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterstatechangedevent/starttime
func (m_ MTRMediaPlaybackClusterStateChangedEvent) StartTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startTime"))
	return rv
}/* debug [instance_properties/getter]: startTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterstatechangedevent/starttime
func (m_ MTRMediaPlaybackClusterStateChangedEvent) SetStartTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:"), value)
}/* debug [instance_properties/setter]: startTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRMediaPlaybackClusterStateChangedEvent */



