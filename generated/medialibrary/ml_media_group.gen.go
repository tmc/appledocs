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

/* debug [class.gen.go]: Generating class MLMediaGroup */


/* debug [class_header]: Header for MLMediaGroup */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MediaGroup */
// An interface definition for the [MediaGroup] class.
type IMediaGroup interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MediaGroup */
	// properties:
	Attributes() foundation.IDictionary
	ChildGroups() []MediaGroup
	IconImage() appkit.Image
	Identifier() objc.IObject /* cross-framework: NSString */
	MediaLibrary() IMLMediaLibrary
	MediaObjects() []MediaObject
	MediaSourceIdentifier() objc.IObject /* cross-framework: NSString */
	ModificationDate() objc.IObject /* cross-framework: NSDate */
	Name() objc.IObject /* cross-framework: NSString */
	Parent() IMLMediaGroup
	TypeIdentifier() objc.IObject /* cross-framework: NSString */
	URL() objc.IObject /* cross-framework: NSURL */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MediaGroup */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MediaGroup */
// Alloc allocates a new instance without initialization.
func (mc _MediaGroupClass) Alloc() MediaGroup {
	rv := objc.Send[MediaGroup](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MediaGroup */
// The class provides groupings for media objects from a single source of media, such as iTunes or Aperture. The media objects—individual files containing a piece of media such as a photo, song, or movie—are referenced by one or more groups within each media source. These groupings serve as filters, providing hierarchical structure to the collection of objects in each source.
//
// The structure of the group hierarchy is specific to each media source, but all sources have certain commonalities. For example, every source has a single root media group, which contains all groups and objects within that source. It is the highest-level parent group in the hierarchy and each of its descendant groups contains its own subgroups and their objects. All groups have a reference to their parent within the hierarchy. A group with no descendants contains only its own objects. If a media group does not contain any objects, it is not visible in the hierarchy. A media group has an array of attributes which can change at any point. For example, a media group may have certain attributes that describe its objects, but these attributes appear only after the objects for that group have been loaded. When any media group attribute changes, observers are notified via KVO notification. For information about handling attributes that change, see . Every media group has a unique identifier as well as a type identifier. In certain cases, multiple groups within a source can have the same type identifier. For descriptions of group type identifiers, see . All properties are read-only, so this information can be accessed but not altered.


// The class provides groupings for media objects from a single source of media, such as iTunes or Aperture. The media objects—individual files containing a piece of media such as a photo, song, or movie—are referenced by one or more groups within each media source. These groupings serve as filters, providing hierarchical structure to the collection of objects in each source.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MediaGroup *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MediaGroup */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MediaGroup */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MediaGroup */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MediaGroup */

// A dictionary of attributes describing the media group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaGroup/attributes
func (m_ MediaGroup) Attributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("attributes"))
	return rv
}/* debug [instance_properties/getter]: attributes */


// A list of child groups contained in the media group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaGroup/childGroups
func (m_ MediaGroup) ChildGroups() []MediaGroup {
	rv := objc.Send[[]MediaGroup](m_.ID, objc.Sel("childGroups"))
	return rv
}/* debug [instance_properties/getter]: childGroups */


// The media group’s icon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaGroup/iconImage
func (m_ MediaGroup) IconImage() appkit.Image {
	rv := objc.Send[appkit.Image](m_.ID, objc.Sel("iconImage"))
	return rv
}/* debug [instance_properties/getter]: iconImage */


// An identifier for the media group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaGroup/identifier
func (m_ MediaGroup) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// A pointer to the media library instance that loaded the media group’s source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaGroup/mediaLibrary
func (m_ MediaGroup) MediaLibrary() IMLMediaLibrary {
	rv := objc.Send[MediaLibrary](m_.ID, objc.Sel("mediaLibrary"))
	return rv
}/* debug [instance_properties/getter]: mediaLibrary */


// A list of media objects in the media group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaGroup/mediaObjects
func (m_ MediaGroup) MediaObjects() []MediaObject {
	rv := objc.Send[[]MediaObject](m_.ID, objc.Sel("mediaObjects"))
	return rv
}/* debug [instance_properties/getter]: mediaObjects */


// An identifier for the source that loaded the media group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaGroup/mediaSourceIdentifier
func (m_ MediaGroup) MediaSourceIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("mediaSourceIdentifier"))
	return rv
}/* debug [instance_properties/getter]: mediaSourceIdentifier */


// The date and time when the media group was last altered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaGroup/modificationDate
func (m_ MediaGroup) ModificationDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("modificationDate"))
	return rv
}/* debug [instance_properties/getter]: modificationDate */


// The name of the media group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaGroup/name
func (m_ MediaGroup) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The media group’s parent group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaGroup/parent
func (m_ MediaGroup) Parent() IMLMediaGroup {
	rv := objc.Send[MediaGroup](m_.ID, objc.Sel("parent"))
	return rv
}/* debug [instance_properties/getter]: parent */


// An identifier for the media group’s type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaGroup/typeIdentifier
func (m_ MediaGroup) TypeIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("typeIdentifier"))
	return rv
}/* debug [instance_properties/getter]: typeIdentifier */


// The location of the media group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaGroup/url
func (m_ MediaGroup) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](m_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_properties/getter]: URL */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLMediaGroup */



