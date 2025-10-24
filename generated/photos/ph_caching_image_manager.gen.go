// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHCachingImageManager] class.
var (
	PHCachingImageManagerClass     _PHCachingImageManagerClass
	PHCachingImageManagerClassOnce sync.Once
)

func getPHCachingImageManagerClass() _PHCachingImageManagerClass {
	PHCachingImageManagerClassOnce.Do(func() {
		PHCachingImageManagerClass = _PHCachingImageManagerClass{objc.GetClass("PHCachingImageManager")}
	})
	return PHCachingImageManagerClass
}

type _PHCachingImageManagerClass struct {
	class objc.Class
}

// An interface definition for the [PHCachingImageManager] class.
type IPHCachingImageManager interface {
	IPHImageManager
	// properties:
	AllowsCachingHighQualityImages() bool
	SetAllowsCachingHighQualityImages(value bool)
	// methods:
}

// An object that facilitates retrieving or generating preview thumbnails, optimized for batch preloading large numbers of assets.
//
// For quick performance when you are working with many assets, a caching image manager can prepare asset images in the background in order to eliminate delays when you later request individual images. For example, use a caching image manager when you want to populate a collection view or similar UI with thumbnails of photo or video assets. Much of the key functionality of the class is defined by its superclass, . For details, see . To use a caching image manager: Create a instance. (This step replaces using the shared instance.) Use class methods to fetch the assets you’re interested in. To prepare images for those assets, call the method with the target size, content mode, and options you plan to use when later requesting images for each individual asset. When you need an image for an individual asset, call the method, and pass the same parameters you used when preparing that asset. If the image you request is among those already prepared, the object immediately returns that image. Otherwise, Photos prepares the image on demand and caches it for later use.


// An object that facilitates retrieving or generating preview thumbnails, optimized for batch preloading large numbers of assets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCachingImageManager
type PHCachingImageManager struct {
	PHImageManager
}

// PHCachingImageManagerFrom constructs a [PHCachingImageManager] from an unsafe.Pointer.
//
// An object that facilitates retrieving or generating preview thumbnails, optimized for batch preloading large numbers of assets.
func PHCachingImageManagerFrom(ptr unsafe.Pointer) PHCachingImageManager {
	return PHCachingImageManager{
		PHImageManager: PHImageManagerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHCachingImageManagerClass) Alloc() PHCachingImageManager {
	rv := objc.Send[PHCachingImageManager](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHCachingImageManagerClass) New() PHCachingImageManager {
	rv := objc.Send[PHCachingImageManager](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHCachingImageManager) Init() PHCachingImageManager {
	rv := objc.Send[PHCachingImageManager](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHCachingImageManager) Autorelease() PHCachingImageManager {
	rv := objc.Send[PHCachingImageManager](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHCachingImageManager creates a new PHCachingImageManager instance.
func NewPHCachingImageManager() PHCachingImageManager {
	return getPHCachingImageManagerClass().New()
}



// A Boolean value that determines whether the image manager prepares high-quality images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcachingimagemanager/allowscachinghighqualityimages
func (p_ PHCachingImageManager) AllowsCachingHighQualityImages() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsCachingHighQualityImages"))
	return rv
}


// A Boolean value that determines whether the image manager prepares high-quality images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcachingimagemanager/allowscachinghighqualityimages
func (p_ PHCachingImageManager) SetAllowsCachingHighQualityImages(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowsCachingHighQualityImages:"), value)
}



