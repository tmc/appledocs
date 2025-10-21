// Code generated from Apple documentation for MediaLibrary. DO NOT EDIT.

package medialibrary

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MediaGroup] class.
var (
	MediaGroupClass     _MediaGroupClass
	MediaGroupClassOnce sync.Once
)

func getMediaGroupClass() _MediaGroupClass {
	MediaGroupClassOnce.Do(func() {
		MediaGroupClass = _MediaGroupClass{objc.GetClass("MLMediaGroup")}
	})
	return MediaGroupClass
}

type _MediaGroupClass struct {
	class objc.Class
}

// An interface definition for the [MediaGroup] class.
type IMediaGroup interface {
	objectivec.IObject
}

// The class provides groupings for media objects from a single source of media, such as iTunes or Aperture. The media objects—individual files containing a piece of media such as a photo, song, or movie—are referenced by one or more groups within each media source. These groupings serve as filters, providing hierarchical structure to the collection of objects in each source.
//
// The structure of the group hierarchy is specific to each media source, but all sources have certain commonalities. For example, every source has a single root media group, which contains all groups and objects within that source. It is the highest-level parent group in the hierarchy and each of its descendant groups contains its own subgroups and their objects. All groups have a reference to their parent within the hierarchy. A group with no descendants contains only its own objects. If a media group does not contain any objects, it is not visible in the hierarchy. A media group has an array of attributes which can change at any point. For example, a media group may have certain attributes that describe its objects, but these attributes appear only after the objects for that group have been loaded. When any media group attribute changes, observers are notified via KVO notification. For information about handling attributes that change, see . Every media group has a unique identifier as well as a type identifier. In certain cases, multiple groups within a source can have the same type identifier. For descriptions of group type identifiers, see . All properties are read-only, so this information can be accessed but not altered.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaGroup
type MediaGroup struct {
	objectivec.Object
}

// MediaGroupFrom constructs a [MediaGroup] from an unsafe.Pointer.
//
// The class provides groupings for media objects from a single source of media, such as iTunes or Aperture. The media objects—individual files containing a piece of media such as a photo, song, or movie—are referenced by one or more groups within each media source. These groupings serve as filters, providing hierarchical structure to the collection of objects in each source.
func MediaGroupFrom(ptr unsafe.Pointer) MediaGroup {
	return MediaGroup{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MediaGroupClass) Alloc() MediaGroup {
	rv := objc.Send[MediaGroup](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MediaGroupClass) New() MediaGroup {
	rv := objc.Send[MediaGroup](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaGroup) Init() MediaGroup {
	rv := objc.Send[MediaGroup](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaGroup) Autorelease() MediaGroup {
	rv := objc.Send[MediaGroup](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaGroup creates a new MediaGroup instance.
func NewMediaGroup() MediaGroup {
	return getMediaGroupClass().New()
}


// The name of the media group.
//
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediagroup/name
func (m_ MediaGroup) Name() string {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The name of the media group.

//
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediagroup/name
func (m_ MediaGroup) SetName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), objc.String(value))
}

// An identifier for the source that loaded the media group.
//
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediagroup/mediasourceidentifier
func (m_ MediaGroup) MediaSourceIdentifier() string {
	rv := objc.Send[string](m_.ID, objc.Sel("mediaSourceIdentifier"))
	return rv
}


// SetMediaSourceIdentifier sets the value of the mediaSourceIdentifier property.
// An identifier for the source that loaded the media group.

//
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediagroup/mediasourceidentifier
func (m_ MediaGroup) SetMediaSourceIdentifier(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMediaSourceIdentifier:"), objc.String(value))
}

// A dictionary of attributes describing the media group.
//
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediagroup/attributes
func (m_ MediaGroup) Attributes() string {
	rv := objc.Send[string](m_.ID, objc.Sel("attributes"))
	return rv
}


// SetAttributes sets the value of the attributes property.
// A dictionary of attributes describing the media group.

//
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediagroup/attributes
func (m_ MediaGroup) SetAttributes(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttributes:"), objc.String(value))
}

// An identifier for the media group.
//
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediagroup/identifier
func (m_ MediaGroup) Identifier() string {
	rv := objc.Send[string](m_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// An identifier for the media group.

//
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmediagroup/identifier
func (m_ MediaGroup) SetIdentifier(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}

// A list of child groups contained in the media group.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaGroup/childGroups
func (m_ MediaGroup) ChildGroups() []MediaGroup {
	rv := objc.Send[[]MediaGroup](m_.ID, objc.Sel("childGroups"))
	return rv
}

// The media group’s icon.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaGroup/iconImage
func (m_ MediaGroup) IconImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("iconImage"))
	return rv
}

// A pointer to the media library instance that loaded the media group’s source.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaGroup/mediaLibrary
func (m_ MediaGroup) MediaLibrary() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mediaLibrary"))
	return rv
}

// A list of media objects in the media group.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaGroup/mediaObjects
func (m_ MediaGroup) MediaObjects() []MediaObject {
	rv := objc.Send[[]MediaObject](m_.ID, objc.Sel("mediaObjects"))
	return rv
}

// The date and time when the media group was last altered.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaGroup/modificationDate
func (m_ MediaGroup) ModificationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("modificationDate"))
	return rv
}

// The media group’s parent group.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaGroup/parent
func (m_ MediaGroup) Parent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("parent"))
	return rv
}

// An identifier for the media group’s type.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaGroup/typeIdentifier
func (m_ MediaGroup) TypeIdentifier() string {
	rv := objc.Send[string](m_.ID, objc.Sel("typeIdentifier"))
	return rv
}

// The location of the media group.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaGroup/url
func (m_ MediaGroup) URL() foundation.URL {
	rv := objc.Send[foundation.URL](m_.ID, objc.Sel("URL"))
	return rv
}



