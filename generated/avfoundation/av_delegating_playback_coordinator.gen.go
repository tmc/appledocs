// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [DelegatingPlaybackCoordinator] class.
type IDelegatingPlaybackCoordinator interface {
	IPlaybackCoordinator
	

	// properties:
	CurrentItemIdentifier() foundation.foundation.INSString
	PlaybackControlDelegate() unsafe.Pointer


	

	// methods:
	CoordinateRateChangeToRateOptions(rate float32, options DelegatingPlaybackCoordinatorRateChangeOptions)
	CoordinateSeekToTimeOptions(time objectivec.IObject, options DelegatingPlaybackCoordinatorSeekOptions)
	ReapplyCurrentItemStateToPlaybackControlDelegate()
	TransitionToItemWithIdentifierProposingInitialTimingBasedOnTimebase(itemIdentifier foundation.foundation.INSString, snapshotTimebase TimebaseRef /* not a class type */)


}





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






// Creates a playback coordinator for a custom playback object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator/init(playbackControlDelegate:)
func NewDelegatingPlaybackCoordinatorWithPlaybackControlDelegate(playbackControlDelegate unsafe.Pointer) DelegatingPlaybackCoordinator {
	instance := getDelegatingPlaybackCoordinatorClass().Alloc()
	rv := objc.Send[DelegatingPlaybackCoordinator](instance.ID, objc.Sel("initWithPlaybackControlDelegate:"), playbackControlDelegate)
	rv.Autorelease()
	return rv
}

















// Coordinates a rate change across all participants, waiting for others to become ready, if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator/coordinateRateChange(to:options:)
func (d_ DelegatingPlaybackCoordinator) CoordinateRateChangeToRateOptions(rate float32, options DelegatingPlaybackCoordinatorRateChangeOptions) {
	objc.Send[objc.ID](d_.ID, objc.Sel("coordinateRateChangeToRate:options:"), rate, options)
}


// Coordinates a seek to the specified time for all connected participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator/coordinateSeek(to:options:)
func (d_ DelegatingPlaybackCoordinator) CoordinateSeekToTimeOptions(time objectivec.IObject, options DelegatingPlaybackCoordinatorSeekOptions) {
	objc.Send[objc.ID](d_.ID, objc.Sel("coordinateSeekToTime:options:"), time, options)
}


// Tells the coordinator to reissue current play state commands to synchronize the current item to the state of other participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator/reapplyCurrentItemStateToPlaybackControlDelegate()
func (d_ DelegatingPlaybackCoordinator) ReapplyCurrentItemStateToPlaybackControlDelegate() {
	objc.Send[objc.ID](d_.ID, objc.Sel("reapplyCurrentItemStateToPlaybackControlDelegate"))
}


// Tells the coordinator to transition to a new item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator/transitionToItem(withIdentifier:proposingInitialTimingBasedOn:)
func (d_ DelegatingPlaybackCoordinator) TransitionToItemWithIdentifierProposingInitialTimingBasedOnTimebase(itemIdentifier foundation.foundation.INSString, snapshotTimebase TimebaseRef /* not a class type */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("transitionToItemWithIdentifier:proposingInitialTimingBasedOnTimebase:"), itemIdentifier, snapshotTimebase)
}







// An identifier of the current item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator/currentItemIdentifier
func (d_ DelegatingPlaybackCoordinator) CurrentItemIdentifier() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("currentItemIdentifier"))
	return rv
}


// The delegate object for the playback coordinator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator/playbackControlDelegate
func (d_ DelegatingPlaybackCoordinator) PlaybackControlDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("playbackControlDelegate"))
	return rv
}







