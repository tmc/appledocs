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

/* debug [class.gen.go]: Generating class MPMediaItemArtwork */


/* debug [class_header]: Header for MPMediaItemArtwork */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MediaItemArtwork */
// An interface definition for the [MediaItemArtwork] class.
type IMediaItemArtwork interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MediaItemArtwork */
	// properties:
	Bounds() corefoundation.CGRect
	ImageCropRect() corefoundation.CGRect
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MediaItemArtwork */
	// methods:
	ImageWithSize(size corefoundation.CGSize) appkit.Image
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MediaItemArtwork */
// Alloc allocates a new instance without initialization.
func (mc _MediaItemArtworkClass) Alloc() MediaItemArtwork {
	rv := objc.Send[MediaItemArtwork](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MediaItemArtwork */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MediaItemArtwork */

// Creates a new image from existing artwork with the specified bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemArtwork/init(boundsSize:requestHandler:)
func NewMediaItemArtworkWithBoundsSizeRequestHandler(boundsSize corefoundation.CGSize, requestHandler unsafe.Pointer) MediaItemArtwork {
	instance := getMediaItemArtworkClass().Alloc()
	rv := objc.Send[MediaItemArtwork](instance.ID, objc.Sel("initWithBoundsSize:requestHandler:"), boundsSize, requestHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMediaItemArtworkWithBoundsSizeRequestHandler */


// Initializes a media item artwork instance with a full-size image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemArtwork/init(image:)
func NewMediaItemArtworkWithImage(image appkit.Image) MediaItemArtwork {
	instance := getMediaItemArtworkClass().Alloc()
	rv := objc.Send[MediaItemArtwork](instance.ID, objc.Sel("initWithImage:"), image)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMediaItemArtworkWithImage */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MediaItemArtwork */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MediaItemArtwork */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MediaItemArtwork */

// Returns the artwork image for an item at the given size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemArtwork/image(at:)
func (m_ MediaItemArtwork) ImageWithSize(size corefoundation.CGSize) appkit.Image {
	rv := objc.Send[appkit.Image](m_.ID, objc.Sel("imageWithSize:"), size)
	return rv
}/* debug [instance_methods/method]: ImageWithSize */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MediaItemArtwork */

// The maximum size, in points, of the image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemArtwork/bounds
func (m_ MediaItemArtwork) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](m_.ID, objc.Sel("bounds"))
	return rv
}/* debug [instance_properties/getter]: bounds */


// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaItemArtwork/imageCropRect
func (m_ MediaItemArtwork) ImageCropRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](m_.ID, objc.Sel("imageCropRect"))
	return rv
}/* debug [instance_properties/getter]: imageCropRect */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPMediaItemArtwork */


