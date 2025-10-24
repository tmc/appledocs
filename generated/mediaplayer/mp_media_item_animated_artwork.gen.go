// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPMediaItemAnimatedArtwork */


/* debug [class_header]: Header for MPMediaItemAnimatedArtwork */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MediaItemAnimatedArtwork */
// An interface definition for the [MediaItemAnimatedArtwork] class.
type IMediaItemAnimatedArtwork interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MediaItemAnimatedArtwork */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MediaItemAnimatedArtwork */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MediaItemAnimatedArtwork */
// Alloc allocates a new instance without initialization.
func (mc _MediaItemAnimatedArtworkClass) Alloc() MediaItemAnimatedArtwork {
	rv := objc.Send[MediaItemAnimatedArtwork](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MediaItemAnimatedArtwork */
// An animated image, such as an animated music album cover art, for a media item.
//
// A single instance of animated artwork is comprised of two assets: an artwork video asset, and a preview image which should match the first frame of the artwork video. The preview image may be used when displaying the animated artwork whilst the video becomes available. Both the preview image and artwork video can be fetched asynchronously and will only be requested when required at point of display. Aim to provide preview images as quickly as possible once requested, and ideally synchronously. Video asset s you provide must be local file s. You should make the associated assets available locally before providing them via the relevant handler, for example by fetching the associated video asset over the network. The s should remain valid for the lifetime of the , once provided. should not be subclassed.


// An animated image, such as an animated music album cover art, for a media item.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MediaItemAnimatedArtwork */

// Creates an animated artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemAnimatedArtwork/init(artworkID:previewImageRequestHandler:videoAssetFileURLRequestHandler:)-ieue
func NewMediaItemAnimatedArtworkWithArtworkIDPreviewImageRequestHandlerVideoAssetFileURLRequestHandler(artworkID objc.IObject /* cross-framework: NSString */, previewImageRequestHandler unsafe.Pointer, videoAssetFileURLRequestHandler unsafe.Pointer) MediaItemAnimatedArtwork {
	instance := getMediaItemAnimatedArtworkClass().Alloc()
	rv := objc.Send[MediaItemAnimatedArtwork](instance.ID, objc.Sel("initWithArtworkID:previewImageRequestHandler:videoAssetFileURLRequestHandler:"), artworkID, previewImageRequestHandler, videoAssetFileURLRequestHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMediaItemAnimatedArtworkWithArtworkIDPreviewImageRequestHandlerVideoAssetFileURLRequestHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MediaItemAnimatedArtwork */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MediaItemAnimatedArtwork */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MediaItemAnimatedArtwork */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MediaItemAnimatedArtwork */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPMediaItemAnimatedArtwork */


