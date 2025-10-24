//go:build darwin && ios

// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MediaLibrary


// Adds the designated item to the user’s music library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibrary/addItem(withProductID:completionHandler:)
func (m_ MediaLibrary) AddItemWithProductIDCompletionHandler(productID objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addItemWithProductID:completionHandler:"), productID, completionHandler)
}

// Asks the media library to turn on notifications for whenever the library changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibrary/beginGeneratingLibraryChangeNotifications()
func (m_ MediaLibrary) BeginGeneratingLibraryChangeNotifications() {
	objc.Send[objc.ID](m_.ID, objc.Sel("beginGeneratingLibraryChangeNotifications"))
}

// Asks the media library to turn off notifications for whenever the library changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibrary/endGeneratingLibraryChangeNotifications()
func (m_ MediaLibrary) EndGeneratingLibraryChangeNotifications() {
	objc.Send[objc.ID](m_.ID, objc.Sel("endGeneratingLibraryChangeNotifications"))
}

// Retrieves an app maintained existing playlist or creates a new playlist when no playlist exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibrary/getPlaylist(with:creationMetadata:completionHandler:)
func (m_ MediaLibrary) GetPlaylistWithUUIDCreationMetadataCompletionHandler(uuid objc.IObject /* cross-framework: UUID */, creationMetadata IMPMediaPlaylistCreationMetadata, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getPlaylistWithUUID:creationMetadata:completionHandler:"), uuid, creationMetadata, completionHandler)
}

// iOS-only properties

// The calendar date on which the media library was last modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibrary/lastModifiedDate
func (m_ MediaLibrary) LastModifiedDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("lastModifiedDate"))
	return rv
}





