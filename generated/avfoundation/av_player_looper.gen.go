// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [PlayerLooper] class.
type IPlayerLooper interface {
	objectivec.IObject
	

	// properties:
	Error() foundation.foundation.INSError
	LoopCount() int
	LoopingPlayerItems() []PlayerItem
	Status() PlayerLooperStatus


	

	// methods:
	DisableLooping()


}





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






// Creates a player looper that continuously plays the full duration of a player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/init(player:templateItem:)
func NewPlayerLooperWithPlayerTemplateItem(player IAVQueuePlayer, itemToLoop IAVPlayerItem) PlayerLooper {
	rv := objc.Send[PlayerLooper](objc.ID(getPlayerLooperClass().class), objc.Sel("playerLooperWithPlayer:templateItem:"), player, itemToLoop)
	return rv
}


// Creates a player looper that continuously plays the specified time range of a player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/init(player:templateItem:timeRange:)
func NewPlayerLooperWithPlayerTemplateItemTimeRange(player IAVQueuePlayer, itemToLoop IAVPlayerItem, loopRange objectivec.IObject) PlayerLooper {
	instance := getPlayerLooperClass().Alloc()
	rv := objc.Send[PlayerLooper](instance.ID, objc.Sel("initWithPlayer:templateItem:timeRange:"), player, itemToLoop, loopRange)
	rv.Autorelease()
	return rv
}


// Creates a player looper that continuously plays the full duration of a player item while adhering to the specified ordering of existing items in the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/init(player:templateItem:timeRange:existingItemsOrdering:)
func NewPlayerLooperWithPlayerTemplateItemTimeRangeExistingItemsOrdering(player IAVQueuePlayer, itemToLoop IAVPlayerItem, loopRange objectivec.IObject, itemOrdering PlayerLooperItemOrdering) PlayerLooper {
	instance := getPlayerLooperClass().Alloc()
	rv := objc.Send[PlayerLooper](instance.ID, objc.Sel("initWithPlayer:templateItem:timeRange:existingItemsOrdering:"), player, itemToLoop, loopRange, itemOrdering)
	rv.Autorelease()
	return rv
}







// Creates a player looper that continuously plays the full duration of a player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/init(player:templateItem:)
func (pc _PlayerLooperClass) PlayerLooperWithPlayerTemplateItem(player IAVQueuePlayer, itemToLoop IAVPlayerItem) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("playerLooperWithPlayer:templateItem:"), player, itemToLoop)
	return rv
}


// Returns player looper that continuously plays the specified time range of a player item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/playerLooperWithPlayer:templateItem:timeRange:
func (pc _PlayerLooperClass) PlayerLooperWithPlayerTemplateItemTimeRange(player IAVQueuePlayer, itemToLoop IAVPlayerItem, loopRange objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("playerLooperWithPlayer:templateItem:timeRange:"), player, itemToLoop, loopRange)
	return rv
}












// Disables looping for the player queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/disableLooping()
func (p_ PlayerLooper) DisableLooping() {
	objc.Send[objc.ID](p_.ID, objc.Sel("disableLooping"))
}







// An error that describes the reason looping failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/error
func (p_ PlayerLooper) Error() foundation.foundation.INSError {
	rv := objc.Send[foundation.NSError](p_.ID, objc.Sel("error"))
	return rv
}


// The number of times the object played the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/loopCount
func (p_ PlayerLooper) LoopCount() int {
	rv := objc.Send[int](p_.ID, objc.Sel("loopCount"))
	return rv
}


// An array containing replicas of the template player item used to accomplish the looping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/loopingPlayerItems
func (p_ PlayerLooper) LoopingPlayerItems() []PlayerItem {
	rv := objc.Send[[]PlayerItem](p_.ID, objc.Sel("loopingPlayerItems"))
	return rv
}


// A status that indicates the object’s ability to loop playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/status-swift.property
func (p_ PlayerLooper) Status() PlayerLooperStatus {
	rv := objc.Send[PlayerLooperStatus](p_.ID, objc.Sel("status"))
	return rv
}







