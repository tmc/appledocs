// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MusicPlayerStoreQueueDescriptor] class.
var (
	MusicPlayerStoreQueueDescriptorClass     _MusicPlayerStoreQueueDescriptorClass
	MusicPlayerStoreQueueDescriptorClassOnce sync.Once
)

func getMusicPlayerStoreQueueDescriptorClass() _MusicPlayerStoreQueueDescriptorClass {
	MusicPlayerStoreQueueDescriptorClassOnce.Do(func() {
		MusicPlayerStoreQueueDescriptorClass = _MusicPlayerStoreQueueDescriptorClass{objc.GetClass("MPMusicPlayerStoreQueueDescriptor")}
	})
	return MusicPlayerStoreQueueDescriptorClass
}

type _MusicPlayerStoreQueueDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MusicPlayerStoreQueueDescriptor] class.
type IMusicPlayerStoreQueueDescriptor interface {
	IMusicPlayerQueueDescriptor
	SetEndTimeForItemWithStoreID(endTime foundation.TimeInterval, storeID string)
	SetStartTimeForItemWithStoreID(startTime foundation.TimeInterval, storeID string)
}

// A set of properties and methods for modifying items, based on their store identifier, in the player’s queue.
//
// Use this class to modify the player queue created by a query before the queue begins to play. You can modify when individual items start and stop playing, along with setting the first item to play.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerStoreQueueDescriptor
type MusicPlayerStoreQueueDescriptor struct {
	MusicPlayerQueueDescriptor
}

// MusicPlayerStoreQueueDescriptorFrom constructs a [MusicPlayerStoreQueueDescriptor] from an unsafe.Pointer.
//
// A set of properties and methods for modifying items, based on their store identifier, in the player’s queue.
func MusicPlayerStoreQueueDescriptorFrom(ptr unsafe.Pointer) MusicPlayerStoreQueueDescriptor {
	return MusicPlayerStoreQueueDescriptor{
		MusicPlayerQueueDescriptor: MusicPlayerQueueDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MusicPlayerStoreQueueDescriptorClass) Alloc() MusicPlayerStoreQueueDescriptor {
	rv := objc.Send[MusicPlayerStoreQueueDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MusicPlayerStoreQueueDescriptorClass) New() MusicPlayerStoreQueueDescriptor {
	rv := objc.Send[MusicPlayerStoreQueueDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MusicPlayerStoreQueueDescriptor) Init() MusicPlayerStoreQueueDescriptor {
	rv := objc.Send[MusicPlayerStoreQueueDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MusicPlayerStoreQueueDescriptor) Autorelease() MusicPlayerStoreQueueDescriptor {
	rv := objc.Send[MusicPlayerStoreQueueDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMusicPlayerStoreQueueDescriptor creates a new MusicPlayerStoreQueueDescriptor instance.
func NewMusicPlayerStoreQueueDescriptor() MusicPlayerStoreQueueDescriptor {
	return getMusicPlayerStoreQueueDescriptorClass().New()
}




// Creates a new queue descriptor using the designated store identifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerStoreQueueDescriptor/init(storeIDs:)
func NewMusicPlayerStoreQueueDescriptorWithStoreIDs(storeIDs unsafe.Pointer) MusicPlayerStoreQueueDescriptor {
	instance := getMusicPlayerStoreQueueDescriptorClass().Alloc()
	rv := objc.Send[MusicPlayerStoreQueueDescriptor](instance.ID, objc.Sel("initWithStoreIDs:"), storeIDs)
	rv.Autorelease()
	return rv
}


// Sets the time the designated store item is to stop playing.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerStoreQueueDescriptor/setEndTime(_:forItemWithStoreID:)
func (m_ MusicPlayerStoreQueueDescriptor) SetEndTimeForItemWithStoreID(endTime foundation.TimeInterval, storeID string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndTime:forItemWithStoreID:"), endTime, objc.String(storeID))
}

// Sets the time the designated store item is to start playing.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerStoreQueueDescriptor/setStartTime(_:forItemWithStoreID:)
func (m_ MusicPlayerStoreQueueDescriptor) SetStartTimeForItemWithStoreID(startTime foundation.TimeInterval, storeID string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:forItemWithStoreID:"), startTime, objc.String(storeID))
}

// The item identified by the store identifier to play first.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerStoreQueueDescriptor/startItemID
func (m_ MusicPlayerStoreQueueDescriptor) StartItemID() string {
	rv := objc.Send[string](m_.ID, objc.Sel("startItemID"))
	return rv
}


// SetStartItemID sets the value of the startItemID property.
// The item identified by the store identifier to play first.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerStoreQueueDescriptor/startItemID
func (m_ MusicPlayerStoreQueueDescriptor) SetStartItemID(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartItemID:"), objc.String(value))
}

// An array containing the store identifiers found by the query used to create the queue descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerStoreQueueDescriptor/storeIDs
func (m_ MusicPlayerStoreQueueDescriptor) StoreIDs() []string {
	rv := objc.Send[[]string](m_.ID, objc.Sel("storeIDs"))
	return rv
}


// SetStoreIDs sets the value of the storeIDs property.
// An array containing the store identifiers found by the query used to create the queue descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerStoreQueueDescriptor/storeIDs
func (m_ MusicPlayerStoreQueueDescriptor) SetStoreIDs(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setStoreIDs:"), nsArray)
}


