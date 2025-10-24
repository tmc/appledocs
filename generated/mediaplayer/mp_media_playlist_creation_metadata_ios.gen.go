//go:build darwin && ios

// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MediaPlaylistCreationMetadata


// iOS-only properties

// App defined display name for the playlist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylistCreationMetadata/authorDisplayName
func (m_ MediaPlaylistCreationMetadata) AuthorDisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("authorDisplayName"))
	return rv
}
func (m_ MediaPlaylistCreationMetadata) SetAuthorDisplayName(value objc.IObject /* cross-framework: NSString */) {
	m_.ID.Send(objc.RegisterName("setAuthorDisplayName:"), value)
}

// The descriptive text for the playlist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylistCreationMetadata/descriptionText
func (m_ MediaPlaylistCreationMetadata) DescriptionText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("descriptionText"))
	return rv
}
func (m_ MediaPlaylistCreationMetadata) SetDescriptionText(value objc.IObject /* cross-framework: NSString */) {
	m_.ID.Send(objc.RegisterName("setDescriptionText:"), value)
}

// The playlist’s displayed name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylistCreationMetadata/name
func (m_ MediaPlaylistCreationMetadata) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}




