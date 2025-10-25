// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVDelegatingPlaybackCoordinator */


/* debug [class_header]: Header for AVDelegatingPlaybackCoordinator */
// The class instance for the [DelegatingPlaybackCoordinator] class.
var (
	DelegatingPlaybackCoordinatorClass     _DelegatingPlaybackCoordinatorClass
	DelegatingPlaybackCoordinatorClassOnce sync.Once
)

func getDelegatingPlaybackCoordinatorClass() _DelegatingPlaybackCoordinatorClass {
	DelegatingPlaybackCoordinatorClassOnce.Do(func() {
		DelegatingPlaybackCoordinatorClass = _DelegatingPlaybackCoordinatorClass{objc.GetClass("AVDelegatingPlaybackCoordinator")}
	})
	return DelegatingPlaybackCoordinatorClass
}

type _DelegatingPlaybackCoordinatorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DelegatingPlaybackCoordinator */
// An interface definition for the [DelegatingPlaybackCoordinator] class.
type IDelegatingPlaybackCoordinator interface {
	IPlaybackCoordinator
	
/* debug [class_interface_properties]: Properties for DelegatingPlaybackCoordinator */
	// properties:
	CurrentItemIdentifier() objc.IObject /* cross-framework: NSString */
	PlaybackControlDelegate() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DelegatingPlaybackCoordinator */
	// methods:
	CoordinateRateChangeToRateOptions(rate float32, options DelegatingPlaybackCoordinatorRateChangeOptions)
	CoordinateSeekToTimeOptions(time objc.IObject /* cross-framework: Time */, options DelegatingPlaybackCoordinatorSeekOptions)
	ReapplyCurrentItemStateToPlaybackControlDelegate()
	TransitionToItemWithIdentifierProposingInitialTimingBasedOnTimebase(itemIdentifier objc.IObject /* cross-framework: NSString */, snapshotTimebase TimebaseRef /* not a class type */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DelegatingPlaybackCoordinator */
// Alloc allocates a new instance without initialization.
func (dc _DelegatingPlaybackCoordinatorClass) Alloc() DelegatingPlaybackCoordinator {
	rv := objc.Send[DelegatingPlaybackCoordinator](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DelegatingPlaybackCoordinatorClass) New() DelegatingPlaybackCoordinator {
	rv := objc.Send[DelegatingPlaybackCoordinator](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DelegatingPlaybackCoordinator) Init() DelegatingPlaybackCoordinator {
	rv := objc.Send[DelegatingPlaybackCoordinator](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DelegatingPlaybackCoordinator) Autorelease() DelegatingPlaybackCoordinator {
	rv := objc.Send[DelegatingPlaybackCoordinator](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDelegatingPlaybackCoordinator creates a new DelegatingPlaybackCoordinator instance.
func NewDelegatingPlaybackCoordinator() DelegatingPlaybackCoordinator {
	return getDelegatingPlaybackCoordinatorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DelegatingPlaybackCoordinator */
// A playback coordinator subclass that coordinates the playback of custom player objects in a connected group.
//
// This object coordinates the state of custom player objects, such as those that render media using and , or that play audio using . Adopt the protocol so that your app responds to playback commands from the coordinator. The commands provide the details of a requested state change so you can control your player object accordingly.


// A playback coordinator subclass that coordinates the playback of custom player objects in a connected group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator
type DelegatingPlaybackCoordinator struct {
	PlaybackCoordinator
}

// DelegatingPlaybackCoordinatorFrom constructs a [DelegatingPlaybackCoordinator] from an unsafe.Pointer.
//
// A playback coordinator subclass that coordinates the playback of custom player objects in a connected group.
func DelegatingPlaybackCoordinatorFrom(ptr unsafe.Pointer) DelegatingPlaybackCoordinator {
	return DelegatingPlaybackCoordinator{
		PlaybackCoordinator: PlaybackCoordinatorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DelegatingPlaybackCoordinator */

// Creates a playback coordinator for a custom playback object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator/init(playbackControlDelegate:)
func NewDelegatingPlaybackCoordinatorWithPlaybackControlDelegate(playbackControlDelegate unsafe.Pointer) DelegatingPlaybackCoordinator {
	instance := getDelegatingPlaybackCoordinatorClass().Alloc()
	rv := objc.Send[DelegatingPlaybackCoordinator](instance.ID, objc.Sel("initWithPlaybackControlDelegate:"), playbackControlDelegate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDelegatingPlaybackCoordinatorWithPlaybackControlDelegate */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DelegatingPlaybackCoordinator */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DelegatingPlaybackCoordinator */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DelegatingPlaybackCoordinator */

// Coordinates a rate change across all participants, waiting for others to become ready, if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator/coordinateRateChange(to:options:)
func (d_ DelegatingPlaybackCoordinator) CoordinateRateChangeToRateOptions(rate float32, options DelegatingPlaybackCoordinatorRateChangeOptions) {
	objc.Send[objc.ID](d_.ID, objc.Sel("coordinateRateChangeToRate:options:"), rate, options)
}/* debug [instance_methods/method]: CoordinateRateChangeToRateOptions */


// Coordinates a seek to the specified time for all connected participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator/coordinateSeek(to:options:)
func (d_ DelegatingPlaybackCoordinator) CoordinateSeekToTimeOptions(time objc.IObject /* cross-framework: Time */, options DelegatingPlaybackCoordinatorSeekOptions) {
	objc.Send[objc.ID](d_.ID, objc.Sel("coordinateSeekToTime:options:"), time, options)
}/* debug [instance_methods/method]: CoordinateSeekToTimeOptions */


// Tells the coordinator to reissue current play state commands to synchronize the current item to the state of other participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator/reapplyCurrentItemStateToPlaybackControlDelegate()
func (d_ DelegatingPlaybackCoordinator) ReapplyCurrentItemStateToPlaybackControlDelegate() {
	objc.Send[objc.ID](d_.ID, objc.Sel("reapplyCurrentItemStateToPlaybackControlDelegate"))
}/* debug [instance_methods/method]: ReapplyCurrentItemStateToPlaybackControlDelegate */


// Tells the coordinator to transition to a new item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator/transitionToItem(withIdentifier:proposingInitialTimingBasedOn:)
func (d_ DelegatingPlaybackCoordinator) TransitionToItemWithIdentifierProposingInitialTimingBasedOnTimebase(itemIdentifier objc.IObject /* cross-framework: NSString */, snapshotTimebase TimebaseRef /* not a class type */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("transitionToItemWithIdentifier:proposingInitialTimingBasedOnTimebase:"), itemIdentifier, snapshotTimebase)
}/* debug [instance_methods/method]: TransitionToItemWithIdentifierProposingInitialTimingBasedOnTimebase */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DelegatingPlaybackCoordinator */

// An identifier of the current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator/currentItemIdentifier
func (d_ DelegatingPlaybackCoordinator) CurrentItemIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("currentItemIdentifier"))
	return rv
}/* debug [instance_properties/getter]: currentItemIdentifier */


// The delegate object for the playback coordinator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator/playbackControlDelegate
func (d_ DelegatingPlaybackCoordinator) PlaybackControlDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("playbackControlDelegate"))
	return rv
}/* debug [instance_properties/getter]: playbackControlDelegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVDelegatingPlaybackCoordinator */


