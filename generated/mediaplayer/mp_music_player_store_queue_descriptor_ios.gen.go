//go:build darwin && ios

// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for MusicPlayerStoreQueueDescriptor


// Sets the time the designated store item is to stop playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerStoreQueueDescriptor/setEndTime(_:forItemWithStoreID:)
func (m_ MusicPlayerStoreQueueDescriptor) SetEndTimeForItemWithStoreID(endTime float64, storeID objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndTime:forItemWithStoreID:"), endTime, storeID)
}

// Sets the time the designated store item is to start playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerStoreQueueDescriptor/setStartTime(_:forItemWithStoreID:)
func (m_ MusicPlayerStoreQueueDescriptor) SetStartTimeForItemWithStoreID(startTime float64, storeID objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:forItemWithStoreID:"), startTime, storeID)
}

// iOS-only properties

// The item identified by the store identifier to play first.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerStoreQueueDescriptor/startItemID
func (m_ MusicPlayerStoreQueueDescriptor) StartItemID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("startItemID"))
	return rv
}
func (m_ MusicPlayerStoreQueueDescriptor) SetStartItemID(value objc.IObject /* cross-framework: NSString */) {
	m_.ID.Send(objc.RegisterName("setStartItemID:"), value)
}

// An array containing the store identifiers found by the query used to create the queue descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerStoreQueueDescriptor/storeIDs
func (m_ MusicPlayerStoreQueueDescriptor) StoreIDs() []string {
	rv := objc.Send[[]string](m_.ID, objc.Sel("storeIDs"))
	return rv
}
func (m_ MusicPlayerStoreQueueDescriptor) SetStoreIDs(value []string) {
	m_.ID.Send(objc.RegisterName("setStoreIDs:"), value)
}




