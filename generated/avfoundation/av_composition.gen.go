// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Composition] class.
var (
	CompositionClass     _CompositionClass
	CompositionClassOnce sync.Once
)

func getCompositionClass() _CompositionClass {
	CompositionClassOnce.Do(func() {
		CompositionClass = _CompositionClass{objc.GetClass("AVComposition")}
	})
	return CompositionClass
}

type _CompositionClass struct {
	class objc.Class
}

// An interface definition for the [Composition] class.
type IComposition interface {
	IAsset
	ChapterMetadataGroupsWithTitleLocaleContainingItemsWithCommonKeys(locale unsafe.Pointer, commonKeys unsafe.Pointer) []TimedMetadataGroup
	LoadTracksWithMediaTypeCompletionHandler(mediaType unsafe.Pointer, completionHandler unsafe.Pointer)
	MetadataForFormat(format unsafe.Pointer) []MetadataItem
}

// An object that combines and arranges media from multiple assets into a single composite asset that you can play or process.
//
// A composition is a container for one or more tracks of media. Its tracks are instances of that present media of a uniform type like audio or video. A track itself is a container for one or more segments of media, which are instances of , a type that represents a region of media in the source track.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition
type Composition struct {
	Asset
}

// CompositionFrom constructs a [Composition] from an unsafe.Pointer.
//
// An object that combines and arranges media from multiple assets into a single composite asset that you can play or process.
func CompositionFrom(ptr unsafe.Pointer) Composition {
	return Composition{
		Asset: AssetFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CompositionClass) Alloc() Composition {
	rv := objc.Send[Composition](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CompositionClass) New() Composition {
	rv := objc.Send[Composition](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Composition) Init() Composition {
	rv := objc.Send[Composition](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Composition) Autorelease() Composition {
	rv := objc.Send[Composition](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewComposition creates a new Composition instance.
func NewComposition() Composition {
	return getCompositionClass().New()
}


// Returns an array of chapters that contain the specified title locale and common keys.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/chapterMetadataGroups(withTitleLocale:containingItemsWithCommonKeys:)
func (c_ Composition) ChapterMetadataGroupsWithTitleLocaleContainingItemsWithCommonKeys(locale unsafe.Pointer, commonKeys unsafe.Pointer) []TimedMetadataGroup {
	rv := objc.Send[[]TimedMetadataGroup](c_.ID, objc.Sel("chapterMetadataGroupsWithTitleLocale:containingItemsWithCommonKeys:"), locale, commonKeys)
	return rv
}

// Loads tracks that contain media of a specified type.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/loadTracks(withMediaType:completionHandler:)
func (c_ Composition) LoadTracksWithMediaTypeCompletionHandler(mediaType unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("loadTracksWithMediaType:completionHandler:"), mediaType, completionHandler)
}

// Returns an array of metadata items from the container with the specified format.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/metadata(forFormat:)
func (c_ Composition) MetadataForFormat(format unsafe.Pointer) []MetadataItem {
	rv := objc.Send[[]MetadataItem](c_.ID, objc.Sel("metadataForFormat:"), format)
	return rv
}

// The tracks that a composition contains.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVComposition/tracks
func (c_ Composition) Tracks() []CompositionTrack {
	rv := objc.Send[[]CompositionTrack](c_.ID, objc.Sel("tracks"))
	return rv
}



