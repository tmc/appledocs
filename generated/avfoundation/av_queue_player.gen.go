// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [QueuePlayer] class.
var (
	QueuePlayerClass     _QueuePlayerClass
	QueuePlayerClassOnce sync.Once
)

func getQueuePlayerClass() _QueuePlayerClass {
	QueuePlayerClassOnce.Do(func() {
		QueuePlayerClass = _QueuePlayerClass{objc.GetClass("AVQueuePlayer")}
	})
	return QueuePlayerClass
}

type _QueuePlayerClass struct {
	class objc.Class
}





// An interface definition for the [QueuePlayer] class.
type IQueuePlayer interface {
	IPlayer
	

	// properties:


	

	// methods:
	AdvanceToNextItem()
	CanInsertItemAfterItem(item IAVPlayerItem, afterItem IAVPlayerItem) bool
	InsertItemAfterItem(item IAVPlayerItem, afterItem IAVPlayerItem)
	Items() []PlayerItem
	RemoveItem(item IAVPlayerItem)
	RemoveAllItems()


}





// Alloc allocates a new instance without initialization.
func (qc _QueuePlayerClass) Alloc() QueuePlayer {
	rv := objc.Send[QueuePlayer](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (qc _QueuePlayerClass) New() QueuePlayer {
	rv := objc.Send[QueuePlayer](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QueuePlayer) Init() QueuePlayer {
	rv := objc.Send[QueuePlayer](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QueuePlayer) Autorelease() QueuePlayer {
	rv := objc.Send[QueuePlayer](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQueuePlayer creates a new QueuePlayer instance.
func NewQueuePlayer() QueuePlayer {
	return getQueuePlayerClass().New()
}





// An object that plays a sequence of player items.
//
// Use an instance of this class to manage a queue of player items.


// An object that plays a sequence of player items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVQueuePlayer
type QueuePlayer struct {
	Player
}

// QueuePlayerFrom constructs a [QueuePlayer] from an unsafe.Pointer.
//
// An object that plays a sequence of player items.
func QueuePlayerFrom(ptr unsafe.Pointer) QueuePlayer {
	return QueuePlayer{
		Player: PlayerFrom(ptr),
	}
}






// Creates an object that plays a queue of items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVQueuePlayer/init(items:)
func NewQueuePlayerWithItems(items []PlayerItem) QueuePlayer {
	instance := getQueuePlayerClass().Alloc()
	rv := objc.Send[QueuePlayer](instance.ID, objc.Sel("initWithItems:"), items)
	rv.Autorelease()
	return rv
}







// Returns an object that plays a queue of items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVQueuePlayer/queuePlayerWithItems:
func (qc _QueuePlayerClass) QueuePlayerWithItems(items []PlayerItem) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(qc.class), objc.Sel("queuePlayerWithItems:"), items)
	return rv
}












// Ends playback of the current item and starts playback of the next item in the player’s queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVQueuePlayer/advanceToNextItem()
func (q_ QueuePlayer) AdvanceToNextItem() {
	objc.Send[objc.ID](q_.ID, objc.Sel("advanceToNextItem"))
}


// Returns a Boolean value that indicates whether you can insert a player item into the player’s queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVQueuePlayer/canInsert(_:after:)
func (q_ QueuePlayer) CanInsertItemAfterItem(item IAVPlayerItem, afterItem IAVPlayerItem) bool {
	rv := objc.Send[bool](q_.ID, objc.Sel("canInsertItem:afterItem:"), item, afterItem)
	return rv
}


// Inserts a player item after another player item in the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVQueuePlayer/insert(_:after:)
func (q_ QueuePlayer) InsertItemAfterItem(item IAVPlayerItem, afterItem IAVPlayerItem) {
	objc.Send[objc.ID](q_.ID, objc.Sel("insertItem:afterItem:"), item, afterItem)
}


// Returns an array of the currently enqueued items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVQueuePlayer/items()
func (q_ QueuePlayer) Items() []PlayerItem {
	rv := objc.Send[[]PlayerItem](q_.ID, objc.Sel("items"))
	return rv
}


// Removes a given player item from the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVQueuePlayer/remove(_:)
func (q_ QueuePlayer) RemoveItem(item IAVPlayerItem) {
	objc.Send[objc.ID](q_.ID, objc.Sel("removeItem:"), item)
}


// Removes all player items from the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVQueuePlayer/removeAllItems()
func (q_ QueuePlayer) RemoveAllItems() {
	objc.Send[objc.ID](q_.ID, objc.Sel("removeAllItems"))
}












