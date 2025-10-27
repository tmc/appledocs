// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [MediaDataStorage] class.
type IMediaDataStorage interface {
	objectivec.IObject
	

	// properties:


	

	// methods:
	URL() foundation.URL


}





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






// Creates a media data storage object associated with a file URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaDataStorage/init(url:options:)
func NewMediaDataStorageWithURLOptions(URL foundation.foundation.INSURL, options foundation.IDictionary) MediaDataStorage {
	instance := getMediaDataStorageClass().Alloc()
	rv := objc.Send[MediaDataStorage](instance.ID, objc.Sel("initWithURL:options:"), URL, options)
	rv.Autorelease()
	return rv
}

















// Returns the URL used to initialize the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaDataStorage/url()
func (m_ MediaDataStorage) URL() foundation.URL {
	rv := objc.Send[foundation.URL](m_.ID, objc.Sel("URL"))
	return rv
}












