// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVMediaDataStorage */


/* debug [class_header]: Header for AVMediaDataStorage */
// The class instance for the [MediaDataStorage] class.
var (
	MediaDataStorageClass     _MediaDataStorageClass
	MediaDataStorageClassOnce sync.Once
)

func getMediaDataStorageClass() _MediaDataStorageClass {
	MediaDataStorageClassOnce.Do(func() {
		MediaDataStorageClass = _MediaDataStorageClass{objc.GetClass("AVMediaDataStorage")}
	})
	return MediaDataStorageClass
}

type _MediaDataStorageClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MediaDataStorage */
// An interface definition for the [MediaDataStorage] class.
type IMediaDataStorage interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MediaDataStorage */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MediaDataStorage */
	// methods:
	URL() foundation.URL
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MediaDataStorage */
// Alloc allocates a new instance without initialization.
func (mc _MediaDataStorageClass) Alloc() MediaDataStorage {
	rv := objc.Send[MediaDataStorage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MediaDataStorageClass) New() MediaDataStorage {
	rv := objc.Send[MediaDataStorage](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaDataStorage) Init() MediaDataStorage {
	rv := objc.Send[MediaDataStorage](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaDataStorage) Autorelease() MediaDataStorage {
	rv := objc.Send[MediaDataStorage](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaDataStorage creates a new MediaDataStorage instance.
func NewMediaDataStorage() MediaDataStorage {
	return getMediaDataStorageClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MediaDataStorage */
// An object that represents the media sample data storage file.


// An object that represents the media sample data storage file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaDataStorage
type MediaDataStorage struct {
	objectivec.Object
}

// MediaDataStorageFrom constructs a [MediaDataStorage] from an unsafe.Pointer.
//
// An object that represents the media sample data storage file.
func MediaDataStorageFrom(ptr unsafe.Pointer) MediaDataStorage {
	return MediaDataStorage{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MediaDataStorage */

// Creates a media data storage object associated with a file URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaDataStorage/init(url:options:)
func NewMediaDataStorageWithURLOptions(URL objc.IObject /* cross-framework: NSURL */, options foundation.IDictionary) MediaDataStorage {
	instance := getMediaDataStorageClass().Alloc()
	rv := objc.Send[MediaDataStorage](instance.ID, objc.Sel("initWithURL:options:"), URL, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMediaDataStorageWithURLOptions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MediaDataStorage */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MediaDataStorage */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MediaDataStorage */

// Returns the URL used to initialize the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaDataStorage/url()
func (m_ MediaDataStorage) URL() foundation.URL {
	rv := objc.Send[foundation.URL](m_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_methods/method]: URL */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MediaDataStorage */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMediaDataStorage */


