// Code generated from Apple documentation for MediaLibrary. DO NOT EDIT.

package medialibrary

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MediaSource] class.
var (
	MediaSourceClass     _MediaSourceClass
	MediaSourceClassOnce sync.Once
)

func getMediaSourceClass() _MediaSourceClass {
	MediaSourceClassOnce.Do(func() {
		MediaSourceClass = _MediaSourceClass{objc.GetClass("MLMediaSource")}
	})
	return MediaSourceClass
}

type _MediaSourceClass struct {
	class objc.Class
}

// An interface definition for the [MediaSource] class.
type IMediaSource interface {
	objectivec.IObject
	// properties:
	Attributes() foundation.IDictionary
	MediaLibrary() IMLMediaLibrary
	SetMediaLibrary(value IMLMediaLibrary)
	MediaSourceIdentifier() objc.IObject /* cross-framework: NSString */
	SetMediaSourceIdentifier(value objc.IObject /* cross-framework: NSString */)
	RootMediaGroup() IMLMediaGroup
	SetRootMediaGroup(value IMLMediaGroup)
	// methods:
}

// The class identifies a specific provider of media. Conceptually, a media source respresents a single app, such as iTunes or Aperture. Each media source contains multiple groups of media objects—individual files containing a piece of media such as a photo, song, or movie.
//
// The structure of the group hierarchy is specific to each media source, but all sources have certain commonalities. For example, every source has a single root media group, which contains all groups and objects within that source. It is the highest-level parent group in the hierarchy and each of its descendant groups contains its own subgroups and their objects. All groups have a reference to their parent within the hierarchy. A group with no descendants contains only its own objects. If a media group does not contain any objects, it is not visible in the hierarchy. Every media source has a unique media source identifier within a single media library instance. For a list of possible media source identifiers, see . All properties are read-only, so this information can be accessed but not altered.


// The class identifies a specific provider of media. Conceptually, a media source respresents a single app, such as iTunes or Aperture. Each media source contains multiple groups of media objects—individual files containing a piece of media such as a photo, song, or movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaSource
type MediaSource struct {
	objectivec.Object
}

// MediaSourceFrom constructs a [MediaSource] from an unsafe.Pointer.
//
// The class identifies a specific provider of media. Conceptually, a media source respresents a single app, such as iTunes or Aperture. Each media source contains multiple groups of media objects—individual files containing a piece of media such as a photo, song, or movie.
func MediaSourceFrom(ptr unsafe.Pointer) MediaSource {
	return MediaSource{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MediaSourceClass) Alloc() MediaSource {
	rv := objc.Send[MediaSource](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MediaSourceClass) New() MediaSource {
	rv := objc.Send[MediaSource](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaSource) Init() MediaSource {
	rv := objc.Send[MediaSource](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaSource) Autorelease() MediaSource {
	rv := objc.Send[MediaSource](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaSource creates a new MediaSource instance.
func NewMediaSource() MediaSource {
	return getMediaSourceClass().New()
}



// A list of attributes describing the media source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaSource/attributes
func (m_ MediaSource) Attributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("attributes"))
	return rv
}


// A pointer to the media library instance that loaded this media source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediasource/medialibrary
func (m_ MediaSource) MediaLibrary() IMLMediaLibrary {
	rv := objc.Send[MediaLibrary](m_.ID, objc.Sel("mediaLibrary"))
	return rv
}


// A pointer to the media library instance that loaded this media source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediasource/medialibrary
func (m_ MediaSource) SetMediaLibrary(value IMLMediaLibrary) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMediaLibrary:"), value)
}


// A unique identifier for the media source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediasource/mediasourceidentifier
func (m_ MediaSource) MediaSourceIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("mediaSourceIdentifier"))
	return rv
}


// A unique identifier for the media source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediasource/mediasourceidentifier
func (m_ MediaSource) SetMediaSourceIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMediaSourceIdentifier:"), value)
}


// The base media group in the media source that contains all other groups within the source as descendant elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediasource/rootmediagroup
func (m_ MediaSource) RootMediaGroup() IMLMediaGroup {
	rv := objc.Send[MediaGroup](m_.ID, objc.Sel("rootMediaGroup"))
	return rv
}


// The base media group in the media source that contains all other groups within the source as descendant elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediasource/rootmediagroup
func (m_ MediaSource) SetRootMediaGroup(value IMLMediaGroup) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootMediaGroup:"), value)
}




