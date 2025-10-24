// Code generated from Apple documentation for MediaLibrary. DO NOT EDIT.

package medialibrary

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLMediaLibrary */


/* debug [class_header]: Header for MLMediaLibrary */
// The class instance for the [MediaLibrary] class.
var (
	MediaLibraryClass     _MediaLibraryClass
	MediaLibraryClassOnce sync.Once
)

func getMediaLibraryClass() _MediaLibraryClass {
	MediaLibraryClassOnce.Do(func() {
		MediaLibraryClass = _MediaLibraryClass{objc.GetClass("MLMediaLibrary")}
	})
	return MediaLibraryClass
}

type _MediaLibraryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MediaLibrary */
// An interface definition for the [MediaLibrary] class.
type IMediaLibrary interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MediaLibrary */
	// properties:
	MediaSources() foundation.IDictionary
	MLMediaLoadFoldersKey() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MediaLibrary */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MediaLibrary */
// Alloc allocates a new instance without initialization.
func (mc _MediaLibraryClass) Alloc() MediaLibrary {
	rv := objc.Send[MediaLibrary](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MediaLibraryClass) New() MediaLibrary {
	rv := objc.Send[MediaLibrary](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaLibrary) Init() MediaLibrary {
	rv := objc.Send[MediaLibrary](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaLibrary) Autorelease() MediaLibrary {
	rv := objc.Send[MediaLibrary](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaLibrary creates a new MediaLibrary instance.
func NewMediaLibrary() MediaLibrary {
	return getMediaLibraryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MediaLibrary */
// The class provides an interface for accessing a collection of media objects from various sources. It serves as the initial access point of the Media Library framework.
//
// The media library structure is defined by , , and classes. At the highest level, all content within a media library instance is categorized by media source. Conceptually, a media source represents a single app, such as iTunes or Aperture. Each source contains a hierarchy of media groups that originates from a root group. These groups consist of media objects—individual files containing a piece of media such as a photo, song, or movie. Only one copy of each object exists within a media library instance, but an object can be referenced by multiple groups from a single source. The structure of the group hierarchy is specific to each media source. A media library is initialized using the method. The options argument to this method serves as a filter. By specifying which folders or sources to include or exclude during load, you can view a particular subset of groups and objects from your collection. All objects provided are thread-safe. For descriptions of possible load options, see . The typical and most efficient use case is to create and use one instance of for the lifetime of an app. When the underlying media files and metadata on the user’s system change, the corresponding data model objects (media groups and media objects) are automatically updated and KVO notifications are sent to notify the calling code of any changes. Multiple instances of can be created and used, but their sources, groups, and objects will be independent of those provided by other instances of .


// The class provides an interface for accessing a collection of media objects from various sources. It serves as the initial access point of the Media Library framework.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaLibrary
type MediaLibrary struct {
	objectivec.Object
}

// MediaLibraryFrom constructs a [MediaLibrary] from an unsafe.Pointer.
//
// The class provides an interface for accessing a collection of media objects from various sources. It serves as the initial access point of the Media Library framework.
func MediaLibraryFrom(ptr unsafe.Pointer) MediaLibrary {
	return MediaLibrary{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MediaLibrary */

// Initializes the media library based on the specified load options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaLibrary/init(options:)
func NewMediaLibraryWithOptions(options foundation.IDictionary) MediaLibrary {
	instance := getMediaLibraryClass().Alloc()
	rv := objc.Send[MediaLibrary](instance.ID, objc.Sel("initWithOptions:"), options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMediaLibraryWithOptions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MediaLibrary */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MediaLibrary */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MediaLibrary */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MediaLibrary */

// Returns a dictionary of media sources by identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaLibrary/MLMediaLibrary/mediaSources
func (m_ MediaLibrary) MediaSources() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("mediaSources"))
	return rv
}/* debug [instance_properties/getter]: mediaSources */


// Specifies the well-known folders that should be searched for media files. If this key is not present, none of the well-known folders will be provided. The value for this key is an array of strings (identifiers that correspond to well-known folder locations). For a list of well-known folder identifiers, see
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/medialibrary/mlmedialoadfolderskey
func (m_ MediaLibrary) MLMediaLoadFoldersKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MLMediaLoadFoldersKey"))
	return rv
}/* debug [instance_properties/getter]: MLMediaLoadFoldersKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLMediaLibrary */


