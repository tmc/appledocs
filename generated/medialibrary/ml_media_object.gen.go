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

// An interface definition for the [MediaObject] class.
type IMediaObject interface {
	objectivec.IObject
	// properties:
	Attributes() foundation.IDictionary
	Name() objc.IObject /* cross-framework: NSString */
	ArtworkImage() objc.IObject /* cross-framework: Image */
	SetArtworkImage(value objc.IObject /* cross-framework: Image */)
	ContentType() objc.IObject /* cross-framework: NSString */
	SetContentType(value objc.IObject /* cross-framework: NSString */)
	FileSize() int
	SetFileSize(value int)
	Identifier() objc.IObject /* cross-framework: NSString */
	SetIdentifier(value objc.IObject /* cross-framework: NSString */)
	MediaLibrary() IMLMediaLibrary
	SetMediaLibrary(value IMLMediaLibrary)
	MediaSourceIdentifier() objc.IObject /* cross-framework: NSString */
	SetMediaSourceIdentifier(value objc.IObject /* cross-framework: NSString */)
	MediaType() MediaType /* not a class type */
	SetMediaType(value MediaType /* not a class type */)
	ModificationDate() objc.IObject /* cross-framework: Date */
	SetModificationDate(value objc.IObject /* cross-framework: Date */)
	OriginalURL() objc.IObject /* cross-framework: URL */
	SetOriginalURL(value objc.IObject /* cross-framework: URL */)
	ThumbnailURL() objc.IObject /* cross-framework: URL */
	SetThumbnailURL(value objc.IObject /* cross-framework: URL */)
	Url() objc.IObject /* cross-framework: URL */
	SetUrl(value objc.IObject /* cross-framework: URL */)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (mc _MediaObjectClass) Alloc() MediaObject {
	rv := objc.Send[MediaObject](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A dictionary of attributes describing the media object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaObject/attributes
func (m_ MediaObject) Attributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("attributes"))
	return rv
}


// The name of the media object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaObject/name
func (m_ MediaObject) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}


// Album artwork associated with the media object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediaobject/artworkimage
func (m_ MediaObject) ArtworkImage() objc.IObject /* cross-framework: Image */ {
	rv := objc.Send[appkit.Image](m_.ID, objc.Sel("artworkImage"))
	return rv
}


// Album artwork associated with the media object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediaobject/artworkimage
func (m_ MediaObject) SetArtworkImage(value objc.IObject /* cross-framework: Image */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArtworkImage:"), value)
}


// The UTI associated with the media object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediaobject/contenttype
func (m_ MediaObject) ContentType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("contentType"))
	return rv
}


// The UTI associated with the media object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediaobject/contenttype
func (m_ MediaObject) SetContentType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setContentType:"), value)
}


// The size, in bytes, of the media object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediaobject/filesize
func (m_ MediaObject) FileSize() int {
	rv := objc.Send[int](m_.ID, objc.Sel("fileSize"))
	return rv
}


// The size, in bytes, of the media object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediaobject/filesize
func (m_ MediaObject) SetFileSize(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFileSize:"), value)
}


// An identifier for the media object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediaobject/identifier
func (m_ MediaObject) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("identifier"))
	return rv
}


// An identifier for the media object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediaobject/identifier
func (m_ MediaObject) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIdentifier:"), value)
}


// A pointer to the media library instance that loaded the media object’s source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediaobject/medialibrary
func (m_ MediaObject) MediaLibrary() IMLMediaLibrary {
	rv := objc.Send[MediaLibrary](m_.ID, objc.Sel("mediaLibrary"))
	return rv
}


// A pointer to the media library instance that loaded the media object’s source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediaobject/medialibrary
func (m_ MediaObject) SetMediaLibrary(value IMLMediaLibrary) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMediaLibrary:"), value)
}


// An identifier for the source that loaded the media object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediaobject/mediasourceidentifier
func (m_ MediaObject) MediaSourceIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("mediaSourceIdentifier"))
	return rv
}


// An identifier for the source that loaded the media object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediaobject/mediasourceidentifier
func (m_ MediaObject) SetMediaSourceIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMediaSourceIdentifier:"), value)
}


// The media object’s type of media (image, audio, or movie).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediaobject/mediatype
func (m_ MediaObject) MediaType() MediaType /* not a class type */ {
	rv := objc.Send[MediaType](m_.ID, objc.Sel("mediaType"))
	return rv
}


// The media object’s type of media (image, audio, or movie).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediaobject/mediatype
func (m_ MediaObject) SetMediaType(value MediaType /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMediaType:"), value)
}


// The date and time when the media object was last altered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediaobject/modificationdate
func (m_ MediaObject) ModificationDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](m_.ID, objc.Sel("modificationDate"))
	return rv
}


// The date and time when the media object was last altered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediaobject/modificationdate
func (m_ MediaObject) SetModificationDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModificationDate:"), value)
}


// The location of the original media object, if
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediaobject/originalurl
func (m_ MediaObject) OriginalURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](m_.ID, objc.Sel("originalURL"))
	return rv
}


// The location of the original media object, if
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediaobject/originalurl
func (m_ MediaObject) SetOriginalURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOriginalURL:"), value)
}


// The location of the media object’s thumbnail image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediaobject/thumbnailurl
func (m_ MediaObject) ThumbnailURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](m_.ID, objc.Sel("thumbnailURL"))
	return rv
}


// The location of the media object’s thumbnail image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediaobject/thumbnailurl
func (m_ MediaObject) SetThumbnailURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setThumbnailURL:"), value)
}


// The location of the media object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediaobject/url
func (m_ MediaObject) Url() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](m_.ID, objc.Sel("url"))
	return rv
}


// The location of the media object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediaobject/url
func (m_ MediaObject) SetUrl(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUrl:"), value)
}



