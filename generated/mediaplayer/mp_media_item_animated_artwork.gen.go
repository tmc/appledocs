// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MediaItemAnimatedArtwork] class.
var (
	MediaItemAnimatedArtworkClass     _MediaItemAnimatedArtworkClass
	MediaItemAnimatedArtworkClassOnce sync.Once
)

func getMediaItemAnimatedArtworkClass() _MediaItemAnimatedArtworkClass {
	MediaItemAnimatedArtworkClassOnce.Do(func() {
		MediaItemAnimatedArtworkClass = _MediaItemAnimatedArtworkClass{objc.GetClass("MPMediaItemAnimatedArtwork")}
	})
	return MediaItemAnimatedArtworkClass
}

type _MediaItemAnimatedArtworkClass struct {
	class objc.Class
}

// An interface definition for the [MediaItemAnimatedArtwork] class.
type IMediaItemAnimatedArtwork interface {
	objectivec.IObject
}

// An animated image, such as an animated music album cover art, for a media item.
//
// A single instance of animated artwork is comprised of two assets: an artwork video asset, and a preview image which should match the first frame of the artwork video. The preview image may be used when displaying the animated artwork whilst the video becomes available. Both the preview image and artwork video can be fetched asynchronously and will only be requested when required at point of display. Aim to provide preview images as quickly as possible once requested, and ideally synchronously. Video asset s you provide must be local file s. You should make the associated assets available locally before providing them via the relevant handler, for example by fetching the associated video asset over the network. The s should remain valid for the lifetime of the , once provided. should not be subclassed.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemAnimatedArtwork
type MediaItemAnimatedArtwork struct {
	objectivec.Object
}

// MediaItemAnimatedArtworkFrom constructs a [MediaItemAnimatedArtwork] from an unsafe.Pointer.
//
// An animated image, such as an animated music album cover art, for a media item.
func MediaItemAnimatedArtworkFrom(ptr unsafe.Pointer) MediaItemAnimatedArtwork {
	return MediaItemAnimatedArtwork{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MediaItemAnimatedArtworkClass) Alloc() MediaItemAnimatedArtwork {
	rv := objc.Send[MediaItemAnimatedArtwork](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MediaItemAnimatedArtworkClass) New() MediaItemAnimatedArtwork {
	rv := objc.Send[MediaItemAnimatedArtwork](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaItemAnimatedArtwork) Init() MediaItemAnimatedArtwork {
	rv := objc.Send[MediaItemAnimatedArtwork](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaItemAnimatedArtwork) Autorelease() MediaItemAnimatedArtwork {
	rv := objc.Send[MediaItemAnimatedArtwork](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaItemAnimatedArtwork creates a new MediaItemAnimatedArtwork instance.
func NewMediaItemAnimatedArtwork() MediaItemAnimatedArtwork {
	return getMediaItemAnimatedArtworkClass().New()
}


// Creates an animated artwork.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemAnimatedArtwork/init(artworkID:previewImageRequestHandler:videoAssetFileURLRequestHandler:)-ieue
func NewMediaItemAnimatedArtworkWithArtworkIDPreviewImageRequestHandlerVideoAssetFileURLRequestHandler(artworkID string, previewImageRequestHandler unsafe.Pointer, videoAssetFileURLRequestHandler unsafe.Pointer) MediaItemAnimatedArtwork {
	instance := getMediaItemAnimatedArtworkClass().Alloc()
	rv := objc.Send[MediaItemAnimatedArtwork](instance.ID, objc.Sel("initWithArtworkID:previewImageRequestHandler:videoAssetFileURLRequestHandler:"), objc.String(artworkID), previewImageRequestHandler, videoAssetFileURLRequestHandler)
	rv.Autorelease()
	return rv
}



