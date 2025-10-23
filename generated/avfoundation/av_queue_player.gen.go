// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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

// Alloc allocates a new instance without initialization.
func (qc _QueuePlayerClass) Alloc() QueuePlayer {
	rv := objc.Send[QueuePlayer](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




