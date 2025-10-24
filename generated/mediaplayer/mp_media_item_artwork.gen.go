// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MediaItemArtwork] class.
var (
	MediaItemArtworkClass     _MediaItemArtworkClass
	MediaItemArtworkClassOnce sync.Once
)

func getMediaItemArtworkClass() _MediaItemArtworkClass {
	MediaItemArtworkClassOnce.Do(func() {
		MediaItemArtworkClass = _MediaItemArtworkClass{objc.GetClass("MPMediaItemArtwork")}
	})
	return MediaItemArtworkClass
}

type _MediaItemArtworkClass struct {
	class objc.Class
}

// An interface definition for the [MediaItemArtwork] class.
type IMediaItemArtwork interface {
	objectivec.IObject
	// properties:
	Bounds() objc.IObject /* cross-framework: Rect */
	ImageCropRect() objc.IObject /* cross-framework: Rect */
	// methods:
	ImageWithSize(size objc.IObject /* cross-framework: Size */) objc.IObject /* cross-framework: Image */
}

// A graphical image, such as music album cover art, associated with a media item.


// A graphical image, such as music album cover art, associated with a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemArtwork
type MediaItemArtwork struct {
	objectivec.Object
}

// MediaItemArtworkFrom constructs a [MediaItemArtwork] from an unsafe.Pointer.
//
// A graphical image, such as music album cover art, associated with a media item.
func MediaItemArtworkFrom(ptr unsafe.Pointer) MediaItemArtwork {
	return MediaItemArtwork{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MediaItemArtworkClass) Alloc() MediaItemArtwork {
	rv := objc.Send[MediaItemArtwork](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MediaItemArtworkClass) New() MediaItemArtwork {
	rv := objc.Send[MediaItemArtwork](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaItemArtwork) Init() MediaItemArtwork {
	rv := objc.Send[MediaItemArtwork](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaItemArtwork) Autorelease() MediaItemArtwork {
	rv := objc.Send[MediaItemArtwork](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaItemArtwork creates a new MediaItemArtwork instance.
func NewMediaItemArtwork() MediaItemArtwork {
	return getMediaItemArtworkClass().New()
}



// Creates a new image from existing artwork with the specified bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemArtwork/init(boundsSize:requestHandler:)
func NewMediaItemArtworkWithBoundsSizeRequestHandler(boundsSize objc.IObject /* cross-framework: Size */, requestHandler Image  * (^)( CGSize size /* not a class type */) MediaItemArtwork {
	instance := getMediaItemArtworkClass().Alloc()
	rv := objc.Send[MediaItemArtwork](instance.ID, objc.Sel("initWithBoundsSize:requestHandler:"), boundsSize, requestHandler)
	rv.Autorelease()
	return rv
}


// Initializes a media item artwork instance with a full-size image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemArtwork/init(image:)
func NewMediaItemArtworkWithImage(image objc.IObject /* cross-framework: Image */) MediaItemArtwork {
	instance := getMediaItemArtworkClass().Alloc()
	rv := objc.Send[MediaItemArtwork](instance.ID, objc.Sel("initWithImage:"), image)
	rv.Autorelease()
	return rv
}



// Returns the artwork image for an item at the given size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemArtwork/image(at:)
func (m_ MediaItemArtwork) ImageWithSize(size objc.IObject /* cross-framework: Size */) objc.IObject /* cross-framework: Image */ {
	rv := objc.Send[appkit.Image](m_.ID, objc.Sel("imageWithSize:"), size)
	return rv
}


// The maximum size, in points, of the image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemArtwork/bounds
func (m_ MediaItemArtwork) Bounds() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](m_.ID, objc.Sel("bounds"))
	return rv
}


// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemArtwork/imageCropRect
func (m_ MediaItemArtwork) ImageCropRect() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](m_.ID, objc.Sel("imageCropRect"))
	return rv
}


