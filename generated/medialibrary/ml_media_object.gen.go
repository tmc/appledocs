// Code generated from Apple documentation for MediaLibrary. DO NOT EDIT.

package medialibrary

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLMediaObject */


/* debug [class_header]: Header for MLMediaObject */
// The class instance for the [MediaObject] class.
var (
	MediaObjectClass     _MediaObjectClass
	MediaObjectClassOnce sync.Once
)

func getMediaObjectClass() _MediaObjectClass {
	MediaObjectClassOnce.Do(func() {
		MediaObjectClass = _MediaObjectClass{objc.GetClass("MLMediaObject")}
	})
	return MediaObjectClass
}

type _MediaObjectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MediaObject */
// An interface definition for the [MediaObject] class.
type IMediaObject interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MediaObject */
	// properties:
	ArtworkImage() appkit.Image
	Attributes() foundation.IDictionary
	ContentType() objc.IObject /* cross-framework: NSString */
	FileSize() uint
	Identifier() objc.IObject /* cross-framework: NSString */
	MediaLibrary() IMLMediaLibrary
	MediaSourceIdentifier() objc.IObject /* cross-framework: NSString */
	MediaType() MediaType
	ModificationDate() objc.IObject /* cross-framework: NSDate */
	Name() objc.IObject /* cross-framework: NSString */
	OriginalURL() objc.IObject /* cross-framework: NSURL */
	ThumbnailURL() objc.IObject /* cross-framework: NSURL */
	URL() objc.IObject /* cross-framework: NSURL */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MediaObject */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MediaObject */
// Alloc allocates a new instance without initialization.
func (mc _MediaObjectClass) Alloc() MediaObject {
	rv := objc.Send[MediaObject](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MediaObjectClass) New() MediaObject {
	rv := objc.Send[MediaObject](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaObject) Init() MediaObject {
	rv := objc.Send[MediaObject](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaObject) Autorelease() MediaObject {
	rv := objc.Send[MediaObject](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaObject creates a new MediaObject instance.
func NewMediaObject() MediaObject {
	return getMediaObjectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MediaObject */
// The class describes a single media file, such as a photo, song, or movie. Each media object contains basic metadata including a name, media type, URL, and so on. Additional information about each object is stored in its list of attributes. For a list of possible object attribute keys, see .
//
// A media object belongs to a single media source but can be referenced by several groups within that source. In other words, an object can appear in multiple places in the group hierarchy under a single media source. In iTunes, a movie that was purchased through the iTunes Store is referenced by both the Purchased playlist and the Movies playlist. If a user adds the movie to his own playlist, the group respresenting that playlist will also reference the movie. All three groups reference the same media object. All properties are read-only, so this information can be accessed but not altered.


// The class describes a single media file, such as a photo, song, or movie. Each media object contains basic metadata including a name, media type, URL, and so on. Additional information about each object is stored in its list of attributes. For a list of possible object attribute keys, see .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaObject
type MediaObject struct {
	objectivec.Object
}

// MediaObjectFrom constructs a [MediaObject] from an unsafe.Pointer.
//
// The class describes a single media file, such as a photo, song, or movie. Each media object contains basic metadata including a name, media type, URL, and so on. Additional information about each object is stored in its list of attributes. For a list of possible object attribute keys, see .
func MediaObjectFrom(ptr unsafe.Pointer) MediaObject {
	return MediaObject{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MediaObject *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MediaObject */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MediaObject */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MediaObject */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MediaObject */

// Album artwork associated with the media object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaObject/artworkImage
func (m_ MediaObject) ArtworkImage() appkit.Image {
	rv := objc.Send[appkit.Image](m_.ID, objc.Sel("artworkImage"))
	return rv
}/* debug [instance_properties/getter]: artworkImage */


// A dictionary of attributes describing the media object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaObject/attributes
func (m_ MediaObject) Attributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("attributes"))
	return rv
}/* debug [instance_properties/getter]: attributes */


// The UTI associated with the media object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaObject/contentType
func (m_ MediaObject) ContentType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("contentType"))
	return rv
}/* debug [instance_properties/getter]: contentType */


// The size, in bytes, of the media object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaObject/fileSize
func (m_ MediaObject) FileSize() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("fileSize"))
	return rv
}/* debug [instance_properties/getter]: fileSize */


// An identifier for the media object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaObject/identifier
func (m_ MediaObject) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// A pointer to the media library instance that loaded the media object’s source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaObject/mediaLibrary
func (m_ MediaObject) MediaLibrary() IMLMediaLibrary {
	rv := objc.Send[MediaLibrary](m_.ID, objc.Sel("mediaLibrary"))
	return rv
}/* debug [instance_properties/getter]: mediaLibrary */


// An identifier for the source that loaded the media object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaObject/mediaSourceIdentifier
func (m_ MediaObject) MediaSourceIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("mediaSourceIdentifier"))
	return rv
}/* debug [instance_properties/getter]: mediaSourceIdentifier */


// The media object’s type of media (image, audio, or movie).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaObject/mediaType
func (m_ MediaObject) MediaType() MediaType {
	rv := objc.Send[MediaType](m_.ID, objc.Sel("mediaType"))
	return rv
}/* debug [instance_properties/getter]: mediaType */


// The date and time when the media object was last altered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaObject/modificationDate
func (m_ MediaObject) ModificationDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("modificationDate"))
	return rv
}/* debug [instance_properties/getter]: modificationDate */


// The name of the media object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaObject/name
func (m_ MediaObject) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The location of the original media object, if is not the original location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaObject/originalURL
func (m_ MediaObject) OriginalURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](m_.ID, objc.Sel("originalURL"))
	return rv
}/* debug [instance_properties/getter]: originalURL */


// The location of the media object’s thumbnail image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaObject/thumbnailURL
func (m_ MediaObject) ThumbnailURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](m_.ID, objc.Sel("thumbnailURL"))
	return rv
}/* debug [instance_properties/getter]: thumbnailURL */


// The location of the media object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaObject/url
func (m_ MediaObject) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](m_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_properties/getter]: URL */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLMediaObject */



