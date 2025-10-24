// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPMediaPlaylistCreationMetadata */


/* debug [class_header]: Header for MPMediaPlaylistCreationMetadata */
// The class instance for the [MediaPlaylistCreationMetadata] class.
var (
	MediaPlaylistCreationMetadataClass     _MediaPlaylistCreationMetadataClass
	MediaPlaylistCreationMetadataClassOnce sync.Once
)

func getMediaPlaylistCreationMetadataClass() _MediaPlaylistCreationMetadataClass {
	MediaPlaylistCreationMetadataClassOnce.Do(func() {
		MediaPlaylistCreationMetadataClass = _MediaPlaylistCreationMetadataClass{objc.GetClass("MPMediaPlaylistCreationMetadata")}
	})
	return MediaPlaylistCreationMetadataClass
}

type _MediaPlaylistCreationMetadataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MediaPlaylistCreationMetadata */
// An interface definition for the [MediaPlaylistCreationMetadata] class.
type IMediaPlaylistCreationMetadata interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MediaPlaylistCreationMetadata */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MediaPlaylistCreationMetadata */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MediaPlaylistCreationMetadata */
// Alloc allocates a new instance without initialization.
func (mc _MediaPlaylistCreationMetadataClass) Alloc() MediaPlaylistCreationMetadata {
	rv := objc.Send[MediaPlaylistCreationMetadata](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MediaPlaylistCreationMetadataClass) New() MediaPlaylistCreationMetadata {
	rv := objc.Send[MediaPlaylistCreationMetadata](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaPlaylistCreationMetadata) Init() MediaPlaylistCreationMetadata {
	rv := objc.Send[MediaPlaylistCreationMetadata](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaPlaylistCreationMetadata) Autorelease() MediaPlaylistCreationMetadata {
	rv := objc.Send[MediaPlaylistCreationMetadata](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaPlaylistCreationMetadata creates a new MediaPlaylistCreationMetadata instance.
func NewMediaPlaylistCreationMetadata() MediaPlaylistCreationMetadata {
	return getMediaPlaylistCreationMetadataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MediaPlaylistCreationMetadata */
// A set of attributes for describing a playlist when creating it.
//
// Use this class when creating a new playlist using the method. The system adds the metadata to the playlist when you create it, however it ignores the metadata if the playlist already exists.


// A set of attributes for describing a playlist when creating it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylistCreationMetadata
type MediaPlaylistCreationMetadata struct {
	objectivec.Object
}

// MediaPlaylistCreationMetadataFrom constructs a [MediaPlaylistCreationMetadata] from an unsafe.Pointer.
//
// A set of attributes for describing a playlist when creating it.
func MediaPlaylistCreationMetadataFrom(ptr unsafe.Pointer) MediaPlaylistCreationMetadata {
	return MediaPlaylistCreationMetadata{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MediaPlaylistCreationMetadata */

// Creates a new playlist metadata object with the designated name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaPlaylistCreationMetadata/init(name:)
func NewMediaPlaylistCreationMetadataWithName(name objc.IObject /* cross-framework: NSString */) MediaPlaylistCreationMetadata {
	instance := getMediaPlaylistCreationMetadataClass().Alloc()
	rv := objc.Send[MediaPlaylistCreationMetadata](instance.ID, objc.Sel("initWithName:"), name)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMediaPlaylistCreationMetadataWithName */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MediaPlaylistCreationMetadata */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MediaPlaylistCreationMetadata */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MediaPlaylistCreationMetadata */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MediaPlaylistCreationMetadata */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPMediaPlaylistCreationMetadata */


