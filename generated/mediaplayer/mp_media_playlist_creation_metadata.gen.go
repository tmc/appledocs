// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MediaPlaylistCreationMetadata] class.
var (
	MediaPlaylistCreationMetadataClass     _MediaPlaylistCreationMetadataClass
	MediaPlaylistCreationMetadataClassOnce sync.Once
)

func getMediaPlaylistCreationMetadataClass() _MediaPlaylistCreationMetadataClass {
	MediaPlaylistCreationMetadataClassOnce.Do(func() {
		MediaPlaylistCreationMetadataClass = _MediaPlaylistCreationMetadataClass{objc.GetClass("MPMediaPlaylistCreationMetadata")}
	})
	return MediaPlaylistCreationMetadataClass
}

type _MediaPlaylistCreationMetadataClass struct {
	class objc.Class
}

// An interface definition for the [MediaPlaylistCreationMetadata] class.
type IMediaPlaylistCreationMetadata interface {
	objectivec.IObject
}

// A set of attributes for describing a playlist when creating it.
//
// Use this class when creating a new playlist using the method. The system adds the metadata to the playlist when you create it, however it ignores the metadata if the playlist already exists.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylistCreationMetadata
type MediaPlaylistCreationMetadata struct {
	objectivec.Object
}

// MediaPlaylistCreationMetadataFrom constructs a [MediaPlaylistCreationMetadata] from an unsafe.Pointer.
//
// A set of attributes for describing a playlist when creating it.
func MediaPlaylistCreationMetadataFrom(ptr unsafe.Pointer) MediaPlaylistCreationMetadata {
	return MediaPlaylistCreationMetadata{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MediaPlaylistCreationMetadataClass) Alloc() MediaPlaylistCreationMetadata {
	rv := objc.Send[MediaPlaylistCreationMetadata](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MediaPlaylistCreationMetadataClass) New() MediaPlaylistCreationMetadata {
	rv := objc.Send[MediaPlaylistCreationMetadata](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaPlaylistCreationMetadata) Init() MediaPlaylistCreationMetadata {
	rv := objc.Send[MediaPlaylistCreationMetadata](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaPlaylistCreationMetadata) Autorelease() MediaPlaylistCreationMetadata {
	rv := objc.Send[MediaPlaylistCreationMetadata](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaPlaylistCreationMetadata creates a new MediaPlaylistCreationMetadata instance.
func NewMediaPlaylistCreationMetadata() MediaPlaylistCreationMetadata {
	return getMediaPlaylistCreationMetadataClass().New()
}




// Creates a new playlist metadata object with the designated name.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylistCreationMetadata/init(name:)
func NewMediaPlaylistCreationMetadataWithName(name string) MediaPlaylistCreationMetadata {
	instance := getMediaPlaylistCreationMetadataClass().Alloc()
	rv := objc.Send[MediaPlaylistCreationMetadata](instance.ID, objc.Sel("initWithName:"), objc.String(name))
	rv.Autorelease()
	return rv
}


// App defined display name for the playlist.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylistCreationMetadata/authorDisplayName
func (m_ MediaPlaylistCreationMetadata) AuthorDisplayName() string {
	rv := objc.Send[string](m_.ID, objc.Sel("authorDisplayName"))
	return rv
}


// SetAuthorDisplayName sets the value of the authorDisplayName property.
// App defined display name for the playlist.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylistCreationMetadata/authorDisplayName
func (m_ MediaPlaylistCreationMetadata) SetAuthorDisplayName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAuthorDisplayName:"), objc.String(value))
}

// The descriptive text for the playlist.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylistCreationMetadata/descriptionText
func (m_ MediaPlaylistCreationMetadata) DescriptionText() string {
	rv := objc.Send[string](m_.ID, objc.Sel("descriptionText"))
	return rv
}


// SetDescriptionText sets the value of the descriptionText property.
// The descriptive text for the playlist.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylistCreationMetadata/descriptionText
func (m_ MediaPlaylistCreationMetadata) SetDescriptionText(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDescriptionText:"), objc.String(value))
}

// The playlist’s displayed name.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylistCreationMetadata/name
func (m_ MediaPlaylistCreationMetadata) Name() string {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}


