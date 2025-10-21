// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
	URL() unsafe.Pointer
}

// An object that represents the media sample data storage file.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MediaDataStorageClass) Alloc() MediaDataStorage {
	rv := objc.Send[MediaDataStorage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns the URL used to initialize the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaDataStorage/url()
func (m_ MediaDataStorage) URL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("URL"))
	return rv
}



