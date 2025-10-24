// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVPlayerLooper */


/* debug [class_header]: Header for AVPlayerLooper */
// The class instance for the [PlayerLooper] class.
var (
	PlayerLooperClass     _PlayerLooperClass
	PlayerLooperClassOnce sync.Once
)

func getPlayerLooperClass() _PlayerLooperClass {
	PlayerLooperClassOnce.Do(func() {
		PlayerLooperClass = _PlayerLooperClass{objc.GetClass("AVPlayerLooper")}
	})
	return PlayerLooperClass
}

type _PlayerLooperClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PlayerLooper */
// An interface definition for the [PlayerLooper] class.
type IPlayerLooper interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PlayerLooper */
	// properties:
	Error() Error
	LoopCount() int
	LoopingPlayerItems() []PlayerItem
	Status() PlayerLooperStatus
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PlayerLooper */
	// methods:
	DisableLooping()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PlayerLooper */
// Alloc allocates a new instance without initialization.
func (pc _PlayerLooperClass) Alloc() PlayerLooper {
	rv := objc.Send[PlayerLooper](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayerLooperClass) New() PlayerLooper {
	rv := objc.Send[PlayerLooper](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerLooper) Init() PlayerLooper {
	rv := objc.Send[PlayerLooper](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerLooper) Autorelease() PlayerLooper {
	rv := objc.Send[PlayerLooper](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerLooper creates a new PlayerLooper instance.
func NewPlayerLooper() PlayerLooper {
	return getPlayerLooperClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PlayerLooper */
// An object that loops media content using a queue player.
//
// You can manually implement looping playback in your app using , but provides a much simpler interface to loop a single . You create a player looper by passing it a reference to your and a template and the looper automatically manages the looping playback of this content (see example).


// An object that loops media content using a queue player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper
type PlayerLooper struct {
	objectivec.Object
}

// PlayerLooperFrom constructs a [PlayerLooper] from an unsafe.Pointer.
//
// An object that loops media content using a queue player.
func PlayerLooperFrom(ptr unsafe.Pointer) PlayerLooper {
	return PlayerLooper{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PlayerLooper */

// Creates a player looper that continuously plays the full duration of a player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/init(player:templateItem:)
func NewPlayerLooperWithPlayerTemplateItem(player IAVQueuePlayer, itemToLoop IAVPlayerItem) PlayerLooper {
	rv := objc.Send[PlayerLooper](objc.ID(getPlayerLooperClass().class), objc.Sel("playerLooperWithPlayer:templateItem:"), player, itemToLoop)
	return rv
}/* debug [class_init_methods/constructor]: NewPlayerLooperWithPlayerTemplateItem */


// Creates a player looper that continuously plays the specified time range of a player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/init(player:templateItem:timeRange:)
func NewPlayerLooperWithPlayerTemplateItemTimeRange(player IAVQueuePlayer, itemToLoop IAVPlayerItem, loopRange TimeRange /* not a class type */) PlayerLooper {
	instance := getPlayerLooperClass().Alloc()
	rv := objc.Send[PlayerLooper](instance.ID, objc.Sel("initWithPlayer:templateItem:timeRange:"), player, itemToLoop, loopRange)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPlayerLooperWithPlayerTemplateItemTimeRange */


// Creates a player looper that continuously plays the full duration of a player item while adhering to the specified ordering of existing items in the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/init(player:templateItem:timeRange:existingItemsOrdering:)
func NewPlayerLooperWithPlayerTemplateItemTimeRangeExistingItemsOrdering(player IAVQueuePlayer, itemToLoop IAVPlayerItem, loopRange TimeRange /* not a class type */, itemOrdering PlayerLooperItemOrdering) PlayerLooper {
	instance := getPlayerLooperClass().Alloc()
	rv := objc.Send[PlayerLooper](instance.ID, objc.Sel("initWithPlayer:templateItem:timeRange:existingItemsOrdering:"), player, itemToLoop, loopRange, itemOrdering)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPlayerLooperWithPlayerTemplateItemTimeRangeExistingItemsOrdering */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PlayerLooper */

// Creates a player looper that continuously plays the full duration of a player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/init(player:templateItem:)
func (pc _PlayerLooperClass) PlayerLooperWithPlayerTemplateItem(player IAVQueuePlayer, itemToLoop IAVPlayerItem) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("playerLooperWithPlayer:templateItem:"), player, itemToLoop)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PlayerLooperWithPlayerTemplateItem) */


// Returns player looper that continuously plays the specified time range of a player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/playerLooperWithPlayer:templateItem:timeRange:
func (pc _PlayerLooperClass) PlayerLooperWithPlayerTemplateItemTimeRange(player IAVQueuePlayer, itemToLoop IAVPlayerItem, loopRange TimeRange /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("playerLooperWithPlayer:templateItem:timeRange:"), player, itemToLoop, loopRange)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PlayerLooperWithPlayerTemplateItemTimeRange) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PlayerLooper */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PlayerLooper */

// Disables looping for the player queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/disableLooping()
func (p_ PlayerLooper) DisableLooping() {
	objc.Send[objc.ID](p_.ID, objc.Sel("disableLooping"))
}/* debug [instance_methods/method]: DisableLooping */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PlayerLooper */

// An error that describes the reason looping failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/error
func (p_ PlayerLooper) Error() Error {
	rv := objc.Send[Error](p_.ID, objc.Sel("error"))
	return rv
}/* debug [instance_properties/getter]: error */


// The number of times the object played the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/loopCount
func (p_ PlayerLooper) LoopCount() int {
	rv := objc.Send[int](p_.ID, objc.Sel("loopCount"))
	return rv
}/* debug [instance_properties/getter]: loopCount */


// An array containing replicas of the template player item used to accomplish the looping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/loopingPlayerItems
func (p_ PlayerLooper) LoopingPlayerItems() []PlayerItem {
	rv := objc.Send[[]PlayerItem](p_.ID, objc.Sel("loopingPlayerItems"))
	return rv
}/* debug [instance_properties/getter]: loopingPlayerItems */


// A status that indicates the object’s ability to loop playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/status-swift.property
func (p_ PlayerLooper) Status() PlayerLooperStatus {
	rv := objc.Send[PlayerLooperStatus](p_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVPlayerLooper */


