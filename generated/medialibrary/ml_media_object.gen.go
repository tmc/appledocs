// Code generated from Apple documentation for MediaLibrary. DO NOT EDIT.

package medialibrary

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// The class describes a single media file, such as a photo, song, or movie. Each media object contains basic metadata including a name, media type, URL, and so on. Additional information about each object is stored in its list of attributes. For a list of possible object attribute keys, see .
//
// A media object belongs to a single media source but can be referenced by several groups within that source. In other words, an object can appear in multiple places in the group hierarchy under a single media source. In iTunes, a movie that was purchased through the iTunes Store is referenced by both the Purchased playlist and the Movies playlist. If a user adds the movie to his own playlist, the group respresenting that playlist will also reference the movie. All three groups reference the same media object. All properties are read-only, so this information can be accessed but not altered.
//
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
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaObject/attributes
func (m_ MediaObject) Attributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("attributes"))
	return rv
}

// The UTI associated with the media object.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaObject/contentType
func (m_ MediaObject) ContentType() string {
	rv := objc.Send[string](m_.ID, objc.Sel("contentType"))
	return rv
}

// The size, in bytes, of the media object.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaObject/fileSize
func (m_ MediaObject) FileSize() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("fileSize"))
	return rv
}

// An identifier for the media object.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaObject/identifier
func (m_ MediaObject) Identifier() string {
	rv := objc.Send[string](m_.ID, objc.Sel("identifier"))
	return rv
}

// An identifier for the source that loaded the media object.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaObject/mediaSourceIdentifier
func (m_ MediaObject) MediaSourceIdentifier() string {
	rv := objc.Send[string](m_.ID, objc.Sel("mediaSourceIdentifier"))
	return rv
}

// The name of the media object.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaObject/name
func (m_ MediaObject) Name() string {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}

// The location of the media object’s thumbnail image.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaObject/thumbnailURL
func (m_ MediaObject) ThumbnailURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("thumbnailURL"))
	return rv
}



